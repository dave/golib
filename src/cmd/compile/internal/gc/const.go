package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"math/big"
	"strings"
)

// Ctype describes the constant kind of an "ideal" (untyped) constant.
type Ctype uint8

const (
	CTxxx Ctype = iota

	CTINT
	CTRUNE
	CTFLT
	CTCPLX
	CTSTR
	CTBOOL
	CTNIL
)

type Val struct {
	// U contains one of:
	// bool     bool when n.ValCtype() == CTBOOL
	// *Mpint   int when n.ValCtype() == CTINT, rune when n.ValCtype() == CTRUNE
	// *Mpflt   float when n.ValCtype() == CTFLT
	// *Mpcplx  pair of floats when n.ValCtype() == CTCPLX
	// string   string when n.ValCtype() == CTSTR
	// *Nilval  when n.ValCtype() == CTNIL
	U interface{}
}

func (v Val) Ctype(psess *PackageSession) Ctype {
	switch x := v.U.(type) {
	default:
		psess.
			Fatalf("unexpected Ctype for %T", v.U)
		panic("not reached")
	case nil:
		return 0
	case *NilVal:
		return CTNIL
	case bool:
		return CTBOOL
	case *Mpint:
		if x.Rune {
			return CTRUNE
		}
		return CTINT
	case *Mpflt:
		return CTFLT
	case *Mpcplx:
		return CTCPLX
	case string:
		return CTSTR
	}
}

func (psess *PackageSession) eqval(a, b Val) bool {
	if a.Ctype(psess) != b.Ctype(psess) {
		return false
	}
	switch x := a.U.(type) {
	default:
		psess.
			Fatalf("unexpected Ctype for %T", a.U)
		panic("not reached")
	case *NilVal:
		return true
	case bool:
		y := b.U.(bool)
		return x == y
	case *Mpint:
		y := b.U.(*Mpint)
		return x.Cmp(y) == 0
	case *Mpflt:
		y := b.U.(*Mpflt)
		return x.Cmp(y) == 0
	case *Mpcplx:
		y := b.U.(*Mpcplx)
		return x.Real.Cmp(&y.Real) == 0 && x.Imag.Cmp(&y.Imag) == 0
	case string:
		y := b.U.(string)
		return x == y
	}
}

// Interface returns the constant value stored in v as an interface{}.
// It returns int64s for ints and runes, float64s for floats,
// complex128s for complex values, and nil for constant nils.
func (v Val) Interface(psess *PackageSession) interface{} {
	switch x := v.U.(type) {
	default:
		psess.
			Fatalf("unexpected Interface for %T", v.U)
		panic("not reached")
	case *NilVal:
		return nil
	case bool, string:
		return x
	case *Mpint:
		return x.Int64(psess)
	case *Mpflt:
		return x.Float64(psess)
	case *Mpcplx:
		return complex(x.Real.Float64(psess), x.Imag.Float64(psess))
	}
}

type NilVal struct{}

// Int64 returns n as an int64.
// n must be an integer or rune constant.
func (n *Node) Int64(psess *PackageSession) int64 {
	if !psess.Isconst(n, CTINT) {
		psess.
			Fatalf("Int64(%v)", n)
	}
	return n.Val().U.(*Mpint).Int64(psess)
}

// Bool returns n as a bool.
// n must be a boolean constant.
func (n *Node) Bool(psess *PackageSession) bool {
	if !psess.Isconst(n, CTBOOL) {
		psess.
			Fatalf("Bool(%v)", n)
	}
	return n.Val().U.(bool)
}

// truncate float literal fv to 32-bit or 64-bit precision
// according to type; return truncated value.
func (psess *PackageSession) truncfltlit(oldv *Mpflt, t *types.Type) *Mpflt {
	if t == nil {
		return oldv
	}

	if psess.overflow(Val{oldv}, t) {

		return oldv
	}

	fv := newMpflt()

	switch t.Etype {
	case types.TFLOAT32:
		fv.SetFloat64(oldv.Float32(psess))
	case types.TFLOAT64:
		fv.SetFloat64(oldv.Float64(psess))
	default:
		psess.
			Fatalf("truncfltlit: unexpected Etype %v", t.Etype)
	}

	return fv
}

// truncate Real and Imag parts of Mpcplx to 32-bit or 64-bit
// precision, according to type; return truncated value. In case of
// overflow, calls yyerror but does not truncate the input value.
func (psess *PackageSession) trunccmplxlit(oldv *Mpcplx, t *types.Type) *Mpcplx {
	if t == nil {
		return oldv
	}

	if psess.overflow(Val{oldv}, t) {

		return oldv
	}

	cv := newMpcmplx()

	switch t.Etype {
	case types.TCOMPLEX64:
		cv.Real.SetFloat64(oldv.Real.Float32(psess))
		cv.Imag.SetFloat64(oldv.Imag.Float32(psess))
	case types.TCOMPLEX128:
		cv.Real.SetFloat64(oldv.Real.Float64(psess))
		cv.Imag.SetFloat64(oldv.Imag.Float64(psess))
	default:
		psess.
			Fatalf("trunccplxlit: unexpected Etype %v", t.Etype)
	}

	return cv
}

// canReuseNode indicates whether it is known to be safe
// to reuse a Node.
type canReuseNode bool

const (
	noReuse canReuseNode = false // not necessarily safe to reuse
	reuseOK canReuseNode = true  // safe to reuse
)

// convert n, if literal, to type t.
// implicit conversion.
// The result of convlit MUST be assigned back to n, e.g.
// 	n.Left = convlit(n.Left, t)
func (psess *PackageSession) convlit(n *Node, t *types.Type) *Node {
	return psess.convlit1(n, t, false, noReuse)
}

// convlit1 converts n, if literal, to type t.
// It returns a new node if necessary.
// The result of convlit1 MUST be assigned back to n, e.g.
// 	n.Left = convlit1(n.Left, t, explicit, reuse)
func (psess *PackageSession) convlit1(n *Node, t *types.Type, explicit bool, reuse canReuseNode) *Node {
	if n == nil || t == nil || n.Type == nil || t.IsUntyped(psess.types) || n.Type == t {
		return n
	}
	if !explicit && !n.Type.IsUntyped(psess.types) {
		return n
	}

	if n.Op == OLITERAL && !reuse {

		n = n.copy()
		reuse = true
	}

	switch n.Op {
	default:
		if n.Type == psess.types.Idealbool {
			if !t.IsBoolean() {
				t = psess.types.Types[TBOOL]
			}
			switch n.Op {
			case ONOT:
				n.Left = psess.convlit(n.Left, t)
			case OANDAND, OOROR:
				n.Left = psess.convlit(n.Left, t)
				n.Right = psess.convlit(n.Right, t)
			}
			n.Type = t
		}

		if n.Type.Etype == TIDEAL {
			n.Left = psess.convlit(n.Left, t)
			n.Right = psess.convlit(n.Right, t)
			n.Type = t
		}

		return n

	case OLITERAL:
		if !psess.okforconst[t.Etype] && n.Type.Etype != TNIL {
			return psess.defaultlitreuse(n, nil, reuse)
		}

	case OLSH, ORSH:
		n.Left = psess.convlit1(n.Left, t, explicit && n.Left.Type.IsUntyped(psess.types), noReuse)
		t = n.Left.Type
		if t != nil && t.Etype == TIDEAL && n.Val().Ctype(psess) != CTINT {
			n.SetVal(psess, psess.toint(n.Val()))
		}
		if t != nil && !t.IsInteger() {
			psess.
				yyerror("invalid operation: %v (shift of type %v)", n, t)
			t = nil
		}

		n.Type = t
		return n

	case OCOMPLEX:
		if n.Type.Etype == TIDEAL {
			switch t.Etype {
			default:

				t = psess.types.Types[TCOMPLEX128]
				fallthrough
			case types.TCOMPLEX128:
				n.Type = t
				n.Left = psess.convlit(n.Left, psess.types.Types[TFLOAT64])
				n.Right = psess.convlit(n.Right, psess.types.Types[TFLOAT64])

			case TCOMPLEX64:
				n.Type = t
				n.Left = psess.convlit(n.Left, psess.types.Types[TFLOAT32])
				n.Right = psess.convlit(n.Right, psess.types.Types[TFLOAT32])
			}
		}

		return n
	}

	if psess.eqtype(n.Type, t) {
		return n
	}

	ct := psess.consttype(n)
	var et types.EType
	if ct == 0 {
		goto bad
	}

	et = t.Etype
	if et == TINTER {
		if ct == CTNIL && n.Type == psess.types.Types[TNIL] {
			n.Type = t
			return n
		}
		return psess.defaultlitreuse(n, nil, reuse)
	}

	switch ct {
	default:
		goto bad

	case CTNIL:
		switch et {
		default:
			n.Type = nil
			goto bad

		case TSTRING:
			return n

		case TARRAY:
			goto bad

		case TPTR32, TPTR64, TUNSAFEPTR:
			n.SetVal(psess, Val{new(Mpint)})

		case TCHAN, TFUNC, TINTER, TMAP, TSLICE:
			break
		}

	case CTSTR, CTBOOL:
		if et != n.Type.Etype {
			goto bad
		}

	case CTINT, CTRUNE, CTFLT, CTCPLX:
		if n.Type.Etype == TUNSAFEPTR && t.Etype != TUINTPTR {
			goto bad
		}
		ct := n.Val().Ctype(psess)
		if psess.isInt[et] {
			switch ct {
			default:
				goto bad

			case CTCPLX, CTFLT, CTRUNE:
				n.SetVal(psess, psess.toint(n.Val()))
				fallthrough

			case CTINT:
				psess.
					overflow(n.Val(), t)
			}
		} else if psess.isFloat[et] {
			switch ct {
			default:
				goto bad

			case CTCPLX, CTINT, CTRUNE:
				n.SetVal(psess, psess.toflt(n.Val()))
				fallthrough

			case CTFLT:
				n.SetVal(psess, Val{psess.truncfltlit(n.Val().U.(*Mpflt), t)})
			}
		} else if psess.isComplex[et] {
			switch ct {
			default:
				goto bad

			case CTFLT, CTINT, CTRUNE:
				n.SetVal(psess, tocplx(n.Val()))
				fallthrough

			case CTCPLX:
				n.SetVal(psess, Val{psess.trunccmplxlit(n.Val().U.(*Mpcplx), t)})
			}
		} else if et == types.TSTRING && (ct == CTINT || ct == CTRUNE) && explicit {
			n.SetVal(psess, psess.tostr(n.Val()))
		} else {
			goto bad
		}
	}

	n.Type = t
	return n

bad:
	if !n.Diag() {
		if !t.Broke() {
			psess.
				yyerror("cannot convert %L to type %v", n, t)
		}
		n.SetDiag(true)
	}

	if n.Type.IsUntyped(psess.types) {
		n = psess.defaultlitreuse(n, nil, reuse)
	}
	return n
}

func copyval(v Val) Val {
	switch u := v.U.(type) {
	case *Mpint:
		i := new(Mpint)
		i.Set(u)
		i.Rune = u.Rune
		v.U = i

	case *Mpflt:
		f := newMpflt()
		f.Set(u)
		v.U = f

	case *Mpcplx:
		c := new(Mpcplx)
		c.Real.Set(&u.Real)
		c.Imag.Set(&u.Imag)
		v.U = c
	}

	return v
}

func tocplx(v Val) Val {
	switch u := v.U.(type) {
	case *Mpint:
		c := new(Mpcplx)
		c.Real.SetInt(u)
		c.Imag.SetFloat64(0.0)
		v.U = c

	case *Mpflt:
		c := new(Mpcplx)
		c.Real.Set(u)
		c.Imag.SetFloat64(0.0)
		v.U = c
	}

	return v
}

func (psess *PackageSession) toflt(v Val) Val {
	switch u := v.U.(type) {
	case *Mpint:
		f := newMpflt()
		f.SetInt(u)
		v.U = f

	case *Mpcplx:
		f := newMpflt()
		f.Set(&u.Real)
		if u.Imag.CmpFloat64(0) != 0 {
			psess.
				yyerror("constant %v%vi truncated to real", fconv(&u.Real, FmtSharp), fconv(&u.Imag, FmtSharp|FmtSign))
		}
		v.U = f
	}

	return v
}

func (psess *PackageSession) toint(v Val) Val {
	switch u := v.U.(type) {
	case *Mpint:
		if u.Rune {
			i := new(Mpint)
			i.Set(u)
			v.U = i
		}

	case *Mpflt:
		i := new(Mpint)
		if !i.SetFloat(u) {
			if i.checkOverflow(0) {
				psess.
					yyerror("integer too large")
			} else {
				// The value of u cannot be represented as an integer;
				// so we need to print an error message.
				// Unfortunately some float values cannot be
				// reasonably formatted for inclusion in an error
				// message (example: 1 + 1e-100), so first we try to
				// format the float; if the truncation resulted in
				// something that looks like an integer we omit the
				// value from the error message.
				// (See issue #11371).
				var t big.Float
				t.Parse(fconv(u, FmtSharp), 10)
				if t.IsInt() {
					psess.
						yyerror("constant truncated to integer")
				} else {
					psess.
						yyerror("constant %v truncated to integer", fconv(u, FmtSharp))
				}
			}
		}
		v.U = i

	case *Mpcplx:
		i := new(Mpint)
		if !i.SetFloat(&u.Real) || u.Imag.CmpFloat64(0) != 0 {
			psess.
				yyerror("constant %v%vi truncated to integer", fconv(&u.Real, FmtSharp), fconv(&u.Imag, FmtSharp|FmtSign))
		}

		v.U = i
	}

	return v
}

func (psess *PackageSession) doesoverflow(v Val, t *types.Type) bool {
	switch u := v.U.(type) {
	case *Mpint:
		if !t.IsInteger() {
			psess.
				Fatalf("overflow: %v integer constant", t)
		}
		return u.Cmp(psess.minintval[t.Etype]) < 0 || u.Cmp(psess.maxintval[t.Etype]) > 0

	case *Mpflt:
		if !t.IsFloat() {
			psess.
				Fatalf("overflow: %v floating-point constant", t)
		}
		return u.Cmp(psess.minfltval[t.Etype]) <= 0 || u.Cmp(psess.maxfltval[t.Etype]) >= 0

	case *Mpcplx:
		if !t.IsComplex() {
			psess.
				Fatalf("overflow: %v complex constant", t)
		}
		return u.Real.Cmp(psess.minfltval[t.Etype]) <= 0 || u.Real.Cmp(psess.maxfltval[t.Etype]) >= 0 ||
			u.Imag.Cmp(psess.minfltval[t.Etype]) <= 0 || u.Imag.Cmp(psess.maxfltval[t.Etype]) >= 0
	}

	return false
}

func (psess *PackageSession) overflow(v Val, t *types.Type) bool {

	if t == nil || t.Etype == TIDEAL {
		return false
	}

	if t.IsPtr() || t.IsUnsafePtr() {
		return false
	}

	if psess.doesoverflow(v, t) {
		psess.
			yyerror("constant %v overflows %v", v, t)
		return true
	}

	return false

}

func (psess *PackageSession) tostr(v Val) Val {
	switch u := v.U.(type) {
	case *Mpint:
		var i int64 = 0xFFFD
		if u.Cmp(psess.minintval[TUINT32]) >= 0 && u.Cmp(psess.maxintval[TUINT32]) <= 0 {
			i = u.Int64(psess)
		}
		v.U = string(i)

	case *NilVal:

		v.U = ""
	}

	return v
}

func (psess *PackageSession) consttype(n *Node) Ctype {
	if n == nil || n.Op != OLITERAL {
		return 0
	}
	return n.Val().Ctype(psess)
}

func (psess *PackageSession) Isconst(n *Node, ct Ctype) bool {
	t := psess.consttype(n)

	return t == ct || (ct == CTINT && t == CTRUNE)
}

// if n is constant, rewrite as OLITERAL node.
func (psess *PackageSession) evconst(n *Node) {

	switch n.Op {
	default:
		return

	case OADD,
		OAND,
		OANDAND,
		OANDNOT,
		OARRAYBYTESTR,
		OCOM,
		ODIV,
		OEQ,
		OGE,
		OGT,
		OLE,
		OLSH,
		OLT,
		OMINUS,
		OMOD,
		OMUL,
		ONE,
		ONOT,
		OOR,
		OOROR,
		OPLUS,
		ORSH,
		OSUB,
		OXOR:
		break

	case OCONV:
		if n.Type == nil {
			return
		}
		if !psess.okforconst[n.Type.Etype] && n.Type.Etype != TNIL {
			return
		}

	case OADDSTR:
		s := n.List.Slice()
		for i1 := 0; i1 < len(s); i1++ {
			if psess.Isconst(s[i1], CTSTR) && i1+1 < len(s) && psess.Isconst(s[i1+1], CTSTR) {
				// merge from i1 up to but not including i2
				var strs []string
				i2 := i1
				for i2 < len(s) && psess.Isconst(s[i2], CTSTR) {
					strs = append(strs, s[i2].Val().U.(string))
					i2++
				}

				nl := *s[i1]
				nl.Orig = &nl
				nl.SetVal(psess, Val{strings.Join(strs, "")})
				s[i1] = &nl
				s = append(s[:i1+1], s[i2:]...)
			}
		}

		if len(s) == 1 && psess.Isconst(s[0], CTSTR) {
			n.Op = OLITERAL
			n.SetVal(psess, s[0].Val())
		} else {
			n.List.Set(s)
		}

		return
	}

	nl := n.Left
	if nl == nil || nl.Type == nil {
		return
	}
	if psess.consttype(nl) == 0 {
		return
	}
	wl := nl.Type.Etype
	if psess.isInt[wl] || psess.isFloat[wl] || psess.isComplex[wl] {
		wl = TIDEAL
	}

	// avoid constant conversions in switches below
	const (
		CTINT_         = uint32(CTINT)
		CTRUNE_        = uint32(CTRUNE)
		CTFLT_         = uint32(CTFLT)
		CTCPLX_        = uint32(CTCPLX)
		CTSTR_         = uint32(CTSTR)
		CTBOOL_        = uint32(CTBOOL)
		CTNIL_         = uint32(CTNIL)
		OCONV_         = uint32(OCONV) << 16
		OARRAYBYTESTR_ = uint32(OARRAYBYTESTR) << 16
		OPLUS_         = uint32(OPLUS) << 16
		OMINUS_        = uint32(OMINUS) << 16
		OCOM_          = uint32(OCOM) << 16
		ONOT_          = uint32(ONOT) << 16
		OLSH_          = uint32(OLSH) << 16
		ORSH_          = uint32(ORSH) << 16
		OADD_          = uint32(OADD) << 16
		OSUB_          = uint32(OSUB) << 16
		OMUL_          = uint32(OMUL) << 16
		ODIV_          = uint32(ODIV) << 16
		OMOD_          = uint32(OMOD) << 16
		OOR_           = uint32(OOR) << 16
		OAND_          = uint32(OAND) << 16
		OANDNOT_       = uint32(OANDNOT) << 16
		OXOR_          = uint32(OXOR) << 16
		OEQ_           = uint32(OEQ) << 16
		ONE_           = uint32(ONE) << 16
		OLT_           = uint32(OLT) << 16
		OLE_           = uint32(OLE) << 16
		OGE_           = uint32(OGE) << 16
		OGT_           = uint32(OGT) << 16
		OOROR_         = uint32(OOROR) << 16
		OANDAND_       = uint32(OANDAND) << 16
	)

	nr := n.Right
	var rv Val
	var wr types.EType
	var ctype uint32
	var v Val
	if nr == nil {

		v = copyval(nl.Val())

		ctype = uint32(v.Ctype(psess))
		if ctype == CTRUNE_ {
			ctype = CTINT_
		}

		switch uint32(n.Op)<<16 | ctype {
		default:
			if !n.Diag() {
				psess.
					yyerror("illegal constant expression %v %v", n.Op, nl.Type)
				n.SetDiag(true)
			}
			return

		case OCONV_ | CTNIL_,
			OARRAYBYTESTR_ | CTNIL_:
			if n.Type.IsString() {
				v = psess.tostr(v)
				nl.Type = n.Type
				break
			}
			fallthrough
		case OCONV_ | CTINT_,
			OCONV_ | CTFLT_,
			OCONV_ | CTCPLX_,
			OCONV_ | CTSTR_,
			OCONV_ | CTBOOL_:
			nl = psess.convlit1(nl, n.Type, true, false)
			v = nl.Val()

		case OPLUS_ | CTINT_:
			break

		case OMINUS_ | CTINT_:
			v.U.(*Mpint).Neg()

		case OCOM_ | CTINT_:
			var et types.EType = Txxx
			if nl.Type != nil {
				et = nl.Type.Etype
			}

			// calculate the mask in b
			// result will be (a ^ mask)
			var b Mpint
			switch et {

			default:
				b.SetInt64(-1)

			case TUINT8,
				TUINT16,
				TUINT32,
				TUINT64,
				TUINT,
				TUINTPTR:
				b.Set(psess.maxintval[et])
			}

			v.U.(*Mpint).Xor(psess, &b)

		case OPLUS_ | CTFLT_:
			break

		case OMINUS_ | CTFLT_:
			v.U.(*Mpflt).Neg()

		case OPLUS_ | CTCPLX_:
			break

		case OMINUS_ | CTCPLX_:
			v.U.(*Mpcplx).Real.Neg()
			v.U.(*Mpcplx).Imag.Neg()

		case ONOT_ | CTBOOL_:
			if !v.U.(bool) {
				goto settrue
			}
			goto setfalse
		}
		goto ret
	}
	if nr.Type == nil {
		return
	}
	if psess.consttype(nr) == 0 {
		return
	}
	wr = nr.Type.Etype
	if psess.isInt[wr] || psess.isFloat[wr] || psess.isComplex[wr] {
		wr = TIDEAL
	}

	if wl != wr {
		if wl == TINTER || wr == TINTER {
			if n.Op == ONE {
				goto settrue
			}
			goto setfalse
		}
		goto illegal
	}

	switch n.Op {

	default:
		if nl.Type.Etype != TIDEAL {
			nr = psess.defaultlit(nr, nl.Type)
			n.Right = nr
		}

		if nr.Type.Etype != TIDEAL {
			nl = psess.defaultlit(nl, nr.Type)
			n.Left = nl
		}

		if nl.Type.Etype != nr.Type.Etype {
			goto illegal
		}

	case OLSH, ORSH:
		nr = psess.defaultlit(nr, psess.types.Types[TUINT])

		n.Right = nr
		if nr.Type != nil && (nr.Type.IsSigned() || !nr.Type.IsInteger()) {
			goto illegal
		}
		if nl.Val().Ctype(psess) != CTRUNE {
			nl.SetVal(psess, psess.toint(nl.Val()))
		}
		nr.SetVal(psess, psess.toint(nr.Val()))
	}

	v = copyval(nl.Val())
	rv = nr.Val()

	if v.Ctype(psess) == CTCPLX || rv.Ctype(psess) == CTCPLX {
		v = tocplx(v)
		rv = tocplx(rv)
	}

	if v.Ctype(psess) == CTFLT || rv.Ctype(psess) == CTFLT {
		v = psess.toflt(v)
		rv = psess.toflt(rv)
	}

	if v.Ctype(psess) == CTRUNE && rv.Ctype(psess) == CTINT {
		i := new(Mpint)
		i.Set(rv.U.(*Mpint))
		i.Rune = true
		rv.U = i
	}
	if v.Ctype(psess) == CTINT && rv.Ctype(psess) == CTRUNE {
		if n.Op == OLSH || n.Op == ORSH {
			i := new(Mpint)
			i.Set(rv.U.(*Mpint))
			rv.U = i
		} else {
			i := new(Mpint)
			i.Set(v.U.(*Mpint))
			i.Rune = true
			v.U = i
		}
	}

	if v.Ctype(psess) != rv.Ctype(psess) {

		if (v.Ctype(psess) == 0 || rv.Ctype(psess) == 0) && psess.nerrors > 0 {
			return
		}
		psess.
			Fatalf("constant type mismatch %v(%d) %v(%d)", nl.Type, v.Ctype(psess), nr.Type, rv.Ctype(psess))
	}

	ctype = uint32(v.Ctype(psess))
	if ctype == CTRUNE_ {
		ctype = CTINT_
	}

	switch uint32(n.Op)<<16 | ctype {
	default:
		goto illegal

	case OADD_ | CTINT_:
		v.U.(*Mpint).Add(psess, rv.U.(*Mpint))

	case OSUB_ | CTINT_:
		v.U.(*Mpint).Sub(psess, rv.U.(*Mpint))

	case OMUL_ | CTINT_:
		v.U.(*Mpint).Mul(psess, rv.U.(*Mpint))

	case ODIV_ | CTINT_:
		if rv.U.(*Mpint).CmpInt64(0) == 0 {
			psess.
				yyerror("division by zero")
			v.U.(*Mpint).SetOverflow()
			break
		}

		v.U.(*Mpint).Quo(psess, rv.U.(*Mpint))

	case OMOD_ | CTINT_:
		if rv.U.(*Mpint).CmpInt64(0) == 0 {
			psess.
				yyerror("division by zero")
			v.U.(*Mpint).SetOverflow()
			break
		}

		v.U.(*Mpint).Rem(psess, rv.U.(*Mpint))

	case OLSH_ | CTINT_:
		v.U.(*Mpint).Lsh(psess, rv.U.(*Mpint))

	case ORSH_ | CTINT_:
		v.U.(*Mpint).Rsh(psess, rv.U.(*Mpint))

	case OOR_ | CTINT_:
		v.U.(*Mpint).Or(psess, rv.U.(*Mpint))

	case OAND_ | CTINT_:
		v.U.(*Mpint).And(psess, rv.U.(*Mpint))

	case OANDNOT_ | CTINT_:
		v.U.(*Mpint).AndNot(psess, rv.U.(*Mpint))

	case OXOR_ | CTINT_:
		v.U.(*Mpint).Xor(psess, rv.U.(*Mpint))

	case OADD_ | CTFLT_:
		v.U.(*Mpflt).Add(rv.U.(*Mpflt))

	case OSUB_ | CTFLT_:
		v.U.(*Mpflt).Sub(rv.U.(*Mpflt))

	case OMUL_ | CTFLT_:
		v.U.(*Mpflt).Mul(rv.U.(*Mpflt))

	case ODIV_ | CTFLT_:
		if rv.U.(*Mpflt).CmpFloat64(0) == 0 {
			psess.
				yyerror("division by zero")
			v.U.(*Mpflt).SetFloat64(1.0)
			break
		}

		v.U.(*Mpflt).Quo(rv.U.(*Mpflt))

	case OMOD_ | CTFLT_:
		if !n.Diag() {
			psess.
				yyerror("illegal constant expression: floating-point %% operation")
			n.SetDiag(true)
		}

		return

	case OADD_ | CTCPLX_:
		v.U.(*Mpcplx).Real.Add(&rv.U.(*Mpcplx).Real)
		v.U.(*Mpcplx).Imag.Add(&rv.U.(*Mpcplx).Imag)

	case OSUB_ | CTCPLX_:
		v.U.(*Mpcplx).Real.Sub(&rv.U.(*Mpcplx).Real)
		v.U.(*Mpcplx).Imag.Sub(&rv.U.(*Mpcplx).Imag)

	case OMUL_ | CTCPLX_:
		v.U.(*Mpcplx).Mul(rv.U.(*Mpcplx))

	case ODIV_ | CTCPLX_:
		if !v.U.(*Mpcplx).Div(rv.U.(*Mpcplx)) {
			psess.
				yyerror("complex division by zero")
			rv.U.(*Mpcplx).Real.SetFloat64(1.0)
			rv.U.(*Mpcplx).Imag.SetFloat64(0.0)
			break
		}

	case OEQ_ | CTNIL_:
		goto settrue

	case ONE_ | CTNIL_:
		goto setfalse

	case OEQ_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) == 0 {
			goto settrue
		}
		goto setfalse

	case ONE_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) != 0 {
			goto settrue
		}
		goto setfalse

	case OLT_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) < 0 {
			goto settrue
		}
		goto setfalse

	case OLE_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) <= 0 {
			goto settrue
		}
		goto setfalse

	case OGE_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) >= 0 {
			goto settrue
		}
		goto setfalse

	case OGT_ | CTINT_:
		if v.U.(*Mpint).Cmp(rv.U.(*Mpint)) > 0 {
			goto settrue
		}
		goto setfalse

	case OEQ_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) == 0 {
			goto settrue
		}
		goto setfalse

	case ONE_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) != 0 {
			goto settrue
		}
		goto setfalse

	case OLT_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) < 0 {
			goto settrue
		}
		goto setfalse

	case OLE_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) <= 0 {
			goto settrue
		}
		goto setfalse

	case OGE_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) >= 0 {
			goto settrue
		}
		goto setfalse

	case OGT_ | CTFLT_:
		if v.U.(*Mpflt).Cmp(rv.U.(*Mpflt)) > 0 {
			goto settrue
		}
		goto setfalse

	case OEQ_ | CTCPLX_:
		if v.U.(*Mpcplx).Real.Cmp(&rv.U.(*Mpcplx).Real) == 0 && v.U.(*Mpcplx).Imag.Cmp(&rv.U.(*Mpcplx).Imag) == 0 {
			goto settrue
		}
		goto setfalse

	case ONE_ | CTCPLX_:
		if v.U.(*Mpcplx).Real.Cmp(&rv.U.(*Mpcplx).Real) != 0 || v.U.(*Mpcplx).Imag.Cmp(&rv.U.(*Mpcplx).Imag) != 0 {
			goto settrue
		}
		goto setfalse

	case OEQ_ | CTSTR_:
		if strlit(nl) == strlit(nr) {
			goto settrue
		}
		goto setfalse

	case ONE_ | CTSTR_:
		if strlit(nl) != strlit(nr) {
			goto settrue
		}
		goto setfalse

	case OLT_ | CTSTR_:
		if strlit(nl) < strlit(nr) {
			goto settrue
		}
		goto setfalse

	case OLE_ | CTSTR_:
		if strlit(nl) <= strlit(nr) {
			goto settrue
		}
		goto setfalse

	case OGE_ | CTSTR_:
		if strlit(nl) >= strlit(nr) {
			goto settrue
		}
		goto setfalse

	case OGT_ | CTSTR_:
		if strlit(nl) > strlit(nr) {
			goto settrue
		}
		goto setfalse

	case OOROR_ | CTBOOL_:
		if v.U.(bool) || rv.U.(bool) {
			goto settrue
		}
		goto setfalse

	case OANDAND_ | CTBOOL_:
		if v.U.(bool) && rv.U.(bool) {
			goto settrue
		}
		goto setfalse

	case OEQ_ | CTBOOL_:
		if v.U.(bool) == rv.U.(bool) {
			goto settrue
		}
		goto setfalse

	case ONE_ | CTBOOL_:
		if v.U.(bool) != rv.U.(bool) {
			goto settrue
		}
		goto setfalse
	}

ret:
	psess.
		setconst(n, v)
	return

settrue:
	psess.
		setconst(n, Val{true})
	return

setfalse:
	psess.
		setconst(n, Val{false})
	return

illegal:
	if !n.Diag() {
		psess.
			yyerror("illegal constant expression: %v %v %v", nl.Type, n.Op, nr.Type)
		n.SetDiag(true)
	}
}

// setconst rewrites n as an OLITERAL with value v.
func (psess *PackageSession) setconst(n *Node, v Val) {

	if n.Orig == n {
		n.Orig = n.copy()
		n.Orig.Orig = n.Orig
	}

	*n = Node{
		Op:      OLITERAL,
		Pos:     n.Pos,
		Orig:    n.Orig,
		Type:    n.Type,
		Xoffset: BADWIDTH,
	}
	n.SetVal(psess, v)

	lno := psess.setlineno(n)
	psess.
		overflow(v, n.Type)
	psess.
		lineno = lno

	if v.Ctype(psess) == CTFLT && n.Type.Etype != TIDEAL {
		n.SetVal(psess, Val{psess.truncfltlit(v.U.(*Mpflt), n.Type)})
	}
}

func (psess *PackageSession) setintconst(n *Node, v int64) {
	u := new(Mpint)
	u.SetInt64(v)
	psess.
		setconst(n, Val{u})
}

// nodlit returns a new untyped constant with value v.
func (psess *PackageSession) nodlit(v Val) *Node {
	n := psess.nod(OLITERAL, nil, nil)
	n.SetVal(psess, v)
	switch v.Ctype(psess) {
	default:
		psess.
			Fatalf("nodlit ctype %d", v.Ctype(psess))

	case CTSTR:
		n.Type = psess.types.Idealstring

	case CTBOOL:
		n.Type = psess.types.Idealbool

	case CTINT, CTRUNE, CTFLT, CTCPLX:
		n.Type = psess.types.Types[TIDEAL]

	case CTNIL:
		n.Type = psess.types.Types[TNIL]
	}

	return n
}

// idealkind returns a constant kind like consttype
// but for an arbitrary "ideal" (untyped constant) expression.
func (psess *PackageSession) idealkind(n *Node) Ctype {
	if n == nil || !n.Type.IsUntyped(psess.types) {
		return CTxxx
	}

	switch n.Op {
	default:
		return CTxxx

	case OLITERAL:
		return n.Val().Ctype(psess)

	case OADD,
		OAND,
		OANDNOT,
		OCOM,
		ODIV,
		OMINUS,
		OMOD,
		OMUL,
		OSUB,
		OXOR,
		OOR,
		OPLUS:
		k1 := psess.idealkind(n.Left)

		k2 := psess.idealkind(n.Right)
		if k1 > k2 {
			return k1
		} else {
			return k2
		}

	case OREAL, OIMAG:
		return CTFLT

	case OCOMPLEX:
		return CTCPLX

	case OADDSTR:
		return CTSTR

	case OANDAND,
		OEQ,
		OGE,
		OGT,
		OLE,
		OLT,
		ONE,
		ONOT,
		OOROR,
		OCMPSTR,
		OCMPIFACE:
		return CTBOOL

	case OLSH, ORSH:
		return psess.idealkind(n.Left)
	}
}

// The result of defaultlit MUST be assigned back to n, e.g.
// 	n.Left = defaultlit(n.Left, t)
func (psess *PackageSession) defaultlit(n *Node, t *types.Type) *Node {
	return psess.defaultlitreuse(n, t, noReuse)
}

// The result of defaultlitreuse MUST be assigned back to n, e.g.
// 	n.Left = defaultlitreuse(n.Left, t, reuse)
func (psess *PackageSession) defaultlitreuse(n *Node, t *types.Type, reuse canReuseNode) *Node {
	if n == nil || !n.Type.IsUntyped(psess.types) {
		return n
	}

	if n.Op == OLITERAL && !reuse {
		n = n.copy()
		reuse = true
	}

	lno := psess.setlineno(n)
	ctype := psess.idealkind(n)
	var t1 *types.Type
	switch ctype {
	default:
		if t != nil {
			return psess.convlit(n, t)
		}

		switch n.Val().Ctype(psess) {
		case CTNIL:
			psess.
				lineno = lno
			if !n.Diag() {
				psess.
					yyerror("use of untyped nil")
				n.SetDiag(true)
			}

			n.Type = nil
		case CTSTR:
			t1 := psess.types.Types[TSTRING]
			n = psess.convlit1(n, t1, false, reuse)
		default:
			psess.
				yyerror("defaultlit: unknown literal: %v", n)
		}
		psess.
			lineno = lno
		return n

	case CTxxx:
		psess.
			Fatalf("defaultlit: idealkind is CTxxx: %+v", n)

	case CTBOOL:
		t1 := psess.types.Types[TBOOL]
		if t != nil && t.IsBoolean() {
			t1 = t
		}
		n = psess.convlit1(n, t1, false, reuse)
		psess.
			lineno = lno
		return n

	case CTINT:
		t1 = psess.types.Types[TINT]
	case CTRUNE:
		t1 = psess.types.Runetype
	case CTFLT:
		t1 = psess.types.Types[TFLOAT64]
	case CTCPLX:
		t1 = psess.types.Types[TCOMPLEX128]
	}

	v1 := n.Val()
	if t != nil {
		if t.IsInteger() {
			t1 = t
			v1 = psess.toint(n.Val())
		} else if t.IsFloat() {
			t1 = t
			v1 = psess.toflt(n.Val())
		} else if t.IsComplex() {
			t1 = t
			v1 = tocplx(n.Val())
		}
		if n.Val().Ctype(psess) != CTxxx {
			n.SetVal(psess, v1)
		}
	}

	if n.Val().Ctype(psess) != CTxxx {
		psess.
			overflow(n.Val(), t1)
	}
	n = psess.convlit1(n, t1, false, reuse)
	psess.
		lineno = lno
	return n
}

// defaultlit on both nodes simultaneously;
// if they're both ideal going in they better
// get the same type going out.
// force means must assign concrete (non-ideal) type.
// The results of defaultlit2 MUST be assigned back to l and r, e.g.
// 	n.Left, n.Right = defaultlit2(n.Left, n.Right, force)
func (psess *PackageSession) defaultlit2(l *Node, r *Node, force bool) (*Node, *Node) {
	if l.Type == nil || r.Type == nil {
		return l, r
	}
	if !l.Type.IsUntyped(psess.types) {
		r = psess.convlit(r, l.Type)
		return l, r
	}

	if !r.Type.IsUntyped(psess.types) {
		l = psess.convlit(l, r.Type)
		return l, r
	}

	if !force {
		return l, r
	}

	if l.Type.IsBoolean() {
		l = psess.convlit(l, psess.types.Types[TBOOL])
		r = psess.convlit(r, psess.types.Types[TBOOL])
	}

	lkind := psess.idealkind(l)
	rkind := psess.idealkind(r)
	if lkind == CTCPLX || rkind == CTCPLX {
		l = psess.convlit(l, psess.types.Types[TCOMPLEX128])
		r = psess.convlit(r, psess.types.Types[TCOMPLEX128])
		return l, r
	}

	if lkind == CTFLT || rkind == CTFLT {
		l = psess.convlit(l, psess.types.Types[TFLOAT64])
		r = psess.convlit(r, psess.types.Types[TFLOAT64])
		return l, r
	}

	if lkind == CTRUNE || rkind == CTRUNE {
		l = psess.convlit(l, psess.types.Runetype)
		r = psess.convlit(r, psess.types.Runetype)
		return l, r
	}

	l = psess.convlit(l, psess.types.Types[TINT])
	r = psess.convlit(r, psess.types.Types[TINT])

	return l, r
}

// strlit returns the value of a literal string Node as a string.
func strlit(n *Node) string {
	return n.Val().U.(string)
}

func (psess *PackageSession) smallintconst(n *Node) bool {
	if n.Op == OLITERAL && psess.Isconst(n, CTINT) && n.Type != nil {
		switch psess.simtype[n.Type.Etype] {
		case TINT8,
			TUINT8,
			TINT16,
			TUINT16,
			TINT32,
			TUINT32,
			TBOOL,
			TPTR32:
			return true

		case TIDEAL, TINT64, TUINT64, TPTR64:
			v, ok := n.Val().U.(*Mpint)
			if ok && v.Cmp(psess.minintval[TINT32]) > 0 && v.Cmp(psess.maxintval[TINT32]) < 0 {
				return true
			}
		}
	}

	return false
}

// nonnegintconst checks if Node n contains a constant expression
// representable as a non-negative small integer, and returns its
// (integer) value if that's the case. Otherwise, it returns -1.
func (psess *PackageSession) nonnegintconst(n *Node) int64 {
	if n.Op != OLITERAL {
		return -1
	}

	v := psess.toint(n.Val())
	vi, ok := v.U.(*Mpint)
	if !ok || vi.CmpInt64(0) < 0 || vi.Cmp(psess.maxintval[TINT32]) > 0 {
		return -1
	}

	return vi.Int64(psess)
}

// isGoConst reports whether n is a Go language constant (as opposed to a
// compile-time constant).
//
// Expressions derived from nil, like string([]byte(nil)), while they
// may be known at compile time, are not Go language constants.
// Only called for expressions known to evaluated to compile-time
// constants.
func (n *Node) isGoConst(psess *PackageSession) bool {
	if n.Orig != nil {
		n = n.Orig
	}

	switch n.Op {
	case OADD,
		OADDSTR,
		OAND,
		OANDAND,
		OANDNOT,
		OCOM,
		ODIV,
		OEQ,
		OGE,
		OGT,
		OLE,
		OLSH,
		OLT,
		OMINUS,
		OMOD,
		OMUL,
		ONE,
		ONOT,
		OOR,
		OOROR,
		OPLUS,
		ORSH,
		OSUB,
		OXOR,
		OIOTA,
		OCOMPLEX,
		OREAL,
		OIMAG:
		if n.Left.isGoConst(psess) && (n.Right == nil || n.Right.isGoConst(psess)) {
			return true
		}

	case OCONV:
		if psess.okforconst[n.Type.Etype] && n.Left.isGoConst(psess) {
			return true
		}

	case OLEN, OCAP:
		l := n.Left
		if l.isGoConst(psess) {
			return true
		}

		t := l.Type

		if t != nil && t.IsPtr() {
			t = t.Elem(psess.types)
		}
		if t != nil && t.IsArray() && !hascallchan(l) {
			return true
		}

	case OLITERAL:
		if n.Val().Ctype(psess) != CTNIL {
			return true
		}

	case ONAME:
		l := asNode(n.Sym.Def)
		if l != nil && l.Op == OLITERAL && n.Val().Ctype(psess) != CTNIL {
			return true
		}

	case ONONAME:
		if asNode(n.Sym.Def) != nil && asNode(n.Sym.Def).Op == OIOTA {
			return true
		}

	case OALIGNOF, OOFFSETOF, OSIZEOF:
		return true
	}

	return false
}

func hascallchan(n *Node) bool {
	if n == nil {
		return false
	}
	switch n.Op {
	case OAPPEND,
		OCALL,
		OCALLFUNC,
		OCALLINTER,
		OCALLMETH,
		OCAP,
		OCLOSE,
		OCOMPLEX,
		OCOPY,
		ODELETE,
		OIMAG,
		OLEN,
		OMAKE,
		ONEW,
		OPANIC,
		OPRINT,
		OPRINTN,
		OREAL,
		ORECOVER,
		ORECV:
		return true
	}

	if hascallchan(n.Left) || hascallchan(n.Right) {
		return true
	}
	for _, n1 := range n.List.Slice() {
		if hascallchan(n1) {
			return true
		}
	}
	for _, n2 := range n.Rlist.Slice() {
		if hascallchan(n2) {
			return true
		}
	}

	return false
}
