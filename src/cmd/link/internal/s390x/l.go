package s390x

const (
	maxAlign  = 32 // max data alignment
	minAlign  = 2  // min data alignment
	funcAlign = 16
)

/* Used by ../internal/ld/dwarf.go */
const (
	dwarfRegSP = 15
	dwarfRegLR = 14
)
