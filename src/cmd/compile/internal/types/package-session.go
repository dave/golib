package types

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
)

type PackageSession struct {
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	src    *src.PackageSession

	Block int32

	Bytetype *Type
	Ctxt     *obj.Link
	Dowidth  func(*Type)

	Errortype *Type
	FErr      int
	Fatalf    func(string, []interface {
	})

	FmtLeft     int
	FmtUnsigned int
	FormatSym   func(*Sym, fmt.State, rune, int)
	FormatType  func(*Type, fmt.State, rune, int)

	Idealbool *Type

	Idealcomplex *Type
	Idealfloat   *Type
	Idealint     *Type
	Idealrune    *Type
	Idealstring  *Type
	InitSyms     []*Sym

	NewPtrCacheEnabled bool

	ParamsResults      [2]func(*Type) *Type
	RecvsParams        [2]func(*Type) *Type
	RecvsParamsResults [3]func(*Type) *Type
	Runetype           *Type
	Sconv              func(*Sym, int, int) string
	Tconv              func(*Type, int, int, int) string
	Tptr               EType

	TypeFlags *Type

	TypeInt128  *Type
	TypeInvalid *Type
	TypeLinkSym func(*Type) *obj.LSym

	TypeMem *Type

	TypeVoid     *Type
	Types        [38]*Type
	Widthptr     int
	_EType_index [40]uint8
	blockgen     int32

	dclstack []dsym

	internedStrings   map[string]string
	internedStringsmu sync.Mutex
	nopkg             *Pkg
	pkgMap            map[string]*Pkg

	recvType      *Type
	unsignedEType [12]EType
}

func NewPackageSession(obj_psess *obj.PackageSession, objabi_psess *objabi.PackageSession, src_psess *src.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.obj = obj_psess
	psess.objabi = objabi_psess
	psess.src = src_psess
	psess._EType_index = [...]uint8{0, 3, 7, 12, 17, 23, 28, 34, 39, 45, 48, 52, 59, 68, 78, 85, 92, 96, 101, 106, 110, 115, 120, 126, 130, 133, 138, 142, 145, 151, 160, 165, 168, 173, 181, 189, 197, 200, 205, 210}
	psess.pkgMap = make(map[string]*Pkg)
	psess.nopkg = &Pkg{
		Syms: make(map[string]*Sym),
	}
	psess.internedStrings = map[string]string{}
	psess.blockgen = 1
	psess.Idealint = New(TIDEAL)
	psess.Idealrune = New(TIDEAL)
	psess.Idealfloat = New(TIDEAL)
	psess.Idealcomplex = New(TIDEAL)
	psess.NewPtrCacheEnabled = true
	psess.RecvsParamsResults = [3]func(*Type) *Type{
		(*Type).Recvs, (*Type).Params, (*Type).Results,
	}
	psess.RecvsParams = [2]func(*Type) *Type{
		(*Type).Recvs, (*Type).Params,
	}
	psess.ParamsResults = [2]func(*Type) *Type{
		(*Type).Params, (*Type).Results,
	}
	psess.unsignedEType = [...]EType{
		TINT8:    TUINT8,
		TUINT8:   TUINT8,
		TINT16:   TUINT16,
		TUINT16:  TUINT16,
		TINT32:   TUINT32,
		TUINT32:  TUINT32,
		TINT64:   TUINT64,
		TUINT64:  TUINT64,
		TINT:     TUINT,
		TUINT:    TUINT,
		TUINTPTR: TUINTPTR,
	}
	psess.TypeInvalid = newSSA("invalid")
	psess.TypeMem = newSSA("mem")
	psess.TypeFlags = newSSA("flags")
	psess.TypeVoid = newSSA("void")
	psess.TypeInt128 = newSSA("int128")
	return psess
}
