package ssa

import (
	"bytes"
	"fmt"
	"io"
)

func printFunc(f *Func) {
	f.Logf("%s", f)
}

func (f *Func) String(psess *PackageSession) string {
	var buf bytes.Buffer
	p := stringFuncPrinter{w: &buf}
	psess.
		fprintFunc(p, f)
	return buf.String()
}

type funcPrinter interface {
	header(f *Func)
	startBlock(b *Block, reachable bool)
	endBlock(b *Block)
	value(v *Value, live bool)
	startDepCycle()
	endDepCycle()
	named(n LocalSlot, vals []*Value)
}

type stringFuncPrinter struct {
	w io.Writer
}

func (p stringFuncPrinter) header(f *Func) {
	fmt.Fprint(p.w, f.Name)
	fmt.Fprint(p.w, " ")
	fmt.Fprintln(p.w, f.Type)
}

func (p stringFuncPrinter) startBlock(b *Block, reachable bool) {
	fmt.Fprintf(p.w, "  b%d:", b.ID)
	if len(b.Preds) > 0 {
		io.WriteString(p.w, " <-")
		for _, e := range b.Preds {
			pred := e.b
			fmt.Fprintf(p.w, " b%d", pred.ID)
		}
	}
	if !reachable {
		fmt.Fprint(p.w, " DEAD")
	}
	io.WriteString(p.w, "\n")
}

func (p stringFuncPrinter) endBlock(psess *PackageSession, b *Block) {
	fmt.Fprintln(p.w, "    "+b.LongString(psess))
}

func (p stringFuncPrinter) value(psess *PackageSession, v *Value, live bool) {
	fmt.Fprint(p.w, "    ")

	fmt.Fprint(p.w, v.LongString(psess))
	if !live {
		fmt.Fprint(p.w, " DEAD")
	}
	fmt.Fprintln(p.w)
}

func (p stringFuncPrinter) startDepCycle() {
	fmt.Fprintln(p.w, "dependency cycle!")
}

func (p stringFuncPrinter) endDepCycle() {}

func (p stringFuncPrinter) named(n LocalSlot, vals []*Value) {
	fmt.Fprintf(p.w, "name %s: %v\n", n, vals)
}

func (psess *PackageSession) fprintFunc(p funcPrinter, f *Func) {
	reachable, live := psess.findlive(f)
	p.header(f)
	printed := make([]bool, f.NumValues())
	for _, b := range f.Blocks {
		p.startBlock(b, reachable[b.ID])

		if f.scheduled {

			for _, v := range b.Values {
				p.value(v, live[v.ID])
				printed[v.ID] = true
			}
			p.endBlock(b)
			continue
		}

		n := 0
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			p.value(v, live[v.ID])
			printed[v.ID] = true
			n++
		}

		for n < len(b.Values) {
			m := n
		outer:
			for _, v := range b.Values {
				if printed[v.ID] {
					continue
				}
				for _, w := range v.Args {

					if w != nil && w.Block == b && !printed[w.ID] {
						continue outer
					}
				}
				p.value(v, live[v.ID])
				printed[v.ID] = true
				n++
			}
			if m == n {
				p.startDepCycle()
				for _, v := range b.Values {
					if printed[v.ID] {
						continue
					}
					p.value(v, live[v.ID])
					printed[v.ID] = true
					n++
				}
				p.endDepCycle()
			}
		}

		p.endBlock(b)
	}
	for _, name := range f.Names {
		p.named(name, f.NamedValues[name])
	}
}
