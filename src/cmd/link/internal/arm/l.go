package arm

const (
	maxAlign  = 8 // max data alignment
	minAlign  = 1 // min data alignment
	funcAlign = 4 // single-instruction alignment
)

/* Used by ../internal/ld/dwarf.go */
const (
	dwarfRegSP = 13
	dwarfRegLR = 14
)
