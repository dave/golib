package ssa

import (
	"fmt"
	"strings"
)

type SparseTreeNode struct {
	child   *Block
	sibling *Block
	parent  *Block

	// Every block has 6 numbers associated with it:
	// entry-1, entry, entry+1, exit-1, and exit, exit+1.
	// entry and exit are conceptually the top of the block (phi functions)
	// entry+1 and exit-1 are conceptually the bottom of the block (ordinary defs)
	// entry-1 and exit+1 are conceptually "just before" the block (conditions flowing in)
	//
	// This simplifies life if we wish to query information about x
	// when x is both an input to and output of a block.
	entry, exit int32
}

func (s *SparseTreeNode) String() string {
	return fmt.Sprintf("[%d,%d]", s.entry, s.exit)
}

func (s *SparseTreeNode) Entry() int32 {
	return s.entry
}

func (s *SparseTreeNode) Exit() int32 {
	return s.exit
}

const (
	// When used to lookup up definitions in a sparse tree,
	// these adjustments to a block's entry (+adjust) and
	// exit (-adjust) numbers allow a distinction to be made
	// between assignments (typically branch-dependent
	// conditionals) occurring "before" the block (e.g., as inputs
	// to the block and its phi functions), "within" the block,
	// and "after" the block.
	AdjustBefore = -1 // defined before phi
	AdjustWithin = 0  // defined by phi
	AdjustAfter  = 1  // defined within block
)

// A SparseTree is a tree of Blocks.
// It allows rapid ancestor queries,
// such as whether one block dominates another.
type SparseTree []SparseTreeNode

// newSparseTree creates a SparseTree from a block-to-parent map (array indexed by Block.ID)
func newSparseTree(f *Func, parentOf []*Block) SparseTree {
	t := make(SparseTree, f.NumBlocks())
	for _, b := range f.Blocks {
		n := &t[b.ID]
		if p := parentOf[b.ID]; p != nil {
			n.parent = p
			n.sibling = t[p.ID].child
			t[p.ID].child = b
		}
	}
	t.numberBlock(f.Entry, 1)
	return t
}

// newSparseOrderedTree creates a SparseTree from a block-to-parent map (array indexed by Block.ID)
// children will appear in the reverse of their order in reverseOrder
// in particular, if reverseOrder is a dfs-reversePostOrder, then the root-to-children
// walk of the tree will yield a pre-order.
func newSparseOrderedTree(f *Func, parentOf, reverseOrder []*Block) SparseTree {
	t := make(SparseTree, f.NumBlocks())
	for _, b := range reverseOrder {
		n := &t[b.ID]
		if p := parentOf[b.ID]; p != nil {
			n.parent = p
			n.sibling = t[p.ID].child
			t[p.ID].child = b
		}
	}
	t.numberBlock(f.Entry, 1)
	return t
}

// treestructure provides a string description of the dominator
// tree and flow structure of block b and all blocks that it
// dominates.
func (t SparseTree) treestructure(b *Block) string {
	return t.treestructure1(b, 0)
}
func (t SparseTree) treestructure1(b *Block, i int) string {
	s := "\n" + strings.Repeat("\t", i) + b.String() + "->["
	for i, e := range b.Succs {
		if i > 0 {
			s = s + ","
		}
		s = s + e.b.String()
	}
	s += "]"
	if c0 := t[b.ID].child; c0 != nil {
		s += "("
		for c := c0; c != nil; c = t[c.ID].sibling {
			if c != c0 {
				s += " "
			}
			s += t.treestructure1(c, i+1)
		}
		s += ")"
	}
	return s
}

func (t SparseTree) numberBlock(b *Block, n int32) int32 {

	n++
	t[b.ID].entry = n

	n += 2
	for c := t[b.ID].child; c != nil; c = t[c.ID].sibling {
		n = t.numberBlock(c, n)
	}

	n++
	t[b.ID].exit = n

	return n + 2
}

// Sibling returns a sibling of x in the dominator tree (i.e.,
// a node with the same immediate dominator) or nil if there
// are no remaining siblings in the arbitrary but repeatable
// order chosen. Because the Child-Sibling order is used
// to assign entry and exit numbers in the treewalk, those
// numbers are also consistent with this order (i.e.,
// Sibling(x) has entry number larger than x's exit number).
func (t SparseTree) Sibling(x *Block) *Block {
	return t[x.ID].sibling
}

// Child returns a child of x in the dominator tree, or
// nil if there are none. The choice of first child is
// arbitrary but repeatable.
func (t SparseTree) Child(x *Block) *Block {
	return t[x.ID].child
}

// isAncestorEq reports whether x is an ancestor of or equal to y.
func (t SparseTree) isAncestorEq(x, y *Block) bool {
	if x == y {
		return true
	}
	xx := &t[x.ID]
	yy := &t[y.ID]
	return xx.entry <= yy.entry && yy.exit <= xx.exit
}

// isAncestor reports whether x is a strict ancestor of y.
func (t SparseTree) isAncestor(x, y *Block) bool {
	if x == y {
		return false
	}
	xx := &t[x.ID]
	yy := &t[y.ID]
	return xx.entry < yy.entry && yy.exit < xx.exit
}

// domorder returns a value for dominator-oriented sorting.
// Block domination does not provide a total ordering,
// but domorder two has useful properties.
// (1) If domorder(x) > domorder(y) then x does not dominate y.
// (2) If domorder(x) < domorder(y) and domorder(y) < domorder(z) and x does not dominate y,
//     then x does not dominate z.
// Property (1) means that blocks sorted by domorder always have a maximal dominant block first.
// Property (2) allows searches for dominated blocks to exit early.
func (t SparseTree) domorder(x *Block) int32 {

	return t[x.ID].entry
}
