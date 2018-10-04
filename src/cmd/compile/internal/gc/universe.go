// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(gri) This file should probably become part of package types.

package gc

import "github.com/dave/golib/src/cmd/compile/internal/types"

// isBuiltinFuncName reports whether name matches a builtin function
// name.
func (pstate *PackageState) isBuiltinFuncName(name string) bool {
	for _, fn := range pstate.builtinFuncs {
		if fn.name == name {
			return true
		}
	}
	return false
}

// initUniverse initializes the universe block.
func (pstate *PackageState) initUniverse() {
	pstate.lexinit()
	pstate.typeinit()
	pstate.lexinit1()
}

// lexinit initializes known symbols and the basic types.
func (pstate *PackageState) lexinit() {
	for _, s := range pstate.basicTypes {
		etype := s.etype
		if int(etype) >= len(pstate.types.Types) {
			pstate.Fatalf("lexinit: %s bad etype", s.name)
		}
		s2 := pstate.builtinpkg.Lookup(pstate.types, s.name)
		t := pstate.types.Types[etype]
		if t == nil {
			t = types.New(etype)
			t.Sym = s2
			if etype != TANY && etype != TSTRING {
				pstate.dowidth(t)
			}
			pstate.types.Types[etype] = t
		}
		s2.Def = asTypesNode(pstate.typenod(t))
		asNode(s2.Def).Name = new(Name)
	}

	for _, s := range pstate.builtinFuncs {
		s2 := pstate.builtinpkg.Lookup(pstate.types, s.name)
		s2.Def = asTypesNode(pstate.newname(s2))
		asNode(s2.Def).SetSubOp(pstate, s.op)
	}

	for _, s := range pstate.unsafeFuncs {
		s2 := pstate.unsafepkg.Lookup(pstate.types, s.name)
		s2.Def = asTypesNode(pstate.newname(s2))
		asNode(s2.Def).SetSubOp(pstate, s.op)
	}

	pstate.types.Idealstring = types.New(TSTRING)
	pstate.types.Idealbool = types.New(TBOOL)
	pstate.types.Types[TANY] = types.New(TANY)

	s := pstate.builtinpkg.Lookup(pstate.types, "true")
	s.Def = asTypesNode(pstate.nodbool(true))
	asNode(s.Def).Sym = pstate.lookup("true")
	asNode(s.Def).Name = new(Name)
	asNode(s.Def).Type = pstate.types.Idealbool

	s = pstate.builtinpkg.Lookup(pstate.types, "false")
	s.Def = asTypesNode(pstate.nodbool(false))
	asNode(s.Def).Sym = pstate.lookup("false")
	asNode(s.Def).Name = new(Name)
	asNode(s.Def).Type = pstate.types.Idealbool

	s = pstate.lookup("_")
	s.Block = -100
	s.Def = asTypesNode(pstate.newname(s))
	pstate.types.Types[TBLANK] = types.New(TBLANK)
	asNode(s.Def).Type = pstate.types.Types[TBLANK]
	pstate.nblank = asNode(s.Def)

	s = pstate.builtinpkg.Lookup(pstate.types, "_")
	s.Block = -100
	s.Def = asTypesNode(pstate.newname(s))
	pstate.types.Types[TBLANK] = types.New(TBLANK)
	asNode(s.Def).Type = pstate.types.Types[TBLANK]

	pstate.types.Types[TNIL] = types.New(TNIL)
	s = pstate.builtinpkg.Lookup(pstate.types, "nil")
	var v Val
	v.U = new(NilVal)
	s.Def = asTypesNode(pstate.nodlit(v))
	asNode(s.Def).Sym = s
	asNode(s.Def).Name = new(Name)

	s = pstate.builtinpkg.Lookup(pstate.types, "iota")
	s.Def = asTypesNode(pstate.nod(OIOTA, nil, nil))
	asNode(s.Def).Sym = s
	asNode(s.Def).Name = new(Name)
}

func (pstate *PackageState) typeinit() {
	if pstate.Widthptr == 0 {
		pstate.Fatalf("typeinit before betypeinit")
	}

	for et := types.EType(0); et < NTYPE; et++ {
		pstate.simtype[et] = et
	}

	pstate.types.Types[TPTR32] = types.New(TPTR32)
	pstate.dowidth(pstate.types.Types[TPTR32])

	pstate.types.Types[TPTR64] = types.New(TPTR64)
	pstate.dowidth(pstate.types.Types[TPTR64])

	t := types.New(TUNSAFEPTR)
	pstate.types.Types[TUNSAFEPTR] = t
	t.Sym = pstate.unsafepkg.Lookup(pstate.types, "Pointer")
	t.Sym.Def = asTypesNode(pstate.typenod(t))
	asNode(t.Sym.Def).Name = new(Name)
	pstate.dowidth(pstate.types.Types[TUNSAFEPTR])

	pstate.types.Tptr = TPTR32
	if pstate.Widthptr == 8 {
		pstate.types.Tptr = TPTR64
	}

	for et := TINT8; et <= TUINT64; et++ {
		pstate.isInt[et] = true
	}
	pstate.isInt[TINT] = true
	pstate.isInt[TUINT] = true
	pstate.isInt[TUINTPTR] = true

	pstate.isFloat[TFLOAT32] = true
	pstate.isFloat[TFLOAT64] = true

	pstate.isComplex[TCOMPLEX64] = true
	pstate.isComplex[TCOMPLEX128] = true

	pstate.isforw[TFORW] = true

	// initialize okfor
	for et := types.EType(0); et < NTYPE; et++ {
		if pstate.isInt[et] || et == TIDEAL {
			pstate.okforeq[et] = true
			pstate.okforcmp[et] = true
			pstate.okforarith[et] = true
			pstate.okforadd[et] = true
			pstate.okforand[et] = true
			pstate.okforconst[et] = true
			pstate.issimple[et] = true
			pstate.minintval[et] = new(Mpint)
			pstate.maxintval[et] = new(Mpint)
		}

		if pstate.isFloat[et] {
			pstate.okforeq[et] = true
			pstate.okforcmp[et] = true
			pstate.okforadd[et] = true
			pstate.okforarith[et] = true
			pstate.okforconst[et] = true
			pstate.issimple[et] = true
			pstate.minfltval[et] = newMpflt()
			pstate.maxfltval[et] = newMpflt()
		}

		if pstate.isComplex[et] {
			pstate.okforeq[et] = true
			pstate.okforadd[et] = true
			pstate.okforarith[et] = true
			pstate.okforconst[et] = true
			pstate.issimple[et] = true
		}
	}

	pstate.issimple[TBOOL] = true

	pstate.okforadd[TSTRING] = true

	pstate.okforbool[TBOOL] = true

	pstate.okforcap[TARRAY] = true
	pstate.okforcap[TCHAN] = true
	pstate.okforcap[TSLICE] = true

	pstate.okforconst[TBOOL] = true
	pstate.okforconst[TSTRING] = true

	pstate.okforlen[TARRAY] = true
	pstate.okforlen[TCHAN] = true
	pstate.okforlen[TMAP] = true
	pstate.okforlen[TSLICE] = true
	pstate.okforlen[TSTRING] = true

	pstate.okforeq[TPTR32] = true
	pstate.okforeq[TPTR64] = true
	pstate.okforeq[TUNSAFEPTR] = true
	pstate.okforeq[TINTER] = true
	pstate.okforeq[TCHAN] = true
	pstate.okforeq[TSTRING] = true
	pstate.okforeq[TBOOL] = true
	pstate.okforeq[TMAP] = true    // nil only; refined in typecheck
	pstate.okforeq[TFUNC] = true   // nil only; refined in typecheck
	pstate.okforeq[TSLICE] = true  // nil only; refined in typecheck
	pstate.okforeq[TARRAY] = true  // only if element type is comparable; refined in typecheck
	pstate.okforeq[TSTRUCT] = true // only if all struct fields are comparable; refined in typecheck

	pstate.okforcmp[TSTRING] = true

	var i int
	for i = 0; i < len(pstate.okfor); i++ {
		pstate.okfor[i] = pstate.okfornone[:]
	}

	// binary
	pstate.okfor[OADD] = pstate.okforadd[:]
	pstate.okfor[OAND] = pstate.okforand[:]
	pstate.okfor[OANDAND] = pstate.okforbool[:]
	pstate.okfor[OANDNOT] = pstate.okforand[:]
	pstate.okfor[ODIV] = pstate.okforarith[:]
	pstate.okfor[OEQ] = pstate.okforeq[:]
	pstate.okfor[OGE] = pstate.okforcmp[:]
	pstate.okfor[OGT] = pstate.okforcmp[:]
	pstate.okfor[OLE] = pstate.okforcmp[:]
	pstate.okfor[OLT] = pstate.okforcmp[:]
	pstate.okfor[OMOD] = pstate.okforand[:]
	pstate.okfor[OMUL] = pstate.okforarith[:]
	pstate.okfor[ONE] = pstate.okforeq[:]
	pstate.okfor[OOR] = pstate.okforand[:]
	pstate.okfor[OOROR] = pstate.okforbool[:]
	pstate.okfor[OSUB] = pstate.okforarith[:]
	pstate.okfor[OXOR] = pstate.okforand[:]
	pstate.okfor[OLSH] = pstate.okforand[:]
	pstate.okfor[ORSH] = pstate.okforand[:]

	// unary
	pstate.okfor[OCOM] = pstate.okforand[:]
	pstate.okfor[OMINUS] = pstate.okforarith[:]
	pstate.okfor[ONOT] = pstate.okforbool[:]
	pstate.okfor[OPLUS] = pstate.okforarith[:]

	// special
	pstate.okfor[OCAP] = pstate.okforcap[:]
	pstate.okfor[OLEN] = pstate.okforlen[:]

	// comparison
	pstate.iscmp[OLT] = true
	pstate.iscmp[OGT] = true
	pstate.iscmp[OGE] = true
	pstate.iscmp[OLE] = true
	pstate.iscmp[OEQ] = true
	pstate.iscmp[ONE] = true

	pstate.maxintval[TINT8].SetString(pstate, "0x7f")
	pstate.minintval[TINT8].SetString(pstate, "-0x80")
	pstate.maxintval[TINT16].SetString(pstate, "0x7fff")
	pstate.minintval[TINT16].SetString(pstate, "-0x8000")
	pstate.maxintval[TINT32].SetString(pstate, "0x7fffffff")
	pstate.minintval[TINT32].SetString(pstate, "-0x80000000")
	pstate.maxintval[TINT64].SetString(pstate, "0x7fffffffffffffff")
	pstate.minintval[TINT64].SetString(pstate, "-0x8000000000000000")

	pstate.maxintval[TUINT8].SetString(pstate, "0xff")
	pstate.maxintval[TUINT16].SetString(pstate, "0xffff")
	pstate.maxintval[TUINT32].SetString(pstate, "0xffffffff")
	pstate.maxintval[TUINT64].SetString(pstate, "0xffffffffffffffff")

	// f is valid float if min < f < max.  (min and max are not themselves valid.)
	pstate.maxfltval[TFLOAT32].SetString(pstate, "33554431p103") // 2^24-1 p (127-23) + 1/2 ulp
	pstate.minfltval[TFLOAT32].SetString(pstate, "-33554431p103")
	pstate.maxfltval[TFLOAT64].SetString(pstate, "18014398509481983p970") // 2^53-1 p (1023-52) + 1/2 ulp
	pstate.minfltval[TFLOAT64].SetString(pstate, "-18014398509481983p970")

	pstate.maxfltval[TCOMPLEX64] = pstate.maxfltval[TFLOAT32]
	pstate.minfltval[TCOMPLEX64] = pstate.minfltval[TFLOAT32]
	pstate.maxfltval[TCOMPLEX128] = pstate.maxfltval[TFLOAT64]
	pstate.minfltval[TCOMPLEX128] = pstate.minfltval[TFLOAT64]

	// for walk to use in error messages
	pstate.types.Types[TFUNC] = pstate.functype(nil, nil, nil)

	// types used in front end
	// types.Types[TNIL] got set early in lexinit
	pstate.types.Types[TIDEAL] = types.New(TIDEAL)

	pstate.types.Types[TINTER] = types.New(TINTER)

	// simple aliases
	pstate.simtype[TMAP] = pstate.types.Tptr
	pstate.simtype[TCHAN] = pstate.types.Tptr
	pstate.simtype[TFUNC] = pstate.types.Tptr
	pstate.simtype[TUNSAFEPTR] = pstate.types.Tptr

	pstate.array_array = int(pstate.Rnd(0, int64(pstate.Widthptr)))
	pstate.array_nel = int(pstate.Rnd(int64(pstate.array_array)+int64(pstate.Widthptr), int64(pstate.Widthptr)))
	pstate.array_cap = int(pstate.Rnd(int64(pstate.array_nel)+int64(pstate.Widthptr), int64(pstate.Widthptr)))
	pstate.sizeof_Array = int(pstate.Rnd(int64(pstate.array_cap)+int64(pstate.Widthptr), int64(pstate.Widthptr)))

	// string is same as slice wo the cap
	pstate.sizeof_String = int(pstate.Rnd(int64(pstate.array_nel)+int64(pstate.Widthptr), int64(pstate.Widthptr)))

	pstate.dowidth(pstate.types.Types[TSTRING])
	pstate.dowidth(pstate.types.Idealstring)

	pstate.itable = pstate.types.NewPtr(pstate.types.Types[TUINT8])
}

func (pstate *PackageState) makeErrorInterface() *types.Type {
	field := types.NewField()
	field.Type = pstate.types.Types[TSTRING]
	f := pstate.functypefield(pstate.fakeRecvField(), nil, []*types.Field{field})

	field = types.NewField()
	field.Sym = pstate.lookup("Error")
	field.Type = f

	t := types.New(TINTER)
	t.SetInterface(pstate.types, []*types.Field{field})
	return t
}

func (pstate *PackageState) lexinit1() {
	// error type
	s := pstate.builtinpkg.Lookup(pstate.types, "error")
	pstate.types.Errortype = pstate.makeErrorInterface()
	pstate.types.Errortype.Sym = s
	pstate.types.Errortype.Orig = pstate.makeErrorInterface()
	s.Def = asTypesNode(pstate.typenod(pstate.types.Errortype))

	// We create separate byte and rune types for better error messages
	// rather than just creating type alias *types.Sym's for the uint8 and
	// int32 types. Hence, (bytetype|runtype).Sym.isAlias() is false.
	// TODO(gri) Should we get rid of this special case (at the cost
	// of less informative error messages involving bytes and runes)?
	// (Alternatively, we could introduce an OTALIAS node representing
	// type aliases, albeit at the cost of having to deal with it everywhere).

	// byte alias
	s = pstate.builtinpkg.Lookup(pstate.types, "byte")
	pstate.types.Bytetype = types.New(TUINT8)
	pstate.types.Bytetype.Sym = s
	s.Def = asTypesNode(pstate.typenod(pstate.types.Bytetype))
	asNode(s.Def).Name = new(Name)

	// rune alias
	s = pstate.builtinpkg.Lookup(pstate.types, "rune")
	pstate.types.Runetype = types.New(TINT32)
	pstate.types.Runetype.Sym = s
	s.Def = asTypesNode(pstate.typenod(pstate.types.Runetype))
	asNode(s.Def).Name = new(Name)

	// backend-dependent builtin types (e.g. int).
	for _, s := range pstate.typedefs {
		s1 := pstate.builtinpkg.Lookup(pstate.types, s.name)

		sameas := s.sameas32
		if pstate.Widthptr == 8 {
			sameas = s.sameas64
		}

		pstate.simtype[s.etype] = sameas
		pstate.minfltval[s.etype] = pstate.minfltval[sameas]
		pstate.maxfltval[s.etype] = pstate.maxfltval[sameas]
		pstate.minintval[s.etype] = pstate.minintval[sameas]
		pstate.maxintval[s.etype] = pstate.maxintval[sameas]

		t := types.New(s.etype)
		t.Sym = s1
		pstate.types.Types[s.etype] = t
		s1.Def = asTypesNode(pstate.typenod(t))
		asNode(s1.Def).Name = new(Name)
		s1.Origpkg = pstate.builtinpkg

		pstate.dowidth(t)
	}
}

// finishUniverse makes the universe block visible within the current package.
func (pstate *PackageState) finishUniverse() {
	// Operationally, this is similar to a dot import of builtinpkg, except
	// that we silently skip symbols that are already declared in the
	// package block rather than emitting a redeclared symbol error.

	for _, s := range pstate.builtinpkg.Syms {
		if s.Def == nil {
			continue
		}
		s1 := pstate.lookup(s.Name)
		if s1.Def != nil {
			continue
		}

		s1.Def = s.Def
		s1.Block = s.Block
	}

	pstate.nodfp = pstate.newname(pstate.lookup(".fp"))
	pstate.nodfp.Type = pstate.types.Types[TINT32]
	pstate.nodfp.SetClass(PPARAM)
	pstate.nodfp.Name.SetUsed(true)
}
