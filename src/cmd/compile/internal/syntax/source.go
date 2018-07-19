package syntax

import (
	"io"
	"unicode/utf8"
)

// starting points for line and column numbers
const linebase = 1
const colbase = 1

type source struct {
	src  io.Reader
	errh func(line, pos uint, msg string)

	// source buffer
	buf         [4 << 10]byte
	offs        int   // source offset of buf
	r0, r, w    int   // previous/current read and write buf positions, excluding sentinel
	line0, line uint  // previous/current line
	col0, col   uint  // previous/current column (byte offsets from line start)
	ioerr       error // pending io error

	// literal buffer
	lit []byte // literal prefix
	suf int    // literal suffix; suf >= 0 means we are scanning a literal
}

// init initializes source to read from src and to report errors via errh.
// errh must not be nil.
func (s *source) init(src io.Reader, errh func(line, pos uint, msg string)) {
	s.src = src
	s.errh = errh

	s.buf[0] = utf8.RuneSelf
	s.offs = 0
	s.r0, s.r, s.w = 0, 0, 0
	s.line0, s.line = 0, linebase
	s.col0, s.col = 0, colbase
	s.ioerr = nil

	s.lit = s.lit[:0]
	s.suf = -1
}

// ungetr ungets the most recently read rune.
func (s *source) ungetr() {
	s.r, s.line, s.col = s.r0, s.line0, s.col0
}

// ungetr2 is like ungetr but enables a 2nd ungetr.
// It must not be called if one of the runes seen
// was a newline.
func (s *source) ungetr2() {
	s.ungetr()

	s.r0--
	s.col0--
}

func (s *source) error(msg string) {
	s.errh(s.line0, s.col0, msg)
}

// getr reads and returns the next rune.
//
// If a read or source encoding error occurs, getr
// calls the error handler installed with init.
// The handler must exist.
//
// The (line, col) position passed to the error handler
// is always at the current source reading position.
func (s *source) getr() rune {
redo:
	s.r0, s.line0, s.col0 = s.r, s.line, s.col

	for s.r+utf8.UTFMax > s.w && !utf8.FullRune(s.buf[s.r:s.w]) && s.ioerr == nil && s.w-s.r < len(s.buf) {
		s.fill()
	}

	if b := s.buf[s.r]; b < utf8.RuneSelf {
		s.r++

		s.col++
		if b == 0 {
			s.error("invalid NUL character")
			goto redo
		}
		if b == '\n' {
			s.line++
			s.col = colbase
		}
		return rune(b)
	}

	if s.r == s.w {
		if s.ioerr != io.EOF {

			s.error("I/O error: " + s.ioerr.Error())
		}
		return -1
	}

	r, w := utf8.DecodeRune(s.buf[s.r:s.w])
	s.r += w
	s.col += uint(w)

	if r == utf8.RuneError && w == 1 {
		s.error("invalid UTF-8 encoding")
		goto redo
	}

	// BOM's are only allowed as the first character in a file
	const BOM = 0xfeff
	if r == BOM {
		if s.r0 > 0 {
			s.error("invalid BOM in the middle of the file")
		}
		goto redo
	}

	return r
}

func (s *source) fill() {

	if s.r0 > 1 {

		if s.suf >= 0 {
			s.lit = append(s.lit, s.buf[s.suf:s.r0]...)
			s.suf = 1
		}
		n := s.r0 - 1
		copy(s.buf[:], s.buf[n:s.w])
		s.offs += n
		s.r0 = 1
		s.r -= n
		s.w -= n
	}

	for i := 100; i > 0; i-- {
		n, err := s.src.Read(s.buf[s.w : len(s.buf)-1])
		if n < 0 {
			panic("negative read")
		}
		s.w += n
		if n > 0 || err != nil {
			s.buf[s.w] = utf8.RuneSelf
			if err != nil {
				s.ioerr = err
			}
			return
		}
	}

	s.ioerr = io.ErrNoProgress
}

func (s *source) startLit() {
	s.suf = s.r0
	s.lit = s.lit[:0]
}

func (s *source) stopLit() []byte {
	lit := s.buf[s.suf:s.r]
	if len(s.lit) > 0 {
		lit = append(s.lit, lit...)
	}
	s.killLit()
	return lit
}

func (s *source) killLit() {
	s.suf = -1
}
