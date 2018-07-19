package ssa

// branchelim tries to eliminate branches by
// generating CondSelect instructions.
//
// Search for basic blocks that look like
//
// bb0            bb0
//  | \          /   \
//  | bb1  or  bb1   bb2    <- trivial if/else blocks
//  | /          \   /
// bb2            bb3
//
// where the intermediate blocks are mostly empty (with no side-effects);
// rewrite Phis in the postdominator as CondSelects.
func (psess *PackageSession) branchelim(f *Func) {

	switch f.Config.arch {
	case "arm64", "amd64":

	default:
		return
	}

	change := true
	for change {
		change = false
		for _, b := range f.Blocks {
			change = psess.elimIf(f, b) || psess.elimIfElse(f, b) || change
		}
	}
}

func (psess *PackageSession) canCondSelect(v *Value, arch string) bool {

	switch {
	case v.Type.Size(psess.types) > v.Block.Func.Config.RegSize:
		return false
	case v.Type.IsPtrShaped():
		return true
	case v.Type.IsInteger():
		if arch == "amd64" && v.Type.Size(psess.types) < 2 {

			return false
		}
		return true
	default:
		return false
	}
}

func (psess *PackageSession) elimIf(f *Func, dom *Block) bool {

	if dom.Kind != BlockIf || dom.Likely != BranchUnknown {
		return false
	}
	var simple, post *Block
	for i := range dom.Succs {
		bb, other := dom.Succs[i].Block(), dom.Succs[i^1].Block()
		if isLeafPlain(bb) && bb.Succs[0].Block() == other {
			simple = bb
			post = other
			break
		}
	}
	if simple == nil || len(post.Preds) != 2 || post == dom {
		return false
	}

	hasphis := false
	for _, v := range post.Values {
		if v.Op == OpPhi {
			hasphis = true
			if !psess.canCondSelect(v, f.Config.arch) {
				return false
			}
		}
	}
	if !hasphis {
		return false
	}

	// Pick some upper bound for the number of instructions
	// we'd be willing to execute just to generate a dead
	// argument to CondSelect. In the worst case, this is
	// the number of useless instructions executed.
	const maxfuseinsts = 2

	if len(simple.Values) > maxfuseinsts || !psess.allTrivial(simple) {
		return false
	}

	swap := (post.Preds[0].Block() == dom) != (dom.Succs[0].Block() == post)
	for _, v := range post.Values {
		if v.Op != OpPhi {
			continue
		}
		v.Op = OpCondSelect
		if swap {
			v.Args[0], v.Args[1] = v.Args[1], v.Args[0]
		}
		v.AddArg(dom.Control)
	}

	dom.Kind = post.Kind
	dom.SetControl(post.Control)
	dom.Aux = post.Aux
	dom.Succs = append(dom.Succs[:0], post.Succs...)
	for i := range dom.Succs {
		e := dom.Succs[i]
		e.b.Preds[e.i].b = dom
	}

	for i := range simple.Values {
		simple.Values[i].Block = dom
	}
	for i := range post.Values {
		post.Values[i].Block = dom
	}
	dom.Values = append(dom.Values, simple.Values...)
	dom.Values = append(dom.Values, post.Values...)

	clobberBlock(post)
	clobberBlock(simple)

	f.invalidateCFG()
	return true
}

// is this a BlockPlain with one predecessor?
func isLeafPlain(b *Block) bool {
	return b.Kind == BlockPlain && len(b.Preds) == 1
}

func clobberBlock(b *Block) {
	b.Values = nil
	b.Preds = nil
	b.Succs = nil
	b.Aux = nil
	b.SetControl(nil)
	b.Likely = BranchUnknown
	b.Kind = BlockInvalid
}

func (psess *PackageSession) elimIfElse(f *Func, b *Block) bool {

	if b.Kind != BlockIf || b.Likely != BranchUnknown {
		return false
	}
	yes, no := b.Succs[0].Block(), b.Succs[1].Block()
	if !isLeafPlain(yes) || len(yes.Values) > 1 || !psess.allTrivial(yes) {
		return false
	}
	if !isLeafPlain(no) || len(no.Values) > 1 || !psess.allTrivial(no) {
		return false
	}
	if b.Succs[0].Block().Succs[0].Block() != b.Succs[1].Block().Succs[0].Block() {
		return false
	}

	post := b.Succs[0].Block().Succs[0].Block()
	if len(post.Preds) != 2 || post == b {
		return false
	}
	hasphis := false
	for _, v := range post.Values {
		if v.Op == OpPhi {
			hasphis = true
			if !psess.canCondSelect(v, f.Config.arch) {
				return false
			}
		}
	}
	if !hasphis {
		return false
	}

	if !shouldElimIfElse(no, yes, post, f.Config.arch) {
		return false
	}

	swap := post.Preds[0].Block() != b.Succs[0].Block()
	for _, v := range post.Values {
		if v.Op != OpPhi {
			continue
		}
		v.Op = OpCondSelect
		if swap {
			v.Args[0], v.Args[1] = v.Args[1], v.Args[0]
		}
		v.AddArg(b.Control)
	}

	b.Kind = post.Kind
	b.SetControl(post.Control)
	b.Aux = post.Aux
	b.Succs = append(b.Succs[:0], post.Succs...)
	for i := range b.Succs {
		e := b.Succs[i]
		e.b.Preds[e.i].b = b
	}
	for i := range post.Values {
		post.Values[i].Block = b
	}
	for i := range yes.Values {
		yes.Values[i].Block = b
	}
	for i := range no.Values {
		no.Values[i].Block = b
	}
	b.Values = append(b.Values, yes.Values...)
	b.Values = append(b.Values, no.Values...)
	b.Values = append(b.Values, post.Values...)

	clobberBlock(yes)
	clobberBlock(no)
	clobberBlock(post)

	f.invalidateCFG()
	return true
}

// shouldElimIfElse reports whether estimated cost of eliminating branch
// is lower than threshold.
func shouldElimIfElse(no, yes, post *Block, arch string) bool {
	switch arch {
	default:
		return true
	case "amd64":
		const maxcost = 2
		phi := 0
		other := 0
		for _, v := range post.Values {
			if v.Op == OpPhi {

				phi++
			}
			for _, x := range v.Args {
				if x.Block == no || x.Block == yes {
					other++
				}
			}
		}
		cost := phi * 1
		if phi > 1 {

			cost += other * 1
		}
		return cost < maxcost
	}
}

func (psess *PackageSession) allTrivial(b *Block) bool {

	for _, v := range b.Values {
		if v.Op == OpPhi || isDivMod(v.Op) || v.Type.IsMemory(psess.types) ||
			v.MemoryArg(psess) != nil || psess.opcodeTable[v.Op].hasSideEffects {
			return false
		}
	}
	return true
}

func isDivMod(op Op) bool {
	switch op {
	case OpDiv8, OpDiv8u, OpDiv16, OpDiv16u,
		OpDiv32, OpDiv32u, OpDiv64, OpDiv64u, OpDiv128u,
		OpDiv32F, OpDiv64F,
		OpMod8, OpMod8u, OpMod16, OpMod16u,
		OpMod32, OpMod32u, OpMod64, OpMod64u:
		return true
	default:
		return false
	}
}
