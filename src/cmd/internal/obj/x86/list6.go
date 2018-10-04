// Inferno utils/6c/list.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6c/list.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package x86

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
)

func (pstate *PackageState) init() {
	pstate.obj.RegisterRegister(REG_AL, REG_AL+len(pstate.Register), pstate.rconv)
	pstate.obj.RegisterOpcode(obj.ABaseAMD64, pstate.Anames)
	pstate.obj.RegisterRegisterList(obj.RegListX86Lo, obj.RegListX86Hi, pstate.rlconv)
	pstate.obj.RegisterOpSuffix("386", pstate.opSuffixString)
	pstate.obj.RegisterOpSuffix("amd64", pstate.opSuffixString)
}

func (pstate *PackageState) rconv(r int) string {
	if REG_AL <= r && r-REG_AL < len(pstate.Register) {
		return pstate.Register[r-REG_AL]
	}
	return fmt.Sprintf("Rgok(%d)", r-obj.RBaseAMD64)
}

func (pstate *PackageState) rlconv(bits int64) string {
	reg0, reg1 := decodeRegisterRange(bits)
	return fmt.Sprintf("[%s-%s]", pstate.rconv(reg0), pstate.rconv(reg1))
}

func (pstate *PackageState) opSuffixString(s uint8) string {
	return "." + opSuffix(s).String(pstate)
}
