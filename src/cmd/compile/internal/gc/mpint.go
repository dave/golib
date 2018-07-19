package gc

import (
	"fmt"
	"math/big"
)

// Mpint represents an integer constant.
type Mpint struct {
	Val  big.Int
	Ovf  bool // set if Val overflowed compiler limit (sticky)
	Rune bool // set if syntax indicates default type rune
}

func (a *Mpint) SetOverflow() {
	a.Val.SetUint64(1)
	a.Ovf = true
}

func (a *Mpint) checkOverflow(extra int) bool {

	if a.Val.BitLen()+extra > Mpprec {
		a.SetOverflow()
	}
	return a.Ovf
}

func (a *Mpint) Set(b *Mpint) {
	a.Val.Set(&b.Val)
}

func (a *Mpint) SetFloat(b *Mpflt) bool {

	if b.Val.MantExp(nil) > 2*Mpprec {
		a.SetOverflow()
		return false
	}

	if _, acc := b.Val.Int(&a.Val); acc == big.Exact {
		return true
	}

	const delta = 16 // a reasonably small number of bits > 0
	var t big.Float
	t.SetPrec(Mpprec - delta)

	t.SetMode(big.ToZero)
	t.Set(&b.Val)
	if _, acc := t.Int(&a.Val); acc == big.Exact {
		return true
	}

	t.SetMode(big.AwayFromZero)
	t.Set(&b.Val)
	if _, acc := t.Int(&a.Val); acc == big.Exact {
		return true
	}

	a.Ovf = false
	return false
}

func (a *Mpint) Add(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Add")
		}
		a.SetOverflow()
		return
	}

	a.Val.Add(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		psess.
			yyerror("constant addition overflow")
	}
}

func (a *Mpint) Sub(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Sub")
		}
		a.SetOverflow()
		return
	}

	a.Val.Sub(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		psess.
			yyerror("constant subtraction overflow")
	}
}

func (a *Mpint) Mul(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Mul")
		}
		a.SetOverflow()
		return
	}

	a.Val.Mul(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		psess.
			yyerror("constant multiplication overflow")
	}
}

func (a *Mpint) Quo(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Quo")
		}
		a.SetOverflow()
		return
	}

	a.Val.Quo(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		psess.
			yyerror("constant division overflow")
	}
}

func (a *Mpint) Rem(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Rem")
		}
		a.SetOverflow()
		return
	}

	a.Val.Rem(&a.Val, &b.Val)

	if a.checkOverflow(0) {
		psess.
			yyerror("constant modulo overflow")
	}
}

func (a *Mpint) Or(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Or")
		}
		a.SetOverflow()
		return
	}

	a.Val.Or(&a.Val, &b.Val)
}

func (a *Mpint) And(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint And")
		}
		a.SetOverflow()
		return
	}

	a.Val.And(&a.Val, &b.Val)
}

func (a *Mpint) AndNot(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint AndNot")
		}
		a.SetOverflow()
		return
	}

	a.Val.AndNot(&a.Val, &b.Val)
}

func (a *Mpint) Xor(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Xor")
		}
		a.SetOverflow()
		return
	}

	a.Val.Xor(&a.Val, &b.Val)
}

func (a *Mpint) Lsh(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Lsh")
		}
		a.SetOverflow()
		return
	}

	s := b.Int64(psess)
	if s < 0 || s >= Mpprec {
		msg := "shift count too large"
		if s < 0 {
			msg = "invalid negative shift count"
		}
		psess.
			yyerror("%s: %d", msg, s)
		a.SetInt64(0)
		return
	}

	if a.checkOverflow(int(s)) {
		psess.
			yyerror("constant shift overflow")
		return
	}
	a.Val.Lsh(&a.Val, uint(s))
}

func (a *Mpint) Rsh(psess *PackageSession, b *Mpint) {
	if a.Ovf || b.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("ovf in Mpint Rsh")
		}
		a.SetOverflow()
		return
	}

	s := b.Int64(psess)
	if s < 0 {
		psess.
			yyerror("invalid negative shift count: %d", s)
		if a.Val.Sign() < 0 {
			a.SetInt64(-1)
		} else {
			a.SetInt64(0)
		}
		return
	}

	a.Val.Rsh(&a.Val, uint(s))
}

func (a *Mpint) Cmp(b *Mpint) int {
	return a.Val.Cmp(&b.Val)
}

func (a *Mpint) CmpInt64(c int64) int {
	if c == 0 {
		return a.Val.Sign()
	}
	return a.Val.Cmp(big.NewInt(c))
}

func (a *Mpint) Neg() {
	a.Val.Neg(&a.Val)
}

func (a *Mpint) Int64(psess *PackageSession) int64 {
	if a.Ovf {
		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				Fatalf("constant overflow")
		}
		return 0
	}

	return a.Val.Int64()
}

func (a *Mpint) SetInt64(c int64) {
	a.Val.SetInt64(c)
}

func (a *Mpint) SetString(psess *PackageSession, as string) {
	_, ok := a.Val.SetString(as, 0)
	if !ok {
		psess.
			yyerror("malformed integer constant: %s", as)
		a.Val.SetUint64(0)
		return
	}
	if a.checkOverflow(0) {
		psess.
			yyerror("constant too large: %s", as)
	}
}

func (x *Mpint) String() string {
	return bconv(x, 0)
}

func bconv(xval *Mpint, flag FmtFlag) string {
	if flag&FmtSharp != 0 {
		return fmt.Sprintf("%#x", &xval.Val)
	}
	return xval.Val.String()
}
