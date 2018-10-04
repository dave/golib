// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"sort"
)

func (pstate *PackageState) Rnd(o int64, r int64) int64 {
	if r < 1 || r > 8 || r&(r-1) != 0 {
		pstate.Fatalf("rnd %d", r)
	}
	return (o + r - 1) &^ (r - 1)
}

// expandiface computes the method set for interface type t by
// expanding embedded interfaces.
func (pstate *PackageState) expandiface(t *types.Type) {
	var fields []*types.Field
	for _, m := range t.Methods().Slice() {
		if m.Sym != nil {
			fields = append(fields, m)
			pstate.checkwidth(m.Type)
			continue
		}

		if !m.Type.IsInterface() {
			pstate.yyerrorl(m.Pos, "interface contains embedded non-interface %v", m.Type)
			m.SetBroke(true)
			t.SetBroke(true)
			// Add to fields so that error messages
			// include the broken embedded type when
			// printing t.
			// TODO(mdempsky): Revisit this.
			fields = append(fields, m)
			continue
		}

		// Embedded interface: duplicate all methods
		// (including broken ones, if any) and add to t's
		// method set.
		for _, t1 := range m.Type.Fields(pstate.types).Slice() {
			f := types.NewField()
			f.Pos = m.Pos // preserve embedding position
			f.Sym = t1.Sym
			f.Type = t1.Type
			f.SetBroke(t1.Broke())
			fields = append(fields, f)
		}
	}
	sort.Sort(methcmp(fields))

	// Access fields directly to avoid recursively calling dowidth
	// within Type.Fields().
	t.Extra.(*types.Interface).Fields.Set(fields)
}

func (pstate *PackageState) offmod(t *types.Type) {
	o := int32(0)
	for _, f := range t.Fields(pstate.types).Slice() {
		f.Offset = int64(o)
		o += int32(pstate.Widthptr)
		if int64(o) >= pstate.thearch.MAXWIDTH {
			pstate.yyerror("interface too large")
			o = int32(pstate.Widthptr)
		}
	}
}

func (pstate *PackageState) widstruct(errtype *types.Type, t *types.Type, o int64, flag int) int64 {
	starto := o
	maxalign := int32(flag)
	if maxalign < 1 {
		maxalign = 1
	}
	lastzero := int64(0)
	for _, f := range t.Fields(pstate.types).Slice() {
		if f.Type == nil {
			// broken field, just skip it so that other valid fields
			// get a width.
			continue
		}

		pstate.dowidth(f.Type)
		if int32(f.Type.Align) > maxalign {
			maxalign = int32(f.Type.Align)
		}
		if f.Type.Align > 0 {
			o = pstate.Rnd(o, int64(f.Type.Align))
		}
		f.Offset = o
		if n := asNode(f.Nname); n != nil {
			// addrescapes has similar code to update these offsets.
			// Usually addrescapes runs after widstruct,
			// in which case we could drop this,
			// but function closure functions are the exception.
			// NOTE(rsc): This comment may be stale.
			// It's possible the ordering has changed and this is
			// now the common case. I'm not sure.
			if n.Name.Param.Stackcopy != nil {
				n.Name.Param.Stackcopy.Xoffset = o
				n.Xoffset = 0
			} else {
				n.Xoffset = o
			}
		}

		w := f.Type.Width
		if w < 0 {
			pstate.Fatalf("invalid width %d", f.Type.Width)
		}
		if w == 0 {
			lastzero = o
		}
		o += w
		maxwidth := pstate.thearch.MAXWIDTH
		// On 32-bit systems, reflect tables impose an additional constraint
		// that each field start offset must fit in 31 bits.
		if maxwidth < 1<<32 {
			maxwidth = 1<<31 - 1
		}
		if o >= maxwidth {
			pstate.yyerror("type %L too large", errtype)
			o = 8 // small but nonzero
		}
	}

	// For nonzero-sized structs which end in a zero-sized thing, we add
	// an extra byte of padding to the type. This padding ensures that
	// taking the address of the zero-sized thing can't manufacture a
	// pointer to the next object in the heap. See issue 9401.
	if flag == 1 && o > starto && o == lastzero {
		o++
	}

	// final width is rounded
	if flag != 0 {
		o = pstate.Rnd(o, int64(maxalign))
	}
	t.Align = uint8(maxalign)

	// type width only includes back to first field's offset
	t.Width = o - starto

	return o
}

// dowidth calculates and stores the size and alignment for t.
// If sizeCalculationDisabled is set, and the size/alignment
// have not already been calculated, it calls Fatal.
// This is used to prevent data races in the back end.
func (pstate *PackageState) dowidth(t *types.Type) {
	if pstate.Widthptr == 0 {
		pstate.Fatalf("dowidth without betypeinit")
	}

	if t == nil {
		return
	}

	if t.Width == -2 {
		if !t.Broke() {
			t.SetBroke(true)
			pstate.yyerrorl(asNode(t.Nod).Pos, "invalid recursive type %v", t)
		}

		t.Width = 0
		t.Align = 1
		return
	}

	if t.WidthCalculated() {
		return
	}

	if pstate.sizeCalculationDisabled {
		if t.Broke() {
			// break infinite recursion from Fatal call below
			return
		}
		t.SetBroke(true)
		pstate.Fatalf("width not calculated: %v", t)
	}

	// break infinite recursion if the broken recursive type
	// is referenced again
	if t.Broke() && t.Width == 0 {
		return
	}

	// defer checkwidth calls until after we're done
	pstate.defercalc++

	lno := pstate.lineno
	if asNode(t.Nod) != nil {
		pstate.lineno = asNode(t.Nod).Pos
	}

	t.Width = -2
	t.Align = 0

	et := t.Etype
	switch et {
	case TFUNC, TCHAN, TMAP, TSTRING:
		break

	// simtype == 0 during bootstrap
	default:
		if pstate.simtype[t.Etype] != 0 {
			et = pstate.simtype[t.Etype]
		}
	}

	w := int64(0)
	switch et {
	default:
		pstate.Fatalf("dowidth: unknown type: %v", t)

	// compiler-specific stuff
	case TINT8, TUINT8, TBOOL:
		// bool is int8
		w = 1

	case TINT16, TUINT16:
		w = 2

	case TINT32, TUINT32, TFLOAT32:
		w = 4

	case TINT64, TUINT64, TFLOAT64:
		w = 8
		t.Align = uint8(pstate.Widthreg)

	case TCOMPLEX64:
		w = 8
		t.Align = 4

	case TCOMPLEX128:
		w = 16
		t.Align = uint8(pstate.Widthreg)

	case TPTR32:
		w = 4
		pstate.checkwidth(t.Elem(pstate.types))

	case TPTR64:
		w = 8
		pstate.checkwidth(t.Elem(pstate.types))

	case TUNSAFEPTR:
		w = int64(pstate.Widthptr)

	case TINTER: // implemented as 2 pointers
		w = 2 * int64(pstate.Widthptr)
		t.Align = uint8(pstate.Widthptr)
		pstate.expandiface(t)

	case TCHAN: // implemented as pointer
		w = int64(pstate.Widthptr)

		pstate.checkwidth(t.Elem(pstate.types))

		// make fake type to check later to
		// trigger channel argument check.
		t1 := types.NewChanArgs(t)
		pstate.checkwidth(t1)

	case TCHANARGS:
		t1 := t.ChanArgs(pstate.types)
		pstate.dowidth(t1) // just in case
		if t1.Elem(pstate.types).Width >= 1<<16 {
			pstate.yyerror("channel element type too large (>64kB)")
		}
		w = 1 // anything will do

	case TMAP: // implemented as pointer
		w = int64(pstate.Widthptr)
		pstate.checkwidth(t.Elem(pstate.types))
		pstate.checkwidth(t.Key(pstate.types))

	case TFORW: // should have been filled in
		if !t.Broke() {
			t.SetBroke(true)
			pstate.yyerror("invalid recursive type %v", t)
		}
		w = 1 // anything will do

	case TANY:
		// dummy type; should be replaced before use.
		pstate.Fatalf("dowidth any")

	case TSTRING:
		if pstate.sizeof_String == 0 {
			pstate.Fatalf("early dowidth string")
		}
		w = int64(pstate.sizeof_String)
		t.Align = uint8(pstate.Widthptr)

	case TARRAY:
		if t.Elem(pstate.types) == nil {
			break
		}
		if t.IsDDDArray() {
			if !t.Broke() {
				pstate.yyerror("use of [...] array outside of array literal")
				t.SetBroke(true)
			}
			break
		}

		pstate.dowidth(t.Elem(pstate.types))
		if t.Elem(pstate.types).Width != 0 {
			cap := (uint64(pstate.thearch.MAXWIDTH) - 1) / uint64(t.Elem(pstate.types).Width)
			if uint64(t.NumElem(pstate.types)) > cap {
				pstate.yyerror("type %L larger than address space", t)
			}
		}
		w = t.NumElem(pstate.types) * t.Elem(pstate.types).Width
		t.Align = t.Elem(pstate.types).Align

	case TSLICE:
		if t.Elem(pstate.types) == nil {
			break
		}
		w = int64(pstate.sizeof_Array)
		pstate.checkwidth(t.Elem(pstate.types))
		t.Align = uint8(pstate.Widthptr)

	case TSTRUCT:
		if t.IsFuncArgStruct() {
			pstate.Fatalf("dowidth fn struct %v", t)
		}
		w = pstate.widstruct(t, t, 0, 1)

	// make fake type to check later to
	// trigger function argument computation.
	case TFUNC:
		t1 := types.NewFuncArgs(t)
		pstate.checkwidth(t1)
		w = int64(pstate.Widthptr) // width of func type is pointer

	// function is 3 cated structures;
	// compute their widths as side-effect.
	case TFUNCARGS:
		t1 := t.FuncArgs(pstate.types)
		w = pstate.widstruct(t1, t1.Recvs(pstate.types), 0, 0)
		w = pstate.widstruct(t1, t1.Params(pstate.types), w, pstate.Widthreg)
		w = pstate.widstruct(t1, t1.Results(pstate.types), w, pstate.Widthreg)
		t1.Extra.(*types.Func).Argwid = w
		if w%int64(pstate.Widthreg) != 0 {
			pstate.Warn("bad type %v %d\n", t1, w)
		}
		t.Align = 1
	}

	if pstate.Widthptr == 4 && w != int64(int32(w)) {
		pstate.yyerror("type %v too large", t)
	}

	t.Width = w
	if t.Align == 0 {
		if w > 8 || w&(w-1) != 0 || w == 0 {
			pstate.Fatalf("invalid alignment for %v", t)
		}
		t.Align = uint8(w)
	}

	if t.Etype == TINTER {
		// We defer calling these functions until after
		// setting t.Width and t.Align so the recursive calls
		// to dowidth within t.Fields() will succeed.
		pstate.checkdupfields("method", t)
		pstate.offmod(t)
	}

	pstate.lineno = lno

	if pstate.defercalc == 1 {
		pstate.resumecheckwidth()
	} else {
		pstate.defercalc--
	}
}

func (pstate *PackageState) checkwidth(t *types.Type) {
	if t == nil {
		return
	}

	// function arg structs should not be checked
	// outside of the enclosing function.
	if t.IsFuncArgStruct() {
		pstate.Fatalf("checkwidth %v", t)
	}

	if pstate.defercalc == 0 {
		pstate.dowidth(t)
		return
	}

	if t.Deferwidth() {
		return
	}
	t.SetDeferwidth(true)

	pstate.deferredTypeStack = append(pstate.deferredTypeStack, t)
}

func (pstate *PackageState) defercheckwidth() {
	// we get out of sync on syntax errors, so don't be pedantic.
	if pstate.defercalc != 0 && pstate.nerrors == 0 {
		pstate.Fatalf("defercheckwidth")
	}
	pstate.defercalc = 1
}

func (pstate *PackageState) resumecheckwidth() {
	if pstate.defercalc == 0 {
		pstate.Fatalf("resumecheckwidth")
	}
	for len(pstate.deferredTypeStack) > 0 {
		t := pstate.deferredTypeStack[len(pstate.deferredTypeStack)-1]
		pstate.deferredTypeStack = pstate.deferredTypeStack[:len(pstate.deferredTypeStack)-1]
		t.SetDeferwidth(false)
		pstate.dowidth(t)
	}

	pstate.defercalc = 0
}
