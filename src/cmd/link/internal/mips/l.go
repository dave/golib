package mips

const (
	MaxAlign  = 32 // max data alignment
	MinAlign  = 1  // min data alignment
	FuncAlign = 4
)

/* Used by ../internal/ld/dwarf.go */
const (
	DWARFREGSP = 29
	DWARFREGLR = 31
)
