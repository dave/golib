package sys

import "encoding/binary"

// ArchFamily represents a family of one or more related architectures.
// For example, amd64 and amd64p32 are both members of the AMD64 family,
// and ppc64 and ppc64le are both members of the PPC64 family.
type ArchFamily byte

const (
	NoArch ArchFamily = iota
	AMD64
	ARM
	ARM64
	I386
	MIPS
	MIPS64
	PPC64
	S390X
	Wasm
)

// Arch represents an individual architecture.
type Arch struct {
	Name   string
	Family ArchFamily

	ByteOrder binary.ByteOrder

	// PtrSize is the size in bytes of pointers and the
	// predeclared "int", "uint", and "uintptr" types.
	PtrSize int

	// RegSize is the size in bytes of general purpose registers.
	RegSize int

	// MinLC is the minimum length of an instruction code.
	MinLC int
}

// InFamily reports whether a is a member of any of the specified
// architecture families.
func (a *Arch) InFamily(xs ...ArchFamily) bool {
	for _, x := range xs {
		if a.Family == x {
			return true
		}
	}
	return false
}
