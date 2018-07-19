package ssa

import "fmt"

const (
	rankLeaf rbrank = 1
	rankZero rbrank = 0
)

type rbrank int8

// RBTint32 is a red-black tree with data stored at internal nodes,
// following Tarjan, Data Structures and Network Algorithms,
// pp 48-52, using explicit rank instead of red and black.
// Deletion is not yet implemented because it is not yet needed.
// Extra operations glb, lub, glbEq, lubEq are provided for
// use in sparse lookup algorithms.
type RBTint32 struct {
	root *node32
}

func (t *RBTint32) String() string {
	if t.root == nil {
		return "[]"
	}
	return "[" + t.root.String() + "]"
}

func (t *node32) String() string {
	s := ""
	if t.left != nil {
		s = t.left.String() + " "
	}
	s = s + fmt.Sprintf("k=%d,d=%v", t.key, t.data)
	if t.right != nil {
		s = s + " " + t.right.String()
	}
	return s
}

type node32 struct {
	// Standard conventions hold for left = smaller, right = larger
	left, right, parent *node32
	data                interface{}
	key                 int32
	rank                rbrank // From Tarjan pp 48-49:

}

// makeNode returns a new leaf node with the given key and nil data.
func (t *RBTint32) makeNode(key int32) *node32 {
	return &node32{key: key, rank: rankLeaf}
}

// IsEmpty reports whether t is empty.
func (t *RBTint32) IsEmpty() bool {
	return t.root == nil
}

// IsSingle reports whether t is a singleton (leaf).
func (t *RBTint32) IsSingle() bool {
	return t.root != nil && t.root.isLeaf()
}

// VisitInOrder applies f to the key and data pairs in t,
// with keys ordered from smallest to largest.
func (t *RBTint32) VisitInOrder(f func(int32, interface{})) {
	if t.root == nil {
		return
	}
	t.root.visitInOrder(f)
}

func (n *node32) Data() interface{} {
	if n == nil {
		return nil
	}
	return n.data
}

func (n *node32) keyAndData() (k int32, d interface{}) {
	if n == nil {
		k = 0
		d = nil
	} else {
		k = n.key
		d = n.data
	}
	return
}

func (n *node32) Rank() rbrank {
	if n == nil {
		return 0
	}
	return n.rank
}

// Find returns the data associated with key in the tree, or
// nil if key is not in the tree.
func (t *RBTint32) Find(key int32) interface{} {
	return t.root.find(key).Data()
}

// Insert adds key to the tree and associates key with data.
// If key was already in the tree, it updates the associated data.
// Insert returns the previous data associated with key,
// or nil if key was not present.
// Insert panics if data is nil.
func (t *RBTint32) Insert(key int32, data interface{}) interface{} {
	if data == nil {
		panic("Cannot insert nil data into tree")
	}
	n := t.root
	var newroot *node32
	if n == nil {
		n = t.makeNode(key)
		newroot = n
	} else {
		newroot, n = n.insert(key, t)
	}
	r := n.data
	n.data = data
	t.root = newroot
	return r
}

// Min returns the minimum element of t and its associated data.
// If t is empty, then (0, nil) is returned.
func (t *RBTint32) Min() (k int32, d interface{}) {
	return t.root.min().keyAndData()
}

// Max returns the maximum element of t and its associated data.
// If t is empty, then (0, nil) is returned.
func (t *RBTint32) Max() (k int32, d interface{}) {
	return t.root.max().keyAndData()
}

// Glb returns the greatest-lower-bound-exclusive of x and its associated
// data.  If x has no glb in the tree, then (0, nil) is returned.
func (t *RBTint32) Glb(x int32) (k int32, d interface{}) {
	return t.root.glb(x, false).keyAndData()
}

// GlbEq returns the greatest-lower-bound-inclusive of x and its associated
// data.  If x has no glbEQ in the tree, then (0, nil) is returned.
func (t *RBTint32) GlbEq(x int32) (k int32, d interface{}) {
	return t.root.glb(x, true).keyAndData()
}

// Lub returns the least-upper-bound-exclusive of x and its associated
// data.  If x has no lub in the tree, then (0, nil) is returned.
func (t *RBTint32) Lub(x int32) (k int32, d interface{}) {
	return t.root.lub(x, false).keyAndData()
}

// LubEq returns the least-upper-bound-inclusive of x and its associated
// data.  If x has no lubEq in the tree, then (0, nil) is returned.
func (t *RBTint32) LubEq(x int32) (k int32, d interface{}) {
	return t.root.lub(x, true).keyAndData()
}

func (t *node32) isLeaf() bool {
	return t.left == nil && t.right == nil
}

func (t *node32) visitInOrder(f func(int32, interface{})) {
	if t.left != nil {
		t.left.visitInOrder(f)
	}
	f(t.key, t.data)
	if t.right != nil {
		t.right.visitInOrder(f)
	}
}

func (t *node32) maxChildRank() rbrank {
	if t.left == nil {
		if t.right == nil {
			return rankZero
		}
		return t.right.rank
	}
	if t.right == nil {
		return t.left.rank
	}
	if t.right.rank > t.left.rank {
		return t.right.rank
	}
	return t.left.rank
}

func (t *node32) minChildRank() rbrank {
	if t.left == nil || t.right == nil {
		return rankZero
	}
	if t.right.rank < t.left.rank {
		return t.right.rank
	}
	return t.left.rank
}

func (t *node32) find(key int32) *node32 {
	for t != nil {
		if key < t.key {
			t = t.left
		} else if key > t.key {
			t = t.right
		} else {
			return t
		}
	}
	return nil
}

func (t *node32) min() *node32 {
	if t == nil {
		return t
	}
	for t.left != nil {
		t = t.left
	}
	return t
}

func (t *node32) max() *node32 {
	if t == nil {
		return t
	}
	for t.right != nil {
		t = t.right
	}
	return t
}

func (t *node32) glb(key int32, allow_eq bool) *node32 {
	var best *node32
	for t != nil {
		if key <= t.key {
			if key == t.key && allow_eq {
				return t
			}

			t = t.left
		} else {

			best = t
			t = t.right
		}
	}
	return best
}

func (t *node32) lub(key int32, allow_eq bool) *node32 {
	var best *node32
	for t != nil {
		if key >= t.key {
			if key == t.key && allow_eq {
				return t
			}

			t = t.right
		} else {

			best = t
			t = t.left
		}
	}
	return best
}

func (t *node32) insert(x int32, w *RBTint32) (newroot, newnode *node32) {

	newroot = t
	newnode = t
	if x == t.key {
		return
	}
	if x < t.key {
		if t.left == nil {
			n := w.makeNode(x)
			n.parent = t
			t.left = n
			newnode = n
			return
		}
		var new_l *node32
		new_l, newnode = t.left.insert(x, w)
		t.left = new_l
		new_l.parent = t
		newrank := 1 + new_l.maxChildRank()
		if newrank > t.rank {
			if newrank > 1+t.right.Rank() {
				if new_l.left.Rank() < new_l.right.Rank() {

					t.left = new_l.rightToRoot()
				}
				newroot = t.leftToRoot()
				return
			} else {
				t.rank = newrank
			}
		}
	} else {
		if t.right == nil {
			n := w.makeNode(x)
			n.parent = t
			t.right = n
			newnode = n
			return
		}
		var new_r *node32
		new_r, newnode = t.right.insert(x, w)
		t.right = new_r
		new_r.parent = t
		newrank := 1 + new_r.maxChildRank()
		if newrank > t.rank {
			if newrank > 1+t.left.Rank() {
				if new_r.right.Rank() < new_r.left.Rank() {

					t.right = new_r.leftToRoot()
				}
				newroot = t.rightToRoot()
				return
			} else {
				t.rank = newrank
			}
		}
	}
	return
}

func (t *node32) rightToRoot() *node32 {

	right := t.right
	rl := right.left
	right.parent = t.parent
	right.left = t
	t.parent = right

	t.right = rl
	if rl != nil {
		rl.parent = t
	}
	return right
}

func (t *node32) leftToRoot() *node32 {

	left := t.left
	lr := left.right
	left.parent = t.parent
	left.right = t
	t.parent = left

	t.left = lr
	if lr != nil {
		lr.parent = t
	}
	return left
}

// next returns the successor of t in a left-to-right
// walk of the tree in which t is embedded.
func (t *node32) next() *node32 {

	r := t.right
	if r != nil {
		return r.min()
	}

	p := t.parent
	for p != nil {
		if p.left == t {
			return p
		}
		t = p
		p = t.parent
	}
	return nil
}

// prev returns the predecessor of t in a left-to-right
// walk of the tree in which t is embedded.
func (t *node32) prev() *node32 {

	l := t.left
	if l != nil {
		return l.max()
	}

	p := t.parent
	for p != nil {
		if p.right == t {
			return p
		}
		t = p
		p = t.parent
	}
	return nil
}
