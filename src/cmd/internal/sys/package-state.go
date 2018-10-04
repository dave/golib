package sys

type PackageState struct {
	Arch386      *Arch
	ArchAMD64    *Arch
	ArchAMD64P32 *Arch
	ArchARM      *Arch
	ArchARM64    *Arch
	ArchMIPS     *Arch
	ArchMIPS64   *Arch
	ArchMIPS64LE *Arch
	ArchMIPSLE   *Arch
	ArchPPC64    *Arch
	ArchPPC64LE  *Arch
	ArchS390X    *Arch
	ArchWasm     *Arch
	Archs        [13]*Arch
}

func NewPackageState() *PackageState {
	pstate := &PackageState{}
	pstate.Arch386 = &Arch{
		Name:      "386",
		Family:    I386,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     1,
	}
	pstate.ArchAMD64 = &Arch{
		Name:      "amd64",
		Family:    AMD64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     1,
	}
	pstate.ArchAMD64P32 = &Arch{
		Name:      "amd64p32",
		Family:    AMD64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   8,
		MinLC:     1,
	}
	pstate.ArchARM = &Arch{
		Name:      "arm",
		Family:    ARM,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	pstate.ArchARM64 = &Arch{
		Name:      "arm64",
		Family:    ARM64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	pstate.ArchMIPS = &Arch{
		Name:      "mips",
		Family:    MIPS,
		ByteOrder: binary.BigEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	pstate.ArchMIPSLE = &Arch{
		Name:      "mipsle",
		Family:    MIPS,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	pstate.ArchMIPS64 = &Arch{
		Name:      "mips64",
		Family:    MIPS64,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	pstate.ArchMIPS64LE = &Arch{
		Name:      "mips64le",
		Family:    MIPS64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	pstate.ArchPPC64 = &Arch{
		Name:      "ppc64",
		Family:    PPC64,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	pstate.ArchPPC64LE = &Arch{
		Name:      "ppc64le",
		Family:    PPC64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	pstate.ArchS390X = &Arch{
		Name:      "s390x",
		Family:    S390X,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     2,
	}
	pstate.ArchWasm = &Arch{
		Name:      "wasm",
		Family:    Wasm,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     1,
	}
	pstate.Archs = [...]*Arch{
		pstate.Arch386,
		pstate.ArchAMD64,
		pstate.ArchAMD64P32,
		pstate.ArchARM,
		pstate.ArchARM64,
		pstate.ArchMIPS,
		pstate.ArchMIPSLE,
		pstate.ArchMIPS64,
		pstate.ArchMIPS64LE,
		pstate.ArchPPC64,
		pstate.ArchPPC64LE,
		pstate.ArchS390X,
		pstate.ArchWasm,
	}
	return pstate
}
