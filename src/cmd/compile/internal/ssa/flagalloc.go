package ssa

// flagalloc allocates the flag register among all the flag-generating
// instructions. Flag values are recomputed if they need to be
// spilled/restored.
func (psess *PackageSession) flagalloc(f *Func) {

	end := make([]*Value, f.NumBlocks())
	po := f.postorder()
	for n := 0; n < 2; n++ {
		for _, b := range po {

			flag := end[b.ID]
			if b.Control != nil && b.Control.Type.IsFlags(psess.types) {
				flag = b.Control
			}
			for j := len(b.Values) - 1; j >= 0; j-- {
				v := b.Values[j]
				if v == flag {
					flag = nil
				}
				if v.clobbersFlags(psess) {
					flag = nil
				}
				for _, a := range v.Args {
					if a.Type.IsFlags(psess.types) {
						flag = a
					}
				}
			}
			if flag != nil {
				for _, e := range b.Preds {
					p := e.b
					end[p.ID] = flag
				}
			}
		}
	}

	for _, b := range f.Blocks {
		v := b.Control
		if v != nil && v.Type.IsFlags(psess.types) && end[b.ID] != v {
			end[b.ID] = nil
		}
		if b.Kind == BlockDefer {

			end[b.ID] = nil
		}
	}

	spill := map[ID]bool{}
	for _, b := range f.Blocks {
		var flag *Value
		if len(b.Preds) > 0 {
			flag = end[b.Preds[0].b.ID]
		}
		for _, v := range b.Values {
			for _, a := range v.Args {
				if !a.Type.IsFlags(psess.types) {
					continue
				}
				if a == flag {
					continue
				}

				spill[a.ID] = true
				flag = a
			}
			if v.clobbersFlags(psess) {
				flag = nil
			}
			if v.Type.IsFlags(psess.types) {
				flag = v
			}
		}
		if v := b.Control; v != nil && v != flag && v.Type.IsFlags(psess.types) {
			spill[v.ID] = true
		}
		if v := end[b.ID]; v != nil && v != flag {
			spill[v.ID] = true
		}
	}

	// Add flag spill and recomputation where they are needed.
	// TODO: Remove original instructions if they are never used.
	var oldSched []*Value
	for _, b := range f.Blocks {
		oldSched = append(oldSched[:0], b.Values...)
		b.Values = b.Values[:0]
		// The current live flag value (the pre-flagalloc copy).
		var flag *Value
		if len(b.Preds) > 0 {
			flag = end[b.Preds[0].b.ID]

			for _, e := range b.Preds[1:] {
				p := e.b
				if end[p.ID] != flag {
					f.Fatalf("live flag in %s's predecessors not consistent", b)
				}
			}
		}
		for _, v := range oldSched {
			if v.Op == OpPhi && v.Type.IsFlags(psess.types) {
				f.Fatalf("phi of flags not supported: %s", v.LongString(psess))
			}

			if spill[v.ID] && v.MemoryArg(psess) != nil {
				switch v.Op {
				case OpAMD64CMPQload:
					load := b.NewValue2IA(v.Pos, OpAMD64MOVQload, f.Config.Types.UInt64, v.AuxInt, v.Aux, v.Args[0], v.Args[2])
					v.Op = OpAMD64CMPQ
					v.AuxInt = 0
					v.Aux = nil
					v.SetArgs2(load, v.Args[1])
				case OpAMD64CMPLload:
					load := b.NewValue2IA(v.Pos, OpAMD64MOVLload, f.Config.Types.UInt32, v.AuxInt, v.Aux, v.Args[0], v.Args[2])
					v.Op = OpAMD64CMPL
					v.AuxInt = 0
					v.Aux = nil
					v.SetArgs2(load, v.Args[1])
				case OpAMD64CMPWload:
					load := b.NewValue2IA(v.Pos, OpAMD64MOVWload, f.Config.Types.UInt16, v.AuxInt, v.Aux, v.Args[0], v.Args[2])
					v.Op = OpAMD64CMPW
					v.AuxInt = 0
					v.Aux = nil
					v.SetArgs2(load, v.Args[1])
				case OpAMD64CMPBload:
					load := b.NewValue2IA(v.Pos, OpAMD64MOVBload, f.Config.Types.UInt8, v.AuxInt, v.Aux, v.Args[0], v.Args[2])
					v.Op = OpAMD64CMPB
					v.AuxInt = 0
					v.Aux = nil
					v.SetArgs2(load, v.Args[1])

				case OpAMD64CMPQconstload:
					vo := v.AuxValAndOff(psess)
					load := b.NewValue2IA(v.Pos, OpAMD64MOVQload, f.Config.Types.UInt64, vo.Off(), v.Aux, v.Args[0], v.Args[1])
					v.Op = OpAMD64CMPQconst
					v.AuxInt = vo.Val()
					v.Aux = nil
					v.SetArgs1(load)
				case OpAMD64CMPLconstload:
					vo := v.AuxValAndOff(psess)
					load := b.NewValue2IA(v.Pos, OpAMD64MOVLload, f.Config.Types.UInt32, vo.Off(), v.Aux, v.Args[0], v.Args[1])
					v.Op = OpAMD64CMPLconst
					v.AuxInt = vo.Val()
					v.Aux = nil
					v.SetArgs1(load)
				case OpAMD64CMPWconstload:
					vo := v.AuxValAndOff(psess)
					load := b.NewValue2IA(v.Pos, OpAMD64MOVWload, f.Config.Types.UInt16, vo.Off(), v.Aux, v.Args[0], v.Args[1])
					v.Op = OpAMD64CMPWconst
					v.AuxInt = vo.Val()
					v.Aux = nil
					v.SetArgs1(load)
				case OpAMD64CMPBconstload:
					vo := v.AuxValAndOff(psess)
					load := b.NewValue2IA(v.Pos, OpAMD64MOVBload, f.Config.Types.UInt8, vo.Off(), v.Aux, v.Args[0], v.Args[1])
					v.Op = OpAMD64CMPBconst
					v.AuxInt = vo.Val()
					v.Aux = nil
					v.SetArgs1(load)

				default:
					f.Fatalf("can't split flag generator: %s", v.LongString(psess))
				}

			}

			for i, a := range v.Args {
				if !a.Type.IsFlags(psess.types) {
					continue
				}
				if a == flag {
					continue
				}

				c := psess.copyFlags(a, b)

				v.SetArg(i, c)

				flag = a
			}

			b.Values = append(b.Values, v)
			if v.clobbersFlags(psess) {
				flag = nil
			}
			if v.Type.IsFlags(psess.types) {
				flag = v
			}
		}
		if v := b.Control; v != nil && v != flag && v.Type.IsFlags(psess.types) {

			c := psess.copyFlags(v, b)
			b.SetControl(c)
			flag = v
		}
		if v := end[b.ID]; v != nil && v != flag {
			psess.
				copyFlags(v, b)

		}
	}

	for _, b := range f.Blocks {
		b.FlagsLiveAtEnd = end[b.ID] != nil
	}
}

func (v *Value) clobbersFlags(psess *PackageSession,) bool {
	if psess.opcodeTable[v.Op].clobberFlags {
		return true
	}
	if v.Type.IsTuple() && (v.Type.FieldType(psess.types, 0).IsFlags(psess.types) || v.Type.FieldType(psess.types, 1).IsFlags(psess.types)) {

		return true
	}
	return false
}

// copyFlags copies v (flag generator) into b, returns the copy.
// If v's arg is also flags, copy recursively.
func (psess *PackageSession) copyFlags(v *Value, b *Block) *Value {
	flagsArgs := make(map[int]*Value)
	for i, a := range v.Args {
		if a.Type.IsFlags(psess.types) || a.Type.IsTuple() {
			flagsArgs[i] = psess.copyFlags(a, b)
		}
	}
	c := v.copyInto(psess, b)
	for i, a := range flagsArgs {
		c.SetArg(i, a)
	}
	return c
}
