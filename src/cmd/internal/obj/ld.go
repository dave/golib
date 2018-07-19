package obj

const (
	LOG = 5
)

func mkfwd(sym *LSym) {
	var dwn [LOG]int32
	var cnt [LOG]int32
	var lst [LOG]*Prog

	for i := 0; i < LOG; i++ {
		if i == 0 {
			cnt[i] = 1
		} else {
			cnt[i] = LOG * cnt[i-1]
		}
		dwn[i] = 1
		lst[i] = nil
	}

	i := 0
	for p := sym.Func.Text; p != nil && p.Link != nil; p = p.Link {
		i--
		if i < 0 {
			i = LOG - 1
		}
		p.Forwd = nil
		dwn[i]--
		if dwn[i] <= 0 {
			dwn[i] = cnt[i]
			if lst[i] != nil {
				lst[i].Forwd = p
			}
			lst[i] = p
		}
	}
}

func Appendp(q *Prog, newprog ProgAlloc) *Prog {
	p := newprog()
	p.Link = q.Link
	q.Link = p
	p.Pos = q.Pos
	return p
}
