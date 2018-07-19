package ssa

// stackframe calls back into the frontend to assign frame offsets.
func stackframe(f *Func) {
	f.fe.AllocFrame(f)
}
