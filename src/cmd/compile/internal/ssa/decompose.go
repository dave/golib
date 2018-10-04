// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

// decompose converts phi ops on compound builtin types into phi
// ops on simple types, then invokes rewrite rules to decompose
// other ops on those types.
func (pstate *PackageState) decomposeBuiltIn(f *Func) {
	// Decompose phis
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			pstate.decomposeBuiltInPhi(v)
		}
	}

	// Decompose other values
	pstate.applyRewrite(f, rewriteBlockdec, pstate.rewriteValuedec)
	if f.Config.RegSize == 4 {
		pstate.applyRewrite(f, rewriteBlockdec64, pstate.rewriteValuedec64)
	}

	// Split up named values into their components.
	var newNames []LocalSlot
	for _, name := range f.Names {
		t := name.Type
		switch {
		case t.IsInteger() && t.Size(pstate.types) > f.Config.RegSize:
			hiName, loName := f.fe.SplitInt64(name)
			newNames = append(newNames, hiName, loName)
			for _, v := range f.NamedValues[name] {
				if v.Op != OpInt64Make {
					continue
				}
				f.NamedValues[hiName] = append(f.NamedValues[hiName], v.Args[0])
				f.NamedValues[loName] = append(f.NamedValues[loName], v.Args[1])
			}
			delete(f.NamedValues, name)
		case t.IsComplex():
			rName, iName := f.fe.SplitComplex(name)
			newNames = append(newNames, rName, iName)
			for _, v := range f.NamedValues[name] {
				if v.Op != OpComplexMake {
					continue
				}
				f.NamedValues[rName] = append(f.NamedValues[rName], v.Args[0])
				f.NamedValues[iName] = append(f.NamedValues[iName], v.Args[1])

			}
			delete(f.NamedValues, name)
		case t.IsString():
			ptrName, lenName := f.fe.SplitString(name)
			newNames = append(newNames, ptrName, lenName)
			for _, v := range f.NamedValues[name] {
				if v.Op != OpStringMake {
					continue
				}
				f.NamedValues[ptrName] = append(f.NamedValues[ptrName], v.Args[0])
				f.NamedValues[lenName] = append(f.NamedValues[lenName], v.Args[1])
			}
			delete(f.NamedValues, name)
		case t.IsSlice():
			ptrName, lenName, capName := f.fe.SplitSlice(name)
			newNames = append(newNames, ptrName, lenName, capName)
			for _, v := range f.NamedValues[name] {
				if v.Op != OpSliceMake {
					continue
				}
				f.NamedValues[ptrName] = append(f.NamedValues[ptrName], v.Args[0])
				f.NamedValues[lenName] = append(f.NamedValues[lenName], v.Args[1])
				f.NamedValues[capName] = append(f.NamedValues[capName], v.Args[2])
			}
			delete(f.NamedValues, name)
		case t.IsInterface():
			typeName, dataName := f.fe.SplitInterface(name)
			newNames = append(newNames, typeName, dataName)
			for _, v := range f.NamedValues[name] {
				if v.Op != OpIMake {
					continue
				}
				f.NamedValues[typeName] = append(f.NamedValues[typeName], v.Args[0])
				f.NamedValues[dataName] = append(f.NamedValues[dataName], v.Args[1])
			}
			delete(f.NamedValues, name)
		case t.IsFloat():
			// floats are never decomposed, even ones bigger than RegSize
			newNames = append(newNames, name)
		case t.Size(pstate.types) > f.Config.RegSize:
			f.Fatalf("undecomposed named type %s %v", name, t)
		default:
			newNames = append(newNames, name)
		}
	}
	f.Names = newNames
}

func (pstate *PackageState) decomposeBuiltInPhi(v *Value) {
	switch {
	case v.Type.IsInteger() && v.Type.Size(pstate.types) > v.Block.Func.Config.RegSize:
		decomposeInt64Phi(v)
	case v.Type.IsComplex():
		pstate.decomposeComplexPhi(v)
	case v.Type.IsString():
		decomposeStringPhi(v)
	case v.Type.IsSlice():
		decomposeSlicePhi(v)
	case v.Type.IsInterface():
		decomposeInterfacePhi(v)
	case v.Type.IsFloat():
	// floats are never decomposed, even ones bigger than RegSize
	case v.Type.Size(pstate.types) > v.Block.Func.Config.RegSize:
		v.Fatalf("undecomposed type %s", v.Type)
	}
}

func decomposeStringPhi(v *Value) {
	types := &v.Block.Func.Config.Types
	ptrType := types.BytePtr
	lenType := types.Int

	ptr := v.Block.NewValue0(v.Pos, OpPhi, ptrType)
	len := v.Block.NewValue0(v.Pos, OpPhi, lenType)
	for _, a := range v.Args {
		ptr.AddArg(a.Block.NewValue1(v.Pos, OpStringPtr, ptrType, a))
		len.AddArg(a.Block.NewValue1(v.Pos, OpStringLen, lenType, a))
	}
	v.reset(OpStringMake)
	v.AddArg(ptr)
	v.AddArg(len)
}

func decomposeSlicePhi(v *Value) {
	types := &v.Block.Func.Config.Types
	ptrType := types.BytePtr
	lenType := types.Int

	ptr := v.Block.NewValue0(v.Pos, OpPhi, ptrType)
	len := v.Block.NewValue0(v.Pos, OpPhi, lenType)
	cap := v.Block.NewValue0(v.Pos, OpPhi, lenType)
	for _, a := range v.Args {
		ptr.AddArg(a.Block.NewValue1(v.Pos, OpSlicePtr, ptrType, a))
		len.AddArg(a.Block.NewValue1(v.Pos, OpSliceLen, lenType, a))
		cap.AddArg(a.Block.NewValue1(v.Pos, OpSliceCap, lenType, a))
	}
	v.reset(OpSliceMake)
	v.AddArg(ptr)
	v.AddArg(len)
	v.AddArg(cap)
}

func decomposeInt64Phi(v *Value) {
	cfgtypes := &v.Block.Func.Config.Types
	var partType *types.Type
	if v.Type.IsSigned() {
		partType = cfgtypes.Int32
	} else {
		partType = cfgtypes.UInt32
	}

	hi := v.Block.NewValue0(v.Pos, OpPhi, partType)
	lo := v.Block.NewValue0(v.Pos, OpPhi, cfgtypes.UInt32)
	for _, a := range v.Args {
		hi.AddArg(a.Block.NewValue1(v.Pos, OpInt64Hi, partType, a))
		lo.AddArg(a.Block.NewValue1(v.Pos, OpInt64Lo, cfgtypes.UInt32, a))
	}
	v.reset(OpInt64Make)
	v.AddArg(hi)
	v.AddArg(lo)
}

func (pstate *PackageState) decomposeComplexPhi(v *Value) {
	cfgtypes := &v.Block.Func.Config.Types
	var partType *types.Type
	switch z := v.Type.Size(pstate.types); z {
	case 8:
		partType = cfgtypes.Float32
	case 16:
		partType = cfgtypes.Float64
	default:
		v.Fatalf("decomposeComplexPhi: bad complex size %d", z)
	}

	real := v.Block.NewValue0(v.Pos, OpPhi, partType)
	imag := v.Block.NewValue0(v.Pos, OpPhi, partType)
	for _, a := range v.Args {
		real.AddArg(a.Block.NewValue1(v.Pos, OpComplexReal, partType, a))
		imag.AddArg(a.Block.NewValue1(v.Pos, OpComplexImag, partType, a))
	}
	v.reset(OpComplexMake)
	v.AddArg(real)
	v.AddArg(imag)
}

func decomposeInterfacePhi(v *Value) {
	uintptrType := v.Block.Func.Config.Types.Uintptr
	ptrType := v.Block.Func.Config.Types.BytePtr

	itab := v.Block.NewValue0(v.Pos, OpPhi, uintptrType)
	data := v.Block.NewValue0(v.Pos, OpPhi, ptrType)
	for _, a := range v.Args {
		itab.AddArg(a.Block.NewValue1(v.Pos, OpITab, uintptrType, a))
		data.AddArg(a.Block.NewValue1(v.Pos, OpIData, ptrType, a))
	}
	v.reset(OpIMake)
	v.AddArg(itab)
	v.AddArg(data)
}

func (pstate *PackageState) decomposeUser(f *Func) {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			pstate.decomposeUserPhi(v)
		}
	}
	// Split up named values into their components.
	i := 0
	var newNames []LocalSlot
	for _, name := range f.Names {
		t := name.Type
		switch {
		case t.IsStruct():
			newNames = pstate.decomposeUserStructInto(f, name, newNames)
		case t.IsArray():
			newNames = pstate.decomposeUserArrayInto(f, name, newNames)
		default:
			f.Names[i] = name
			i++
		}
	}
	f.Names = f.Names[:i]
	f.Names = append(f.Names, newNames...)
}

// decomposeUserArrayInto creates names for the element(s) of arrays referenced
// by name where possible, and appends those new names to slots, which is then
// returned.
func (pstate *PackageState) decomposeUserArrayInto(f *Func, name LocalSlot, slots []LocalSlot) []LocalSlot {
	t := name.Type
	if t.NumElem(pstate.types) == 0 {
		// TODO(khr): Not sure what to do here.  Probably nothing.
		// Names for empty arrays aren't important.
		return slots
	}
	if t.NumElem(pstate.types) != 1 {
		// shouldn't get here due to CanSSA
		f.Fatalf("array not of size 1")
	}
	elemName := f.fe.SplitArray(name)
	for _, v := range f.NamedValues[name] {
		if v.Op != OpArrayMake1 {
			continue
		}
		f.NamedValues[elemName] = append(f.NamedValues[elemName], v.Args[0])
	}
	// delete the name for the array as a whole
	delete(f.NamedValues, name)

	if t.Elem(pstate.types).IsArray() {
		return pstate.decomposeUserArrayInto(f, elemName, slots)
	} else if t.Elem(pstate.types).IsStruct() {
		return pstate.decomposeUserStructInto(f, elemName, slots)
	}

	return append(slots, elemName)
}

// decomposeUserStructInto creates names for the fields(s) of structs referenced
// by name where possible, and appends those new names to slots, which is then
// returned.
func (pstate *PackageState) decomposeUserStructInto(f *Func, name LocalSlot, slots []LocalSlot) []LocalSlot {
	fnames := []LocalSlot{} // slots for struct in name
	t := name.Type
	n := t.NumFields(pstate.types)

	for i := 0; i < n; i++ {
		fs := f.fe.SplitStruct(name, i)
		fnames = append(fnames, fs)
		// arrays and structs will be decomposed further, so
		// there's no need to record a name
		if !fs.Type.IsArray() && !fs.Type.IsStruct() {
			slots = append(slots, fs)
		}
	}

	makeOp := StructMakeOp(n)
	// create named values for each struct field
	for _, v := range f.NamedValues[name] {
		if v.Op != makeOp {
			continue
		}
		for i := 0; i < len(fnames); i++ {
			f.NamedValues[fnames[i]] = append(f.NamedValues[fnames[i]], v.Args[i])
		}
	}
	// remove the name of the struct as a whole
	delete(f.NamedValues, name)

	// now that this f.NamedValues contains values for the struct
	// fields, recurse into nested structs
	for i := 0; i < n; i++ {
		if name.Type.FieldType(pstate.types, i).IsStruct() {
			slots = pstate.decomposeUserStructInto(f, fnames[i], slots)
			delete(f.NamedValues, fnames[i])
		} else if name.Type.FieldType(pstate.types, i).IsArray() {
			slots = pstate.decomposeUserArrayInto(f, fnames[i], slots)
			delete(f.NamedValues, fnames[i])
		}
	}
	return slots
}
func (pstate *PackageState) decomposeUserPhi(v *Value) {
	switch {
	case v.Type.IsStruct():
		pstate.decomposeStructPhi(v)
	case v.Type.IsArray():
		pstate.decomposeArrayPhi(v)
	}
}

// decomposeStructPhi replaces phi-of-struct with structmake(phi-for-each-field),
// and then recursively decomposes the phis for each field.
func (pstate *PackageState) decomposeStructPhi(v *Value) {
	t := v.Type
	n := t.NumFields(pstate.types)
	var fields [MaxStruct]*Value
	for i := 0; i < n; i++ {
		fields[i] = v.Block.NewValue0(v.Pos, OpPhi, t.FieldType(pstate.types, i))
	}
	for _, a := range v.Args {
		for i := 0; i < n; i++ {
			fields[i].AddArg(a.Block.NewValue1I(v.Pos, OpStructSelect, t.FieldType(pstate.types, i), int64(i), a))
		}
	}
	v.reset(StructMakeOp(n))
	v.AddArgs(fields[:n]...)

	// Recursively decompose phis for each field.
	for _, f := range fields[:n] {
		pstate.decomposeUserPhi(f)
	}
}

// decomposeArrayPhi replaces phi-of-array with arraymake(phi-of-array-element),
// and then recursively decomposes the element phi.
func (pstate *PackageState) decomposeArrayPhi(v *Value) {
	t := v.Type
	if t.NumElem(pstate.types) == 0 {
		v.reset(OpArrayMake0)
		return
	}
	if t.NumElem(pstate.types) != 1 {
		v.Fatalf("SSAable array must have no more than 1 element")
	}
	elem := v.Block.NewValue0(v.Pos, OpPhi, t.Elem(pstate.types))
	for _, a := range v.Args {
		elem.AddArg(a.Block.NewValue1I(v.Pos, OpArraySelect, t.Elem(pstate.types), 0, a))
	}
	v.reset(OpArrayMake1)
	v.AddArg(elem)

	// Recursively decompose elem phi.
	pstate.decomposeUserPhi(elem)
}

// MaxStruct is the maximum number of fields a struct
// can have and still be SSAable.
const MaxStruct = 4

// StructMakeOp returns the opcode to construct a struct with the
// given number of fields.
func StructMakeOp(nf int) Op {
	switch nf {
	case 0:
		return OpStructMake0
	case 1:
		return OpStructMake1
	case 2:
		return OpStructMake2
	case 3:
		return OpStructMake3
	case 4:
		return OpStructMake4
	}
	panic("too many fields in an SSAable struct")
}
