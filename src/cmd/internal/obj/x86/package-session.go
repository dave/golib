package x86

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
)

type PackageSession struct {
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	src    *src.PackageSession
	sys    *sys.PackageSession

	AMD64DWARFRegisters map[int16]int16
	Anames              []string

	Link386   obj.LinkArch
	Linkamd64 obj.LinkArch

	Linkamd64p32 obj.LinkArch
	Register     []string

	X86DWARFRegisters map[int16]int16
	_yandnl           []ytab

	_ybextrl []ytab

	_yblsil []ytab

	_ykaddb []ytab

	_ykmovb []ytab

	_yknotb []ytab

	_ykshiftlb []ytab

	_yrorxl []ytab

	_yv4fmaddps []ytab

	_yv4fmaddss []ytab

	_yvaddpd []ytab

	_yvaddsd []ytab

	_yvaddsubpd []ytab

	_yvaesdec []ytab

	_yvaesimc []ytab

	_yvaeskeygenassist []ytab

	_yvalignd []ytab

	_yvandnpd []ytab

	_yvblendmpd []ytab

	_yvblendpd []ytab

	_yvblendvpd []ytab

	_yvbroadcastf128 []ytab

	_yvbroadcastf32x2 []ytab

	_yvbroadcastf32x4 []ytab

	_yvbroadcastf32x8 []ytab

	_yvbroadcasti32x2 []ytab

	_yvbroadcastsd []ytab

	_yvbroadcastss []ytab

	_yvcmppd []ytab

	_yvcmpsd []ytab

	_yvcomisd []ytab

	_yvcompresspd []ytab

	_yvcvtdq2pd []ytab

	_yvcvtdq2ps []ytab

	_yvcvtpd2dq []ytab

	_yvcvtpd2dqx []ytab

	_yvcvtpd2dqy []ytab

	_yvcvtpd2qq []ytab

	_yvcvtpd2udqx []ytab

	_yvcvtpd2udqy []ytab

	_yvcvtph2ps []ytab

	_yvcvtps2ph []ytab

	_yvcvtps2qq []ytab

	_yvcvtsd2si []ytab

	_yvcvtsd2usil []ytab

	_yvcvtsi2sdl []ytab

	_yvcvtudq2pd []ytab

	_yvcvtusi2sdl []ytab

	_yvdppd []ytab

	_yvexp2pd []ytab

	_yvexpandpd []ytab

	_yvextractf128 []ytab

	_yvextractf32x4 []ytab

	_yvextractf32x8 []ytab

	_yvextractps []ytab

	_yvfixupimmpd []ytab

	_yvfixupimmsd []ytab

	_yvfpclasspdx []ytab

	_yvfpclasspdy []ytab

	_yvfpclasspdz []ytab

	_yvgatherdpd []ytab

	_yvgatherdps []ytab

	_yvgatherpf0dpd []ytab

	_yvgatherpf0dps []ytab

	_yvgatherqps []ytab

	_yvgetexpsd []ytab

	_yvgetmantpd []ytab

	_yvgf2p8affineinvqb []ytab

	_yvinsertf128 []ytab

	_yvinsertf32x4 []ytab

	_yvinsertf32x8 []ytab

	_yvinsertps []ytab

	_yvlddqu []ytab

	_yvldmxcsr []ytab

	_yvmaskmovdqu []ytab

	_yvmaskmovpd []ytab

	_yvmovapd []ytab

	_yvmovd []ytab

	_yvmovddup []ytab

	_yvmovdqa []ytab

	_yvmovdqa32 []ytab

	_yvmovhlps []ytab

	_yvmovhpd []ytab

	_yvmovmskpd []ytab

	_yvmovntdq []ytab

	_yvmovntdqa []ytab

	_yvmovq []ytab

	_yvmovsd []ytab

	_yvpbroadcastb []ytab

	_yvpbroadcastmb2q []ytab

	_yvpclmulqdq []ytab

	_yvpcmpb []ytab

	_yvpcmpeqb []ytab

	_yvperm2f128 []ytab

	_yvpermd []ytab

	_yvpermilpd []ytab

	_yvpermpd []ytab

	_yvpermq []ytab

	_yvpextrw []ytab

	_yvpinsrb []ytab

	_yvpmovb2m []ytab

	_yvpmovdb []ytab

	_yvpmovdw []ytab

	_yvprold []ytab

	_yvpscatterdd []ytab

	_yvpscatterdq []ytab

	_yvpscatterqd []ytab

	_yvpshufbitqmb []ytab

	_yvpshufd []ytab

	_yvpslld []ytab

	_yvpslldq []ytab

	_yvpsraq []ytab

	_yvptest []ytab

	_yvrcpss []ytab

	_yvroundpd []ytab

	_yvscalefpd []ytab

	_yvshuff32x4 []ytab

	_yvzeroall []ytab

	avxOptab [735]Optab

	bpduff1 []byte

	bpduff2     []byte
	deferreturn *obj.LSym

	evexSuffixMap [255]evexSuffix

	isAndroid bool

	naclbpfix []uint8

	naclmovs []uint8
	naclret  []uint8

	naclret8 []uint8

	naclspfix []uint8

	naclstos      []uint8
	nop           [][16]uint8
	opSuffixTable [15]string

	opindex       [1608]*Optab
	optab         [864]Optab
	plan9privates *obj.LSym

	reg [2248]int

	regrex [2249]int

	unaryDst map[obj.As]bool
	yaddl    []ytab

	yaes []ytab

	yblendvpd []ytab
	ybswap    []ytab
	ybtl      []ytab

	ybyte []ytab

	ycall    []ytab
	yclflush []ytab
	ycmpb    []ytab

	ycmpl []ytab

	ycompp []ytab
	ycover [8100]uint8

	ycrc32b []ytab
	ycrc32l []ytab
	ydivb   []ytab
	ydivl   []ytab

	yduff []ytab

	yextr []ytab

	yextractps []ytab
	yextrw     []ytab
	yfadd      []ytab
	yfcmv      []ytab
	yfmvd      []ytab

	yfmvdp []ytab

	yfmvf []ytab

	yfmvp     []ytab
	yfmvx     []ytab
	yfuncdata []ytab

	yfxch []ytab
	yimul []ytab

	yimul3 []ytab

	yin   []ytab
	yincl []ytab

	yincq []ytab

	yinsr  []ytab
	yinsrw []ytab
	yint   []ytab

	yjcond []ytab

	yjmp []ytab

	ylddqu []ytab
	yloop  []ytab
	ym_rl  []ytab

	ymb_rl []ytab

	yml_mb []ytab
	yml_rl []ytab

	ymm []ytab

	ymmxmm0f38 []ytab
	ymovb      []ytab

	ymovbe []ytab
	ymovl  []ytab

	ymovq []ytab

	ymovtab []Movtab
	ymovw   []ytab

	ymr []ytab

	ymr_ml []ytab

	ymrxr []ytab

	ymshuf []ytab

	ymshufb []ytab

	ymskb []ytab
	ynone []ytab

	ynop []ytab

	ypalignr []ytab
	ypcdata  []ytab

	ypopl []ytab

	yprefetch []ytab
	yps       []ytab

	ypsdq  []ytab
	ypushl []ytab
	yrb_mb []ytab

	yrdrand []ytab
	yret    []ytab

	yrl_m []ytab

	yrl_ml []ytab

	yscond []ytab

	ysha1rnds4   []ytab
	ysha256rnds2 []ytab
	yshb         []ytab

	yshl []ytab

	ystsw []ytab

	ysvrs_mo []ytab

	ysvrs_om []ytab
	ytestl   []ytab
	ytext    []ytab

	ywrfsbase []ytab

	yxabort []ytab
	yxbegin []ytab
	yxchg   []ytab

	yxcmpi []ytab

	yxcvfl []ytab

	yxcvfq []ytab
	yxcvlf []ytab
	yxcvm1 []ytab

	yxcvm2 []ytab

	yxcvqf []ytab
	yxm    []ytab

	yxm_q4 []ytab

	yxmov []ytab
	yxorb []ytab

	yxr []ytab

	yxr_ml []ytab

	yxrrl []ytab

	yxshuf []ytab
}

func NewPackageSession(obj_psess *obj.PackageSession, objabi_psess *objabi.PackageSession, src_psess *src.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.obj = obj_psess
	psess.objabi = objabi_psess
	psess.src = src_psess
	psess.sys = sys_psess
	psess.Register = []string{
		"AL",
		"CL",
		"DL",
		"BL",
		"SPB",
		"BPB",
		"SIB",
		"DIB",
		"R8B",
		"R9B",
		"R10B",
		"R11B",
		"R12B",
		"R13B",
		"R14B",
		"R15B",
		"AX",
		"CX",
		"DX",
		"BX",
		"SP",
		"BP",
		"SI",
		"DI",
		"R8",
		"R9",
		"R10",
		"R11",
		"R12",
		"R13",
		"R14",
		"R15",
		"AH",
		"CH",
		"DH",
		"BH",
		"F0",
		"F1",
		"F2",
		"F3",
		"F4",
		"F5",
		"F6",
		"F7",
		"M0",
		"M1",
		"M2",
		"M3",
		"M4",
		"M5",
		"M6",
		"M7",
		"K0",
		"K1",
		"K2",
		"K3",
		"K4",
		"K5",
		"K6",
		"K7",
		"X0",
		"X1",
		"X2",
		"X3",
		"X4",
		"X5",
		"X6",
		"X7",
		"X8",
		"X9",
		"X10",
		"X11",
		"X12",
		"X13",
		"X14",
		"X15",
		"X16",
		"X17",
		"X18",
		"X19",
		"X20",
		"X21",
		"X22",
		"X23",
		"X24",
		"X25",
		"X26",
		"X27",
		"X28",
		"X29",
		"X30",
		"X31",
		"Y0",
		"Y1",
		"Y2",
		"Y3",
		"Y4",
		"Y5",
		"Y6",
		"Y7",
		"Y8",
		"Y9",
		"Y10",
		"Y11",
		"Y12",
		"Y13",
		"Y14",
		"Y15",
		"Y16",
		"Y17",
		"Y18",
		"Y19",
		"Y20",
		"Y21",
		"Y22",
		"Y23",
		"Y24",
		"Y25",
		"Y26",
		"Y27",
		"Y28",
		"Y29",
		"Y30",
		"Y31",
		"Z0",
		"Z1",
		"Z2",
		"Z3",
		"Z4",
		"Z5",
		"Z6",
		"Z7",
		"Z8",
		"Z9",
		"Z10",
		"Z11",
		"Z12",
		"Z13",
		"Z14",
		"Z15",
		"Z16",
		"Z17",
		"Z18",
		"Z19",
		"Z20",
		"Z21",
		"Z22",
		"Z23",
		"Z24",
		"Z25",
		"Z26",
		"Z27",
		"Z28",
		"Z29",
		"Z30",
		"Z31",
		"CS",
		"SS",
		"DS",
		"ES",
		"FS",
		"GS",
		"GDTR",
		"IDTR",
		"LDTR",
		"MSW",
		"TASK",
		"CR0",
		"CR1",
		"CR2",
		"CR3",
		"CR4",
		"CR5",
		"CR6",
		"CR7",
		"CR8",
		"CR9",
		"CR10",
		"CR11",
		"CR12",
		"CR13",
		"CR14",
		"CR15",
		"DR0",
		"DR1",
		"DR2",
		"DR3",
		"DR4",
		"DR5",
		"DR6",
		"DR7",
		"TR0",
		"TR1",
		"TR2",
		"TR3",
		"TR4",
		"TR5",
		"TR6",
		"TR7",
		"TLS",
		"MAXREG",
	}
	psess.AMD64DWARFRegisters = map[int16]int16{
		REG_AX:  0,
		REG_DX:  1,
		REG_CX:  2,
		REG_BX:  3,
		REG_SI:  4,
		REG_DI:  5,
		REG_BP:  6,
		REG_SP:  7,
		REG_R8:  8,
		REG_R9:  9,
		REG_R10: 10,
		REG_R11: 11,
		REG_R12: 12,
		REG_R13: 13,
		REG_R14: 14,
		REG_R15: 15,

		REG_X0: 17,
		REG_X1: 18,
		REG_X2: 19,
		REG_X3: 20,
		REG_X4: 21,
		REG_X5: 22,
		REG_X6: 23,
		REG_X7: 24,

		REG_X8:  25,
		REG_X9:  26,
		REG_X10: 27,
		REG_X11: 28,
		REG_X12: 29,
		REG_X13: 30,
		REG_X14: 31,
		REG_X15: 32,

		REG_F0: 33,
		REG_F1: 34,
		REG_F2: 35,
		REG_F3: 36,
		REG_F4: 37,
		REG_F5: 38,
		REG_F6: 39,
		REG_F7: 40,

		REG_M0: 41,
		REG_M1: 42,
		REG_M2: 43,
		REG_M3: 44,
		REG_M4: 45,
		REG_M5: 46,
		REG_M6: 47,
		REG_M7: 48,

		REG_ES: 50,
		REG_CS: 51,
		REG_SS: 52,
		REG_DS: 53,
		REG_FS: 54,
		REG_GS: 55,

		REG_TR:   62,
		REG_LDTR: 63,

		REG_X16: 67,
		REG_X17: 68,
		REG_X18: 69,
		REG_X19: 70,
		REG_X20: 71,
		REG_X21: 72,
		REG_X22: 73,
		REG_X23: 74,
		REG_X24: 75,
		REG_X25: 76,
		REG_X26: 77,
		REG_X27: 78,
		REG_X28: 79,
		REG_X29: 80,
		REG_X30: 81,
		REG_X31: 82,

		REG_K0: 118,
		REG_K1: 119,
		REG_K2: 120,
		REG_K3: 121,
		REG_K4: 122,
		REG_K5: 123,
		REG_K6: 124,
		REG_K7: 125,
	}
	psess.X86DWARFRegisters = map[int16]int16{
		REG_AX: 0,
		REG_CX: 1,
		REG_DX: 2,
		REG_BX: 3,
		REG_SP: 4,
		REG_BP: 5,
		REG_SI: 6,
		REG_DI: 7,

		REG_F0: 11,
		REG_F1: 12,
		REG_F2: 13,
		REG_F3: 14,
		REG_F4: 15,
		REG_F5: 16,
		REG_F6: 17,
		REG_F7: 18,

		REG_X0: 21,
		REG_X1: 22,
		REG_X2: 23,
		REG_X3: 24,
		REG_X4: 25,
		REG_X5: 26,
		REG_X6: 27,
		REG_X7: 28,

		REG_M0: 29,
		REG_M1: 30,
		REG_M2: 31,
		REG_M3: 32,
		REG_M4: 33,
		REG_M5: 34,
		REG_M6: 35,
		REG_M7: 36,

		REG_ES:   40,
		REG_CS:   41,
		REG_SS:   42,
		REG_DS:   43,
		REG_FS:   44,
		REG_GS:   45,
		REG_TR:   48,
		REG_LDTR: 49,
	}
	psess.Anames = []string{
		obj.A_ARCHSPECIFIC: "AAA",
		"AAD",
		"AAM",
		"AAS",
		"ADCB",
		"ADCL",
		"ADCQ",
		"ADCW",
		"ADCXL",
		"ADCXQ",
		"ADDB",
		"ADDL",
		"ADDPD",
		"ADDPS",
		"ADDQ",
		"ADDSD",
		"ADDSS",
		"ADDSUBPD",
		"ADDSUBPS",
		"ADDW",
		"ADJSP",
		"ADOXL",
		"ADOXQ",
		"AESDEC",
		"AESDECLAST",
		"AESENC",
		"AESENCLAST",
		"AESIMC",
		"AESKEYGENASSIST",
		"ANDB",
		"ANDL",
		"ANDNL",
		"ANDNPD",
		"ANDNPS",
		"ANDNQ",
		"ANDPD",
		"ANDPS",
		"ANDQ",
		"ANDW",
		"ARPL",
		"BEXTRL",
		"BEXTRQ",
		"BLENDPD",
		"BLENDPS",
		"BLENDVPD",
		"BLENDVPS",
		"BLSIL",
		"BLSIQ",
		"BLSMSKL",
		"BLSMSKQ",
		"BLSRL",
		"BLSRQ",
		"BOUNDL",
		"BOUNDW",
		"BSFL",
		"BSFQ",
		"BSFW",
		"BSRL",
		"BSRQ",
		"BSRW",
		"BSWAPL",
		"BSWAPQ",
		"BSWAPW",
		"BTCL",
		"BTCQ",
		"BTCW",
		"BTL",
		"BTQ",
		"BTRL",
		"BTRQ",
		"BTRW",
		"BTSL",
		"BTSQ",
		"BTSW",
		"BTW",
		"BYTE",
		"BZHIL",
		"BZHIQ",
		"CBW",
		"CDQ",
		"CDQE",
		"CLAC",
		"CLC",
		"CLD",
		"CLFLUSH",
		"CLFLUSHOPT",
		"CLI",
		"CLTS",
		"CMC",
		"CMOVLCC",
		"CMOVLCS",
		"CMOVLEQ",
		"CMOVLGE",
		"CMOVLGT",
		"CMOVLHI",
		"CMOVLLE",
		"CMOVLLS",
		"CMOVLLT",
		"CMOVLMI",
		"CMOVLNE",
		"CMOVLOC",
		"CMOVLOS",
		"CMOVLPC",
		"CMOVLPL",
		"CMOVLPS",
		"CMOVQCC",
		"CMOVQCS",
		"CMOVQEQ",
		"CMOVQGE",
		"CMOVQGT",
		"CMOVQHI",
		"CMOVQLE",
		"CMOVQLS",
		"CMOVQLT",
		"CMOVQMI",
		"CMOVQNE",
		"CMOVQOC",
		"CMOVQOS",
		"CMOVQPC",
		"CMOVQPL",
		"CMOVQPS",
		"CMOVWCC",
		"CMOVWCS",
		"CMOVWEQ",
		"CMOVWGE",
		"CMOVWGT",
		"CMOVWHI",
		"CMOVWLE",
		"CMOVWLS",
		"CMOVWLT",
		"CMOVWMI",
		"CMOVWNE",
		"CMOVWOC",
		"CMOVWOS",
		"CMOVWPC",
		"CMOVWPL",
		"CMOVWPS",
		"CMPB",
		"CMPL",
		"CMPPD",
		"CMPPS",
		"CMPQ",
		"CMPSB",
		"CMPSD",
		"CMPSL",
		"CMPSQ",
		"CMPSS",
		"CMPSW",
		"CMPW",
		"CMPXCHG16B",
		"CMPXCHG8B",
		"CMPXCHGB",
		"CMPXCHGL",
		"CMPXCHGQ",
		"CMPXCHGW",
		"COMISD",
		"COMISS",
		"CPUID",
		"CQO",
		"CRC32B",
		"CRC32L",
		"CRC32Q",
		"CRC32W",
		"CVTPD2PL",
		"CVTPD2PS",
		"CVTPL2PD",
		"CVTPL2PS",
		"CVTPS2PD",
		"CVTPS2PL",
		"CVTSD2SL",
		"CVTSD2SQ",
		"CVTSD2SS",
		"CVTSL2SD",
		"CVTSL2SS",
		"CVTSQ2SD",
		"CVTSQ2SS",
		"CVTSS2SD",
		"CVTSS2SL",
		"CVTSS2SQ",
		"CVTTPD2PL",
		"CVTTPS2PL",
		"CVTTSD2SL",
		"CVTTSD2SQ",
		"CVTTSS2SL",
		"CVTTSS2SQ",
		"CWD",
		"CWDE",
		"DAA",
		"DAS",
		"DECB",
		"DECL",
		"DECQ",
		"DECW",
		"DIVB",
		"DIVL",
		"DIVPD",
		"DIVPS",
		"DIVQ",
		"DIVSD",
		"DIVSS",
		"DIVW",
		"DPPD",
		"DPPS",
		"EMMS",
		"ENTER",
		"EXTRACTPS",
		"F2XM1",
		"FABS",
		"FADDD",
		"FADDDP",
		"FADDF",
		"FADDL",
		"FADDW",
		"FBLD",
		"FBSTP",
		"FCHS",
		"FCLEX",
		"FCMOVB",
		"FCMOVBE",
		"FCMOVCC",
		"FCMOVCS",
		"FCMOVE",
		"FCMOVEQ",
		"FCMOVHI",
		"FCMOVLS",
		"FCMOVNB",
		"FCMOVNBE",
		"FCMOVNE",
		"FCMOVNU",
		"FCMOVU",
		"FCMOVUN",
		"FCOMD",
		"FCOMDP",
		"FCOMDPP",
		"FCOMF",
		"FCOMFP",
		"FCOMI",
		"FCOMIP",
		"FCOML",
		"FCOMLP",
		"FCOMW",
		"FCOMWP",
		"FCOS",
		"FDECSTP",
		"FDIVD",
		"FDIVDP",
		"FDIVF",
		"FDIVL",
		"FDIVRD",
		"FDIVRDP",
		"FDIVRF",
		"FDIVRL",
		"FDIVRW",
		"FDIVW",
		"FFREE",
		"FINCSTP",
		"FINIT",
		"FLD1",
		"FLDCW",
		"FLDENV",
		"FLDL2E",
		"FLDL2T",
		"FLDLG2",
		"FLDLN2",
		"FLDPI",
		"FLDZ",
		"FMOVB",
		"FMOVBP",
		"FMOVD",
		"FMOVDP",
		"FMOVF",
		"FMOVFP",
		"FMOVL",
		"FMOVLP",
		"FMOVV",
		"FMOVVP",
		"FMOVW",
		"FMOVWP",
		"FMOVX",
		"FMOVXP",
		"FMULD",
		"FMULDP",
		"FMULF",
		"FMULL",
		"FMULW",
		"FNOP",
		"FPATAN",
		"FPREM",
		"FPREM1",
		"FPTAN",
		"FRNDINT",
		"FRSTOR",
		"FSAVE",
		"FSCALE",
		"FSIN",
		"FSINCOS",
		"FSQRT",
		"FSTCW",
		"FSTENV",
		"FSTSW",
		"FSUBD",
		"FSUBDP",
		"FSUBF",
		"FSUBL",
		"FSUBRD",
		"FSUBRDP",
		"FSUBRF",
		"FSUBRL",
		"FSUBRW",
		"FSUBW",
		"FTST",
		"FUCOM",
		"FUCOMI",
		"FUCOMIP",
		"FUCOMP",
		"FUCOMPP",
		"FXAM",
		"FXCHD",
		"FXRSTOR",
		"FXRSTOR64",
		"FXSAVE",
		"FXSAVE64",
		"FXTRACT",
		"FYL2X",
		"FYL2XP1",
		"HADDPD",
		"HADDPS",
		"HLT",
		"HSUBPD",
		"HSUBPS",
		"ICEBP",
		"IDIVB",
		"IDIVL",
		"IDIVQ",
		"IDIVW",
		"IMUL3L",
		"IMUL3Q",
		"IMUL3W",
		"IMULB",
		"IMULL",
		"IMULQ",
		"IMULW",
		"INB",
		"INCB",
		"INCL",
		"INCQ",
		"INCW",
		"INL",
		"INSB",
		"INSERTPS",
		"INSL",
		"INSW",
		"INT",
		"INTO",
		"INVD",
		"INVLPG",
		"INVPCID",
		"INW",
		"IRETL",
		"IRETQ",
		"IRETW",
		"JCC",
		"JCS",
		"JCXZL",
		"JCXZQ",
		"JCXZW",
		"JEQ",
		"JGE",
		"JGT",
		"JHI",
		"JLE",
		"JLS",
		"JLT",
		"JMI",
		"JNE",
		"JOC",
		"JOS",
		"JPC",
		"JPL",
		"JPS",
		"KADDB",
		"KADDD",
		"KADDQ",
		"KADDW",
		"KANDB",
		"KANDD",
		"KANDNB",
		"KANDND",
		"KANDNQ",
		"KANDNW",
		"KANDQ",
		"KANDW",
		"KMOVB",
		"KMOVD",
		"KMOVQ",
		"KMOVW",
		"KNOTB",
		"KNOTD",
		"KNOTQ",
		"KNOTW",
		"KORB",
		"KORD",
		"KORQ",
		"KORTESTB",
		"KORTESTD",
		"KORTESTQ",
		"KORTESTW",
		"KORW",
		"KSHIFTLB",
		"KSHIFTLD",
		"KSHIFTLQ",
		"KSHIFTLW",
		"KSHIFTRB",
		"KSHIFTRD",
		"KSHIFTRQ",
		"KSHIFTRW",
		"KTESTB",
		"KTESTD",
		"KTESTQ",
		"KTESTW",
		"KUNPCKBW",
		"KUNPCKDQ",
		"KUNPCKWD",
		"KXNORB",
		"KXNORD",
		"KXNORQ",
		"KXNORW",
		"KXORB",
		"KXORD",
		"KXORQ",
		"KXORW",
		"LAHF",
		"LARL",
		"LARQ",
		"LARW",
		"LDDQU",
		"LDMXCSR",
		"LEAL",
		"LEAQ",
		"LEAVEL",
		"LEAVEQ",
		"LEAVEW",
		"LEAW",
		"LFENCE",
		"LFSL",
		"LFSQ",
		"LFSW",
		"LGDT",
		"LGSL",
		"LGSQ",
		"LGSW",
		"LIDT",
		"LLDT",
		"LMSW",
		"LOCK",
		"LODSB",
		"LODSL",
		"LODSQ",
		"LODSW",
		"LONG",
		"LOOP",
		"LOOPEQ",
		"LOOPNE",
		"LSLL",
		"LSLQ",
		"LSLW",
		"LSSL",
		"LSSQ",
		"LSSW",
		"LTR",
		"LZCNTL",
		"LZCNTQ",
		"LZCNTW",
		"MASKMOVOU",
		"MASKMOVQ",
		"MAXPD",
		"MAXPS",
		"MAXSD",
		"MAXSS",
		"MFENCE",
		"MINPD",
		"MINPS",
		"MINSD",
		"MINSS",
		"MONITOR",
		"MOVAPD",
		"MOVAPS",
		"MOVB",
		"MOVBELL",
		"MOVBEQQ",
		"MOVBEWW",
		"MOVBLSX",
		"MOVBLZX",
		"MOVBQSX",
		"MOVBQZX",
		"MOVBWSX",
		"MOVBWZX",
		"MOVDDUP",
		"MOVHLPS",
		"MOVHPD",
		"MOVHPS",
		"MOVL",
		"MOVLHPS",
		"MOVLPD",
		"MOVLPS",
		"MOVLQSX",
		"MOVLQZX",
		"MOVMSKPD",
		"MOVMSKPS",
		"MOVNTDQA",
		"MOVNTIL",
		"MOVNTIQ",
		"MOVNTO",
		"MOVNTPD",
		"MOVNTPS",
		"MOVNTQ",
		"MOVO",
		"MOVOU",
		"MOVQ",
		"MOVQL",
		"MOVQOZX",
		"MOVSB",
		"MOVSD",
		"MOVSHDUP",
		"MOVSL",
		"MOVSLDUP",
		"MOVSQ",
		"MOVSS",
		"MOVSW",
		"MOVSWW",
		"MOVUPD",
		"MOVUPS",
		"MOVW",
		"MOVWLSX",
		"MOVWLZX",
		"MOVWQSX",
		"MOVWQZX",
		"MOVZWW",
		"MPSADBW",
		"MULB",
		"MULL",
		"MULPD",
		"MULPS",
		"MULQ",
		"MULSD",
		"MULSS",
		"MULW",
		"MULXL",
		"MULXQ",
		"MWAIT",
		"NEGB",
		"NEGL",
		"NEGQ",
		"NEGW",
		"NOPL",
		"NOPW",
		"NOTB",
		"NOTL",
		"NOTQ",
		"NOTW",
		"ORB",
		"ORL",
		"ORPD",
		"ORPS",
		"ORQ",
		"ORW",
		"OUTB",
		"OUTL",
		"OUTSB",
		"OUTSL",
		"OUTSW",
		"OUTW",
		"PABSB",
		"PABSD",
		"PABSW",
		"PACKSSLW",
		"PACKSSWB",
		"PACKUSDW",
		"PACKUSWB",
		"PADDB",
		"PADDL",
		"PADDQ",
		"PADDSB",
		"PADDSW",
		"PADDUSB",
		"PADDUSW",
		"PADDW",
		"PALIGNR",
		"PAND",
		"PANDN",
		"PAUSE",
		"PAVGB",
		"PAVGW",
		"PBLENDVB",
		"PBLENDW",
		"PCLMULQDQ",
		"PCMPEQB",
		"PCMPEQL",
		"PCMPEQQ",
		"PCMPEQW",
		"PCMPESTRI",
		"PCMPESTRM",
		"PCMPGTB",
		"PCMPGTL",
		"PCMPGTQ",
		"PCMPGTW",
		"PCMPISTRI",
		"PCMPISTRM",
		"PDEPL",
		"PDEPQ",
		"PEXTL",
		"PEXTQ",
		"PEXTRB",
		"PEXTRD",
		"PEXTRQ",
		"PEXTRW",
		"PHADDD",
		"PHADDSW",
		"PHADDW",
		"PHMINPOSUW",
		"PHSUBD",
		"PHSUBSW",
		"PHSUBW",
		"PINSRB",
		"PINSRD",
		"PINSRQ",
		"PINSRW",
		"PMADDUBSW",
		"PMADDWL",
		"PMAXSB",
		"PMAXSD",
		"PMAXSW",
		"PMAXUB",
		"PMAXUD",
		"PMAXUW",
		"PMINSB",
		"PMINSD",
		"PMINSW",
		"PMINUB",
		"PMINUD",
		"PMINUW",
		"PMOVMSKB",
		"PMOVSXBD",
		"PMOVSXBQ",
		"PMOVSXBW",
		"PMOVSXDQ",
		"PMOVSXWD",
		"PMOVSXWQ",
		"PMOVZXBD",
		"PMOVZXBQ",
		"PMOVZXBW",
		"PMOVZXDQ",
		"PMOVZXWD",
		"PMOVZXWQ",
		"PMULDQ",
		"PMULHRSW",
		"PMULHUW",
		"PMULHW",
		"PMULLD",
		"PMULLW",
		"PMULULQ",
		"POPAL",
		"POPAW",
		"POPCNTL",
		"POPCNTQ",
		"POPCNTW",
		"POPFL",
		"POPFQ",
		"POPFW",
		"POPL",
		"POPQ",
		"POPW",
		"POR",
		"PREFETCHNTA",
		"PREFETCHT0",
		"PREFETCHT1",
		"PREFETCHT2",
		"PSADBW",
		"PSHUFB",
		"PSHUFD",
		"PSHUFHW",
		"PSHUFL",
		"PSHUFLW",
		"PSHUFW",
		"PSIGNB",
		"PSIGND",
		"PSIGNW",
		"PSLLL",
		"PSLLO",
		"PSLLQ",
		"PSLLW",
		"PSRAL",
		"PSRAW",
		"PSRLL",
		"PSRLO",
		"PSRLQ",
		"PSRLW",
		"PSUBB",
		"PSUBL",
		"PSUBQ",
		"PSUBSB",
		"PSUBSW",
		"PSUBUSB",
		"PSUBUSW",
		"PSUBW",
		"PTEST",
		"PUNPCKHBW",
		"PUNPCKHLQ",
		"PUNPCKHQDQ",
		"PUNPCKHWL",
		"PUNPCKLBW",
		"PUNPCKLLQ",
		"PUNPCKLQDQ",
		"PUNPCKLWL",
		"PUSHAL",
		"PUSHAW",
		"PUSHFL",
		"PUSHFQ",
		"PUSHFW",
		"PUSHL",
		"PUSHQ",
		"PUSHW",
		"PXOR",
		"QUAD",
		"RCLB",
		"RCLL",
		"RCLQ",
		"RCLW",
		"RCPPS",
		"RCPSS",
		"RCRB",
		"RCRL",
		"RCRQ",
		"RCRW",
		"RDFSBASEL",
		"RDFSBASEQ",
		"RDGSBASEL",
		"RDGSBASEQ",
		"RDMSR",
		"RDPKRU",
		"RDPMC",
		"RDRANDL",
		"RDRANDQ",
		"RDRANDW",
		"RDSEEDL",
		"RDSEEDQ",
		"RDSEEDW",
		"RDTSC",
		"RDTSCP",
		"REP",
		"REPN",
		"RETFL",
		"RETFQ",
		"RETFW",
		"ROLB",
		"ROLL",
		"ROLQ",
		"ROLW",
		"RORB",
		"RORL",
		"RORQ",
		"RORW",
		"RORXL",
		"RORXQ",
		"ROUNDPD",
		"ROUNDPS",
		"ROUNDSD",
		"ROUNDSS",
		"RSM",
		"RSQRTPS",
		"RSQRTSS",
		"SAHF",
		"SALB",
		"SALL",
		"SALQ",
		"SALW",
		"SARB",
		"SARL",
		"SARQ",
		"SARW",
		"SARXL",
		"SARXQ",
		"SBBB",
		"SBBL",
		"SBBQ",
		"SBBW",
		"SCASB",
		"SCASL",
		"SCASQ",
		"SCASW",
		"SETCC",
		"SETCS",
		"SETEQ",
		"SETGE",
		"SETGT",
		"SETHI",
		"SETLE",
		"SETLS",
		"SETLT",
		"SETMI",
		"SETNE",
		"SETOC",
		"SETOS",
		"SETPC",
		"SETPL",
		"SETPS",
		"SFENCE",
		"SGDT",
		"SHA1MSG1",
		"SHA1MSG2",
		"SHA1NEXTE",
		"SHA1RNDS4",
		"SHA256MSG1",
		"SHA256MSG2",
		"SHA256RNDS2",
		"SHLB",
		"SHLL",
		"SHLQ",
		"SHLW",
		"SHLXL",
		"SHLXQ",
		"SHRB",
		"SHRL",
		"SHRQ",
		"SHRW",
		"SHRXL",
		"SHRXQ",
		"SHUFPD",
		"SHUFPS",
		"SIDT",
		"SLDTL",
		"SLDTQ",
		"SLDTW",
		"SMSWL",
		"SMSWQ",
		"SMSWW",
		"SQRTPD",
		"SQRTPS",
		"SQRTSD",
		"SQRTSS",
		"STAC",
		"STC",
		"STD",
		"STI",
		"STMXCSR",
		"STOSB",
		"STOSL",
		"STOSQ",
		"STOSW",
		"STRL",
		"STRQ",
		"STRW",
		"SUBB",
		"SUBL",
		"SUBPD",
		"SUBPS",
		"SUBQ",
		"SUBSD",
		"SUBSS",
		"SUBW",
		"SWAPGS",
		"SYSCALL",
		"SYSENTER",
		"SYSENTER64",
		"SYSEXIT",
		"SYSEXIT64",
		"SYSRET",
		"TESTB",
		"TESTL",
		"TESTQ",
		"TESTW",
		"TZCNTL",
		"TZCNTQ",
		"TZCNTW",
		"UCOMISD",
		"UCOMISS",
		"UD1",
		"UD2",
		"UNPCKHPD",
		"UNPCKHPS",
		"UNPCKLPD",
		"UNPCKLPS",
		"V4FMADDPS",
		"V4FMADDSS",
		"V4FNMADDPS",
		"V4FNMADDSS",
		"VADDPD",
		"VADDPS",
		"VADDSD",
		"VADDSS",
		"VADDSUBPD",
		"VADDSUBPS",
		"VAESDEC",
		"VAESDECLAST",
		"VAESENC",
		"VAESENCLAST",
		"VAESIMC",
		"VAESKEYGENASSIST",
		"VALIGND",
		"VALIGNQ",
		"VANDNPD",
		"VANDNPS",
		"VANDPD",
		"VANDPS",
		"VBLENDMPD",
		"VBLENDMPS",
		"VBLENDPD",
		"VBLENDPS",
		"VBLENDVPD",
		"VBLENDVPS",
		"VBROADCASTF128",
		"VBROADCASTF32X2",
		"VBROADCASTF32X4",
		"VBROADCASTF32X8",
		"VBROADCASTF64X2",
		"VBROADCASTF64X4",
		"VBROADCASTI128",
		"VBROADCASTI32X2",
		"VBROADCASTI32X4",
		"VBROADCASTI32X8",
		"VBROADCASTI64X2",
		"VBROADCASTI64X4",
		"VBROADCASTSD",
		"VBROADCASTSS",
		"VCMPPD",
		"VCMPPS",
		"VCMPSD",
		"VCMPSS",
		"VCOMISD",
		"VCOMISS",
		"VCOMPRESSPD",
		"VCOMPRESSPS",
		"VCVTDQ2PD",
		"VCVTDQ2PS",
		"VCVTPD2DQ",
		"VCVTPD2DQX",
		"VCVTPD2DQY",
		"VCVTPD2PS",
		"VCVTPD2PSX",
		"VCVTPD2PSY",
		"VCVTPD2QQ",
		"VCVTPD2UDQ",
		"VCVTPD2UDQX",
		"VCVTPD2UDQY",
		"VCVTPD2UQQ",
		"VCVTPH2PS",
		"VCVTPS2DQ",
		"VCVTPS2PD",
		"VCVTPS2PH",
		"VCVTPS2QQ",
		"VCVTPS2UDQ",
		"VCVTPS2UQQ",
		"VCVTQQ2PD",
		"VCVTQQ2PS",
		"VCVTQQ2PSX",
		"VCVTQQ2PSY",
		"VCVTSD2SI",
		"VCVTSD2SIQ",
		"VCVTSD2SS",
		"VCVTSD2USI",
		"VCVTSD2USIL",
		"VCVTSD2USIQ",
		"VCVTSI2SDL",
		"VCVTSI2SDQ",
		"VCVTSI2SSL",
		"VCVTSI2SSQ",
		"VCVTSS2SD",
		"VCVTSS2SI",
		"VCVTSS2SIQ",
		"VCVTSS2USI",
		"VCVTSS2USIL",
		"VCVTSS2USIQ",
		"VCVTTPD2DQ",
		"VCVTTPD2DQX",
		"VCVTTPD2DQY",
		"VCVTTPD2QQ",
		"VCVTTPD2UDQ",
		"VCVTTPD2UDQX",
		"VCVTTPD2UDQY",
		"VCVTTPD2UQQ",
		"VCVTTPS2DQ",
		"VCVTTPS2QQ",
		"VCVTTPS2UDQ",
		"VCVTTPS2UQQ",
		"VCVTTSD2SI",
		"VCVTTSD2SIQ",
		"VCVTTSD2USI",
		"VCVTTSD2USIL",
		"VCVTTSD2USIQ",
		"VCVTTSS2SI",
		"VCVTTSS2SIQ",
		"VCVTTSS2USI",
		"VCVTTSS2USIL",
		"VCVTTSS2USIQ",
		"VCVTUDQ2PD",
		"VCVTUDQ2PS",
		"VCVTUQQ2PD",
		"VCVTUQQ2PS",
		"VCVTUQQ2PSX",
		"VCVTUQQ2PSY",
		"VCVTUSI2SD",
		"VCVTUSI2SDL",
		"VCVTUSI2SDQ",
		"VCVTUSI2SS",
		"VCVTUSI2SSL",
		"VCVTUSI2SSQ",
		"VDBPSADBW",
		"VDIVPD",
		"VDIVPS",
		"VDIVSD",
		"VDIVSS",
		"VDPPD",
		"VDPPS",
		"VERR",
		"VERW",
		"VEXP2PD",
		"VEXP2PS",
		"VEXPANDPD",
		"VEXPANDPS",
		"VEXTRACTF128",
		"VEXTRACTF32X4",
		"VEXTRACTF32X8",
		"VEXTRACTF64X2",
		"VEXTRACTF64X4",
		"VEXTRACTI128",
		"VEXTRACTI32X4",
		"VEXTRACTI32X8",
		"VEXTRACTI64X2",
		"VEXTRACTI64X4",
		"VEXTRACTPS",
		"VFIXUPIMMPD",
		"VFIXUPIMMPS",
		"VFIXUPIMMSD",
		"VFIXUPIMMSS",
		"VFMADD132PD",
		"VFMADD132PS",
		"VFMADD132SD",
		"VFMADD132SS",
		"VFMADD213PD",
		"VFMADD213PS",
		"VFMADD213SD",
		"VFMADD213SS",
		"VFMADD231PD",
		"VFMADD231PS",
		"VFMADD231SD",
		"VFMADD231SS",
		"VFMADDSUB132PD",
		"VFMADDSUB132PS",
		"VFMADDSUB213PD",
		"VFMADDSUB213PS",
		"VFMADDSUB231PD",
		"VFMADDSUB231PS",
		"VFMSUB132PD",
		"VFMSUB132PS",
		"VFMSUB132SD",
		"VFMSUB132SS",
		"VFMSUB213PD",
		"VFMSUB213PS",
		"VFMSUB213SD",
		"VFMSUB213SS",
		"VFMSUB231PD",
		"VFMSUB231PS",
		"VFMSUB231SD",
		"VFMSUB231SS",
		"VFMSUBADD132PD",
		"VFMSUBADD132PS",
		"VFMSUBADD213PD",
		"VFMSUBADD213PS",
		"VFMSUBADD231PD",
		"VFMSUBADD231PS",
		"VFNMADD132PD",
		"VFNMADD132PS",
		"VFNMADD132SD",
		"VFNMADD132SS",
		"VFNMADD213PD",
		"VFNMADD213PS",
		"VFNMADD213SD",
		"VFNMADD213SS",
		"VFNMADD231PD",
		"VFNMADD231PS",
		"VFNMADD231SD",
		"VFNMADD231SS",
		"VFNMSUB132PD",
		"VFNMSUB132PS",
		"VFNMSUB132SD",
		"VFNMSUB132SS",
		"VFNMSUB213PD",
		"VFNMSUB213PS",
		"VFNMSUB213SD",
		"VFNMSUB213SS",
		"VFNMSUB231PD",
		"VFNMSUB231PS",
		"VFNMSUB231SD",
		"VFNMSUB231SS",
		"VFPCLASSPD",
		"VFPCLASSPDX",
		"VFPCLASSPDY",
		"VFPCLASSPDZ",
		"VFPCLASSPS",
		"VFPCLASSPSX",
		"VFPCLASSPSY",
		"VFPCLASSPSZ",
		"VFPCLASSSD",
		"VFPCLASSSS",
		"VGATHERDPD",
		"VGATHERDPS",
		"VGATHERPF0DPD",
		"VGATHERPF0DPS",
		"VGATHERPF0QPD",
		"VGATHERPF0QPS",
		"VGATHERPF1DPD",
		"VGATHERPF1DPS",
		"VGATHERPF1QPD",
		"VGATHERPF1QPS",
		"VGATHERQPD",
		"VGATHERQPS",
		"VGETEXPPD",
		"VGETEXPPS",
		"VGETEXPSD",
		"VGETEXPSS",
		"VGETMANTPD",
		"VGETMANTPS",
		"VGETMANTSD",
		"VGETMANTSS",
		"VGF2P8AFFINEINVQB",
		"VGF2P8AFFINEQB",
		"VGF2P8MULB",
		"VHADDPD",
		"VHADDPS",
		"VHSUBPD",
		"VHSUBPS",
		"VINSERTF128",
		"VINSERTF32X4",
		"VINSERTF32X8",
		"VINSERTF64X2",
		"VINSERTF64X4",
		"VINSERTI128",
		"VINSERTI32X4",
		"VINSERTI32X8",
		"VINSERTI64X2",
		"VINSERTI64X4",
		"VINSERTPS",
		"VLDDQU",
		"VLDMXCSR",
		"VMASKMOVDQU",
		"VMASKMOVPD",
		"VMASKMOVPS",
		"VMAXPD",
		"VMAXPS",
		"VMAXSD",
		"VMAXSS",
		"VMINPD",
		"VMINPS",
		"VMINSD",
		"VMINSS",
		"VMOVAPD",
		"VMOVAPS",
		"VMOVD",
		"VMOVDDUP",
		"VMOVDQA",
		"VMOVDQA32",
		"VMOVDQA64",
		"VMOVDQU",
		"VMOVDQU16",
		"VMOVDQU32",
		"VMOVDQU64",
		"VMOVDQU8",
		"VMOVHLPS",
		"VMOVHPD",
		"VMOVHPS",
		"VMOVLHPS",
		"VMOVLPD",
		"VMOVLPS",
		"VMOVMSKPD",
		"VMOVMSKPS",
		"VMOVNTDQ",
		"VMOVNTDQA",
		"VMOVNTPD",
		"VMOVNTPS",
		"VMOVQ",
		"VMOVSD",
		"VMOVSHDUP",
		"VMOVSLDUP",
		"VMOVSS",
		"VMOVUPD",
		"VMOVUPS",
		"VMPSADBW",
		"VMULPD",
		"VMULPS",
		"VMULSD",
		"VMULSS",
		"VORPD",
		"VORPS",
		"VP4DPWSSD",
		"VP4DPWSSDS",
		"VPABSB",
		"VPABSD",
		"VPABSQ",
		"VPABSW",
		"VPACKSSDW",
		"VPACKSSWB",
		"VPACKUSDW",
		"VPACKUSWB",
		"VPADDB",
		"VPADDD",
		"VPADDQ",
		"VPADDSB",
		"VPADDSW",
		"VPADDUSB",
		"VPADDUSW",
		"VPADDW",
		"VPALIGNR",
		"VPAND",
		"VPANDD",
		"VPANDN",
		"VPANDND",
		"VPANDNQ",
		"VPANDQ",
		"VPAVGB",
		"VPAVGW",
		"VPBLENDD",
		"VPBLENDMB",
		"VPBLENDMD",
		"VPBLENDMQ",
		"VPBLENDMW",
		"VPBLENDVB",
		"VPBLENDW",
		"VPBROADCASTB",
		"VPBROADCASTD",
		"VPBROADCASTMB2Q",
		"VPBROADCASTMW2D",
		"VPBROADCASTQ",
		"VPBROADCASTW",
		"VPCLMULQDQ",
		"VPCMPB",
		"VPCMPD",
		"VPCMPEQB",
		"VPCMPEQD",
		"VPCMPEQQ",
		"VPCMPEQW",
		"VPCMPESTRI",
		"VPCMPESTRM",
		"VPCMPGTB",
		"VPCMPGTD",
		"VPCMPGTQ",
		"VPCMPGTW",
		"VPCMPISTRI",
		"VPCMPISTRM",
		"VPCMPQ",
		"VPCMPUB",
		"VPCMPUD",
		"VPCMPUQ",
		"VPCMPUW",
		"VPCMPW",
		"VPCOMPRESSB",
		"VPCOMPRESSD",
		"VPCOMPRESSQ",
		"VPCOMPRESSW",
		"VPCONFLICTD",
		"VPCONFLICTQ",
		"VPDPBUSD",
		"VPDPBUSDS",
		"VPDPWSSD",
		"VPDPWSSDS",
		"VPERM2F128",
		"VPERM2I128",
		"VPERMB",
		"VPERMD",
		"VPERMI2B",
		"VPERMI2D",
		"VPERMI2PD",
		"VPERMI2PS",
		"VPERMI2Q",
		"VPERMI2W",
		"VPERMILPD",
		"VPERMILPS",
		"VPERMPD",
		"VPERMPS",
		"VPERMQ",
		"VPERMT2B",
		"VPERMT2D",
		"VPERMT2PD",
		"VPERMT2PS",
		"VPERMT2Q",
		"VPERMT2W",
		"VPERMW",
		"VPEXPANDB",
		"VPEXPANDD",
		"VPEXPANDQ",
		"VPEXPANDW",
		"VPEXTRB",
		"VPEXTRD",
		"VPEXTRQ",
		"VPEXTRW",
		"VPGATHERDD",
		"VPGATHERDQ",
		"VPGATHERQD",
		"VPGATHERQQ",
		"VPHADDD",
		"VPHADDSW",
		"VPHADDW",
		"VPHMINPOSUW",
		"VPHSUBD",
		"VPHSUBSW",
		"VPHSUBW",
		"VPINSRB",
		"VPINSRD",
		"VPINSRQ",
		"VPINSRW",
		"VPLZCNTD",
		"VPLZCNTQ",
		"VPMADD52HUQ",
		"VPMADD52LUQ",
		"VPMADDUBSW",
		"VPMADDWD",
		"VPMASKMOVD",
		"VPMASKMOVQ",
		"VPMAXSB",
		"VPMAXSD",
		"VPMAXSQ",
		"VPMAXSW",
		"VPMAXUB",
		"VPMAXUD",
		"VPMAXUQ",
		"VPMAXUW",
		"VPMINSB",
		"VPMINSD",
		"VPMINSQ",
		"VPMINSW",
		"VPMINUB",
		"VPMINUD",
		"VPMINUQ",
		"VPMINUW",
		"VPMOVB2M",
		"VPMOVD2M",
		"VPMOVDB",
		"VPMOVDW",
		"VPMOVM2B",
		"VPMOVM2D",
		"VPMOVM2Q",
		"VPMOVM2W",
		"VPMOVMSKB",
		"VPMOVQ2M",
		"VPMOVQB",
		"VPMOVQD",
		"VPMOVQW",
		"VPMOVSDB",
		"VPMOVSDW",
		"VPMOVSQB",
		"VPMOVSQD",
		"VPMOVSQW",
		"VPMOVSWB",
		"VPMOVSXBD",
		"VPMOVSXBQ",
		"VPMOVSXBW",
		"VPMOVSXDQ",
		"VPMOVSXWD",
		"VPMOVSXWQ",
		"VPMOVUSDB",
		"VPMOVUSDW",
		"VPMOVUSQB",
		"VPMOVUSQD",
		"VPMOVUSQW",
		"VPMOVUSWB",
		"VPMOVW2M",
		"VPMOVWB",
		"VPMOVZXBD",
		"VPMOVZXBQ",
		"VPMOVZXBW",
		"VPMOVZXDQ",
		"VPMOVZXWD",
		"VPMOVZXWQ",
		"VPMULDQ",
		"VPMULHRSW",
		"VPMULHUW",
		"VPMULHW",
		"VPMULLD",
		"VPMULLQ",
		"VPMULLW",
		"VPMULTISHIFTQB",
		"VPMULUDQ",
		"VPOPCNTB",
		"VPOPCNTD",
		"VPOPCNTQ",
		"VPOPCNTW",
		"VPOR",
		"VPORD",
		"VPORQ",
		"VPROLD",
		"VPROLQ",
		"VPROLVD",
		"VPROLVQ",
		"VPRORD",
		"VPRORQ",
		"VPRORVD",
		"VPRORVQ",
		"VPSADBW",
		"VPSCATTERDD",
		"VPSCATTERDQ",
		"VPSCATTERQD",
		"VPSCATTERQQ",
		"VPSHLDD",
		"VPSHLDQ",
		"VPSHLDVD",
		"VPSHLDVQ",
		"VPSHLDVW",
		"VPSHLDW",
		"VPSHRDD",
		"VPSHRDQ",
		"VPSHRDVD",
		"VPSHRDVQ",
		"VPSHRDVW",
		"VPSHRDW",
		"VPSHUFB",
		"VPSHUFBITQMB",
		"VPSHUFD",
		"VPSHUFHW",
		"VPSHUFLW",
		"VPSIGNB",
		"VPSIGND",
		"VPSIGNW",
		"VPSLLD",
		"VPSLLDQ",
		"VPSLLQ",
		"VPSLLVD",
		"VPSLLVQ",
		"VPSLLVW",
		"VPSLLW",
		"VPSRAD",
		"VPSRAQ",
		"VPSRAVD",
		"VPSRAVQ",
		"VPSRAVW",
		"VPSRAW",
		"VPSRLD",
		"VPSRLDQ",
		"VPSRLQ",
		"VPSRLVD",
		"VPSRLVQ",
		"VPSRLVW",
		"VPSRLW",
		"VPSUBB",
		"VPSUBD",
		"VPSUBQ",
		"VPSUBSB",
		"VPSUBSW",
		"VPSUBUSB",
		"VPSUBUSW",
		"VPSUBW",
		"VPTERNLOGD",
		"VPTERNLOGQ",
		"VPTEST",
		"VPTESTMB",
		"VPTESTMD",
		"VPTESTMQ",
		"VPTESTMW",
		"VPTESTNMB",
		"VPTESTNMD",
		"VPTESTNMQ",
		"VPTESTNMW",
		"VPUNPCKHBW",
		"VPUNPCKHDQ",
		"VPUNPCKHQDQ",
		"VPUNPCKHWD",
		"VPUNPCKLBW",
		"VPUNPCKLDQ",
		"VPUNPCKLQDQ",
		"VPUNPCKLWD",
		"VPXOR",
		"VPXORD",
		"VPXORQ",
		"VRANGEPD",
		"VRANGEPS",
		"VRANGESD",
		"VRANGESS",
		"VRCP14PD",
		"VRCP14PS",
		"VRCP14SD",
		"VRCP14SS",
		"VRCP28PD",
		"VRCP28PS",
		"VRCP28SD",
		"VRCP28SS",
		"VRCPPS",
		"VRCPSS",
		"VREDUCEPD",
		"VREDUCEPS",
		"VREDUCESD",
		"VREDUCESS",
		"VRNDSCALEPD",
		"VRNDSCALEPS",
		"VRNDSCALESD",
		"VRNDSCALESS",
		"VROUNDPD",
		"VROUNDPS",
		"VROUNDSD",
		"VROUNDSS",
		"VRSQRT14PD",
		"VRSQRT14PS",
		"VRSQRT14SD",
		"VRSQRT14SS",
		"VRSQRT28PD",
		"VRSQRT28PS",
		"VRSQRT28SD",
		"VRSQRT28SS",
		"VRSQRTPS",
		"VRSQRTSS",
		"VSCALEFPD",
		"VSCALEFPS",
		"VSCALEFSD",
		"VSCALEFSS",
		"VSCATTERDPD",
		"VSCATTERDPS",
		"VSCATTERPF0DPD",
		"VSCATTERPF0DPS",
		"VSCATTERPF0QPD",
		"VSCATTERPF0QPS",
		"VSCATTERPF1DPD",
		"VSCATTERPF1DPS",
		"VSCATTERPF1QPD",
		"VSCATTERPF1QPS",
		"VSCATTERQPD",
		"VSCATTERQPS",
		"VSHUFF32X4",
		"VSHUFF64X2",
		"VSHUFI32X4",
		"VSHUFI64X2",
		"VSHUFPD",
		"VSHUFPS",
		"VSQRTPD",
		"VSQRTPS",
		"VSQRTSD",
		"VSQRTSS",
		"VSTMXCSR",
		"VSUBPD",
		"VSUBPS",
		"VSUBSD",
		"VSUBSS",
		"VTESTPD",
		"VTESTPS",
		"VUCOMISD",
		"VUCOMISS",
		"VUNPCKHPD",
		"VUNPCKHPS",
		"VUNPCKLPD",
		"VUNPCKLPS",
		"VXORPD",
		"VXORPS",
		"VZEROALL",
		"VZEROUPPER",
		"WAIT",
		"WBINVD",
		"WORD",
		"WRFSBASEL",
		"WRFSBASEQ",
		"WRGSBASEL",
		"WRGSBASEQ",
		"WRMSR",
		"WRPKRU",
		"XABORT",
		"XACQUIRE",
		"XADDB",
		"XADDL",
		"XADDQ",
		"XADDW",
		"XBEGIN",
		"XCHGB",
		"XCHGL",
		"XCHGQ",
		"XCHGW",
		"XEND",
		"XGETBV",
		"XLAT",
		"XORB",
		"XORL",
		"XORPD",
		"XORPS",
		"XORQ",
		"XORW",
		"XRELEASE",
		"XRSTOR",
		"XRSTOR64",
		"XRSTORS",
		"XRSTORS64",
		"XSAVE",
		"XSAVE64",
		"XSAVEC",
		"XSAVEC64",
		"XSAVEOPT",
		"XSAVEOPT64",
		"XSAVES",
		"XSAVES64",
		"XSETBV",
		"XTEST",
		"LAST",
	}
	psess.ynone = []ytab{
		{Zlit, 1, argList{}},
	}
	psess.ytext = []ytab{
		{Zpseudo, 0, argList{Ymb, Ytextsize}},
		{Zpseudo, 1, argList{Ymb, Yi32, Ytextsize}},
	}
	psess.ynop = []ytab{
		{Zpseudo, 0, argList{}},
		{Zpseudo, 0, argList{Yiauto}},
		{Zpseudo, 0, argList{Yml}},
		{Zpseudo, 0, argList{Yrf}},
		{Zpseudo, 0, argList{Yxr}},
		{Zpseudo, 0, argList{Yiauto}},
		{Zpseudo, 0, argList{Yml}},
		{Zpseudo, 0, argList{Yrf}},
		{Zpseudo, 1, argList{Yxr}},
	}
	psess.yfuncdata = []ytab{
		{Zpseudo, 0, argList{Yi32, Ym}},
	}
	psess.ypcdata = []ytab{
		{Zpseudo, 0, argList{Yi32, Yi32}},
	}
	psess.yxorb = []ytab{
		{Zib_, 1, argList{Yi32, Yal}},
		{Zibo_m, 2, argList{Yi32, Ymb}},
		{Zr_m, 1, argList{Yrb, Ymb}},
		{Zm_r, 1, argList{Ymb, Yrb}},
	}
	psess.yaddl = []ytab{
		{Zibo_m, 2, argList{Yi8, Yml}},
		{Zil_, 1, argList{Yi32, Yax}},
		{Zilo_m, 2, argList{Yi32, Yml}},
		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
	}
	psess.yincl = []ytab{
		{Z_rp, 1, argList{Yrl}},
		{Zo_m, 2, argList{Yml}},
	}
	psess.yincq = []ytab{
		{Zo_m, 2, argList{Yml}},
	}
	psess.ycmpb = []ytab{
		{Z_ib, 1, argList{Yal, Yi32}},
		{Zm_ibo, 2, argList{Ymb, Yi32}},
		{Zm_r, 1, argList{Ymb, Yrb}},
		{Zr_m, 1, argList{Yrb, Ymb}},
	}
	psess.ycmpl = []ytab{
		{Zm_ibo, 2, argList{Yml, Yi8}},
		{Z_il, 1, argList{Yax, Yi32}},
		{Zm_ilo, 2, argList{Yml, Yi32}},
		{Zm_r, 1, argList{Yml, Yrl}},
		{Zr_m, 1, argList{Yrl, Yml}},
	}
	psess.yshb = []ytab{
		{Zo_m, 2, argList{Yi1, Ymb}},
		{Zibo_m, 2, argList{Yu8, Ymb}},
		{Zo_m, 2, argList{Ycx, Ymb}},
	}
	psess.yshl = []ytab{
		{Zo_m, 2, argList{Yi1, Yml}},
		{Zibo_m, 2, argList{Yu8, Yml}},
		{Zo_m, 2, argList{Ycl, Yml}},
		{Zo_m, 2, argList{Ycx, Yml}},
	}
	psess.ytestl = []ytab{
		{Zil_, 1, argList{Yi32, Yax}},
		{Zilo_m, 2, argList{Yi32, Yml}},
		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
	}
	psess.ymovb = []ytab{
		{Zr_m, 1, argList{Yrb, Ymb}},
		{Zm_r, 1, argList{Ymb, Yrb}},
		{Zib_rp, 1, argList{Yi32, Yrb}},
		{Zibo_m, 2, argList{Yi32, Ymb}},
	}
	psess.ybtl = []ytab{
		{Zibo_m, 2, argList{Yi8, Yml}},
		{Zr_m, 1, argList{Yrl, Yml}},
	}
	psess.ymovw = []ytab{
		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
		{Zil_rp, 1, argList{Yi32, Yrl}},
		{Zilo_m, 2, argList{Yi32, Yml}},
		{Zaut_r, 2, argList{Yiauto, Yrl}},
	}
	psess.ymovl = []ytab{
		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
		{Zil_rp, 1, argList{Yi32, Yrl}},
		{Zilo_m, 2, argList{Yi32, Yml}},
		{Zm_r_xm, 1, argList{Yml, Ymr}},
		{Zr_m_xm, 1, argList{Ymr, Yml}},
		{Zm_r_xm, 2, argList{Yml, Yxr}},
		{Zr_m_xm, 2, argList{Yxr, Yml}},
		{Zaut_r, 2, argList{Yiauto, Yrl}},
	}
	psess.yret = []ytab{
		{Zo_iw, 1, argList{}},
		{Zo_iw, 1, argList{Yi32}},
	}
	psess.ymovq = []ytab{

		{Zm_r_xm_nr, 1, argList{Ym, Ymr}},
		{Zr_m_xm_nr, 1, argList{Ymr, Ym}},
		{Zm_r_xm_nr, 2, argList{Yxr, Ymr}},
		{Zm_r_xm_nr, 2, argList{Yxm, Yxr}},
		{Zr_m_xm_nr, 2, argList{Yxr, Yxm}},

		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
		{Zilo_m, 2, argList{Ys32, Yrl}},
		{Ziq_rp, 1, argList{Yi64, Yrl}},
		{Zilo_m, 2, argList{Yi32, Yml}},
		{Zm_r_xm, 1, argList{Ymm, Ymr}},
		{Zr_m_xm, 1, argList{Ymr, Ymm}},
		{Zm_r_xm, 2, argList{Yml, Yxr}},
		{Zr_m_xm, 2, argList{Yxr, Yml}},
		{Zaut_r, 1, argList{Yiauto, Yrl}},
	}
	psess.ymovbe = []ytab{
		{Zlitm_r, 3, argList{Ym, Yrl}},
		{Zlitr_m, 3, argList{Yrl, Ym}},
	}
	psess.ym_rl = []ytab{
		{Zm_r, 1, argList{Ym, Yrl}},
	}
	psess.yrl_m = []ytab{
		{Zr_m, 1, argList{Yrl, Ym}},
	}
	psess.ymb_rl = []ytab{
		{Zmb_r, 1, argList{Ymb, Yrl}},
	}
	psess.yml_rl = []ytab{
		{Zm_r, 1, argList{Yml, Yrl}},
	}
	psess.yrl_ml = []ytab{
		{Zr_m, 1, argList{Yrl, Yml}},
	}
	psess.yml_mb = []ytab{
		{Zr_m, 1, argList{Yrb, Ymb}},
		{Zm_r, 1, argList{Ymb, Yrb}},
	}
	psess.yrb_mb = []ytab{
		{Zr_m, 1, argList{Yrb, Ymb}},
	}
	psess.yxchg = []ytab{
		{Z_rp, 1, argList{Yax, Yrl}},
		{Zrp_, 1, argList{Yrl, Yax}},
		{Zr_m, 1, argList{Yrl, Yml}},
		{Zm_r, 1, argList{Yml, Yrl}},
	}
	psess.ydivl = []ytab{
		{Zm_o, 2, argList{Yml}},
	}
	psess.ydivb = []ytab{
		{Zm_o, 2, argList{Ymb}},
	}
	psess.yimul = []ytab{
		{Zm_o, 2, argList{Yml}},
		{Zib_rr, 1, argList{Yi8, Yrl}},
		{Zil_rr, 1, argList{Yi32, Yrl}},
		{Zm_r, 2, argList{Yml, Yrl}},
	}
	psess.yimul3 = []ytab{
		{Zibm_r, 2, argList{Yi8, Yml, Yrl}},
		{Zibm_r, 2, argList{Yi32, Yml, Yrl}},
	}
	psess.ybyte = []ytab{
		{Zbyte, 1, argList{Yi64}},
	}
	psess.yin = []ytab{
		{Zib_, 1, argList{Yi32}},
		{Zlit, 1, argList{}},
	}
	psess.yint = []ytab{
		{Zib_, 1, argList{Yi32}},
	}
	psess.ypushl = []ytab{
		{Zrp_, 1, argList{Yrl}},
		{Zm_o, 2, argList{Ym}},
		{Zib_, 1, argList{Yi8}},
		{Zil_, 1, argList{Yi32}},
	}
	psess.ypopl = []ytab{
		{Z_rp, 1, argList{Yrl}},
		{Zo_m, 2, argList{Ym}},
	}
	psess.ywrfsbase = []ytab{
		{Zm_o, 2, argList{Yrl}},
	}
	psess.yrdrand = []ytab{
		{Zo_m, 2, argList{Yrl}},
	}
	psess.yclflush = []ytab{
		{Zo_m, 2, argList{Ym}},
	}
	psess.ybswap = []ytab{
		{Z_rp, 2, argList{Yrl}},
	}
	psess.yscond = []ytab{
		{Zo_m, 2, argList{Ymb}},
	}
	psess.yjcond = []ytab{
		{Zbr, 0, argList{Ybr}},
		{Zbr, 0, argList{Yi0, Ybr}},
		{Zbr, 1, argList{Yi1, Ybr}},
	}
	psess.yloop = []ytab{
		{Zloop, 1, argList{Ybr}},
	}
	psess.ycall = []ytab{
		{Zcallindreg, 0, argList{Yml}},
		{Zcallindreg, 2, argList{Yrx, Yrx}},
		{Zcallind, 2, argList{Yindir}},
		{Zcall, 0, argList{Ybr}},
		{Zcallcon, 1, argList{Yi32}},
	}
	psess.yduff = []ytab{
		{Zcallduff, 1, argList{Yi32}},
	}
	psess.yjmp = []ytab{
		{Zo_m64, 2, argList{Yml}},
		{Zjmp, 0, argList{Ybr}},
		{Zjmpcon, 1, argList{Yi32}},
	}
	psess.yfmvd = []ytab{
		{Zm_o, 2, argList{Ym, Yf0}},
		{Zo_m, 2, argList{Yf0, Ym}},
		{Zm_o, 2, argList{Yrf, Yf0}},
		{Zo_m, 2, argList{Yf0, Yrf}},
	}
	psess.yfmvdp = []ytab{
		{Zo_m, 2, argList{Yf0, Ym}},
		{Zo_m, 2, argList{Yf0, Yrf}},
	}
	psess.yfmvf = []ytab{
		{Zm_o, 2, argList{Ym, Yf0}},
		{Zo_m, 2, argList{Yf0, Ym}},
	}
	psess.yfmvx = []ytab{
		{Zm_o, 2, argList{Ym, Yf0}},
	}
	psess.yfmvp = []ytab{
		{Zo_m, 2, argList{Yf0, Ym}},
	}
	psess.yfcmv = []ytab{
		{Zm_o, 2, argList{Yrf, Yf0}},
	}
	psess.yfadd = []ytab{
		{Zm_o, 2, argList{Ym, Yf0}},
		{Zm_o, 2, argList{Yrf, Yf0}},
		{Zo_m, 2, argList{Yf0, Yrf}},
	}
	psess.yfxch = []ytab{
		{Zo_m, 2, argList{Yf0, Yrf}},
		{Zm_o, 2, argList{Yrf, Yf0}},
	}
	psess.ycompp = []ytab{
		{Zo_m, 2, argList{Yf0, Yrf}},
	}
	psess.ystsw = []ytab{
		{Zo_m, 2, argList{Ym}},
		{Zlit, 1, argList{Yax}},
	}
	psess.ysvrs_mo = []ytab{
		{Zm_o, 2, argList{Ym}},
	}
	psess.ysvrs_om = []ytab{
		{Zo_m, 2, argList{Ym}},
	}
	psess.ymm = []ytab{
		{Zm_r_xm, 1, argList{Ymm, Ymr}},
		{Zm_r_xm, 2, argList{Yxm, Yxr}},
	}
	psess.yxm = []ytab{
		{Zm_r_xm, 1, argList{Yxm, Yxr}},
	}
	psess.yxm_q4 = []ytab{
		{Zm_r, 1, argList{Yxm, Yxr}},
	}
	psess.yxcvm1 = []ytab{
		{Zm_r_xm, 2, argList{Yxm, Yxr}},
		{Zm_r_xm, 2, argList{Yxm, Ymr}},
	}
	psess.yxcvm2 = []ytab{
		{Zm_r_xm, 2, argList{Yxm, Yxr}},
		{Zm_r_xm, 2, argList{Ymm, Yxr}},
	}
	psess.yxr = []ytab{
		{Zm_r_xm, 1, argList{Yxr, Yxr}},
	}
	psess.yxr_ml = []ytab{
		{Zr_m_xm, 1, argList{Yxr, Yml}},
	}
	psess.ymr = []ytab{
		{Zm_r, 1, argList{Ymr, Ymr}},
	}
	psess.ymr_ml = []ytab{
		{Zr_m_xm, 1, argList{Ymr, Yml}},
	}
	psess.yxcmpi = []ytab{
		{Zm_r_i_xm, 2, argList{Yxm, Yxr, Yi8}},
	}
	psess.yxmov = []ytab{
		{Zm_r_xm, 1, argList{Yxm, Yxr}},
		{Zr_m_xm, 1, argList{Yxr, Yxm}},
	}
	psess.yxcvfl = []ytab{
		{Zm_r_xm, 1, argList{Yxm, Yrl}},
	}
	psess.yxcvlf = []ytab{
		{Zm_r_xm, 1, argList{Yml, Yxr}},
	}
	psess.yxcvfq = []ytab{
		{Zm_r_xm, 2, argList{Yxm, Yrl}},
	}
	psess.yxcvqf = []ytab{
		{Zm_r_xm, 2, argList{Yml, Yxr}},
	}
	psess.yps = []ytab{
		{Zm_r_xm, 1, argList{Ymm, Ymr}},
		{Zibo_m_xm, 2, argList{Yi8, Ymr}},
		{Zm_r_xm, 2, argList{Yxm, Yxr}},
		{Zibo_m_xm, 3, argList{Yi8, Yxr}},
	}
	psess.yxrrl = []ytab{
		{Zm_r, 1, argList{Yxr, Yrl}},
	}
	psess.ymrxr = []ytab{
		{Zm_r, 1, argList{Ymr, Yxr}},
		{Zm_r_xm, 1, argList{Yxm, Yxr}},
	}
	psess.ymshuf = []ytab{
		{Zibm_r, 2, argList{Yi8, Ymm, Ymr}},
	}
	psess.ymshufb = []ytab{
		{Zm2_r, 2, argList{Yxm, Yxr}},
	}
	psess.yxshuf = []ytab{
		{Zibm_r, 2, argList{Yu8, Yxm, Yxr}},
	}
	psess.yextrw = []ytab{
		{Zibm_r, 2, argList{Yu8, Yxr, Yrl}},
		{Zibr_m, 2, argList{Yu8, Yxr, Yml}},
	}
	psess.yextr = []ytab{
		{Zibr_m, 3, argList{Yu8, Yxr, Ymm}},
	}
	psess.yinsrw = []ytab{
		{Zibm_r, 2, argList{Yu8, Yml, Yxr}},
	}
	psess.yinsr = []ytab{
		{Zibm_r, 3, argList{Yu8, Ymm, Yxr}},
	}
	psess.ypsdq = []ytab{
		{Zibo_m, 2, argList{Yi8, Yxr}},
	}
	psess.ymskb = []ytab{
		{Zm_r_xm, 2, argList{Yxr, Yrl}},
		{Zm_r_xm, 1, argList{Ymr, Yrl}},
	}
	psess.ycrc32l = []ytab{
		{Zlitm_r, 0, argList{Yml, Yrl}},
	}
	psess.ycrc32b = []ytab{
		{Zlitm_r, 0, argList{Ymb, Yrl}},
	}
	psess.yprefetch = []ytab{
		{Zm_o, 2, argList{Ym}},
	}
	psess.yaes = []ytab{
		{Zlitm_r, 2, argList{Yxm, Yxr}},
	}
	psess.yxbegin = []ytab{
		{Zjmp, 1, argList{Ybr}},
	}
	psess.yxabort = []ytab{
		{Zib_, 1, argList{Yu8}},
	}
	psess.ylddqu = []ytab{
		{Zm_r, 1, argList{Ym, Yxr}},
	}
	psess.ypalignr = []ytab{
		{Zibm_r, 2, argList{Yu8, Yxm, Yxr}},
	}
	psess.ysha256rnds2 = []ytab{
		{Zlit_m_r, 0, argList{Yxr0, Yxm, Yxr}},
	}
	psess.yblendvpd = []ytab{
		{Z_m_r, 1, argList{Yxr0, Yxm, Yxr}},
	}
	psess.ymmxmm0f38 = []ytab{
		{Zlitm_r, 3, argList{Ymm, Ymr}},
		{Zlitm_r, 5, argList{Yxm, Yxr}},
	}
	psess.yextractps = []ytab{
		{Zibr_m, 2, argList{Yu2, Yxr, Yml}},
	}
	psess.ysha1rnds4 = []ytab{
		{Zibm_r, 2, argList{Yu2, Yxm, Yxr}},
	}
	psess.nop = [][16]uint8{
		{0x90},
		{0x66, 0x90},
		{0x0F, 0x1F, 0x00},
		{0x0F, 0x1F, 0x40, 0x00},
		{0x0F, 0x1F, 0x44, 0x00, 0x00},
		{0x66, 0x0F, 0x1F, 0x44, 0x00, 0x00},
		{0x0F, 0x1F, 0x80, 0x00, 0x00, 0x00, 0x00},
		{0x0F, 0x1F, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x66, 0x0F, 0x1F, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00},
	}
	psess.isAndroid = (psess.objabi.GOOS == "android")
	psess.bpduff1 = []byte{
		0x48, 0x89, 0x6c, 0x24, 0xf0,
		0x48, 0x8d, 0x6c, 0x24, 0xf0,
	}
	psess.bpduff2 = []byte{
		0x48, 0x8b, 0x6d, 0x00,
	}
	psess.naclret = []uint8{
		0x5e,

		0x83,
		0xe6,
		0xe0,
		0x4c,
		0x01,
		0xfe,
		0xff,
		0xe6,
	}
	psess.naclret8 = []uint8{
		0x5d,

		0x83,
		0xe5,
		0xe0,
		0xff,
		0xe5,
	}
	psess.naclspfix = []uint8{0x4c, 0x01, 0xfc}
	psess.naclbpfix = []uint8{0x4c, 0x01, 0xfd}
	psess.naclmovs = []uint8{
		0x89,
		0xf6,
		0x49,
		0x8d,
		0x34,
		0x37,
		0x89,
		0xff,
		0x49,
		0x8d,
		0x3c,
		0x3f,
	}
	psess.naclstos = []uint8{
		0x89,
		0xff,
		0x49,
		0x8d,
		0x3c,
		0x3f,
	}
	psess.ymovtab = []Movtab{

		{APUSHL, Ycs, Ynone, Ynone, movLit, [4]uint8{0x0e, 0}},
		{APUSHL, Yss, Ynone, Ynone, movLit, [4]uint8{0x16, 0}},
		{APUSHL, Yds, Ynone, Ynone, movLit, [4]uint8{0x1e, 0}},
		{APUSHL, Yes, Ynone, Ynone, movLit, [4]uint8{0x06, 0}},
		{APUSHL, Yfs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa0, 0}},
		{APUSHL, Ygs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa8, 0}},
		{APUSHQ, Yfs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa0, 0}},
		{APUSHQ, Ygs, Ynone, Ynone, movLit, [4]uint8{0x0f, 0xa8, 0}},
		{APUSHW, Ycs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0e, 0}},
		{APUSHW, Yss, Ynone, Ynone, movLit, [4]uint8{Pe, 0x16, 0}},
		{APUSHW, Yds, Ynone, Ynone, movLit, [4]uint8{Pe, 0x1e, 0}},
		{APUSHW, Yes, Ynone, Ynone, movLit, [4]uint8{Pe, 0x06, 0}},
		{APUSHW, Yfs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0f, 0xa0, 0}},
		{APUSHW, Ygs, Ynone, Ynone, movLit, [4]uint8{Pe, 0x0f, 0xa8, 0}},

		{APOPL, Ynone, Ynone, Yds, movLit, [4]uint8{0x1f, 0}},
		{APOPL, Ynone, Ynone, Yes, movLit, [4]uint8{0x07, 0}},
		{APOPL, Ynone, Ynone, Yss, movLit, [4]uint8{0x17, 0}},
		{APOPL, Ynone, Ynone, Yfs, movLit, [4]uint8{0x0f, 0xa1, 0}},
		{APOPL, Ynone, Ynone, Ygs, movLit, [4]uint8{0x0f, 0xa9, 0}},
		{APOPQ, Ynone, Ynone, Yfs, movLit, [4]uint8{0x0f, 0xa1, 0}},
		{APOPQ, Ynone, Ynone, Ygs, movLit, [4]uint8{0x0f, 0xa9, 0}},
		{APOPW, Ynone, Ynone, Yds, movLit, [4]uint8{Pe, 0x1f, 0}},
		{APOPW, Ynone, Ynone, Yes, movLit, [4]uint8{Pe, 0x07, 0}},
		{APOPW, Ynone, Ynone, Yss, movLit, [4]uint8{Pe, 0x17, 0}},
		{APOPW, Ynone, Ynone, Yfs, movLit, [4]uint8{Pe, 0x0f, 0xa1, 0}},
		{APOPW, Ynone, Ynone, Ygs, movLit, [4]uint8{Pe, 0x0f, 0xa9, 0}},

		{AMOVW, Yes, Ynone, Yml, movRegMem, [4]uint8{0x8c, 0, 0, 0}},
		{AMOVW, Ycs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 1, 0, 0}},
		{AMOVW, Yss, Ynone, Yml, movRegMem, [4]uint8{0x8c, 2, 0, 0}},
		{AMOVW, Yds, Ynone, Yml, movRegMem, [4]uint8{0x8c, 3, 0, 0}},
		{AMOVW, Yfs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 4, 0, 0}},
		{AMOVW, Ygs, Ynone, Yml, movRegMem, [4]uint8{0x8c, 5, 0, 0}},
		{AMOVW, Yml, Ynone, Yes, movMemReg, [4]uint8{0x8e, 0, 0, 0}},
		{AMOVW, Yml, Ynone, Ycs, movMemReg, [4]uint8{0x8e, 1, 0, 0}},
		{AMOVW, Yml, Ynone, Yss, movMemReg, [4]uint8{0x8e, 2, 0, 0}},
		{AMOVW, Yml, Ynone, Yds, movMemReg, [4]uint8{0x8e, 3, 0, 0}},
		{AMOVW, Yml, Ynone, Yfs, movMemReg, [4]uint8{0x8e, 4, 0, 0}},
		{AMOVW, Yml, Ynone, Ygs, movMemReg, [4]uint8{0x8e, 5, 0, 0}},

		{AMOVL, Ycr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 0, 0}},
		{AMOVL, Ycr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 2, 0}},
		{AMOVL, Ycr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 3, 0}},
		{AMOVL, Ycr4, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 4, 0}},
		{AMOVL, Ycr8, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 8, 0}},
		{AMOVQ, Ycr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 0, 0}},
		{AMOVQ, Ycr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 2, 0}},
		{AMOVQ, Ycr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 3, 0}},
		{AMOVQ, Ycr4, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 4, 0}},
		{AMOVQ, Ycr8, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x20, 8, 0}},
		{AMOVL, Yrl, Ynone, Ycr0, movMemReg2op, [4]uint8{0x0f, 0x22, 0, 0}},
		{AMOVL, Yrl, Ynone, Ycr2, movMemReg2op, [4]uint8{0x0f, 0x22, 2, 0}},
		{AMOVL, Yrl, Ynone, Ycr3, movMemReg2op, [4]uint8{0x0f, 0x22, 3, 0}},
		{AMOVL, Yrl, Ynone, Ycr4, movMemReg2op, [4]uint8{0x0f, 0x22, 4, 0}},
		{AMOVL, Yrl, Ynone, Ycr8, movMemReg2op, [4]uint8{0x0f, 0x22, 8, 0}},
		{AMOVQ, Yrl, Ynone, Ycr0, movMemReg2op, [4]uint8{0x0f, 0x22, 0, 0}},
		{AMOVQ, Yrl, Ynone, Ycr2, movMemReg2op, [4]uint8{0x0f, 0x22, 2, 0}},
		{AMOVQ, Yrl, Ynone, Ycr3, movMemReg2op, [4]uint8{0x0f, 0x22, 3, 0}},
		{AMOVQ, Yrl, Ynone, Ycr4, movMemReg2op, [4]uint8{0x0f, 0x22, 4, 0}},
		{AMOVQ, Yrl, Ynone, Ycr8, movMemReg2op, [4]uint8{0x0f, 0x22, 8, 0}},

		{AMOVL, Ydr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 0, 0}},
		{AMOVL, Ydr6, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 6, 0}},
		{AMOVL, Ydr7, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 7, 0}},
		{AMOVQ, Ydr0, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 0, 0}},
		{AMOVQ, Ydr2, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 2, 0}},
		{AMOVQ, Ydr3, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 3, 0}},
		{AMOVQ, Ydr6, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 6, 0}},
		{AMOVQ, Ydr7, Ynone, Yrl, movRegMem2op, [4]uint8{0x0f, 0x21, 7, 0}},
		{AMOVL, Yrl, Ynone, Ydr0, movMemReg2op, [4]uint8{0x0f, 0x23, 0, 0}},
		{AMOVL, Yrl, Ynone, Ydr6, movMemReg2op, [4]uint8{0x0f, 0x23, 6, 0}},
		{AMOVL, Yrl, Ynone, Ydr7, movMemReg2op, [4]uint8{0x0f, 0x23, 7, 0}},
		{AMOVQ, Yrl, Ynone, Ydr0, movMemReg2op, [4]uint8{0x0f, 0x23, 0, 0}},
		{AMOVQ, Yrl, Ynone, Ydr2, movMemReg2op, [4]uint8{0x0f, 0x23, 2, 0}},
		{AMOVQ, Yrl, Ynone, Ydr3, movMemReg2op, [4]uint8{0x0f, 0x23, 3, 0}},
		{AMOVQ, Yrl, Ynone, Ydr6, movMemReg2op, [4]uint8{0x0f, 0x23, 6, 0}},
		{AMOVQ, Yrl, Ynone, Ydr7, movMemReg2op, [4]uint8{0x0f, 0x23, 7, 0}},

		{AMOVL, Ytr6, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x24, 6, 0}},
		{AMOVL, Ytr7, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x24, 7, 0}},
		{AMOVL, Yml, Ynone, Ytr6, movMemReg2op, [4]uint8{0x0f, 0x26, 6, 0xff}},
		{AMOVL, Yml, Ynone, Ytr7, movMemReg2op, [4]uint8{0x0f, 0x26, 7, 0xff}},

		{AMOVL, Ym, Ynone, Ygdtr, movMemReg2op, [4]uint8{0x0f, 0x01, 2, 0}},
		{AMOVL, Ygdtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 0, 0}},
		{AMOVL, Ym, Ynone, Yidtr, movMemReg2op, [4]uint8{0x0f, 0x01, 3, 0}},
		{AMOVL, Yidtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 1, 0}},
		{AMOVQ, Ym, Ynone, Ygdtr, movMemReg2op, [4]uint8{0x0f, 0x01, 2, 0}},
		{AMOVQ, Ygdtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 0, 0}},
		{AMOVQ, Ym, Ynone, Yidtr, movMemReg2op, [4]uint8{0x0f, 0x01, 3, 0}},
		{AMOVQ, Yidtr, Ynone, Ym, movRegMem2op, [4]uint8{0x0f, 0x01, 1, 0}},

		{AMOVW, Yml, Ynone, Yldtr, movMemReg2op, [4]uint8{0x0f, 0x00, 2, 0}},
		{AMOVW, Yldtr, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x00, 0, 0}},

		{AMOVW, Yml, Ynone, Ymsw, movMemReg2op, [4]uint8{0x0f, 0x01, 6, 0}},
		{AMOVW, Ymsw, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x01, 4, 0}},

		{AMOVW, Yml, Ynone, Ytask, movMemReg2op, [4]uint8{0x0f, 0x00, 3, 0}},
		{AMOVW, Ytask, Ynone, Yml, movRegMem2op, [4]uint8{0x0f, 0x00, 1, 0}},

		{ASHLL, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
		{ASHLL, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
		{ASHLL, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{0xa4, 0xa5, 0, 0}},
		{ASHRL, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
		{ASHRL, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
		{ASHRL, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{0xac, 0xad, 0, 0}},
		{ASHLQ, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
		{ASHLQ, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
		{ASHLQ, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xa4, 0xa5, 0}},
		{ASHRQ, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
		{ASHRQ, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
		{ASHRQ, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pw, 0xac, 0xad, 0}},
		{ASHLW, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
		{ASHLW, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
		{ASHLW, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xa4, 0xa5, 0}},
		{ASHRW, Yi8, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},
		{ASHRW, Ycl, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},
		{ASHRW, Ycx, Yrl, Yml, movDoubleShift, [4]uint8{Pe, 0xac, 0xad, 0}},

		{AMOVL, Ytls, Ynone, Yrl, movTLSReg, [4]uint8{0, 0, 0, 0}},
		{AMOVQ, Ytls, Ynone, Yrl, movTLSReg, [4]uint8{0, 0, 0, 0}},
		{0, 0, 0, 0, 0, [4]uint8{}},
	}
	psess.unaryDst = map[obj.As]bool{
		ABSWAPL:     true,
		ABSWAPQ:     true,
		ABSWAPW:     true,
		ACLFLUSH:    true,
		ACLFLUSHOPT: true,
		ACMPXCHG16B: true,
		ACMPXCHG8B:  true,
		ADECB:       true,
		ADECL:       true,
		ADECQ:       true,
		ADECW:       true,
		AFBSTP:      true,
		AFFREE:      true,
		AFLDENV:     true,
		AFSAVE:      true,
		AFSTCW:      true,
		AFSTENV:     true,
		AFSTSW:      true,
		AFXSAVE64:   true,
		AFXSAVE:     true,
		AINCB:       true,
		AINCL:       true,
		AINCQ:       true,
		AINCW:       true,
		ANEGB:       true,
		ANEGL:       true,
		ANEGQ:       true,
		ANEGW:       true,
		ANOTB:       true,
		ANOTL:       true,
		ANOTQ:       true,
		ANOTW:       true,
		APOPL:       true,
		APOPQ:       true,
		APOPW:       true,
		ARDFSBASEL:  true,
		ARDFSBASEQ:  true,
		ARDGSBASEL:  true,
		ARDGSBASEQ:  true,
		ARDRANDL:    true,
		ARDRANDQ:    true,
		ARDRANDW:    true,
		ARDSEEDL:    true,
		ARDSEEDQ:    true,
		ARDSEEDW:    true,
		ASETCC:      true,
		ASETCS:      true,
		ASETEQ:      true,
		ASETGE:      true,
		ASETGT:      true,
		ASETHI:      true,
		ASETLE:      true,
		ASETLS:      true,
		ASETLT:      true,
		ASETMI:      true,
		ASETNE:      true,
		ASETOC:      true,
		ASETOS:      true,
		ASETPC:      true,
		ASETPL:      true,
		ASETPS:      true,
		ASGDT:       true,
		ASIDT:       true,
		ASLDTL:      true,
		ASLDTQ:      true,
		ASLDTW:      true,
		ASMSWL:      true,
		ASMSWQ:      true,
		ASMSWW:      true,
		ASTMXCSR:    true,
		ASTRL:       true,
		ASTRQ:       true,
		ASTRW:       true,
		AXSAVE64:    true,
		AXSAVE:      true,
		AXSAVEC64:   true,
		AXSAVEC:     true,
		AXSAVEOPT64: true,
		AXSAVEOPT:   true,
		AXSAVES64:   true,
		AXSAVES:     true,
	}
	psess.optab = [...]Optab{
		{obj.AXXX, nil, 0, opBytes{}},
		{AAAA, psess.ynone, P32, opBytes{0x37}},
		{AAAD, psess.ynone, P32, opBytes{0xd5, 0x0a}},
		{AAAM, psess.ynone, P32, opBytes{0xd4, 0x0a}},
		{AAAS, psess.ynone, P32, opBytes{0x3f}},
		{AADCB, psess.yxorb, Pb, opBytes{0x14, 0x80, 02, 0x10, 0x12}},
		{AADCL, psess.yaddl, Px, opBytes{0x83, 02, 0x15, 0x81, 02, 0x11, 0x13}},
		{AADCQ, psess.yaddl, Pw, opBytes{0x83, 02, 0x15, 0x81, 02, 0x11, 0x13}},
		{AADCW, psess.yaddl, Pe, opBytes{0x83, 02, 0x15, 0x81, 02, 0x11, 0x13}},
		{AADCXL, psess.yml_rl, Pq4, opBytes{0xf6}},
		{AADCXQ, psess.yml_rl, Pq4w, opBytes{0xf6}},
		{AADDB, psess.yxorb, Pb, opBytes{0x04, 0x80, 00, 0x00, 0x02}},
		{AADDL, psess.yaddl, Px, opBytes{0x83, 00, 0x05, 0x81, 00, 0x01, 0x03}},
		{AADDPD, psess.yxm, Pq, opBytes{0x58}},
		{AADDPS, psess.yxm, Pm, opBytes{0x58}},
		{AADDQ, psess.yaddl, Pw, opBytes{0x83, 00, 0x05, 0x81, 00, 0x01, 0x03}},
		{AADDSD, psess.yxm, Pf2, opBytes{0x58}},
		{AADDSS, psess.yxm, Pf3, opBytes{0x58}},
		{AADDSUBPD, psess.yxm, Pq, opBytes{0xd0}},
		{AADDSUBPS, psess.yxm, Pf2, opBytes{0xd0}},
		{AADDW, psess.yaddl, Pe, opBytes{0x83, 00, 0x05, 0x81, 00, 0x01, 0x03}},
		{AADOXL, psess.yml_rl, Pq5, opBytes{0xf6}},
		{AADOXQ, psess.yml_rl, Pq5w, opBytes{0xf6}},
		{AADJSP, nil, 0, opBytes{}},
		{AANDB, psess.yxorb, Pb, opBytes{0x24, 0x80, 04, 0x20, 0x22}},
		{AANDL, psess.yaddl, Px, opBytes{0x83, 04, 0x25, 0x81, 04, 0x21, 0x23}},
		{AANDNPD, psess.yxm, Pq, opBytes{0x55}},
		{AANDNPS, psess.yxm, Pm, opBytes{0x55}},
		{AANDPD, psess.yxm, Pq, opBytes{0x54}},
		{AANDPS, psess.yxm, Pm, opBytes{0x54}},
		{AANDQ, psess.yaddl, Pw, opBytes{0x83, 04, 0x25, 0x81, 04, 0x21, 0x23}},
		{AANDW, psess.yaddl, Pe, opBytes{0x83, 04, 0x25, 0x81, 04, 0x21, 0x23}},
		{AARPL, psess.yrl_ml, P32, opBytes{0x63}},
		{ABOUNDL, psess.yrl_m, P32, opBytes{0x62}},
		{ABOUNDW, psess.yrl_m, Pe, opBytes{0x62}},
		{ABSFL, psess.yml_rl, Pm, opBytes{0xbc}},
		{ABSFQ, psess.yml_rl, Pw, opBytes{0x0f, 0xbc}},
		{ABSFW, psess.yml_rl, Pq, opBytes{0xbc}},
		{ABSRL, psess.yml_rl, Pm, opBytes{0xbd}},
		{ABSRQ, psess.yml_rl, Pw, opBytes{0x0f, 0xbd}},
		{ABSRW, psess.yml_rl, Pq, opBytes{0xbd}},
		{ABSWAPW, psess.ybswap, Pe, opBytes{0x0f, 0xc8}},
		{ABSWAPL, psess.ybswap, Px, opBytes{0x0f, 0xc8}},
		{ABSWAPQ, psess.ybswap, Pw, opBytes{0x0f, 0xc8}},
		{ABTCL, psess.ybtl, Pm, opBytes{0xba, 07, 0xbb}},
		{ABTCQ, psess.ybtl, Pw, opBytes{0x0f, 0xba, 07, 0x0f, 0xbb}},
		{ABTCW, psess.ybtl, Pq, opBytes{0xba, 07, 0xbb}},
		{ABTL, psess.ybtl, Pm, opBytes{0xba, 04, 0xa3}},
		{ABTQ, psess.ybtl, Pw, opBytes{0x0f, 0xba, 04, 0x0f, 0xa3}},
		{ABTRL, psess.ybtl, Pm, opBytes{0xba, 06, 0xb3}},
		{ABTRQ, psess.ybtl, Pw, opBytes{0x0f, 0xba, 06, 0x0f, 0xb3}},
		{ABTRW, psess.ybtl, Pq, opBytes{0xba, 06, 0xb3}},
		{ABTSL, psess.ybtl, Pm, opBytes{0xba, 05, 0xab}},
		{ABTSQ, psess.ybtl, Pw, opBytes{0x0f, 0xba, 05, 0x0f, 0xab}},
		{ABTSW, psess.ybtl, Pq, opBytes{0xba, 05, 0xab}},
		{ABTW, psess.ybtl, Pq, opBytes{0xba, 04, 0xa3}},
		{ABYTE, psess.ybyte, Px, opBytes{1}},
		{obj.ACALL, psess.ycall, Px, opBytes{0xff, 02, 0xff, 0x15, 0xe8}},
		{ACBW, psess.ynone, Pe, opBytes{0x98}},
		{ACDQ, psess.ynone, Px, opBytes{0x99}},
		{ACDQE, psess.ynone, Pw, opBytes{0x98}},
		{ACLAC, psess.ynone, Pm, opBytes{01, 0xca}},
		{ACLC, psess.ynone, Px, opBytes{0xf8}},
		{ACLD, psess.ynone, Px, opBytes{0xfc}},
		{ACLFLUSH, psess.yclflush, Pm, opBytes{0xae, 07}},
		{ACLFLUSHOPT, psess.yclflush, Pq, opBytes{0xae, 07}},
		{ACLI, psess.ynone, Px, opBytes{0xfa}},
		{ACLTS, psess.ynone, Pm, opBytes{0x06}},
		{ACMC, psess.ynone, Px, opBytes{0xf5}},
		{ACMOVLCC, psess.yml_rl, Pm, opBytes{0x43}},
		{ACMOVLCS, psess.yml_rl, Pm, opBytes{0x42}},
		{ACMOVLEQ, psess.yml_rl, Pm, opBytes{0x44}},
		{ACMOVLGE, psess.yml_rl, Pm, opBytes{0x4d}},
		{ACMOVLGT, psess.yml_rl, Pm, opBytes{0x4f}},
		{ACMOVLHI, psess.yml_rl, Pm, opBytes{0x47}},
		{ACMOVLLE, psess.yml_rl, Pm, opBytes{0x4e}},
		{ACMOVLLS, psess.yml_rl, Pm, opBytes{0x46}},
		{ACMOVLLT, psess.yml_rl, Pm, opBytes{0x4c}},
		{ACMOVLMI, psess.yml_rl, Pm, opBytes{0x48}},
		{ACMOVLNE, psess.yml_rl, Pm, opBytes{0x45}},
		{ACMOVLOC, psess.yml_rl, Pm, opBytes{0x41}},
		{ACMOVLOS, psess.yml_rl, Pm, opBytes{0x40}},
		{ACMOVLPC, psess.yml_rl, Pm, opBytes{0x4b}},
		{ACMOVLPL, psess.yml_rl, Pm, opBytes{0x49}},
		{ACMOVLPS, psess.yml_rl, Pm, opBytes{0x4a}},
		{ACMOVQCC, psess.yml_rl, Pw, opBytes{0x0f, 0x43}},
		{ACMOVQCS, psess.yml_rl, Pw, opBytes{0x0f, 0x42}},
		{ACMOVQEQ, psess.yml_rl, Pw, opBytes{0x0f, 0x44}},
		{ACMOVQGE, psess.yml_rl, Pw, opBytes{0x0f, 0x4d}},
		{ACMOVQGT, psess.yml_rl, Pw, opBytes{0x0f, 0x4f}},
		{ACMOVQHI, psess.yml_rl, Pw, opBytes{0x0f, 0x47}},
		{ACMOVQLE, psess.yml_rl, Pw, opBytes{0x0f, 0x4e}},
		{ACMOVQLS, psess.yml_rl, Pw, opBytes{0x0f, 0x46}},
		{ACMOVQLT, psess.yml_rl, Pw, opBytes{0x0f, 0x4c}},
		{ACMOVQMI, psess.yml_rl, Pw, opBytes{0x0f, 0x48}},
		{ACMOVQNE, psess.yml_rl, Pw, opBytes{0x0f, 0x45}},
		{ACMOVQOC, psess.yml_rl, Pw, opBytes{0x0f, 0x41}},
		{ACMOVQOS, psess.yml_rl, Pw, opBytes{0x0f, 0x40}},
		{ACMOVQPC, psess.yml_rl, Pw, opBytes{0x0f, 0x4b}},
		{ACMOVQPL, psess.yml_rl, Pw, opBytes{0x0f, 0x49}},
		{ACMOVQPS, psess.yml_rl, Pw, opBytes{0x0f, 0x4a}},
		{ACMOVWCC, psess.yml_rl, Pq, opBytes{0x43}},
		{ACMOVWCS, psess.yml_rl, Pq, opBytes{0x42}},
		{ACMOVWEQ, psess.yml_rl, Pq, opBytes{0x44}},
		{ACMOVWGE, psess.yml_rl, Pq, opBytes{0x4d}},
		{ACMOVWGT, psess.yml_rl, Pq, opBytes{0x4f}},
		{ACMOVWHI, psess.yml_rl, Pq, opBytes{0x47}},
		{ACMOVWLE, psess.yml_rl, Pq, opBytes{0x4e}},
		{ACMOVWLS, psess.yml_rl, Pq, opBytes{0x46}},
		{ACMOVWLT, psess.yml_rl, Pq, opBytes{0x4c}},
		{ACMOVWMI, psess.yml_rl, Pq, opBytes{0x48}},
		{ACMOVWNE, psess.yml_rl, Pq, opBytes{0x45}},
		{ACMOVWOC, psess.yml_rl, Pq, opBytes{0x41}},
		{ACMOVWOS, psess.yml_rl, Pq, opBytes{0x40}},
		{ACMOVWPC, psess.yml_rl, Pq, opBytes{0x4b}},
		{ACMOVWPL, psess.yml_rl, Pq, opBytes{0x49}},
		{ACMOVWPS, psess.yml_rl, Pq, opBytes{0x4a}},
		{ACMPB, psess.ycmpb, Pb, opBytes{0x3c, 0x80, 07, 0x38, 0x3a}},
		{ACMPL, psess.ycmpl, Px, opBytes{0x83, 07, 0x3d, 0x81, 07, 0x39, 0x3b}},
		{ACMPPD, psess.yxcmpi, Px, opBytes{Pe, 0xc2}},
		{ACMPPS, psess.yxcmpi, Pm, opBytes{0xc2, 0}},
		{ACMPQ, psess.ycmpl, Pw, opBytes{0x83, 07, 0x3d, 0x81, 07, 0x39, 0x3b}},
		{ACMPSB, psess.ynone, Pb, opBytes{0xa6}},
		{ACMPSD, psess.yxcmpi, Px, opBytes{Pf2, 0xc2}},
		{ACMPSL, psess.ynone, Px, opBytes{0xa7}},
		{ACMPSQ, psess.ynone, Pw, opBytes{0xa7}},
		{ACMPSS, psess.yxcmpi, Px, opBytes{Pf3, 0xc2}},
		{ACMPSW, psess.ynone, Pe, opBytes{0xa7}},
		{ACMPW, psess.ycmpl, Pe, opBytes{0x83, 07, 0x3d, 0x81, 07, 0x39, 0x3b}},
		{ACOMISD, psess.yxm, Pe, opBytes{0x2f}},
		{ACOMISS, psess.yxm, Pm, opBytes{0x2f}},
		{ACPUID, psess.ynone, Pm, opBytes{0xa2}},
		{ACVTPL2PD, psess.yxcvm2, Px, opBytes{Pf3, 0xe6, Pe, 0x2a}},
		{ACVTPL2PS, psess.yxcvm2, Pm, opBytes{0x5b, 0, 0x2a, 0}},
		{ACVTPD2PL, psess.yxcvm1, Px, opBytes{Pf2, 0xe6, Pe, 0x2d}},
		{ACVTPD2PS, psess.yxm, Pe, opBytes{0x5a}},
		{ACVTPS2PL, psess.yxcvm1, Px, opBytes{Pe, 0x5b, Pm, 0x2d}},
		{ACVTPS2PD, psess.yxm, Pm, opBytes{0x5a}},
		{ACVTSD2SL, psess.yxcvfl, Pf2, opBytes{0x2d}},
		{ACVTSD2SQ, psess.yxcvfq, Pw, opBytes{Pf2, 0x2d}},
		{ACVTSD2SS, psess.yxm, Pf2, opBytes{0x5a}},
		{ACVTSL2SD, psess.yxcvlf, Pf2, opBytes{0x2a}},
		{ACVTSQ2SD, psess.yxcvqf, Pw, opBytes{Pf2, 0x2a}},
		{ACVTSL2SS, psess.yxcvlf, Pf3, opBytes{0x2a}},
		{ACVTSQ2SS, psess.yxcvqf, Pw, opBytes{Pf3, 0x2a}},
		{ACVTSS2SD, psess.yxm, Pf3, opBytes{0x5a}},
		{ACVTSS2SL, psess.yxcvfl, Pf3, opBytes{0x2d}},
		{ACVTSS2SQ, psess.yxcvfq, Pw, opBytes{Pf3, 0x2d}},
		{ACVTTPD2PL, psess.yxcvm1, Px, opBytes{Pe, 0xe6, Pe, 0x2c}},
		{ACVTTPS2PL, psess.yxcvm1, Px, opBytes{Pf3, 0x5b, Pm, 0x2c}},
		{ACVTTSD2SL, psess.yxcvfl, Pf2, opBytes{0x2c}},
		{ACVTTSD2SQ, psess.yxcvfq, Pw, opBytes{Pf2, 0x2c}},
		{ACVTTSS2SL, psess.yxcvfl, Pf3, opBytes{0x2c}},
		{ACVTTSS2SQ, psess.yxcvfq, Pw, opBytes{Pf3, 0x2c}},
		{ACWD, psess.ynone, Pe, opBytes{0x99}},
		{ACWDE, psess.ynone, Px, opBytes{0x98}},
		{ACQO, psess.ynone, Pw, opBytes{0x99}},
		{ADAA, psess.ynone, P32, opBytes{0x27}},
		{ADAS, psess.ynone, P32, opBytes{0x2f}},
		{ADECB, psess.yscond, Pb, opBytes{0xfe, 01}},
		{ADECL, psess.yincl, Px1, opBytes{0x48, 0xff, 01}},
		{ADECQ, psess.yincq, Pw, opBytes{0xff, 01}},
		{ADECW, psess.yincq, Pe, opBytes{0xff, 01}},
		{ADIVB, psess.ydivb, Pb, opBytes{0xf6, 06}},
		{ADIVL, psess.ydivl, Px, opBytes{0xf7, 06}},
		{ADIVPD, psess.yxm, Pe, opBytes{0x5e}},
		{ADIVPS, psess.yxm, Pm, opBytes{0x5e}},
		{ADIVQ, psess.ydivl, Pw, opBytes{0xf7, 06}},
		{ADIVSD, psess.yxm, Pf2, opBytes{0x5e}},
		{ADIVSS, psess.yxm, Pf3, opBytes{0x5e}},
		{ADIVW, psess.ydivl, Pe, opBytes{0xf7, 06}},
		{ADPPD, psess.yxshuf, Pq, opBytes{0x3a, 0x41, 0}},
		{ADPPS, psess.yxshuf, Pq, opBytes{0x3a, 0x40, 0}},
		{AEMMS, psess.ynone, Pm, opBytes{0x77}},
		{AEXTRACTPS, psess.yextractps, Pq, opBytes{0x3a, 0x17, 0}},
		{AENTER, nil, 0, opBytes{}},
		{AFXRSTOR, psess.ysvrs_mo, Pm, opBytes{0xae, 01, 0xae, 01}},
		{AFXSAVE, psess.ysvrs_om, Pm, opBytes{0xae, 00, 0xae, 00}},
		{AFXRSTOR64, psess.ysvrs_mo, Pw, opBytes{0x0f, 0xae, 01, 0x0f, 0xae, 01}},
		{AFXSAVE64, psess.ysvrs_om, Pw, opBytes{0x0f, 0xae, 00, 0x0f, 0xae, 00}},
		{AHLT, psess.ynone, Px, opBytes{0xf4}},
		{AIDIVB, psess.ydivb, Pb, opBytes{0xf6, 07}},
		{AIDIVL, psess.ydivl, Px, opBytes{0xf7, 07}},
		{AIDIVQ, psess.ydivl, Pw, opBytes{0xf7, 07}},
		{AIDIVW, psess.ydivl, Pe, opBytes{0xf7, 07}},
		{AIMULB, psess.ydivb, Pb, opBytes{0xf6, 05}},
		{AIMULL, psess.yimul, Px, opBytes{0xf7, 05, 0x6b, 0x69, Pm, 0xaf}},
		{AIMULQ, psess.yimul, Pw, opBytes{0xf7, 05, 0x6b, 0x69, Pm, 0xaf}},
		{AIMULW, psess.yimul, Pe, opBytes{0xf7, 05, 0x6b, 0x69, Pm, 0xaf}},
		{AIMUL3W, psess.yimul3, Pe, opBytes{0x6b, 00, 0x69, 00}},
		{AIMUL3L, psess.yimul3, Px, opBytes{0x6b, 00, 0x69, 00}},
		{AIMUL3Q, psess.yimul3, Pw, opBytes{0x6b, 00, 0x69, 00}},
		{AINB, psess.yin, Pb, opBytes{0xe4, 0xec}},
		{AINW, psess.yin, Pe, opBytes{0xe5, 0xed}},
		{AINL, psess.yin, Px, opBytes{0xe5, 0xed}},
		{AINCB, psess.yscond, Pb, opBytes{0xfe, 00}},
		{AINCL, psess.yincl, Px1, opBytes{0x40, 0xff, 00}},
		{AINCQ, psess.yincq, Pw, opBytes{0xff, 00}},
		{AINCW, psess.yincq, Pe, opBytes{0xff, 00}},
		{AINSB, psess.ynone, Pb, opBytes{0x6c}},
		{AINSL, psess.ynone, Px, opBytes{0x6d}},
		{AINSERTPS, psess.yxshuf, Pq, opBytes{0x3a, 0x21, 0}},
		{AINSW, psess.ynone, Pe, opBytes{0x6d}},
		{AICEBP, psess.ynone, Px, opBytes{0xf1}},
		{AINT, psess.yint, Px, opBytes{0xcd}},
		{AINTO, psess.ynone, P32, opBytes{0xce}},
		{AIRETL, psess.ynone, Px, opBytes{0xcf}},
		{AIRETQ, psess.ynone, Pw, opBytes{0xcf}},
		{AIRETW, psess.ynone, Pe, opBytes{0xcf}},
		{AJCC, psess.yjcond, Px, opBytes{0x73, 0x83, 00}},
		{AJCS, psess.yjcond, Px, opBytes{0x72, 0x82}},
		{AJCXZL, psess.yloop, Px, opBytes{0xe3}},
		{AJCXZW, psess.yloop, Px, opBytes{0xe3}},
		{AJCXZQ, psess.yloop, Px, opBytes{0xe3}},
		{AJEQ, psess.yjcond, Px, opBytes{0x74, 0x84}},
		{AJGE, psess.yjcond, Px, opBytes{0x7d, 0x8d}},
		{AJGT, psess.yjcond, Px, opBytes{0x7f, 0x8f}},
		{AJHI, psess.yjcond, Px, opBytes{0x77, 0x87}},
		{AJLE, psess.yjcond, Px, opBytes{0x7e, 0x8e}},
		{AJLS, psess.yjcond, Px, opBytes{0x76, 0x86}},
		{AJLT, psess.yjcond, Px, opBytes{0x7c, 0x8c}},
		{AJMI, psess.yjcond, Px, opBytes{0x78, 0x88}},
		{obj.AJMP, psess.yjmp, Px, opBytes{0xff, 04, 0xeb, 0xe9}},
		{AJNE, psess.yjcond, Px, opBytes{0x75, 0x85}},
		{AJOC, psess.yjcond, Px, opBytes{0x71, 0x81, 00}},
		{AJOS, psess.yjcond, Px, opBytes{0x70, 0x80, 00}},
		{AJPC, psess.yjcond, Px, opBytes{0x7b, 0x8b}},
		{AJPL, psess.yjcond, Px, opBytes{0x79, 0x89}},
		{AJPS, psess.yjcond, Px, opBytes{0x7a, 0x8a}},
		{AHADDPD, psess.yxm, Pq, opBytes{0x7c}},
		{AHADDPS, psess.yxm, Pf2, opBytes{0x7c}},
		{AHSUBPD, psess.yxm, Pq, opBytes{0x7d}},
		{AHSUBPS, psess.yxm, Pf2, opBytes{0x7d}},
		{ALAHF, psess.ynone, Px, opBytes{0x9f}},
		{ALARL, psess.yml_rl, Pm, opBytes{0x02}},
		{ALARQ, psess.yml_rl, Pw, opBytes{0x0f, 0x02}},
		{ALARW, psess.yml_rl, Pq, opBytes{0x02}},
		{ALDDQU, psess.ylddqu, Pf2, opBytes{0xf0}},
		{ALDMXCSR, psess.ysvrs_mo, Pm, opBytes{0xae, 02, 0xae, 02}},
		{ALEAL, psess.ym_rl, Px, opBytes{0x8d}},
		{ALEAQ, psess.ym_rl, Pw, opBytes{0x8d}},
		{ALEAVEL, psess.ynone, P32, opBytes{0xc9}},
		{ALEAVEQ, psess.ynone, Py, opBytes{0xc9}},
		{ALEAVEW, psess.ynone, Pe, opBytes{0xc9}},
		{ALEAW, psess.ym_rl, Pe, opBytes{0x8d}},
		{ALOCK, psess.ynone, Px, opBytes{0xf0}},
		{ALODSB, psess.ynone, Pb, opBytes{0xac}},
		{ALODSL, psess.ynone, Px, opBytes{0xad}},
		{ALODSQ, psess.ynone, Pw, opBytes{0xad}},
		{ALODSW, psess.ynone, Pe, opBytes{0xad}},
		{ALONG, psess.ybyte, Px, opBytes{4}},
		{ALOOP, psess.yloop, Px, opBytes{0xe2}},
		{ALOOPEQ, psess.yloop, Px, opBytes{0xe1}},
		{ALOOPNE, psess.yloop, Px, opBytes{0xe0}},
		{ALTR, psess.ydivl, Pm, opBytes{0x00, 03}},
		{ALZCNTL, psess.yml_rl, Pf3, opBytes{0xbd}},
		{ALZCNTQ, psess.yml_rl, Pfw, opBytes{0xbd}},
		{ALZCNTW, psess.yml_rl, Pef3, opBytes{0xbd}},
		{ALSLL, psess.yml_rl, Pm, opBytes{0x03}},
		{ALSLW, psess.yml_rl, Pq, opBytes{0x03}},
		{ALSLQ, psess.yml_rl, Pw, opBytes{0x0f, 0x03}},
		{AMASKMOVOU, psess.yxr, Pe, opBytes{0xf7}},
		{AMASKMOVQ, psess.ymr, Pm, opBytes{0xf7}},
		{AMAXPD, psess.yxm, Pe, opBytes{0x5f}},
		{AMAXPS, psess.yxm, Pm, opBytes{0x5f}},
		{AMAXSD, psess.yxm, Pf2, opBytes{0x5f}},
		{AMAXSS, psess.yxm, Pf3, opBytes{0x5f}},
		{AMINPD, psess.yxm, Pe, opBytes{0x5d}},
		{AMINPS, psess.yxm, Pm, opBytes{0x5d}},
		{AMINSD, psess.yxm, Pf2, opBytes{0x5d}},
		{AMINSS, psess.yxm, Pf3, opBytes{0x5d}},
		{AMONITOR, psess.ynone, Px, opBytes{0x0f, 0x01, 0xc8, 0}},
		{AMWAIT, psess.ynone, Px, opBytes{0x0f, 0x01, 0xc9, 0}},
		{AMOVAPD, psess.yxmov, Pe, opBytes{0x28, 0x29}},
		{AMOVAPS, psess.yxmov, Pm, opBytes{0x28, 0x29}},
		{AMOVB, psess.ymovb, Pb, opBytes{0x88, 0x8a, 0xb0, 0xc6, 00}},
		{AMOVBLSX, psess.ymb_rl, Pm, opBytes{0xbe}},
		{AMOVBLZX, psess.ymb_rl, Pm, opBytes{0xb6}},
		{AMOVBQSX, psess.ymb_rl, Pw, opBytes{0x0f, 0xbe}},
		{AMOVBQZX, psess.ymb_rl, Pw, opBytes{0x0f, 0xb6}},
		{AMOVBWSX, psess.ymb_rl, Pq, opBytes{0xbe}},
		{AMOVSWW, psess.ymb_rl, Pe, opBytes{0x0f, 0xbf}},
		{AMOVBWZX, psess.ymb_rl, Pq, opBytes{0xb6}},
		{AMOVZWW, psess.ymb_rl, Pe, opBytes{0x0f, 0xb7}},
		{AMOVO, psess.yxmov, Pe, opBytes{0x6f, 0x7f}},
		{AMOVOU, psess.yxmov, Pf3, opBytes{0x6f, 0x7f}},
		{AMOVHLPS, psess.yxr, Pm, opBytes{0x12}},
		{AMOVHPD, psess.yxmov, Pe, opBytes{0x16, 0x17}},
		{AMOVHPS, psess.yxmov, Pm, opBytes{0x16, 0x17}},
		{AMOVL, psess.ymovl, Px, opBytes{0x89, 0x8b, 0xb8, 0xc7, 00, 0x6e, 0x7e, Pe, 0x6e, Pe, 0x7e, 0}},
		{AMOVLHPS, psess.yxr, Pm, opBytes{0x16}},
		{AMOVLPD, psess.yxmov, Pe, opBytes{0x12, 0x13}},
		{AMOVLPS, psess.yxmov, Pm, opBytes{0x12, 0x13}},
		{AMOVLQSX, psess.yml_rl, Pw, opBytes{0x63}},
		{AMOVLQZX, psess.yml_rl, Px, opBytes{0x8b}},
		{AMOVMSKPD, psess.yxrrl, Pq, opBytes{0x50}},
		{AMOVMSKPS, psess.yxrrl, Pm, opBytes{0x50}},
		{AMOVNTO, psess.yxr_ml, Pe, opBytes{0xe7}},
		{AMOVNTDQA, psess.ylddqu, Pq4, opBytes{0x2a}},
		{AMOVNTPD, psess.yxr_ml, Pe, opBytes{0x2b}},
		{AMOVNTPS, psess.yxr_ml, Pm, opBytes{0x2b}},
		{AMOVNTQ, psess.ymr_ml, Pm, opBytes{0xe7}},
		{AMOVQ, psess.ymovq, Pw8, opBytes{0x6f, 0x7f, Pf2, 0xd6, Pf3, 0x7e, Pe, 0xd6, 0x89, 0x8b, 0xc7, 00, 0xb8, 0xc7, 00, 0x6e, 0x7e, Pe, 0x6e, Pe, 0x7e, 0}},
		{AMOVQOZX, psess.ymrxr, Pf3, opBytes{0xd6, 0x7e}},
		{AMOVSB, psess.ynone, Pb, opBytes{0xa4}},
		{AMOVSD, psess.yxmov, Pf2, opBytes{0x10, 0x11}},
		{AMOVSL, psess.ynone, Px, opBytes{0xa5}},
		{AMOVSQ, psess.ynone, Pw, opBytes{0xa5}},
		{AMOVSS, psess.yxmov, Pf3, opBytes{0x10, 0x11}},
		{AMOVSW, psess.ynone, Pe, opBytes{0xa5}},
		{AMOVUPD, psess.yxmov, Pe, opBytes{0x10, 0x11}},
		{AMOVUPS, psess.yxmov, Pm, opBytes{0x10, 0x11}},
		{AMOVW, psess.ymovw, Pe, opBytes{0x89, 0x8b, 0xb8, 0xc7, 00, 0}},
		{AMOVWLSX, psess.yml_rl, Pm, opBytes{0xbf}},
		{AMOVWLZX, psess.yml_rl, Pm, opBytes{0xb7}},
		{AMOVWQSX, psess.yml_rl, Pw, opBytes{0x0f, 0xbf}},
		{AMOVWQZX, psess.yml_rl, Pw, opBytes{0x0f, 0xb7}},
		{AMPSADBW, psess.yxshuf, Pq, opBytes{0x3a, 0x42, 0}},
		{AMULB, psess.ydivb, Pb, opBytes{0xf6, 04}},
		{AMULL, psess.ydivl, Px, opBytes{0xf7, 04}},
		{AMULPD, psess.yxm, Pe, opBytes{0x59}},
		{AMULPS, psess.yxm, Ym, opBytes{0x59}},
		{AMULQ, psess.ydivl, Pw, opBytes{0xf7, 04}},
		{AMULSD, psess.yxm, Pf2, opBytes{0x59}},
		{AMULSS, psess.yxm, Pf3, opBytes{0x59}},
		{AMULW, psess.ydivl, Pe, opBytes{0xf7, 04}},
		{ANEGB, psess.yscond, Pb, opBytes{0xf6, 03}},
		{ANEGL, psess.yscond, Px, opBytes{0xf7, 03}},
		{ANEGQ, psess.yscond, Pw, opBytes{0xf7, 03}},
		{ANEGW, psess.yscond, Pe, opBytes{0xf7, 03}},
		{obj.ANOP, psess.ynop, Px, opBytes{0, 0}},
		{ANOTB, psess.yscond, Pb, opBytes{0xf6, 02}},
		{ANOTL, psess.yscond, Px, opBytes{0xf7, 02}},
		{ANOTQ, psess.yscond, Pw, opBytes{0xf7, 02}},
		{ANOTW, psess.yscond, Pe, opBytes{0xf7, 02}},
		{AORB, psess.yxorb, Pb, opBytes{0x0c, 0x80, 01, 0x08, 0x0a}},
		{AORL, psess.yaddl, Px, opBytes{0x83, 01, 0x0d, 0x81, 01, 0x09, 0x0b}},
		{AORPD, psess.yxm, Pq, opBytes{0x56}},
		{AORPS, psess.yxm, Pm, opBytes{0x56}},
		{AORQ, psess.yaddl, Pw, opBytes{0x83, 01, 0x0d, 0x81, 01, 0x09, 0x0b}},
		{AORW, psess.yaddl, Pe, opBytes{0x83, 01, 0x0d, 0x81, 01, 0x09, 0x0b}},
		{AOUTB, psess.yin, Pb, opBytes{0xe6, 0xee}},
		{AOUTL, psess.yin, Px, opBytes{0xe7, 0xef}},
		{AOUTW, psess.yin, Pe, opBytes{0xe7, 0xef}},
		{AOUTSB, psess.ynone, Pb, opBytes{0x6e}},
		{AOUTSL, psess.ynone, Px, opBytes{0x6f}},
		{AOUTSW, psess.ynone, Pe, opBytes{0x6f}},
		{APABSB, psess.yxm_q4, Pq4, opBytes{0x1c}},
		{APABSD, psess.yxm_q4, Pq4, opBytes{0x1e}},
		{APABSW, psess.yxm_q4, Pq4, opBytes{0x1d}},
		{APACKSSLW, psess.ymm, Py1, opBytes{0x6b, Pe, 0x6b}},
		{APACKSSWB, psess.ymm, Py1, opBytes{0x63, Pe, 0x63}},
		{APACKUSDW, psess.yxm_q4, Pq4, opBytes{0x2b}},
		{APACKUSWB, psess.ymm, Py1, opBytes{0x67, Pe, 0x67}},
		{APADDB, psess.ymm, Py1, opBytes{0xfc, Pe, 0xfc}},
		{APADDL, psess.ymm, Py1, opBytes{0xfe, Pe, 0xfe}},
		{APADDQ, psess.yxm, Pe, opBytes{0xd4}},
		{APADDSB, psess.ymm, Py1, opBytes{0xec, Pe, 0xec}},
		{APADDSW, psess.ymm, Py1, opBytes{0xed, Pe, 0xed}},
		{APADDUSB, psess.ymm, Py1, opBytes{0xdc, Pe, 0xdc}},
		{APADDUSW, psess.ymm, Py1, opBytes{0xdd, Pe, 0xdd}},
		{APADDW, psess.ymm, Py1, opBytes{0xfd, Pe, 0xfd}},
		{APALIGNR, psess.ypalignr, Pq, opBytes{0x3a, 0x0f}},
		{APAND, psess.ymm, Py1, opBytes{0xdb, Pe, 0xdb}},
		{APANDN, psess.ymm, Py1, opBytes{0xdf, Pe, 0xdf}},
		{APAUSE, psess.ynone, Px, opBytes{0xf3, 0x90}},
		{APAVGB, psess.ymm, Py1, opBytes{0xe0, Pe, 0xe0}},
		{APAVGW, psess.ymm, Py1, opBytes{0xe3, Pe, 0xe3}},
		{APBLENDW, psess.yxshuf, Pq, opBytes{0x3a, 0x0e, 0}},
		{APCMPEQB, psess.ymm, Py1, opBytes{0x74, Pe, 0x74}},
		{APCMPEQL, psess.ymm, Py1, opBytes{0x76, Pe, 0x76}},
		{APCMPEQQ, psess.yxm_q4, Pq4, opBytes{0x29}},
		{APCMPEQW, psess.ymm, Py1, opBytes{0x75, Pe, 0x75}},
		{APCMPGTB, psess.ymm, Py1, opBytes{0x64, Pe, 0x64}},
		{APCMPGTL, psess.ymm, Py1, opBytes{0x66, Pe, 0x66}},
		{APCMPGTQ, psess.yxm_q4, Pq4, opBytes{0x37}},
		{APCMPGTW, psess.ymm, Py1, opBytes{0x65, Pe, 0x65}},
		{APCMPISTRI, psess.yxshuf, Pq, opBytes{0x3a, 0x63, 0}},
		{APCMPISTRM, psess.yxshuf, Pq, opBytes{0x3a, 0x62, 0}},
		{APEXTRW, psess.yextrw, Pq, opBytes{0xc5, 0, 0x3a, 0x15, 0}},
		{APEXTRB, psess.yextr, Pq, opBytes{0x3a, 0x14, 00}},
		{APEXTRD, psess.yextr, Pq, opBytes{0x3a, 0x16, 00}},
		{APEXTRQ, psess.yextr, Pq3, opBytes{0x3a, 0x16, 00}},
		{APHADDD, psess.ymmxmm0f38, Px, opBytes{0x0F, 0x38, 0x02, 0, 0x66, 0x0F, 0x38, 0x02, 0}},
		{APHADDSW, psess.yxm_q4, Pq4, opBytes{0x03}},
		{APHADDW, psess.yxm_q4, Pq4, opBytes{0x01}},
		{APHMINPOSUW, psess.yxm_q4, Pq4, opBytes{0x41}},
		{APHSUBD, psess.yxm_q4, Pq4, opBytes{0x06}},
		{APHSUBSW, psess.yxm_q4, Pq4, opBytes{0x07}},
		{APHSUBW, psess.yxm_q4, Pq4, opBytes{0x05}},
		{APINSRW, psess.yinsrw, Pq, opBytes{0xc4, 00}},
		{APINSRB, psess.yinsr, Pq, opBytes{0x3a, 0x20, 00}},
		{APINSRD, psess.yinsr, Pq, opBytes{0x3a, 0x22, 00}},
		{APINSRQ, psess.yinsr, Pq3, opBytes{0x3a, 0x22, 00}},
		{APMADDUBSW, psess.yxm_q4, Pq4, opBytes{0x04}},
		{APMADDWL, psess.ymm, Py1, opBytes{0xf5, Pe, 0xf5}},
		{APMAXSB, psess.yxm_q4, Pq4, opBytes{0x3c}},
		{APMAXSD, psess.yxm_q4, Pq4, opBytes{0x3d}},
		{APMAXSW, psess.yxm, Pe, opBytes{0xee}},
		{APMAXUB, psess.yxm, Pe, opBytes{0xde}},
		{APMAXUD, psess.yxm_q4, Pq4, opBytes{0x3f}},
		{APMAXUW, psess.yxm_q4, Pq4, opBytes{0x3e}},
		{APMINSB, psess.yxm_q4, Pq4, opBytes{0x38}},
		{APMINSD, psess.yxm_q4, Pq4, opBytes{0x39}},
		{APMINSW, psess.yxm, Pe, opBytes{0xea}},
		{APMINUB, psess.yxm, Pe, opBytes{0xda}},
		{APMINUD, psess.yxm_q4, Pq4, opBytes{0x3b}},
		{APMINUW, psess.yxm_q4, Pq4, opBytes{0x3a}},
		{APMOVMSKB, psess.ymskb, Px, opBytes{Pe, 0xd7, 0xd7}},
		{APMOVSXBD, psess.yxm_q4, Pq4, opBytes{0x21}},
		{APMOVSXBQ, psess.yxm_q4, Pq4, opBytes{0x22}},
		{APMOVSXBW, psess.yxm_q4, Pq4, opBytes{0x20}},
		{APMOVSXDQ, psess.yxm_q4, Pq4, opBytes{0x25}},
		{APMOVSXWD, psess.yxm_q4, Pq4, opBytes{0x23}},
		{APMOVSXWQ, psess.yxm_q4, Pq4, opBytes{0x24}},
		{APMOVZXBD, psess.yxm_q4, Pq4, opBytes{0x31}},
		{APMOVZXBQ, psess.yxm_q4, Pq4, opBytes{0x32}},
		{APMOVZXBW, psess.yxm_q4, Pq4, opBytes{0x30}},
		{APMOVZXDQ, psess.yxm_q4, Pq4, opBytes{0x35}},
		{APMOVZXWD, psess.yxm_q4, Pq4, opBytes{0x33}},
		{APMOVZXWQ, psess.yxm_q4, Pq4, opBytes{0x34}},
		{APMULDQ, psess.yxm_q4, Pq4, opBytes{0x28}},
		{APMULHRSW, psess.yxm_q4, Pq4, opBytes{0x0b}},
		{APMULHUW, psess.ymm, Py1, opBytes{0xe4, Pe, 0xe4}},
		{APMULHW, psess.ymm, Py1, opBytes{0xe5, Pe, 0xe5}},
		{APMULLD, psess.yxm_q4, Pq4, opBytes{0x40}},
		{APMULLW, psess.ymm, Py1, opBytes{0xd5, Pe, 0xd5}},
		{APMULULQ, psess.ymm, Py1, opBytes{0xf4, Pe, 0xf4}},
		{APOPAL, psess.ynone, P32, opBytes{0x61}},
		{APOPAW, psess.ynone, Pe, opBytes{0x61}},
		{APOPCNTW, psess.yml_rl, Pef3, opBytes{0xb8}},
		{APOPCNTL, psess.yml_rl, Pf3, opBytes{0xb8}},
		{APOPCNTQ, psess.yml_rl, Pfw, opBytes{0xb8}},
		{APOPFL, psess.ynone, P32, opBytes{0x9d}},
		{APOPFQ, psess.ynone, Py, opBytes{0x9d}},
		{APOPFW, psess.ynone, Pe, opBytes{0x9d}},
		{APOPL, psess.ypopl, P32, opBytes{0x58, 0x8f, 00}},
		{APOPQ, psess.ypopl, Py, opBytes{0x58, 0x8f, 00}},
		{APOPW, psess.ypopl, Pe, opBytes{0x58, 0x8f, 00}},
		{APOR, psess.ymm, Py1, opBytes{0xeb, Pe, 0xeb}},
		{APSADBW, psess.yxm, Pq, opBytes{0xf6}},
		{APSHUFHW, psess.yxshuf, Pf3, opBytes{0x70, 00}},
		{APSHUFL, psess.yxshuf, Pq, opBytes{0x70, 00}},
		{APSHUFLW, psess.yxshuf, Pf2, opBytes{0x70, 00}},
		{APSHUFW, psess.ymshuf, Pm, opBytes{0x70, 00}},
		{APSHUFB, psess.ymshufb, Pq, opBytes{0x38, 0x00}},
		{APSIGNB, psess.yxm_q4, Pq4, opBytes{0x08}},
		{APSIGND, psess.yxm_q4, Pq4, opBytes{0x0a}},
		{APSIGNW, psess.yxm_q4, Pq4, opBytes{0x09}},
		{APSLLO, psess.ypsdq, Pq, opBytes{0x73, 07}},
		{APSLLL, psess.yps, Py3, opBytes{0xf2, 0x72, 06, Pe, 0xf2, Pe, 0x72, 06}},
		{APSLLQ, psess.yps, Py3, opBytes{0xf3, 0x73, 06, Pe, 0xf3, Pe, 0x73, 06}},
		{APSLLW, psess.yps, Py3, opBytes{0xf1, 0x71, 06, Pe, 0xf1, Pe, 0x71, 06}},
		{APSRAL, psess.yps, Py3, opBytes{0xe2, 0x72, 04, Pe, 0xe2, Pe, 0x72, 04}},
		{APSRAW, psess.yps, Py3, opBytes{0xe1, 0x71, 04, Pe, 0xe1, Pe, 0x71, 04}},
		{APSRLO, psess.ypsdq, Pq, opBytes{0x73, 03}},
		{APSRLL, psess.yps, Py3, opBytes{0xd2, 0x72, 02, Pe, 0xd2, Pe, 0x72, 02}},
		{APSRLQ, psess.yps, Py3, opBytes{0xd3, 0x73, 02, Pe, 0xd3, Pe, 0x73, 02}},
		{APSRLW, psess.yps, Py3, opBytes{0xd1, 0x71, 02, Pe, 0xd1, Pe, 0x71, 02}},
		{APSUBB, psess.yxm, Pe, opBytes{0xf8}},
		{APSUBL, psess.yxm, Pe, opBytes{0xfa}},
		{APSUBQ, psess.yxm, Pe, opBytes{0xfb}},
		{APSUBSB, psess.yxm, Pe, opBytes{0xe8}},
		{APSUBSW, psess.yxm, Pe, opBytes{0xe9}},
		{APSUBUSB, psess.yxm, Pe, opBytes{0xd8}},
		{APSUBUSW, psess.yxm, Pe, opBytes{0xd9}},
		{APSUBW, psess.yxm, Pe, opBytes{0xf9}},
		{APTEST, psess.yxm_q4, Pq4, opBytes{0x17}},
		{APUNPCKHBW, psess.ymm, Py1, opBytes{0x68, Pe, 0x68}},
		{APUNPCKHLQ, psess.ymm, Py1, opBytes{0x6a, Pe, 0x6a}},
		{APUNPCKHQDQ, psess.yxm, Pe, opBytes{0x6d}},
		{APUNPCKHWL, psess.ymm, Py1, opBytes{0x69, Pe, 0x69}},
		{APUNPCKLBW, psess.ymm, Py1, opBytes{0x60, Pe, 0x60}},
		{APUNPCKLLQ, psess.ymm, Py1, opBytes{0x62, Pe, 0x62}},
		{APUNPCKLQDQ, psess.yxm, Pe, opBytes{0x6c}},
		{APUNPCKLWL, psess.ymm, Py1, opBytes{0x61, Pe, 0x61}},
		{APUSHAL, psess.ynone, P32, opBytes{0x60}},
		{APUSHAW, psess.ynone, Pe, opBytes{0x60}},
		{APUSHFL, psess.ynone, P32, opBytes{0x9c}},
		{APUSHFQ, psess.ynone, Py, opBytes{0x9c}},
		{APUSHFW, psess.ynone, Pe, opBytes{0x9c}},
		{APUSHL, psess.ypushl, P32, opBytes{0x50, 0xff, 06, 0x6a, 0x68}},
		{APUSHQ, psess.ypushl, Py, opBytes{0x50, 0xff, 06, 0x6a, 0x68}},
		{APUSHW, psess.ypushl, Pe, opBytes{0x50, 0xff, 06, 0x6a, 0x68}},
		{APXOR, psess.ymm, Py1, opBytes{0xef, Pe, 0xef}},
		{AQUAD, psess.ybyte, Px, opBytes{8}},
		{ARCLB, psess.yshb, Pb, opBytes{0xd0, 02, 0xc0, 02, 0xd2, 02}},
		{ARCLL, psess.yshl, Px, opBytes{0xd1, 02, 0xc1, 02, 0xd3, 02, 0xd3, 02}},
		{ARCLQ, psess.yshl, Pw, opBytes{0xd1, 02, 0xc1, 02, 0xd3, 02, 0xd3, 02}},
		{ARCLW, psess.yshl, Pe, opBytes{0xd1, 02, 0xc1, 02, 0xd3, 02, 0xd3, 02}},
		{ARCPPS, psess.yxm, Pm, opBytes{0x53}},
		{ARCPSS, psess.yxm, Pf3, opBytes{0x53}},
		{ARCRB, psess.yshb, Pb, opBytes{0xd0, 03, 0xc0, 03, 0xd2, 03}},
		{ARCRL, psess.yshl, Px, opBytes{0xd1, 03, 0xc1, 03, 0xd3, 03, 0xd3, 03}},
		{ARCRQ, psess.yshl, Pw, opBytes{0xd1, 03, 0xc1, 03, 0xd3, 03, 0xd3, 03}},
		{ARCRW, psess.yshl, Pe, opBytes{0xd1, 03, 0xc1, 03, 0xd3, 03, 0xd3, 03}},
		{AREP, psess.ynone, Px, opBytes{0xf3}},
		{AREPN, psess.ynone, Px, opBytes{0xf2}},
		{obj.ARET, psess.ynone, Px, opBytes{0xc3}},
		{ARETFW, psess.yret, Pe, opBytes{0xcb, 0xca}},
		{ARETFL, psess.yret, Px, opBytes{0xcb, 0xca}},
		{ARETFQ, psess.yret, Pw, opBytes{0xcb, 0xca}},
		{AROLB, psess.yshb, Pb, opBytes{0xd0, 00, 0xc0, 00, 0xd2, 00}},
		{AROLL, psess.yshl, Px, opBytes{0xd1, 00, 0xc1, 00, 0xd3, 00, 0xd3, 00}},
		{AROLQ, psess.yshl, Pw, opBytes{0xd1, 00, 0xc1, 00, 0xd3, 00, 0xd3, 00}},
		{AROLW, psess.yshl, Pe, opBytes{0xd1, 00, 0xc1, 00, 0xd3, 00, 0xd3, 00}},
		{ARORB, psess.yshb, Pb, opBytes{0xd0, 01, 0xc0, 01, 0xd2, 01}},
		{ARORL, psess.yshl, Px, opBytes{0xd1, 01, 0xc1, 01, 0xd3, 01, 0xd3, 01}},
		{ARORQ, psess.yshl, Pw, opBytes{0xd1, 01, 0xc1, 01, 0xd3, 01, 0xd3, 01}},
		{ARORW, psess.yshl, Pe, opBytes{0xd1, 01, 0xc1, 01, 0xd3, 01, 0xd3, 01}},
		{ARSQRTPS, psess.yxm, Pm, opBytes{0x52}},
		{ARSQRTSS, psess.yxm, Pf3, opBytes{0x52}},
		{ASAHF, psess.ynone, Px, opBytes{0x9e, 00, 0x86, 0xe0, 0x50, 0x9d}},
		{ASALB, psess.yshb, Pb, opBytes{0xd0, 04, 0xc0, 04, 0xd2, 04}},
		{ASALL, psess.yshl, Px, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASALQ, psess.yshl, Pw, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASALW, psess.yshl, Pe, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASARB, psess.yshb, Pb, opBytes{0xd0, 07, 0xc0, 07, 0xd2, 07}},
		{ASARL, psess.yshl, Px, opBytes{0xd1, 07, 0xc1, 07, 0xd3, 07, 0xd3, 07}},
		{ASARQ, psess.yshl, Pw, opBytes{0xd1, 07, 0xc1, 07, 0xd3, 07, 0xd3, 07}},
		{ASARW, psess.yshl, Pe, opBytes{0xd1, 07, 0xc1, 07, 0xd3, 07, 0xd3, 07}},
		{ASBBB, psess.yxorb, Pb, opBytes{0x1c, 0x80, 03, 0x18, 0x1a}},
		{ASBBL, psess.yaddl, Px, opBytes{0x83, 03, 0x1d, 0x81, 03, 0x19, 0x1b}},
		{ASBBQ, psess.yaddl, Pw, opBytes{0x83, 03, 0x1d, 0x81, 03, 0x19, 0x1b}},
		{ASBBW, psess.yaddl, Pe, opBytes{0x83, 03, 0x1d, 0x81, 03, 0x19, 0x1b}},
		{ASCASB, psess.ynone, Pb, opBytes{0xae}},
		{ASCASL, psess.ynone, Px, opBytes{0xaf}},
		{ASCASQ, psess.ynone, Pw, opBytes{0xaf}},
		{ASCASW, psess.ynone, Pe, opBytes{0xaf}},
		{ASETCC, psess.yscond, Pb, opBytes{0x0f, 0x93, 00}},
		{ASETCS, psess.yscond, Pb, opBytes{0x0f, 0x92, 00}},
		{ASETEQ, psess.yscond, Pb, opBytes{0x0f, 0x94, 00}},
		{ASETGE, psess.yscond, Pb, opBytes{0x0f, 0x9d, 00}},
		{ASETGT, psess.yscond, Pb, opBytes{0x0f, 0x9f, 00}},
		{ASETHI, psess.yscond, Pb, opBytes{0x0f, 0x97, 00}},
		{ASETLE, psess.yscond, Pb, opBytes{0x0f, 0x9e, 00}},
		{ASETLS, psess.yscond, Pb, opBytes{0x0f, 0x96, 00}},
		{ASETLT, psess.yscond, Pb, opBytes{0x0f, 0x9c, 00}},
		{ASETMI, psess.yscond, Pb, opBytes{0x0f, 0x98, 00}},
		{ASETNE, psess.yscond, Pb, opBytes{0x0f, 0x95, 00}},
		{ASETOC, psess.yscond, Pb, opBytes{0x0f, 0x91, 00}},
		{ASETOS, psess.yscond, Pb, opBytes{0x0f, 0x90, 00}},
		{ASETPC, psess.yscond, Pb, opBytes{0x0f, 0x9b, 00}},
		{ASETPL, psess.yscond, Pb, opBytes{0x0f, 0x99, 00}},
		{ASETPS, psess.yscond, Pb, opBytes{0x0f, 0x9a, 00}},
		{ASHLB, psess.yshb, Pb, opBytes{0xd0, 04, 0xc0, 04, 0xd2, 04}},
		{ASHLL, psess.yshl, Px, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASHLQ, psess.yshl, Pw, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASHLW, psess.yshl, Pe, opBytes{0xd1, 04, 0xc1, 04, 0xd3, 04, 0xd3, 04}},
		{ASHRB, psess.yshb, Pb, opBytes{0xd0, 05, 0xc0, 05, 0xd2, 05}},
		{ASHRL, psess.yshl, Px, opBytes{0xd1, 05, 0xc1, 05, 0xd3, 05, 0xd3, 05}},
		{ASHRQ, psess.yshl, Pw, opBytes{0xd1, 05, 0xc1, 05, 0xd3, 05, 0xd3, 05}},
		{ASHRW, psess.yshl, Pe, opBytes{0xd1, 05, 0xc1, 05, 0xd3, 05, 0xd3, 05}},
		{ASHUFPD, psess.yxshuf, Pq, opBytes{0xc6, 00}},
		{ASHUFPS, psess.yxshuf, Pm, opBytes{0xc6, 00}},
		{ASQRTPD, psess.yxm, Pe, opBytes{0x51}},
		{ASQRTPS, psess.yxm, Pm, opBytes{0x51}},
		{ASQRTSD, psess.yxm, Pf2, opBytes{0x51}},
		{ASQRTSS, psess.yxm, Pf3, opBytes{0x51}},
		{ASTC, psess.ynone, Px, opBytes{0xf9}},
		{ASTD, psess.ynone, Px, opBytes{0xfd}},
		{ASTI, psess.ynone, Px, opBytes{0xfb}},
		{ASTMXCSR, psess.ysvrs_om, Pm, opBytes{0xae, 03, 0xae, 03}},
		{ASTOSB, psess.ynone, Pb, opBytes{0xaa}},
		{ASTOSL, psess.ynone, Px, opBytes{0xab}},
		{ASTOSQ, psess.ynone, Pw, opBytes{0xab}},
		{ASTOSW, psess.ynone, Pe, opBytes{0xab}},
		{ASUBB, psess.yxorb, Pb, opBytes{0x2c, 0x80, 05, 0x28, 0x2a}},
		{ASUBL, psess.yaddl, Px, opBytes{0x83, 05, 0x2d, 0x81, 05, 0x29, 0x2b}},
		{ASUBPD, psess.yxm, Pe, opBytes{0x5c}},
		{ASUBPS, psess.yxm, Pm, opBytes{0x5c}},
		{ASUBQ, psess.yaddl, Pw, opBytes{0x83, 05, 0x2d, 0x81, 05, 0x29, 0x2b}},
		{ASUBSD, psess.yxm, Pf2, opBytes{0x5c}},
		{ASUBSS, psess.yxm, Pf3, opBytes{0x5c}},
		{ASUBW, psess.yaddl, Pe, opBytes{0x83, 05, 0x2d, 0x81, 05, 0x29, 0x2b}},
		{ASWAPGS, psess.ynone, Pm, opBytes{0x01, 0xf8}},
		{ASYSCALL, psess.ynone, Px, opBytes{0x0f, 0x05}},
		{ATESTB, psess.yxorb, Pb, opBytes{0xa8, 0xf6, 00, 0x84, 0x84}},
		{ATESTL, psess.ytestl, Px, opBytes{0xa9, 0xf7, 00, 0x85, 0x85}},
		{ATESTQ, psess.ytestl, Pw, opBytes{0xa9, 0xf7, 00, 0x85, 0x85}},
		{ATESTW, psess.ytestl, Pe, opBytes{0xa9, 0xf7, 00, 0x85, 0x85}},
		{obj.ATEXT, psess.ytext, Px, opBytes{}},
		{AUCOMISD, psess.yxm, Pe, opBytes{0x2e}},
		{AUCOMISS, psess.yxm, Pm, opBytes{0x2e}},
		{AUNPCKHPD, psess.yxm, Pe, opBytes{0x15}},
		{AUNPCKHPS, psess.yxm, Pm, opBytes{0x15}},
		{AUNPCKLPD, psess.yxm, Pe, opBytes{0x14}},
		{AUNPCKLPS, psess.yxm, Pm, opBytes{0x14}},
		{AVERR, psess.ydivl, Pm, opBytes{0x00, 04}},
		{AVERW, psess.ydivl, Pm, opBytes{0x00, 05}},
		{AWAIT, psess.ynone, Px, opBytes{0x9b}},
		{AWORD, psess.ybyte, Px, opBytes{2}},
		{AXCHGB, psess.yml_mb, Pb, opBytes{0x86, 0x86}},
		{AXCHGL, psess.yxchg, Px, opBytes{0x90, 0x90, 0x87, 0x87}},
		{AXCHGQ, psess.yxchg, Pw, opBytes{0x90, 0x90, 0x87, 0x87}},
		{AXCHGW, psess.yxchg, Pe, opBytes{0x90, 0x90, 0x87, 0x87}},
		{AXLAT, psess.ynone, Px, opBytes{0xd7}},
		{AXORB, psess.yxorb, Pb, opBytes{0x34, 0x80, 06, 0x30, 0x32}},
		{AXORL, psess.yaddl, Px, opBytes{0x83, 06, 0x35, 0x81, 06, 0x31, 0x33}},
		{AXORPD, psess.yxm, Pe, opBytes{0x57}},
		{AXORPS, psess.yxm, Pm, opBytes{0x57}},
		{AXORQ, psess.yaddl, Pw, opBytes{0x83, 06, 0x35, 0x81, 06, 0x31, 0x33}},
		{AXORW, psess.yaddl, Pe, opBytes{0x83, 06, 0x35, 0x81, 06, 0x31, 0x33}},
		{AFMOVB, psess.yfmvx, Px, opBytes{0xdf, 04}},
		{AFMOVBP, psess.yfmvp, Px, opBytes{0xdf, 06}},
		{AFMOVD, psess.yfmvd, Px, opBytes{0xdd, 00, 0xdd, 02, 0xd9, 00, 0xdd, 02}},
		{AFMOVDP, psess.yfmvdp, Px, opBytes{0xdd, 03, 0xdd, 03}},
		{AFMOVF, psess.yfmvf, Px, opBytes{0xd9, 00, 0xd9, 02}},
		{AFMOVFP, psess.yfmvp, Px, opBytes{0xd9, 03}},
		{AFMOVL, psess.yfmvf, Px, opBytes{0xdb, 00, 0xdb, 02}},
		{AFMOVLP, psess.yfmvp, Px, opBytes{0xdb, 03}},
		{AFMOVV, psess.yfmvx, Px, opBytes{0xdf, 05}},
		{AFMOVVP, psess.yfmvp, Px, opBytes{0xdf, 07}},
		{AFMOVW, psess.yfmvf, Px, opBytes{0xdf, 00, 0xdf, 02}},
		{AFMOVWP, psess.yfmvp, Px, opBytes{0xdf, 03}},
		{AFMOVX, psess.yfmvx, Px, opBytes{0xdb, 05}},
		{AFMOVXP, psess.yfmvp, Px, opBytes{0xdb, 07}},
		{AFCMOVCC, psess.yfcmv, Px, opBytes{0xdb, 00}},
		{AFCMOVCS, psess.yfcmv, Px, opBytes{0xda, 00}},
		{AFCMOVEQ, psess.yfcmv, Px, opBytes{0xda, 01}},
		{AFCMOVHI, psess.yfcmv, Px, opBytes{0xdb, 02}},
		{AFCMOVLS, psess.yfcmv, Px, opBytes{0xda, 02}},
		{AFCMOVB, psess.yfcmv, Px, opBytes{0xda, 00}},
		{AFCMOVBE, psess.yfcmv, Px, opBytes{0xda, 02}},
		{AFCMOVNB, psess.yfcmv, Px, opBytes{0xdb, 00}},
		{AFCMOVNBE, psess.yfcmv, Px, opBytes{0xdb, 02}},
		{AFCMOVE, psess.yfcmv, Px, opBytes{0xda, 01}},
		{AFCMOVNE, psess.yfcmv, Px, opBytes{0xdb, 01}},
		{AFCMOVNU, psess.yfcmv, Px, opBytes{0xdb, 03}},
		{AFCMOVU, psess.yfcmv, Px, opBytes{0xda, 03}},
		{AFCMOVUN, psess.yfcmv, Px, opBytes{0xda, 03}},
		{AFCOMD, psess.yfadd, Px, opBytes{0xdc, 02, 0xd8, 02, 0xdc, 02}},
		{AFCOMDP, psess.yfadd, Px, opBytes{0xdc, 03, 0xd8, 03, 0xdc, 03}},
		{AFCOMDPP, psess.ycompp, Px, opBytes{0xde, 03}},
		{AFCOMF, psess.yfmvx, Px, opBytes{0xd8, 02}},
		{AFCOMFP, psess.yfmvx, Px, opBytes{0xd8, 03}},
		{AFCOMI, psess.yfcmv, Px, opBytes{0xdb, 06}},
		{AFCOMIP, psess.yfcmv, Px, opBytes{0xdf, 06}},
		{AFCOML, psess.yfmvx, Px, opBytes{0xda, 02}},
		{AFCOMLP, psess.yfmvx, Px, opBytes{0xda, 03}},
		{AFCOMW, psess.yfmvx, Px, opBytes{0xde, 02}},
		{AFCOMWP, psess.yfmvx, Px, opBytes{0xde, 03}},
		{AFUCOM, psess.ycompp, Px, opBytes{0xdd, 04}},
		{AFUCOMI, psess.ycompp, Px, opBytes{0xdb, 05}},
		{AFUCOMIP, psess.ycompp, Px, opBytes{0xdf, 05}},
		{AFUCOMP, psess.ycompp, Px, opBytes{0xdd, 05}},
		{AFUCOMPP, psess.ycompp, Px, opBytes{0xda, 13}},
		{AFADDDP, psess.ycompp, Px, opBytes{0xde, 00}},
		{AFADDW, psess.yfmvx, Px, opBytes{0xde, 00}},
		{AFADDL, psess.yfmvx, Px, opBytes{0xda, 00}},
		{AFADDF, psess.yfmvx, Px, opBytes{0xd8, 00}},
		{AFADDD, psess.yfadd, Px, opBytes{0xdc, 00, 0xd8, 00, 0xdc, 00}},
		{AFMULDP, psess.ycompp, Px, opBytes{0xde, 01}},
		{AFMULW, psess.yfmvx, Px, opBytes{0xde, 01}},
		{AFMULL, psess.yfmvx, Px, opBytes{0xda, 01}},
		{AFMULF, psess.yfmvx, Px, opBytes{0xd8, 01}},
		{AFMULD, psess.yfadd, Px, opBytes{0xdc, 01, 0xd8, 01, 0xdc, 01}},
		{AFSUBDP, psess.ycompp, Px, opBytes{0xde, 05}},
		{AFSUBW, psess.yfmvx, Px, opBytes{0xde, 04}},
		{AFSUBL, psess.yfmvx, Px, opBytes{0xda, 04}},
		{AFSUBF, psess.yfmvx, Px, opBytes{0xd8, 04}},
		{AFSUBD, psess.yfadd, Px, opBytes{0xdc, 04, 0xd8, 04, 0xdc, 05}},
		{AFSUBRDP, psess.ycompp, Px, opBytes{0xde, 04}},
		{AFSUBRW, psess.yfmvx, Px, opBytes{0xde, 05}},
		{AFSUBRL, psess.yfmvx, Px, opBytes{0xda, 05}},
		{AFSUBRF, psess.yfmvx, Px, opBytes{0xd8, 05}},
		{AFSUBRD, psess.yfadd, Px, opBytes{0xdc, 05, 0xd8, 05, 0xdc, 04}},
		{AFDIVDP, psess.ycompp, Px, opBytes{0xde, 07}},
		{AFDIVW, psess.yfmvx, Px, opBytes{0xde, 06}},
		{AFDIVL, psess.yfmvx, Px, opBytes{0xda, 06}},
		{AFDIVF, psess.yfmvx, Px, opBytes{0xd8, 06}},
		{AFDIVD, psess.yfadd, Px, opBytes{0xdc, 06, 0xd8, 06, 0xdc, 07}},
		{AFDIVRDP, psess.ycompp, Px, opBytes{0xde, 06}},
		{AFDIVRW, psess.yfmvx, Px, opBytes{0xde, 07}},
		{AFDIVRL, psess.yfmvx, Px, opBytes{0xda, 07}},
		{AFDIVRF, psess.yfmvx, Px, opBytes{0xd8, 07}},
		{AFDIVRD, psess.yfadd, Px, opBytes{0xdc, 07, 0xd8, 07, 0xdc, 06}},
		{AFXCHD, psess.yfxch, Px, opBytes{0xd9, 01, 0xd9, 01}},
		{AFFREE, nil, 0, opBytes{}},
		{AFLDCW, psess.ysvrs_mo, Px, opBytes{0xd9, 05, 0xd9, 05}},
		{AFLDENV, psess.ysvrs_mo, Px, opBytes{0xd9, 04, 0xd9, 04}},
		{AFRSTOR, psess.ysvrs_mo, Px, opBytes{0xdd, 04, 0xdd, 04}},
		{AFSAVE, psess.ysvrs_om, Px, opBytes{0xdd, 06, 0xdd, 06}},
		{AFSTCW, psess.ysvrs_om, Px, opBytes{0xd9, 07, 0xd9, 07}},
		{AFSTENV, psess.ysvrs_om, Px, opBytes{0xd9, 06, 0xd9, 06}},
		{AFSTSW, psess.ystsw, Px, opBytes{0xdd, 07, 0xdf, 0xe0}},
		{AF2XM1, psess.ynone, Px, opBytes{0xd9, 0xf0}},
		{AFABS, psess.ynone, Px, opBytes{0xd9, 0xe1}},
		{AFBLD, psess.ysvrs_mo, Px, opBytes{0xdf, 04}},
		{AFBSTP, psess.yclflush, Px, opBytes{0xdf, 06}},
		{AFCHS, psess.ynone, Px, opBytes{0xd9, 0xe0}},
		{AFCLEX, psess.ynone, Px, opBytes{0xdb, 0xe2}},
		{AFCOS, psess.ynone, Px, opBytes{0xd9, 0xff}},
		{AFDECSTP, psess.ynone, Px, opBytes{0xd9, 0xf6}},
		{AFINCSTP, psess.ynone, Px, opBytes{0xd9, 0xf7}},
		{AFINIT, psess.ynone, Px, opBytes{0xdb, 0xe3}},
		{AFLD1, psess.ynone, Px, opBytes{0xd9, 0xe8}},
		{AFLDL2E, psess.ynone, Px, opBytes{0xd9, 0xea}},
		{AFLDL2T, psess.ynone, Px, opBytes{0xd9, 0xe9}},
		{AFLDLG2, psess.ynone, Px, opBytes{0xd9, 0xec}},
		{AFLDLN2, psess.ynone, Px, opBytes{0xd9, 0xed}},
		{AFLDPI, psess.ynone, Px, opBytes{0xd9, 0xeb}},
		{AFLDZ, psess.ynone, Px, opBytes{0xd9, 0xee}},
		{AFNOP, psess.ynone, Px, opBytes{0xd9, 0xd0}},
		{AFPATAN, psess.ynone, Px, opBytes{0xd9, 0xf3}},
		{AFPREM, psess.ynone, Px, opBytes{0xd9, 0xf8}},
		{AFPREM1, psess.ynone, Px, opBytes{0xd9, 0xf5}},
		{AFPTAN, psess.ynone, Px, opBytes{0xd9, 0xf2}},
		{AFRNDINT, psess.ynone, Px, opBytes{0xd9, 0xfc}},
		{AFSCALE, psess.ynone, Px, opBytes{0xd9, 0xfd}},
		{AFSIN, psess.ynone, Px, opBytes{0xd9, 0xfe}},
		{AFSINCOS, psess.ynone, Px, opBytes{0xd9, 0xfb}},
		{AFSQRT, psess.ynone, Px, opBytes{0xd9, 0xfa}},
		{AFTST, psess.ynone, Px, opBytes{0xd9, 0xe4}},
		{AFXAM, psess.ynone, Px, opBytes{0xd9, 0xe5}},
		{AFXTRACT, psess.ynone, Px, opBytes{0xd9, 0xf4}},
		{AFYL2X, psess.ynone, Px, opBytes{0xd9, 0xf1}},
		{AFYL2XP1, psess.ynone, Px, opBytes{0xd9, 0xf9}},
		{ACMPXCHGB, psess.yrb_mb, Pb, opBytes{0x0f, 0xb0}},
		{ACMPXCHGL, psess.yrl_ml, Px, opBytes{0x0f, 0xb1}},
		{ACMPXCHGW, psess.yrl_ml, Pe, opBytes{0x0f, 0xb1}},
		{ACMPXCHGQ, psess.yrl_ml, Pw, opBytes{0x0f, 0xb1}},
		{ACMPXCHG8B, psess.yscond, Pm, opBytes{0xc7, 01}},
		{ACMPXCHG16B, psess.yscond, Pw, opBytes{0x0f, 0xc7, 01}},
		{AINVD, psess.ynone, Pm, opBytes{0x08}},
		{AINVLPG, psess.ydivb, Pm, opBytes{0x01, 07}},
		{AINVPCID, psess.ycrc32l, Pe, opBytes{0x0f, 0x38, 0x82, 0}},
		{ALFENCE, psess.ynone, Pm, opBytes{0xae, 0xe8}},
		{AMFENCE, psess.ynone, Pm, opBytes{0xae, 0xf0}},
		{AMOVNTIL, psess.yrl_ml, Pm, opBytes{0xc3}},
		{AMOVNTIQ, psess.yrl_ml, Pw, opBytes{0x0f, 0xc3}},
		{ARDPKRU, psess.ynone, Pm, opBytes{0x01, 0xee, 0}},
		{ARDMSR, psess.ynone, Pm, opBytes{0x32}},
		{ARDPMC, psess.ynone, Pm, opBytes{0x33}},
		{ARDTSC, psess.ynone, Pm, opBytes{0x31}},
		{ARSM, psess.ynone, Pm, opBytes{0xaa}},
		{ASFENCE, psess.ynone, Pm, opBytes{0xae, 0xf8}},
		{ASYSRET, psess.ynone, Pm, opBytes{0x07}},
		{AWBINVD, psess.ynone, Pm, opBytes{0x09}},
		{AWRMSR, psess.ynone, Pm, opBytes{0x30}},
		{AWRPKRU, psess.ynone, Pm, opBytes{0x01, 0xef, 0}},
		{AXADDB, psess.yrb_mb, Pb, opBytes{0x0f, 0xc0}},
		{AXADDL, psess.yrl_ml, Px, opBytes{0x0f, 0xc1}},
		{AXADDQ, psess.yrl_ml, Pw, opBytes{0x0f, 0xc1}},
		{AXADDW, psess.yrl_ml, Pe, opBytes{0x0f, 0xc1}},
		{ACRC32B, psess.ycrc32b, Px, opBytes{0xf2, 0x0f, 0x38, 0xf0, 0}},
		{ACRC32L, psess.ycrc32l, Px, opBytes{0xf2, 0x0f, 0x38, 0xf1, 0}},
		{ACRC32Q, psess.ycrc32l, Pw, opBytes{0xf2, 0x0f, 0x38, 0xf1, 0}},
		{ACRC32W, psess.ycrc32l, Pe, opBytes{0xf2, 0x0f, 0x38, 0xf1, 0}},
		{APREFETCHT0, psess.yprefetch, Pm, opBytes{0x18, 01}},
		{APREFETCHT1, psess.yprefetch, Pm, opBytes{0x18, 02}},
		{APREFETCHT2, psess.yprefetch, Pm, opBytes{0x18, 03}},
		{APREFETCHNTA, psess.yprefetch, Pm, opBytes{0x18, 00}},
		{AMOVQL, psess.yrl_ml, Px, opBytes{0x89}},
		{obj.AUNDEF, psess.ynone, Px, opBytes{0x0f, 0x0b}},
		{AAESENC, psess.yaes, Pq, opBytes{0x38, 0xdc, 0}},
		{AAESENCLAST, psess.yaes, Pq, opBytes{0x38, 0xdd, 0}},
		{AAESDEC, psess.yaes, Pq, opBytes{0x38, 0xde, 0}},
		{AAESDECLAST, psess.yaes, Pq, opBytes{0x38, 0xdf, 0}},
		{AAESIMC, psess.yaes, Pq, opBytes{0x38, 0xdb, 0}},
		{AAESKEYGENASSIST, psess.yxshuf, Pq, opBytes{0x3a, 0xdf, 0}},
		{AROUNDPD, psess.yxshuf, Pq, opBytes{0x3a, 0x09, 0}},
		{AROUNDPS, psess.yxshuf, Pq, opBytes{0x3a, 0x08, 0}},
		{AROUNDSD, psess.yxshuf, Pq, opBytes{0x3a, 0x0b, 0}},
		{AROUNDSS, psess.yxshuf, Pq, opBytes{0x3a, 0x0a, 0}},
		{APSHUFD, psess.yxshuf, Pq, opBytes{0x70, 0}},
		{APCLMULQDQ, psess.yxshuf, Pq, opBytes{0x3a, 0x44, 0}},
		{APCMPESTRI, psess.yxshuf, Pq, opBytes{0x3a, 0x61, 0}},
		{APCMPESTRM, psess.yxshuf, Pq, opBytes{0x3a, 0x60, 0}},
		{AMOVDDUP, psess.yxm, Pf2, opBytes{0x12}},
		{AMOVSHDUP, psess.yxm, Pf3, opBytes{0x16}},
		{AMOVSLDUP, psess.yxm, Pf3, opBytes{0x12}},

		{ARDTSCP, psess.ynone, Pm, opBytes{0x01, 0xf9, 0}},
		{ASTAC, psess.ynone, Pm, opBytes{0x01, 0xcb, 0}},
		{AUD1, psess.ynone, Pm, opBytes{0xb9, 0}},
		{AUD2, psess.ynone, Pm, opBytes{0x0b, 0}},
		{ASYSENTER, psess.ynone, Px, opBytes{0x0f, 0x34, 0}},
		{ASYSENTER64, psess.ynone, Pw, opBytes{0x0f, 0x34, 0}},
		{ASYSEXIT, psess.ynone, Px, opBytes{0x0f, 0x35, 0}},
		{ASYSEXIT64, psess.ynone, Pw, opBytes{0x0f, 0x35, 0}},
		{ALMSW, psess.ydivl, Pm, opBytes{0x01, 06}},
		{ALLDT, psess.ydivl, Pm, opBytes{0x00, 02}},
		{ALIDT, psess.ysvrs_mo, Pm, opBytes{0x01, 03}},
		{ALGDT, psess.ysvrs_mo, Pm, opBytes{0x01, 02}},
		{ATZCNTW, psess.ycrc32l, Pe, opBytes{0xf3, 0x0f, 0xbc, 0}},
		{ATZCNTL, psess.ycrc32l, Px, opBytes{0xf3, 0x0f, 0xbc, 0}},
		{ATZCNTQ, psess.ycrc32l, Pw, opBytes{0xf3, 0x0f, 0xbc, 0}},
		{AXRSTOR, psess.ydivl, Px, opBytes{0x0f, 0xae, 05}},
		{AXRSTOR64, psess.ydivl, Pw, opBytes{0x0f, 0xae, 05}},
		{AXRSTORS, psess.ydivl, Px, opBytes{0x0f, 0xc7, 03}},
		{AXRSTORS64, psess.ydivl, Pw, opBytes{0x0f, 0xc7, 03}},
		{AXSAVE, psess.yclflush, Px, opBytes{0x0f, 0xae, 04}},
		{AXSAVE64, psess.yclflush, Pw, opBytes{0x0f, 0xae, 04}},
		{AXSAVEOPT, psess.yclflush, Px, opBytes{0x0f, 0xae, 06}},
		{AXSAVEOPT64, psess.yclflush, Pw, opBytes{0x0f, 0xae, 06}},
		{AXSAVEC, psess.yclflush, Px, opBytes{0x0f, 0xc7, 04}},
		{AXSAVEC64, psess.yclflush, Pw, opBytes{0x0f, 0xc7, 04}},
		{AXSAVES, psess.yclflush, Px, opBytes{0x0f, 0xc7, 05}},
		{AXSAVES64, psess.yclflush, Pw, opBytes{0x0f, 0xc7, 05}},
		{ASGDT, psess.yclflush, Pm, opBytes{0x01, 00}},
		{ASIDT, psess.yclflush, Pm, opBytes{0x01, 01}},
		{ARDRANDW, psess.yrdrand, Pe, opBytes{0x0f, 0xc7, 06}},
		{ARDRANDL, psess.yrdrand, Px, opBytes{0x0f, 0xc7, 06}},
		{ARDRANDQ, psess.yrdrand, Pw, opBytes{0x0f, 0xc7, 06}},
		{ARDSEEDW, psess.yrdrand, Pe, opBytes{0x0f, 0xc7, 07}},
		{ARDSEEDL, psess.yrdrand, Px, opBytes{0x0f, 0xc7, 07}},
		{ARDSEEDQ, psess.yrdrand, Pw, opBytes{0x0f, 0xc7, 07}},
		{ASTRW, psess.yincq, Pe, opBytes{0x0f, 0x00, 01}},
		{ASTRL, psess.yincq, Px, opBytes{0x0f, 0x00, 01}},
		{ASTRQ, psess.yincq, Pw, opBytes{0x0f, 0x00, 01}},
		{AXSETBV, psess.ynone, Pm, opBytes{0x01, 0xd1, 0}},
		{AMOVBEWW, psess.ymovbe, Pq, opBytes{0x38, 0xf0, 0, 0x38, 0xf1, 0}},
		{AMOVBELL, psess.ymovbe, Pm, opBytes{0x38, 0xf0, 0, 0x38, 0xf1, 0}},
		{AMOVBEQQ, psess.ymovbe, Pw, opBytes{0x0f, 0x38, 0xf0, 0, 0x0f, 0x38, 0xf1, 0}},
		{ANOPW, psess.ydivl, Pe, opBytes{0x0f, 0x1f, 00}},
		{ANOPL, psess.ydivl, Px, opBytes{0x0f, 0x1f, 00}},
		{ASLDTW, psess.yincq, Pe, opBytes{0x0f, 0x00, 00}},
		{ASLDTL, psess.yincq, Px, opBytes{0x0f, 0x00, 00}},
		{ASLDTQ, psess.yincq, Pw, opBytes{0x0f, 0x00, 00}},
		{ASMSWW, psess.yincq, Pe, opBytes{0x0f, 0x01, 04}},
		{ASMSWL, psess.yincq, Px, opBytes{0x0f, 0x01, 04}},
		{ASMSWQ, psess.yincq, Pw, opBytes{0x0f, 0x01, 04}},
		{ABLENDVPS, psess.yblendvpd, Pq4, opBytes{0x14}},
		{ABLENDVPD, psess.yblendvpd, Pq4, opBytes{0x15}},
		{APBLENDVB, psess.yblendvpd, Pq4, opBytes{0x10}},
		{ASHA1MSG1, psess.yaes, Px, opBytes{0x0f, 0x38, 0xc9, 0}},
		{ASHA1MSG2, psess.yaes, Px, opBytes{0x0f, 0x38, 0xca, 0}},
		{ASHA1NEXTE, psess.yaes, Px, opBytes{0x0f, 0x38, 0xc8, 0}},
		{ASHA256MSG1, psess.yaes, Px, opBytes{0x0f, 0x38, 0xcc, 0}},
		{ASHA256MSG2, psess.yaes, Px, opBytes{0x0f, 0x38, 0xcd, 0}},
		{ASHA1RNDS4, psess.ysha1rnds4, Pm, opBytes{0x3a, 0xcc, 0}},
		{ASHA256RNDS2, psess.ysha256rnds2, Px, opBytes{0x0f, 0x38, 0xcb, 0}},
		{ARDFSBASEL, psess.yrdrand, Pf3, opBytes{0xae, 00}},
		{ARDFSBASEQ, psess.yrdrand, Pfw, opBytes{0xae, 00}},
		{ARDGSBASEL, psess.yrdrand, Pf3, opBytes{0xae, 01}},
		{ARDGSBASEQ, psess.yrdrand, Pfw, opBytes{0xae, 01}},
		{AWRFSBASEL, psess.ywrfsbase, Pf3, opBytes{0xae, 02}},
		{AWRFSBASEQ, psess.ywrfsbase, Pfw, opBytes{0xae, 02}},
		{AWRGSBASEL, psess.ywrfsbase, Pf3, opBytes{0xae, 03}},
		{AWRGSBASEQ, psess.ywrfsbase, Pfw, opBytes{0xae, 03}},
		{ALFSW, psess.ym_rl, Pe, opBytes{0x0f, 0xb4}},
		{ALFSL, psess.ym_rl, Px, opBytes{0x0f, 0xb4}},
		{ALFSQ, psess.ym_rl, Pw, opBytes{0x0f, 0xb4}},
		{ALGSW, psess.ym_rl, Pe, opBytes{0x0f, 0xb5}},
		{ALGSL, psess.ym_rl, Px, opBytes{0x0f, 0xb5}},
		{ALGSQ, psess.ym_rl, Pw, opBytes{0x0f, 0xb5}},
		{ALSSW, psess.ym_rl, Pe, opBytes{0x0f, 0xb2}},
		{ALSSL, psess.ym_rl, Px, opBytes{0x0f, 0xb2}},
		{ALSSQ, psess.ym_rl, Pw, opBytes{0x0f, 0xb2}},

		{ABLENDPD, psess.yxshuf, Pq, opBytes{0x3a, 0x0d, 0}},
		{ABLENDPS, psess.yxshuf, Pq, opBytes{0x3a, 0x0c, 0}},
		{AXACQUIRE, psess.ynone, Px, opBytes{0xf2}},
		{AXRELEASE, psess.ynone, Px, opBytes{0xf3}},
		{AXBEGIN, psess.yxbegin, Px, opBytes{0xc7, 0xf8}},
		{AXABORT, psess.yxabort, Px, opBytes{0xc6, 0xf8}},
		{AXEND, psess.ynone, Px, opBytes{0x0f, 01, 0xd5}},
		{AXTEST, psess.ynone, Px, opBytes{0x0f, 01, 0xd6}},
		{AXGETBV, psess.ynone, Pm, opBytes{01, 0xd0}},
		{obj.AFUNCDATA, psess.yfuncdata, Px, opBytes{0, 0}},
		{obj.APCDATA, psess.ypcdata, Px, opBytes{0, 0}},
		{obj.ADUFFCOPY, psess.yduff, Px, opBytes{0xe8}},
		{obj.ADUFFZERO, psess.yduff, Px, opBytes{0xe8}},

		{obj.AEND, nil, 0, opBytes{}},
		{0, nil, 0, opBytes{}},
	}
	psess._yandnl = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yml, Yrl, Yrl}},
	}
	psess._ybextrl = []ytab{
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yrl, Yml, Yrl}},
	}
	psess._yblsil = []ytab{
		{zcase: Zvex_rm_r_vo, zoffset: 3, args: argList{Yml, Yrl}},
	}
	psess._ykaddb = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yk, Yk, Yk}},
	}
	psess._ykmovb = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yk, Ym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yk, Yrl}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ykm, Yk}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yrl, Yk}},
	}
	psess._yknotb = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yk, Yk}},
	}
	psess._ykshiftlb = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yu8, Yk, Yk}},
	}
	psess._yrorxl = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yml, Yrl}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yml, Yrl}},
	}
	psess._yv4fmaddps = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, YzrMulti4, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Ym, YzrMulti4, Yknot0, Yzr}},
	}
	psess._yv4fmaddss = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, YxrEvexMulti4, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Ym, YxrEvexMulti4, Yknot0, YxrEvex}},
	}
	psess._yvaddpd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
	}
	psess._yvaddsd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
	}
	psess._yvaddsubpd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
	}
	psess._yvaesdec = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yzm, Yzr, Yzr}},
	}
	psess._yvaesimc = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
	}
	psess._yvaeskeygenassist = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yxm, Yxr}},
	}
	psess._yvalignd = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvandnpd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvblendmpd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvblendpd = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yym, Yyr, Yyr}},
	}
	psess._yvblendvpd = []ytab{
		{zcase: Zvex_hr_rm_v_r, zoffset: 2, args: argList{Yxr, Yxm, Yxr, Yxr}},
		{zcase: Zvex_hr_rm_v_r, zoffset: 2, args: argList{Yyr, Yym, Yyr, Yyr}},
	}
	psess._yvbroadcastf128 = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yyr}},
	}
	psess._yvbroadcastf32x2 = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, Yzr}},
	}
	psess._yvbroadcastf32x4 = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Ym, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Ym, Yknot0, Yzr}},
	}
	psess._yvbroadcastf32x8 = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Ym, Yknot0, Yzr}},
	}
	psess._yvbroadcasti32x2 = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, Yzr}},
	}
	psess._yvbroadcastsd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, Yzr}},
	}
	psess._yvbroadcastss = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, Yzr}},
	}
	psess._yvcmppd = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yym, Yyr, Yyr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yk}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, Yk}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, Yk}},
	}
	psess._yvcmpsd = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, Yk}},
	}
	psess._yvcomisd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxmEvex, YxrEvex}},
	}
	psess._yvcompresspd = []ytab{
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YyrEvex, YymEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YymEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{Yzr, Yzm}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, Yzm}},
	}
	psess._yvcvtdq2pd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, Yzr}},
	}
	psess._yvcvtdq2ps = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
	}
	psess._yvcvtpd2dq = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, YyrEvex}},
	}
	psess._yvcvtpd2dqx = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
	}
	psess._yvcvtpd2dqy = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YxrEvex}},
	}
	psess._yvcvtpd2qq = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
	}
	psess._yvcvtpd2udqx = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
	}
	psess._yvcvtpd2udqy = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YxrEvex}},
	}
	psess._yvcvtph2ps = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
	}
	psess._yvcvtps2ph = []ytab{
		{zcase: Zvex_i_r_rm, zoffset: 0, args: argList{Yu8, Yxr, Yxm}},
		{zcase: Zvex_i_r_rm, zoffset: 2, args: argList{Yi8, Yxr, Yxm}},
		{zcase: Zvex_i_r_rm, zoffset: 0, args: argList{Yu8, Yyr, Yxm}},
		{zcase: Zvex_i_r_rm, zoffset: 2, args: argList{Yi8, Yyr, Yxm}},
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, Yzr, YymEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, Yzr, Yknot0, YymEvex}},
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, YxrEvex, YxmEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, YyrEvex, YxmEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, YyrEvex, Yknot0, YxmEvex}},
	}
	psess._yvcvtps2qq = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
	}
	psess._yvcvtsd2si = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yrl}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxmEvex, Yrl}},
	}
	psess._yvcvtsd2usil = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxmEvex, Yrl}},
	}
	psess._yvcvtsi2sdl = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yml, Yxr, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yml, YxrEvex, YxrEvex}},
	}
	psess._yvcvtudq2pd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, Yzr}},
	}
	psess._yvcvtusi2sdl = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yml, YxrEvex, YxrEvex}},
	}
	psess._yvdppd = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
	}
	psess._yvexp2pd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
	}
	psess._yvexpandpd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
	}
	psess._yvextractf128 = []ytab{
		{zcase: Zvex_i_r_rm, zoffset: 0, args: argList{Yu8, Yyr, Yxm}},
		{zcase: Zvex_i_r_rm, zoffset: 2, args: argList{Yi8, Yyr, Yxm}},
	}
	psess._yvextractf32x4 = []ytab{
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, YyrEvex, YxmEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, YyrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, Yzr, YxmEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, Yzr, Yknot0, YxmEvex}},
	}
	psess._yvextractf32x8 = []ytab{
		{zcase: Zevex_i_r_rm, zoffset: 0, args: argList{Yu8, Yzr, YymEvex}},
		{zcase: Zevex_i_r_k_rm, zoffset: 3, args: argList{Yu8, Yzr, Yknot0, YymEvex}},
	}
	psess._yvextractps = []ytab{
		{zcase: Zvex_i_r_rm, zoffset: 0, args: argList{Yu8, Yxr, Yml}},
		{zcase: Zvex_i_r_rm, zoffset: 2, args: argList{Yi8, Yxr, Yml}},
		{zcase: Zevex_i_r_rm, zoffset: 3, args: argList{Yu8, YxrEvex, Yml}},
	}
	psess._yvfixupimmpd = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yzr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, YyrEvex}},
	}
	psess._yvfixupimmsd = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, YxrEvex}},
	}
	psess._yvfpclasspdx = []ytab{
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YxmEvex, Yk}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, Yknot0, Yk}},
	}
	psess._yvfpclasspdy = []ytab{
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, Yk}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, Yk}},
	}
	psess._yvfpclasspdz = []ytab{
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yk}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yk}},
	}
	psess._yvgatherdpd = []ytab{
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yxr, Yxvm, Yxr}},
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yyr, Yxvm, Yyr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxvmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxvmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YyvmEvex, Yknot0, Yzr}},
	}
	psess._yvgatherdps = []ytab{
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yxr, Yxvm, Yxr}},
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yyr, Yyvm, Yyr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxvmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YyvmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzvm, Yknot0, Yzr}},
	}
	psess._yvgatherpf0dpd = []ytab{
		{zcase: Zevex_k_rmo, zoffset: 4, args: argList{Yknot0, YyvmEvex}},
	}
	psess._yvgatherpf0dps = []ytab{
		{zcase: Zevex_k_rmo, zoffset: 4, args: argList{Yknot0, Yzvm}},
	}
	psess._yvgatherqps = []ytab{
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yxr, Yxvm, Yxr}},
		{zcase: Zvex_v_rm_r, zoffset: 2, args: argList{Yxr, Yyvm, Yxr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxvmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YyvmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzvm, Yknot0, YyrEvex}},
	}
	psess._yvgetexpsd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
	}
	psess._yvgetmantpd = []ytab{
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
	}
	psess._yvgf2p8affineinvqb = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yym, Yyr, Yyr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvinsertf128 = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yyr, Yyr}},
	}
	psess._yvinsertf32x4 = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, Yzr, Yknot0, Yzr}},
	}
	psess._yvinsertf32x8 = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yzr, Yknot0, Yzr}},
	}
	psess._yvinsertps = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
	}
	psess._yvlddqu = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yyr}},
	}
	psess._yvldmxcsr = []ytab{
		{zcase: Zvex_rm_v_ro, zoffset: 3, args: argList{Ym}},
	}
	psess._yvmaskmovdqu = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxr, Yxr}},
	}
	psess._yvmaskmovpd = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yxr, Ym}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yyr, Yyr, Ym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yyr, Yyr}},
	}
	psess._yvmovapd = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yxm}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yyr, Yym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YyrEvex, YymEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YymEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{Yzr, Yzm}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, Yzm}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
	}
	psess._yvmovd = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yml}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yml, Yxr}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YxrEvex, Yml}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yml, YxrEvex}},
	}
	psess._yvmovddup = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
	}
	psess._yvmovdqa = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yxm}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yyr, Yym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr}},
	}
	psess._yvmovdqa32 = []ytab{
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YyrEvex, YymEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YymEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{Yzr, Yzm}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, Yzm}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yzm, Yknot0, Yzr}},
	}
	psess._yvmovhlps = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxr, Yxr, Yxr}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxrEvex, YxrEvex, YxrEvex}},
	}
	psess._yvmovhpd = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Ym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yxr, Yxr}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YxrEvex, Ym}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Ym, YxrEvex, YxrEvex}},
	}
	psess._yvmovmskpd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxr, Yrl}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yyr, Yrl}},
	}
	psess._yvmovntdq = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Ym}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yyr, Ym}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YxrEvex, Ym}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YyrEvex, Ym}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{Yzr, Ym}},
	}
	psess._yvmovntdqa = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Ym, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Ym, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Ym, Yzr}},
	}
	psess._yvmovq = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yml}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yxm}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yml, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YxrEvex, Yml}},
		{zcase: Zevex_r_v_rm, zoffset: 3, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yml, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxmEvex, YxrEvex}},
	}
	psess._yvmovsd = []ytab{
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Yxr, Yxr}},
		{zcase: Zvex_r_v_rm, zoffset: 2, args: argList{Yxr, Ym}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Ym, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxr, Yxr, Yxr}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_r_v_k_rm, zoffset: 3, args: argList{YxrEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, Ym}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, Ym}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Ym, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Ym, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxrEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxrEvex, YxrEvex, Yknot0, YxrEvex}},
	}
	psess._yvpbroadcastb = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yrl, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yrl, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yrl, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yrl, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yrl, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{Yrl, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr}},
		{zcase: Zevex_rm_k_r, zoffset: 3, args: argList{YxmEvex, Yknot0, Yzr}},
	}
	psess._yvpbroadcastmb2q = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yk, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yk, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yk, Yzr}},
	}
	psess._yvpclmulqdq = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yxm, Yxr, Yxr}},
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yym, Yyr, Yyr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yzr}},
	}
	psess._yvpcmpb = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, YxrEvex, Yknot0, Yk}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, Yk}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yk}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yk}},
	}
	psess._yvpcmpeqb = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yk}},
	}
	psess._yvperm2f128 = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yym, Yyr, Yyr}},
	}
	psess._yvpermd = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvpermilpd = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yym, Yyr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yym, Yyr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr, Yyr}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvpermpd = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yu8, Yym, Yyr}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvpermq = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yym, Yyr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yym, Yyr}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvpextrw = []ytab{
		{zcase: Zvex_i_r_rm, zoffset: 0, args: argList{Yu8, Yxr, Yml}},
		{zcase: Zvex_i_r_rm, zoffset: 2, args: argList{Yi8, Yxr, Yml}},
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yxr, Yrl}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yxr, Yrl}},
		{zcase: Zevex_i_r_rm, zoffset: 3, args: argList{Yu8, YxrEvex, Yml}},
		{zcase: Zevex_i_rm_r, zoffset: 3, args: argList{Yu8, YxrEvex, Yrl}},
	}
	psess._yvpinsrb = []ytab{
		{zcase: Zvex_i_rm_v_r, zoffset: 2, args: argList{Yu8, Yml, Yxr, Yxr}},
		{zcase: Zevex_i_rm_v_r, zoffset: 3, args: argList{Yu8, Yml, YxrEvex, YxrEvex}},
	}
	psess._yvpmovb2m = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YxrEvex, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{YyrEvex, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 3, args: argList{Yzr, Yk}},
	}
	psess._yvpmovdb = []ytab{
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YyrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{Yzr, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, YxmEvex}},
	}
	psess._yvpmovdw = []ytab{
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YxrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{YyrEvex, YxmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YxmEvex}},
		{zcase: Zevex_r_v_rm, zoffset: 0, args: argList{Yzr, YymEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, YymEvex}},
	}
	psess._yvprold = []ytab{
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, Yzm, Yknot0, Yzr}},
	}
	psess._yvpscatterdd = []ytab{
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YyvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, Yzvm}},
	}
	psess._yvpscatterdq = []ytab{
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, YxvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{Yzr, Yknot0, YyvmEvex}},
	}
	psess._yvpscatterqd = []ytab{
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YxvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YxrEvex, Yknot0, YyvmEvex}},
		{zcase: Zevex_r_k_rm, zoffset: 3, args: argList{YyrEvex, Yknot0, Yzvm}},
	}
	psess._yvpshufbitqmb = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, Yk}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yk}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yk}},
	}
	psess._yvpshufd = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yym, Yyr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yym, Yyr}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yknot0, Yzr}},
	}
	psess._yvpslld = []ytab{
		{zcase: Zvex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yxr, Yxr}},
		{zcase: Zvex_i_rm_vo, zoffset: 3, args: argList{Yi8, Yxr, Yxr}},
		{zcase: Zvex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yyr, Yyr}},
		{zcase: Zvex_i_rm_vo, zoffset: 3, args: argList{Yi8, Yyr, Yyr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yyr, Yyr}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, Yzr, Yknot0, Yzr}},
	}
	psess._yvpslldq = []ytab{
		{zcase: Zvex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yxr, Yxr}},
		{zcase: Zvex_i_rm_vo, zoffset: 3, args: argList{Yi8, Yxr, Yxr}},
		{zcase: Zvex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yyr, Yyr}},
		{zcase: Zvex_i_rm_vo, zoffset: 3, args: argList{Yi8, Yyr, Yyr}},
		{zcase: Zevex_i_rm_vo, zoffset: 4, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 4, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 4, args: argList{Yu8, Yzm, Yzr}},
	}
	psess._yvpsraq = []ytab{
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YxmEvex, YxrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YxmEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, YymEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_vo, zoffset: 0, args: argList{Yu8, Yzm, Yzr}},
		{zcase: Zevex_i_rm_k_vo, zoffset: 4, args: argList{Yu8, Yzm, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, Yzr, Yknot0, Yzr}},
	}
	psess._yvptest = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr}},
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yym, Yyr}},
	}
	psess._yvrcpss = []ytab{
		{zcase: Zvex_rm_v_r, zoffset: 2, args: argList{Yxm, Yxr, Yxr}},
	}
	psess._yvroundpd = []ytab{
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yxm, Yxr}},
		{zcase: Zvex_i_rm_r, zoffset: 0, args: argList{Yu8, Yym, Yyr}},
		{zcase: Zvex_i_rm_r, zoffset: 2, args: argList{Yi8, Yym, Yyr}},
	}
	psess._yvscalefpd = []ytab{
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{Yzm, Yzr, Yzr}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{Yzm, Yzr, Yknot0, Yzr}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YxmEvex, YxrEvex, YxrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YxmEvex, YxrEvex, Yknot0, YxrEvex}},
		{zcase: Zevex_rm_v_r, zoffset: 0, args: argList{YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_rm_v_k_r, zoffset: 3, args: argList{YymEvex, YyrEvex, Yknot0, YyrEvex}},
	}
	psess._yvshuff32x4 = []ytab{
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, YymEvex, YyrEvex, YyrEvex}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, YymEvex, YyrEvex, Yknot0, YyrEvex}},
		{zcase: Zevex_i_rm_v_r, zoffset: 0, args: argList{Yu8, Yzm, Yzr, Yzr}},
		{zcase: Zevex_i_rm_v_k_r, zoffset: 3, args: argList{Yu8, Yzm, Yzr, Yknot0, Yzr}},
	}
	psess._yvzeroall = []ytab{
		{zcase: Zvex, zoffset: 2, args: argList{}},
	}
	psess.avxOptab = [...]Optab{
		{as: AANDNL, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF2,
		}},
		{as: AANDNQ, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF2,
		}},
		{as: ABEXTRL, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF7,
		}},
		{as: ABEXTRQ, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF7,
		}},
		{as: ABLSIL, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF3, 03,
		}},
		{as: ABLSIQ, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF3, 03,
		}},
		{as: ABLSMSKL, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF3, 02,
		}},
		{as: ABLSMSKQ, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF3, 02,
		}},
		{as: ABLSRL, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF3, 01,
		}},
		{as: ABLSRQ, ytab: psess._yblsil, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF3, 01,
		}},
		{as: ABZHIL, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW0, 0xF5,
		}},
		{as: ABZHIQ, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F38 | vexW1, 0xF5,
		}},
		{as: AKADDB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x4A,
		}},
		{as: AKADDD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x4A,
		}},
		{as: AKADDQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x4A,
		}},
		{as: AKADDW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x4A,
		}},
		{as: AKANDB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x41,
		}},
		{as: AKANDD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x41,
		}},
		{as: AKANDNB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x42,
		}},
		{as: AKANDND, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x42,
		}},
		{as: AKANDNQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x42,
		}},
		{as: AKANDNW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x42,
		}},
		{as: AKANDQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x41,
		}},
		{as: AKANDW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x41,
		}},
		{as: AKMOVB, ytab: psess._ykmovb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x91,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x93,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x90,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x92,
		}},
		{as: AKMOVD, ytab: psess._ykmovb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x91,
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x93,
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x90,
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x92,
		}},
		{as: AKMOVQ, ytab: psess._ykmovb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW1, 0x91,
			avxEscape | vex128 | vexF2 | vex0F | vexW1, 0x93,
			avxEscape | vex128 | vex0F | vexW1, 0x90,
			avxEscape | vex128 | vexF2 | vex0F | vexW1, 0x92,
		}},
		{as: AKMOVW, ytab: psess._ykmovb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x91,
			avxEscape | vex128 | vex0F | vexW0, 0x93,
			avxEscape | vex128 | vex0F | vexW0, 0x90,
			avxEscape | vex128 | vex0F | vexW0, 0x92,
		}},
		{as: AKNOTB, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x44,
		}},
		{as: AKNOTD, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x44,
		}},
		{as: AKNOTQ, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW1, 0x44,
		}},
		{as: AKNOTW, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x44,
		}},
		{as: AKORB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x45,
		}},
		{as: AKORD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x45,
		}},
		{as: AKORQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x45,
		}},
		{as: AKORTESTB, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x98,
		}},
		{as: AKORTESTD, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x98,
		}},
		{as: AKORTESTQ, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW1, 0x98,
		}},
		{as: AKORTESTW, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x98,
		}},
		{as: AKORW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x45,
		}},
		{as: AKSHIFTLB, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x32,
		}},
		{as: AKSHIFTLD, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x33,
		}},
		{as: AKSHIFTLQ, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x33,
		}},
		{as: AKSHIFTLW, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x32,
		}},
		{as: AKSHIFTRB, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x30,
		}},
		{as: AKSHIFTRD, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x31,
		}},
		{as: AKSHIFTRQ, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x31,
		}},
		{as: AKSHIFTRW, ytab: psess._ykshiftlb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x30,
		}},
		{as: AKTESTB, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x99,
		}},
		{as: AKTESTD, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x99,
		}},
		{as: AKTESTQ, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW1, 0x99,
		}},
		{as: AKTESTW, ytab: psess._yknotb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x99,
		}},
		{as: AKUNPCKBW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x4B,
		}},
		{as: AKUNPCKDQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x4B,
		}},
		{as: AKUNPCKWD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x4B,
		}},
		{as: AKXNORB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x46,
		}},
		{as: AKXNORD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x46,
		}},
		{as: AKXNORQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x46,
		}},
		{as: AKXNORW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x46,
		}},
		{as: AKXORB, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x47,
		}},
		{as: AKXORD, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW1, 0x47,
		}},
		{as: AKXORQ, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW1, 0x47,
		}},
		{as: AKXORW, ytab: psess._ykaddb, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x47,
		}},
		{as: AMULXL, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW0, 0xF6,
		}},
		{as: AMULXQ, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW1, 0xF6,
		}},
		{as: APDEPL, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW0, 0xF5,
		}},
		{as: APDEPQ, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW1, 0xF5,
		}},
		{as: APEXTL, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F38 | vexW0, 0xF5,
		}},
		{as: APEXTQ, ytab: psess._yandnl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F38 | vexW1, 0xF5,
		}},
		{as: ARORXL, ytab: psess._yrorxl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F3A | vexW0, 0xF0,
		}},
		{as: ARORXQ, ytab: psess._yrorxl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F3A | vexW1, 0xF0,
		}},
		{as: ASARXL, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F38 | vexW0, 0xF7,
		}},
		{as: ASARXQ, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F38 | vexW1, 0xF7,
		}},
		{as: ASHLXL, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xF7,
		}},
		{as: ASHLXQ, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xF7,
		}},
		{as: ASHRXL, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW0, 0xF7,
		}},
		{as: ASHRXQ, ytab: psess._ybextrl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F38 | vexW1, 0xF7,
		}},
		{as: AV4FMADDPS, ytab: psess._yv4fmaddps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x9A,
		}},
		{as: AV4FMADDSS, ytab: psess._yv4fmaddss, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x9B,
		}},
		{as: AV4FNMADDPS, ytab: psess._yv4fmaddps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0xAA,
		}},
		{as: AV4FNMADDSS, ytab: psess._yv4fmaddss, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0xAB,
		}},
		{as: AVADDPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x58,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x58,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x58,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x58,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x58,
		}},
		{as: AVADDPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x58,
			avxEscape | vex256 | vex0F | vexW0, 0x58,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x58,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x58,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x58,
		}},
		{as: AVADDSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x58,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x58,
		}},
		{as: AVADDSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x58,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x58,
		}},
		{as: AVADDSUBPD, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD0,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD0,
		}},
		{as: AVADDSUBPS, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0xD0,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0xD0,
		}},
		{as: AVAESDEC, ytab: psess._yvaesdec, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xDE,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xDE,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0xDE,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0xDE,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0xDE,
		}},
		{as: AVAESDECLAST, ytab: psess._yvaesdec, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xDF,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xDF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0xDF,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0xDF,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0xDF,
		}},
		{as: AVAESENC, ytab: psess._yvaesdec, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xDC,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xDC,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0xDC,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0xDC,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0xDC,
		}},
		{as: AVAESENCLAST, ytab: psess._yvaesdec, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xDD,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xDD,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0xDD,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0xDD,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0xDD,
		}},
		{as: AVAESIMC, ytab: psess._yvaesimc, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xDB,
		}},
		{as: AVAESKEYGENASSIST, ytab: psess._yvaeskeygenassist, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0xDF,
		}},
		{as: AVALIGND, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x03,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x03,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x03,
		}},
		{as: AVALIGNQ, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x03,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x03,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x03,
		}},
		{as: AVANDNPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x55,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x55,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x55,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x55,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x55,
		}},
		{as: AVANDNPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x55,
			avxEscape | vex256 | vex0F | vexW0, 0x55,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x55,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x55,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x55,
		}},
		{as: AVANDPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x54,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x54,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x54,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x54,
		}},
		{as: AVANDPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x54,
			avxEscape | vex256 | vex0F | vexW0, 0x54,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x54,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x54,
		}},
		{as: AVBLENDMPD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x65,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x65,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x65,
		}},
		{as: AVBLENDMPS, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x65,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x65,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x65,
		}},
		{as: AVBLENDPD, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0D,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x0D,
		}},
		{as: AVBLENDPS, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0C,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x0C,
		}},
		{as: AVBLENDVPD, ytab: psess._yvblendvpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x4B,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x4B,
		}},
		{as: AVBLENDVPS, ytab: psess._yvblendvpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x4A,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x4A,
		}},
		{as: AVBROADCASTF128, ytab: psess._yvbroadcastf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x1A,
		}},
		{as: AVBROADCASTF32X2, ytab: psess._yvbroadcastf32x2, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x19,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x19,
		}},
		{as: AVBROADCASTF32X4, ytab: psess._yvbroadcastf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x1A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x1A,
		}},
		{as: AVBROADCASTF32X8, ytab: psess._yvbroadcastf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x1B,
		}},
		{as: AVBROADCASTF64X2, ytab: psess._yvbroadcastf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x1A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x1A,
		}},
		{as: AVBROADCASTF64X4, ytab: psess._yvbroadcastf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x1B,
		}},
		{as: AVBROADCASTI128, ytab: psess._yvbroadcastf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x5A,
		}},
		{as: AVBROADCASTI32X2, ytab: psess._yvbroadcasti32x2, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x59,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x59,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x59,
		}},
		{as: AVBROADCASTI32X4, ytab: psess._yvbroadcastf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x5A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x5A,
		}},
		{as: AVBROADCASTI32X8, ytab: psess._yvbroadcastf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVBROADCASTI64X2, ytab: psess._yvbroadcastf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x5A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x5A,
		}},
		{as: AVBROADCASTI64X4, ytab: psess._yvbroadcastf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVBROADCASTSD, ytab: psess._yvbroadcastsd, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x19,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x19,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x19,
		}},
		{as: AVBROADCASTSS, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x18,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x18,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x18,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x18,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x18,
		}},
		{as: AVCMPPD, ytab: psess._yvcmppd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xC2,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xC2,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled, 0xC2,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8, 0xC2,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8, 0xC2,
		}},
		{as: AVCMPPS, ytab: psess._yvcmppd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0xC2,
			avxEscape | vex256 | vex0F | vexW0, 0xC2,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled, 0xC2,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4, 0xC2,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4, 0xC2,
		}},
		{as: AVCMPSD, ytab: psess._yvcmpsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0xC2,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexSaeEnabled, 0xC2,
		}},
		{as: AVCMPSS, ytab: psess._yvcmpsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0xC2,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled, 0xC2,
		}},
		{as: AVCOMISD, ytab: psess._yvcomisd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x2F,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8 | evexSaeEnabled, 0x2F,
		}},
		{as: AVCOMISS, ytab: psess._yvcomisd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x2F,
			avxEscape | evex128 | evex0F | evexW0, evexN4 | evexSaeEnabled, 0x2F,
		}},
		{as: AVCOMPRESSPD, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8A,
		}},
		{as: AVCOMPRESSPS, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8A,
		}},
		{as: AVCVTDQ2PD, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0xE6,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0xE6,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0xE6,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xE6,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTDQ2PS, ytab: psess._yvcvtdq2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5B,
			avxEscape | vex256 | vex0F | vexW0, 0x5B,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5B,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5B,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTPD2DQ, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTPD2DQX, ytab: psess._yvcvtpd2dqx, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0xE6,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTPD2DQY, ytab: psess._yvcvtpd2dqy, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0xE6,
			avxEscape | evex256 | evexF2 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTPD2PS, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTPD2PSX, ytab: psess._yvcvtpd2dqx, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5A,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTPD2PSY, ytab: psess._yvcvtpd2dqy, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5A,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTPD2QQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x7B,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7B,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7B,
		}},
		{as: AVCVTPD2UDQ, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTPD2UDQX, ytab: psess._yvcvtpd2udqx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTPD2UDQY, ytab: psess._yvcvtpd2udqy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTPD2UQQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x79,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x79,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTPH2PS, ytab: psess._yvcvtph2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x13,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x13,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexSaeEnabled | evexZeroingEnabled, 0x13,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x13,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x13,
		}},
		{as: AVCVTPS2DQ, ytab: psess._yvcvtdq2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5B,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5B,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5B,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5B,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTPS2PD, ytab: psess._yvcvtph2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5A,
			avxEscape | vex256 | vex0F | vexW0, 0x5A,
			avxEscape | evex512 | evex0F | evexW0, evexN32 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x5A,
			avxEscape | evex128 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x5A,
			avxEscape | evex256 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTPS2PH, ytab: psess._yvcvtps2ph, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x1D,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x1D,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN32 | evexSaeEnabled | evexZeroingEnabled, 0x1D,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN8 | evexZeroingEnabled, 0x1D,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x1D,
		}},
		{as: AVCVTPS2QQ, ytab: psess._yvcvtps2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x7B,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x7B,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7B,
		}},
		{as: AVCVTPS2UDQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x79,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x79,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTPS2UQQ, ytab: psess._yvcvtps2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x79,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x79,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x79,
		}},
		{as: AVCVTQQ2PD, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF3 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xE6,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xE6,
			avxEscape | evex256 | evexF3 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTQQ2PS, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTQQ2PSX, ytab: psess._yvcvtpd2udqx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTQQ2PSY, ytab: psess._yvcvtpd2udqy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTSD2SI, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x2D,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN8 | evexRoundingEnabled, 0x2D,
		}},
		{as: AVCVTSD2SIQ, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW1, 0x2D,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x2D,
		}},
		{as: AVCVTSD2SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x5A,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTSD2USIL, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN8 | evexRoundingEnabled, 0x79,
		}},
		{as: AVCVTSD2USIQ, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x79,
		}},
		{as: AVCVTSI2SDL, ytab: psess._yvcvtsi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x2A,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN4, 0x2A,
		}},
		{as: AVCVTSI2SDQ, ytab: psess._yvcvtsi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW1, 0x2A,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x2A,
		}},
		{as: AVCVTSI2SSL, ytab: psess._yvcvtsi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x2A,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled, 0x2A,
		}},
		{as: AVCVTSI2SSQ, ytab: psess._yvcvtsi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW1, 0x2A,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x2A,
		}},
		{as: AVCVTSS2SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5A,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x5A,
		}},
		{as: AVCVTSS2SI, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x2D,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled, 0x2D,
		}},
		{as: AVCVTSS2SIQ, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW1, 0x2D,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN4 | evexRoundingEnabled, 0x2D,
		}},
		{as: AVCVTSS2USIL, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled, 0x79,
		}},
		{as: AVCVTSS2USIQ, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN4 | evexRoundingEnabled, 0x79,
		}},
		{as: AVCVTTPD2DQ, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTTPD2DQX, ytab: psess._yvcvtpd2dqx, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE6,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTTPD2DQY, ytab: psess._yvcvtpd2dqy, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE6,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xE6,
		}},
		{as: AVCVTTPD2QQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x7A,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTTPD2UDQ, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTPD2UDQX, ytab: psess._yvcvtpd2udqx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTPD2UDQY, ytab: psess._yvcvtpd2udqy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTPD2UQQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x78,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x78,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTPS2DQ, ytab: psess._yvcvtdq2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5B,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x5B,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x5B,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5B,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5B,
		}},
		{as: AVCVTTPS2QQ, ytab: psess._yvcvtps2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x7A,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTTPS2UDQ, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x78,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x78,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTPS2UQQ, ytab: psess._yvcvtps2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x78,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x78,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x78,
		}},
		{as: AVCVTTSD2SI, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x2C,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN8 | evexSaeEnabled, 0x2C,
		}},
		{as: AVCVTTSD2SIQ, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW1, 0x2C,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexSaeEnabled, 0x2C,
		}},
		{as: AVCVTTSD2USIL, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN8 | evexSaeEnabled, 0x78,
		}},
		{as: AVCVTTSD2USIQ, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexSaeEnabled, 0x78,
		}},
		{as: AVCVTTSS2SI, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x2C,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled, 0x2C,
		}},
		{as: AVCVTTSS2SIQ, ytab: psess._yvcvtsd2si, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW1, 0x2C,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN4 | evexSaeEnabled, 0x2C,
		}},
		{as: AVCVTTSS2USIL, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled, 0x78,
		}},
		{as: AVCVTTSS2USIQ, ytab: psess._yvcvtsd2usil, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN4 | evexSaeEnabled, 0x78,
		}},
		{as: AVCVTUDQ2PD, ytab: psess._yvcvtudq2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN8 | evexBcstN4 | evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7A,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUDQ2PS, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x7A,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evexF2 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUQQ2PD, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF3 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x7A,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evexF3 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUQQ2PS, ytab: psess._yvcvtpd2dq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUQQ2PSX, ytab: psess._yvcvtpd2udqx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUQQ2PSY, ytab: psess._yvcvtpd2udqy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evexF2 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7A,
		}},
		{as: AVCVTUSI2SDL, ytab: psess._yvcvtusi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN4, 0x7B,
		}},
		{as: AVCVTUSI2SDQ, ytab: psess._yvcvtusi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x7B,
		}},
		{as: AVCVTUSI2SSL, ytab: psess._yvcvtusi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled, 0x7B,
		}},
		{as: AVCVTUSI2SSQ, ytab: psess._yvcvtusi2sdl, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN8 | evexRoundingEnabled, 0x7B,
		}},
		{as: AVDBPSADBW, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x42,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x42,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexZeroingEnabled, 0x42,
		}},
		{as: AVDIVPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5E,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5E,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5E,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5E,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5E,
		}},
		{as: AVDIVPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5E,
			avxEscape | vex256 | vex0F | vexW0, 0x5E,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5E,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5E,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5E,
		}},
		{as: AVDIVSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x5E,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5E,
		}},
		{as: AVDIVSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5E,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5E,
		}},
		{as: AVDPPD, ytab: psess._yvdppd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x41,
		}},
		{as: AVDPPS, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x40,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x40,
		}},
		{as: AVEXP2PD, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0xC8,
		}},
		{as: AVEXP2PS, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0xC8,
		}},
		{as: AVEXPANDPD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x88,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x88,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x88,
		}},
		{as: AVEXPANDPS, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x88,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x88,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x88,
		}},
		{as: AVEXTRACTF128, ytab: psess._yvextractf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x19,
		}},
		{as: AVEXTRACTF32X4, ytab: psess._yvextractf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x19,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x19,
		}},
		{as: AVEXTRACTF32X8, ytab: psess._yvextractf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x1B,
		}},
		{as: AVEXTRACTF64X2, ytab: psess._yvextractf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x19,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x19,
		}},
		{as: AVEXTRACTF64X4, ytab: psess._yvextractf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x1B,
		}},
		{as: AVEXTRACTI128, ytab: psess._yvextractf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x39,
		}},
		{as: AVEXTRACTI32X4, ytab: psess._yvextractf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x39,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x39,
		}},
		{as: AVEXTRACTI32X8, ytab: psess._yvextractf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x3B,
		}},
		{as: AVEXTRACTI64X2, ytab: psess._yvextractf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x39,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x39,
		}},
		{as: AVEXTRACTI64X4, ytab: psess._yvextractf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x3B,
		}},
		{as: AVEXTRACTPS, ytab: psess._yvextractps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x17,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4, 0x17,
		}},
		{as: AVFIXUPIMMPD, ytab: psess._yvfixupimmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x54,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x54,
		}},
		{as: AVFIXUPIMMPS, ytab: psess._yvfixupimmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x54,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x54,
		}},
		{as: AVFIXUPIMMSD, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x55,
		}},
		{as: AVFIXUPIMMSS, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x55,
		}},
		{as: AVFMADD132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x98,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x98,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x98,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x98,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x98,
		}},
		{as: AVFMADD132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x98,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x98,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x98,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x98,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x98,
		}},
		{as: AVFMADD132SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x99,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x99,
		}},
		{as: AVFMADD132SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x99,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x99,
		}},
		{as: AVFMADD213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xA8,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xA8,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xA8,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xA8,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xA8,
		}},
		{as: AVFMADD213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xA8,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xA8,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xA8,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xA8,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xA8,
		}},
		{as: AVFMADD213SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xA9,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xA9,
		}},
		{as: AVFMADD213SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xA9,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xA9,
		}},
		{as: AVFMADD231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xB8,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xB8,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xB8,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xB8,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xB8,
		}},
		{as: AVFMADD231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xB8,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xB8,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xB8,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xB8,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xB8,
		}},
		{as: AVFMADD231SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xB9,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xB9,
		}},
		{as: AVFMADD231SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xB9,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xB9,
		}},
		{as: AVFMADDSUB132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x96,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x96,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x96,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x96,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x96,
		}},
		{as: AVFMADDSUB132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x96,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x96,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x96,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x96,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x96,
		}},
		{as: AVFMADDSUB213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xA6,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xA6,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xA6,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xA6,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xA6,
		}},
		{as: AVFMADDSUB213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xA6,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xA6,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xA6,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xA6,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xA6,
		}},
		{as: AVFMADDSUB231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xB6,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xB6,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xB6,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xB6,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xB6,
		}},
		{as: AVFMADDSUB231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xB6,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xB6,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xB6,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xB6,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xB6,
		}},
		{as: AVFMSUB132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9A,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x9A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9A,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x9A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x9A,
		}},
		{as: AVFMSUB132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9A,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x9A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9A,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x9A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x9A,
		}},
		{as: AVFMSUB132SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9B,
		}},
		{as: AVFMSUB132SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9B,
		}},
		{as: AVFMSUB213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAA,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xAA,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAA,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xAA,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xAA,
		}},
		{as: AVFMSUB213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAA,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xAA,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAA,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xAA,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xAA,
		}},
		{as: AVFMSUB213SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAB,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAB,
		}},
		{as: AVFMSUB213SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAB,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAB,
		}},
		{as: AVFMSUB231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBA,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xBA,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBA,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xBA,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xBA,
		}},
		{as: AVFMSUB231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBA,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xBA,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBA,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xBA,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xBA,
		}},
		{as: AVFMSUB231SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBB,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBB,
		}},
		{as: AVFMSUB231SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBB,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBB,
		}},
		{as: AVFMSUBADD132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x97,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x97,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x97,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x97,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x97,
		}},
		{as: AVFMSUBADD132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x97,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x97,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x97,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x97,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x97,
		}},
		{as: AVFMSUBADD213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xA7,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xA7,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xA7,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xA7,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xA7,
		}},
		{as: AVFMSUBADD213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xA7,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xA7,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xA7,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xA7,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xA7,
		}},
		{as: AVFMSUBADD231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xB7,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xB7,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xB7,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xB7,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xB7,
		}},
		{as: AVFMSUBADD231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xB7,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xB7,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xB7,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xB7,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xB7,
		}},
		{as: AVFNMADD132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x9C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x9C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x9C,
		}},
		{as: AVFNMADD132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x9C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x9C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x9C,
		}},
		{as: AVFNMADD132SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9D,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9D,
		}},
		{as: AVFNMADD132SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9D,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9D,
		}},
		{as: AVFNMADD213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAC,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xAC,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAC,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xAC,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xAC,
		}},
		{as: AVFNMADD213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAC,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xAC,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAC,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xAC,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xAC,
		}},
		{as: AVFNMADD213SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAD,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAD,
		}},
		{as: AVFNMADD213SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAD,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAD,
		}},
		{as: AVFNMADD231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBC,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xBC,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBC,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xBC,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xBC,
		}},
		{as: AVFNMADD231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBC,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xBC,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBC,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xBC,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xBC,
		}},
		{as: AVFNMADD231SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBD,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBD,
		}},
		{as: AVFNMADD231SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBD,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBD,
		}},
		{as: AVFNMSUB132PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x9E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9E,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x9E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x9E,
		}},
		{as: AVFNMSUB132PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x9E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9E,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x9E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x9E,
		}},
		{as: AVFNMSUB132SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x9F,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x9F,
		}},
		{as: AVFNMSUB132SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x9F,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x9F,
		}},
		{as: AVFNMSUB213PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAE,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xAE,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAE,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xAE,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xAE,
		}},
		{as: AVFNMSUB213PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAE,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xAE,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAE,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xAE,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xAE,
		}},
		{as: AVFNMSUB213SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xAF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xAF,
		}},
		{as: AVFNMSUB213SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xAF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xAF,
		}},
		{as: AVFNMSUB231PD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBE,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0xBE,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBE,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xBE,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xBE,
		}},
		{as: AVFNMSUB231PS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBE,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xBE,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBE,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xBE,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xBE,
		}},
		{as: AVFNMSUB231SD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0xBF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0xBF,
		}},
		{as: AVFNMSUB231SS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xBF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0xBF,
		}},
		{as: AVFPCLASSPDX, ytab: psess._yvfpclasspdx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8, 0x66,
		}},
		{as: AVFPCLASSPDY, ytab: psess._yvfpclasspdy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8, 0x66,
		}},
		{as: AVFPCLASSPDZ, ytab: psess._yvfpclasspdz, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8, 0x66,
		}},
		{as: AVFPCLASSPSX, ytab: psess._yvfpclasspdx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4, 0x66,
		}},
		{as: AVFPCLASSPSY, ytab: psess._yvfpclasspdy, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4, 0x66,
		}},
		{as: AVFPCLASSPSZ, ytab: psess._yvfpclasspdz, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4, 0x66,
		}},
		{as: AVFPCLASSSD, ytab: psess._yvfpclasspdx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8, 0x67,
		}},
		{as: AVFPCLASSSS, ytab: psess._yvfpclasspdx, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4, 0x67,
		}},
		{as: AVGATHERDPD, ytab: psess._yvgatherdpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x92,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x92,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0x92,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0x92,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0x92,
		}},
		{as: AVGATHERDPS, ytab: psess._yvgatherdps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x92,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x92,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0x92,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0x92,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0x92,
		}},
		{as: AVGATHERPF0DPD, ytab: psess._yvgatherpf0dpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC6, 01,
		}},
		{as: AVGATHERPF0DPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC6, 01,
		}},
		{as: AVGATHERPF0QPD, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC7, 01,
		}},
		{as: AVGATHERPF0QPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC7, 01,
		}},
		{as: AVGATHERPF1DPD, ytab: psess._yvgatherpf0dpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC6, 02,
		}},
		{as: AVGATHERPF1DPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC6, 02,
		}},
		{as: AVGATHERPF1QPD, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC7, 02,
		}},
		{as: AVGATHERPF1QPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC7, 02,
		}},
		{as: AVGATHERQPD, ytab: psess._yvgatherdps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x93,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x93,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0x93,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0x93,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0x93,
		}},
		{as: AVGATHERQPS, ytab: psess._yvgatherqps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x93,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x93,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0x93,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0x93,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0x93,
		}},
		{as: AVGETEXPPD, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x42,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x42,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x42,
		}},
		{as: AVGETEXPPS, ytab: psess._yvcvtpd2qq, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x42,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x42,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x42,
		}},
		{as: AVGETEXPSD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x43,
		}},
		{as: AVGETEXPSS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x43,
		}},
		{as: AVGETMANTPD, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x26,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x26,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x26,
		}},
		{as: AVGETMANTPS, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x26,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x26,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x26,
		}},
		{as: AVGETMANTSD, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x27,
		}},
		{as: AVGETMANTSS, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x27,
		}},
		{as: AVGF2P8AFFINEINVQB, ytab: psess._yvgf2p8affineinvqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0xCF,
			avxEscape | vex256 | vex66 | vex0F3A | vexW1, 0xCF,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xCF,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xCF,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xCF,
		}},
		{as: AVGF2P8AFFINEQB, ytab: psess._yvgf2p8affineinvqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0xCE,
			avxEscape | vex256 | vex66 | vex0F3A | vexW1, 0xCE,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xCE,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xCE,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xCE,
		}},
		{as: AVGF2P8MULB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0xCF,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0xCF,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0xCF,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0xCF,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0xCF,
		}},
		{as: AVHADDPD, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x7C,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x7C,
		}},
		{as: AVHADDPS, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x7C,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0x7C,
		}},
		{as: AVHSUBPD, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x7D,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x7D,
		}},
		{as: AVHSUBPS, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x7D,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0x7D,
		}},
		{as: AVINSERTF128, ytab: psess._yvinsertf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x18,
		}},
		{as: AVINSERTF32X4, ytab: psess._yvinsertf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x18,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x18,
		}},
		{as: AVINSERTF32X8, ytab: psess._yvinsertf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x1A,
		}},
		{as: AVINSERTF64X2, ytab: psess._yvinsertf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x18,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x18,
		}},
		{as: AVINSERTF64X4, ytab: psess._yvinsertf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x1A,
		}},
		{as: AVINSERTI128, ytab: psess._yvinsertf128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x38,
		}},
		{as: AVINSERTI32X4, ytab: psess._yvinsertf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x38,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x38,
		}},
		{as: AVINSERTI32X8, ytab: psess._yvinsertf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x3A,
		}},
		{as: AVINSERTI64X2, ytab: psess._yvinsertf32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x38,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x38,
		}},
		{as: AVINSERTI64X4, ytab: psess._yvinsertf32x8, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x3A,
		}},
		{as: AVINSERTPS, ytab: psess._yvinsertps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x21,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4, 0x21,
		}},
		{as: AVLDDQU, ytab: psess._yvlddqu, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0xF0,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0xF0,
		}},
		{as: AVLDMXCSR, ytab: psess._yvldmxcsr, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0xAE, 02,
		}},
		{as: AVMASKMOVDQU, ytab: psess._yvmaskmovdqu, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF7,
		}},
		{as: AVMASKMOVPD, ytab: psess._yvmaskmovpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2F,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2F,
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2D,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2D,
		}},
		{as: AVMASKMOVPS, ytab: psess._yvmaskmovpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2E,
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2C,
		}},
		{as: AVMAXPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5F,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5F,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x5F,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5F,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5F,
		}},
		{as: AVMAXPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5F,
			avxEscape | vex256 | vex0F | vexW0, 0x5F,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x5F,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5F,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5F,
		}},
		{as: AVMAXSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x5F,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x5F,
		}},
		{as: AVMAXSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5F,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x5F,
		}},
		{as: AVMINPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5D,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5D,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x5D,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5D,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5D,
		}},
		{as: AVMINPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5D,
			avxEscape | vex256 | vex0F | vexW0, 0x5D,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x5D,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5D,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5D,
		}},
		{as: AVMINSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x5D,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x5D,
		}},
		{as: AVMINSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5D,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x5D,
		}},
		{as: AVMOVAPD, ytab: psess._yvmovapd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x29,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x29,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x28,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x28,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x29,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x29,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x29,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x28,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x28,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x28,
		}},
		{as: AVMOVAPS, ytab: psess._yvmovapd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x29,
			avxEscape | vex256 | vex0F | vexW0, 0x29,
			avxEscape | vex128 | vex0F | vexW0, 0x28,
			avxEscape | vex256 | vex0F | vexW0, 0x28,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x29,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x29,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x29,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x28,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x28,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x28,
		}},
		{as: AVMOVD, ytab: psess._yvmovd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x7E,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6E,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN4, 0x7E,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN4, 0x6E,
		}},
		{as: AVMOVDDUP, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x12,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0x12,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexZeroingEnabled, 0x12,
			avxEscape | evex256 | evexF2 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x12,
			avxEscape | evex512 | evexF2 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x12,
		}},
		{as: AVMOVDQA, ytab: psess._yvmovdqa, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x7F,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x7F,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6F,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x6F,
		}},
		{as: AVMOVDQA32, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVDQA64, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVDQU, ytab: psess._yvmovdqa, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x7F,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x7F,
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x6F,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x6F,
		}},
		{as: AVMOVDQU16, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evexF2 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evexF2 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evexF2 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evexF2 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVDQU32, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVDQU64, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evexF3 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evexF3 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evexF3 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evexF3 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVDQU8, ytab: psess._yvmovdqa32, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evexF2 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evexF2 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x7F,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x6F,
			avxEscape | evex256 | evexF2 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x6F,
			avxEscape | evex512 | evexF2 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x6F,
		}},
		{as: AVMOVHLPS, ytab: psess._yvmovhlps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x12,
			avxEscape | evex128 | evex0F | evexW0, 0, 0x12,
		}},
		{as: AVMOVHPD, ytab: psess._yvmovhpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x17,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x16,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x17,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x16,
		}},
		{as: AVMOVHPS, ytab: psess._yvmovhpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x17,
			avxEscape | vex128 | vex0F | vexW0, 0x16,
			avxEscape | evex128 | evex0F | evexW0, evexN8, 0x17,
			avxEscape | evex128 | evex0F | evexW0, evexN8, 0x16,
		}},
		{as: AVMOVLHPS, ytab: psess._yvmovhlps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x16,
			avxEscape | evex128 | evex0F | evexW0, 0, 0x16,
		}},
		{as: AVMOVLPD, ytab: psess._yvmovhpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x13,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x12,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x13,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x12,
		}},
		{as: AVMOVLPS, ytab: psess._yvmovhpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x13,
			avxEscape | vex128 | vex0F | vexW0, 0x12,
			avxEscape | evex128 | evex0F | evexW0, evexN8, 0x13,
			avxEscape | evex128 | evex0F | evexW0, evexN8, 0x12,
		}},
		{as: AVMOVMSKPD, ytab: psess._yvmovmskpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x50,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x50,
		}},
		{as: AVMOVMSKPS, ytab: psess._yvmovmskpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x50,
			avxEscape | vex256 | vex0F | vexW0, 0x50,
		}},
		{as: AVMOVNTDQ, ytab: psess._yvmovntdq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE7,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE7,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0xE7,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0xE7,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0xE7,
		}},
		{as: AVMOVNTDQA, ytab: psess._yvmovntdqa, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2A,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2A,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0x2A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0x2A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0x2A,
		}},
		{as: AVMOVNTPD, ytab: psess._yvmovntdq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x2B,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x2B,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16, 0x2B,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32, 0x2B,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64, 0x2B,
		}},
		{as: AVMOVNTPS, ytab: psess._yvmovntdq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x2B,
			avxEscape | vex256 | vex0F | vexW0, 0x2B,
			avxEscape | evex128 | evex0F | evexW0, evexN16, 0x2B,
			avxEscape | evex256 | evex0F | evexW0, evexN32, 0x2B,
			avxEscape | evex512 | evex0F | evexW0, evexN64, 0x2B,
		}},
		{as: AVMOVQ, ytab: psess._yvmovq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x7E,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD6,
			avxEscape | vex128 | vex66 | vex0F | vexW1, 0x6E,
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x7E,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x7E,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0xD6,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8, 0x6E,
			avxEscape | evex128 | evexF3 | evex0F | evexW1, evexN8, 0x7E,
		}},
		{as: AVMOVSD, ytab: psess._yvmovsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x10,
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x10,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexZeroingEnabled, 0x11,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8, 0x11,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexZeroingEnabled, 0x10,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexZeroingEnabled, 0x10,
		}},
		{as: AVMOVSHDUP, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x16,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x16,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x16,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x16,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x16,
		}},
		{as: AVMOVSLDUP, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x12,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x12,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x12,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x12,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x12,
		}},
		{as: AVMOVSS, ytab: psess._yvmovsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x10,
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x10,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexZeroingEnabled, 0x11,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4, 0x11,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexZeroingEnabled, 0x10,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexZeroingEnabled, 0x10,
		}},
		{as: AVMOVUPD, ytab: psess._yvmovapd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x11,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x10,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x10,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x11,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x11,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x11,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0x10,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexZeroingEnabled, 0x10,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexZeroingEnabled, 0x10,
		}},
		{as: AVMOVUPS, ytab: psess._yvmovapd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x11,
			avxEscape | vex256 | vex0F | vexW0, 0x11,
			avxEscape | vex128 | vex0F | vexW0, 0x10,
			avxEscape | vex256 | vex0F | vexW0, 0x10,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x11,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x11,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x11,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x10,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x10,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x10,
		}},
		{as: AVMPSADBW, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x42,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x42,
		}},
		{as: AVMULPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x59,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x59,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x59,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x59,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x59,
		}},
		{as: AVMULPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x59,
			avxEscape | vex256 | vex0F | vexW0, 0x59,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x59,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x59,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x59,
		}},
		{as: AVMULSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x59,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x59,
		}},
		{as: AVMULSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x59,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x59,
		}},
		{as: AVORPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x56,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x56,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x56,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x56,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x56,
		}},
		{as: AVORPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x56,
			avxEscape | vex256 | vex0F | vexW0, 0x56,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x56,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x56,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x56,
		}},
		{as: AVP4DPWSSD, ytab: psess._yv4fmaddps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x52,
		}},
		{as: AVP4DPWSSDS, ytab: psess._yv4fmaddps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evexF2 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x53,
		}},
		{as: AVPABSB, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x1C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x1C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x1C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x1C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x1C,
		}},
		{as: AVPABSD, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x1E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x1E,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x1E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x1E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x1E,
		}},
		{as: AVPABSQ, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x1F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x1F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x1F,
		}},
		{as: AVPABSW, ytab: psess._yvmovddup, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x1D,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x1D,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x1D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x1D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x1D,
		}},
		{as: AVPACKSSDW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6B,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x6B,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x6B,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x6B,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x6B,
		}},
		{as: AVPACKSSWB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x63,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x63,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x63,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x63,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x63,
		}},
		{as: AVPACKUSDW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x2B,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x2B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x2B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x2B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x2B,
		}},
		{as: AVPACKUSWB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x67,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x67,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x67,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x67,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x67,
		}},
		{as: AVPADDB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xFC,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xFC,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xFC,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xFC,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xFC,
		}},
		{as: AVPADDD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xFE,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xFE,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xFE,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xFE,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xFE,
		}},
		{as: AVPADDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD4,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD4,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xD4,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xD4,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xD4,
		}},
		{as: AVPADDSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xEC,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xEC,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xEC,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xEC,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xEC,
		}},
		{as: AVPADDSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xED,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xED,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xED,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xED,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xED,
		}},
		{as: AVPADDUSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDC,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDC,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xDC,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xDC,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xDC,
		}},
		{as: AVPADDUSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDD,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDD,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xDD,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xDD,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xDD,
		}},
		{as: AVPADDW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xFD,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xFD,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xFD,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xFD,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xFD,
		}},
		{as: AVPALIGNR, ytab: psess._yvgf2p8affineinvqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0F,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x0F,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexZeroingEnabled, 0x0F,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexZeroingEnabled, 0x0F,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexZeroingEnabled, 0x0F,
		}},
		{as: AVPAND, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDB,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDB,
		}},
		{as: AVPANDD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xDB,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xDB,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xDB,
		}},
		{as: AVPANDN, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDF,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDF,
		}},
		{as: AVPANDND, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xDF,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xDF,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xDF,
		}},
		{as: AVPANDNQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xDF,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xDF,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xDF,
		}},
		{as: AVPANDQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xDB,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xDB,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xDB,
		}},
		{as: AVPAVGB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE0,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE0,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE0,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE0,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE0,
		}},
		{as: AVPAVGW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE3,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE3,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE3,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE3,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE3,
		}},
		{as: AVPBLENDD, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x02,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x02,
		}},
		{as: AVPBLENDMB, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x66,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x66,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x66,
		}},
		{as: AVPBLENDMD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x64,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x64,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x64,
		}},
		{as: AVPBLENDMQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x64,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x64,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x64,
		}},
		{as: AVPBLENDMW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x66,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x66,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x66,
		}},
		{as: AVPBLENDVB, ytab: psess._yvblendvpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x4C,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x4C,
		}},
		{as: AVPBLENDW, ytab: psess._yvblendpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0E,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x0E,
		}},
		{as: AVPBROADCASTB, ytab: psess._yvpbroadcastb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x78,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x78,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7A,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x78,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x78,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x78,
		}},
		{as: AVPBROADCASTD, ytab: psess._yvpbroadcastb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x58,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x58,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x58,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x58,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x58,
		}},
		{as: AVPBROADCASTMB2Q, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, 0, 0x2A,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, 0, 0x2A,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, 0, 0x2A,
		}},
		{as: AVPBROADCASTMW2D, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, 0, 0x3A,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, 0, 0x3A,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, 0, 0x3A,
		}},
		{as: AVPBROADCASTQ, ytab: psess._yvpbroadcastb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x59,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x59,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexZeroingEnabled, 0x7C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexZeroingEnabled, 0x7C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexZeroingEnabled, 0x7C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x59,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x59,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x59,
		}},
		{as: AVPBROADCASTW, ytab: psess._yvpbroadcastb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x79,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x79,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexZeroingEnabled, 0x7B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x79,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x79,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x79,
		}},
		{as: AVPCLMULQDQ, ytab: psess._yvpclmulqdq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x44,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x44,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16, 0x44,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32, 0x44,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64, 0x44,
		}},
		{as: AVPCMPB, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16, 0x3F,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32, 0x3F,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64, 0x3F,
		}},
		{as: AVPCMPD, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4, 0x1F,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4, 0x1F,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4, 0x1F,
		}},
		{as: AVPCMPEQB, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x74,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x74,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x74,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x74,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x74,
		}},
		{as: AVPCMPEQD, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x76,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x76,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4, 0x76,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4, 0x76,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4, 0x76,
		}},
		{as: AVPCMPEQQ, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x29,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x29,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8, 0x29,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8, 0x29,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8, 0x29,
		}},
		{as: AVPCMPEQW, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x75,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x75,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x75,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x75,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x75,
		}},
		{as: AVPCMPESTRI, ytab: psess._yvaeskeygenassist, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexWIG, 0x61,
		}},
		{as: AVPCMPESTRM, ytab: psess._yvaeskeygenassist, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexWIG, 0x60,
		}},
		{as: AVPCMPGTB, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x64,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x64,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x64,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x64,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x64,
		}},
		{as: AVPCMPGTD, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x66,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x66,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4, 0x66,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4, 0x66,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4, 0x66,
		}},
		{as: AVPCMPGTQ, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x37,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x37,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8, 0x37,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8, 0x37,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8, 0x37,
		}},
		{as: AVPCMPGTW, ytab: psess._yvpcmpeqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x65,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x65,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x65,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x65,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x65,
		}},
		{as: AVPCMPISTRI, ytab: psess._yvaeskeygenassist, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexWIG, 0x63,
		}},
		{as: AVPCMPISTRM, ytab: psess._yvaeskeygenassist, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x62,
		}},
		{as: AVPCMPQ, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8, 0x1F,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8, 0x1F,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8, 0x1F,
		}},
		{as: AVPCMPUB, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16, 0x3E,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32, 0x3E,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64, 0x3E,
		}},
		{as: AVPCMPUD, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4, 0x1E,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4, 0x1E,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4, 0x1E,
		}},
		{as: AVPCMPUQ, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8, 0x1E,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8, 0x1E,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8, 0x1E,
		}},
		{as: AVPCMPUW, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16, 0x3E,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32, 0x3E,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64, 0x3E,
		}},
		{as: AVPCMPW, ytab: psess._yvpcmpb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16, 0x3F,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32, 0x3F,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64, 0x3F,
		}},
		{as: AVPCOMPRESSB, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x63,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x63,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x63,
		}},
		{as: AVPCOMPRESSD, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x8B,
		}},
		{as: AVPCOMPRESSQ, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x8B,
		}},
		{as: AVPCOMPRESSW, ytab: psess._yvcompresspd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x63,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x63,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x63,
		}},
		{as: AVPCONFLICTD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xC4,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xC4,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xC4,
		}},
		{as: AVPCONFLICTQ, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xC4,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xC4,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xC4,
		}},
		{as: AVPDPBUSD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x50,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x50,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x50,
		}},
		{as: AVPDPBUSDS, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x51,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x51,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x51,
		}},
		{as: AVPDPWSSD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x52,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x52,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x52,
		}},
		{as: AVPDPWSSDS, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x53,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x53,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x53,
		}},
		{as: AVPERM2F128, ytab: psess._yvperm2f128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x06,
		}},
		{as: AVPERM2I128, ytab: psess._yvperm2f128, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x46,
		}},
		{as: AVPERMB, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x8D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x8D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x8D,
		}},
		{as: AVPERMD, ytab: psess._yvpermd, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x36,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x36,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x36,
		}},
		{as: AVPERMI2B, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x75,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x75,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x75,
		}},
		{as: AVPERMI2D, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x76,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x76,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x76,
		}},
		{as: AVPERMI2PD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x77,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x77,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x77,
		}},
		{as: AVPERMI2PS, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x77,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x77,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x77,
		}},
		{as: AVPERMI2Q, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x76,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x76,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x76,
		}},
		{as: AVPERMI2W, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x75,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x75,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x75,
		}},
		{as: AVPERMILPD, ytab: psess._yvpermilpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x05,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x05,
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0D,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0D,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x05,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x05,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x05,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x0D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x0D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x0D,
		}},
		{as: AVPERMILPS, ytab: psess._yvpermilpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x04,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x04,
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0C,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x04,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x04,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x04,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x0C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x0C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x0C,
		}},
		{as: AVPERMPD, ytab: psess._yvpermq, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW1, 0x01,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x01,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x01,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x16,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x16,
		}},
		{as: AVPERMPS, ytab: psess._yvpermd, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x16,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x16,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x16,
		}},
		{as: AVPERMQ, ytab: psess._yvpermq, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex66 | vex0F3A | vexW1, 0x00,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x00,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x00,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x36,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x36,
		}},
		{as: AVPERMT2B, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x7D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x7D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x7D,
		}},
		{as: AVPERMT2D, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x7E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x7E,
		}},
		{as: AVPERMT2PD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x7F,
		}},
		{as: AVPERMT2PS, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x7F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x7F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x7F,
		}},
		{as: AVPERMT2Q, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x7E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x7E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x7E,
		}},
		{as: AVPERMT2W, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x7D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x7D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x7D,
		}},
		{as: AVPERMW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x8D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x8D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x8D,
		}},
		{as: AVPEXPANDB, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x62,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x62,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN1 | evexZeroingEnabled, 0x62,
		}},
		{as: AVPEXPANDD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x89,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x89,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x89,
		}},
		{as: AVPEXPANDQ, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x89,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x89,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x89,
		}},
		{as: AVPEXPANDW, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x62,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x62,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN2 | evexZeroingEnabled, 0x62,
		}},
		{as: AVPEXTRB, ytab: psess._yvextractps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x14,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN1, 0x14,
		}},
		{as: AVPEXTRD, ytab: psess._yvextractps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x16,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4, 0x16,
		}},
		{as: AVPEXTRQ, ytab: psess._yvextractps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x16,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8, 0x16,
		}},
		{as: AVPEXTRW, ytab: psess._yvpextrw, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x15,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xC5,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN2, 0x15,
			avxEscape | evex128 | evex66 | evex0F | evexW0, 0, 0xC5,
		}},
		{as: AVPGATHERDD, ytab: psess._yvgatherdps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x90,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x90,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0x90,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0x90,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0x90,
		}},
		{as: AVPGATHERDQ, ytab: psess._yvgatherdpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x90,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x90,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0x90,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0x90,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0x90,
		}},
		{as: AVPGATHERQD, ytab: psess._yvgatherqps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x91,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x91,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0x91,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0x91,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0x91,
		}},
		{as: AVPGATHERQQ, ytab: psess._yvgatherdps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x91,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x91,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0x91,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0x91,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0x91,
		}},
		{as: AVPHADDD, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x02,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x02,
		}},
		{as: AVPHADDSW, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x03,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x03,
		}},
		{as: AVPHADDW, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x01,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x01,
		}},
		{as: AVPHMINPOSUW, ytab: psess._yvaesimc, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x41,
		}},
		{as: AVPHSUBD, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x06,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x06,
		}},
		{as: AVPHSUBSW, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x07,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x07,
		}},
		{as: AVPHSUBW, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x05,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x05,
		}},
		{as: AVPINSRB, ytab: psess._yvpinsrb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x20,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN1, 0x20,
		}},
		{as: AVPINSRD, ytab: psess._yvpinsrb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x22,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4, 0x22,
		}},
		{as: AVPINSRQ, ytab: psess._yvpinsrb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW1, 0x22,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8, 0x22,
		}},
		{as: AVPINSRW, ytab: psess._yvpinsrb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xC4,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN2, 0xC4,
		}},
		{as: AVPLZCNTD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x44,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x44,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x44,
		}},
		{as: AVPLZCNTQ, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x44,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x44,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x44,
		}},
		{as: AVPMADD52HUQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xB5,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xB5,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xB5,
		}},
		{as: AVPMADD52LUQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xB4,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xB4,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xB4,
		}},
		{as: AVPMADDUBSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x04,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x04,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x04,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x04,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x04,
		}},
		{as: AVPMADDWD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF5,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF5,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF5,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xF5,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xF5,
		}},
		{as: AVPMASKMOVD, ytab: psess._yvmaskmovpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x8E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x8E,
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x8C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x8C,
		}},
		{as: AVPMASKMOVQ, ytab: psess._yvmaskmovpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x8E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x8E,
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x8C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x8C,
		}},
		{as: AVPMAXSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3C,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x3C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x3C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x3C,
		}},
		{as: AVPMAXSD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3D,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3D,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x3D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x3D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x3D,
		}},
		{as: AVPMAXSQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x3D,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x3D,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x3D,
		}},
		{as: AVPMAXSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xEE,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xEE,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xEE,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xEE,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xEE,
		}},
		{as: AVPMAXUB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDE,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDE,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xDE,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xDE,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xDE,
		}},
		{as: AVPMAXUD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3F,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3F,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x3F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x3F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x3F,
		}},
		{as: AVPMAXUQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x3F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x3F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x3F,
		}},
		{as: AVPMAXUW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3E,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x3E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x3E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x3E,
		}},
		{as: AVPMINSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x38,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x38,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x38,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x38,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x38,
		}},
		{as: AVPMINSD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x39,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x39,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x39,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x39,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x39,
		}},
		{as: AVPMINSQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x39,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x39,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x39,
		}},
		{as: AVPMINSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xEA,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xEA,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xEA,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xEA,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xEA,
		}},
		{as: AVPMINUB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xDA,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xDA,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xDA,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xDA,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xDA,
		}},
		{as: AVPMINUD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3B,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x3B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x3B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x3B,
		}},
		{as: AVPMINUQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x3B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x3B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x3B,
		}},
		{as: AVPMINUW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x3A,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x3A,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x3A,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x3A,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x3A,
		}},
		{as: AVPMOVB2M, ytab: psess._yvpmovb2m, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, 0, 0x29,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, 0, 0x29,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, 0, 0x29,
		}},
		{as: AVPMOVD2M, ytab: psess._yvpmovb2m, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, 0, 0x39,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, 0, 0x39,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, 0, 0x39,
		}},
		{as: AVPMOVDB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x31,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x31,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x31,
		}},
		{as: AVPMOVDW, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x33,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x33,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x33,
		}},
		{as: AVPMOVM2B, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, 0, 0x28,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, 0, 0x28,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, 0, 0x28,
		}},
		{as: AVPMOVM2D, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, 0, 0x38,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, 0, 0x38,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, 0, 0x38,
		}},
		{as: AVPMOVM2Q, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, 0, 0x38,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, 0, 0x38,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, 0, 0x38,
		}},
		{as: AVPMOVM2W, ytab: psess._yvpbroadcastmb2q, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, 0, 0x28,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, 0, 0x28,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, 0, 0x28,
		}},
		{as: AVPMOVMSKB, ytab: psess._yvmovmskpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD7,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD7,
		}},
		{as: AVPMOVQ2M, ytab: psess._yvpmovb2m, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, 0, 0x39,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, 0, 0x39,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, 0, 0x39,
		}},
		{as: AVPMOVQB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x32,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x32,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x32,
		}},
		{as: AVPMOVQD, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x35,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x35,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x35,
		}},
		{as: AVPMOVQW, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x34,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x34,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x34,
		}},
		{as: AVPMOVSDB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x21,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x21,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x21,
		}},
		{as: AVPMOVSDW, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x23,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x23,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x23,
		}},
		{as: AVPMOVSQB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x22,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x22,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x22,
		}},
		{as: AVPMOVSQD, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x25,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x25,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x25,
		}},
		{as: AVPMOVSQW, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x24,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x24,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x24,
		}},
		{as: AVPMOVSWB, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x20,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x20,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x20,
		}},
		{as: AVPMOVSXBD, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x21,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x21,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x21,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x21,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x21,
		}},
		{as: AVPMOVSXBQ, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x22,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x22,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x22,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x22,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x22,
		}},
		{as: AVPMOVSXBW, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x20,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x20,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x20,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x20,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x20,
		}},
		{as: AVPMOVSXDQ, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x25,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x25,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x25,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x25,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x25,
		}},
		{as: AVPMOVSXWD, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x23,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x23,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x23,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x23,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x23,
		}},
		{as: AVPMOVSXWQ, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x24,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x24,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x24,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x24,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x24,
		}},
		{as: AVPMOVUSDB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x11,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x11,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x11,
		}},
		{as: AVPMOVUSDW, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x13,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x13,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x13,
		}},
		{as: AVPMOVUSQB, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x12,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x12,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x12,
		}},
		{as: AVPMOVUSQD, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x15,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x15,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x15,
		}},
		{as: AVPMOVUSQW, ytab: psess._yvpmovdb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x14,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x14,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x14,
		}},
		{as: AVPMOVUSWB, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x10,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x10,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x10,
		}},
		{as: AVPMOVW2M, ytab: psess._yvpmovb2m, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, 0, 0x29,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, 0, 0x29,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, 0, 0x29,
		}},
		{as: AVPMOVWB, ytab: psess._yvpmovdw, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x30,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x30,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x30,
		}},
		{as: AVPMOVZXBD, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x31,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x31,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x31,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x31,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x31,
		}},
		{as: AVPMOVZXBQ, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x32,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x32,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN2 | evexZeroingEnabled, 0x32,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x32,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x32,
		}},
		{as: AVPMOVZXBW, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x30,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x30,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x30,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x30,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x30,
		}},
		{as: AVPMOVZXDQ, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x35,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x35,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x35,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x35,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x35,
		}},
		{as: AVPMOVZXWD, ytab: psess._yvcvtdq2pd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x33,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x33,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x33,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x33,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x33,
		}},
		{as: AVPMOVZXWQ, ytab: psess._yvbroadcastss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x34,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x34,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x34,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN8 | evexZeroingEnabled, 0x34,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x34,
		}},
		{as: AVPMULDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x28,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x28,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x28,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x28,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x28,
		}},
		{as: AVPMULHRSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0B,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0B,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x0B,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x0B,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x0B,
		}},
		{as: AVPMULHUW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE4,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE4,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE4,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE4,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE4,
		}},
		{as: AVPMULHW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE5,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE5,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE5,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE5,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE5,
		}},
		{as: AVPMULLD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x40,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x40,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x40,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x40,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x40,
		}},
		{as: AVPMULLQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x40,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x40,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x40,
		}},
		{as: AVPMULLW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD5,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD5,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD5,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xD5,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xD5,
		}},
		{as: AVPMULTISHIFTQB, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x83,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x83,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x83,
		}},
		{as: AVPMULUDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF4,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF4,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xF4,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xF4,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xF4,
		}},
		{as: AVPOPCNTB, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x54,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x54,
		}},
		{as: AVPOPCNTD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x55,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x55,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x55,
		}},
		{as: AVPOPCNTQ, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x55,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x55,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x55,
		}},
		{as: AVPOPCNTW, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x54,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x54,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x54,
		}},
		{as: AVPOR, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xEB,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xEB,
		}},
		{as: AVPORD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xEB,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xEB,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xEB,
		}},
		{as: AVPORQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xEB,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xEB,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xEB,
		}},
		{as: AVPROLD, ytab: psess._yvprold, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x72, 01,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x72, 01,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x72, 01,
		}},
		{as: AVPROLQ, ytab: psess._yvprold, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x72, 01,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x72, 01,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x72, 01,
		}},
		{as: AVPROLVD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x15,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x15,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x15,
		}},
		{as: AVPROLVQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x15,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x15,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x15,
		}},
		{as: AVPRORD, ytab: psess._yvprold, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x72, 00,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x72, 00,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x72, 00,
		}},
		{as: AVPRORQ, ytab: psess._yvprold, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x72, 00,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x72, 00,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x72, 00,
		}},
		{as: AVPRORVD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x14,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x14,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x14,
		}},
		{as: AVPRORVQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x14,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x14,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x14,
		}},
		{as: AVPSADBW, ytab: psess._yvaesdec, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF6,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF6,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0xF6,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0xF6,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0xF6,
		}},
		{as: AVPSCATTERDD, ytab: psess._yvpscatterdd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0xA0,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0xA0,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xA0,
		}},
		{as: AVPSCATTERDQ, ytab: psess._yvpscatterdq, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0xA0,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0xA0,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xA0,
		}},
		{as: AVPSCATTERQD, ytab: psess._yvpscatterqd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0xA1,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0xA1,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xA1,
		}},
		{as: AVPSCATTERQQ, ytab: psess._yvpscatterdd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0xA1,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0xA1,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xA1,
		}},
		{as: AVPSHLDD, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x71,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x71,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x71,
		}},
		{as: AVPSHLDQ, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x71,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x71,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x71,
		}},
		{as: AVPSHLDVD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x71,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x71,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x71,
		}},
		{as: AVPSHLDVQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x71,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x71,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x71,
		}},
		{as: AVPSHLDVW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x70,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x70,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x70,
		}},
		{as: AVPSHLDW, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x70,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x70,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexZeroingEnabled, 0x70,
		}},
		{as: AVPSHRDD, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x73,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x73,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x73,
		}},
		{as: AVPSHRDQ, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x73,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x73,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x73,
		}},
		{as: AVPSHRDVD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x73,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x73,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x73,
		}},
		{as: AVPSHRDVQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x73,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x73,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x73,
		}},
		{as: AVPSHRDVW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x72,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x72,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x72,
		}},
		{as: AVPSHRDW, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexZeroingEnabled, 0x72,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexZeroingEnabled, 0x72,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexZeroingEnabled, 0x72,
		}},
		{as: AVPSHUFB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x00,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x00,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexZeroingEnabled, 0x00,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexZeroingEnabled, 0x00,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexZeroingEnabled, 0x00,
		}},
		{as: AVPSHUFBITQMB, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0x8F,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0x8F,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0x8F,
		}},
		{as: AVPSHUFD, ytab: psess._yvpshufd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x70,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x70,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x70,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x70,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x70,
		}},
		{as: AVPSHUFHW, ytab: psess._yvpshufd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x70,
			avxEscape | vex256 | vexF3 | vex0F | vexW0, 0x70,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x70,
			avxEscape | evex256 | evexF3 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x70,
			avxEscape | evex512 | evexF3 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x70,
		}},
		{as: AVPSHUFLW, ytab: psess._yvpshufd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x70,
			avxEscape | vex256 | vexF2 | vex0F | vexW0, 0x70,
			avxEscape | evex128 | evexF2 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x70,
			avxEscape | evex256 | evexF2 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x70,
			avxEscape | evex512 | evexF2 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x70,
		}},
		{as: AVPSIGNB, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x08,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x08,
		}},
		{as: AVPSIGND, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0A,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0A,
		}},
		{as: AVPSIGNW, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x09,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x09,
		}},
		{as: AVPSLLD, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x72, 06,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x72, 06,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF2,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF2,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x72, 06,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x72, 06,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x72, 06,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF2,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF2,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF2,
		}},
		{as: AVPSLLDQ, ytab: psess._yvpslldq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x73, 07,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x73, 07,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x73, 07,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x73, 07,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x73, 07,
		}},
		{as: AVPSLLQ, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x73, 06,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x73, 06,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF3,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF3,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x73, 06,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x73, 06,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x73, 06,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xF3,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xF3,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xF3,
		}},
		{as: AVPSLLVD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x47,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x47,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x47,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x47,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x47,
		}},
		{as: AVPSLLVQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x47,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x47,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x47,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x47,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x47,
		}},
		{as: AVPSLLVW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x12,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x12,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x12,
		}},
		{as: AVPSLLW, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x71, 06,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x71, 06,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF1,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF1,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x71, 06,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x71, 06,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x71, 06,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF1,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF1,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF1,
		}},
		{as: AVPSRAD, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x72, 04,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x72, 04,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE2,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE2,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE2,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE2,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE2,
		}},
		{as: AVPSRAQ, ytab: psess._yvpsraq, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x72, 04,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xE2,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xE2,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xE2,
		}},
		{as: AVPSRAVD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x46,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x46,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x46,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x46,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x46,
		}},
		{as: AVPSRAVQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x46,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x46,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x46,
		}},
		{as: AVPSRAVW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x11,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x11,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x11,
		}},
		{as: AVPSRAW, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x71, 04,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x71, 04,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE1,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE1,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x71, 04,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x71, 04,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x71, 04,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE1,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE1,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE1,
		}},
		{as: AVPSRLD, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x72, 02,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x72, 02,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD2,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD2,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x72, 02,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x72, 02,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x72, 02,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD2,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD2,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD2,
		}},
		{as: AVPSRLDQ, ytab: psess._yvpslldq, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x73, 03,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x73, 03,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16, 0x73, 03,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32, 0x73, 03,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64, 0x73, 03,
		}},
		{as: AVPSRLQ, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x73, 02,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x73, 02,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD3,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD3,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x73, 02,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x73, 02,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x73, 02,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xD3,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xD3,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN16 | evexZeroingEnabled, 0xD3,
		}},
		{as: AVPSRLVD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x45,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x45,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x45,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x45,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x45,
		}},
		{as: AVPSRLVQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW1, 0x45,
			avxEscape | vex256 | vex66 | vex0F38 | vexW1, 0x45,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x45,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x45,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x45,
		}},
		{as: AVPSRLVW, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexZeroingEnabled, 0x10,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexZeroingEnabled, 0x10,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexZeroingEnabled, 0x10,
		}},
		{as: AVPSRLW, ytab: psess._yvpslld, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x71, 02,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x71, 02,
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD1,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD1,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x71, 02,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x71, 02,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x71, 02,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD1,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD1,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD1,
		}},
		{as: AVPSUBB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF8,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF8,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF8,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xF8,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xF8,
		}},
		{as: AVPSUBD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xFA,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xFA,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xFA,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xFA,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xFA,
		}},
		{as: AVPSUBQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xFB,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xFB,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xFB,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xFB,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xFB,
		}},
		{as: AVPSUBSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE8,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE8,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE8,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE8,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE8,
		}},
		{as: AVPSUBSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xE9,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xE9,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xE9,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xE9,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xE9,
		}},
		{as: AVPSUBUSB, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD8,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD8,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD8,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xD8,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xD8,
		}},
		{as: AVPSUBUSW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xD9,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xD9,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xD9,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xD9,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xD9,
		}},
		{as: AVPSUBW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xF9,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xF9,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0xF9,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0xF9,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0xF9,
		}},
		{as: AVPTERNLOGD, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x25,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x25,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x25,
		}},
		{as: AVPTERNLOGQ, ytab: psess._yvalignd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x25,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x25,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x25,
		}},
		{as: AVPTEST, ytab: psess._yvptest, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x17,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x17,
		}},
		{as: AVPTESTMB, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16, 0x26,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32, 0x26,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64, 0x26,
		}},
		{as: AVPTESTMD, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4, 0x27,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4, 0x27,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4, 0x27,
		}},
		{as: AVPTESTMQ, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8, 0x27,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8, 0x27,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8, 0x27,
		}},
		{as: AVPTESTMW, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16, 0x26,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32, 0x26,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64, 0x26,
		}},
		{as: AVPTESTNMB, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN16, 0x26,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN32, 0x26,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN64, 0x26,
		}},
		{as: AVPTESTNMD, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW0, evexN16 | evexBcstN4, 0x27,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW0, evexN32 | evexBcstN4, 0x27,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW0, evexN64 | evexBcstN4, 0x27,
		}},
		{as: AVPTESTNMQ, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, evexN16 | evexBcstN8, 0x27,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, evexN32 | evexBcstN8, 0x27,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, evexN64 | evexBcstN8, 0x27,
		}},
		{as: AVPTESTNMW, ytab: psess._yvpshufbitqmb, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evexF3 | evex0F38 | evexW1, evexN16, 0x26,
			avxEscape | evex256 | evexF3 | evex0F38 | evexW1, evexN32, 0x26,
			avxEscape | evex512 | evexF3 | evex0F38 | evexW1, evexN64, 0x26,
		}},
		{as: AVPUNPCKHBW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x68,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x68,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x68,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x68,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x68,
		}},
		{as: AVPUNPCKHDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6A,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x6A,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x6A,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x6A,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x6A,
		}},
		{as: AVPUNPCKHQDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6D,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x6D,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x6D,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x6D,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x6D,
		}},
		{as: AVPUNPCKHWD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x69,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x69,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x69,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x69,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x69,
		}},
		{as: AVPUNPCKLBW, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x60,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x60,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x60,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x60,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x60,
		}},
		{as: AVPUNPCKLDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x62,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x62,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x62,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x62,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x62,
		}},
		{as: AVPUNPCKLQDQ, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x6C,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x6C,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x6C,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x6C,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x6C,
		}},
		{as: AVPUNPCKLWD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x61,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x61,
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexZeroingEnabled, 0x61,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexZeroingEnabled, 0x61,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexZeroingEnabled, 0x61,
		}},
		{as: AVPXOR, ytab: psess._yvaddsubpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xEF,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xEF,
		}},
		{as: AVPXORD, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xEF,
			avxEscape | evex256 | evex66 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xEF,
			avxEscape | evex512 | evex66 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xEF,
		}},
		{as: AVPXORQ, ytab: psess._yvblendmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xEF,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xEF,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xEF,
		}},
		{as: AVRANGEPD, ytab: psess._yvfixupimmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x50,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x50,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x50,
		}},
		{as: AVRANGEPS, ytab: psess._yvfixupimmpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x50,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x50,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x50,
		}},
		{as: AVRANGESD, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x51,
		}},
		{as: AVRANGESS, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x51,
		}},
		{as: AVRCP14PD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x4C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x4C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x4C,
		}},
		{as: AVRCP14PS, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x4C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x4C,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x4C,
		}},
		{as: AVRCP14SD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x4D,
		}},
		{as: AVRCP14SS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x4D,
		}},
		{as: AVRCP28PD, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0xCA,
		}},
		{as: AVRCP28PS, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0xCA,
		}},
		{as: AVRCP28SD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0xCB,
		}},
		{as: AVRCP28SS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0xCB,
		}},
		{as: AVRCPPS, ytab: psess._yvptest, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x53,
			avxEscape | vex256 | vex0F | vexW0, 0x53,
		}},
		{as: AVRCPSS, ytab: psess._yvrcpss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x53,
		}},
		{as: AVREDUCEPD, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x56,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x56,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x56,
		}},
		{as: AVREDUCEPS, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x56,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x56,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x56,
		}},
		{as: AVREDUCESD, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x57,
		}},
		{as: AVREDUCESS, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x57,
		}},
		{as: AVRNDSCALEPD, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0x09,
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x09,
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x09,
		}},
		{as: AVRNDSCALEPS, ytab: psess._yvgetmantpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0x08,
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x08,
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x08,
		}},
		{as: AVRNDSCALESD, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0x0B,
		}},
		{as: AVRNDSCALESS, ytab: psess._yvfixupimmsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F3A | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0x0A,
		}},
		{as: AVROUNDPD, ytab: psess._yvroundpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x09,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x09,
		}},
		{as: AVROUNDPS, ytab: psess._yvroundpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x08,
			avxEscape | vex256 | vex66 | vex0F3A | vexW0, 0x08,
		}},
		{as: AVROUNDSD, ytab: psess._yvdppd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0B,
		}},
		{as: AVROUNDSS, ytab: psess._yvdppd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F3A | vexW0, 0x0A,
		}},
		{as: AVRSQRT14PD, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x4E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x4E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x4E,
		}},
		{as: AVRSQRT14PS, ytab: psess._yvexpandpd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x4E,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x4E,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x4E,
		}},
		{as: AVRSQRT14SD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexZeroingEnabled, 0x4F,
		}},
		{as: AVRSQRT14SS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexZeroingEnabled, 0x4F,
		}},
		{as: AVRSQRT28PD, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexSaeEnabled | evexZeroingEnabled, 0xCC,
		}},
		{as: AVRSQRT28PS, ytab: psess._yvexp2pd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexSaeEnabled | evexZeroingEnabled, 0xCC,
		}},
		{as: AVRSQRT28SD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexSaeEnabled | evexZeroingEnabled, 0xCD,
		}},
		{as: AVRSQRT28SS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexSaeEnabled | evexZeroingEnabled, 0xCD,
		}},
		{as: AVRSQRTPS, ytab: psess._yvptest, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x52,
			avxEscape | vex256 | vex0F | vexW0, 0x52,
		}},
		{as: AVRSQRTSS, ytab: psess._yvrcpss, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x52,
		}},
		{as: AVSCALEFPD, ytab: psess._yvscalefpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x2C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x2C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x2C,
		}},
		{as: AVSCALEFPS, ytab: psess._yvscalefpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x2C,
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x2C,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x2C,
		}},
		{as: AVSCALEFSD, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x2D,
		}},
		{as: AVSCALEFSS, ytab: psess._yvgetexpsd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x2D,
		}},
		{as: AVSCATTERDPD, ytab: psess._yvpscatterdq, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0xA2,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0xA2,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xA2,
		}},
		{as: AVSCATTERDPS, ytab: psess._yvpscatterdd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0xA2,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0xA2,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xA2,
		}},
		{as: AVSCATTERPF0DPD, ytab: psess._yvgatherpf0dpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC6, 05,
		}},
		{as: AVSCATTERPF0DPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC6, 05,
		}},
		{as: AVSCATTERPF0QPD, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC7, 05,
		}},
		{as: AVSCATTERPF0QPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC7, 05,
		}},
		{as: AVSCATTERPF1DPD, ytab: psess._yvgatherpf0dpd, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC6, 06,
		}},
		{as: AVSCATTERPF1DPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC6, 06,
		}},
		{as: AVSCATTERPF1QPD, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xC7, 06,
		}},
		{as: AVSCATTERPF1QPS, ytab: psess._yvgatherpf0dps, prefix: Pavx, op: opBytes{
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xC7, 06,
		}},
		{as: AVSCATTERQPD, ytab: psess._yvpscatterdd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW1, evexN8, 0xA3,
			avxEscape | evex256 | evex66 | evex0F38 | evexW1, evexN8, 0xA3,
			avxEscape | evex512 | evex66 | evex0F38 | evexW1, evexN8, 0xA3,
		}},
		{as: AVSCATTERQPS, ytab: psess._yvpscatterqd, prefix: Pavx, op: opBytes{
			avxEscape | evex128 | evex66 | evex0F38 | evexW0, evexN4, 0xA3,
			avxEscape | evex256 | evex66 | evex0F38 | evexW0, evexN4, 0xA3,
			avxEscape | evex512 | evex66 | evex0F38 | evexW0, evexN4, 0xA3,
		}},
		{as: AVSHUFF32X4, ytab: psess._yvshuff32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x23,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x23,
		}},
		{as: AVSHUFF64X2, ytab: psess._yvshuff32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x23,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x23,
		}},
		{as: AVSHUFI32X4, ytab: psess._yvshuff32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x43,
			avxEscape | evex512 | evex66 | evex0F3A | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x43,
		}},
		{as: AVSHUFI64X2, ytab: psess._yvshuff32x4, prefix: Pavx, op: opBytes{
			avxEscape | evex256 | evex66 | evex0F3A | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x43,
			avxEscape | evex512 | evex66 | evex0F3A | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x43,
		}},
		{as: AVSHUFPD, ytab: psess._yvgf2p8affineinvqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0xC6,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0xC6,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0xC6,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0xC6,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0xC6,
		}},
		{as: AVSHUFPS, ytab: psess._yvgf2p8affineinvqb, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0xC6,
			avxEscape | vex256 | vex0F | vexW0, 0xC6,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0xC6,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0xC6,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0xC6,
		}},
		{as: AVSQRTPD, ytab: psess._yvcvtdq2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x51,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x51,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x51,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x51,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x51,
		}},
		{as: AVSQRTPS, ytab: psess._yvcvtdq2ps, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x51,
			avxEscape | vex256 | vex0F | vexW0, 0x51,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x51,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x51,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x51,
		}},
		{as: AVSQRTSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x51,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x51,
		}},
		{as: AVSQRTSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x51,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x51,
		}},
		{as: AVSTMXCSR, ytab: psess._yvldmxcsr, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0xAE, 03,
		}},
		{as: AVSUBPD, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x5C,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x5C,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5C,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x5C,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x5C,
		}},
		{as: AVSUBPS, ytab: psess._yvaddpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x5C,
			avxEscape | vex256 | vex0F | vexW0, 0x5C,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5C,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x5C,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x5C,
		}},
		{as: AVSUBSD, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF2 | vex0F | vexW0, 0x5C,
			avxEscape | evex128 | evexF2 | evex0F | evexW1, evexN8 | evexRoundingEnabled | evexZeroingEnabled, 0x5C,
		}},
		{as: AVSUBSS, ytab: psess._yvaddsd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vexF3 | vex0F | vexW0, 0x5C,
			avxEscape | evex128 | evexF3 | evex0F | evexW0, evexN4 | evexRoundingEnabled | evexZeroingEnabled, 0x5C,
		}},
		{as: AVTESTPD, ytab: psess._yvptest, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0F,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0F,
		}},
		{as: AVTESTPS, ytab: psess._yvptest, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F38 | vexW0, 0x0E,
			avxEscape | vex256 | vex66 | vex0F38 | vexW0, 0x0E,
		}},
		{as: AVUCOMISD, ytab: psess._yvcomisd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x2E,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN8 | evexSaeEnabled, 0x2E,
		}},
		{as: AVUCOMISS, ytab: psess._yvcomisd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x2E,
			avxEscape | evex128 | evex0F | evexW0, evexN4 | evexSaeEnabled, 0x2E,
		}},
		{as: AVUNPCKHPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x15,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x15,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x15,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x15,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x15,
		}},
		{as: AVUNPCKHPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x15,
			avxEscape | vex256 | vex0F | vexW0, 0x15,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x15,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x15,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x15,
		}},
		{as: AVUNPCKLPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x14,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x14,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x14,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x14,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x14,
		}},
		{as: AVUNPCKLPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x14,
			avxEscape | vex256 | vex0F | vexW0, 0x14,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x14,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x14,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x14,
		}},
		{as: AVXORPD, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex66 | vex0F | vexW0, 0x57,
			avxEscape | vex256 | vex66 | vex0F | vexW0, 0x57,
			avxEscape | evex128 | evex66 | evex0F | evexW1, evexN16 | evexBcstN8 | evexZeroingEnabled, 0x57,
			avxEscape | evex256 | evex66 | evex0F | evexW1, evexN32 | evexBcstN8 | evexZeroingEnabled, 0x57,
			avxEscape | evex512 | evex66 | evex0F | evexW1, evexN64 | evexBcstN8 | evexZeroingEnabled, 0x57,
		}},
		{as: AVXORPS, ytab: psess._yvandnpd, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x57,
			avxEscape | vex256 | vex0F | vexW0, 0x57,
			avxEscape | evex128 | evex0F | evexW0, evexN16 | evexBcstN4 | evexZeroingEnabled, 0x57,
			avxEscape | evex256 | evex0F | evexW0, evexN32 | evexBcstN4 | evexZeroingEnabled, 0x57,
			avxEscape | evex512 | evex0F | evexW0, evexN64 | evexBcstN4 | evexZeroingEnabled, 0x57,
		}},
		{as: AVZEROALL, ytab: psess._yvzeroall, prefix: Pavx, op: opBytes{
			avxEscape | vex256 | vex0F | vexW0, 0x77,
		}},
		{as: AVZEROUPPER, ytab: psess._yvzeroall, prefix: Pavx, op: opBytes{
			avxEscape | vex128 | vex0F | vexW0, 0x77,
		}},
	}
	psess.Linkamd64 = obj.LinkArch{
		Arch:           psess.sys.ArchAMD64,
		Init:           psess.instinit,
		Preprocess:     psess.preprocess,
		Assemble:       psess.span6,
		Progedit:       psess.progedit,
		UnaryDst:       psess.unaryDst,
		DWARFRegisters: psess.AMD64DWARFRegisters,
	}
	psess.Linkamd64p32 = obj.LinkArch{
		Arch:           psess.sys.ArchAMD64P32,
		Init:           psess.instinit,
		Preprocess:     psess.preprocess,
		Assemble:       psess.span6,
		Progedit:       psess.progedit,
		UnaryDst:       psess.unaryDst,
		DWARFRegisters: psess.AMD64DWARFRegisters,
	}
	psess.Link386 = obj.LinkArch{
		Arch:           psess.sys.Arch386,
		Init:           psess.instinit,
		Preprocess:     psess.preprocess,
		Assemble:       psess.span6,
		Progedit:       psess.progedit,
		UnaryDst:       psess.unaryDst,
		DWARFRegisters: psess.X86DWARFRegisters,
	}
	psess.opSuffixTable = [...]string{
		"",

		"Z",

		"SAE",
		"SAE.Z",

		"RN_SAE",
		"RZ_SAE",
		"RD_SAE",
		"RU_SAE",
		"RN_SAE.Z",
		"RZ_SAE.Z",
		"RD_SAE.Z",
		"RU_SAE.Z",

		"BCST",
		"BCST.Z",

		"<bad suffix>",
	}
	return psess
}
