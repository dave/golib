package obj

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"log"
	"math"
)

func (psess *PackageSession) Linknew(arch *LinkArch) *Link {
	ctxt := new(Link)
	ctxt.hash = make(map[string]*LSym)
	ctxt.statichash = make(map[string]*LSym)
	ctxt.Arch = arch
	ctxt.Pathname = objabi.WorkingDir()

	if err := ctxt.Headtype.Set(psess.objabi.GOOS); err != nil {
		log.Fatalf("unknown goos %s", psess.objabi.GOOS)
	}

	ctxt.Flag_optimize = true
	ctxt.Framepointer_enabled = psess.objabi.Framepointer_enabled(psess.objabi.GOOS, arch.Name)
	return ctxt
}

// LookupDerived looks up or creates the symbol with name name derived from symbol s.
// The resulting symbol will be static iff s is.
func (ctxt *Link) LookupDerived(s *LSym, name string) *LSym {
	if s.Static() {
		return ctxt.LookupStatic(name)
	}
	return ctxt.Lookup(name)
}

// LookupStatic looks up the static symbol with name name.
// If it does not exist, it creates it.
func (ctxt *Link) LookupStatic(name string) *LSym {
	s := ctxt.statichash[name]
	if s == nil {
		s = &LSym{Name: name, Attribute: AttrStatic}
		ctxt.statichash[name] = s
	}
	return s
}

// Lookup looks up the symbol with name name.
// If it does not exist, it creates it.
func (ctxt *Link) Lookup(name string) *LSym {
	return ctxt.LookupInit(name, nil)
}

// LookupInit looks up the symbol with name name.
// If it does not exist, it creates it and
// passes it to init for one-time initialization.
func (ctxt *Link) LookupInit(name string, init func(s *LSym)) *LSym {
	ctxt.hashmu.Lock()
	s := ctxt.hash[name]
	if s == nil {
		s = &LSym{Name: name}
		ctxt.hash[name] = s
		if init != nil {
			init(s)
		}
	}
	ctxt.hashmu.Unlock()
	return s
}

func (ctxt *Link) Float32Sym(f float32) *LSym {
	i := math.Float32bits(f)
	name := fmt.Sprintf("$f32.%08x", i)
	return ctxt.LookupInit(name, func(s *LSym) {
		s.Size = 4
		s.Set(AttrLocal, true)
	})
}

func (ctxt *Link) Float64Sym(f float64) *LSym {
	i := math.Float64bits(f)
	name := fmt.Sprintf("$f64.%016x", i)
	return ctxt.LookupInit(name, func(s *LSym) {
		s.Size = 8
		s.Set(AttrLocal, true)
	})
}

func (ctxt *Link) Int64Sym(i int64) *LSym {
	name := fmt.Sprintf("$i64.%016x", uint64(i))
	return ctxt.LookupInit(name, func(s *LSym) {
		s.Size = 8
		s.Set(AttrLocal, true)
	})
}
