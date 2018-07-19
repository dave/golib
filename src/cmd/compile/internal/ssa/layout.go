package ssa

// layout orders basic blocks in f with the goal of minimizing control flow instructions.
// After this phase returns, the order of f.Blocks matters and is the order
// in which those blocks will appear in the assembly output.
func layout(f *Func) {
	f.Blocks = layoutOrder(f)
}

// Register allocation may use a different order which has constraints
// imposed by the linear-scan algorithm. Note that that f.pass here is
// regalloc, so the switch is conditional on -d=ssa/regalloc/test=N
func layoutRegallocOrder(f *Func) []*Block {

	switch f.pass.test {
	case 0:
		return layoutOrder(f)
	case 1:
		return f.Blocks
	case 2:
		po := f.postorder()
		visitOrder := make([]*Block, len(po))
		for i, b := range po {
			j := len(po) - i - 1
			visitOrder[j] = b
		}
		return visitOrder
	}

	return nil
}

func layoutOrder(f *Func) []*Block {
	order := make([]*Block, 0, f.NumBlocks())
	scheduled := make([]bool, f.NumBlocks())
	idToBlock := make([]*Block, f.NumBlocks())
	indegree := make([]int, f.NumBlocks())
	posdegree := f.newSparseSet(f.NumBlocks())
	defer f.retSparseSet(posdegree)
	zerodegree := f.newSparseSet(f.NumBlocks())
	defer f.retSparseSet(zerodegree)
	exit := f.newSparseSet(f.NumBlocks())
	defer f.retSparseSet(exit)

	for _, b := range f.Blocks {
		idToBlock[b.ID] = b
		if b.Kind == BlockExit {

			exit.add(b.ID)
			continue
		}
		indegree[b.ID] = len(b.Preds)
		if len(b.Preds) == 0 {
			zerodegree.add(b.ID)
		} else {
			posdegree.add(b.ID)
		}
	}

	bid := f.Entry.ID
blockloop:
	for {

		b := idToBlock[bid]
		order = append(order, b)
		scheduled[bid] = true
		if len(order) == len(f.Blocks) {
			break
		}

		for _, e := range b.Succs {
			c := e.b
			indegree[c.ID]--
			if indegree[c.ID] == 0 {
				posdegree.remove(c.ID)
				zerodegree.add(c.ID)
			}
		}

		// Use likely direction if we have it.
		var likely *Block
		switch b.Likely {
		case BranchLikely:
			likely = b.Succs[0].b
		case BranchUnlikely:
			likely = b.Succs[1].b
		}
		if likely != nil && !scheduled[likely.ID] {
			bid = likely.ID
			continue
		}

		bid = 0
		mindegree := f.NumBlocks()
		for _, e := range order[len(order)-1].Succs {
			c := e.b
			if scheduled[c.ID] || c.Kind == BlockExit {
				continue
			}
			if indegree[c.ID] < mindegree {
				mindegree = indegree[c.ID]
				bid = c.ID
			}
		}
		if bid != 0 {
			continue
		}

		for zerodegree.size() > 0 {
			cid := zerodegree.pop()
			if !scheduled[cid] {
				bid = cid
				continue blockloop
			}
		}

		for posdegree.size() > 0 {
			cid := posdegree.pop()
			if !scheduled[cid] {
				bid = cid
				continue blockloop
			}
		}

		for {
			cid := exit.pop()
			if !scheduled[cid] {
				bid = cid
				continue blockloop
			}
		}
	}
	return order
}
