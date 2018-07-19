package sym

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
)

type PackageSession struct {
	objabi *objabi.PackageSession
	sys    *sys.PackageSession

	AbiSymKindToSymKind [12]SymKind

	ReadOnly []SymKind

	RelROMap       map[SymKind]SymKind
	_SymKind_index [49]uint16
}

func NewPackageSession(objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.AbiSymKindToSymKind = [...]SymKind{
		Sxxx,
		STEXT,
		SRODATA,
		SNOPTRDATA,
		SDATA,
		SBSS,
		SNOPTRBSS,
		STLSBSS,
		SDWARFINFO,
		SDWARFRANGE,
		SDWARFLOC,
		SDWARFMISC,
	}
	psess.ReadOnly = []SymKind{
		STYPE,
		SSTRING,
		SGOSTRING,
		SGOFUNC,
		SGCBITS,
		SRODATA,
		SFUNCTAB,
	}
	psess.RelROMap = map[SymKind]SymKind{
		STYPE:     STYPERELRO,
		SSTRING:   SSTRINGRELRO,
		SGOSTRING: SGOSTRINGRELRO,
		SGOFUNC:   SGOFUNCRELRO,
		SGCBITS:   SGCBITSRELRO,
		SRODATA:   SRODATARELRO,
		SFUNCTAB:  SFUNCTABRELRO,
	}
	psess._SymKind_index = [...]uint16{0, 4, 9, 19, 24, 31, 40, 47, 54, 61, 69, 79, 88, 98, 110, 124, 136, 148, 160, 173, 182, 191, 198, 206, 214, 220, 229, 237, 244, 254, 262, 267, 271, 280, 287, 292, 304, 316, 333, 350, 359, 365, 375, 383, 393, 403, 414, 423, 433}
	return psess
}
