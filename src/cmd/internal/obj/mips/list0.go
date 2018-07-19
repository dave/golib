package mips

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
)

func (psess *PackageSession) init() {
	psess.obj.
		RegisterRegister(obj.RBaseMIPS, REG_LAST+1, rconv)
	psess.obj.
		RegisterOpcode(obj.ABaseMIPS, psess.Anames)
}

func rconv(r int) string {
	if r == 0 {
		return "NONE"
	}
	if r == REGG {

		return "g"
	}
	if REG_R0 <= r && r <= REG_R31 {
		return fmt.Sprintf("R%d", r-REG_R0)
	}
	if REG_F0 <= r && r <= REG_F31 {
		return fmt.Sprintf("F%d", r-REG_F0)
	}
	if REG_M0 <= r && r <= REG_M31 {
		return fmt.Sprintf("M%d", r-REG_M0)
	}
	if REG_FCR0 <= r && r <= REG_FCR31 {
		return fmt.Sprintf("FCR%d", r-REG_FCR0)
	}
	if r == REG_HI {
		return "HI"
	}
	if r == REG_LO {
		return "LO"
	}

	return fmt.Sprintf("Rgok(%d)", r-obj.RBaseMIPS)
}

func (psess *PackageSession) DRconv(a int) string {
	s := "C_??"
	if a >= C_NONE && a <= C_NCLASS {
		s = psess.cnames0[a]
	}
	var fp string
	fp += s
	return fp
}
