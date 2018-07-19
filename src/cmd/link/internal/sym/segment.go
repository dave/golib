package sym

type Segment struct {
	Rwx      uint8  // permission as usual unix bits (5 = r-x etc)
	Vaddr    uint64 // virtual address
	Length   uint64 // length in memory
	Fileoff  uint64 // file offset
	Filelen  uint64 // length on disk
	Sections []*Section
}

type Section struct {
	Rwx     uint8
	Extnum  int16
	Align   int32
	Name    string
	Vaddr   uint64
	Length  uint64
	Seg     *Segment
	Elfsect interface{} // an *ld.ElfShdr
	Reloff  uint64
	Rellen  uint64
}
