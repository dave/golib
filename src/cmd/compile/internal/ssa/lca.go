package ssa

// lcaRange is a data structure that can compute lowest common ancestor queries
// in O(n lg n) precomputed space and O(1) time per query.
type lcaRange struct {
	// Additional information about each block (indexed by block ID).
	blocks []lcaRangeBlock

	// Data structure for range minimum queries.
	// rangeMin[k][i] contains the ID of the minimum depth block
	// in the Euler tour from positions i to i+1<<k-1, inclusive.
	rangeMin [][]ID
}

type lcaRangeBlock struct {
	b          *Block
	parent     ID    // parent in dominator tree.  0 = no parent (entry or unreachable)
	firstChild ID    // first child in dominator tree
	sibling    ID    // next child of parent
	pos        int32 // an index in the Euler tour where this block appears (any one of its occurrences)
	depth      int32 // depth in dominator tree (root=0, its children=1, etc.)
}

func makeLCArange(f *Func) *lcaRange {
	dom := f.Idom()

	blocks := make([]lcaRangeBlock, f.NumBlocks())
	for _, b := range f.Blocks {
		blocks[b.ID].b = b
		if dom[b.ID] == nil {
			continue
		}
		parent := dom[b.ID].ID
		blocks[b.ID].parent = parent
		blocks[b.ID].sibling = blocks[parent].firstChild
		blocks[parent].firstChild = b.ID
	}

	tour := make([]ID, 0, f.NumBlocks()*2-1)
	type queueEntry struct {
		bid ID // block to work on
		cid ID // child we're already working on (0 = haven't started yet)
	}
	q := []queueEntry{{f.Entry.ID, 0}}
	for len(q) > 0 {
		n := len(q) - 1
		bid := q[n].bid
		cid := q[n].cid
		q = q[:n]

		blocks[bid].pos = int32(len(tour))
		tour = append(tour, bid)

		if cid == 0 {

			blocks[bid].depth = blocks[blocks[bid].parent].depth + 1

			cid = blocks[bid].firstChild
		} else {

			cid = blocks[cid].sibling
		}
		if cid != 0 {
			q = append(q, queueEntry{bid, cid}, queueEntry{cid, 0})
		}
	}

	// Compute fast range-minimum query data structure
	var rangeMin [][]ID
	rangeMin = append(rangeMin, tour)
	for logS, s := 1, 2; s < len(tour); logS, s = logS+1, s*2 {
		r := make([]ID, len(tour)-s+1)
		for i := 0; i < len(tour)-s+1; i++ {
			bid := rangeMin[logS-1][i]
			bid2 := rangeMin[logS-1][i+s/2]
			if blocks[bid2].depth < blocks[bid].depth {
				bid = bid2
			}
			r[i] = bid
		}
		rangeMin = append(rangeMin, r)
	}

	return &lcaRange{blocks: blocks, rangeMin: rangeMin}
}

// find returns the lowest common ancestor of a and b.
func (lca *lcaRange) find(a, b *Block) *Block {
	if a == b {
		return a
	}

	p1 := lca.blocks[a.ID].pos
	p2 := lca.blocks[b.ID].pos
	if p1 > p2 {
		p1, p2 = p2, p1
	}

	logS := uint(log2(int64(p2 - p1)))
	bid1 := lca.rangeMin[logS][p1]
	bid2 := lca.rangeMin[logS][p2-1<<logS+1]
	if lca.blocks[bid1].depth < lca.blocks[bid2].depth {
		return lca.blocks[bid1].b
	}
	return lca.blocks[bid2].b
}
