package gc

import (
	"fmt"
	"math"
	"math/big"
)

const (
	// Maximum size in bits for Mpints before signalling
	// overflow and also mantissa precision for Mpflts.
	Mpprec = 512
	// Turn on for constant arithmetic debugging output.
	Mpdebug = false
)

// Mpflt represents a floating-point constant.
type Mpflt struct {
	Val big.Float
}

// Mpcplx represents a complex constant.
type Mpcplx struct {
	Real Mpflt
	Imag Mpflt
}

func newMpflt() *Mpflt {
	var a Mpflt
	a.Val.SetPrec(Mpprec)
	return &a
}

func newMpcmplx() *Mpcplx {
	var a Mpcplx
	a.Real = *newMpflt()
	a.Imag = *newMpflt()
	return &a
}

func (a *Mpflt) SetInt(b *Mpint) {
	if b.checkOverflow(0) {

		a.Val.SetInf(b.Val.Sign() < 0)
		return
	}
	a.Val.SetInt(&b.Val)
}

func (a *Mpflt) Set(b *Mpflt) {
	a.Val.Set(&b.Val)
}

func (a *Mpflt) Add(b *Mpflt) {
	if Mpdebug {
		fmt.Printf("\n%v + %v", a, b)
	}

	a.Val.Add(&a.Val, &b.Val)

	if Mpdebug {
		fmt.Printf(" = %v\n\n", a)
	}
}

func (a *Mpflt) AddFloat64(c float64) {
	var b Mpflt

	b.SetFloat64(c)
	a.Add(&b)
}

func (a *Mpflt) Sub(b *Mpflt) {
	if Mpdebug {
		fmt.Printf("\n%v - %v", a, b)
	}

	a.Val.Sub(&a.Val, &b.Val)

	if Mpdebug {
		fmt.Printf(" = %v\n\n", a)
	}
}

func (a *Mpflt) Mul(b *Mpflt) {
	if Mpdebug {
		fmt.Printf("%v\n * %v\n", a, b)
	}

	a.Val.Mul(&a.Val, &b.Val)

	if Mpdebug {
		fmt.Printf(" = %v\n\n", a)
	}
}

func (a *Mpflt) MulFloat64(c float64) {
	var b Mpflt

	b.SetFloat64(c)
	a.Mul(&b)
}

func (a *Mpflt) Quo(b *Mpflt) {
	if Mpdebug {
		fmt.Printf("%v\n / %v\n", a, b)
	}

	a.Val.Quo(&a.Val, &b.Val)

	if Mpdebug {
		fmt.Printf(" = %v\n\n", a)
	}
}

func (a *Mpflt) Cmp(b *Mpflt) int {
	return a.Val.Cmp(&b.Val)
}

func (a *Mpflt) CmpFloat64(c float64) int {
	if c == 0 {
		return a.Val.Sign()
	}
	return a.Val.Cmp(big.NewFloat(c))
}

func (a *Mpflt) Float64(psess *PackageSession) float64 {
	x, _ := a.Val.Float64()

	if math.IsInf(x, 0) && psess.nsavederrors+psess.nerrors == 0 {
		psess.
			Fatalf("ovf in Mpflt Float64")
	}

	return x + 0
}

func (a *Mpflt) Float32(psess *PackageSession) float64 {
	x32, _ := a.Val.Float32()
	x := float64(x32)

	if math.IsInf(x, 0) && psess.nsavederrors+psess.nerrors == 0 {
		psess.
			Fatalf("ovf in Mpflt Float32")
	}

	return x + 0
}

func (a *Mpflt) SetFloat64(c float64) {
	if Mpdebug {
		fmt.Printf("\nconst %g", c)
	}

	if c == 0 {
		c = 0
	}
	a.Val.SetFloat64(c)

	if Mpdebug {
		fmt.Printf(" = %v\n", a)
	}
}

func (a *Mpflt) Neg() {

	if a.Val.Sign() != 0 {
		a.Val.Neg(&a.Val)
	}
}

func (a *Mpflt) SetString(psess *PackageSession, as string) {
	for len(as) > 0 && (as[0] == ' ' || as[0] == '\t') {
		as = as[1:]
	}

	f, _, err := a.Val.Parse(as, 10)
	if err != nil {
		psess.
			yyerror("malformed constant: %s (%v)", as, err)
		a.Val.SetFloat64(0)
		return
	}

	if f.IsInf() {
		psess.
			yyerror("constant too large: %s", as)
		a.Val.SetFloat64(0)
		return
	}

	if f.Sign() == 0 && f.Signbit() {
		a.Val.SetFloat64(0)
	}
}

func (f *Mpflt) String() string {
	return fconv(f, 0)
}

func fconv(fvp *Mpflt, flag FmtFlag) string {
	if flag&FmtSharp == 0 {
		return fvp.Val.Text('b', 0)
	}

	f := &fvp.Val
	var sign string
	if f.Sign() < 0 {
		sign = "-"
		f = new(big.Float).Abs(f)
	} else if flag&FmtSign != 0 {
		sign = "+"
	}

	if f.IsInf() {
		return sign + "Inf"
	}

	if x, _ := f.Float64(); f.Sign() == 0 == (x == 0) && !math.IsInf(x, 0) {
		return fmt.Sprintf("%s%.6g", sign, x)
	}

	// Out of float64 range. Do approximate manual to decimal
	// conversion to avoid precise but possibly slow Float
	// formatting.
	// f = mant * 2**exp
	var mant big.Float
	exp := f.MantExp(&mant)

	m, _ := mant.Float64()
	d := float64(exp) * (math.Ln2 / math.Ln10)

	e := int64(d)
	m *= math.Pow(10, d-float64(e))

	switch {
	case m < 1-0.5e-6:

		m *= 10
		e--
	case m >= 10:
		m /= 10
		e++
	}

	return fmt.Sprintf("%s%.6ge%+d", sign, m, e)
}

// complex multiply v *= rv
//	(a, b) * (c, d) = (a*c - b*d, b*c + a*d)
func (v *Mpcplx) Mul(rv *Mpcplx) {
	var ac, ad, bc, bd Mpflt

	ac.Set(&v.Real)
	ac.Mul(&rv.Real)

	bd.Set(&v.Imag)
	bd.Mul(&rv.Imag)

	bc.Set(&v.Imag)
	bc.Mul(&rv.Real)

	ad.Set(&v.Real)
	ad.Mul(&rv.Imag)

	v.Real.Set(&ac)
	v.Real.Sub(&bd)

	v.Imag.Set(&bc)
	v.Imag.Add(&ad)
}

// complex divide v /= rv
//	(a, b) / (c, d) = ((a*c + b*d), (b*c - a*d))/(c*c + d*d)
func (v *Mpcplx) Div(rv *Mpcplx) bool {
	if rv.Real.CmpFloat64(0) == 0 && rv.Imag.CmpFloat64(0) == 0 {
		return false
	}

	var ac, ad, bc, bd, cc_plus_dd Mpflt

	cc_plus_dd.Set(&rv.Real)
	cc_plus_dd.Mul(&rv.Real)

	ac.Set(&rv.Imag)
	ac.Mul(&rv.Imag)
	cc_plus_dd.Add(&ac)

	if cc_plus_dd.CmpFloat64(0) == 0 {
		return false
	}

	ac.Set(&v.Real)
	ac.Mul(&rv.Real)

	bd.Set(&v.Imag)
	bd.Mul(&rv.Imag)

	bc.Set(&v.Imag)
	bc.Mul(&rv.Real)

	ad.Set(&v.Real)
	ad.Mul(&rv.Imag)

	v.Real.Set(&ac)
	v.Real.Add(&bd)
	v.Real.Quo(&cc_plus_dd)

	v.Imag.Set(&bc)
	v.Imag.Sub(&ad)
	v.Imag.Quo(&cc_plus_dd)

	return true
}
