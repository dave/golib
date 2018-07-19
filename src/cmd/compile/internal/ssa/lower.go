package ssa

// convert to machine-dependent ops
func (psess *PackageSession) lower(f *Func) {
	psess.
		applyRewrite(f, f.Config.lowerBlock, f.Config.lowerValue)
}

// checkLower checks for unlowered opcodes and fails if we find one.
func (psess *PackageSession) checkLower(f *Func) {

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if !psess.opcodeTable[v.Op].generic {
				continue
			}
			switch v.Op {
			case OpSP, OpSB, OpInitMem, OpArg, OpPhi, OpVarDef, OpVarKill, OpVarLive, OpKeepAlive, OpSelect0, OpSelect1, OpConvert:
				continue
			case OpGetG:
				if f.Config.hasGReg {

					continue
				}
			}
			s := "not lowered: " + v.String() + ", " + v.Op.String(psess) + " " + v.Type.SimpleString(psess.types)
			for _, a := range v.Args {
				s += " " + a.Type.SimpleString(psess.types)
			}
			f.Fatalf("%s", s)
		}
	}
}
