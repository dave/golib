package syntax

import (
	"fmt"
	"io"
	"reflect"
	"unicode"
	"unicode/utf8"
)

// Fdump dumps the structure of the syntax tree rooted at n to w.
// It is intended for debugging purposes; no specific output format
// is guaranteed.
func Fdump(w io.Writer, n Node) (err error) {
	p := dumper{
		output: w,
		ptrmap: make(map[Node]int),
		last:   '\n',
	}

	defer func() {
		if e := recover(); e != nil {
			err = e.(localError).err
		}
	}()

	if n == nil {
		p.printf("nil\n")
		return
	}
	p.dump(reflect.ValueOf(n), n)
	p.printf("\n")

	return
}

type dumper struct {
	output io.Writer
	ptrmap map[Node]int // node -> dump line number
	indent int          // current indentation level
	last   byte         // last byte processed by Write
	line   int          // current line number
}

func (p *dumper) Write(psess *PackageSession, data []byte) (n int, err error) {
	var m int
	for i, b := range data {

		if b == '\n' {
			m, err = p.output.Write(data[n : i+1])
			n += m
			if err != nil {
				return
			}
		} else if p.last == '\n' {
			p.line++
			_, err = fmt.Fprintf(p.output, "%6d  ", p.line)
			if err != nil {
				return
			}
			for j := p.indent; j > 0; j-- {
				_, err = p.output.Write(psess.indentBytes)
				if err != nil {
					return
				}
			}
		}
		p.last = b
	}
	if len(data) > n {
		m, err = p.output.Write(data[n:])
		n += m
	}
	return
}

// localError wraps locally caught errors so we can distinguish
// them from genuine panics which we don't want to return as errors.
type localError struct {
	err error
}

// printf is a convenience wrapper that takes care of print errors.
func (p *dumper) printf(format string, args ...interface{}) {
	if _, err := fmt.Fprintf(p, format, args...); err != nil {
		panic(localError{err})
	}
}

// dump prints the contents of x.
// If x is the reflect.Value of a struct s, where &s
// implements Node, then &s should be passed for n -
// this permits printing of the unexported span and
// comments fields of the embedded isNode field by
// calling the Span() and Comment() instead of using
// reflection.
func (p *dumper) dump(x reflect.Value, n Node) {
	switch x.Kind() {
	case reflect.Interface:
		if x.IsNil() {
			p.printf("nil")
			return
		}
		p.dump(x.Elem(), nil)

	case reflect.Ptr:
		if x.IsNil() {
			p.printf("nil")
			return
		}

		if x, ok := x.Interface().(*Name); ok {
			p.printf("%s @ %v", x.Value, x.Pos())
			return
		}

		p.printf("*")

		if ptr, ok := x.Interface().(Node); ok {
			if line, exists := p.ptrmap[ptr]; exists {
				p.printf("(Node @ %d)", line)
				return
			}
			p.ptrmap[ptr] = p.line
			n = ptr
		}
		p.dump(x.Elem(), n)

	case reflect.Slice:
		if x.IsNil() {
			p.printf("nil")
			return
		}
		p.printf("%s (%d entries) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.dump(x.Index(i), nil)
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Struct:
		typ := x.Type()

		p.printf("%s {", typ)
		p.indent++

		first := true
		if n != nil {
			p.printf("\n")
			first = false

		}

		for i, n := 0, typ.NumField(); i < n; i++ {

			if name := typ.Field(i).Name; isExported(name) {
				if first {
					p.printf("\n")
					first = false
				}
				p.printf("%s: ", name)
				p.dump(x.Field(i), nil)
				p.printf("\n")
			}
		}

		p.indent--
		p.printf("}")

	default:
		switch x := x.Interface().(type) {
		case string:

			p.printf("%q", x)
		default:
			p.printf("%v", x)
		}
	}
}

func isExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}
