package ssa

// machine-independent optimization
func (psess *PackageSession) opt(f *Func) {
	psess.
		applyRewrite(f, rewriteBlockgeneric, psess.rewriteValuegeneric)
}
