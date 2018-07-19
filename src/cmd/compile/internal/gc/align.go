package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"sort"
)

// sizeCalculationDisabled indicates whether it is safe
// to calculate Types' widths and alignments. See dowidth.

// machine size and rounding alignment is dictated around
// the size of a pointer, set in betypeinit (see ../amd64/galign.go).

func (psess *PackageSession) Rnd(o int64, r int64) int64 {
	if r < 1 || r > 8 || r&(r-1) != 0 {
		psess.
			Fatalf("rnd %d", r)
	}
	return (o + r - 1) &^ (r - 1)
}

// expandiface computes the method set for interface type t by
// expanding embedded interfaces.
func (psess *PackageSession) expandiface(t *types.Type) {
	var fields []*types.Field
	for _, m := range t.Methods().Slice() {
		if m.Sym != nil {
			fields = append(fields, m)
			psess.
				checkwidth(m.Type)
			continue
		}

		if !m.Type.IsInterface() {
			psess.
				yyerrorl(m.Pos, "interface contains embedded non-interface %v", m.Type)
			m.SetBroke(true)
			t.SetBroke(true)

			fields = append(fields, m)
			continue
		}

		for _, t1 := range m.Type.Fields(psess.types).Slice() {
			f := types.NewField()
			f.Pos = m.Pos
			f.Sym = t1.Sym
			f.Type = t1.Type
			f.SetBroke(t1.Broke())
			fields = append(fields, f)
		}
	}
	sort.Sort(methcmp(fields))

	t.Extra.(*types.Interface).Fields.Set(fields)
}

func (psess *PackageSession) offmod(t *types.Type) {
	o := int32(0)
	for _, f := range t.Fields(psess.types).Slice() {
		f.Offset = int64(o)
		o += int32(psess.Widthptr)
		if int64(o) >= psess.thearch.MAXWIDTH {
			psess.
				yyerror("interface too large")
			o = int32(psess.Widthptr)
		}
	}
}

func (psess *PackageSession) widstruct(errtype *types.Type, t *types.Type, o int64, flag int) int64 {
	starto := o
	maxalign := int32(flag)
	if maxalign < 1 {
		maxalign = 1
	}
	lastzero := int64(0)
	for _, f := range t.Fields(psess.types).Slice() {
		if f.Type == nil {

			continue
		}
		psess.
			dowidth(f.Type)
		if int32(f.Type.Align) > maxalign {
			maxalign = int32(f.Type.Align)
		}
		if f.Type.Align > 0 {
			o = psess.Rnd(o, int64(f.Type.Align))
		}
		f.Offset = o
		if n := asNode(f.Nname); n != nil {

			if n.Name.Param.Stackcopy != nil {
				n.Name.Param.Stackcopy.Xoffset = o
				n.Xoffset = 0
			} else {
				n.Xoffset = o
			}
		}

		w := f.Type.Width
		if w < 0 {
			psess.
				Fatalf("invalid width %d", f.Type.Width)
		}
		if w == 0 {
			lastzero = o
		}
		o += w
		maxwidth := psess.thearch.MAXWIDTH

		if maxwidth < 1<<32 {
			maxwidth = 1<<31 - 1
		}
		if o >= maxwidth {
			psess.
				yyerror("type %L too large", errtype)
			o = 8
		}
	}

	if flag == 1 && o > starto && o == lastzero {
		o++
	}

	if flag != 0 {
		o = psess.Rnd(o, int64(maxalign))
	}
	t.Align = uint8(maxalign)

	t.Width = o - starto

	return o
}

// dowidth calculates and stores the size and alignment for t.
// If sizeCalculationDisabled is set, and the size/alignment
// have not already been calculated, it calls Fatal.
// This is used to prevent data races in the back end.
func (psess *PackageSession) dowidth(t *types.Type) {
	if psess.Widthptr == 0 {
		psess.
			Fatalf("dowidth without betypeinit")
	}

	if t == nil {
		return
	}

	if t.Width == -2 {
		if !t.Broke() {
			t.SetBroke(true)
			psess.
				yyerrorl(asNode(t.Nod).Pos, "invalid recursive type %v", t)
		}

		t.Width = 0
		t.Align = 1
		return
	}

	if t.WidthCalculated() {
		return
	}

	if psess.sizeCalculationDisabled {
		if t.Broke() {

			return
		}
		t.SetBroke(true)
		psess.
			Fatalf("width not calculated: %v", t)
	}

	if t.Broke() && t.Width == 0 {
		return
	}
	psess.
		defercalc++

	lno := psess.lineno
	if asNode(t.Nod) != nil {
		psess.
			lineno = asNode(t.Nod).Pos
	}

	t.Width = -2
	t.Align = 0

	et := t.Etype
	switch et {
	case TFUNC, TCHAN, TMAP, TSTRING:
		break

	default:
		if psess.simtype[t.Etype] != 0 {
			et = psess.simtype[t.Etype]
		}
	}

	w := int64(0)
	switch et {
	default:
		psess.
			Fatalf("dowidth: unknown type: %v", t)

	case TINT8, TUINT8, TBOOL:

		w = 1

	case TINT16, TUINT16:
		w = 2

	case TINT32, TUINT32, TFLOAT32:
		w = 4

	case TINT64, TUINT64, TFLOAT64:
		w = 8
		t.Align = uint8(psess.Widthreg)

	case TCOMPLEX64:
		w = 8
		t.Align = 4

	case TCOMPLEX128:
		w = 16
		t.Align = uint8(psess.Widthreg)

	case TPTR32:
		w = 4
		psess.
			checkwidth(t.Elem(psess.types))

	case TPTR64:
		w = 8
		psess.
			checkwidth(t.Elem(psess.types))

	case TUNSAFEPTR:
		w = int64(psess.Widthptr)

	case TINTER:
		w = 2 * int64(psess.Widthptr)
		t.Align = uint8(psess.Widthptr)
		psess.
			expandiface(t)

	case TCHAN:
		w = int64(psess.Widthptr)
		psess.
			checkwidth(t.Elem(psess.types))

		t1 := types.NewChanArgs(t)
		psess.
			checkwidth(t1)

	case TCHANARGS:
		t1 := t.ChanArgs(psess.types)
		psess.
			dowidth(t1)
		if t1.Elem(psess.types).Width >= 1<<16 {
			psess.
				yyerror("channel element type too large (>64kB)")
		}
		w = 1

	case TMAP:
		w = int64(psess.Widthptr)
		psess.
			checkwidth(t.Elem(psess.types))
		psess.
			checkwidth(t.Key(psess.types))

	case TFORW:
		if !t.Broke() {
			t.SetBroke(true)
			psess.
				yyerror("invalid recursive type %v", t)
		}
		w = 1

	case TANY:
		psess.
			Fatalf("dowidth any")

	case TSTRING:
		if psess.sizeof_String == 0 {
			psess.
				Fatalf("early dowidth string")
		}
		w = int64(psess.sizeof_String)
		t.Align = uint8(psess.Widthptr)

	case TARRAY:
		if t.Elem(psess.types) == nil {
			break
		}
		if t.IsDDDArray() {
			if !t.Broke() {
				psess.
					yyerror("use of [...] array outside of array literal")
				t.SetBroke(true)
			}
			break
		}
		psess.
			dowidth(t.Elem(psess.types))
		if t.Elem(psess.types).Width != 0 {
			cap := (uint64(psess.thearch.MAXWIDTH) - 1) / uint64(t.Elem(psess.types).Width)
			if uint64(t.NumElem(psess.types)) > cap {
				psess.
					yyerror("type %L larger than address space", t)
			}
		}
		w = t.NumElem(psess.types) * t.Elem(psess.types).Width
		t.Align = t.Elem(psess.types).Align

	case TSLICE:
		if t.Elem(psess.types) == nil {
			break
		}
		w = int64(psess.sizeof_Array)
		psess.
			checkwidth(t.Elem(psess.types))
		t.Align = uint8(psess.Widthptr)

	case TSTRUCT:
		if t.IsFuncArgStruct() {
			psess.
				Fatalf("dowidth fn struct %v", t)
		}
		w = psess.widstruct(t, t, 0, 1)

	case TFUNC:
		t1 := types.NewFuncArgs(t)
		psess.
			checkwidth(t1)
		w = int64(psess.Widthptr)

	case TFUNCARGS:
		t1 := t.FuncArgs(psess.types)
		w = psess.widstruct(t1, t1.Recvs(psess.types), 0, 0)
		w = psess.widstruct(t1, t1.Params(psess.types), w, psess.Widthreg)
		w = psess.widstruct(t1, t1.Results(psess.types), w, psess.Widthreg)
		t1.Extra.(*types.Func).Argwid = w
		if w%int64(psess.Widthreg) != 0 {
			psess.
				Warn("bad type %v %d\n", t1, w)
		}
		t.Align = 1
	}

	if psess.Widthptr == 4 && w != int64(int32(w)) {
		psess.
			yyerror("type %v too large", t)
	}

	t.Width = w
	if t.Align == 0 {
		if w > 8 || w&(w-1) != 0 || w == 0 {
			psess.
				Fatalf("invalid alignment for %v", t)
		}
		t.Align = uint8(w)
	}

	if t.Etype == TINTER {
		psess.
			checkdupfields("method", t)
		psess.
			offmod(t)
	}
	psess.
		lineno = lno

	if psess.defercalc == 1 {
		psess.
			resumecheckwidth()
	} else {
		psess.
			defercalc--
	}
}

func (psess *PackageSession) checkwidth(t *types.Type) {
	if t == nil {
		return
	}

	if t.IsFuncArgStruct() {
		psess.
			Fatalf("checkwidth %v", t)
	}

	if psess.defercalc == 0 {
		psess.
			dowidth(t)
		return
	}

	if t.Deferwidth() {
		return
	}
	t.SetDeferwidth(true)
	psess.
		deferredTypeStack = append(psess.deferredTypeStack, t)
}

func (psess *PackageSession) defercheckwidth() {

	if psess.defercalc != 0 && psess.nerrors == 0 {
		psess.
			Fatalf("defercheckwidth")
	}
	psess.
		defercalc = 1
}

func (psess *PackageSession) resumecheckwidth() {
	if psess.defercalc == 0 {
		psess.
			Fatalf("resumecheckwidth")
	}
	for len(psess.deferredTypeStack) > 0 {
		t := psess.deferredTypeStack[len(psess.deferredTypeStack)-1]
		psess.
			deferredTypeStack = psess.deferredTypeStack[:len(psess.deferredTypeStack)-1]
		t.SetDeferwidth(false)
		psess.
			dowidth(t)
	}
	psess.
		defercalc = 0
}
