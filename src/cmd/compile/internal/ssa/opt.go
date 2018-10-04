// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// machine-independent optimization
func (pstate *PackageState) opt(f *Func) {
	pstate.applyRewrite(f, rewriteBlockgeneric, pstate.rewriteValuegeneric)
}
