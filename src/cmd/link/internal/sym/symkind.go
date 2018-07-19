package sym

// A SymKind describes the kind of memory represented by a symbol.
type SymKind uint8

// Defined SymKind values.
//
// TODO(rsc): Give idiomatic Go names.
//go:generate stringer -type=SymKind
const (
	Sxxx SymKind = iota
	STEXT
	SELFRXSECT

	// Read-only sections.
	STYPE
	SSTRING
	SGOSTRING
	SGOFUNC
	SGCBITS
	SRODATA
	SFUNCTAB

	SELFROSECT
	SMACHOPLT

	// Read-only sections with relocations.
	//
	// Types STYPE-SFUNCTAB above are written to the .rodata section by default.
	// When linking a shared object, some conceptually "read only" types need to
	// be written to by relocations and putting them in a section called
	// ".rodata" interacts poorly with the system linkers. The GNU linkers
	// support this situation by arranging for sections of the name
	// ".data.rel.ro.XXX" to be mprotected read only by the dynamic linker after
	// relocations have applied, so when the Go linker is creating a shared
	// object it checks all objects of the above types and bumps any object that
	// has a relocation to it to the corresponding type below, which are then
	// written to sections with appropriate magic names.
	STYPERELRO
	SSTRINGRELRO
	SGOSTRINGRELRO
	SGOFUNCRELRO
	SGCBITSRELRO
	SRODATARELRO
	SFUNCTABRELRO

	// Part of .data.rel.ro if it exists, otherwise part of .rodata.
	STYPELINK
	SITABLINK
	SSYMTAB
	SPCLNTAB

	// Writable sections.
	SELFSECT
	SMACHO
	SMACHOGOT
	SWINDOWS
	SELFGOT
	SNOPTRDATA
	SINITARR
	SDATA
	SBSS
	SNOPTRBSS
	STLSBSS
	SXREF
	SMACHOSYMSTR
	SMACHOSYMTAB
	SMACHOINDIRECTPLT
	SMACHOINDIRECTGOT
	SFILEPATH
	SCONST
	SDYNIMPORT
	SHOSTOBJ

	// Sections for debugging information
	SDWARFSECT
	SDWARFINFO
	SDWARFRANGE
	SDWARFLOC
	SDWARFMISC // Not really a section; informs/affects other DWARF section generation
)

// AbiSymKindToSymKind maps values read from object files (which are
// of type cmd/internal/objabi.SymKind) to values of type SymKind.

// ReadOnly are the symbol kinds that form read-only sections. In some
// cases, if they will require relocations, they are transformed into
// rel-ro sections using relROMap.

// RelROMap describes the transformation of read-only symbols to rel-ro
// symbols.
