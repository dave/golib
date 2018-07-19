package ssa

import (
	"github.com/dave/golib/src/cmd/internal/src"
)

// fuse simplifies control flow by joining basic blocks.
func fuse(f *Func) {
	for changed := true; changed; {
		changed = false

		for i := len(f.Blocks) - 1; i >= 0; i-- {
			b := f.Blocks[i]
			changed = fuseBlockIf(b) || changed
			changed = fuseBlockPlain(b) || changed
		}
	}
}

// fuseBlockIf handles the following cases where s0 and s1 are empty blocks.
//
//   b        b        b      b
//  / \      | \      / |    | |
// s0  s1    |  s1   s0 |    | |
//  \ /      | /      \ |    | |
//   ss      ss        ss     ss
//
// If all Phi ops in ss have identical variables for slots corresponding to
// s0, s1 and b then the branch can be dropped.
// This optimization often comes up in switch statements with multiple
// expressions in a case clause:
//   switch n {
//     case 1,2,3: return 4
//   }
// TODO: If ss doesn't contain any OpPhis, are s0 and s1 dead code anyway.
func fuseBlockIf(b *Block) bool {
	if b.Kind != BlockIf {
		return false
	}

	var ss0, ss1 *Block
	s0 := b.Succs[0].b
	i0 := b.Succs[0].i
	if s0.Kind != BlockPlain || len(s0.Preds) != 1 || len(s0.Values) != 0 {
		s0, ss0 = b, s0
	} else {
		ss0 = s0.Succs[0].b
		i0 = s0.Succs[0].i
	}
	s1 := b.Succs[1].b
	i1 := b.Succs[1].i
	if s1.Kind != BlockPlain || len(s1.Preds) != 1 || len(s1.Values) != 0 {
		s1, ss1 = b, s1
	} else {
		ss1 = s1.Succs[0].b
		i1 = s1.Succs[0].i
	}

	if ss0 != ss1 {
		return false
	}
	ss := ss0

	for _, v := range ss.Values {
		if v.Op == OpPhi && v.Uses > 0 && v.Args[i0] != v.Args[i1] {
			return false
		}
	}

	if s0 != b && s1 != b {

		b.Succs[0] = Edge{ss, i0}
		ss.Preds[i0] = Edge{b, 0}
		b.removeEdge(1)
		s1.removeEdge(0)
	} else if s0 != b {
		b.removeEdge(0)
		s0.removeEdge(0)
	} else if s1 != b {
		b.removeEdge(1)
		s1.removeEdge(0)
	} else {
		b.removeEdge(1)
	}
	b.Kind = BlockPlain
	b.Likely = BranchUnknown
	b.SetControl(nil)

	if s0 != b {
		s0.Kind = BlockInvalid
		s0.Values = nil
		s0.Succs = nil
		s0.Preds = nil
	}
	if s1 != b {
		s1.Kind = BlockInvalid
		s1.Values = nil
		s1.Succs = nil
		s1.Preds = nil
	}
	return true
}

func fuseBlockPlain(b *Block) bool {
	if b.Kind != BlockPlain {
		return false
	}

	c := b.Succs[0].b
	if len(c.Preds) != 1 {
		return false
	}

	if b.Pos.IsStmt() == src.PosIsStmt {
		l := b.Pos.Line()
		for _, v := range c.Values {
			if v.Pos.IsStmt() == src.PosNotStmt {
				continue
			}
			if l == v.Pos.Line() {
				v.Pos = v.Pos.WithIsStmt()
				l = 0
				break
			}
		}
		if l != 0 && c.Pos.Line() == l {
			c.Pos = c.Pos.WithIsStmt()
		}
	}

	for _, v := range b.Values {
		v.Block = c
	}

	if cap(c.Values) >= cap(b.Values) || len(b.Values) <= len(b.valstorage) {
		bl := len(b.Values)
		cl := len(c.Values)
		var t []*Value // construct t = b.Values followed-by c.Values, but with attention to allocation.
		if cap(c.Values) < bl+cl {

			t = make([]*Value, bl+cl)
		} else {

			t = c.Values[0 : bl+cl]
		}
		copy(t[bl:], c.Values)
		c.Values = t
		copy(c.Values, b.Values)
	} else {
		c.Values = append(b.Values, c.Values...)
	}

	c.predstorage[0] = Edge{}
	if len(b.Preds) > len(b.predstorage) {
		c.Preds = b.Preds
	} else {
		c.Preds = append(c.predstorage[:0], b.Preds...)
	}
	for i, e := range c.Preds {
		p := e.b
		p.Succs[e.i] = Edge{c, i}
	}
	f := b.Func
	if f.Entry == b {
		f.Entry = c
	}
	f.invalidateCFG()

	b.Kind = BlockInvalid
	b.Values = nil
	b.Preds = nil
	b.Succs = nil
	return true
}
