package gc

import "github.com/dave/golib/src/cmd/compile/internal/types"

// builtinpkg is a fake package that declares the universe block.

// distinguished *byte

// isBuiltinFuncName reports whether name matches a builtin function
// name.
func (psess *PackageSession) isBuiltinFuncName(name string) bool {
	for _, fn := range psess.builtinFuncs {
		if fn.name == name {
			return true
		}
	}
	return false
}

// initUniverse initializes the universe block.
func (psess *PackageSession) initUniverse() {
	psess.
		lexinit()
	psess.
		typeinit()
	psess.
		lexinit1()
}

// lexinit initializes known symbols and the basic types.
func (psess *PackageSession) lexinit() {
	for _, s := range psess.basicTypes {
		etype := s.etype
		if int(etype) >= len(psess.types.Types) {
			psess.
				Fatalf("lexinit: %s bad etype", s.name)
		}
		s2 := psess.builtinpkg.Lookup(psess.types, s.name)
		t := psess.types.Types[etype]
		if t == nil {
			t = types.New(etype)
			t.Sym = s2
			if etype != TANY && etype != TSTRING {
				psess.
					dowidth(t)
			}
			psess.types.
				Types[etype] = t
		}
		s2.Def = asTypesNode(psess.typenod(t))
		asNode(s2.Def).Name = new(Name)
	}

	for _, s := range psess.builtinFuncs {
		s2 := psess.builtinpkg.Lookup(psess.types, s.name)
		s2.Def = asTypesNode(psess.newname(s2))
		asNode(s2.Def).SetSubOp(psess, s.op)
	}

	for _, s := range psess.unsafeFuncs {
		s2 := psess.unsafepkg.Lookup(psess.types, s.name)
		s2.Def = asTypesNode(psess.newname(s2))
		asNode(s2.Def).SetSubOp(psess, s.op)
	}
	psess.types.
		Idealstring = types.New(TSTRING)
	psess.types.
		Idealbool = types.New(TBOOL)
	psess.types.
		Types[TANY] = types.New(TANY)

	s := psess.builtinpkg.Lookup(psess.types, "true")
	s.Def = asTypesNode(psess.nodbool(true))
	asNode(s.Def).Sym = psess.lookup("true")
	asNode(s.Def).Name = new(Name)
	asNode(s.Def).Type = psess.types.Idealbool

	s = psess.builtinpkg.Lookup(psess.types, "false")
	s.Def = asTypesNode(psess.nodbool(false))
	asNode(s.Def).Sym = psess.lookup("false")
	asNode(s.Def).Name = new(Name)
	asNode(s.Def).Type = psess.types.Idealbool

	s = psess.lookup("_")
	s.Block = -100
	s.Def = asTypesNode(psess.newname(s))
	psess.types.
		Types[TBLANK] = types.New(TBLANK)
	asNode(s.Def).Type = psess.types.Types[TBLANK]
	psess.
		nblank = asNode(s.Def)

	s = psess.builtinpkg.Lookup(psess.types, "_")
	s.Block = -100
	s.Def = asTypesNode(psess.newname(s))
	psess.types.
		Types[TBLANK] = types.New(TBLANK)
	asNode(s.Def).Type = psess.types.Types[TBLANK]
	psess.types.
		Types[TNIL] = types.New(TNIL)
	s = psess.builtinpkg.Lookup(psess.types, "nil")
	var v Val
	v.U = new(NilVal)
	s.Def = asTypesNode(psess.nodlit(v))
	asNode(s.Def).Sym = s
	asNode(s.Def).Name = new(Name)

	s = psess.builtinpkg.Lookup(psess.types, "iota")
	s.Def = asTypesNode(psess.nod(OIOTA, nil, nil))
	asNode(s.Def).Sym = s
	asNode(s.Def).Name = new(Name)
}

func (psess *PackageSession) typeinit() {
	if psess.Widthptr == 0 {
		psess.
			Fatalf("typeinit before betypeinit")
	}

	for et := types.EType(0); et < NTYPE; et++ {
		psess.
			simtype[et] = et
	}
	psess.types.
		Types[TPTR32] = types.New(TPTR32)
	psess.
		dowidth(psess.types.Types[TPTR32])
	psess.types.
		Types[TPTR64] = types.New(TPTR64)
	psess.
		dowidth(psess.types.Types[TPTR64])

	t := types.New(TUNSAFEPTR)
	psess.types.
		Types[TUNSAFEPTR] = t
	t.Sym = psess.unsafepkg.Lookup(psess.types, "Pointer")
	t.Sym.Def = asTypesNode(psess.typenod(t))
	asNode(t.Sym.Def).Name = new(Name)
	psess.
		dowidth(psess.types.Types[TUNSAFEPTR])
	psess.types.
		Tptr = TPTR32
	if psess.Widthptr == 8 {
		psess.types.
			Tptr = TPTR64
	}

	for et := TINT8; et <= TUINT64; et++ {
		psess.
			isInt[et] = true
	}
	psess.
		isInt[TINT] = true
	psess.
		isInt[TUINT] = true
	psess.
		isInt[TUINTPTR] = true
	psess.
		isFloat[TFLOAT32] = true
	psess.
		isFloat[TFLOAT64] = true
	psess.
		isComplex[TCOMPLEX64] = true
	psess.
		isComplex[TCOMPLEX128] = true
	psess.
		isforw[TFORW] = true

	for et := types.EType(0); et < NTYPE; et++ {
		if psess.isInt[et] || et == TIDEAL {
			psess.
				okforeq[et] = true
			psess.
				okforcmp[et] = true
			psess.
				okforarith[et] = true
			psess.
				okforadd[et] = true
			psess.
				okforand[et] = true
			psess.
				okforconst[et] = true
			psess.
				issimple[et] = true
			psess.
				minintval[et] = new(Mpint)
			psess.
				maxintval[et] = new(Mpint)
		}

		if psess.isFloat[et] {
			psess.
				okforeq[et] = true
			psess.
				okforcmp[et] = true
			psess.
				okforadd[et] = true
			psess.
				okforarith[et] = true
			psess.
				okforconst[et] = true
			psess.
				issimple[et] = true
			psess.
				minfltval[et] = newMpflt()
			psess.
				maxfltval[et] = newMpflt()
		}

		if psess.isComplex[et] {
			psess.
				okforeq[et] = true
			psess.
				okforadd[et] = true
			psess.
				okforarith[et] = true
			psess.
				okforconst[et] = true
			psess.
				issimple[et] = true
		}
	}
	psess.
		issimple[TBOOL] = true
	psess.
		okforadd[TSTRING] = true
	psess.
		okforbool[TBOOL] = true
	psess.
		okforcap[TARRAY] = true
	psess.
		okforcap[TCHAN] = true
	psess.
		okforcap[TSLICE] = true
	psess.
		okforconst[TBOOL] = true
	psess.
		okforconst[TSTRING] = true
	psess.
		okforlen[TARRAY] = true
	psess.
		okforlen[TCHAN] = true
	psess.
		okforlen[TMAP] = true
	psess.
		okforlen[TSLICE] = true
	psess.
		okforlen[TSTRING] = true
	psess.
		okforeq[TPTR32] = true
	psess.
		okforeq[TPTR64] = true
	psess.
		okforeq[TUNSAFEPTR] = true
	psess.
		okforeq[TINTER] = true
	psess.
		okforeq[TCHAN] = true
	psess.
		okforeq[TSTRING] = true
	psess.
		okforeq[TBOOL] = true
	psess.
		okforeq[TMAP] = true
	psess.
		okforeq[TFUNC] = true
	psess.
		okforeq[TSLICE] = true
	psess.
		okforeq[TARRAY] = true
	psess.
		okforeq[TSTRUCT] = true
	psess.
		okforcmp[TSTRING] = true

	var i int
	for i = 0; i < len(psess.okfor); i++ {
		psess.
			okfor[i] = psess.okfornone[:]
	}
	psess.
		okfor[OADD] = psess.okforadd[:]
	psess.
		okfor[OAND] = psess.okforand[:]
	psess.
		okfor[OANDAND] = psess.okforbool[:]
	psess.
		okfor[OANDNOT] = psess.okforand[:]
	psess.
		okfor[ODIV] = psess.okforarith[:]
	psess.
		okfor[OEQ] = psess.okforeq[:]
	psess.
		okfor[OGE] = psess.okforcmp[:]
	psess.
		okfor[OGT] = psess.okforcmp[:]
	psess.
		okfor[OLE] = psess.okforcmp[:]
	psess.
		okfor[OLT] = psess.okforcmp[:]
	psess.
		okfor[OMOD] = psess.okforand[:]
	psess.
		okfor[OMUL] = psess.okforarith[:]
	psess.
		okfor[ONE] = psess.okforeq[:]
	psess.
		okfor[OOR] = psess.okforand[:]
	psess.
		okfor[OOROR] = psess.okforbool[:]
	psess.
		okfor[OSUB] = psess.okforarith[:]
	psess.
		okfor[OXOR] = psess.okforand[:]
	psess.
		okfor[OLSH] = psess.okforand[:]
	psess.
		okfor[ORSH] = psess.okforand[:]
	psess.
		okfor[OCOM] = psess.okforand[:]
	psess.
		okfor[OMINUS] = psess.okforarith[:]
	psess.
		okfor[ONOT] = psess.okforbool[:]
	psess.
		okfor[OPLUS] = psess.okforarith[:]
	psess.
		okfor[OCAP] = psess.okforcap[:]
	psess.
		okfor[OLEN] = psess.okforlen[:]
	psess.
		iscmp[OLT] = true
	psess.
		iscmp[OGT] = true
	psess.
		iscmp[OGE] = true
	psess.
		iscmp[OLE] = true
	psess.
		iscmp[OEQ] = true
	psess.
		iscmp[ONE] = true
	psess.
		maxintval[TINT8].SetString(psess, "0x7f")
	psess.
		minintval[TINT8].SetString(psess, "-0x80")
	psess.
		maxintval[TINT16].SetString(psess, "0x7fff")
	psess.
		minintval[TINT16].SetString(psess, "-0x8000")
	psess.
		maxintval[TINT32].SetString(psess, "0x7fffffff")
	psess.
		minintval[TINT32].SetString(psess, "-0x80000000")
	psess.
		maxintval[TINT64].SetString(psess, "0x7fffffffffffffff")
	psess.
		minintval[TINT64].SetString(psess, "-0x8000000000000000")
	psess.
		maxintval[TUINT8].SetString(psess, "0xff")
	psess.
		maxintval[TUINT16].SetString(psess, "0xffff")
	psess.
		maxintval[TUINT32].SetString(psess, "0xffffffff")
	psess.
		maxintval[TUINT64].SetString(psess, "0xffffffffffffffff")
	psess.
		maxfltval[TFLOAT32].SetString(psess, "33554431p103")
	psess.
		minfltval[TFLOAT32].SetString(psess, "-33554431p103")
	psess.
		maxfltval[TFLOAT64].SetString(psess, "18014398509481983p970")
	psess.
		minfltval[TFLOAT64].SetString(psess, "-18014398509481983p970")
	psess.
		maxfltval[TCOMPLEX64] = psess.maxfltval[TFLOAT32]
	psess.
		minfltval[TCOMPLEX64] = psess.minfltval[TFLOAT32]
	psess.
		maxfltval[TCOMPLEX128] = psess.maxfltval[TFLOAT64]
	psess.
		minfltval[TCOMPLEX128] = psess.minfltval[TFLOAT64]
	psess.types.
		Types[TFUNC] = psess.functype(nil, nil, nil)
	psess.types.
		Types[TIDEAL] = types.New(TIDEAL)
	psess.types.
		Types[TINTER] = types.New(TINTER)
	psess.
		simtype[TMAP] = psess.types.Tptr
	psess.
		simtype[TCHAN] = psess.types.Tptr
	psess.
		simtype[TFUNC] = psess.types.Tptr
	psess.
		simtype[TUNSAFEPTR] = psess.types.Tptr
	psess.
		array_array = int(psess.Rnd(0, int64(psess.Widthptr)))
	psess.
		array_nel = int(psess.Rnd(int64(psess.array_array)+int64(psess.Widthptr), int64(psess.Widthptr)))
	psess.
		array_cap = int(psess.Rnd(int64(psess.array_nel)+int64(psess.Widthptr), int64(psess.Widthptr)))
	psess.
		sizeof_Array = int(psess.Rnd(int64(psess.array_cap)+int64(psess.Widthptr), int64(psess.Widthptr)))
	psess.
		sizeof_String = int(psess.Rnd(int64(psess.array_nel)+int64(psess.Widthptr), int64(psess.Widthptr)))
	psess.
		dowidth(psess.types.Types[TSTRING])
	psess.
		dowidth(psess.types.Idealstring)
	psess.
		itable = psess.types.NewPtr(psess.types.Types[TUINT8])
}

func (psess *PackageSession) makeErrorInterface() *types.Type {
	field := types.NewField()
	field.Type = psess.types.Types[TSTRING]
	f := psess.functypefield(psess.fakeRecvField(), nil, []*types.Field{field})

	field = types.NewField()
	field.Sym = psess.lookup("Error")
	field.Type = f

	t := types.New(TINTER)
	t.SetInterface(psess.types, []*types.Field{field})
	return t
}

func (psess *PackageSession) lexinit1() {

	s := psess.builtinpkg.Lookup(psess.types, "error")
	psess.types.
		Errortype = psess.makeErrorInterface()
	psess.types.
		Errortype.Sym = s
	psess.types.
		Errortype.Orig = psess.makeErrorInterface()
	s.Def = asTypesNode(psess.typenod(psess.types.Errortype))

	s = psess.builtinpkg.Lookup(psess.types, "byte")
	psess.types.
		Bytetype = types.New(TUINT8)
	psess.types.
		Bytetype.Sym = s
	s.Def = asTypesNode(psess.typenod(psess.types.Bytetype))
	asNode(s.Def).Name = new(Name)

	s = psess.builtinpkg.Lookup(psess.types, "rune")
	psess.types.
		Runetype = types.New(TINT32)
	psess.types.
		Runetype.Sym = s
	s.Def = asTypesNode(psess.typenod(psess.types.Runetype))
	asNode(s.Def).Name = new(Name)

	for _, s := range psess.typedefs {
		s1 := psess.builtinpkg.Lookup(psess.types, s.name)

		sameas := s.sameas32
		if psess.Widthptr == 8 {
			sameas = s.sameas64
		}
		psess.
			simtype[s.etype] = sameas
		psess.
			minfltval[s.etype] = psess.minfltval[sameas]
		psess.
			maxfltval[s.etype] = psess.maxfltval[sameas]
		psess.
			minintval[s.etype] = psess.minintval[sameas]
		psess.
			maxintval[s.etype] = psess.maxintval[sameas]

		t := types.New(s.etype)
		t.Sym = s1
		psess.types.
			Types[s.etype] = t
		s1.Def = asTypesNode(psess.typenod(t))
		asNode(s1.Def).Name = new(Name)
		s1.Origpkg = psess.builtinpkg
		psess.
			dowidth(t)
	}
}

// finishUniverse makes the universe block visible within the current package.
func (psess *PackageSession) finishUniverse() {

	for _, s := range psess.builtinpkg.Syms {
		if s.Def == nil {
			continue
		}
		s1 := psess.lookup(s.Name)
		if s1.Def != nil {
			continue
		}

		s1.Def = s.Def
		s1.Block = s.Block
	}
	psess.
		nodfp = psess.newname(psess.lookup(".fp"))
	psess.
		nodfp.Type = psess.types.Types[TINT32]
	psess.
		nodfp.SetClass(PPARAM)
	psess.
		nodfp.Name.SetUsed(true)
}
