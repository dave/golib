package ld

import (
	"bytes"
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
)

// tflag is documented in reflect/type.go.
//
// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	reflect/type.go
//	runtime/type.go
const (
	tflagUncommon  = 1 << 0
	tflagExtraStar = 1 << 1
)

func decodeReloc(s *sym.Symbol, off int32) *sym.Reloc {
	for i := range s.R {
		if s.R[i].Off == off {
			return &s.R[i]
		}
	}
	return nil
}

func decodeRelocSym(s *sym.Symbol, off int32) *sym.Symbol {
	r := decodeReloc(s, off)
	if r == nil {
		return nil
	}
	return r.Sym
}

func (psess *PackageSession) decodeInuxi(arch *sys.Arch, p []byte, sz int) uint64 {
	switch sz {
	case 2:
		return uint64(arch.ByteOrder.Uint16(p))
	case 4:
		return uint64(arch.ByteOrder.Uint32(p))
	case 8:
		return arch.ByteOrder.Uint64(p)
	default:
		psess.
			Exitf("dwarf: decode inuxi %d", sz)
		panic("unreachable")
	}
}

func commonsize(arch *sys.Arch) int      { return 4*arch.PtrSize + 8 + 8 }
func structfieldSize(arch *sys.Arch) int { return 3 * arch.PtrSize }
func uncommonSize() int                  { return 4 + 2 + 2 + 4 + 4 }

// Type.commonType.kind
func decodetypeKind(arch *sys.Arch, s *sym.Symbol) uint8 {
	return s.P[2*arch.PtrSize+7] & objabi.KindMask
}

// Type.commonType.kind
func decodetypeUsegcprog(arch *sys.Arch, s *sym.Symbol) uint8 {
	return s.P[2*arch.PtrSize+7] & objabi.KindGCProg
}

// Type.commonType.size
func (psess *PackageSession) decodetypeSize(arch *sys.Arch, s *sym.Symbol) int64 {
	return int64(psess.decodeInuxi(arch, s.P, arch.PtrSize))
}

// Type.commonType.ptrdata
func (psess *PackageSession) decodetypePtrdata(arch *sys.Arch, s *sym.Symbol) int64 {
	return int64(psess.decodeInuxi(arch, s.P[arch.PtrSize:], arch.PtrSize))
}

// Type.commonType.tflag
func decodetypeHasUncommon(arch *sys.Arch, s *sym.Symbol) bool {
	return s.P[2*arch.PtrSize+4]&tflagUncommon != 0
}

// Find the elf.Section of a given shared library that contains a given address.
func findShlibSection(ctxt *Link, path string, addr uint64) *elf.Section {
	for _, shlib := range ctxt.Shlibs {
		if shlib.Path == path {
			for _, sect := range shlib.File.Sections {
				if sect.Addr <= addr && addr <= sect.Addr+sect.Size {
					return sect
				}
			}
		}
	}
	return nil
}

// Type.commonType.gc
func (psess *PackageSession) decodetypeGcprog(ctxt *Link, s *sym.Symbol) []byte {
	if s.Type == sym.SDYNIMPORT {
		addr := psess.decodetypeGcprogShlib(ctxt, s)
		sect := findShlibSection(ctxt, s.File, addr)
		if sect != nil {

			progsize := make([]byte, 4)
			sect.ReadAt(progsize, int64(addr-sect.Addr))
			progbytes := make([]byte, ctxt.Arch.ByteOrder.Uint32(progsize))
			sect.ReadAt(progbytes, int64(addr-sect.Addr+4))
			return append(progsize, progbytes...)
		}
		psess.
			Exitf("cannot find gcprog for %s", s.Name)
		return nil
	}
	return decodeRelocSym(s, 2*int32(ctxt.Arch.PtrSize)+8+1*int32(ctxt.Arch.PtrSize)).P
}

func (psess *PackageSession) decodetypeGcprogShlib(ctxt *Link, s *sym.Symbol) uint64 {
	if ctxt.Arch.Family == sys.ARM64 {
		for _, shlib := range ctxt.Shlibs {
			if shlib.Path == s.File {
				return shlib.gcdataAddresses[s]
			}
		}
		return 0
	}
	return psess.decodeInuxi(ctxt.Arch, s.P[2*int32(ctxt.Arch.PtrSize)+8+1*int32(ctxt.Arch.PtrSize):], ctxt.Arch.PtrSize)
}

func (psess *PackageSession) decodetypeGcmask(ctxt *Link, s *sym.Symbol) []byte {
	if s.Type == sym.SDYNIMPORT {
		addr := psess.decodetypeGcprogShlib(ctxt, s)
		ptrdata := psess.decodetypePtrdata(ctxt.Arch, s)
		sect := findShlibSection(ctxt, s.File, addr)
		if sect != nil {
			r := make([]byte, ptrdata/int64(ctxt.Arch.PtrSize))
			sect.ReadAt(r, int64(addr-sect.Addr))
			return r
		}
		psess.
			Exitf("cannot find gcmask for %s", s.Name)
		return nil
	}
	mask := decodeRelocSym(s, 2*int32(ctxt.Arch.PtrSize)+8+1*int32(ctxt.Arch.PtrSize))
	return mask.P
}

// Type.ArrayType.elem and Type.SliceType.Elem
func decodetypeArrayElem(arch *sys.Arch, s *sym.Symbol) *sym.Symbol {
	return decodeRelocSym(s, int32(commonsize(arch)))
}

func (psess *PackageSession) decodetypeArrayLen(arch *sys.Arch, s *sym.Symbol) int64 {
	return int64(psess.decodeInuxi(arch, s.P[commonsize(arch)+2*arch.PtrSize:], arch.PtrSize))
}

// Type.PtrType.elem
func decodetypePtrElem(arch *sys.Arch, s *sym.Symbol) *sym.Symbol {
	return decodeRelocSym(s, int32(commonsize(arch)))
}

// Type.MapType.key, elem
func decodetypeMapKey(arch *sys.Arch, s *sym.Symbol) *sym.Symbol {
	return decodeRelocSym(s, int32(commonsize(arch)))
}

func decodetypeMapValue(arch *sys.Arch, s *sym.Symbol) *sym.Symbol {
	return decodeRelocSym(s, int32(commonsize(arch))+int32(arch.PtrSize))
}

// Type.ChanType.elem
func decodetypeChanElem(arch *sys.Arch, s *sym.Symbol) *sym.Symbol {
	return decodeRelocSym(s, int32(commonsize(arch)))
}

// Type.FuncType.dotdotdot
func (psess *PackageSession) decodetypeFuncDotdotdot(arch *sys.Arch, s *sym.Symbol) bool {
	return uint16(psess.decodeInuxi(arch, s.P[commonsize(arch)+2:], 2))&(1<<15) != 0
}

// Type.FuncType.inCount
func (psess *PackageSession) decodetypeFuncInCount(arch *sys.Arch, s *sym.Symbol) int {
	return int(psess.decodeInuxi(arch, s.P[commonsize(arch):], 2))
}

func (psess *PackageSession) decodetypeFuncOutCount(arch *sys.Arch, s *sym.Symbol) int {
	return int(uint16(psess.decodeInuxi(arch, s.P[commonsize(arch)+2:], 2)) & (1<<15 - 1))
}

func decodetypeFuncInType(arch *sys.Arch, s *sym.Symbol, i int) *sym.Symbol {
	uadd := commonsize(arch) + 4
	if arch.PtrSize == 8 {
		uadd += 4
	}
	if decodetypeHasUncommon(arch, s) {
		uadd += uncommonSize()
	}
	return decodeRelocSym(s, int32(uadd+i*arch.PtrSize))
}

func (psess *PackageSession) decodetypeFuncOutType(arch *sys.Arch, s *sym.Symbol, i int) *sym.Symbol {
	return decodetypeFuncInType(arch, s, i+psess.decodetypeFuncInCount(arch, s))
}

// Type.StructType.fields.Slice::length
func (psess *PackageSession) decodetypeStructFieldCount(arch *sys.Arch, s *sym.Symbol) int {
	return int(psess.decodeInuxi(arch, s.P[commonsize(arch)+2*arch.PtrSize:], arch.PtrSize))
}

func decodetypeStructFieldArrayOff(arch *sys.Arch, s *sym.Symbol, i int) int {
	off := commonsize(arch) + 4*arch.PtrSize
	if decodetypeHasUncommon(arch, s) {
		off += uncommonSize()
	}
	off += i * structfieldSize(arch)
	return off
}

// decodetypeStr returns the contents of an rtype's str field (a nameOff).
func decodetypeStr(arch *sys.Arch, s *sym.Symbol) string {
	str := decodetypeName(s, 4*arch.PtrSize+8)
	if s.P[2*arch.PtrSize+4]&tflagExtraStar != 0 {
		return str[1:]
	}
	return str
}

// decodetypeName decodes the name from a reflect.name.
func decodetypeName(s *sym.Symbol, off int) string {
	r := decodeReloc(s, int32(off))
	if r == nil {
		return ""
	}

	data := r.Sym.P
	namelen := int(uint16(data[1])<<8 | uint16(data[2]))
	return string(data[3 : 3+namelen])
}

func decodetypeStructFieldName(arch *sys.Arch, s *sym.Symbol, i int) string {
	off := decodetypeStructFieldArrayOff(arch, s, i)
	return decodetypeName(s, off)
}

func decodetypeStructFieldType(arch *sys.Arch, s *sym.Symbol, i int) *sym.Symbol {
	off := decodetypeStructFieldArrayOff(arch, s, i)
	return decodeRelocSym(s, int32(off+arch.PtrSize))
}

func (psess *PackageSession) decodetypeStructFieldOffs(arch *sys.Arch, s *sym.Symbol, i int) int64 {
	return psess.decodetypeStructFieldOffsAnon(arch, s, i) >> 1
}

func (psess *PackageSession) decodetypeStructFieldOffsAnon(arch *sys.Arch, s *sym.Symbol, i int) int64 {
	off := decodetypeStructFieldArrayOff(arch, s, i)
	return int64(psess.decodeInuxi(arch, s.P[off+2*arch.PtrSize:], arch.PtrSize))
}

// InterfaceType.methods.length
func (psess *PackageSession) decodetypeIfaceMethodCount(arch *sys.Arch, s *sym.Symbol) int64 {
	return int64(psess.decodeInuxi(arch, s.P[commonsize(arch)+2*arch.PtrSize:], arch.PtrSize))
}

// methodsig is a fully qualified typed method signature, like
// "Visit(type.go/ast.Node) (type.go/ast.Visitor)".
type methodsig string

// Matches runtime/typekind.go and reflect.Kind.
const (
	kindArray     = 17
	kindChan      = 18
	kindFunc      = 19
	kindInterface = 20
	kindMap       = 21
	kindPtr       = 22
	kindSlice     = 23
	kindStruct    = 25
	kindMask      = (1 << 5) - 1
)

// decodeMethodSig decodes an array of method signature information.
// Each element of the array is size bytes. The first 4 bytes is a
// nameOff for the method name, and the next 4 bytes is a typeOff for
// the function type.
//
// Conveniently this is the layout of both runtime.method and runtime.imethod.
func (psess *PackageSession) decodeMethodSig(arch *sys.Arch, s *sym.Symbol, off, size, count int) []methodsig {
	var buf bytes.Buffer
	var methods []methodsig
	for i := 0; i < count; i++ {
		buf.WriteString(decodetypeName(s, off))
		mtypSym := decodeRelocSym(s, int32(off+4))

		buf.WriteRune('(')
		inCount := psess.decodetypeFuncInCount(arch, mtypSym)
		for i := 0; i < inCount; i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(decodetypeFuncInType(arch, mtypSym, i).Name)
		}
		buf.WriteString(") (")
		outCount := psess.decodetypeFuncOutCount(arch, mtypSym)
		for i := 0; i < outCount; i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(psess.decodetypeFuncOutType(arch, mtypSym, i).Name)
		}
		buf.WriteRune(')')

		off += size
		methods = append(methods, methodsig(buf.String()))
		buf.Reset()
	}
	return methods
}

func (psess *PackageSession) decodeIfaceMethods(arch *sys.Arch, s *sym.Symbol) []methodsig {
	if decodetypeKind(arch, s)&kindMask != kindInterface {
		panic(fmt.Sprintf("symbol %q is not an interface", s.Name))
	}
	r := decodeReloc(s, int32(commonsize(arch)+arch.PtrSize))
	if r == nil {
		return nil
	}
	if r.Sym != s {
		panic(fmt.Sprintf("imethod slice pointer in %q leads to a different symbol", s.Name))
	}
	off := int(r.Add)
	numMethods := int(psess.decodetypeIfaceMethodCount(arch, s))
	sizeofIMethod := 4 + 4
	return psess.decodeMethodSig(arch, s, off, sizeofIMethod, numMethods)
}

func (psess *PackageSession) decodetypeMethods(arch *sys.Arch, s *sym.Symbol) []methodsig {
	if !decodetypeHasUncommon(arch, s) {
		panic(fmt.Sprintf("no methods on %q", s.Name))
	}
	off := commonsize(arch)
	switch decodetypeKind(arch, s) & kindMask {
	case kindStruct:
		off += 4 * arch.PtrSize
	case kindPtr:
		off += arch.PtrSize
	case kindFunc:
		off += arch.PtrSize
	case kindSlice:
		off += arch.PtrSize
	case kindArray:
		off += 3 * arch.PtrSize
	case kindChan:
		off += 2 * arch.PtrSize
	case kindMap:
		off += 3*arch.PtrSize + 8
	case kindInterface:
		off += 3 * arch.PtrSize
	default:

	}

	mcount := int(psess.decodeInuxi(arch, s.P[off+4:], 2))
	moff := int(psess.decodeInuxi(arch, s.P[off+4+2+2:], 4))
	off += moff
	const sizeofMethod = 4 * 4 // sizeof reflect.method in program
	return psess.decodeMethodSig(arch, s, off, sizeofMethod, mcount)
}
