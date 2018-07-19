package sym

type Symbols struct {
	symbolBatch []Symbol

	// Symbol lookup based on name and indexed by version.
	hash []map[string]*Symbol

	Allsym []*Symbol
}

func NewSymbols() *Symbols {
	return &Symbols{
		hash: []map[string]*Symbol{

			make(map[string]*Symbol, 100000),
		},
		Allsym: make([]*Symbol, 0, 100000),
	}
}

func (syms *Symbols) Newsym(name string, v int) *Symbol {
	batch := syms.symbolBatch
	if len(batch) == 0 {
		batch = make([]Symbol, 1000)
	}
	s := &batch[0]
	syms.symbolBatch = batch[1:]

	s.Dynid = -1
	s.Plt = -1
	s.Got = -1
	s.Name = name
	s.Version = int16(v)
	syms.Allsym = append(syms.Allsym, s)

	return s
}

// Look up the symbol with the given name and version, creating the
// symbol if it is not found.
func (syms *Symbols) Lookup(name string, v int) *Symbol {
	m := syms.hash[v]
	s := m[name]
	if s != nil {
		return s
	}
	s = syms.Newsym(name, v)
	s.Extname = s.Name
	m[name] = s
	return s
}

// Look up the symbol with the given name and version, returning nil
// if it is not found.
func (syms *Symbols) ROLookup(name string, v int) *Symbol {
	return syms.hash[v][name]
}

// Allocate a new version (i.e. symbol namespace).
func (syms *Symbols) IncVersion() int {
	syms.hash = append(syms.hash, make(map[string]*Symbol))
	return len(syms.hash) - 1
}

// Rename renames a symbol.
func (syms *Symbols) Rename(old, new string, v int) {
	s := syms.hash[v][old]
	s.Name = new
	if s.Extname == old {
		s.Extname = new
	}
	delete(syms.hash[v], old)

	dup := syms.hash[v][new]
	if dup == nil {
		syms.hash[v][new] = s
	} else {
		if s.Type == 0 {
			*s = *dup
		} else if dup.Type == 0 {
			*dup = *s
			syms.hash[v][new] = s
		}
	}
}
