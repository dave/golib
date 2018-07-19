package sym

type Library struct {
	Objref        string
	Srcref        string
	File          string
	Pkg           string
	Shlib         string
	Hash          string
	ImportStrings []string
	Imports       []*Library
	Textp         []*Symbol // text symbols defined in this library
	DupTextSyms   []*Symbol // dupok text symbols defined in this library
	Main          bool
	Safe          bool
}

func (l Library) String() string {
	return l.Pkg
}
