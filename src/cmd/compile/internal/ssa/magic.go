package ssa

import "math/big"

// umagicOK returns whether we should strength reduce a n-bit divide by c.
func umagicOK(n uint, c int64) bool {

	d := uint64(c) << (64 - n) >> (64 - n)

	return d&(d-1) != 0
}

type umagicData struct {
	s int64  // ⎡log2(c)⎤
	m uint64 // ⎡2^(n+s)/c⎤ - 2^n
}

// umagic computes the constants needed to strength reduce unsigned n-bit divides by the constant uint64(c).
// The return values satisfy for all 0 <= x < 2^n
//  floor(x / uint64(c)) = x * (m + 2^n) >> (n+s)
func umagic(n uint, c int64) umagicData {

	d := uint64(c) << (64 - n) >> (64 - n)

	C := new(big.Int).SetUint64(d)
	s := C.BitLen()
	M := big.NewInt(1)
	M.Lsh(M, n+uint(s))
	M.Add(M, C)
	M.Sub(M, big.NewInt(1))
	M.Div(M, C)
	if M.Bit(int(n)) != 1 {
		panic("n+1st bit isn't set")
	}
	M.SetBit(M, int(n), 0)
	m := M.Uint64()
	return umagicData{s: int64(s), m: m}
}

func smagicOK(n uint, c int64) bool {
	if c < 0 {

		return false
	}

	return c&(c-1) != 0
}

type smagicData struct {
	s int64  // ⎡log2(c)⎤-1
	m uint64 // ⎡2^(n+s)/c⎤
}

// magic computes the constants needed to strength reduce signed n-bit divides by the constant c.
// Must have c>0.
// The return values satisfy for all -2^(n-1) <= x < 2^(n-1)
//  trunc(x / c) = x * m >> (n+s) + (x < 0 ? 1 : 0)
func smagic(n uint, c int64) smagicData {
	C := new(big.Int).SetInt64(c)
	s := C.BitLen() - 1
	M := big.NewInt(1)
	M.Lsh(M, n+uint(s))
	M.Add(M, C)
	M.Sub(M, big.NewInt(1))
	M.Div(M, C)
	if M.Bit(int(n)) != 0 {
		panic("n+1st bit is set")
	}
	if M.Bit(int(n-1)) == 0 {
		panic("nth bit is not set")
	}
	m := M.Uint64()
	return smagicData{s: int64(s), m: m}
}
