package arm64

const (
	maxAlign  = 32 // max data alignment
	minAlign  = 1  // min data alignment
	funcAlign = 8
)

/* Used by ../internal/ld/dwarf.go */
const (
	dwarfRegSP = 31
	dwarfRegLR = 30
)
