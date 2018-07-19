package x86

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
)

func (psess *PackageSession) init() {
	psess.obj.
		RegisterRegister(REG_AL, REG_AL+len(psess.Register), psess.rconv)
	psess.obj.
		RegisterOpcode(obj.ABaseAMD64, psess.Anames)
	psess.obj.
		RegisterRegisterList(obj.RegListX86Lo, obj.RegListX86Hi, psess.rlconv)
	psess.obj.
		RegisterOpSuffix("386", psess.opSuffixString)
	psess.obj.
		RegisterOpSuffix("amd64", psess.opSuffixString)
}

func (psess *PackageSession) rconv(r int) string {
	if REG_AL <= r && r-REG_AL < len(psess.Register) {
		return psess.Register[r-REG_AL]
	}
	return fmt.Sprintf("Rgok(%d)", r-obj.RBaseAMD64)
}

func (psess *PackageSession) rlconv(bits int64) string {
	reg0, reg1 := decodeRegisterRange(bits)
	return fmt.Sprintf("[%s-%s]", psess.rconv(reg0), psess.rconv(reg1))
}

func (psess *PackageSession) opSuffixString(s uint8) string {
	return "." + opSuffix(s).String(psess)
}
