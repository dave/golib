package ssa

import (
	"fmt"
)

type loop struct {
	header *Block // The header node of this (reducible) loop
	outer  *loop  // loop containing this loop

	// By default, children, exits, and depth are not initialized.
	children []*loop  // loops nested directly within this loop. Initialized by assembleChildren().
	exits    []*Block // exits records blocks reached by exits from this loop. Initialized by findExits().

	// Next three fields used by regalloc and/or
	// aid in computation of inner-ness and list of blocks.
	nBlocks int32 // Number of blocks in this loop but not within inner loops
	depth   int16 // Nesting depth of the loop; 1 is outermost. Initialized by calculateDepths().
	isInner bool  // True if never discovered to contain a loop

	// register allocation uses this.
	containsUnavoidableCall bool // True if all paths through the loop have a call
}

// outerinner records that outer contains inner
func (sdom SparseTree) outerinner(outer, inner *loop) {

	oldouter := inner.outer
	for oldouter != nil && sdom.isAncestor(outer.header, oldouter.header) {
		inner = oldouter
		oldouter = inner.outer
	}
	if outer == oldouter {
		return
	}
	if oldouter != nil {
		sdom.outerinner(oldouter, outer)
	}

	inner.outer = outer
	outer.isInner = false
}

func (psess *PackageSession) checkContainsCall(bb *Block) bool {
	if bb.Kind == BlockDefer {
		return true
	}
	for _, v := range bb.Values {
		if psess.opcodeTable[v.Op].call {
			return true
		}
	}
	return false
}

type loopnest struct {
	f              *Func
	b2l            []*loop
	po             []*Block
	sdom           SparseTree
	loops          []*loop
	hasIrreducible bool // TODO current treatment of irreducible loops is very flaky, if accurate loops are needed, must punt at function level.

	// Record which of the lazily initialized fields have actually been initialized.
	initializedChildren, initializedDepth, initializedExits bool
}

func min8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func max8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

const (
	blDEFAULT = 0
	blMin     = blDEFAULT
	blCALL    = 1
	blRET     = 2
	blEXIT    = 3
)

func describePredictionAgrees(b *Block, prediction BranchPrediction) string {
	s := ""
	if prediction == b.Likely {
		s = " (agrees with previous)"
	} else if b.Likely != BranchUnknown {
		s = " (disagrees with previous, ignored)"
	}
	return s
}

func (psess *PackageSession) describeBranchPrediction(f *Func, b *Block, likely, not int8, prediction BranchPrediction) {
	f.Warnl(b.Pos, "Branch prediction rule %s < %s%s", psess.
		bllikelies[likely-blMin], psess.bllikelies[not-blMin], describePredictionAgrees(b, prediction))
}

func (psess *PackageSession) likelyadjust(f *Func) {

	certain := make([]int8, f.NumBlocks())
	local := make([]int8, f.NumBlocks())

	po := f.postorder()
	nest := f.loopnest(psess)
	b2l := nest.b2l

	for _, b := range po {
		switch b.Kind {
		case BlockExit:

			local[b.ID] = blEXIT
			certain[b.ID] = blEXIT

		case BlockRet, BlockRetJmp:
			local[b.ID] = blRET
			certain[b.ID] = blRET

		case BlockDefer:
			local[b.ID] = blCALL
			certain[b.ID] = max8(blCALL, certain[b.Succs[0].b.ID])

		default:
			if len(b.Succs) == 1 {
				certain[b.ID] = certain[b.Succs[0].b.ID]
			} else if len(b.Succs) == 2 {

				b0 := b.Succs[0].b.ID
				b1 := b.Succs[1].b.ID
				certain[b.ID] = min8(certain[b0], certain[b1])

				l := b2l[b.ID]
				l0 := b2l[b0]
				l1 := b2l[b1]

				prediction := b.Likely

				if l != nil && l0 != l1 {
					noprediction := false
					switch {

					case l1 == nil:
						prediction = BranchLikely
					case l0 == nil:
						prediction = BranchUnlikely

					case l == l0:
						prediction = BranchLikely
					case l == l1:
						prediction = BranchUnlikely
					default:
						noprediction = true
					}
					if f.pass.debug > 0 && !noprediction {
						f.Warnl(b.Pos, "Branch prediction rule stay in loop%s",
							describePredictionAgrees(b, prediction))
					}

				} else {

					if certain[b1] > certain[b0] {
						prediction = BranchLikely
						if f.pass.debug > 0 {
							psess.
								describeBranchPrediction(f, b, certain[b0], certain[b1], prediction)
						}
					} else if certain[b0] > certain[b1] {
						prediction = BranchUnlikely
						if f.pass.debug > 0 {
							psess.
								describeBranchPrediction(f, b, certain[b1], certain[b0], prediction)
						}
					} else if local[b1] > local[b0] {
						prediction = BranchLikely
						if f.pass.debug > 0 {
							psess.
								describeBranchPrediction(f, b, local[b0], local[b1], prediction)
						}
					} else if local[b0] > local[b1] {
						prediction = BranchUnlikely
						if f.pass.debug > 0 {
							psess.
								describeBranchPrediction(f, b, local[b1], local[b0], prediction)
						}
					}
				}
				if b.Likely != prediction {
					if b.Likely == BranchUnknown {
						b.Likely = prediction
					}
				}
			}

			for _, v := range b.Values {
				if psess.opcodeTable[v.Op].call {
					local[b.ID] = blCALL
					certain[b.ID] = max8(blCALL, certain[b.Succs[0].b.ID])
				}
			}
		}
		if f.pass.debug > 2 {
			f.Warnl(b.Pos, "BP: Block %s, local=%s, certain=%s", b, psess.bllikelies[local[b.ID]-blMin], psess.bllikelies[certain[b.ID]-blMin])
		}

	}
}

func (l *loop) String() string {
	return fmt.Sprintf("hdr:%s", l.header)
}

func (l *loop) LongString() string {
	i := ""
	o := ""
	if l.isInner {
		i = ", INNER"
	}
	if l.outer != nil {
		o = ", o=" + l.outer.header.String()
	}
	return fmt.Sprintf("hdr:%s%s%s", l.header, i, o)
}

func (l *loop) isWithinOrEq(ll *loop) bool {
	if ll == nil {
		return true
	}
	for ; l != nil; l = l.outer {
		if l == ll {
			return true
		}
	}
	return false
}

// nearestOuterLoop returns the outer loop of loop most nearly
// containing block b; the header must dominate b.  loop itself
// is assumed to not be that loop. For acceptable performance,
// we're relying on loop nests to not be terribly deep.
func (l *loop) nearestOuterLoop(sdom SparseTree, b *Block) *loop {
	var o *loop
	for o = l.outer; o != nil && !sdom.isAncestorEq(o.header, b); o = o.outer {
	}
	return o
}

func (psess *PackageSession) loopnestfor(f *Func) *loopnest {
	po := f.postorder()
	sdom := f.sdom()
	b2l := make([]*loop, f.NumBlocks())
	loops := make([]*loop, 0)
	visited := make([]bool, f.NumBlocks())
	sawIrred := false

	if f.pass.debug > 2 {
		fmt.Printf("loop finding in %s\n", f.Name)
	}

	for _, b := range po {
		if f.pass != nil && f.pass.debug > 3 {
			fmt.Printf("loop finding at %s\n", b)
		}

		var innermost *loop // innermost header reachable from this block

		for _, e := range b.Succs {
			bb := e.b
			l := b2l[bb.ID]

			if sdom.isAncestorEq(bb, b) {
				if f.pass != nil && f.pass.debug > 4 {
					fmt.Printf("loop finding    succ %s of %s is header\n", bb.String(), b.String())
				}
				if l == nil {
					l = &loop{header: bb, isInner: true}
					loops = append(loops, l)
					b2l[bb.ID] = l
				}
			} else if !visited[bb.ID] {
				sawIrred = true
				if f.pass != nil && f.pass.debug > 4 {
					fmt.Printf("loop finding    succ %s of %s is IRRED, in %s\n", bb.String(), b.String(), f.Name)
				}
			} else if l != nil {

				if !sdom.isAncestorEq(l.header, b) {
					l = l.nearestOuterLoop(sdom, b)
				}
				if f.pass != nil && f.pass.debug > 4 {
					if l == nil {
						fmt.Printf("loop finding    succ %s of %s has no loop\n", bb.String(), b.String())
					} else {
						fmt.Printf("loop finding    succ %s of %s provides loop with header %s\n", bb.String(), b.String(), l.header.String())
					}
				}
			} else {
				if f.pass != nil && f.pass.debug > 4 {
					fmt.Printf("loop finding    succ %s of %s has no loop\n", bb.String(), b.String())
				}

			}

			if l == nil || innermost == l {
				continue
			}

			if innermost == nil {
				innermost = l
				continue
			}

			if sdom.isAncestor(innermost.header, l.header) {
				sdom.outerinner(innermost, l)
				innermost = l
			} else if sdom.isAncestor(l.header, innermost.header) {
				sdom.outerinner(l, innermost)
			}
		}

		if innermost != nil {
			b2l[b.ID] = innermost
			innermost.nBlocks++
		}
		visited[b.ID] = true
	}

	ln := &loopnest{f: f, b2l: b2l, po: po, sdom: sdom, loops: loops, hasIrreducible: sawIrred}

	dominatedByCall := make([]bool, f.NumBlocks())
	for _, b := range po {
		if psess.checkContainsCall(b) {
			dominatedByCall[b.ID] = true
		}
	}

	for _, l := range loops {

		if dominatedByCall[l.header.ID] {
			l.containsUnavoidableCall = true
			continue
		}
		callfreepath := false
		tovisit := make([]*Block, 0, len(l.header.Succs))

		for _, s := range l.header.Succs {
			nb := s.Block()

			if !l.iterationEnd(nb, b2l) {
				tovisit = append(tovisit, nb)
			}
		}
		for len(tovisit) > 0 {
			cur := tovisit[len(tovisit)-1]
			tovisit = tovisit[:len(tovisit)-1]
			if dominatedByCall[cur.ID] {
				continue
			}

			dominatedByCall[cur.ID] = true
			for _, s := range cur.Succs {
				nb := s.Block()
				if l.iterationEnd(nb, b2l) {
					callfreepath = true
				}
				if !dominatedByCall[nb.ID] {
					tovisit = append(tovisit, nb)
				}

			}
			if callfreepath {
				break
			}
		}
		if !callfreepath {
			l.containsUnavoidableCall = true
		}
	}

	if f.pass != nil && f.pass.stats > 0 && len(loops) > 0 {
		ln.assembleChildren()
		ln.calculateDepths()
		ln.findExits()

		for _, l := range loops {
			x := len(l.exits)
			cf := 0
			if !l.containsUnavoidableCall {
				cf = 1
			}
			inner := 0
			if l.isInner {
				inner++
			}

			f.LogStat("loopstats:",
				l.depth, "depth", x, "exits",
				inner, "is_inner", cf, "always_calls", l.nBlocks, "n_blocks")
		}
	}

	if f.pass != nil && f.pass.debug > 1 && len(loops) > 0 {
		fmt.Printf("Loops in %s:\n", f.Name)
		for _, l := range loops {
			fmt.Printf("%s, b=", l.LongString())
			for _, b := range f.Blocks {
				if b2l[b.ID] == l {
					fmt.Printf(" %s", b)
				}
			}
			fmt.Print("\n")
		}
		fmt.Printf("Nonloop blocks in %s:", f.Name)
		for _, b := range f.Blocks {
			if b2l[b.ID] == nil {
				fmt.Printf(" %s", b)
			}
		}
		fmt.Print("\n")
	}
	return ln
}

// assembleChildren initializes the children field of each
// loop in the nest.  Loop A is a child of loop B if A is
// directly nested within B (based on the reducible-loops
// detection above)
func (ln *loopnest) assembleChildren() {
	if ln.initializedChildren {
		return
	}
	for _, l := range ln.loops {
		if l.outer != nil {
			l.outer.children = append(l.outer.children, l)
		}
	}
	ln.initializedChildren = true
}

// calculateDepths uses the children field of loops
// to determine the nesting depth (outer=1) of each
// loop.  This is helpful for finding exit edges.
func (ln *loopnest) calculateDepths() {
	if ln.initializedDepth {
		return
	}
	ln.assembleChildren()
	for _, l := range ln.loops {
		if l.outer == nil {
			l.setDepth(1)
		}
	}
	ln.initializedDepth = true
}

// findExits uses loop depth information to find the
// exits from a loop.
func (ln *loopnest) findExits() {
	if ln.initializedExits {
		return
	}
	ln.calculateDepths()
	b2l := ln.b2l
	for _, b := range ln.po {
		l := b2l[b.ID]
		if l != nil && len(b.Succs) == 2 {
			sl := b2l[b.Succs[0].b.ID]
			if recordIfExit(l, sl, b.Succs[0].b) {
				continue
			}
			sl = b2l[b.Succs[1].b.ID]
			if recordIfExit(l, sl, b.Succs[1].b) {
				continue
			}
		}
	}
	ln.initializedExits = true
}

// depth returns the loop nesting level of block b.
func (ln *loopnest) depth(b ID) int16 {
	if l := ln.b2l[b]; l != nil {
		return l.depth
	}
	return 0
}

// recordIfExit checks sl (the loop containing b) to see if it
// is outside of loop l, and if so, records b as an exit block
// from l and returns true.
func recordIfExit(l, sl *loop, b *Block) bool {
	if sl != l {
		if sl == nil || sl.depth <= l.depth {
			l.exits = append(l.exits, b)
			return true
		}

		for sl.depth > l.depth {
			sl = sl.outer
		}
		if sl != l {
			l.exits = append(l.exits, b)
			return true
		}
	}
	return false
}

func (l *loop) setDepth(d int16) {
	l.depth = d
	for _, c := range l.children {
		c.setDepth(d + 1)
	}
}

// iterationEnd checks if block b ends iteration of loop l.
// Ending iteration means either escaping to outer loop/code or
// going back to header
func (l *loop) iterationEnd(b *Block, b2l []*loop) bool {
	return b == l.header || b2l[b.ID] == nil || (b2l[b.ID] != l && b2l[b.ID].depth <= l.depth)
}
