package ssa

// phielim eliminates redundant phi values from f.
// A phi is redundant if its arguments are all equal. For
// purposes of counting, ignore the phi itself. Both of
// these phis are redundant:
//   v = phi(x,x,x)
//   v = phi(x,v,x,v)
// We repeat this process to also catch situations like:
//   v = phi(x, phi(x, x), phi(x, v))
// TODO: Can we also simplify cases like:
//   v = phi(v, w, x)
//   w = phi(v, w, x)
// and would that be useful?
func phielim(f *Func) {
	for {
		change := false
		for _, b := range f.Blocks {
			for _, v := range b.Values {
				copyelimValue(v)
				change = phielimValue(v) || change
			}
		}
		if !change {
			break
		}
	}
}

// phielimValue tries to convert the phi v to a copy.
func phielimValue(v *Value) bool {
	if v.Op != OpPhi {
		return false
	}

	// If there are two distinct args of v which
	// are not v itself, then the phi must remain.
	// Otherwise, we can replace it with a copy.
	var w *Value
	for _, x := range v.Args {
		if x == v {
			continue
		}
		if x == w {
			continue
		}
		if w != nil {
			return false
		}
		w = x
	}

	if w == nil {

		return false
	}
	v.Op = OpCopy
	v.SetArgs1(w)
	f := v.Block.Func
	if f.pass.debug > 0 {
		f.Warnl(v.Pos, "eliminated phi")
	}
	return true
}
