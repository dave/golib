package sys

import "encoding/binary"

type PackageSession struct {
	Arch386 *Arch

	ArchAMD64 *Arch

	ArchAMD64P32 *Arch

	ArchARM *Arch

	ArchARM64 *Arch

	ArchMIPS *Arch

	ArchMIPS64 *Arch

	ArchMIPS64LE *Arch
	ArchMIPSLE   *Arch

	ArchPPC64 *Arch

	ArchPPC64LE *Arch

	ArchS390X *Arch

	ArchWasm *Arch

	Archs [13]*Arch
}

func NewPackageSession() *PackageSession {
	psess := &PackageSession{}
	psess.Arch386 = &Arch{
		Name:      "386",
		Family:    I386,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     1,
	}
	psess.ArchAMD64 = &Arch{
		Name:      "amd64",
		Family:    AMD64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     1,
	}
	psess.ArchAMD64P32 = &Arch{
		Name:      "amd64p32",
		Family:    AMD64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   8,
		MinLC:     1,
	}
	psess.ArchARM = &Arch{
		Name:      "arm",
		Family:    ARM,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	psess.ArchARM64 = &Arch{
		Name:      "arm64",
		Family:    ARM64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	psess.ArchMIPS = &Arch{
		Name:      "mips",
		Family:    MIPS,
		ByteOrder: binary.BigEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	psess.ArchMIPSLE = &Arch{
		Name:      "mipsle",
		Family:    MIPS,
		ByteOrder: binary.LittleEndian,
		PtrSize:   4,
		RegSize:   4,
		MinLC:     4,
	}
	psess.ArchMIPS64 = &Arch{
		Name:      "mips64",
		Family:    MIPS64,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	psess.ArchMIPS64LE = &Arch{
		Name:      "mips64le",
		Family:    MIPS64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	psess.ArchPPC64 = &Arch{
		Name:      "ppc64",
		Family:    PPC64,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	psess.ArchPPC64LE = &Arch{
		Name:      "ppc64le",
		Family:    PPC64,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     4,
	}
	psess.ArchS390X = &Arch{
		Name:      "s390x",
		Family:    S390X,
		ByteOrder: binary.BigEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     2,
	}
	psess.ArchWasm = &Arch{
		Name:      "wasm",
		Family:    Wasm,
		ByteOrder: binary.LittleEndian,
		PtrSize:   8,
		RegSize:   8,
		MinLC:     1,
	}
	psess.Archs = [...]*Arch{psess.
		Arch386, psess.
		ArchAMD64, psess.
		ArchAMD64P32, psess.
		ArchARM, psess.
		ArchARM64, psess.
		ArchMIPS, psess.
		ArchMIPSLE, psess.
		ArchMIPS64, psess.
		ArchMIPS64LE, psess.
		ArchPPC64, psess.
		ArchPPC64LE, psess.
		ArchS390X, psess.
		ArchWasm,
	}
	return psess
}
