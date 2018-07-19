package objabi

import "runtime"

const defaultGO386 = "sse2"
const defaultGOARM = "7"
const defaultGOMIPS = "hardfloat"
const defaultGOMIPS64 = "hardfloat"
const defaultGOOS = runtime.GOOS
const defaultGOARCH = runtime.GOARCH
const defaultGO_EXTLINK_ENABLED = ""
const version = "devel +4991bc6257 Fri Jun 22 00:28:25 2018 +0000"
const stackGuardMultiplier = 1
const goexperiment = ""
