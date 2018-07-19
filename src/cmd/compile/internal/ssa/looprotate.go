package ssa

// loopRotate converts loops with a check-loop-condition-at-beginning
// to loops with a check-loop-condition-at-end.
// This helps loops avoid extra unnecessary jumps.
//
//   loop:
//     CMPQ ...
//     JGE exit
//     ...
//     JMP loop
//   exit:
//
//    JMP entry
//  loop:
//    ...
//  entry:
//    CMPQ ...
//    JLT loop
func (psess *PackageSession) loopRotate(f *Func) {
	loopnest := f.loopnest(psess)
	if loopnest.hasIrreducible {
		return
	}
	if len(loopnest.loops) == 0 {
		return
	}

	idToIdx := make([]int, f.NumBlocks())
	for i, b := range f.Blocks {
		idToIdx[b.ID] = i
	}

	move := map[ID]struct{}{}

	after := map[ID][]*Block{}

	for _, loop := range loopnest.loops {
		b := loop.header
		var p *Block // b's in-loop predecessor
		for _, e := range b.Preds {
			if e.b.Kind != BlockPlain {
				continue
			}
			if loopnest.b2l[e.b.ID] != loop {
				continue
			}
			p = e.b
		}
		if p == nil || p == b {
			continue
		}
		after[p.ID] = []*Block{b}
		for {
			nextIdx := idToIdx[b.ID] + 1
			if nextIdx >= len(f.Blocks) {
				break
			}
			nextb := f.Blocks[nextIdx]
			if nextb == p {
				break
			}
			if loopnest.b2l[nextb.ID] != loop {
				break
			}
			after[p.ID] = append(after[p.ID], nextb)
			b = nextb
		}

		for _, b := range after[p.ID] {
			move[b.ID] = struct{}{}
		}
	}

	j := 0
	for i, b := range f.Blocks {
		if _, ok := move[b.ID]; ok {
			continue
		}
		f.Blocks[j] = b
		j++
		for _, a := range after[b.ID] {
			if j > i {
				f.Fatalf("head before tail in loop %s", b)
			}
			f.Blocks[j] = a
			j++
		}
	}
	if j != len(f.Blocks) {
		f.Fatalf("bad reordering in looprotate")
	}
}
