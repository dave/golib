package types

type PackageState struct {
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	src       *src.PackageState
	Block     int32
	Bytetype  *Type
	Ctxt      *obj.Link
	Dowidth   func( *Type)
	Errortype *Type
	FErr      int
	Fatalf    func( string,  []interface {
	})
	FmtLeft            int
	FmtUnsigned        int
	FormatSym          func( *Sym,  fmt.State,  rune,  int)
	FormatType         func( *Type,  fmt.State,  rune,  int)
	Idealbool          *Type
	Idealcomplex       *Type
	Idealfloat         *Type
	Idealint           *Type
	Idealrune          *Type
	Idealstring        *Type
	InitSyms           []*Sym
	NewPtrCacheEnabled bool
	ParamsResults      [2]func( *Type) ( *Type)
	RecvsParams        [2]func( *Type) ( *Type)
	RecvsParamsResults [3]func( *Type) ( *Type)
	Runetype           *Type
	Sconv              func( *Sym,  int,  int) ( string)
	Tconv              func( *Type,  int,  int,  int) ( string)
	Tptr               EType
	TypeFlags          *Type
	TypeInt128         *Type
	TypeInvalid        *Type
	TypeLinkSym        func( *Type) ( *obj.LSym)
	TypeMem            *Type
	TypeVoid           *Type
	Types              [38]*Type
	Widthptr           int
	_EType_index       [40]uint8
	blockgen           int32
	dclstack           []dsym
	internedStrings    map[string]string
	internedStringsmu  sync.Mutex
	nopkg              *Pkg
	pkgMap             map[string]*Pkg
	recvType           *Type
	unsignedEType      [12]EType
}

func NewPackageState(obj_pstate *obj.PackageState, objabi_pstate *objabi.PackageState, src_pstate *src.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.obj = obj_pstate
	pstate.objabi = objabi_pstate
	pstate.src = src_pstate
	pstate.pkgMap = make(map[string]*Pkg)
	pstate.nopkg = &Pkg{
		Syms: make(map[string]*Sym),
	}
	pstate.internedStrings = map[string]string{}
	pstate.blockgen = 1
	pstate.NewPtrCacheEnabled = true
	pstate.unsignedEType = [...]EType{
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
	pstate.Idealint = New(TIDEAL)
	pstate.Idealrune = New(TIDEAL)
	pstate.Idealfloat = New(TIDEAL)
	pstate.Idealcomplex = New(TIDEAL)
	pstate.TypeInvalid = newSSA("invalid")
	pstate.TypeMem = newSSA("mem")
	pstate.TypeFlags = newSSA("flags")
	pstate.TypeVoid = newSSA("void")
	pstate.TypeInt128 = newSSA("int128")
	pstate.RecvsParamsResults = [3]func(*Type) *Type{
		(*Type).Recvs, (*Type).Params, (*Type).Results,
	}
	pstate.RecvsParams = [2]func(*Type) *Type{
		(*Type).Recvs, (*Type).Params,
	}
	pstate.ParamsResults = [2]func(*Type) *Type{
		(*Type).Params, (*Type).Results,
	}
	pstate._EType_index = [...]uint8{0, 3, 7, 12, 17, 23, 28, 34, 39, 45, 48, 52, 59, 68, 78, 85, 92, 96, 101, 106, 110, 115, 120, 126, 130, 133, 138, 142, 145, 151, 160, 165, 168, 173, 181, 189, 197, 200, 205, 210}
	return pstate
}
