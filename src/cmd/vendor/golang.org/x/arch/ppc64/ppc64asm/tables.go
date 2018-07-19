package ppc64asm

const (
	_ Op = iota
	CNTLZW
	CNTLZW_
	B
	BA
	BL
	BLA
	BC
	BCA
	BCL
	BCLA
	BCLR
	BCLRL
	BCCTR
	BCCTRL
	BCTAR
	BCTARL
	CRAND
	CROR
	CRNAND
	CRXOR
	CRNOR
	CRANDC
	MCRF
	CREQV
	CRORC
	SC
	CLRBHRB
	MFBHRBE
	LBZ
	LBZU
	LBZX
	LBZUX
	LHZ
	LHZU
	LHZX
	LHZUX
	LHA
	LHAU
	LHAX
	LHAUX
	LWZ
	LWZU
	LWZX
	LWZUX
	LWA
	LWAX
	LWAUX
	LD
	LDU
	LDX
	LDUX
	STB
	STBU
	STBX
	STBUX
	STH
	STHU
	STHX
	STHUX
	STW
	STWU
	STWX
	STWUX
	STD
	STDU
	STDX
	STDUX
	LQ
	STQ
	LHBRX
	LWBRX
	STHBRX
	STWBRX
	LDBRX
	STDBRX
	LMW
	STMW
	LSWI
	LSWX
	STSWI
	STSWX
	LI
	ADDI
	LIS
	ADDIS
	ADD
	ADD_
	ADDO
	ADDO_
	ADDIC
	SUBF
	SUBF_
	SUBFO
	SUBFO_
	ADDIC_
	SUBFIC
	ADDC
	ADDC_
	ADDCO
	ADDCO_
	SUBFC
	SUBFC_
	SUBFCO
	SUBFCO_
	ADDE
	ADDE_
	ADDEO
	ADDEO_
	ADDME
	ADDME_
	ADDMEO
	ADDMEO_
	SUBFE
	SUBFE_
	SUBFEO
	SUBFEO_
	SUBFME
	SUBFME_
	SUBFMEO
	SUBFMEO_
	ADDZE
	ADDZE_
	ADDZEO
	ADDZEO_
	SUBFZE
	SUBFZE_
	SUBFZEO
	SUBFZEO_
	NEG
	NEG_
	NEGO
	NEGO_
	MULLI
	MULLW
	MULLW_
	MULLWO
	MULLWO_
	MULHW
	MULHW_
	MULHWU
	MULHWU_
	DIVW
	DIVW_
	DIVWO
	DIVWO_
	DIVWU
	DIVWU_
	DIVWUO
	DIVWUO_
	DIVWE
	DIVWE_
	DIVWEO
	DIVWEO_
	DIVWEU
	DIVWEU_
	DIVWEUO
	DIVWEUO_
	MULLD
	MULLD_
	MULLDO
	MULLDO_
	MULHDU
	MULHDU_
	MULHD
	MULHD_
	DIVD
	DIVD_
	DIVDO
	DIVDO_
	DIVDU
	DIVDU_
	DIVDUO
	DIVDUO_
	DIVDE
	DIVDE_
	DIVDEO
	DIVDEO_
	DIVDEU
	DIVDEU_
	DIVDEUO
	DIVDEUO_
	CMPWI
	CMPDI
	CMPW
	CMPD
	CMPLWI
	CMPLDI
	CMPLW
	CMPLD
	TWI
	TW
	TDI
	ISEL
	TD
	ANDI_
	ANDIS_
	ORI
	ORIS
	XORI
	XORIS
	AND
	AND_
	XOR
	XOR_
	NAND
	NAND_
	OR
	OR_
	NOR
	NOR_
	ANDC
	ANDC_
	EXTSB
	EXTSB_
	EQV
	EQV_
	ORC
	ORC_
	EXTSH
	EXTSH_
	CMPB
	POPCNTB
	POPCNTW
	PRTYD
	PRTYW
	EXTSW
	EXTSW_
	CNTLZD
	CNTLZD_
	POPCNTD
	BPERMD
	RLWINM
	RLWINM_
	RLWNM
	RLWNM_
	RLWIMI
	RLWIMI_
	RLDICL
	RLDICL_
	RLDICR
	RLDICR_
	RLDIC
	RLDIC_
	RLDCL
	RLDCL_
	RLDCR
	RLDCR_
	RLDIMI
	RLDIMI_
	SLW
	SLW_
	SRW
	SRW_
	SRAWI
	SRAWI_
	SRAW
	SRAW_
	SLD
	SLD_
	SRD
	SRD_
	SRADI
	SRADI_
	SRAD
	SRAD_
	CDTBCD
	CBCDTD
	ADDG6S
	MTSPR
	MFSPR
	MTCRF
	MFCR
	MTSLE
	MFVSRD
	MFVSRWZ
	MTVSRD
	MTVSRWA
	MTVSRWZ
	MTOCRF
	MFOCRF
	MCRXR
	MTDCRUX
	MFDCRUX
	LFS
	LFSU
	LFSX
	LFSUX
	LFD
	LFDU
	LFDX
	LFDUX
	LFIWAX
	LFIWZX
	STFS
	STFSU
	STFSX
	STFSUX
	STFD
	STFDU
	STFDX
	STFDUX
	STFIWX
	LFDP
	LFDPX
	STFDP
	STFDPX
	FMR
	FMR_
	FABS
	FABS_
	FNABS
	FNABS_
	FNEG
	FNEG_
	FCPSGN
	FCPSGN_
	FMRGEW
	FMRGOW
	FADD
	FADD_
	FADDS
	FADDS_
	FSUB
	FSUB_
	FSUBS
	FSUBS_
	FMUL
	FMUL_
	FMULS
	FMULS_
	FDIV
	FDIV_
	FDIVS
	FDIVS_
	FSQRT
	FSQRT_
	FSQRTS
	FSQRTS_
	FRE
	FRE_
	FRES
	FRES_
	FRSQRTE
	FRSQRTE_
	FRSQRTES
	FRSQRTES_
	FTDIV
	FTSQRT
	FMADD
	FMADD_
	FMADDS
	FMADDS_
	FMSUB
	FMSUB_
	FMSUBS
	FMSUBS_
	FNMADD
	FNMADD_
	FNMADDS
	FNMADDS_
	FNMSUB
	FNMSUB_
	FNMSUBS
	FNMSUBS_
	FRSP
	FRSP_
	FCTID
	FCTID_
	FCTIDZ
	FCTIDZ_
	FCTIDU
	FCTIDU_
	FCTIDUZ
	FCTIDUZ_
	FCTIW
	FCTIW_
	FCTIWZ
	FCTIWZ_
	FCTIWU
	FCTIWU_
	FCTIWUZ
	FCTIWUZ_
	FCFID
	FCFID_
	FCFIDU
	FCFIDU_
	FCFIDS
	FCFIDS_
	FCFIDUS
	FCFIDUS_
	FRIN
	FRIN_
	FRIZ
	FRIZ_
	FRIP
	FRIP_
	FRIM
	FRIM_
	FCMPU
	FCMPO
	FSEL
	FSEL_
	MFFS
	MFFS_
	MCRFS
	MTFSFI
	MTFSFI_
	MTFSF
	MTFSF_
	MTFSB0
	MTFSB0_
	MTFSB1
	MTFSB1_
	LVEBX
	LVEHX
	LVEWX
	LVX
	LVXL
	STVEBX
	STVEHX
	STVEWX
	STVX
	STVXL
	LVSL
	LVSR
	VPKPX
	VPKSDSS
	VPKSDUS
	VPKSHSS
	VPKSHUS
	VPKSWSS
	VPKSWUS
	VPKUDUM
	VPKUDUS
	VPKUHUM
	VPKUHUS
	VPKUWUM
	VPKUWUS
	VUPKHPX
	VUPKLPX
	VUPKHSB
	VUPKHSH
	VUPKHSW
	VUPKLSB
	VUPKLSH
	VUPKLSW
	VMRGHB
	VMRGHH
	VMRGLB
	VMRGLH
	VMRGHW
	VMRGLW
	VMRGEW
	VMRGOW
	VSPLTB
	VSPLTH
	VSPLTW
	VSPLTISB
	VSPLTISH
	VSPLTISW
	VPERM
	VSEL
	VSL
	VSLDOI
	VSLO
	VSR
	VSRO
	VADDCUW
	VADDSBS
	VADDSHS
	VADDSWS
	VADDUBM
	VADDUDM
	VADDUHM
	VADDUWM
	VADDUBS
	VADDUHS
	VADDUWS
	VADDUQM
	VADDEUQM
	VADDCUQ
	VADDECUQ
	VSUBCUW
	VSUBSBS
	VSUBSHS
	VSUBSWS
	VSUBUBM
	VSUBUDM
	VSUBUHM
	VSUBUWM
	VSUBUBS
	VSUBUHS
	VSUBUWS
	VSUBUQM
	VSUBEUQM
	VSUBCUQ
	VSUBECUQ
	VMULESB
	VMULEUB
	VMULOSB
	VMULOUB
	VMULESH
	VMULEUH
	VMULOSH
	VMULOUH
	VMULESW
	VMULEUW
	VMULOSW
	VMULOUW
	VMULUWM
	VMHADDSHS
	VMHRADDSHS
	VMLADDUHM
	VMSUMUBM
	VMSUMMBM
	VMSUMSHM
	VMSUMSHS
	VMSUMUHM
	VMSUMUHS
	VSUMSWS
	VSUM2SWS
	VSUM4SBS
	VSUM4SHS
	VSUM4UBS
	VAVGSB
	VAVGSH
	VAVGSW
	VAVGUB
	VAVGUW
	VAVGUH
	VMAXSB
	VMAXSD
	VMAXUB
	VMAXUD
	VMAXSH
	VMAXSW
	VMAXUH
	VMAXUW
	VMINSB
	VMINSD
	VMINUB
	VMINUD
	VMINSH
	VMINSW
	VMINUH
	VMINUW
	VCMPEQUB
	VCMPEQUB_
	VCMPEQUH
	VCMPEQUH_
	VCMPEQUW
	VCMPEQUW_
	VCMPEQUD
	VCMPEQUD_
	VCMPGTSB
	VCMPGTSB_
	VCMPGTSD
	VCMPGTSD_
	VCMPGTSH
	VCMPGTSH_
	VCMPGTSW
	VCMPGTSW_
	VCMPGTUB
	VCMPGTUB_
	VCMPGTUD
	VCMPGTUD_
	VCMPGTUH
	VCMPGTUH_
	VCMPGTUW
	VCMPGTUW_
	VAND
	VANDC
	VEQV
	VNAND
	VORC
	VNOR
	VOR
	VXOR
	VRLB
	VRLH
	VRLW
	VRLD
	VSLB
	VSLH
	VSLW
	VSLD
	VSRB
	VSRH
	VSRW
	VSRD
	VSRAB
	VSRAH
	VSRAW
	VSRAD
	VADDFP
	VSUBFP
	VMADDFP
	VNMSUBFP
	VMAXFP
	VMINFP
	VCTSXS
	VCTUXS
	VCFSX
	VCFUX
	VRFIM
	VRFIN
	VRFIP
	VRFIZ
	VCMPBFP
	VCMPBFP_
	VCMPEQFP
	VCMPEQFP_
	VCMPGEFP
	VCMPGEFP_
	VCMPGTFP
	VCMPGTFP_
	VEXPTEFP
	VLOGEFP
	VREFP
	VRSQRTEFP
	VCIPHER
	VCIPHERLAST
	VNCIPHER
	VNCIPHERLAST
	VSBOX
	VSHASIGMAD
	VSHASIGMAW
	VPMSUMB
	VPMSUMD
	VPMSUMH
	VPMSUMW
	VPERMXOR
	VGBBD
	VCLZB
	VCLZH
	VCLZW
	VCLZD
	VPOPCNTB
	VPOPCNTD
	VPOPCNTH
	VPOPCNTW
	VBPERMQ
	BCDADD_
	BCDSUB_
	MTVSCR
	MFVSCR
	DADD
	DADD_
	DSUB
	DSUB_
	DMUL
	DMUL_
	DDIV
	DDIV_
	DCMPU
	DCMPO
	DTSTDC
	DTSTDG
	DTSTEX
	DTSTSF
	DQUAI
	DQUAI_
	DQUA
	DQUA_
	DRRND
	DRRND_
	DRINTX
	DRINTX_
	DRINTN
	DRINTN_
	DCTDP
	DCTDP_
	DCTQPQ
	DCTQPQ_
	DRSP
	DRSP_
	DRDPQ
	DRDPQ_
	DCFFIX
	DCFFIX_
	DCFFIXQ
	DCFFIXQ_
	DCTFIX
	DCTFIX_
	DDEDPD
	DDEDPD_
	DENBCD
	DENBCD_
	DXEX
	DXEX_
	DIEX
	DIEX_
	DSCLI
	DSCLI_
	DSCRI
	DSCRI_
	LXSDX
	LXSIWAX
	LXSIWZX
	LXSSPX
	LXVD2X
	LXVDSX
	LXVW4X
	STXSDX
	STXSIWX
	STXSSPX
	STXVD2X
	STXVW4X
	XSABSDP
	XSADDDP
	XSADDSP
	XSCMPODP
	XSCMPUDP
	XSCPSGNDP
	XSCVDPSP
	XSCVDPSPN
	XSCVDPSXDS
	XSCVDPSXWS
	XSCVDPUXDS
	XSCVDPUXWS
	XSCVSPDP
	XSCVSPDPN
	XSCVSXDDP
	XSCVSXDSP
	XSCVUXDDP
	XSCVUXDSP
	XSDIVDP
	XSDIVSP
	XSMADDADP
	XSMADDASP
	XSMAXDP
	XSMINDP
	XSMSUBADP
	XSMSUBASP
	XSMULDP
	XSMULSP
	XSNABSDP
	XSNEGDP
	XSNMADDADP
	XSNMADDASP
	XSNMSUBADP
	XSNMSUBASP
	XSRDPI
	XSRDPIC
	XSRDPIM
	XSRDPIP
	XSRDPIZ
	XSREDP
	XSRESP
	XSRSP
	XSRSQRTEDP
	XSRSQRTESP
	XSSQRTDP
	XSSQRTSP
	XSSUBDP
	XSSUBSP
	XSTDIVDP
	XSTSQRTDP
	XVABSDP
	XVABSSP
	XVADDDP
	XVADDSP
	XVCMPEQDP
	XVCMPEQDP_
	XVCMPEQSP
	XVCMPEQSP_
	XVCMPGEDP
	XVCMPGEDP_
	XVCMPGESP
	XVCMPGESP_
	XVCMPGTDP
	XVCMPGTDP_
	XVCMPGTSP
	XVCMPGTSP_
	XVCPSGNDP
	XVCPSGNSP
	XVCVDPSP
	XVCVDPSXDS
	XVCVDPSXWS
	XVCVDPUXDS
	XVCVDPUXWS
	XVCVSPDP
	XVCVSPSXDS
	XVCVSPSXWS
	XVCVSPUXDS
	XVCVSPUXWS
	XVCVSXDDP
	XVCVSXDSP
	XVCVSXWDP
	XVCVSXWSP
	XVCVUXDDP
	XVCVUXDSP
	XVCVUXWDP
	XVCVUXWSP
	XVDIVDP
	XVDIVSP
	XVMADDADP
	XVMADDASP
	XVMAXDP
	XVMAXSP
	XVMINDP
	XVMINSP
	XVMSUBADP
	XVMSUBASP
	XVMULDP
	XVMULSP
	XVNABSDP
	XVNABSSP
	XVNEGDP
	XVNEGSP
	XVNMADDADP
	XVNMADDASP
	XVNMSUBADP
	XVNMSUBASP
	XVRDPI
	XVRDPIC
	XVRDPIM
	XVRDPIP
	XVRDPIZ
	XVREDP
	XVRESP
	XVRSPI
	XVRSPIC
	XVRSPIM
	XVRSPIP
	XVRSPIZ
	XVRSQRTEDP
	XVRSQRTESP
	XVSQRTDP
	XVSQRTSP
	XVSUBDP
	XVSUBSP
	XVTDIVDP
	XVTDIVSP
	XVTSQRTDP
	XVTSQRTSP
	XXLAND
	XXLANDC
	XXLEQV
	XXLNAND
	XXLORC
	XXLNOR
	XXLOR
	XXLXOR
	XXMRGHW
	XXMRGLW
	XXPERMDI
	XXSEL
	XXSLDWI
	XXSPLTW
	BRINC
	EVABS
	EVADDIW
	EVADDSMIAAW
	EVADDSSIAAW
	EVADDUMIAAW
	EVADDUSIAAW
	EVADDW
	EVAND
	EVCMPEQ
	EVANDC
	EVCMPGTS
	EVCMPGTU
	EVCMPLTU
	EVCMPLTS
	EVCNTLSW
	EVCNTLZW
	EVDIVWS
	EVDIVWU
	EVEQV
	EVEXTSB
	EVEXTSH
	EVLDD
	EVLDH
	EVLDDX
	EVLDHX
	EVLDW
	EVLHHESPLAT
	EVLDWX
	EVLHHESPLATX
	EVLHHOSSPLAT
	EVLHHOUSPLAT
	EVLHHOSSPLATX
	EVLHHOUSPLATX
	EVLWHE
	EVLWHOS
	EVLWHEX
	EVLWHOSX
	EVLWHOU
	EVLWHSPLAT
	EVLWHOUX
	EVLWHSPLATX
	EVLWWSPLAT
	EVMERGEHI
	EVLWWSPLATX
	EVMERGELO
	EVMERGEHILO
	EVMHEGSMFAA
	EVMERGELOHI
	EVMHEGSMFAN
	EVMHEGSMIAA
	EVMHEGUMIAA
	EVMHEGSMIAN
	EVMHEGUMIAN
	EVMHESMF
	EVMHESMFAAW
	EVMHESMFA
	EVMHESMFANW
	EVMHESMI
	EVMHESMIAAW
	EVMHESMIA
	EVMHESMIANW
	EVMHESSF
	EVMHESSFA
	EVMHESSFAAW
	EVMHESSFANW
	EVMHESSIAAW
	EVMHESSIANW
	EVMHEUMI
	EVMHEUMIAAW
	EVMHEUMIA
	EVMHEUMIANW
	EVMHEUSIAAW
	EVMHEUSIANW
	EVMHOGSMFAA
	EVMHOGSMIAA
	EVMHOGSMFAN
	EVMHOGSMIAN
	EVMHOGUMIAA
	EVMHOSMF
	EVMHOGUMIAN
	EVMHOSMFA
	EVMHOSMFAAW
	EVMHOSMI
	EVMHOSMFANW
	EVMHOSMIA
	EVMHOSMIAAW
	EVMHOSMIANW
	EVMHOSSF
	EVMHOSSFA
	EVMHOSSFAAW
	EVMHOSSFANW
	EVMHOSSIAAW
	EVMHOUMI
	EVMHOSSIANW
	EVMHOUMIA
	EVMHOUMIAAW
	EVMHOUSIAAW
	EVMHOUMIANW
	EVMHOUSIANW
	EVMRA
	EVMWHSMF
	EVMWHSMI
	EVMWHSMFA
	EVMWHSMIA
	EVMWHSSF
	EVMWHUMI
	EVMWHSSFA
	EVMWHUMIA
	EVMWLSMIAAW
	EVMWLSSIAAW
	EVMWLSMIANW
	EVMWLSSIANW
	EVMWLUMI
	EVMWLUMIAAW
	EVMWLUMIA
	EVMWLUMIANW
	EVMWLUSIAAW
	EVMWSMF
	EVMWLUSIANW
	EVMWSMFA
	EVMWSMFAA
	EVMWSMI
	EVMWSMIAA
	EVMWSMFAN
	EVMWSMIA
	EVMWSMIAN
	EVMWSSF
	EVMWSSFA
	EVMWSSFAA
	EVMWUMI
	EVMWSSFAN
	EVMWUMIA
	EVMWUMIAA
	EVNAND
	EVMWUMIAN
	EVNEG
	EVNOR
	EVORC
	EVOR
	EVRLW
	EVRLWI
	EVSEL
	EVRNDW
	EVSLW
	EVSPLATFI
	EVSRWIS
	EVSLWI
	EVSPLATI
	EVSRWIU
	EVSRWS
	EVSTDD
	EVSRWU
	EVSTDDX
	EVSTDH
	EVSTDW
	EVSTDHX
	EVSTDWX
	EVSTWHE
	EVSTWHO
	EVSTWWE
	EVSTWHEX
	EVSTWHOX
	EVSTWWEX
	EVSTWWO
	EVSUBFSMIAAW
	EVSTWWOX
	EVSUBFSSIAAW
	EVSUBFUMIAAW
	EVSUBFUSIAAW
	EVSUBFW
	EVSUBIFW
	EVXOR
	EVFSABS
	EVFSNABS
	EVFSNEG
	EVFSADD
	EVFSMUL
	EVFSSUB
	EVFSDIV
	EVFSCMPGT
	EVFSCMPLT
	EVFSCMPEQ
	EVFSTSTGT
	EVFSTSTLT
	EVFSTSTEQ
	EVFSCFSI
	EVFSCFSF
	EVFSCFUI
	EVFSCFUF
	EVFSCTSI
	EVFSCTUI
	EVFSCTSIZ
	EVFSCTUIZ
	EVFSCTSF
	EVFSCTUF
	EFSABS
	EFSNEG
	EFSNABS
	EFSADD
	EFSMUL
	EFSSUB
	EFSDIV
	EFSCMPGT
	EFSCMPLT
	EFSCMPEQ
	EFSTSTGT
	EFSTSTLT
	EFSTSTEQ
	EFSCFSI
	EFSCFSF
	EFSCTSI
	EFSCFUI
	EFSCFUF
	EFSCTUI
	EFSCTSIZ
	EFSCTSF
	EFSCTUIZ
	EFSCTUF
	EFDABS
	EFDNEG
	EFDNABS
	EFDADD
	EFDMUL
	EFDSUB
	EFDDIV
	EFDCMPGT
	EFDCMPEQ
	EFDCMPLT
	EFDTSTGT
	EFDTSTLT
	EFDCFSI
	EFDTSTEQ
	EFDCFUI
	EFDCFSID
	EFDCFSF
	EFDCFUF
	EFDCFUID
	EFDCTSI
	EFDCTUI
	EFDCTSIDZ
	EFDCTUIDZ
	EFDCTSIZ
	EFDCTSF
	EFDCTUF
	EFDCTUIZ
	EFDCFS
	EFSCFD
	DLMZB
	DLMZB_
	MACCHW
	MACCHW_
	MACCHWO
	MACCHWO_
	MACCHWS
	MACCHWS_
	MACCHWSO
	MACCHWSO_
	MACCHWU
	MACCHWU_
	MACCHWUO
	MACCHWUO_
	MACCHWSU
	MACCHWSU_
	MACCHWSUO
	MACCHWSUO_
	MACHHW
	MACHHW_
	MACHHWO
	MACHHWO_
	MACHHWS
	MACHHWS_
	MACHHWSO
	MACHHWSO_
	MACHHWU
	MACHHWU_
	MACHHWUO
	MACHHWUO_
	MACHHWSU
	MACHHWSU_
	MACHHWSUO
	MACHHWSUO_
	MACLHW
	MACLHW_
	MACLHWO
	MACLHWO_
	MACLHWS
	MACLHWS_
	MACLHWSO
	MACLHWSO_
	MACLHWU
	MACLHWU_
	MACLHWUO
	MACLHWUO_
	MULCHW
	MULCHW_
	MACLHWSU
	MACLHWSU_
	MACLHWSUO
	MACLHWSUO_
	MULCHWU
	MULCHWU_
	MULHHW
	MULHHW_
	MULLHW
	MULLHW_
	MULHHWU
	MULHHWU_
	MULLHWU
	MULLHWU_
	NMACCHW
	NMACCHW_
	NMACCHWO
	NMACCHWO_
	NMACCHWS
	NMACCHWS_
	NMACCHWSO
	NMACCHWSO_
	NMACHHW
	NMACHHW_
	NMACHHWO
	NMACHHWO_
	NMACHHWS
	NMACHHWS_
	NMACHHWSO
	NMACHHWSO_
	NMACLHW
	NMACLHW_
	NMACLHWO
	NMACLHWO_
	NMACLHWS
	NMACLHWS_
	NMACLHWSO
	NMACLHWSO_
	ICBI
	ICBT
	DCBA
	DCBT
	DCBTST
	DCBZ
	DCBST
	DCBF
	ISYNC
	LBARX
	LHARX
	LWARX
	STBCX_
	STHCX_
	STWCX_
	LDARX
	STDCX_
	LQARX
	STQCX_
	SYNC
	EIEIO
	MBAR
	WAIT
	TBEGIN_
	TEND_
	TABORT_
	TABORTWC_
	TABORTWCI_
	TABORTDC_
	TABORTDCI_
	TSR_
	TCHECK
	MFTB
	RFEBB
	LBDX
	LHDX
	LWDX
	LDDX
	LFDDX
	STBDX
	STHDX
	STWDX
	STDDX
	STFDDX
	DSN
	ECIWX
	ECOWX
	RFID
	HRFID
	DOZE
	NAP
	SLEEP
	RVWINKLE
	LBZCIX
	LWZCIX
	LHZCIX
	LDCIX
	STBCIX
	STWCIX
	STHCIX
	STDCIX
	TRECLAIM_
	TRECHKPT_
	MTMSR
	MTMSRD
	MFMSR
	SLBIE
	SLBIA
	SLBMTE
	SLBMFEV
	SLBMFEE
	SLBFEE_
	MTSR
	MTSRIN
	MFSR
	MFSRIN
	TLBIE
	TLBIEL
	TLBIA
	TLBSYNC
	MSGSND
	MSGCLR
	MSGSNDP
	MSGCLRP
	MTTMR
	RFI
	RFCI
	RFDI
	RFMCI
	RFGI
	EHPRIV
	MTDCR
	MTDCRX
	MFDCR
	MFDCRX
	WRTEE
	WRTEEI
	LBEPX
	LHEPX
	LWEPX
	LDEPX
	STBEPX
	STHEPX
	STWEPX
	STDEPX
	DCBSTEP
	DCBTEP
	DCBFEP
	DCBTSTEP
	ICBIEP
	DCBZEP
	LFDEPX
	STFDEPX
	EVLDDEPX
	EVSTDDEPX
	LVEPX
	LVEPXL
	STVEPX
	STVEPXL
	DCBI
	DCBLQ_
	ICBLQ_
	DCBTLS
	DCBTSTLS
	ICBTLS
	ICBLC
	DCBLC
	TLBIVAX
	TLBILX
	TLBSX
	TLBSRX_
	TLBRE
	TLBWE
	DNH
	DCI
	ICI
	DCREAD
	ICREAD
	MFPMR
	MTPMR
)

var opstr = [...]string{
	CNTLZW:        "cntlzw",
	CNTLZW_:       "cntlzw.",
	B:             "b",
	BA:            "ba",
	BL:            "bl",
	BLA:           "bla",
	BC:            "bc",
	BCA:           "bca",
	BCL:           "bcl",
	BCLA:          "bcla",
	BCLR:          "bclr",
	BCLRL:         "bclrl",
	BCCTR:         "bcctr",
	BCCTRL:        "bcctrl",
	BCTAR:         "bctar",
	BCTARL:        "bctarl",
	CRAND:         "crand",
	CROR:          "cror",
	CRNAND:        "crnand",
	CRXOR:         "crxor",
	CRNOR:         "crnor",
	CRANDC:        "crandc",
	MCRF:          "mcrf",
	CREQV:         "creqv",
	CRORC:         "crorc",
	SC:            "sc",
	CLRBHRB:       "clrbhrb",
	MFBHRBE:       "mfbhrbe",
	LBZ:           "lbz",
	LBZU:          "lbzu",
	LBZX:          "lbzx",
	LBZUX:         "lbzux",
	LHZ:           "lhz",
	LHZU:          "lhzu",
	LHZX:          "lhzx",
	LHZUX:         "lhzux",
	LHA:           "lha",
	LHAU:          "lhau",
	LHAX:          "lhax",
	LHAUX:         "lhaux",
	LWZ:           "lwz",
	LWZU:          "lwzu",
	LWZX:          "lwzx",
	LWZUX:         "lwzux",
	LWA:           "lwa",
	LWAX:          "lwax",
	LWAUX:         "lwaux",
	LD:            "ld",
	LDU:           "ldu",
	LDX:           "ldx",
	LDUX:          "ldux",
	STB:           "stb",
	STBU:          "stbu",
	STBX:          "stbx",
	STBUX:         "stbux",
	STH:           "sth",
	STHU:          "sthu",
	STHX:          "sthx",
	STHUX:         "sthux",
	STW:           "stw",
	STWU:          "stwu",
	STWX:          "stwx",
	STWUX:         "stwux",
	STD:           "std",
	STDU:          "stdu",
	STDX:          "stdx",
	STDUX:         "stdux",
	LQ:            "lq",
	STQ:           "stq",
	LHBRX:         "lhbrx",
	LWBRX:         "lwbrx",
	STHBRX:        "sthbrx",
	STWBRX:        "stwbrx",
	LDBRX:         "ldbrx",
	STDBRX:        "stdbrx",
	LMW:           "lmw",
	STMW:          "stmw",
	LSWI:          "lswi",
	LSWX:          "lswx",
	STSWI:         "stswi",
	STSWX:         "stswx",
	LI:            "li",
	ADDI:          "addi",
	LIS:           "lis",
	ADDIS:         "addis",
	ADD:           "add",
	ADD_:          "add.",
	ADDO:          "addo",
	ADDO_:         "addo.",
	ADDIC:         "addic",
	SUBF:          "subf",
	SUBF_:         "subf.",
	SUBFO:         "subfo",
	SUBFO_:        "subfo.",
	ADDIC_:        "addic.",
	SUBFIC:        "subfic",
	ADDC:          "addc",
	ADDC_:         "addc.",
	ADDCO:         "addco",
	ADDCO_:        "addco.",
	SUBFC:         "subfc",
	SUBFC_:        "subfc.",
	SUBFCO:        "subfco",
	SUBFCO_:       "subfco.",
	ADDE:          "adde",
	ADDE_:         "adde.",
	ADDEO:         "addeo",
	ADDEO_:        "addeo.",
	ADDME:         "addme",
	ADDME_:        "addme.",
	ADDMEO:        "addmeo",
	ADDMEO_:       "addmeo.",
	SUBFE:         "subfe",
	SUBFE_:        "subfe.",
	SUBFEO:        "subfeo",
	SUBFEO_:       "subfeo.",
	SUBFME:        "subfme",
	SUBFME_:       "subfme.",
	SUBFMEO:       "subfmeo",
	SUBFMEO_:      "subfmeo.",
	ADDZE:         "addze",
	ADDZE_:        "addze.",
	ADDZEO:        "addzeo",
	ADDZEO_:       "addzeo.",
	SUBFZE:        "subfze",
	SUBFZE_:       "subfze.",
	SUBFZEO:       "subfzeo",
	SUBFZEO_:      "subfzeo.",
	NEG:           "neg",
	NEG_:          "neg.",
	NEGO:          "nego",
	NEGO_:         "nego.",
	MULLI:         "mulli",
	MULLW:         "mullw",
	MULLW_:        "mullw.",
	MULLWO:        "mullwo",
	MULLWO_:       "mullwo.",
	MULHW:         "mulhw",
	MULHW_:        "mulhw.",
	MULHWU:        "mulhwu",
	MULHWU_:       "mulhwu.",
	DIVW:          "divw",
	DIVW_:         "divw.",
	DIVWO:         "divwo",
	DIVWO_:        "divwo.",
	DIVWU:         "divwu",
	DIVWU_:        "divwu.",
	DIVWUO:        "divwuo",
	DIVWUO_:       "divwuo.",
	DIVWE:         "divwe",
	DIVWE_:        "divwe.",
	DIVWEO:        "divweo",
	DIVWEO_:       "divweo.",
	DIVWEU:        "divweu",
	DIVWEU_:       "divweu.",
	DIVWEUO:       "divweuo",
	DIVWEUO_:      "divweuo.",
	MULLD:         "mulld",
	MULLD_:        "mulld.",
	MULLDO:        "mulldo",
	MULLDO_:       "mulldo.",
	MULHDU:        "mulhdu",
	MULHDU_:       "mulhdu.",
	MULHD:         "mulhd",
	MULHD_:        "mulhd.",
	DIVD:          "divd",
	DIVD_:         "divd.",
	DIVDO:         "divdo",
	DIVDO_:        "divdo.",
	DIVDU:         "divdu",
	DIVDU_:        "divdu.",
	DIVDUO:        "divduo",
	DIVDUO_:       "divduo.",
	DIVDE:         "divde",
	DIVDE_:        "divde.",
	DIVDEO:        "divdeo",
	DIVDEO_:       "divdeo.",
	DIVDEU:        "divdeu",
	DIVDEU_:       "divdeu.",
	DIVDEUO:       "divdeuo",
	DIVDEUO_:      "divdeuo.",
	CMPWI:         "cmpwi",
	CMPDI:         "cmpdi",
	CMPW:          "cmpw",
	CMPD:          "cmpd",
	CMPLWI:        "cmplwi",
	CMPLDI:        "cmpldi",
	CMPLW:         "cmplw",
	CMPLD:         "cmpld",
	TWI:           "twi",
	TW:            "tw",
	TDI:           "tdi",
	ISEL:          "isel",
	TD:            "td",
	ANDI_:         "andi.",
	ANDIS_:        "andis.",
	ORI:           "ori",
	ORIS:          "oris",
	XORI:          "xori",
	XORIS:         "xoris",
	AND:           "and",
	AND_:          "and.",
	XOR:           "xor",
	XOR_:          "xor.",
	NAND:          "nand",
	NAND_:         "nand.",
	OR:            "or",
	OR_:           "or.",
	NOR:           "nor",
	NOR_:          "nor.",
	ANDC:          "andc",
	ANDC_:         "andc.",
	EXTSB:         "extsb",
	EXTSB_:        "extsb.",
	EQV:           "eqv",
	EQV_:          "eqv.",
	ORC:           "orc",
	ORC_:          "orc.",
	EXTSH:         "extsh",
	EXTSH_:        "extsh.",
	CMPB:          "cmpb",
	POPCNTB:       "popcntb",
	POPCNTW:       "popcntw",
	PRTYD:         "prtyd",
	PRTYW:         "prtyw",
	EXTSW:         "extsw",
	EXTSW_:        "extsw.",
	CNTLZD:        "cntlzd",
	CNTLZD_:       "cntlzd.",
	POPCNTD:       "popcntd",
	BPERMD:        "bpermd",
	RLWINM:        "rlwinm",
	RLWINM_:       "rlwinm.",
	RLWNM:         "rlwnm",
	RLWNM_:        "rlwnm.",
	RLWIMI:        "rlwimi",
	RLWIMI_:       "rlwimi.",
	RLDICL:        "rldicl",
	RLDICL_:       "rldicl.",
	RLDICR:        "rldicr",
	RLDICR_:       "rldicr.",
	RLDIC:         "rldic",
	RLDIC_:        "rldic.",
	RLDCL:         "rldcl",
	RLDCL_:        "rldcl.",
	RLDCR:         "rldcr",
	RLDCR_:        "rldcr.",
	RLDIMI:        "rldimi",
	RLDIMI_:       "rldimi.",
	SLW:           "slw",
	SLW_:          "slw.",
	SRW:           "srw",
	SRW_:          "srw.",
	SRAWI:         "srawi",
	SRAWI_:        "srawi.",
	SRAW:          "sraw",
	SRAW_:         "sraw.",
	SLD:           "sld",
	SLD_:          "sld.",
	SRD:           "srd",
	SRD_:          "srd.",
	SRADI:         "sradi",
	SRADI_:        "sradi.",
	SRAD:          "srad",
	SRAD_:         "srad.",
	CDTBCD:        "cdtbcd",
	CBCDTD:        "cbcdtd",
	ADDG6S:        "addg6s",
	MTSPR:         "mtspr",
	MFSPR:         "mfspr",
	MTCRF:         "mtcrf",
	MFCR:          "mfcr",
	MTSLE:         "mtsle",
	MFVSRD:        "mfvsrd",
	MFVSRWZ:       "mfvsrwz",
	MTVSRD:        "mtvsrd",
	MTVSRWA:       "mtvsrwa",
	MTVSRWZ:       "mtvsrwz",
	MTOCRF:        "mtocrf",
	MFOCRF:        "mfocrf",
	MCRXR:         "mcrxr",
	MTDCRUX:       "mtdcrux",
	MFDCRUX:       "mfdcrux",
	LFS:           "lfs",
	LFSU:          "lfsu",
	LFSX:          "lfsx",
	LFSUX:         "lfsux",
	LFD:           "lfd",
	LFDU:          "lfdu",
	LFDX:          "lfdx",
	LFDUX:         "lfdux",
	LFIWAX:        "lfiwax",
	LFIWZX:        "lfiwzx",
	STFS:          "stfs",
	STFSU:         "stfsu",
	STFSX:         "stfsx",
	STFSUX:        "stfsux",
	STFD:          "stfd",
	STFDU:         "stfdu",
	STFDX:         "stfdx",
	STFDUX:        "stfdux",
	STFIWX:        "stfiwx",
	LFDP:          "lfdp",
	LFDPX:         "lfdpx",
	STFDP:         "stfdp",
	STFDPX:        "stfdpx",
	FMR:           "fmr",
	FMR_:          "fmr.",
	FABS:          "fabs",
	FABS_:         "fabs.",
	FNABS:         "fnabs",
	FNABS_:        "fnabs.",
	FNEG:          "fneg",
	FNEG_:         "fneg.",
	FCPSGN:        "fcpsgn",
	FCPSGN_:       "fcpsgn.",
	FMRGEW:        "fmrgew",
	FMRGOW:        "fmrgow",
	FADD:          "fadd",
	FADD_:         "fadd.",
	FADDS:         "fadds",
	FADDS_:        "fadds.",
	FSUB:          "fsub",
	FSUB_:         "fsub.",
	FSUBS:         "fsubs",
	FSUBS_:        "fsubs.",
	FMUL:          "fmul",
	FMUL_:         "fmul.",
	FMULS:         "fmuls",
	FMULS_:        "fmuls.",
	FDIV:          "fdiv",
	FDIV_:         "fdiv.",
	FDIVS:         "fdivs",
	FDIVS_:        "fdivs.",
	FSQRT:         "fsqrt",
	FSQRT_:        "fsqrt.",
	FSQRTS:        "fsqrts",
	FSQRTS_:       "fsqrts.",
	FRE:           "fre",
	FRE_:          "fre.",
	FRES:          "fres",
	FRES_:         "fres.",
	FRSQRTE:       "frsqrte",
	FRSQRTE_:      "frsqrte.",
	FRSQRTES:      "frsqrtes",
	FRSQRTES_:     "frsqrtes.",
	FTDIV:         "ftdiv",
	FTSQRT:        "ftsqrt",
	FMADD:         "fmadd",
	FMADD_:        "fmadd.",
	FMADDS:        "fmadds",
	FMADDS_:       "fmadds.",
	FMSUB:         "fmsub",
	FMSUB_:        "fmsub.",
	FMSUBS:        "fmsubs",
	FMSUBS_:       "fmsubs.",
	FNMADD:        "fnmadd",
	FNMADD_:       "fnmadd.",
	FNMADDS:       "fnmadds",
	FNMADDS_:      "fnmadds.",
	FNMSUB:        "fnmsub",
	FNMSUB_:       "fnmsub.",
	FNMSUBS:       "fnmsubs",
	FNMSUBS_:      "fnmsubs.",
	FRSP:          "frsp",
	FRSP_:         "frsp.",
	FCTID:         "fctid",
	FCTID_:        "fctid.",
	FCTIDZ:        "fctidz",
	FCTIDZ_:       "fctidz.",
	FCTIDU:        "fctidu",
	FCTIDU_:       "fctidu.",
	FCTIDUZ:       "fctiduz",
	FCTIDUZ_:      "fctiduz.",
	FCTIW:         "fctiw",
	FCTIW_:        "fctiw.",
	FCTIWZ:        "fctiwz",
	FCTIWZ_:       "fctiwz.",
	FCTIWU:        "fctiwu",
	FCTIWU_:       "fctiwu.",
	FCTIWUZ:       "fctiwuz",
	FCTIWUZ_:      "fctiwuz.",
	FCFID:         "fcfid",
	FCFID_:        "fcfid.",
	FCFIDU:        "fcfidu",
	FCFIDU_:       "fcfidu.",
	FCFIDS:        "fcfids",
	FCFIDS_:       "fcfids.",
	FCFIDUS:       "fcfidus",
	FCFIDUS_:      "fcfidus.",
	FRIN:          "frin",
	FRIN_:         "frin.",
	FRIZ:          "friz",
	FRIZ_:         "friz.",
	FRIP:          "frip",
	FRIP_:         "frip.",
	FRIM:          "frim",
	FRIM_:         "frim.",
	FCMPU:         "fcmpu",
	FCMPO:         "fcmpo",
	FSEL:          "fsel",
	FSEL_:         "fsel.",
	MFFS:          "mffs",
	MFFS_:         "mffs.",
	MCRFS:         "mcrfs",
	MTFSFI:        "mtfsfi",
	MTFSFI_:       "mtfsfi.",
	MTFSF:         "mtfsf",
	MTFSF_:        "mtfsf.",
	MTFSB0:        "mtfsb0",
	MTFSB0_:       "mtfsb0.",
	MTFSB1:        "mtfsb1",
	MTFSB1_:       "mtfsb1.",
	LVEBX:         "lvebx",
	LVEHX:         "lvehx",
	LVEWX:         "lvewx",
	LVX:           "lvx",
	LVXL:          "lvxl",
	STVEBX:        "stvebx",
	STVEHX:        "stvehx",
	STVEWX:        "stvewx",
	STVX:          "stvx",
	STVXL:         "stvxl",
	LVSL:          "lvsl",
	LVSR:          "lvsr",
	VPKPX:         "vpkpx",
	VPKSDSS:       "vpksdss",
	VPKSDUS:       "vpksdus",
	VPKSHSS:       "vpkshss",
	VPKSHUS:       "vpkshus",
	VPKSWSS:       "vpkswss",
	VPKSWUS:       "vpkswus",
	VPKUDUM:       "vpkudum",
	VPKUDUS:       "vpkudus",
	VPKUHUM:       "vpkuhum",
	VPKUHUS:       "vpkuhus",
	VPKUWUM:       "vpkuwum",
	VPKUWUS:       "vpkuwus",
	VUPKHPX:       "vupkhpx",
	VUPKLPX:       "vupklpx",
	VUPKHSB:       "vupkhsb",
	VUPKHSH:       "vupkhsh",
	VUPKHSW:       "vupkhsw",
	VUPKLSB:       "vupklsb",
	VUPKLSH:       "vupklsh",
	VUPKLSW:       "vupklsw",
	VMRGHB:        "vmrghb",
	VMRGHH:        "vmrghh",
	VMRGLB:        "vmrglb",
	VMRGLH:        "vmrglh",
	VMRGHW:        "vmrghw",
	VMRGLW:        "vmrglw",
	VMRGEW:        "vmrgew",
	VMRGOW:        "vmrgow",
	VSPLTB:        "vspltb",
	VSPLTH:        "vsplth",
	VSPLTW:        "vspltw",
	VSPLTISB:      "vspltisb",
	VSPLTISH:      "vspltish",
	VSPLTISW:      "vspltisw",
	VPERM:         "vperm",
	VSEL:          "vsel",
	VSL:           "vsl",
	VSLDOI:        "vsldoi",
	VSLO:          "vslo",
	VSR:           "vsr",
	VSRO:          "vsro",
	VADDCUW:       "vaddcuw",
	VADDSBS:       "vaddsbs",
	VADDSHS:       "vaddshs",
	VADDSWS:       "vaddsws",
	VADDUBM:       "vaddubm",
	VADDUDM:       "vaddudm",
	VADDUHM:       "vadduhm",
	VADDUWM:       "vadduwm",
	VADDUBS:       "vaddubs",
	VADDUHS:       "vadduhs",
	VADDUWS:       "vadduws",
	VADDUQM:       "vadduqm",
	VADDEUQM:      "vaddeuqm",
	VADDCUQ:       "vaddcuq",
	VADDECUQ:      "vaddecuq",
	VSUBCUW:       "vsubcuw",
	VSUBSBS:       "vsubsbs",
	VSUBSHS:       "vsubshs",
	VSUBSWS:       "vsubsws",
	VSUBUBM:       "vsububm",
	VSUBUDM:       "vsubudm",
	VSUBUHM:       "vsubuhm",
	VSUBUWM:       "vsubuwm",
	VSUBUBS:       "vsububs",
	VSUBUHS:       "vsubuhs",
	VSUBUWS:       "vsubuws",
	VSUBUQM:       "vsubuqm",
	VSUBEUQM:      "vsubeuqm",
	VSUBCUQ:       "vsubcuq",
	VSUBECUQ:      "vsubecuq",
	VMULESB:       "vmulesb",
	VMULEUB:       "vmuleub",
	VMULOSB:       "vmulosb",
	VMULOUB:       "vmuloub",
	VMULESH:       "vmulesh",
	VMULEUH:       "vmuleuh",
	VMULOSH:       "vmulosh",
	VMULOUH:       "vmulouh",
	VMULESW:       "vmulesw",
	VMULEUW:       "vmuleuw",
	VMULOSW:       "vmulosw",
	VMULOUW:       "vmulouw",
	VMULUWM:       "vmuluwm",
	VMHADDSHS:     "vmhaddshs",
	VMHRADDSHS:    "vmhraddshs",
	VMLADDUHM:     "vmladduhm",
	VMSUMUBM:      "vmsumubm",
	VMSUMMBM:      "vmsummbm",
	VMSUMSHM:      "vmsumshm",
	VMSUMSHS:      "vmsumshs",
	VMSUMUHM:      "vmsumuhm",
	VMSUMUHS:      "vmsumuhs",
	VSUMSWS:       "vsumsws",
	VSUM2SWS:      "vsum2sws",
	VSUM4SBS:      "vsum4sbs",
	VSUM4SHS:      "vsum4shs",
	VSUM4UBS:      "vsum4ubs",
	VAVGSB:        "vavgsb",
	VAVGSH:        "vavgsh",
	VAVGSW:        "vavgsw",
	VAVGUB:        "vavgub",
	VAVGUW:        "vavguw",
	VAVGUH:        "vavguh",
	VMAXSB:        "vmaxsb",
	VMAXSD:        "vmaxsd",
	VMAXUB:        "vmaxub",
	VMAXUD:        "vmaxud",
	VMAXSH:        "vmaxsh",
	VMAXSW:        "vmaxsw",
	VMAXUH:        "vmaxuh",
	VMAXUW:        "vmaxuw",
	VMINSB:        "vminsb",
	VMINSD:        "vminsd",
	VMINUB:        "vminub",
	VMINUD:        "vminud",
	VMINSH:        "vminsh",
	VMINSW:        "vminsw",
	VMINUH:        "vminuh",
	VMINUW:        "vminuw",
	VCMPEQUB:      "vcmpequb",
	VCMPEQUB_:     "vcmpequb.",
	VCMPEQUH:      "vcmpequh",
	VCMPEQUH_:     "vcmpequh.",
	VCMPEQUW:      "vcmpequw",
	VCMPEQUW_:     "vcmpequw.",
	VCMPEQUD:      "vcmpequd",
	VCMPEQUD_:     "vcmpequd.",
	VCMPGTSB:      "vcmpgtsb",
	VCMPGTSB_:     "vcmpgtsb.",
	VCMPGTSD:      "vcmpgtsd",
	VCMPGTSD_:     "vcmpgtsd.",
	VCMPGTSH:      "vcmpgtsh",
	VCMPGTSH_:     "vcmpgtsh.",
	VCMPGTSW:      "vcmpgtsw",
	VCMPGTSW_:     "vcmpgtsw.",
	VCMPGTUB:      "vcmpgtub",
	VCMPGTUB_:     "vcmpgtub.",
	VCMPGTUD:      "vcmpgtud",
	VCMPGTUD_:     "vcmpgtud.",
	VCMPGTUH:      "vcmpgtuh",
	VCMPGTUH_:     "vcmpgtuh.",
	VCMPGTUW:      "vcmpgtuw",
	VCMPGTUW_:     "vcmpgtuw.",
	VAND:          "vand",
	VANDC:         "vandc",
	VEQV:          "veqv",
	VNAND:         "vnand",
	VORC:          "vorc",
	VNOR:          "vnor",
	VOR:           "vor",
	VXOR:          "vxor",
	VRLB:          "vrlb",
	VRLH:          "vrlh",
	VRLW:          "vrlw",
	VRLD:          "vrld",
	VSLB:          "vslb",
	VSLH:          "vslh",
	VSLW:          "vslw",
	VSLD:          "vsld",
	VSRB:          "vsrb",
	VSRH:          "vsrh",
	VSRW:          "vsrw",
	VSRD:          "vsrd",
	VSRAB:         "vsrab",
	VSRAH:         "vsrah",
	VSRAW:         "vsraw",
	VSRAD:         "vsrad",
	VADDFP:        "vaddfp",
	VSUBFP:        "vsubfp",
	VMADDFP:       "vmaddfp",
	VNMSUBFP:      "vnmsubfp",
	VMAXFP:        "vmaxfp",
	VMINFP:        "vminfp",
	VCTSXS:        "vctsxs",
	VCTUXS:        "vctuxs",
	VCFSX:         "vcfsx",
	VCFUX:         "vcfux",
	VRFIM:         "vrfim",
	VRFIN:         "vrfin",
	VRFIP:         "vrfip",
	VRFIZ:         "vrfiz",
	VCMPBFP:       "vcmpbfp",
	VCMPBFP_:      "vcmpbfp.",
	VCMPEQFP:      "vcmpeqfp",
	VCMPEQFP_:     "vcmpeqfp.",
	VCMPGEFP:      "vcmpgefp",
	VCMPGEFP_:     "vcmpgefp.",
	VCMPGTFP:      "vcmpgtfp",
	VCMPGTFP_:     "vcmpgtfp.",
	VEXPTEFP:      "vexptefp",
	VLOGEFP:       "vlogefp",
	VREFP:         "vrefp",
	VRSQRTEFP:     "vrsqrtefp",
	VCIPHER:       "vcipher",
	VCIPHERLAST:   "vcipherlast",
	VNCIPHER:      "vncipher",
	VNCIPHERLAST:  "vncipherlast",
	VSBOX:         "vsbox",
	VSHASIGMAD:    "vshasigmad",
	VSHASIGMAW:    "vshasigmaw",
	VPMSUMB:       "vpmsumb",
	VPMSUMD:       "vpmsumd",
	VPMSUMH:       "vpmsumh",
	VPMSUMW:       "vpmsumw",
	VPERMXOR:      "vpermxor",
	VGBBD:         "vgbbd",
	VCLZB:         "vclzb",
	VCLZH:         "vclzh",
	VCLZW:         "vclzw",
	VCLZD:         "vclzd",
	VPOPCNTB:      "vpopcntb",
	VPOPCNTD:      "vpopcntd",
	VPOPCNTH:      "vpopcnth",
	VPOPCNTW:      "vpopcntw",
	VBPERMQ:       "vbpermq",
	BCDADD_:       "bcdadd.",
	BCDSUB_:       "bcdsub.",
	MTVSCR:        "mtvscr",
	MFVSCR:        "mfvscr",
	DADD:          "dadd",
	DADD_:         "dadd.",
	DSUB:          "dsub",
	DSUB_:         "dsub.",
	DMUL:          "dmul",
	DMUL_:         "dmul.",
	DDIV:          "ddiv",
	DDIV_:         "ddiv.",
	DCMPU:         "dcmpu",
	DCMPO:         "dcmpo",
	DTSTDC:        "dtstdc",
	DTSTDG:        "dtstdg",
	DTSTEX:        "dtstex",
	DTSTSF:        "dtstsf",
	DQUAI:         "dquai",
	DQUAI_:        "dquai.",
	DQUA:          "dqua",
	DQUA_:         "dqua.",
	DRRND:         "drrnd",
	DRRND_:        "drrnd.",
	DRINTX:        "drintx",
	DRINTX_:       "drintx.",
	DRINTN:        "drintn",
	DRINTN_:       "drintn.",
	DCTDP:         "dctdp",
	DCTDP_:        "dctdp.",
	DCTQPQ:        "dctqpq",
	DCTQPQ_:       "dctqpq.",
	DRSP:          "drsp",
	DRSP_:         "drsp.",
	DRDPQ:         "drdpq",
	DRDPQ_:        "drdpq.",
	DCFFIX:        "dcffix",
	DCFFIX_:       "dcffix.",
	DCFFIXQ:       "dcffixq",
	DCFFIXQ_:      "dcffixq.",
	DCTFIX:        "dctfix",
	DCTFIX_:       "dctfix.",
	DDEDPD:        "ddedpd",
	DDEDPD_:       "ddedpd.",
	DENBCD:        "denbcd",
	DENBCD_:       "denbcd.",
	DXEX:          "dxex",
	DXEX_:         "dxex.",
	DIEX:          "diex",
	DIEX_:         "diex.",
	DSCLI:         "dscli",
	DSCLI_:        "dscli.",
	DSCRI:         "dscri",
	DSCRI_:        "dscri.",
	LXSDX:         "lxsdx",
	LXSIWAX:       "lxsiwax",
	LXSIWZX:       "lxsiwzx",
	LXSSPX:        "lxsspx",
	LXVD2X:        "lxvd2x",
	LXVDSX:        "lxvdsx",
	LXVW4X:        "lxvw4x",
	STXSDX:        "stxsdx",
	STXSIWX:       "stxsiwx",
	STXSSPX:       "stxsspx",
	STXVD2X:       "stxvd2x",
	STXVW4X:       "stxvw4x",
	XSABSDP:       "xsabsdp",
	XSADDDP:       "xsadddp",
	XSADDSP:       "xsaddsp",
	XSCMPODP:      "xscmpodp",
	XSCMPUDP:      "xscmpudp",
	XSCPSGNDP:     "xscpsgndp",
	XSCVDPSP:      "xscvdpsp",
	XSCVDPSPN:     "xscvdpspn",
	XSCVDPSXDS:    "xscvdpsxds",
	XSCVDPSXWS:    "xscvdpsxws",
	XSCVDPUXDS:    "xscvdpuxds",
	XSCVDPUXWS:    "xscvdpuxws",
	XSCVSPDP:      "xscvspdp",
	XSCVSPDPN:     "xscvspdpn",
	XSCVSXDDP:     "xscvsxddp",
	XSCVSXDSP:     "xscvsxdsp",
	XSCVUXDDP:     "xscvuxddp",
	XSCVUXDSP:     "xscvuxdsp",
	XSDIVDP:       "xsdivdp",
	XSDIVSP:       "xsdivsp",
	XSMADDADP:     "xsmaddadp",
	XSMADDASP:     "xsmaddasp",
	XSMAXDP:       "xsmaxdp",
	XSMINDP:       "xsmindp",
	XSMSUBADP:     "xsmsubadp",
	XSMSUBASP:     "xsmsubasp",
	XSMULDP:       "xsmuldp",
	XSMULSP:       "xsmulsp",
	XSNABSDP:      "xsnabsdp",
	XSNEGDP:       "xsnegdp",
	XSNMADDADP:    "xsnmaddadp",
	XSNMADDASP:    "xsnmaddasp",
	XSNMSUBADP:    "xsnmsubadp",
	XSNMSUBASP:    "xsnmsubasp",
	XSRDPI:        "xsrdpi",
	XSRDPIC:       "xsrdpic",
	XSRDPIM:       "xsrdpim",
	XSRDPIP:       "xsrdpip",
	XSRDPIZ:       "xsrdpiz",
	XSREDP:        "xsredp",
	XSRESP:        "xsresp",
	XSRSP:         "xsrsp",
	XSRSQRTEDP:    "xsrsqrtedp",
	XSRSQRTESP:    "xsrsqrtesp",
	XSSQRTDP:      "xssqrtdp",
	XSSQRTSP:      "xssqrtsp",
	XSSUBDP:       "xssubdp",
	XSSUBSP:       "xssubsp",
	XSTDIVDP:      "xstdivdp",
	XSTSQRTDP:     "xstsqrtdp",
	XVABSDP:       "xvabsdp",
	XVABSSP:       "xvabssp",
	XVADDDP:       "xvadddp",
	XVADDSP:       "xvaddsp",
	XVCMPEQDP:     "xvcmpeqdp",
	XVCMPEQDP_:    "xvcmpeqdp.",
	XVCMPEQSP:     "xvcmpeqsp",
	XVCMPEQSP_:    "xvcmpeqsp.",
	XVCMPGEDP:     "xvcmpgedp",
	XVCMPGEDP_:    "xvcmpgedp.",
	XVCMPGESP:     "xvcmpgesp",
	XVCMPGESP_:    "xvcmpgesp.",
	XVCMPGTDP:     "xvcmpgtdp",
	XVCMPGTDP_:    "xvcmpgtdp.",
	XVCMPGTSP:     "xvcmpgtsp",
	XVCMPGTSP_:    "xvcmpgtsp.",
	XVCPSGNDP:     "xvcpsgndp",
	XVCPSGNSP:     "xvcpsgnsp",
	XVCVDPSP:      "xvcvdpsp",
	XVCVDPSXDS:    "xvcvdpsxds",
	XVCVDPSXWS:    "xvcvdpsxws",
	XVCVDPUXDS:    "xvcvdpuxds",
	XVCVDPUXWS:    "xvcvdpuxws",
	XVCVSPDP:      "xvcvspdp",
	XVCVSPSXDS:    "xvcvspsxds",
	XVCVSPSXWS:    "xvcvspsxws",
	XVCVSPUXDS:    "xvcvspuxds",
	XVCVSPUXWS:    "xvcvspuxws",
	XVCVSXDDP:     "xvcvsxddp",
	XVCVSXDSP:     "xvcvsxdsp",
	XVCVSXWDP:     "xvcvsxwdp",
	XVCVSXWSP:     "xvcvsxwsp",
	XVCVUXDDP:     "xvcvuxddp",
	XVCVUXDSP:     "xvcvuxdsp",
	XVCVUXWDP:     "xvcvuxwdp",
	XVCVUXWSP:     "xvcvuxwsp",
	XVDIVDP:       "xvdivdp",
	XVDIVSP:       "xvdivsp",
	XVMADDADP:     "xvmaddadp",
	XVMADDASP:     "xvmaddasp",
	XVMAXDP:       "xvmaxdp",
	XVMAXSP:       "xvmaxsp",
	XVMINDP:       "xvmindp",
	XVMINSP:       "xvminsp",
	XVMSUBADP:     "xvmsubadp",
	XVMSUBASP:     "xvmsubasp",
	XVMULDP:       "xvmuldp",
	XVMULSP:       "xvmulsp",
	XVNABSDP:      "xvnabsdp",
	XVNABSSP:      "xvnabssp",
	XVNEGDP:       "xvnegdp",
	XVNEGSP:       "xvnegsp",
	XVNMADDADP:    "xvnmaddadp",
	XVNMADDASP:    "xvnmaddasp",
	XVNMSUBADP:    "xvnmsubadp",
	XVNMSUBASP:    "xvnmsubasp",
	XVRDPI:        "xvrdpi",
	XVRDPIC:       "xvrdpic",
	XVRDPIM:       "xvrdpim",
	XVRDPIP:       "xvrdpip",
	XVRDPIZ:       "xvrdpiz",
	XVREDP:        "xvredp",
	XVRESP:        "xvresp",
	XVRSPI:        "xvrspi",
	XVRSPIC:       "xvrspic",
	XVRSPIM:       "xvrspim",
	XVRSPIP:       "xvrspip",
	XVRSPIZ:       "xvrspiz",
	XVRSQRTEDP:    "xvrsqrtedp",
	XVRSQRTESP:    "xvrsqrtesp",
	XVSQRTDP:      "xvsqrtdp",
	XVSQRTSP:      "xvsqrtsp",
	XVSUBDP:       "xvsubdp",
	XVSUBSP:       "xvsubsp",
	XVTDIVDP:      "xvtdivdp",
	XVTDIVSP:      "xvtdivsp",
	XVTSQRTDP:     "xvtsqrtdp",
	XVTSQRTSP:     "xvtsqrtsp",
	XXLAND:        "xxland",
	XXLANDC:       "xxlandc",
	XXLEQV:        "xxleqv",
	XXLNAND:       "xxlnand",
	XXLORC:        "xxlorc",
	XXLNOR:        "xxlnor",
	XXLOR:         "xxlor",
	XXLXOR:        "xxlxor",
	XXMRGHW:       "xxmrghw",
	XXMRGLW:       "xxmrglw",
	XXPERMDI:      "xxpermdi",
	XXSEL:         "xxsel",
	XXSLDWI:       "xxsldwi",
	XXSPLTW:       "xxspltw",
	BRINC:         "brinc",
	EVABS:         "evabs",
	EVADDIW:       "evaddiw",
	EVADDSMIAAW:   "evaddsmiaaw",
	EVADDSSIAAW:   "evaddssiaaw",
	EVADDUMIAAW:   "evaddumiaaw",
	EVADDUSIAAW:   "evaddusiaaw",
	EVADDW:        "evaddw",
	EVAND:         "evand",
	EVCMPEQ:       "evcmpeq",
	EVANDC:        "evandc",
	EVCMPGTS:      "evcmpgts",
	EVCMPGTU:      "evcmpgtu",
	EVCMPLTU:      "evcmpltu",
	EVCMPLTS:      "evcmplts",
	EVCNTLSW:      "evcntlsw",
	EVCNTLZW:      "evcntlzw",
	EVDIVWS:       "evdivws",
	EVDIVWU:       "evdivwu",
	EVEQV:         "eveqv",
	EVEXTSB:       "evextsb",
	EVEXTSH:       "evextsh",
	EVLDD:         "evldd",
	EVLDH:         "evldh",
	EVLDDX:        "evlddx",
	EVLDHX:        "evldhx",
	EVLDW:         "evldw",
	EVLHHESPLAT:   "evlhhesplat",
	EVLDWX:        "evldwx",
	EVLHHESPLATX:  "evlhhesplatx",
	EVLHHOSSPLAT:  "evlhhossplat",
	EVLHHOUSPLAT:  "evlhhousplat",
	EVLHHOSSPLATX: "evlhhossplatx",
	EVLHHOUSPLATX: "evlhhousplatx",
	EVLWHE:        "evlwhe",
	EVLWHOS:       "evlwhos",
	EVLWHEX:       "evlwhex",
	EVLWHOSX:      "evlwhosx",
	EVLWHOU:       "evlwhou",
	EVLWHSPLAT:    "evlwhsplat",
	EVLWHOUX:      "evlwhoux",
	EVLWHSPLATX:   "evlwhsplatx",
	EVLWWSPLAT:    "evlwwsplat",
	EVMERGEHI:     "evmergehi",
	EVLWWSPLATX:   "evlwwsplatx",
	EVMERGELO:     "evmergelo",
	EVMERGEHILO:   "evmergehilo",
	EVMHEGSMFAA:   "evmhegsmfaa",
	EVMERGELOHI:   "evmergelohi",
	EVMHEGSMFAN:   "evmhegsmfan",
	EVMHEGSMIAA:   "evmhegsmiaa",
	EVMHEGUMIAA:   "evmhegumiaa",
	EVMHEGSMIAN:   "evmhegsmian",
	EVMHEGUMIAN:   "evmhegumian",
	EVMHESMF:      "evmhesmf",
	EVMHESMFAAW:   "evmhesmfaaw",
	EVMHESMFA:     "evmhesmfa",
	EVMHESMFANW:   "evmhesmfanw",
	EVMHESMI:      "evmhesmi",
	EVMHESMIAAW:   "evmhesmiaaw",
	EVMHESMIA:     "evmhesmia",
	EVMHESMIANW:   "evmhesmianw",
	EVMHESSF:      "evmhessf",
	EVMHESSFA:     "evmhessfa",
	EVMHESSFAAW:   "evmhessfaaw",
	EVMHESSFANW:   "evmhessfanw",
	EVMHESSIAAW:   "evmhessiaaw",
	EVMHESSIANW:   "evmhessianw",
	EVMHEUMI:      "evmheumi",
	EVMHEUMIAAW:   "evmheumiaaw",
	EVMHEUMIA:     "evmheumia",
	EVMHEUMIANW:   "evmheumianw",
	EVMHEUSIAAW:   "evmheusiaaw",
	EVMHEUSIANW:   "evmheusianw",
	EVMHOGSMFAA:   "evmhogsmfaa",
	EVMHOGSMIAA:   "evmhogsmiaa",
	EVMHOGSMFAN:   "evmhogsmfan",
	EVMHOGSMIAN:   "evmhogsmian",
	EVMHOGUMIAA:   "evmhogumiaa",
	EVMHOSMF:      "evmhosmf",
	EVMHOGUMIAN:   "evmhogumian",
	EVMHOSMFA:     "evmhosmfa",
	EVMHOSMFAAW:   "evmhosmfaaw",
	EVMHOSMI:      "evmhosmi",
	EVMHOSMFANW:   "evmhosmfanw",
	EVMHOSMIA:     "evmhosmia",
	EVMHOSMIAAW:   "evmhosmiaaw",
	EVMHOSMIANW:   "evmhosmianw",
	EVMHOSSF:      "evmhossf",
	EVMHOSSFA:     "evmhossfa",
	EVMHOSSFAAW:   "evmhossfaaw",
	EVMHOSSFANW:   "evmhossfanw",
	EVMHOSSIAAW:   "evmhossiaaw",
	EVMHOUMI:      "evmhoumi",
	EVMHOSSIANW:   "evmhossianw",
	EVMHOUMIA:     "evmhoumia",
	EVMHOUMIAAW:   "evmhoumiaaw",
	EVMHOUSIAAW:   "evmhousiaaw",
	EVMHOUMIANW:   "evmhoumianw",
	EVMHOUSIANW:   "evmhousianw",
	EVMRA:         "evmra",
	EVMWHSMF:      "evmwhsmf",
	EVMWHSMI:      "evmwhsmi",
	EVMWHSMFA:     "evmwhsmfa",
	EVMWHSMIA:     "evmwhsmia",
	EVMWHSSF:      "evmwhssf",
	EVMWHUMI:      "evmwhumi",
	EVMWHSSFA:     "evmwhssfa",
	EVMWHUMIA:     "evmwhumia",
	EVMWLSMIAAW:   "evmwlsmiaaw",
	EVMWLSSIAAW:   "evmwlssiaaw",
	EVMWLSMIANW:   "evmwlsmianw",
	EVMWLSSIANW:   "evmwlssianw",
	EVMWLUMI:      "evmwlumi",
	EVMWLUMIAAW:   "evmwlumiaaw",
	EVMWLUMIA:     "evmwlumia",
	EVMWLUMIANW:   "evmwlumianw",
	EVMWLUSIAAW:   "evmwlusiaaw",
	EVMWSMF:       "evmwsmf",
	EVMWLUSIANW:   "evmwlusianw",
	EVMWSMFA:      "evmwsmfa",
	EVMWSMFAA:     "evmwsmfaa",
	EVMWSMI:       "evmwsmi",
	EVMWSMIAA:     "evmwsmiaa",
	EVMWSMFAN:     "evmwsmfan",
	EVMWSMIA:      "evmwsmia",
	EVMWSMIAN:     "evmwsmian",
	EVMWSSF:       "evmwssf",
	EVMWSSFA:      "evmwssfa",
	EVMWSSFAA:     "evmwssfaa",
	EVMWUMI:       "evmwumi",
	EVMWSSFAN:     "evmwssfan",
	EVMWUMIA:      "evmwumia",
	EVMWUMIAA:     "evmwumiaa",
	EVNAND:        "evnand",
	EVMWUMIAN:     "evmwumian",
	EVNEG:         "evneg",
	EVNOR:         "evnor",
	EVORC:         "evorc",
	EVOR:          "evor",
	EVRLW:         "evrlw",
	EVRLWI:        "evrlwi",
	EVSEL:         "evsel",
	EVRNDW:        "evrndw",
	EVSLW:         "evslw",
	EVSPLATFI:     "evsplatfi",
	EVSRWIS:       "evsrwis",
	EVSLWI:        "evslwi",
	EVSPLATI:      "evsplati",
	EVSRWIU:       "evsrwiu",
	EVSRWS:        "evsrws",
	EVSTDD:        "evstdd",
	EVSRWU:        "evsrwu",
	EVSTDDX:       "evstddx",
	EVSTDH:        "evstdh",
	EVSTDW:        "evstdw",
	EVSTDHX:       "evstdhx",
	EVSTDWX:       "evstdwx",
	EVSTWHE:       "evstwhe",
	EVSTWHO:       "evstwho",
	EVSTWWE:       "evstwwe",
	EVSTWHEX:      "evstwhex",
	EVSTWHOX:      "evstwhox",
	EVSTWWEX:      "evstwwex",
	EVSTWWO:       "evstwwo",
	EVSUBFSMIAAW:  "evsubfsmiaaw",
	EVSTWWOX:      "evstwwox",
	EVSUBFSSIAAW:  "evsubfssiaaw",
	EVSUBFUMIAAW:  "evsubfumiaaw",
	EVSUBFUSIAAW:  "evsubfusiaaw",
	EVSUBFW:       "evsubfw",
	EVSUBIFW:      "evsubifw",
	EVXOR:         "evxor",
	EVFSABS:       "evfsabs",
	EVFSNABS:      "evfsnabs",
	EVFSNEG:       "evfsneg",
	EVFSADD:       "evfsadd",
	EVFSMUL:       "evfsmul",
	EVFSSUB:       "evfssub",
	EVFSDIV:       "evfsdiv",
	EVFSCMPGT:     "evfscmpgt",
	EVFSCMPLT:     "evfscmplt",
	EVFSCMPEQ:     "evfscmpeq",
	EVFSTSTGT:     "evfststgt",
	EVFSTSTLT:     "evfststlt",
	EVFSTSTEQ:     "evfststeq",
	EVFSCFSI:      "evfscfsi",
	EVFSCFSF:      "evfscfsf",
	EVFSCFUI:      "evfscfui",
	EVFSCFUF:      "evfscfuf",
	EVFSCTSI:      "evfsctsi",
	EVFSCTUI:      "evfsctui",
	EVFSCTSIZ:     "evfsctsiz",
	EVFSCTUIZ:     "evfsctuiz",
	EVFSCTSF:      "evfsctsf",
	EVFSCTUF:      "evfsctuf",
	EFSABS:        "efsabs",
	EFSNEG:        "efsneg",
	EFSNABS:       "efsnabs",
	EFSADD:        "efsadd",
	EFSMUL:        "efsmul",
	EFSSUB:        "efssub",
	EFSDIV:        "efsdiv",
	EFSCMPGT:      "efscmpgt",
	EFSCMPLT:      "efscmplt",
	EFSCMPEQ:      "efscmpeq",
	EFSTSTGT:      "efststgt",
	EFSTSTLT:      "efststlt",
	EFSTSTEQ:      "efststeq",
	EFSCFSI:       "efscfsi",
	EFSCFSF:       "efscfsf",
	EFSCTSI:       "efsctsi",
	EFSCFUI:       "efscfui",
	EFSCFUF:       "efscfuf",
	EFSCTUI:       "efsctui",
	EFSCTSIZ:      "efsctsiz",
	EFSCTSF:       "efsctsf",
	EFSCTUIZ:      "efsctuiz",
	EFSCTUF:       "efsctuf",
	EFDABS:        "efdabs",
	EFDNEG:        "efdneg",
	EFDNABS:       "efdnabs",
	EFDADD:        "efdadd",
	EFDMUL:        "efdmul",
	EFDSUB:        "efdsub",
	EFDDIV:        "efddiv",
	EFDCMPGT:      "efdcmpgt",
	EFDCMPEQ:      "efdcmpeq",
	EFDCMPLT:      "efdcmplt",
	EFDTSTGT:      "efdtstgt",
	EFDTSTLT:      "efdtstlt",
	EFDCFSI:       "efdcfsi",
	EFDTSTEQ:      "efdtsteq",
	EFDCFUI:       "efdcfui",
	EFDCFSID:      "efdcfsid",
	EFDCFSF:       "efdcfsf",
	EFDCFUF:       "efdcfuf",
	EFDCFUID:      "efdcfuid",
	EFDCTSI:       "efdctsi",
	EFDCTUI:       "efdctui",
	EFDCTSIDZ:     "efdctsidz",
	EFDCTUIDZ:     "efdctuidz",
	EFDCTSIZ:      "efdctsiz",
	EFDCTSF:       "efdctsf",
	EFDCTUF:       "efdctuf",
	EFDCTUIZ:      "efdctuiz",
	EFDCFS:        "efdcfs",
	EFSCFD:        "efscfd",
	DLMZB:         "dlmzb",
	DLMZB_:        "dlmzb.",
	MACCHW:        "macchw",
	MACCHW_:       "macchw.",
	MACCHWO:       "macchwo",
	MACCHWO_:      "macchwo.",
	MACCHWS:       "macchws",
	MACCHWS_:      "macchws.",
	MACCHWSO:      "macchwso",
	MACCHWSO_:     "macchwso.",
	MACCHWU:       "macchwu",
	MACCHWU_:      "macchwu.",
	MACCHWUO:      "macchwuo",
	MACCHWUO_:     "macchwuo.",
	MACCHWSU:      "macchwsu",
	MACCHWSU_:     "macchwsu.",
	MACCHWSUO:     "macchwsuo",
	MACCHWSUO_:    "macchwsuo.",
	MACHHW:        "machhw",
	MACHHW_:       "machhw.",
	MACHHWO:       "machhwo",
	MACHHWO_:      "machhwo.",
	MACHHWS:       "machhws",
	MACHHWS_:      "machhws.",
	MACHHWSO:      "machhwso",
	MACHHWSO_:     "machhwso.",
	MACHHWU:       "machhwu",
	MACHHWU_:      "machhwu.",
	MACHHWUO:      "machhwuo",
	MACHHWUO_:     "machhwuo.",
	MACHHWSU:      "machhwsu",
	MACHHWSU_:     "machhwsu.",
	MACHHWSUO:     "machhwsuo",
	MACHHWSUO_:    "machhwsuo.",
	MACLHW:        "maclhw",
	MACLHW_:       "maclhw.",
	MACLHWO:       "maclhwo",
	MACLHWO_:      "maclhwo.",
	MACLHWS:       "maclhws",
	MACLHWS_:      "maclhws.",
	MACLHWSO:      "maclhwso",
	MACLHWSO_:     "maclhwso.",
	MACLHWU:       "maclhwu",
	MACLHWU_:      "maclhwu.",
	MACLHWUO:      "maclhwuo",
	MACLHWUO_:     "maclhwuo.",
	MULCHW:        "mulchw",
	MULCHW_:       "mulchw.",
	MACLHWSU:      "maclhwsu",
	MACLHWSU_:     "maclhwsu.",
	MACLHWSUO:     "maclhwsuo",
	MACLHWSUO_:    "maclhwsuo.",
	MULCHWU:       "mulchwu",
	MULCHWU_:      "mulchwu.",
	MULHHW:        "mulhhw",
	MULHHW_:       "mulhhw.",
	MULLHW:        "mullhw",
	MULLHW_:       "mullhw.",
	MULHHWU:       "mulhhwu",
	MULHHWU_:      "mulhhwu.",
	MULLHWU:       "mullhwu",
	MULLHWU_:      "mullhwu.",
	NMACCHW:       "nmacchw",
	NMACCHW_:      "nmacchw.",
	NMACCHWO:      "nmacchwo",
	NMACCHWO_:     "nmacchwo.",
	NMACCHWS:      "nmacchws",
	NMACCHWS_:     "nmacchws.",
	NMACCHWSO:     "nmacchwso",
	NMACCHWSO_:    "nmacchwso.",
	NMACHHW:       "nmachhw",
	NMACHHW_:      "nmachhw.",
	NMACHHWO:      "nmachhwo",
	NMACHHWO_:     "nmachhwo.",
	NMACHHWS:      "nmachhws",
	NMACHHWS_:     "nmachhws.",
	NMACHHWSO:     "nmachhwso",
	NMACHHWSO_:    "nmachhwso.",
	NMACLHW:       "nmaclhw",
	NMACLHW_:      "nmaclhw.",
	NMACLHWO:      "nmaclhwo",
	NMACLHWO_:     "nmaclhwo.",
	NMACLHWS:      "nmaclhws",
	NMACLHWS_:     "nmaclhws.",
	NMACLHWSO:     "nmaclhwso",
	NMACLHWSO_:    "nmaclhwso.",
	ICBI:          "icbi",
	ICBT:          "icbt",
	DCBA:          "dcba",
	DCBT:          "dcbt",
	DCBTST:        "dcbtst",
	DCBZ:          "dcbz",
	DCBST:         "dcbst",
	DCBF:          "dcbf",
	ISYNC:         "isync",
	LBARX:         "lbarx",
	LHARX:         "lharx",
	LWARX:         "lwarx",
	STBCX_:        "stbcx.",
	STHCX_:        "sthcx.",
	STWCX_:        "stwcx.",
	LDARX:         "ldarx",
	STDCX_:        "stdcx.",
	LQARX:         "lqarx",
	STQCX_:        "stqcx.",
	SYNC:          "sync",
	EIEIO:         "eieio",
	MBAR:          "mbar",
	WAIT:          "wait",
	TBEGIN_:       "tbegin.",
	TEND_:         "tend.",
	TABORT_:       "tabort.",
	TABORTWC_:     "tabortwc.",
	TABORTWCI_:    "tabortwci.",
	TABORTDC_:     "tabortdc.",
	TABORTDCI_:    "tabortdci.",
	TSR_:          "tsr.",
	TCHECK:        "tcheck",
	MFTB:          "mftb",
	RFEBB:         "rfebb",
	LBDX:          "lbdx",
	LHDX:          "lhdx",
	LWDX:          "lwdx",
	LDDX:          "lddx",
	LFDDX:         "lfddx",
	STBDX:         "stbdx",
	STHDX:         "sthdx",
	STWDX:         "stwdx",
	STDDX:         "stddx",
	STFDDX:        "stfddx",
	DSN:           "dsn",
	ECIWX:         "eciwx",
	ECOWX:         "ecowx",
	RFID:          "rfid",
	HRFID:         "hrfid",
	DOZE:          "doze",
	NAP:           "nap",
	SLEEP:         "sleep",
	RVWINKLE:      "rvwinkle",
	LBZCIX:        "lbzcix",
	LWZCIX:        "lwzcix",
	LHZCIX:        "lhzcix",
	LDCIX:         "ldcix",
	STBCIX:        "stbcix",
	STWCIX:        "stwcix",
	STHCIX:        "sthcix",
	STDCIX:        "stdcix",
	TRECLAIM_:     "treclaim.",
	TRECHKPT_:     "trechkpt.",
	MTMSR:         "mtmsr",
	MTMSRD:        "mtmsrd",
	MFMSR:         "mfmsr",
	SLBIE:         "slbie",
	SLBIA:         "slbia",
	SLBMTE:        "slbmte",
	SLBMFEV:       "slbmfev",
	SLBMFEE:       "slbmfee",
	SLBFEE_:       "slbfee.",
	MTSR:          "mtsr",
	MTSRIN:        "mtsrin",
	MFSR:          "mfsr",
	MFSRIN:        "mfsrin",
	TLBIE:         "tlbie",
	TLBIEL:        "tlbiel",
	TLBIA:         "tlbia",
	TLBSYNC:       "tlbsync",
	MSGSND:        "msgsnd",
	MSGCLR:        "msgclr",
	MSGSNDP:       "msgsndp",
	MSGCLRP:       "msgclrp",
	MTTMR:         "mttmr",
	RFI:           "rfi",
	RFCI:          "rfci",
	RFDI:          "rfdi",
	RFMCI:         "rfmci",
	RFGI:          "rfgi",
	EHPRIV:        "ehpriv",
	MTDCR:         "mtdcr",
	MTDCRX:        "mtdcrx",
	MFDCR:         "mfdcr",
	MFDCRX:        "mfdcrx",
	WRTEE:         "wrtee",
	WRTEEI:        "wrteei",
	LBEPX:         "lbepx",
	LHEPX:         "lhepx",
	LWEPX:         "lwepx",
	LDEPX:         "ldepx",
	STBEPX:        "stbepx",
	STHEPX:        "sthepx",
	STWEPX:        "stwepx",
	STDEPX:        "stdepx",
	DCBSTEP:       "dcbstep",
	DCBTEP:        "dcbtep",
	DCBFEP:        "dcbfep",
	DCBTSTEP:      "dcbtstep",
	ICBIEP:        "icbiep",
	DCBZEP:        "dcbzep",
	LFDEPX:        "lfdepx",
	STFDEPX:       "stfdepx",
	EVLDDEPX:      "evlddepx",
	EVSTDDEPX:     "evstddepx",
	LVEPX:         "lvepx",
	LVEPXL:        "lvepxl",
	STVEPX:        "stvepx",
	STVEPXL:       "stvepxl",
	DCBI:          "dcbi",
	DCBLQ_:        "dcblq.",
	ICBLQ_:        "icblq.",
	DCBTLS:        "dcbtls",
	DCBTSTLS:      "dcbtstls",
	ICBTLS:        "icbtls",
	ICBLC:         "icblc",
	DCBLC:         "dcblc",
	TLBIVAX:       "tlbivax",
	TLBILX:        "tlbilx",
	TLBSX:         "tlbsx",
	TLBSRX_:       "tlbsrx.",
	TLBRE:         "tlbre",
	TLBWE:         "tlbwe",
	DNH:           "dnh",
	DCI:           "dci",
	ICI:           "ici",
	DCREAD:        "dcread",
	ICREAD:        "icread",
	MFPMR:         "mfpmr",
	MTPMR:         "mtpmr",
}

var (
	ap_Reg_11_15               = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_Reg_6_10                = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_PCRel_6_29_shift2       = &argField{Type: TypePCRel, Shift: 2, BitFields: BitFields{{6, 24}}}
	ap_Label_6_29_shift2       = &argField{Type: TypeLabel, Shift: 2, BitFields: BitFields{{6, 24}}}
	ap_ImmUnsigned_6_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_CondRegBit_11_15        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_PCRel_16_29_shift2      = &argField{Type: TypePCRel, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_Label_16_29_shift2      = &argField{Type: TypeLabel, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_ImmUnsigned_19_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{19, 2}}}
	ap_CondRegBit_6_10         = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_CondRegBit_16_20        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_CondRegField_6_8        = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{6, 3}}}
	ap_CondRegField_11_13      = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{11, 3}}}
	ap_ImmUnsigned_20_26       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{20, 7}}}
	ap_SpReg_11_20             = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{11, 10}}}
	ap_Offset_16_31            = &argField{Type: TypeOffset, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_Reg_16_20               = &argField{Type: TypeReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_Offset_16_29_shift2     = &argField{Type: TypeOffset, Shift: 2, BitFields: BitFields{{16, 14}}}
	ap_Offset_16_27_shift4     = &argField{Type: TypeOffset, Shift: 4, BitFields: BitFields{{16, 12}}}
	ap_ImmUnsigned_16_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmSigned_16_31         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_ImmUnsigned_16_31       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 16}}}
	ap_CondRegBit_21_25        = &argField{Type: TypeCondRegBit, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_21_25       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_26_30       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{26, 5}}}
	ap_ImmUnsigned_30_30_16_20 = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{30, 1}, {16, 5}}}
	ap_ImmUnsigned_26_26_21_25 = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{26, 1}, {21, 5}}}
	ap_SpReg_16_20_11_15       = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{16, 5}, {11, 5}}}
	ap_ImmUnsigned_12_19       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{12, 8}}}
	ap_ImmUnsigned_10_10       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{10, 1}}}
	ap_VecSReg_31_31_6_10      = &argField{Type: TypeVecSReg, Shift: 0, BitFields: BitFields{{31, 1}, {6, 5}}}
	ap_FPReg_6_10              = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_FPReg_16_20             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_FPReg_11_15             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_FPReg_21_25             = &argField{Type: TypeFPReg, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_16_19       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 4}}}
	ap_ImmUnsigned_15_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{15, 1}}}
	ap_ImmUnsigned_7_14        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{7, 8}}}
	ap_ImmUnsigned_6_6         = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 1}}}
	ap_VecReg_6_10             = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{6, 5}}}
	ap_VecReg_11_15            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_VecReg_16_20            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmUnsigned_12_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{12, 4}}}
	ap_ImmUnsigned_13_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{13, 3}}}
	ap_ImmUnsigned_14_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{14, 2}}}
	ap_ImmSigned_11_15         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_VecReg_21_25            = &argField{Type: TypeVecReg, Shift: 0, BitFields: BitFields{{21, 5}}}
	ap_ImmUnsigned_22_25       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 4}}}
	ap_ImmUnsigned_11_15       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 5}}}
	ap_ImmUnsigned_16_16       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 1}}}
	ap_ImmUnsigned_17_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{17, 4}}}
	ap_ImmUnsigned_22_22       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 1}}}
	ap_ImmUnsigned_16_21       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{16, 6}}}
	ap_ImmUnsigned_21_22       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{21, 2}}}
	ap_ImmUnsigned_11_12       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 2}}}
	ap_ImmUnsigned_11_11       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 1}}}
	ap_VecSReg_30_30_16_20     = &argField{Type: TypeVecSReg, Shift: 0, BitFields: BitFields{{30, 1}, {16, 5}}}
	ap_VecSReg_29_29_11_15     = &argField{Type: TypeVecSReg, Shift: 0, BitFields: BitFields{{29, 1}, {11, 5}}}
	ap_ImmUnsigned_22_23       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{22, 2}}}
	ap_VecSReg_28_28_21_25     = &argField{Type: TypeVecSReg, Shift: 0, BitFields: BitFields{{28, 1}, {21, 5}}}
	ap_CondRegField_29_31      = &argField{Type: TypeCondRegField, Shift: 0, BitFields: BitFields{{29, 3}}}
	ap_ImmUnsigned_7_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{7, 4}}}
	ap_ImmUnsigned_9_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{9, 2}}}
	ap_ImmUnsigned_31_31       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{31, 1}}}
	ap_ImmSigned_16_20         = &argField{Type: TypeImmSigned, Shift: 0, BitFields: BitFields{{16, 5}}}
	ap_ImmUnsigned_20_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{20, 1}}}
	ap_ImmUnsigned_8_10        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{8, 3}}}
	ap_SpReg_12_15             = &argField{Type: TypeSpReg, Shift: 0, BitFields: BitFields{{12, 4}}}
	ap_ImmUnsigned_6_20        = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{6, 15}}}
	ap_ImmUnsigned_11_20       = &argField{Type: TypeImmUnsigned, Shift: 0, BitFields: BitFields{{11, 10}}}
)

var instFormats = [...]instFormat{
	{CNTLZW, 0xfc0007ff, 0x7c000034, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZW_, 0xfc0007ff, 0x7c000035, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{B, 0xfc000003, 0x48000000, 0x0,
		[5]*argField{ap_PCRel_6_29_shift2}},
	{BA, 0xfc000003, 0x48000002, 0x0,
		[5]*argField{ap_Label_6_29_shift2}},
	{BL, 0xfc000003, 0x48000001, 0x0,
		[5]*argField{ap_PCRel_6_29_shift2}},
	{BLA, 0xfc000003, 0x48000003, 0x0,
		[5]*argField{ap_Label_6_29_shift2}},
	{BC, 0xfc000003, 0x40000000, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_PCRel_16_29_shift2}},
	{BCA, 0xfc000003, 0x40000002, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_Label_16_29_shift2}},
	{BCL, 0xfc000003, 0x40000001, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_PCRel_16_29_shift2}},
	{BCLA, 0xfc000003, 0x40000003, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_Label_16_29_shift2}},
	{BCLR, 0xfc0007ff, 0x4c000020, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCLRL, 0xfc0007ff, 0x4c000021, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCCTR, 0xfc0007ff, 0x4c000420, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCCTRL, 0xfc0007ff, 0x4c000421, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCTAR, 0xfc0007ff, 0x4c000460, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{BCTARL, 0xfc0007ff, 0x4c000461, 0xe000,
		[5]*argField{ap_ImmUnsigned_6_10, ap_CondRegBit_11_15, ap_ImmUnsigned_19_20}},
	{CRAND, 0xfc0007fe, 0x4c000202, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CROR, 0xfc0007fe, 0x4c000382, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRNAND, 0xfc0007fe, 0x4c0001c2, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRXOR, 0xfc0007fe, 0x4c000182, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRNOR, 0xfc0007fe, 0x4c000042, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRANDC, 0xfc0007fe, 0x4c000102, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{MCRF, 0xfc0007fe, 0x4c000000, 0x63f801,
		[5]*argField{ap_CondRegField_6_8, ap_CondRegField_11_13}},
	{CREQV, 0xfc0007fe, 0x4c000242, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{CRORC, 0xfc0007fe, 0x4c000342, 0x1,
		[5]*argField{ap_CondRegBit_6_10, ap_CondRegBit_11_15, ap_CondRegBit_16_20}},
	{SC, 0xfc000002, 0x44000002, 0x3fff01d,
		[5]*argField{ap_ImmUnsigned_20_26}},
	{CLRBHRB, 0xfc0007fe, 0x7c00035c, 0x3fff801,
		[5]*argField{}},
	{MFBHRBE, 0xfc0007fe, 0x7c00025c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_11_20}},
	{LBZ, 0xfc000000, 0x88000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LBZU, 0xfc000000, 0x8c000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LBZX, 0xfc0007fe, 0x7c0000ae, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LBZUX, 0xfc0007fe, 0x7c0000ee, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZ, 0xfc000000, 0xa0000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHZU, 0xfc000000, 0xa4000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHZX, 0xfc0007fe, 0x7c00022e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZUX, 0xfc0007fe, 0x7c00026e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHA, 0xfc000000, 0xa8000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHAU, 0xfc000000, 0xac000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LHAX, 0xfc0007fe, 0x7c0002ae, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHAUX, 0xfc0007fe, 0x7c0002ee, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZ, 0xfc000000, 0x80000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LWZU, 0xfc000000, 0x84000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LWZX, 0xfc0007fe, 0x7c00002e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZUX, 0xfc0007fe, 0x7c00006e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWA, 0xfc000003, 0xe8000002, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LWAX, 0xfc0007fe, 0x7c0002aa, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWAUX, 0xfc0007fe, 0x7c0002ea, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LD, 0xfc000003, 0xe8000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LDU, 0xfc000003, 0xe8000001, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LDX, 0xfc0007fe, 0x7c00002a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDUX, 0xfc0007fe, 0x7c00006a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STB, 0xfc000000, 0x98000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STBU, 0xfc000000, 0x9c000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STBX, 0xfc0007fe, 0x7c0001ae, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBUX, 0xfc0007fe, 0x7c0001ee, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STH, 0xfc000000, 0xb0000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STHU, 0xfc000000, 0xb4000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STHX, 0xfc0007fe, 0x7c00032e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHUX, 0xfc0007fe, 0x7c00036e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STW, 0xfc000000, 0x90000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STWU, 0xfc000000, 0x94000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STWX, 0xfc0007fe, 0x7c00012e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWUX, 0xfc0007fe, 0x7c00016e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STD, 0xfc000003, 0xf8000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STDU, 0xfc000003, 0xf8000001, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STDX, 0xfc0007fe, 0x7c00012a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDUX, 0xfc0007fe, 0x7c00016a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQ, 0xfc000000, 0xe0000000, 0xf,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_27_shift4, ap_Reg_11_15}},
	{STQ, 0xfc000003, 0xf8000002, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LHBRX, 0xfc0007fe, 0x7c00062c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWBRX, 0xfc0007fe, 0x7c00042c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHBRX, 0xfc0007fe, 0x7c00072c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWBRX, 0xfc0007fe, 0x7c00052c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDBRX, 0xfc0007fe, 0x7c000428, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDBRX, 0xfc0007fe, 0x7c000528, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LMW, 0xfc000000, 0xb8000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STMW, 0xfc000000, 0xbc000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LSWI, 0xfc0007fe, 0x7c0004aa, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{LSWX, 0xfc0007fe, 0x7c00042a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STSWI, 0xfc0007fe, 0x7c0005aa, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{STSWX, 0xfc0007fe, 0x7c00052a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LI, 0xfc1f0000, 0x38000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_16_31}},
	{ADDI, 0xfc000000, 0x38000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{LIS, 0xfc1f0000, 0x3c000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_16_31}},
	{ADDIS, 0xfc000000, 0x3c000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ADD, 0xfc0007ff, 0x7c000214, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADD_, 0xfc0007ff, 0x7c000215, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDO, 0xfc0007ff, 0x7c000614, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDO_, 0xfc0007ff, 0x7c000615, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDIC, 0xfc000000, 0x30000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{SUBF, 0xfc0007ff, 0x7c000050, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBF_, 0xfc0007ff, 0x7c000051, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFO, 0xfc0007ff, 0x7c000450, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFO_, 0xfc0007ff, 0x7c000451, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDIC_, 0xfc000000, 0x34000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{SUBFIC, 0xfc000000, 0x20000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ADDC, 0xfc0007ff, 0x7c000014, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDC_, 0xfc0007ff, 0x7c000015, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDCO, 0xfc0007ff, 0x7c000414, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDCO_, 0xfc0007ff, 0x7c000415, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFC, 0xfc0007ff, 0x7c000010, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFC_, 0xfc0007ff, 0x7c000011, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFCO, 0xfc0007ff, 0x7c000410, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFCO_, 0xfc0007ff, 0x7c000411, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDE, 0xfc0007ff, 0x7c000114, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDE_, 0xfc0007ff, 0x7c000115, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDEO, 0xfc0007ff, 0x7c000514, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDEO_, 0xfc0007ff, 0x7c000515, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ADDME, 0xfc0007ff, 0x7c0001d4, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDME_, 0xfc0007ff, 0x7c0001d5, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDMEO, 0xfc0007ff, 0x7c0005d4, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDMEO_, 0xfc0007ff, 0x7c0005d5, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFE, 0xfc0007ff, 0x7c000110, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFE_, 0xfc0007ff, 0x7c000111, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFEO, 0xfc0007ff, 0x7c000510, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFEO_, 0xfc0007ff, 0x7c000511, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SUBFME, 0xfc0007ff, 0x7c0001d0, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFME_, 0xfc0007ff, 0x7c0001d1, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFMEO, 0xfc0007ff, 0x7c0005d0, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFMEO_, 0xfc0007ff, 0x7c0005d1, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZE, 0xfc0007ff, 0x7c000194, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZE_, 0xfc0007ff, 0x7c000195, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZEO, 0xfc0007ff, 0x7c000594, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{ADDZEO_, 0xfc0007ff, 0x7c000595, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZE, 0xfc0007ff, 0x7c000190, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZE_, 0xfc0007ff, 0x7c000191, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZEO, 0xfc0007ff, 0x7c000590, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{SUBFZEO_, 0xfc0007ff, 0x7c000591, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEG, 0xfc0007ff, 0x7c0000d0, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEG_, 0xfc0007ff, 0x7c0000d1, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEGO, 0xfc0007ff, 0x7c0004d0, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{NEGO_, 0xfc0007ff, 0x7c0004d1, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MULLI, 0xfc000000, 0x1c000000, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{MULLW, 0xfc0007ff, 0x7c0001d6, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLW_, 0xfc0007ff, 0x7c0001d7, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLWO, 0xfc0007ff, 0x7c0005d6, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLWO_, 0xfc0007ff, 0x7c0005d7, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHW, 0xfc0003ff, 0x7c000096, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHW_, 0xfc0003ff, 0x7c000097, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHWU, 0xfc0003ff, 0x7c000016, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHWU_, 0xfc0003ff, 0x7c000017, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVW, 0xfc0007ff, 0x7c0003d6, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVW_, 0xfc0007ff, 0x7c0003d7, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWO, 0xfc0007ff, 0x7c0007d6, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWO_, 0xfc0007ff, 0x7c0007d7, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWU, 0xfc0007ff, 0x7c000396, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWU_, 0xfc0007ff, 0x7c000397, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWUO, 0xfc0007ff, 0x7c000796, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWUO_, 0xfc0007ff, 0x7c000797, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWE, 0xfc0007ff, 0x7c000356, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWE_, 0xfc0007ff, 0x7c000357, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEO, 0xfc0007ff, 0x7c000756, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEO_, 0xfc0007ff, 0x7c000757, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEU, 0xfc0007ff, 0x7c000316, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEU_, 0xfc0007ff, 0x7c000317, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEUO, 0xfc0007ff, 0x7c000716, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVWEUO_, 0xfc0007ff, 0x7c000717, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLD, 0xfc0007ff, 0x7c0001d2, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLD_, 0xfc0007ff, 0x7c0001d3, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLDO, 0xfc0007ff, 0x7c0005d2, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLDO_, 0xfc0007ff, 0x7c0005d3, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHDU, 0xfc0003ff, 0x7c000012, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHDU_, 0xfc0003ff, 0x7c000013, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHD, 0xfc0003ff, 0x7c000092, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHD_, 0xfc0003ff, 0x7c000093, 0x400,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVD, 0xfc0007ff, 0x7c0003d2, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVD_, 0xfc0007ff, 0x7c0003d3, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDO, 0xfc0007ff, 0x7c0007d2, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDO_, 0xfc0007ff, 0x7c0007d3, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDU, 0xfc0007ff, 0x7c000392, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDU_, 0xfc0007ff, 0x7c000393, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDUO, 0xfc0007ff, 0x7c000792, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDUO_, 0xfc0007ff, 0x7c000793, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDE, 0xfc0007ff, 0x7c000352, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDE_, 0xfc0007ff, 0x7c000353, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEO, 0xfc0007ff, 0x7c000752, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEO_, 0xfc0007ff, 0x7c000753, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEU, 0xfc0007ff, 0x7c000312, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEU_, 0xfc0007ff, 0x7c000313, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEUO, 0xfc0007ff, 0x7c000712, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DIVDEUO_, 0xfc0007ff, 0x7c000713, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPWI, 0xfc200000, 0x2c000000, 0x400000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{CMPDI, 0xfc200000, 0x2c200000, 0x400000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{CMPW, 0xfc2007fe, 0x7c000000, 0x400001,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPD, 0xfc2007fe, 0x7c200000, 0x400001,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPLWI, 0xfc200000, 0x28000000, 0x400000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmUnsigned_16_31}},
	{CMPLDI, 0xfc200000, 0x28200000, 0x400000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_ImmUnsigned_16_31}},
	{CMPLW, 0xfc2007fe, 0x7c000040, 0x400001,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{CMPLD, 0xfc2007fe, 0x7c200040, 0x400001,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{TWI, 0xfc000000, 0xc000000, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{TW, 0xfc0007fe, 0x7c000008, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TDI, 0xfc000000, 0x8000000, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_31}},
	{ISEL, 0xfc00003e, 0x7c00001e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_CondRegBit_21_25}},
	{TD, 0xfc0007fe, 0x7c000088, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ANDI_, 0xfc000000, 0x70000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ANDIS_, 0xfc000000, 0x74000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ORI, 0xfc000000, 0x60000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{ORIS, 0xfc000000, 0x64000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{XORI, 0xfc000000, 0x68000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{XORIS, 0xfc000000, 0x6c000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_31}},
	{AND, 0xfc0007ff, 0x7c000038, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{AND_, 0xfc0007ff, 0x7c000039, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{XOR, 0xfc0007ff, 0x7c000278, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{XOR_, 0xfc0007ff, 0x7c000279, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NAND, 0xfc0007ff, 0x7c0003b8, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NAND_, 0xfc0007ff, 0x7c0003b9, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{OR, 0xfc0007ff, 0x7c000378, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{OR_, 0xfc0007ff, 0x7c000379, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NOR, 0xfc0007ff, 0x7c0000f8, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{NOR_, 0xfc0007ff, 0x7c0000f9, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ANDC, 0xfc0007ff, 0x7c000078, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ANDC_, 0xfc0007ff, 0x7c000079, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EXTSB, 0xfc0007ff, 0x7c000774, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSB_, 0xfc0007ff, 0x7c000775, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EQV, 0xfc0007ff, 0x7c000238, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EQV_, 0xfc0007ff, 0x7c000239, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ORC, 0xfc0007ff, 0x7c000338, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{ORC_, 0xfc0007ff, 0x7c000339, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{EXTSH, 0xfc0007ff, 0x7c000734, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSH_, 0xfc0007ff, 0x7c000735, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CMPB, 0xfc0007fe, 0x7c0003f8, 0x1,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{POPCNTB, 0xfc0007fe, 0x7c0000f4, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{POPCNTW, 0xfc0007fe, 0x7c0002f4, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{PRTYD, 0xfc0007fe, 0x7c000174, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{PRTYW, 0xfc0007fe, 0x7c000134, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSW, 0xfc0007ff, 0x7c0007b4, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{EXTSW_, 0xfc0007ff, 0x7c0007b5, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZD, 0xfc0007ff, 0x7c000074, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CNTLZD_, 0xfc0007ff, 0x7c000075, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{POPCNTD, 0xfc0007fe, 0x7c0003f4, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{BPERMD, 0xfc0007fe, 0x7c0001f8, 0x1,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{RLWINM, 0xfc000001, 0x54000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWINM_, 0xfc000001, 0x54000001, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWNM, 0xfc000001, 0x5c000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWNM_, 0xfc000001, 0x5c000001, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWIMI, 0xfc000001, 0x50000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLWIMI_, 0xfc000001, 0x50000001, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_ImmUnsigned_21_25, ap_ImmUnsigned_26_30}},
	{RLDICL, 0xfc00001d, 0x78000000, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICL_, 0xfc00001d, 0x78000001, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICR, 0xfc00001d, 0x78000004, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDICR_, 0xfc00001d, 0x78000005, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIC, 0xfc00001d, 0x78000008, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIC_, 0xfc00001d, 0x78000009, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCL, 0xfc00001f, 0x78000010, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCL_, 0xfc00001f, 0x78000011, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCR, 0xfc00001f, 0x78000012, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDCR_, 0xfc00001f, 0x78000013, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIMI, 0xfc00001d, 0x7800000c, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{RLDIMI_, 0xfc00001d, 0x7800000d, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20, ap_ImmUnsigned_26_26_21_25}},
	{SLW, 0xfc0007ff, 0x7c000030, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLW_, 0xfc0007ff, 0x7c000031, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRW, 0xfc0007ff, 0x7c000430, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRW_, 0xfc0007ff, 0x7c000431, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAWI, 0xfc0007ff, 0x7c000670, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20}},
	{SRAWI_, 0xfc0007ff, 0x7c000671, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_16_20}},
	{SRAW, 0xfc0007ff, 0x7c000630, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAW_, 0xfc0007ff, 0x7c000631, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLD, 0xfc0007ff, 0x7c000036, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SLD_, 0xfc0007ff, 0x7c000037, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRD, 0xfc0007ff, 0x7c000436, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRD_, 0xfc0007ff, 0x7c000437, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRADI, 0xfc0007fd, 0x7c000674, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20}},
	{SRADI_, 0xfc0007fd, 0x7c000675, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_ImmUnsigned_30_30_16_20}},
	{SRAD, 0xfc0007ff, 0x7c000634, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{SRAD_, 0xfc0007ff, 0x7c000635, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{CDTBCD, 0xfc0007fe, 0x7c000234, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{CBCDTD, 0xfc0007fe, 0x7c000274, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{ADDG6S, 0xfc0003fe, 0x7c000094, 0x401,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1,
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTCRF, 0xfc1007fe, 0x7c000120, 0x801,
		[5]*argField{ap_ImmUnsigned_12_19, ap_Reg_6_10}},
	{MFCR, 0xfc1007fe, 0x7c000026, 0xff801,
		[5]*argField{ap_Reg_6_10}},
	{MTSLE, 0xfc0007fe, 0x7c000126, 0x3dff801,
		[5]*argField{ap_ImmUnsigned_10_10}},
	{MFVSRD, 0xfc0007fe, 0x7c000066, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_VecSReg_31_31_6_10}},
	{MFVSRWZ, 0xfc0007fe, 0x7c0000e6, 0xf800,
		[5]*argField{ap_Reg_11_15, ap_VecSReg_31_31_6_10}},
	{MTVSRD, 0xfc0007fe, 0x7c000166, 0xf800,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15}},
	{MTVSRWA, 0xfc0007fe, 0x7c0001a6, 0xf800,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15}},
	{MTVSRWZ, 0xfc0007fe, 0x7c0001e6, 0xf800,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15}},
	{MTOCRF, 0xfc1007fe, 0x7c100120, 0x801,
		[5]*argField{ap_ImmUnsigned_12_19, ap_Reg_6_10}},
	{MFOCRF, 0xfc1007fe, 0x7c100026, 0x801,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_12_19}},
	{MCRXR, 0xfc0007fe, 0x7c000400, 0x7ff801,
		[5]*argField{ap_CondRegField_6_8}},
	{MTDCRUX, 0xfc0007fe, 0x7c000346, 0xf801,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MFDCRUX, 0xfc0007fe, 0x7c000246, 0xf801,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{LFS, 0xfc000000, 0xc0000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFSU, 0xfc000000, 0xc4000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFSX, 0xfc0007fe, 0x7c00042e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFSUX, 0xfc0007fe, 0x7c00046e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFD, 0xfc000000, 0xc8000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFDU, 0xfc000000, 0xcc000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{LFDX, 0xfc0007fe, 0x7c0004ae, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDUX, 0xfc0007fe, 0x7c0004ee, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFIWAX, 0xfc0007fe, 0x7c0006ae, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFIWZX, 0xfc0007fe, 0x7c0006ee, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFS, 0xfc000000, 0xd0000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFSU, 0xfc000000, 0xd4000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFSX, 0xfc0007fe, 0x7c00052e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFSUX, 0xfc0007fe, 0x7c00056e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFD, 0xfc000000, 0xd8000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFDU, 0xfc000000, 0xdc000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_31, ap_Reg_11_15}},
	{STFDX, 0xfc0007fe, 0x7c0005ae, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDUX, 0xfc0007fe, 0x7c0005ee, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFIWX, 0xfc0007fe, 0x7c0007ae, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDP, 0xfc000003, 0xe4000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{LFDPX, 0xfc0007fe, 0x7c00062e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDP, 0xfc000003, 0xf4000000, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_Offset_16_29_shift2, ap_Reg_11_15}},
	{STFDPX, 0xfc0007fe, 0x7c00072e, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{FMR, 0xfc0007ff, 0xfc000090, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FMR_, 0xfc0007ff, 0xfc000091, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FABS, 0xfc0007ff, 0xfc000210, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FABS_, 0xfc0007ff, 0xfc000211, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNABS, 0xfc0007ff, 0xfc000110, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNABS_, 0xfc0007ff, 0xfc000111, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNEG, 0xfc0007ff, 0xfc000050, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FNEG_, 0xfc0007ff, 0xfc000051, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCPSGN, 0xfc0007ff, 0xfc000010, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FCPSGN_, 0xfc0007ff, 0xfc000011, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMRGEW, 0xfc0007fe, 0xfc00078c, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMRGOW, 0xfc0007fe, 0xfc00068c, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADD, 0xfc00003f, 0xfc00002a, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADD_, 0xfc00003f, 0xfc00002b, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADDS, 0xfc00003f, 0xec00002a, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FADDS_, 0xfc00003f, 0xec00002b, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUB, 0xfc00003f, 0xfc000028, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUB_, 0xfc00003f, 0xfc000029, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUBS, 0xfc00003f, 0xec000028, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSUBS_, 0xfc00003f, 0xec000029, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FMUL, 0xfc00003f, 0xfc000032, 0xf800,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMUL_, 0xfc00003f, 0xfc000033, 0xf800,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMULS, 0xfc00003f, 0xec000032, 0xf800,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FMULS_, 0xfc00003f, 0xec000033, 0xf800,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25}},
	{FDIV, 0xfc00003f, 0xfc000024, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIV_, 0xfc00003f, 0xfc000025, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIVS, 0xfc00003f, 0xec000024, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FDIVS_, 0xfc00003f, 0xec000025, 0x7c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSQRT, 0xfc00003f, 0xfc00002c, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRT_, 0xfc00003f, 0xfc00002d, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRTS, 0xfc00003f, 0xec00002c, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FSQRTS_, 0xfc00003f, 0xec00002d, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRE, 0xfc00003f, 0xfc000030, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRE_, 0xfc00003f, 0xfc000031, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRES, 0xfc00003f, 0xec000030, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRES_, 0xfc00003f, 0xec000031, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTE, 0xfc00003f, 0xfc000034, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTE_, 0xfc00003f, 0xfc000035, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTES, 0xfc00003f, 0xec000034, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSQRTES_, 0xfc00003f, 0xec000035, 0x1f07c0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FTDIV, 0xfc0007fe, 0xfc000100, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FTSQRT, 0xfc0007fe, 0xfc000140, 0x7f0001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_16_20}},
	{FMADD, 0xfc00003f, 0xfc00003a, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADD_, 0xfc00003f, 0xfc00003b, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADDS, 0xfc00003f, 0xec00003a, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMADDS_, 0xfc00003f, 0xec00003b, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUB, 0xfc00003f, 0xfc000038, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUB_, 0xfc00003f, 0xfc000039, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUBS, 0xfc00003f, 0xec000038, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FMSUBS_, 0xfc00003f, 0xec000039, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADD, 0xfc00003f, 0xfc00003e, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADD_, 0xfc00003f, 0xfc00003f, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADDS, 0xfc00003f, 0xec00003e, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMADDS_, 0xfc00003f, 0xec00003f, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUB, 0xfc00003f, 0xfc00003c, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUB_, 0xfc00003f, 0xfc00003d, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUBS, 0xfc00003f, 0xec00003c, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FNMSUBS_, 0xfc00003f, 0xec00003d, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FRSP, 0xfc0007ff, 0xfc000018, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRSP_, 0xfc0007ff, 0xfc000019, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTID, 0xfc0007ff, 0xfc00065c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTID_, 0xfc0007ff, 0xfc00065d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDZ, 0xfc0007ff, 0xfc00065e, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDZ_, 0xfc0007ff, 0xfc00065f, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDU, 0xfc0007ff, 0xfc00075c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDU_, 0xfc0007ff, 0xfc00075d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDUZ, 0xfc0007ff, 0xfc00075e, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIDUZ_, 0xfc0007ff, 0xfc00075f, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIW, 0xfc0007ff, 0xfc00001c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIW_, 0xfc0007ff, 0xfc00001d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWZ, 0xfc0007ff, 0xfc00001e, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWZ_, 0xfc0007ff, 0xfc00001f, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWU, 0xfc0007ff, 0xfc00011c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWU_, 0xfc0007ff, 0xfc00011d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWUZ, 0xfc0007ff, 0xfc00011e, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCTIWUZ_, 0xfc0007ff, 0xfc00011f, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFID, 0xfc0007ff, 0xfc00069c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFID_, 0xfc0007ff, 0xfc00069d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDU, 0xfc0007ff, 0xfc00079c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDU_, 0xfc0007ff, 0xfc00079d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDS, 0xfc0007ff, 0xec00069c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDS_, 0xfc0007ff, 0xec00069d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDUS, 0xfc0007ff, 0xec00079c, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCFIDUS_, 0xfc0007ff, 0xec00079d, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIN, 0xfc0007ff, 0xfc000310, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIN_, 0xfc0007ff, 0xfc000311, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIZ, 0xfc0007ff, 0xfc000350, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIZ_, 0xfc0007ff, 0xfc000351, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIP, 0xfc0007ff, 0xfc000390, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIP_, 0xfc0007ff, 0xfc000391, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIM, 0xfc0007ff, 0xfc0003d0, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FRIM_, 0xfc0007ff, 0xfc0003d1, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{FCMPU, 0xfc0007fe, 0xfc000000, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FCMPO, 0xfc0007fe, 0xfc000040, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{FSEL, 0xfc00003f, 0xfc00002e, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{FSEL_, 0xfc00003f, 0xfc00002f, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_21_25, ap_FPReg_16_20}},
	{MFFS, 0xfc0007ff, 0xfc00048e, 0x1ff800,
		[5]*argField{ap_FPReg_6_10}},
	{MFFS_, 0xfc0007ff, 0xfc00048f, 0x1ff800,
		[5]*argField{ap_FPReg_6_10}},
	{MCRFS, 0xfc0007fe, 0xfc000080, 0x63f801,
		[5]*argField{ap_CondRegField_6_8, ap_CondRegField_11_13}},
	{MTFSFI, 0xfc0007ff, 0xfc00010c, 0x7e0800,
		[5]*argField{ap_CondRegField_6_8, ap_ImmUnsigned_16_19, ap_ImmUnsigned_15_15}},
	{MTFSFI_, 0xfc0007ff, 0xfc00010d, 0x7e0800,
		[5]*argField{ap_CondRegField_6_8, ap_ImmUnsigned_16_19, ap_ImmUnsigned_15_15}},
	{MTFSF, 0xfc0007ff, 0xfc00058e, 0x0,
		[5]*argField{ap_ImmUnsigned_7_14, ap_FPReg_16_20, ap_ImmUnsigned_6_6, ap_ImmUnsigned_15_15}},
	{MTFSF_, 0xfc0007ff, 0xfc00058f, 0x0,
		[5]*argField{ap_ImmUnsigned_7_14, ap_FPReg_16_20, ap_ImmUnsigned_6_6, ap_ImmUnsigned_15_15}},
	{MTFSB0, 0xfc0007ff, 0xfc00008c, 0x1ff800,
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB0_, 0xfc0007ff, 0xfc00008d, 0x1ff800,
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB1, 0xfc0007ff, 0xfc00004c, 0x1ff800,
		[5]*argField{ap_CondRegBit_6_10}},
	{MTFSB1_, 0xfc0007ff, 0xfc00004d, 0x1ff800,
		[5]*argField{ap_CondRegBit_6_10}},
	{LVEBX, 0xfc0007fe, 0x7c00000e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEHX, 0xfc0007fe, 0x7c00004e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEWX, 0xfc0007fe, 0x7c00008e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVX, 0xfc0007fe, 0x7c0000ce, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVXL, 0xfc0007fe, 0x7c0002ce, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEBX, 0xfc0007fe, 0x7c00010e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEHX, 0xfc0007fe, 0x7c00014e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEWX, 0xfc0007fe, 0x7c00018e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVX, 0xfc0007fe, 0x7c0001ce, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVXL, 0xfc0007fe, 0x7c0003ce, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVSL, 0xfc0007fe, 0x7c00000c, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVSR, 0xfc0007fe, 0x7c00004c, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{VPKPX, 0xfc0007ff, 0x1000030e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSDSS, 0xfc0007ff, 0x100005ce, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSDUS, 0xfc0007ff, 0x1000054e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSHSS, 0xfc0007ff, 0x1000018e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSHUS, 0xfc0007ff, 0x1000010e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSWSS, 0xfc0007ff, 0x100001ce, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKSWUS, 0xfc0007ff, 0x1000014e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUDUM, 0xfc0007ff, 0x1000044e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUDUS, 0xfc0007ff, 0x100004ce, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUHUM, 0xfc0007ff, 0x1000000e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUHUS, 0xfc0007ff, 0x1000008e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUWUM, 0xfc0007ff, 0x1000004e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPKUWUS, 0xfc0007ff, 0x100000ce, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VUPKHPX, 0xfc0007ff, 0x1000034e, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLPX, 0xfc0007ff, 0x100003ce, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSB, 0xfc0007ff, 0x1000020e, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSH, 0xfc0007ff, 0x1000024e, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKHSW, 0xfc0007ff, 0x1000064e, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSB, 0xfc0007ff, 0x1000028e, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSH, 0xfc0007ff, 0x100002ce, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VUPKLSW, 0xfc0007ff, 0x100006ce, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VMRGHB, 0xfc0007ff, 0x1000000c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGHH, 0xfc0007ff, 0x1000004c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLB, 0xfc0007ff, 0x1000010c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLH, 0xfc0007ff, 0x1000014c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGHW, 0xfc0007ff, 0x1000008c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGLW, 0xfc0007ff, 0x1000018c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGEW, 0xfc0007ff, 0x1000078c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMRGOW, 0xfc0007ff, 0x1000068c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSPLTB, 0xfc0007ff, 0x1000020c, 0x100000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_12_15}},
	{VSPLTH, 0xfc0007ff, 0x1000024c, 0x180000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_13_15}},
	{VSPLTW, 0xfc0007ff, 0x1000028c, 0x1c0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_14_15}},
	{VSPLTISB, 0xfc0007ff, 0x1000030c, 0xf800,
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VSPLTISH, 0xfc0007ff, 0x1000034c, 0xf800,
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VSPLTISW, 0xfc0007ff, 0x1000038c, 0xf800,
		[5]*argField{ap_VecReg_6_10, ap_ImmSigned_11_15}},
	{VPERM, 0xfc00003f, 0x1000002b, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSEL, 0xfc00003f, 0x1000002a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSL, 0xfc0007ff, 0x100001c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLDOI, 0xfc00003f, 0x1000002c, 0x400,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_25}},
	{VSLO, 0xfc0007ff, 0x1000040c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSR, 0xfc0007ff, 0x100002c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRO, 0xfc0007ff, 0x1000044c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDCUW, 0xfc0007ff, 0x10000180, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSBS, 0xfc0007ff, 0x10000300, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSHS, 0xfc0007ff, 0x10000340, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDSWS, 0xfc0007ff, 0x10000380, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUBM, 0xfc0007ff, 0x10000000, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUDM, 0xfc0007ff, 0x100000c0, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUHM, 0xfc0007ff, 0x10000040, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUWM, 0xfc0007ff, 0x10000080, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUBS, 0xfc0007ff, 0x10000200, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUHS, 0xfc0007ff, 0x10000240, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUWS, 0xfc0007ff, 0x10000280, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDUQM, 0xfc0007ff, 0x10000100, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDEUQM, 0xfc00003f, 0x1000003c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VADDCUQ, 0xfc0007ff, 0x10000140, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDECUQ, 0xfc00003f, 0x1000003d, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUBCUW, 0xfc0007ff, 0x10000580, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSBS, 0xfc0007ff, 0x10000700, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSHS, 0xfc0007ff, 0x10000740, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBSWS, 0xfc0007ff, 0x10000780, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUBM, 0xfc0007ff, 0x10000400, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUDM, 0xfc0007ff, 0x100004c0, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUHM, 0xfc0007ff, 0x10000440, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUWM, 0xfc0007ff, 0x10000480, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUBS, 0xfc0007ff, 0x10000600, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUHS, 0xfc0007ff, 0x10000640, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUWS, 0xfc0007ff, 0x10000680, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBUQM, 0xfc0007ff, 0x10000500, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBEUQM, 0xfc00003f, 0x1000003e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUBCUQ, 0xfc0007ff, 0x10000540, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBECUQ, 0xfc00003f, 0x1000003f, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMULESB, 0xfc0007ff, 0x10000308, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUB, 0xfc0007ff, 0x10000208, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSB, 0xfc0007ff, 0x10000108, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUB, 0xfc0007ff, 0x10000008, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULESH, 0xfc0007ff, 0x10000348, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUH, 0xfc0007ff, 0x10000248, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSH, 0xfc0007ff, 0x10000148, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUH, 0xfc0007ff, 0x10000048, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULESW, 0xfc0007ff, 0x10000388, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULEUW, 0xfc0007ff, 0x10000288, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOSW, 0xfc0007ff, 0x10000188, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULOUW, 0xfc0007ff, 0x10000088, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMULUWM, 0xfc0007ff, 0x10000089, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMHADDSHS, 0xfc00003f, 0x10000020, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMHRADDSHS, 0xfc00003f, 0x10000021, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMLADDUHM, 0xfc00003f, 0x10000022, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUBM, 0xfc00003f, 0x10000024, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMMBM, 0xfc00003f, 0x10000025, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMSHM, 0xfc00003f, 0x10000028, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMSHS, 0xfc00003f, 0x10000029, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUHM, 0xfc00003f, 0x10000026, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VMSUMUHS, 0xfc00003f, 0x10000027, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VSUMSWS, 0xfc0007ff, 0x10000788, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM2SWS, 0xfc0007ff, 0x10000688, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4SBS, 0xfc0007ff, 0x10000708, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4SHS, 0xfc0007ff, 0x10000648, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUM4UBS, 0xfc0007ff, 0x10000608, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSB, 0xfc0007ff, 0x10000502, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSH, 0xfc0007ff, 0x10000542, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGSW, 0xfc0007ff, 0x10000582, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUB, 0xfc0007ff, 0x10000402, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUW, 0xfc0007ff, 0x10000482, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAVGUH, 0xfc0007ff, 0x10000442, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSB, 0xfc0007ff, 0x10000102, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSD, 0xfc0007ff, 0x100001c2, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUB, 0xfc0007ff, 0x10000002, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUD, 0xfc0007ff, 0x100000c2, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSH, 0xfc0007ff, 0x10000142, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXSW, 0xfc0007ff, 0x10000182, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUH, 0xfc0007ff, 0x10000042, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMAXUW, 0xfc0007ff, 0x10000082, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSB, 0xfc0007ff, 0x10000302, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSD, 0xfc0007ff, 0x100003c2, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUB, 0xfc0007ff, 0x10000202, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUD, 0xfc0007ff, 0x100002c2, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSH, 0xfc0007ff, 0x10000342, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINSW, 0xfc0007ff, 0x10000382, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUH, 0xfc0007ff, 0x10000242, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINUW, 0xfc0007ff, 0x10000282, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUB, 0xfc0007ff, 0x10000006, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUB_, 0xfc0007ff, 0x10000406, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUH, 0xfc0007ff, 0x10000046, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUH_, 0xfc0007ff, 0x10000446, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUW, 0xfc0007ff, 0x10000086, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUW_, 0xfc0007ff, 0x10000486, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUD, 0xfc0007ff, 0x100000c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQUD_, 0xfc0007ff, 0x100004c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSB, 0xfc0007ff, 0x10000306, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSB_, 0xfc0007ff, 0x10000706, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSD, 0xfc0007ff, 0x100003c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSD_, 0xfc0007ff, 0x100007c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSH, 0xfc0007ff, 0x10000346, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSH_, 0xfc0007ff, 0x10000746, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSW, 0xfc0007ff, 0x10000386, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTSW_, 0xfc0007ff, 0x10000786, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUB, 0xfc0007ff, 0x10000206, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUB_, 0xfc0007ff, 0x10000606, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUD, 0xfc0007ff, 0x100002c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUD_, 0xfc0007ff, 0x100006c7, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUH, 0xfc0007ff, 0x10000246, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUH_, 0xfc0007ff, 0x10000646, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUW, 0xfc0007ff, 0x10000286, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTUW_, 0xfc0007ff, 0x10000686, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VAND, 0xfc0007ff, 0x10000404, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VANDC, 0xfc0007ff, 0x10000444, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VEQV, 0xfc0007ff, 0x10000684, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNAND, 0xfc0007ff, 0x10000584, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VORC, 0xfc0007ff, 0x10000544, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNOR, 0xfc0007ff, 0x10000504, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VOR, 0xfc0007ff, 0x10000484, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VXOR, 0xfc0007ff, 0x100004c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLB, 0xfc0007ff, 0x10000004, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLH, 0xfc0007ff, 0x10000044, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLW, 0xfc0007ff, 0x10000084, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VRLD, 0xfc0007ff, 0x100000c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLB, 0xfc0007ff, 0x10000104, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLH, 0xfc0007ff, 0x10000144, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLW, 0xfc0007ff, 0x10000184, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSLD, 0xfc0007ff, 0x100005c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRB, 0xfc0007ff, 0x10000204, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRH, 0xfc0007ff, 0x10000244, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRW, 0xfc0007ff, 0x10000284, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRD, 0xfc0007ff, 0x100006c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAB, 0xfc0007ff, 0x10000304, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAH, 0xfc0007ff, 0x10000344, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAW, 0xfc0007ff, 0x10000384, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSRAD, 0xfc0007ff, 0x100003c4, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VADDFP, 0xfc0007ff, 0x1000000a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSUBFP, 0xfc0007ff, 0x1000004a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMADDFP, 0xfc00003f, 0x1000002e, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_21_25, ap_VecReg_16_20}},
	{VNMSUBFP, 0xfc00003f, 0x1000002f, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_21_25, ap_VecReg_16_20}},
	{VMAXFP, 0xfc0007ff, 0x1000040a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VMINFP, 0xfc0007ff, 0x1000044a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCTSXS, 0xfc0007ff, 0x100003ca, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCTUXS, 0xfc0007ff, 0x1000038a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCFSX, 0xfc0007ff, 0x1000034a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VCFUX, 0xfc0007ff, 0x1000030a, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20, ap_ImmUnsigned_11_15}},
	{VRFIM, 0xfc0007ff, 0x100002ca, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIN, 0xfc0007ff, 0x1000020a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIP, 0xfc0007ff, 0x1000028a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRFIZ, 0xfc0007ff, 0x1000024a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCMPBFP, 0xfc0007ff, 0x100003c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPBFP_, 0xfc0007ff, 0x100007c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQFP, 0xfc0007ff, 0x100000c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPEQFP_, 0xfc0007ff, 0x100004c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGEFP, 0xfc0007ff, 0x100001c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGEFP_, 0xfc0007ff, 0x100005c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTFP, 0xfc0007ff, 0x100002c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCMPGTFP_, 0xfc0007ff, 0x100006c6, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VEXPTEFP, 0xfc0007ff, 0x1000018a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VLOGEFP, 0xfc0007ff, 0x100001ca, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VREFP, 0xfc0007ff, 0x1000010a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VRSQRTEFP, 0xfc0007ff, 0x1000014a, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCIPHER, 0xfc0007ff, 0x10000508, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VCIPHERLAST, 0xfc0007ff, 0x10000509, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNCIPHER, 0xfc0007ff, 0x10000548, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VNCIPHERLAST, 0xfc0007ff, 0x10000549, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VSBOX, 0xfc0007ff, 0x100005c8, 0xf800,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15}},
	{VSHASIGMAD, 0xfc0007ff, 0x100006c2, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_ImmUnsigned_16_16, ap_ImmUnsigned_17_20}},
	{VSHASIGMAW, 0xfc0007ff, 0x10000682, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_ImmUnsigned_16_16, ap_ImmUnsigned_17_20}},
	{VPMSUMB, 0xfc0007ff, 0x10000408, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMD, 0xfc0007ff, 0x100004c8, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMH, 0xfc0007ff, 0x10000448, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPMSUMW, 0xfc0007ff, 0x10000488, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{VPERMXOR, 0xfc00003f, 0x1000002d, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_VecReg_21_25}},
	{VGBBD, 0xfc0007ff, 0x1000050c, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZB, 0xfc0007ff, 0x10000702, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZH, 0xfc0007ff, 0x10000742, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZW, 0xfc0007ff, 0x10000782, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VCLZD, 0xfc0007ff, 0x100007c2, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTB, 0xfc0007ff, 0x10000703, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTD, 0xfc0007ff, 0x100007c3, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTH, 0xfc0007ff, 0x10000743, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VPOPCNTW, 0xfc0007ff, 0x10000783, 0x1f0000,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_16_20}},
	{VBPERMQ, 0xfc0007ff, 0x1000054c, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20}},
	{BCDADD_, 0xfc0005ff, 0x10000401, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_22}},
	{BCDSUB_, 0xfc0005ff, 0x10000441, 0x0,
		[5]*argField{ap_VecReg_6_10, ap_VecReg_11_15, ap_VecReg_16_20, ap_ImmUnsigned_22_22}},
	{MTVSCR, 0xfc0007ff, 0x10000644, 0x3ff0000,
		[5]*argField{ap_VecReg_16_20}},
	{MFVSCR, 0xfc0007ff, 0x10000604, 0x1ff800,
		[5]*argField{ap_VecReg_6_10}},
	{DADD, 0xfc0007ff, 0xec000004, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DADD_, 0xfc0007ff, 0xec000005, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSUB, 0xfc0007ff, 0xec000404, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSUB_, 0xfc0007ff, 0xec000405, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DMUL, 0xfc0007ff, 0xec000044, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DMUL_, 0xfc0007ff, 0xec000045, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DDIV, 0xfc0007ff, 0xec000444, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DDIV_, 0xfc0007ff, 0xec000445, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DCMPU, 0xfc0007fe, 0xec000504, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DCMPO, 0xfc0007fe, 0xec000104, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DTSTDC, 0xfc0003fe, 0xec000184, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DTSTDG, 0xfc0003fe, 0xec0001c4, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DTSTEX, 0xfc0007fe, 0xec000144, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DTSTSF, 0xfc0007fe, 0xec000544, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DQUAI, 0xfc0001ff, 0xec000086, 0x0,
		[5]*argField{ap_ImmSigned_11_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUAI_, 0xfc0001ff, 0xec000087, 0x0,
		[5]*argField{ap_ImmSigned_11_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUA, 0xfc0001ff, 0xec000006, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DQUA_, 0xfc0001ff, 0xec000007, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRRND, 0xfc0001ff, 0xec000046, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRRND_, 0xfc0001ff, 0xec000047, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTX, 0xfc0001ff, 0xec0000c6, 0x1e0000,
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTX_, 0xfc0001ff, 0xec0000c7, 0x1e0000,
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTN, 0xfc0001ff, 0xec0001c6, 0x1e0000,
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DRINTN_, 0xfc0001ff, 0xec0001c7, 0x1e0000,
		[5]*argField{ap_ImmUnsigned_15_15, ap_FPReg_6_10, ap_FPReg_16_20, ap_ImmUnsigned_21_22}},
	{DCTDP, 0xfc0007ff, 0xec000204, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTDP_, 0xfc0007ff, 0xec000205, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTQPQ, 0xfc0007ff, 0xfc000204, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTQPQ_, 0xfc0007ff, 0xfc000205, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRSP, 0xfc0007ff, 0xec000604, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRSP_, 0xfc0007ff, 0xec000605, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRDPQ, 0xfc0007ff, 0xfc000604, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DRDPQ_, 0xfc0007ff, 0xfc000605, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIX, 0xfc0007ff, 0xec000644, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIX_, 0xfc0007ff, 0xec000645, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIXQ, 0xfc0007ff, 0xfc000644, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCFFIXQ_, 0xfc0007ff, 0xfc000645, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTFIX, 0xfc0007ff, 0xec000244, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DCTFIX_, 0xfc0007ff, 0xec000245, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DDEDPD, 0xfc0007ff, 0xec000284, 0x70000,
		[5]*argField{ap_ImmUnsigned_11_12, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DDEDPD_, 0xfc0007ff, 0xec000285, 0x70000,
		[5]*argField{ap_ImmUnsigned_11_12, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DENBCD, 0xfc0007ff, 0xec000684, 0xf0000,
		[5]*argField{ap_ImmUnsigned_11_11, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DENBCD_, 0xfc0007ff, 0xec000685, 0xf0000,
		[5]*argField{ap_ImmUnsigned_11_11, ap_FPReg_6_10, ap_FPReg_16_20}},
	{DXEX, 0xfc0007ff, 0xec0002c4, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DXEX_, 0xfc0007ff, 0xec0002c5, 0x1f0000,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_16_20}},
	{DIEX, 0xfc0007ff, 0xec0006c4, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DIEX_, 0xfc0007ff, 0xec0006c5, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_FPReg_16_20}},
	{DSCLI, 0xfc0003ff, 0xec000084, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCLI_, 0xfc0003ff, 0xec000085, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCRI, 0xfc0003ff, 0xec0000c4, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{DSCRI_, 0xfc0003ff, 0xec0000c5, 0x0,
		[5]*argField{ap_FPReg_6_10, ap_FPReg_11_15, ap_ImmUnsigned_16_21}},
	{LXSDX, 0xfc0007fe, 0x7c000498, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSIWAX, 0xfc0007fe, 0x7c000098, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSIWZX, 0xfc0007fe, 0x7c000018, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXSSPX, 0xfc0007fe, 0x7c000418, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVD2X, 0xfc0007fe, 0x7c000698, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVDSX, 0xfc0007fe, 0x7c000298, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LXVW4X, 0xfc0007fe, 0x7c000618, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSDX, 0xfc0007fe, 0x7c000598, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSIWX, 0xfc0007fe, 0x7c000118, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXSSPX, 0xfc0007fe, 0x7c000518, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXVD2X, 0xfc0007fe, 0x7c000798, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STXVW4X, 0xfc0007fe, 0x7c000718, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{XSABSDP, 0xfc0007fc, 0xf0000564, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSADDDP, 0xfc0007f8, 0xf0000100, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSADDSP, 0xfc0007f8, 0xf0000000, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSCMPODP, 0xfc0007f8, 0xf0000158, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSCMPUDP, 0xfc0007f8, 0xf0000118, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSCPSGNDP, 0xfc0007f8, 0xf0000580, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSCVDPSP, 0xfc0007fc, 0xf0000424, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVDPSPN, 0xfc0007fc, 0xf000042c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVDPSXDS, 0xfc0007fc, 0xf0000560, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVDPSXWS, 0xfc0007fc, 0xf0000160, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVDPUXDS, 0xfc0007fc, 0xf0000520, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVDPUXWS, 0xfc0007fc, 0xf0000120, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVSPDP, 0xfc0007fc, 0xf0000524, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVSPDPN, 0xfc0007fc, 0xf000052c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVSXDDP, 0xfc0007fc, 0xf00005e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVSXDSP, 0xfc0007fc, 0xf00004e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVUXDDP, 0xfc0007fc, 0xf00005a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSCVUXDSP, 0xfc0007fc, 0xf00004a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSDIVDP, 0xfc0007f8, 0xf00001c0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSDIVSP, 0xfc0007f8, 0xf00000c0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMADDADP, 0xfc0007f8, 0xf0000108, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMADDASP, 0xfc0007f8, 0xf0000008, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMAXDP, 0xfc0007f8, 0xf0000500, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMINDP, 0xfc0007f8, 0xf0000540, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMSUBADP, 0xfc0007f8, 0xf0000188, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMSUBASP, 0xfc0007f8, 0xf0000088, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMULDP, 0xfc0007f8, 0xf0000180, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSMULSP, 0xfc0007f8, 0xf0000080, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSNABSDP, 0xfc0007fc, 0xf00005a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSNEGDP, 0xfc0007fc, 0xf00005e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSNMADDADP, 0xfc0007f8, 0xf0000508, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSNMADDASP, 0xfc0007f8, 0xf0000408, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSNMSUBADP, 0xfc0007f8, 0xf0000588, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSNMSUBASP, 0xfc0007f8, 0xf0000488, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSRDPI, 0xfc0007fc, 0xf0000124, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRDPIC, 0xfc0007fc, 0xf00001ac, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRDPIM, 0xfc0007fc, 0xf00001e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRDPIP, 0xfc0007fc, 0xf00001a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRDPIZ, 0xfc0007fc, 0xf0000164, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSREDP, 0xfc0007fc, 0xf0000168, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRESP, 0xfc0007fc, 0xf0000068, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRSP, 0xfc0007fc, 0xf0000464, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRSQRTEDP, 0xfc0007fc, 0xf0000128, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSRSQRTESP, 0xfc0007fc, 0xf0000028, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSSQRTDP, 0xfc0007fc, 0xf000012c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSSQRTSP, 0xfc0007fc, 0xf000002c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XSSUBDP, 0xfc0007f8, 0xf0000140, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSSUBSP, 0xfc0007f8, 0xf0000040, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSTDIVDP, 0xfc0007f8, 0xf00001e8, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XSTSQRTDP, 0xfc0007fc, 0xf00001a8, 0x7f0001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_30_30_16_20}},
	{XVABSDP, 0xfc0007fc, 0xf0000764, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVABSSP, 0xfc0007fc, 0xf0000664, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVADDDP, 0xfc0007f8, 0xf0000300, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVADDSP, 0xfc0007f8, 0xf0000200, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPEQDP, 0xfc0007f8, 0xf0000318, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPEQDP_, 0xfc0007f8, 0xf0000718, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPEQSP, 0xfc0007f8, 0xf0000218, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPEQSP_, 0xfc0007f8, 0xf0000618, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGEDP, 0xfc0007f8, 0xf0000398, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGEDP_, 0xfc0007f8, 0xf0000798, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGESP, 0xfc0007f8, 0xf0000298, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGESP_, 0xfc0007f8, 0xf0000698, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGTDP, 0xfc0007f8, 0xf0000358, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGTDP_, 0xfc0007f8, 0xf0000758, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGTSP, 0xfc0007f8, 0xf0000258, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCMPGTSP_, 0xfc0007f8, 0xf0000658, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCPSGNDP, 0xfc0007f8, 0xf0000780, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCPSGNSP, 0xfc0007f8, 0xf0000680, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVCVDPSP, 0xfc0007fc, 0xf0000624, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVDPSXDS, 0xfc0007fc, 0xf0000760, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVDPSXWS, 0xfc0007fc, 0xf0000360, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVDPUXDS, 0xfc0007fc, 0xf0000720, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVDPUXWS, 0xfc0007fc, 0xf0000320, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSPDP, 0xfc0007fc, 0xf0000724, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSPSXDS, 0xfc0007fc, 0xf0000660, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSPSXWS, 0xfc0007fc, 0xf0000260, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSPUXDS, 0xfc0007fc, 0xf0000620, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSPUXWS, 0xfc0007fc, 0xf0000220, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSXDDP, 0xfc0007fc, 0xf00007e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSXDSP, 0xfc0007fc, 0xf00006e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSXWDP, 0xfc0007fc, 0xf00003e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVSXWSP, 0xfc0007fc, 0xf00002e0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVUXDDP, 0xfc0007fc, 0xf00007a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVUXDSP, 0xfc0007fc, 0xf00006a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVUXWDP, 0xfc0007fc, 0xf00003a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVCVUXWSP, 0xfc0007fc, 0xf00002a0, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVDIVDP, 0xfc0007f8, 0xf00003c0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVDIVSP, 0xfc0007f8, 0xf00002c0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMADDADP, 0xfc0007f8, 0xf0000308, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMADDASP, 0xfc0007f8, 0xf0000208, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMAXDP, 0xfc0007f8, 0xf0000700, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMAXSP, 0xfc0007f8, 0xf0000600, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMINDP, 0xfc0007f8, 0xf0000740, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMINSP, 0xfc0007f8, 0xf0000640, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMSUBADP, 0xfc0007f8, 0xf0000388, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMSUBASP, 0xfc0007f8, 0xf0000288, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMULDP, 0xfc0007f8, 0xf0000380, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVMULSP, 0xfc0007f8, 0xf0000280, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVNABSDP, 0xfc0007fc, 0xf00007a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVNABSSP, 0xfc0007fc, 0xf00006a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVNEGDP, 0xfc0007fc, 0xf00007e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVNEGSP, 0xfc0007fc, 0xf00006e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVNMADDADP, 0xfc0007f8, 0xf0000708, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVNMADDASP, 0xfc0007f8, 0xf0000608, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVNMSUBADP, 0xfc0007f8, 0xf0000788, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVNMSUBASP, 0xfc0007f8, 0xf0000688, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVRDPI, 0xfc0007fc, 0xf0000324, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRDPIC, 0xfc0007fc, 0xf00003ac, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRDPIM, 0xfc0007fc, 0xf00003e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRDPIP, 0xfc0007fc, 0xf00003a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRDPIZ, 0xfc0007fc, 0xf0000364, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVREDP, 0xfc0007fc, 0xf0000368, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRESP, 0xfc0007fc, 0xf0000268, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSPI, 0xfc0007fc, 0xf0000224, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSPIC, 0xfc0007fc, 0xf00002ac, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSPIM, 0xfc0007fc, 0xf00002e4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSPIP, 0xfc0007fc, 0xf00002a4, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSPIZ, 0xfc0007fc, 0xf0000264, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSQRTEDP, 0xfc0007fc, 0xf0000328, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVRSQRTESP, 0xfc0007fc, 0xf0000228, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVSQRTDP, 0xfc0007fc, 0xf000032c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVSQRTSP, 0xfc0007fc, 0xf000022c, 0x1f0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20}},
	{XVSUBDP, 0xfc0007f8, 0xf0000340, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVSUBSP, 0xfc0007f8, 0xf0000240, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVTDIVDP, 0xfc0007f8, 0xf00003e8, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVTDIVSP, 0xfc0007f8, 0xf00002e8, 0x600001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XVTSQRTDP, 0xfc0007fc, 0xf00003a8, 0x7f0001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_30_30_16_20}},
	{XVTSQRTSP, 0xfc0007fc, 0xf00002a8, 0x7f0001,
		[5]*argField{ap_CondRegField_6_8, ap_VecSReg_30_30_16_20}},
	{XXLAND, 0xfc0007f8, 0xf0000410, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLANDC, 0xfc0007f8, 0xf0000450, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLEQV, 0xfc0007f8, 0xf00005d0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLNAND, 0xfc0007f8, 0xf0000590, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLORC, 0xfc0007f8, 0xf0000550, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLNOR, 0xfc0007f8, 0xf0000510, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLOR, 0xfc0007f8, 0xf0000490, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXLXOR, 0xfc0007f8, 0xf00004d0, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXMRGHW, 0xfc0007f8, 0xf0000090, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXMRGLW, 0xfc0007f8, 0xf0000190, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20}},
	{XXPERMDI, 0xfc0004f8, 0xf0000050, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20, ap_ImmUnsigned_22_23}},
	{XXSEL, 0xfc000030, 0xf0000030, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20, ap_VecSReg_28_28_21_25}},
	{XXSLDWI, 0xfc0004f8, 0xf0000010, 0x0,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_29_29_11_15, ap_VecSReg_30_30_16_20, ap_ImmUnsigned_22_23}},
	{XXSPLTW, 0xfc0007fc, 0xf0000290, 0x1c0000,
		[5]*argField{ap_VecSReg_31_31_6_10, ap_VecSReg_30_30_16_20, ap_ImmUnsigned_14_15}},
	{BRINC, 0xfc0007ff, 0x1000020f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVABS, 0xfc0007ff, 0x10000208, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDIW, 0xfc0007ff, 0x10000202, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20, ap_ImmUnsigned_11_15}},
	{EVADDSMIAAW, 0xfc0007ff, 0x100004c9, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDSSIAAW, 0xfc0007ff, 0x100004c1, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDUMIAAW, 0xfc0007ff, 0x100004c8, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDUSIAAW, 0xfc0007ff, 0x100004c0, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVADDW, 0xfc0007ff, 0x10000200, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVAND, 0xfc0007ff, 0x10000211, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPEQ, 0xfc0007ff, 0x10000234, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVANDC, 0xfc0007ff, 0x10000212, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPGTS, 0xfc0007ff, 0x10000231, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPGTU, 0xfc0007ff, 0x10000230, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPLTU, 0xfc0007ff, 0x10000232, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCMPLTS, 0xfc0007ff, 0x10000233, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVCNTLSW, 0xfc0007ff, 0x1000020e, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVCNTLZW, 0xfc0007ff, 0x1000020d, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVDIVWS, 0xfc0007ff, 0x100004c6, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVDIVWU, 0xfc0007ff, 0x100004c7, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVEQV, 0xfc0007ff, 0x10000219, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVEXTSB, 0xfc0007ff, 0x1000020a, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVEXTSH, 0xfc0007ff, 0x1000020b, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVLDD, 0xfc0007ff, 0x10000301, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDH, 0xfc0007ff, 0x10000305, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDDX, 0xfc0007ff, 0x10000300, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDHX, 0xfc0007ff, 0x10000304, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDW, 0xfc0007ff, 0x10000303, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHESPLAT, 0xfc0007ff, 0x10000309, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLDWX, 0xfc0007ff, 0x10000302, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHESPLATX, 0xfc0007ff, 0x10000308, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHOSSPLAT, 0xfc0007ff, 0x1000030f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHOUSPLAT, 0xfc0007ff, 0x1000030d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLHHOSSPLATX, 0xfc0007ff, 0x1000030e, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLHHOUSPLATX, 0xfc0007ff, 0x1000030c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHE, 0xfc0007ff, 0x10000311, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHOS, 0xfc0007ff, 0x10000317, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHEX, 0xfc0007ff, 0x10000310, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHOSX, 0xfc0007ff, 0x10000316, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHOU, 0xfc0007ff, 0x10000315, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHSPLAT, 0xfc0007ff, 0x1000031d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVLWHOUX, 0xfc0007ff, 0x10000314, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWHSPLATX, 0xfc0007ff, 0x1000031c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWWSPLAT, 0xfc0007ff, 0x10000319, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVMERGEHI, 0xfc0007ff, 0x1000022c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLWWSPLATX, 0xfc0007ff, 0x10000318, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGELO, 0xfc0007ff, 0x1000022d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGEHILO, 0xfc0007ff, 0x1000022e, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMFAA, 0xfc0007ff, 0x1000052b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMERGELOHI, 0xfc0007ff, 0x1000022f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMFAN, 0xfc0007ff, 0x100005ab, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMIAA, 0xfc0007ff, 0x10000529, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGUMIAA, 0xfc0007ff, 0x10000528, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGSMIAN, 0xfc0007ff, 0x100005a9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEGUMIAN, 0xfc0007ff, 0x100005a8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMF, 0xfc0007ff, 0x1000040b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFAAW, 0xfc0007ff, 0x1000050b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFA, 0xfc0007ff, 0x1000042b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMFANW, 0xfc0007ff, 0x1000058b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMI, 0xfc0007ff, 0x10000409, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIAAW, 0xfc0007ff, 0x10000509, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIA, 0xfc0007ff, 0x10000429, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESMIANW, 0xfc0007ff, 0x10000589, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSF, 0xfc0007ff, 0x10000403, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFA, 0xfc0007ff, 0x10000423, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFAAW, 0xfc0007ff, 0x10000503, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSFANW, 0xfc0007ff, 0x10000583, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSIAAW, 0xfc0007ff, 0x10000501, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHESSIANW, 0xfc0007ff, 0x10000581, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMI, 0xfc0007ff, 0x10000408, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIAAW, 0xfc0007ff, 0x10000508, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIA, 0xfc0007ff, 0x10000428, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUMIANW, 0xfc0007ff, 0x10000588, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUSIAAW, 0xfc0007ff, 0x10000500, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHEUSIANW, 0xfc0007ff, 0x10000580, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMFAA, 0xfc0007ff, 0x1000052f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMIAA, 0xfc0007ff, 0x1000052d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMFAN, 0xfc0007ff, 0x100005af, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGSMIAN, 0xfc0007ff, 0x100005ad, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGUMIAA, 0xfc0007ff, 0x1000052c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMF, 0xfc0007ff, 0x1000040f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOGUMIAN, 0xfc0007ff, 0x100005ac, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFA, 0xfc0007ff, 0x1000042f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFAAW, 0xfc0007ff, 0x1000050f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMI, 0xfc0007ff, 0x1000040d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMFANW, 0xfc0007ff, 0x1000058f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIA, 0xfc0007ff, 0x1000042d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIAAW, 0xfc0007ff, 0x1000050d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSMIANW, 0xfc0007ff, 0x1000058d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSF, 0xfc0007ff, 0x10000407, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFA, 0xfc0007ff, 0x10000427, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFAAW, 0xfc0007ff, 0x10000507, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSFANW, 0xfc0007ff, 0x10000587, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSIAAW, 0xfc0007ff, 0x10000505, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMI, 0xfc0007ff, 0x1000040c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOSSIANW, 0xfc0007ff, 0x10000585, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIA, 0xfc0007ff, 0x1000042c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIAAW, 0xfc0007ff, 0x1000050c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUSIAAW, 0xfc0007ff, 0x10000504, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUMIANW, 0xfc0007ff, 0x1000058c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMHOUSIANW, 0xfc0007ff, 0x10000584, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMRA, 0xfc0007ff, 0x100004c4, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVMWHSMF, 0xfc0007ff, 0x1000044f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMI, 0xfc0007ff, 0x1000044d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMFA, 0xfc0007ff, 0x1000046f, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSMIA, 0xfc0007ff, 0x1000046d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSSF, 0xfc0007ff, 0x10000447, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHUMI, 0xfc0007ff, 0x1000044c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHSSFA, 0xfc0007ff, 0x10000467, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWHUMIA, 0xfc0007ff, 0x1000046c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSMIAAW, 0xfc0007ff, 0x10000549, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSSIAAW, 0xfc0007ff, 0x10000541, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSMIANW, 0xfc0007ff, 0x100005c9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLSSIANW, 0xfc0007ff, 0x100005c1, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMI, 0xfc0007ff, 0x10000448, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIAAW, 0xfc0007ff, 0x10000548, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIA, 0xfc0007ff, 0x10000468, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUMIANW, 0xfc0007ff, 0x100005c8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUSIAAW, 0xfc0007ff, 0x10000540, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMF, 0xfc0007ff, 0x1000045b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWLUSIANW, 0xfc0007ff, 0x100005c0, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFA, 0xfc0007ff, 0x1000047b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFAA, 0xfc0007ff, 0x1000055b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMI, 0xfc0007ff, 0x10000459, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIAA, 0xfc0007ff, 0x10000559, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMFAN, 0xfc0007ff, 0x100005db, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIA, 0xfc0007ff, 0x10000479, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSMIAN, 0xfc0007ff, 0x100005d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSF, 0xfc0007ff, 0x10000453, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFA, 0xfc0007ff, 0x10000473, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFAA, 0xfc0007ff, 0x10000553, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMI, 0xfc0007ff, 0x10000458, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWSSFAN, 0xfc0007ff, 0x100005d3, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIA, 0xfc0007ff, 0x10000478, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIAA, 0xfc0007ff, 0x10000558, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVNAND, 0xfc0007ff, 0x1000021e, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVMWUMIAN, 0xfc0007ff, 0x100005d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVNEG, 0xfc0007ff, 0x10000209, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVNOR, 0xfc0007ff, 0x10000218, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVORC, 0xfc0007ff, 0x1000021b, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVOR, 0xfc0007ff, 0x10000217, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVRLW, 0xfc0007ff, 0x10000228, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVRLWI, 0xfc0007ff, 0x1000022a, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSEL, 0xfc0007f8, 0x10000278, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_CondRegField_29_31}},
	{EVRNDW, 0xfc0007ff, 0x1000020c, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSLW, 0xfc0007ff, 0x10000224, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSPLATFI, 0xfc0007ff, 0x1000022b, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_11_15}},
	{EVSRWIS, 0xfc0007ff, 0x10000223, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSLWI, 0xfc0007ff, 0x10000226, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSPLATI, 0xfc0007ff, 0x10000229, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_ImmSigned_11_15}},
	{EVSRWIU, 0xfc0007ff, 0x10000222, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_ImmUnsigned_16_20}},
	{EVSRWS, 0xfc0007ff, 0x10000221, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDD, 0xfc0007ff, 0x10000321, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSRWU, 0xfc0007ff, 0x10000220, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDDX, 0xfc0007ff, 0x10000320, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDH, 0xfc0007ff, 0x10000325, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTDW, 0xfc0007ff, 0x10000323, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTDHX, 0xfc0007ff, 0x10000324, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDWX, 0xfc0007ff, 0x10000322, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWHE, 0xfc0007ff, 0x10000331, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWHO, 0xfc0007ff, 0x10000335, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWWE, 0xfc0007ff, 0x10000339, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSTWHEX, 0xfc0007ff, 0x10000330, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWHOX, 0xfc0007ff, 0x10000334, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWWEX, 0xfc0007ff, 0x10000338, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTWWO, 0xfc0007ff, 0x1000033d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_16_20, ap_Reg_11_15}},
	{EVSUBFSMIAAW, 0xfc0007ff, 0x100004cb, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSTWWOX, 0xfc0007ff, 0x1000033c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSUBFSSIAAW, 0xfc0007ff, 0x100004c3, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFUMIAAW, 0xfc0007ff, 0x100004ca, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFUSIAAW, 0xfc0007ff, 0x100004c2, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVSUBFW, 0xfc0007ff, 0x10000204, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSUBIFW, 0xfc0007ff, 0x10000206, 0x0,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_11_15, ap_Reg_16_20}},
	{EVXOR, 0xfc0007ff, 0x10000216, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSABS, 0xfc0007ff, 0x10000284, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSNABS, 0xfc0007ff, 0x10000285, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSNEG, 0xfc0007ff, 0x10000286, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EVFSADD, 0xfc0007ff, 0x10000280, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSMUL, 0xfc0007ff, 0x10000288, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSSUB, 0xfc0007ff, 0x10000281, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSDIV, 0xfc0007ff, 0x10000289, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPGT, 0xfc0007ff, 0x1000028c, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPLT, 0xfc0007ff, 0x1000028d, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCMPEQ, 0xfc0007ff, 0x1000028e, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTGT, 0xfc0007ff, 0x1000029c, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTLT, 0xfc0007ff, 0x1000029d, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSTSTEQ, 0xfc0007ff, 0x1000029e, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EVFSCFSI, 0xfc0007ff, 0x10000291, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFSF, 0xfc0007ff, 0x10000293, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFUI, 0xfc0007ff, 0x10000290, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCFUF, 0xfc0007ff, 0x10000292, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSI, 0xfc0007ff, 0x10000295, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUI, 0xfc0007ff, 0x10000294, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSIZ, 0xfc0007ff, 0x1000029a, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUIZ, 0xfc0007ff, 0x10000298, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTSF, 0xfc0007ff, 0x10000297, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EVFSCTUF, 0xfc0007ff, 0x10000296, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSABS, 0xfc0007ff, 0x100002c4, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSNEG, 0xfc0007ff, 0x100002c6, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSNABS, 0xfc0007ff, 0x100002c5, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFSADD, 0xfc0007ff, 0x100002c0, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSMUL, 0xfc0007ff, 0x100002c8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSSUB, 0xfc0007ff, 0x100002c1, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSDIV, 0xfc0007ff, 0x100002c9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPGT, 0xfc0007ff, 0x100002cc, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPLT, 0xfc0007ff, 0x100002cd, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCMPEQ, 0xfc0007ff, 0x100002ce, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTGT, 0xfc0007ff, 0x100002dc, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTLT, 0xfc0007ff, 0x100002dd, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSTSTEQ, 0xfc0007ff, 0x100002de, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFSCFSI, 0xfc0007ff, 0x100002d1, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFSF, 0xfc0007ff, 0x100002d3, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSI, 0xfc0007ff, 0x100002d5, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFUI, 0xfc0007ff, 0x100002d0, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFUF, 0xfc0007ff, 0x100002d2, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUI, 0xfc0007ff, 0x100002d4, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSIZ, 0xfc0007ff, 0x100002da, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTSF, 0xfc0007ff, 0x100002d7, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUIZ, 0xfc0007ff, 0x100002d8, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCTUF, 0xfc0007ff, 0x100002d6, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDABS, 0xfc0007ff, 0x100002e4, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDNEG, 0xfc0007ff, 0x100002e6, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDNABS, 0xfc0007ff, 0x100002e5, 0xf800,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{EFDADD, 0xfc0007ff, 0x100002e0, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDMUL, 0xfc0007ff, 0x100002e8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDSUB, 0xfc0007ff, 0x100002e1, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDDIV, 0xfc0007ff, 0x100002e9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPGT, 0xfc0007ff, 0x100002ec, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPEQ, 0xfc0007ff, 0x100002ee, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCMPLT, 0xfc0007ff, 0x100002ed, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDTSTGT, 0xfc0007ff, 0x100002fc, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDTSTLT, 0xfc0007ff, 0x100002fd, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCFSI, 0xfc0007ff, 0x100002f1, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDTSTEQ, 0xfc0007ff, 0x100002fe, 0x600000,
		[5]*argField{ap_CondRegField_6_8, ap_Reg_11_15, ap_Reg_16_20}},
	{EFDCFUI, 0xfc0007ff, 0x100002f0, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFSID, 0xfc0007ff, 0x100002e3, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFSF, 0xfc0007ff, 0x100002f3, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFUF, 0xfc0007ff, 0x100002f2, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFUID, 0xfc0007ff, 0x100002e2, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSI, 0xfc0007ff, 0x100002f5, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUI, 0xfc0007ff, 0x100002f4, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSIDZ, 0xfc0007ff, 0x100002eb, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUIDZ, 0xfc0007ff, 0x100002ea, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSIZ, 0xfc0007ff, 0x100002fa, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTSF, 0xfc0007ff, 0x100002f7, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUF, 0xfc0007ff, 0x100002f6, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCTUIZ, 0xfc0007ff, 0x100002f8, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFDCFS, 0xfc0007ff, 0x100002ef, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{EFSCFD, 0xfc0007ff, 0x100002cf, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{DLMZB, 0xfc0007ff, 0x7c00009c, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{DLMZB_, 0xfc0007ff, 0x7c00009d, 0x0,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10, ap_Reg_16_20}},
	{MACCHW, 0xfc0007ff, 0x10000158, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHW_, 0xfc0007ff, 0x10000159, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWO, 0xfc0007ff, 0x10000558, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWO_, 0xfc0007ff, 0x10000559, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWS, 0xfc0007ff, 0x100001d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWS_, 0xfc0007ff, 0x100001d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSO, 0xfc0007ff, 0x100005d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSO_, 0xfc0007ff, 0x100005d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWU, 0xfc0007ff, 0x10000118, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWU_, 0xfc0007ff, 0x10000119, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWUO, 0xfc0007ff, 0x10000518, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWUO_, 0xfc0007ff, 0x10000519, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSU, 0xfc0007ff, 0x10000198, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSU_, 0xfc0007ff, 0x10000199, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSUO, 0xfc0007ff, 0x10000598, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACCHWSUO_, 0xfc0007ff, 0x10000599, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHW, 0xfc0007ff, 0x10000058, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHW_, 0xfc0007ff, 0x10000059, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWO, 0xfc0007ff, 0x10000458, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWO_, 0xfc0007ff, 0x10000459, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWS, 0xfc0007ff, 0x100000d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWS_, 0xfc0007ff, 0x100000d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSO, 0xfc0007ff, 0x100004d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSO_, 0xfc0007ff, 0x100004d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWU, 0xfc0007ff, 0x10000018, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWU_, 0xfc0007ff, 0x10000019, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWUO, 0xfc0007ff, 0x10000418, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWUO_, 0xfc0007ff, 0x10000419, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSU, 0xfc0007ff, 0x10000098, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSU_, 0xfc0007ff, 0x10000099, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSUO, 0xfc0007ff, 0x10000498, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACHHWSUO_, 0xfc0007ff, 0x10000499, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHW, 0xfc0007ff, 0x10000358, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHW_, 0xfc0007ff, 0x10000359, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWO, 0xfc0007ff, 0x10000758, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWO_, 0xfc0007ff, 0x10000759, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWS, 0xfc0007ff, 0x100003d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWS_, 0xfc0007ff, 0x100003d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSO, 0xfc0007ff, 0x100007d8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSO_, 0xfc0007ff, 0x100007d9, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWU, 0xfc0007ff, 0x10000318, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWU_, 0xfc0007ff, 0x10000319, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWUO, 0xfc0007ff, 0x10000718, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWUO_, 0xfc0007ff, 0x10000719, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHW, 0xfc0007ff, 0x10000150, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHW_, 0xfc0007ff, 0x10000151, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSU, 0xfc0007ff, 0x10000398, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSU_, 0xfc0007ff, 0x10000399, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSUO, 0xfc0007ff, 0x10000798, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MACLHWSUO_, 0xfc0007ff, 0x10000799, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHWU, 0xfc0007ff, 0x10000110, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULCHWU_, 0xfc0007ff, 0x10000111, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHW, 0xfc0007ff, 0x10000050, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHW_, 0xfc0007ff, 0x10000051, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHW, 0xfc0007ff, 0x10000350, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHW_, 0xfc0007ff, 0x10000351, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHWU, 0xfc0007ff, 0x10000010, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULHHWU_, 0xfc0007ff, 0x10000011, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHWU, 0xfc0007ff, 0x10000310, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{MULLHWU_, 0xfc0007ff, 0x10000311, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHW, 0xfc0007ff, 0x1000015c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHW_, 0xfc0007ff, 0x1000015d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWO, 0xfc0007ff, 0x1000055c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWO_, 0xfc0007ff, 0x1000055d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWS, 0xfc0007ff, 0x100001dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWS_, 0xfc0007ff, 0x100001dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWSO, 0xfc0007ff, 0x100005dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACCHWSO_, 0xfc0007ff, 0x100005dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHW, 0xfc0007ff, 0x1000005c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHW_, 0xfc0007ff, 0x1000005d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWO, 0xfc0007ff, 0x1000045c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWO_, 0xfc0007ff, 0x1000045d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWS, 0xfc0007ff, 0x100000dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWS_, 0xfc0007ff, 0x100000dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWSO, 0xfc0007ff, 0x100004dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACHHWSO_, 0xfc0007ff, 0x100004dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHW, 0xfc0007ff, 0x1000035c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHW_, 0xfc0007ff, 0x1000035d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWO, 0xfc0007ff, 0x1000075c, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWO_, 0xfc0007ff, 0x1000075d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWS, 0xfc0007ff, 0x100003dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWS_, 0xfc0007ff, 0x100003dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWSO, 0xfc0007ff, 0x100007dc, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{NMACLHWSO_, 0xfc0007ff, 0x100007dd, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBI, 0xfc0007fe, 0x7c0007ac, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{ICBT, 0xfc0007fe, 0x7c00002c, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBA, 0xfc0007fe, 0x7c0005ec, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBT, 0xfc0007fe, 0x7c00022c, 0x1,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_6_10}},
	{DCBT, 0xfc0007fe, 0x7c00022c, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTST, 0xfc0007fe, 0x7c0001ec, 0x1,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_6_10}},
	{DCBTST, 0xfc0007fe, 0x7c0001ec, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBZ, 0xfc0007fe, 0x7c0007ec, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBST, 0xfc0007fe, 0x7c00006c, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBF, 0xfc0007fe, 0x7c0000ac, 0x3800001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_9_10}},
	{ISYNC, 0xfc0007fe, 0x4c00012c, 0x3fff801,
		[5]*argField{}},
	{LBARX, 0xfc0007ff, 0x7c000068, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LBARX, 0xfc0007fe, 0x7c000068, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{LHARX, 0xfc0007ff, 0x7c0000e8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHARX, 0xfc0007fe, 0x7c0000e8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{LWARX, 0xfc0007ff, 0x7c000028, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWARX, 0xfc0007ff, 0x7c000028, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWARX, 0xfc0007fe, 0x7c000028, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STBCX_, 0xfc0007ff, 0x7c00056d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHCX_, 0xfc0007ff, 0x7c0005ad, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWCX_, 0xfc0007ff, 0x7c00012d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDARX, 0xfc0007ff, 0x7c0000a8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDARX, 0xfc0007fe, 0x7c0000a8, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STDCX_, 0xfc0007ff, 0x7c0001ad, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQARX, 0xfc0007ff, 0x7c000228, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LQARX, 0xfc0007fe, 0x7c000228, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_31_31}},
	{STQCX_, 0xfc0007ff, 0x7c00016d, 0x0,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SYNC, 0xfc0007fe, 0x7c0004ac, 0x390f801,
		[5]*argField{ap_ImmUnsigned_9_10, ap_ImmUnsigned_12_15}},
	{EIEIO, 0xfc0007fe, 0x7c0006ac, 0x3fff801,
		[5]*argField{}},
	{MBAR, 0xfc0007fe, 0x7c0006ac, 0x1ff801,
		[5]*argField{ap_ImmUnsigned_6_10}},
	{WAIT, 0xfc0007fe, 0x7c00007c, 0x39ff801,
		[5]*argField{ap_ImmUnsigned_9_10}},
	{TBEGIN_, 0xfc0007ff, 0x7c00051d, 0x1dff800,
		[5]*argField{ap_ImmUnsigned_10_10}},
	{TEND_, 0xfc0007ff, 0x7c00055d, 0x1fff800,
		[5]*argField{ap_ImmUnsigned_6_6}},
	{TABORT_, 0xfc0007ff, 0x7c00071d, 0x3e0f800,
		[5]*argField{ap_Reg_11_15}},
	{TABORTWC_, 0xfc0007ff, 0x7c00061d, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TABORTWCI_, 0xfc0007ff, 0x7c00069d, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_20}},
	{TABORTDC_, 0xfc0007ff, 0x7c00065d, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TABORTDCI_, 0xfc0007ff, 0x7c0006dd, 0x0,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_ImmSigned_16_20}},
	{TSR_, 0xfc0007ff, 0x7c0005dd, 0x3dff800,
		[5]*argField{ap_ImmUnsigned_10_10}},
	{TCHECK, 0xfc0007fe, 0x7c00059c, 0x7ff801,
		[5]*argField{ap_CondRegField_6_8}},
	{MFTB, 0xfc0007fe, 0x7c0002e6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{RFEBB, 0xfc0007fe, 0x4c000124, 0x3fff001,
		[5]*argField{ap_ImmUnsigned_20_20}},
	{LBDX, 0xfc0007fe, 0x7c000406, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHDX, 0xfc0007fe, 0x7c000446, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWDX, 0xfc0007fe, 0x7c000486, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDDX, 0xfc0007fe, 0x7c0004c6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LFDDX, 0xfc0007fe, 0x7c000646, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBDX, 0xfc0007fe, 0x7c000506, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHDX, 0xfc0007fe, 0x7c000546, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWDX, 0xfc0007fe, 0x7c000586, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDDX, 0xfc0007fe, 0x7c0005c6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDDX, 0xfc0007fe, 0x7c000746, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DSN, 0xfc0007fe, 0x7c0003c6, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{ECIWX, 0xfc0007fe, 0x7c00026c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ECOWX, 0xfc0007fe, 0x7c00036c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{SC, 0xfc000002, 0x44000002, 0x3fff01d,
		[5]*argField{ap_ImmUnsigned_20_26}},
	{RFID, 0xfc0007fe, 0x4c000024, 0x3fff801,
		[5]*argField{}},
	{HRFID, 0xfc0007fe, 0x4c000224, 0x3fff801,
		[5]*argField{}},
	{DOZE, 0xfc0007fe, 0x4c000324, 0x3fff801,
		[5]*argField{}},
	{NAP, 0xfc0007fe, 0x4c000364, 0x3fff801,
		[5]*argField{}},
	{SLEEP, 0xfc0007fe, 0x4c0003a4, 0x3fff801,
		[5]*argField{}},
	{RVWINKLE, 0xfc0007fe, 0x4c0003e4, 0x3fff801,
		[5]*argField{}},
	{LBZCIX, 0xfc0007fe, 0x7c0006aa, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWZCIX, 0xfc0007fe, 0x7c00062a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHZCIX, 0xfc0007fe, 0x7c00066a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDCIX, 0xfc0007fe, 0x7c0006ea, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBCIX, 0xfc0007fe, 0x7c0007aa, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWCIX, 0xfc0007fe, 0x7c00072a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHCIX, 0xfc0007fe, 0x7c00076a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDCIX, 0xfc0007fe, 0x7c0007ea, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TRECLAIM_, 0xfc0007ff, 0x7c00075d, 0x3e0f800,
		[5]*argField{ap_Reg_11_15}},
	{TRECHKPT_, 0xfc0007ff, 0x7c0007dd, 0x3fff800,
		[5]*argField{}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1,
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTMSR, 0xfc0007fe, 0x7c000124, 0x1ef801,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_15_15}},
	{MTMSRD, 0xfc0007fe, 0x7c000164, 0x1ef801,
		[5]*argField{ap_Reg_6_10, ap_ImmUnsigned_15_15}},
	{MFMSR, 0xfc0007fe, 0x7c0000a6, 0x1ff801,
		[5]*argField{ap_Reg_6_10}},
	{SLBIE, 0xfc0007fe, 0x7c000364, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{SLBIA, 0xfc0007fe, 0x7c0003e4, 0x31ff801,
		[5]*argField{ap_ImmUnsigned_8_10}},
	{SLBMTE, 0xfc0007fe, 0x7c000324, 0x1f0001,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBMFEV, 0xfc0007fe, 0x7c0006a6, 0x1f0001,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBMFEE, 0xfc0007fe, 0x7c000726, 0x1f0001,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{SLBFEE_, 0xfc0007ff, 0x7c0007a7, 0x1f0000,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{MTSR, 0xfc0007fe, 0x7c0001a4, 0x10f801,
		[5]*argField{ap_SpReg_12_15, ap_Reg_6_10}},
	{MTSRIN, 0xfc0007fe, 0x7c0001e4, 0x1f0001,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{MFSR, 0xfc0007fe, 0x7c0004a6, 0x10f801,
		[5]*argField{ap_Reg_6_10, ap_SpReg_12_15}},
	{MFSRIN, 0xfc0007fe, 0x7c000526, 0x1f0001,
		[5]*argField{ap_Reg_6_10, ap_Reg_16_20}},
	{TLBIE, 0xfc0007fe, 0x7c000264, 0x1f0001,
		[5]*argField{ap_Reg_16_20, ap_Reg_6_10}},
	{TLBIEL, 0xfc0007fe, 0x7c000224, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{TLBIA, 0xfc0007fe, 0x7c0002e4, 0x3fff801,
		[5]*argField{}},
	{TLBSYNC, 0xfc0007fe, 0x7c00046c, 0x3fff801,
		[5]*argField{}},
	{MSGSND, 0xfc0007fe, 0x7c00019c, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{MSGCLR, 0xfc0007fe, 0x7c0001dc, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{MSGSNDP, 0xfc0007fe, 0x7c00011c, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{MSGCLRP, 0xfc0007fe, 0x7c00015c, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{MTTMR, 0xfc0007fe, 0x7c0003dc, 0x1,
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{SC, 0xfc000002, 0x44000002, 0x3fffffd,
		[5]*argField{}},
	{RFI, 0xfc0007fe, 0x4c000064, 0x3fff801,
		[5]*argField{}},
	{RFCI, 0xfc0007fe, 0x4c000066, 0x3fff801,
		[5]*argField{}},
	{RFDI, 0xfc0007fe, 0x4c00004e, 0x3fff801,
		[5]*argField{}},
	{RFMCI, 0xfc0007fe, 0x4c00004c, 0x3fff801,
		[5]*argField{}},
	{RFGI, 0xfc0007fe, 0x4c0000cc, 0x3fff801,
		[5]*argField{}},
	{EHPRIV, 0xfc0007fe, 0x7c00021c, 0x1,
		[5]*argField{ap_ImmUnsigned_6_20}},
	{MTSPR, 0xfc0007fe, 0x7c0003a6, 0x1,
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MFSPR, 0xfc0007fe, 0x7c0002a6, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MTDCR, 0xfc0007fe, 0x7c000386, 0x1,
		[5]*argField{ap_SpReg_16_20_11_15, ap_Reg_6_10}},
	{MTDCRX, 0xfc0007fe, 0x7c000306, 0xf801,
		[5]*argField{ap_Reg_11_15, ap_Reg_6_10}},
	{MFDCR, 0xfc0007fe, 0x7c000286, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_16_20_11_15}},
	{MFDCRX, 0xfc0007fe, 0x7c000206, 0xf801,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15}},
	{MTMSR, 0xfc0007fe, 0x7c000124, 0x1ff801,
		[5]*argField{ap_Reg_6_10}},
	{MFMSR, 0xfc0007fe, 0x7c0000a6, 0x1ff801,
		[5]*argField{ap_Reg_6_10}},
	{WRTEE, 0xfc0007fe, 0x7c000106, 0x1ff801,
		[5]*argField{ap_Reg_6_10}},
	{WRTEEI, 0xfc0007fe, 0x7c000146, 0x3ff7801,
		[5]*argField{ap_ImmUnsigned_16_16}},
	{LBEPX, 0xfc0007fe, 0x7c0000be, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LHEPX, 0xfc0007fe, 0x7c00023e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LWEPX, 0xfc0007fe, 0x7c00003e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LDEPX, 0xfc0007fe, 0x7c00003a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STBEPX, 0xfc0007fe, 0x7c0001be, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STHEPX, 0xfc0007fe, 0x7c00033e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STWEPX, 0xfc0007fe, 0x7c00013e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STDEPX, 0xfc0007fe, 0x7c00013a, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBSTEP, 0xfc0007fe, 0x7c00007e, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTEP, 0xfc0007fe, 0x7c00027e, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBFEP, 0xfc0007fe, 0x7c0000fe, 0x3800001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20, ap_ImmUnsigned_9_10}},
	{DCBTSTEP, 0xfc0007fe, 0x7c0001fe, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBIEP, 0xfc0007fe, 0x7c0007be, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBZEP, 0xfc0007fe, 0x7c0007fe, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{LFDEPX, 0xfc0007fe, 0x7c0004be, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STFDEPX, 0xfc0007fe, 0x7c0005be, 0x1,
		[5]*argField{ap_FPReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVLDDEPX, 0xfc0007fe, 0x7c00063e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{EVSTDDEPX, 0xfc0007fe, 0x7c00073e, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEPX, 0xfc0007fe, 0x7c00024e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{LVEPXL, 0xfc0007fe, 0x7c00020e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEPX, 0xfc0007fe, 0x7c00064e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{STVEPXL, 0xfc0007fe, 0x7c00060e, 0x1,
		[5]*argField{ap_VecReg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBI, 0xfc0007fe, 0x7c0003ac, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{DCBLQ_, 0xfc0007ff, 0x7c00034d, 0x2000000,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBLQ_, 0xfc0007ff, 0x7c00018d, 0x2000000,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTLS, 0xfc0007fe, 0x7c00014c, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBTSTLS, 0xfc0007fe, 0x7c00010c, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBTLS, 0xfc0007fe, 0x7c0003cc, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICBLC, 0xfc0007fe, 0x7c0001cc, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{DCBLC, 0xfc0007fe, 0x7c00030c, 0x2000001,
		[5]*argField{ap_ImmUnsigned_7_10, ap_Reg_11_15, ap_Reg_16_20}},
	{TLBIVAX, 0xfc0007fe, 0x7c000624, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBILX, 0xfc0007fe, 0x7c000024, 0x3800001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBSX, 0xfc0007fe, 0x7c000724, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBSRX_, 0xfc0007ff, 0x7c0006a5, 0x3e00000,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{TLBRE, 0xfc0007fe, 0x7c000764, 0x3fff801,
		[5]*argField{}},
	{TLBSYNC, 0xfc0007fe, 0x7c00046c, 0x3fff801,
		[5]*argField{}},
	{TLBWE, 0xfc0007fe, 0x7c0007a4, 0x3fff801,
		[5]*argField{}},
	{DNH, 0xfc0007fe, 0x4c00018c, 0x1,
		[5]*argField{ap_ImmUnsigned_6_10, ap_ImmUnsigned_11_20}},
	{MSGSND, 0xfc0007fe, 0x7c00019c, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{MSGCLR, 0xfc0007fe, 0x7c0001dc, 0x3ff0001,
		[5]*argField{ap_Reg_16_20}},
	{DCI, 0xfc0007fe, 0x7c00038c, 0x21ff801,
		[5]*argField{ap_ImmUnsigned_7_10}},
	{ICI, 0xfc0007fe, 0x7c00078c, 0x21ff801,
		[5]*argField{ap_ImmUnsigned_7_10}},
	{DCREAD, 0xfc0007fe, 0x7c0003cc, 0x1,
		[5]*argField{ap_Reg_6_10, ap_Reg_11_15, ap_Reg_16_20}},
	{ICREAD, 0xfc0007fe, 0x7c0007cc, 0x3e00001,
		[5]*argField{ap_Reg_11_15, ap_Reg_16_20}},
	{MFPMR, 0xfc0007fe, 0x7c00029c, 0x1,
		[5]*argField{ap_Reg_6_10, ap_SpReg_11_20}},
	{MTPMR, 0xfc0007fe, 0x7c00039c, 0x1,
		[5]*argField{ap_SpReg_11_20, ap_Reg_6_10}},
}
