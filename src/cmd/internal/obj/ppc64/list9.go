package ppc64

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
)

func (psess *PackageSession) init() {
	psess.obj.
		RegisterRegister(obj.RBasePPC64, REG_DCR0+1024, rconv)
	psess.obj.
		RegisterOpcode(obj.ABasePPC64, psess.Anames)
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
	if REG_V0 <= r && r <= REG_V31 {
		return fmt.Sprintf("V%d", r-REG_V0)
	}
	if REG_VS0 <= r && r <= REG_VS63 {
		return fmt.Sprintf("VS%d", r-REG_VS0)
	}
	if REG_CR0 <= r && r <= REG_CR7 {
		return fmt.Sprintf("CR%d", r-REG_CR0)
	}
	if r == REG_CR {
		return "CR"
	}
	if REG_SPR0 <= r && r <= REG_SPR0+1023 {
		switch r {
		case REG_XER:
			return "XER"

		case REG_LR:
			return "LR"

		case REG_CTR:
			return "CTR"
		}

		return fmt.Sprintf("SPR(%d)", r-REG_SPR0)
	}

	if REG_DCR0 <= r && r <= REG_DCR0+1023 {
		return fmt.Sprintf("DCR(%d)", r-REG_DCR0)
	}
	if r == REG_FPSCR {
		return "FPSCR"
	}
	if r == REG_MSR {
		return "MSR"
	}

	return fmt.Sprintf("Rgok(%d)", r-obj.RBasePPC64)
}

func (psess *PackageSession) DRconv(a int) string {
	s := "C_??"
	if a >= C_NONE && a <= C_NCLASS {
		s = psess.cnames9[a]
	}
	var fp string
	fp += s
	return fp
}
