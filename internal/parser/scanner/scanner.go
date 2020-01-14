package scanner

import (
	"io"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

type Scanner interface {
	HasNext() bool
	Next() token.Token
	io.Closer
}

type scanner struct {
	input []rune
	start int
	pos   int

	closed bool
}

func New(input []rune) Scanner {
	return &scanner{
		input: input,
		start: 0,
		pos:   0,

		closed: false,
	}
}

func (s *scanner) HasNext() bool {
	if s.done() {
		return false
	}

	panic("implement me")
}

func (s *scanner) Next() token.Token {
	/*
		Read the next token. This is basically starting from the initial state until a
		token gets emitted. If an error occurs, simply return an error token. Never
		panic. NEVER. PANIC. Errors are no exception, but common. You have to think of
		errors as a language structure that we have to support.
	*/
	panic("implement me")
}

// Close will cause this scanner to not execute any more states. The execution
// of the current state cannot be aborted, but the scanner will stop executing
// states after the current state has finished.
func (s *scanner) Close() error {
	s.closed = true
	return nil
}

// done determines whether the scanner is done with its work. This is the case,
// if either the scanner was closed, or the scanner has reached the end of its
// input.
func (s *scanner) done() bool {
	return s.closed ||
		s.pos >= len(s.input)
}

// next returns the next rune of the scanners input and advances its pointer by
// one position. This method also updates the line and col information of the
// scanner. If the scanner.done()==true and this method is called, it will panic
// with a syntax error.
//
// The process of advancing the pointer and returning the read rune is called
// "consuming a rune" or "accepting a rune".
func (s *scanner) next() rune {
	// get the actual next rune
	next := s.input[s.pos]

	// update current scanner position
	s.pos++

	return next
}

// peek returns the next rune of the scanners input without consuming it.
func (s *scanner) peek() rune {
	return s.input[s.pos]
}

// goback decrements the scanner's position by one and updates its line and col
// information.
func (s *scanner) goback() {
	s.pos--
}

// ignore discards all accepted runes. This is done by simply setting the start
// position of the last read token to the current scanner position.
func (s *scanner) ignore() {
	s.start = s.pos
}

// accept accepts exactly one rune matched by the given matcher. This means,
// that: If the next rune is matched by the scanner, it is consumed and ok=true
// is returned. If the next rune is NOT matched, it is unread and ok=false is
// returned. This implies, that accept(Alphanumeric) will actually do nothing if
// the next rune is not Alphanumeric. However, if the next rune is Alphanumeric,
// it will be accepted.
func (s *scanner) accept(m matcher.M) (ok bool) {
	if m.Matches(s.next()) {
		return true
	}
	s.goback()
	return false
}

// acceptMultiple accepts multiple runes that are matched by the given matcher.
// See the godoc on (*scanner).accept for more information. The amount of
// matched runes is returned.
func (s *scanner) acceptMultiple(m matcher.M) (matched uint) {
	for s.accept(m) {
		matched++
	}
	return
}

// acceptString accepts the exact sequence of runes that the given string
// represents, or does nothing, if the string is not matched.
//
//	input := []rune(".hello")
//	...
//	s.acceptString("hello") // will do nothing, as the next rune is '.'
//	s.next()                // advance the position by one (next rune is now 'h')
//	s.acceptString("hello") // will accept 5 runes, the scanner has reached its EOF now
func (s *scanner) acceptString(str string) bool {
	if s.peekString(str) {
		s.pos += len(str)
		return true
	}
	return false
}

// peekString works like (*scanner).acceptString, except it doesn't consume any
// runes. It just peeks, if the given string lays ahead.
func (s *scanner) peekString(str string) bool {
	for i, r := range str {
		if r != s.input[s.pos+i] {
			return false
		}
	}
	return true
}
