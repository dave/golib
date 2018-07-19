package ssa

// copyelim removes all uses of OpCopy values from f.
// A subsequent deadcode pass is needed to actually remove the copies.
func copyelim(f *Func) {

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			copyelimValue(v)
		}
	}

	for _, b := range f.Blocks {
		if v := b.Control; v != nil && v.Op == OpCopy {
			b.SetControl(v.Args[0])
		}
	}

	for _, name := range f.Names {
		values := f.NamedValues[name]
		for i, v := range values {
			if v.Op == OpCopy {
				values[i] = v.Args[0]
			}
		}
	}
}

// copySource returns the (non-copy) op which is the
// ultimate source of v.  v must be a copy op.
func copySource(v *Value) *Value {
	w := v.Args[0]

	slow := w
	var advance bool
	for w.Op == OpCopy {
		w = w.Args[0]
		if w == slow {
			w.reset(OpUnknown)
			break
		}
		if advance {
			slow = slow.Args[0]
		}
		advance = !advance
	}

	for v != w {
		x := v.Args[0]
		v.SetArg(0, w)
		v = x
	}
	return w
}

// copyelimValue ensures that no args of v are copies.
func copyelimValue(v *Value) {
	for i, a := range v.Args {
		if a.Op == OpCopy {
			v.SetArg(i, copySource(a))
		}
	}
}
