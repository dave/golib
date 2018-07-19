package objabi

// A SymKind describes the kind of memory represented by a symbol.
type SymKind uint8

// Defined SymKind values.
// These are used to index into cmd/link/internal/sym/AbiSymKindToSymKind
//
// TODO(rsc): Give idiomatic Go names.
//go:generate stringer -type=SymKind
const (
	// An otherwise invalid zero value for the type
	Sxxx SymKind = iota
	// Executable instructions
	STEXT
	// Read only static data
	SRODATA
	// Static data that does not contain any pointers
	SNOPTRDATA
	// Static data
	SDATA
	// Statically data that is initially all 0s
	SBSS
	// Statically data that is initially all 0s and does not contain pointers
	SNOPTRBSS
	// Thread-local data that is initially all 0s
	STLSBSS
	// Debugging data
	SDWARFINFO
	SDWARFRANGE
	SDWARFLOC
	SDWARFMISC
)
