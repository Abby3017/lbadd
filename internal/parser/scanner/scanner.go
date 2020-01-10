package scanner

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// state is a type alias for a function that takes a scanner and returns another
// state. Such functions (or states) will be invoked by the scanner. It will
// pass itself as the argument, and the returned state will be chained and
// executed next. If a state evaluates to nil, the initial state will be
// chained, if the scanner has not reached the EOF yet.
//
// It is a state's responsibility to reset the scanner's position etc. to the
// state they were in, when the state was entered, if the state does not
// evaluate to nil. A state can do this by using checkpoints as shown in the
// following example.
//
//	func myState(s *Scanner) state {
//		chck := s.checkpoint()
//		if s.accept(MyMatcher) {
//			// maybe emit a token here
//			return nil
//		}
//		s.restore(chck) // resets start, pos, line, col and the other position attributes
//		return errorf("myerror")
//	}
type state func(*Scanner) state

func (st state) String() string {
	return runtime.FuncForPC(reflect.ValueOf(st).Pointer()).Name()
}

// checkpoint represents the state of a scanner at any given point in time. A
// scanner can be restored to a checkpoint.
//
//	var s *Scanner
//	...
//	chck := s.checkpoint() // create a checkpoint
//	...                     // accept(), next(), goback(), whatever
//	s.restore(chck)        // scanner is in the same state as when the checkpoint was created
//
// This is useful when a state should not return an error if something
// unexpected was read, but for example another grammar production should be
// tried. To guarantee that both tries start with the same scanner state, a
// checkpoint can be used.
type checkpoint struct {
	start int
	pos   int

	startLine, startCol int
	line, lastCol, col  int
}

type Scanner struct {
	input []rune
	start int
	pos   int

	current state
	stream  token.Stream

	startLine, startCol int
	line, lastCol, col  int

	closed bool
	doneCh chan struct{}
}

// New creates a ready to use scanner with the given input data, that will push
// created tokens onto the given token stream. After all input has been handled,
// the scanner will emit a final EOF token.
func New(input []rune, stream token.Stream) *Scanner {
	return &Scanner{
		input: input,
		start: 0,
		pos:   0,

		current: initial,
		stream:  stream,

		startLine: 1,
		startCol:  1,
		line:      1, // line starts at 1, because it should be human readable and editor line and column numbers usually start at 1
		lastCol:   1,
		col:       1, // col starts at 1, because it should be human readable and editor line and column numbers usually start at 1

		doneCh: make(chan struct{}),
	}
}

// Scan starts the scanning process. This method will block until the input has
// been handled, so consider calling it in a separate goroutine. Keep in mind,
// that the default tokn.Stream implementation has a buffer of 5 tokens, so this
// might block if the token stream is full.
func (s *Scanner) Scan() {
	defer func() {
		// ignore all accepted runes that are not emitted as a token when
		// reaching the EOF.
		s.ignore()
		// EOF must be emitted after recovering from a crash, so that the parser
		// doesn't miss the error
		s.emit(token.EOF)
		// the scanner is now done with its work
		close(s.doneCh)
	}()

	// recover from syntax errors that cannot be handled before EOF is emitted
	defer func() {
		if recovered := recover(); recovered != nil {
			if err, ok := recovered.(SyntaxError); ok {
				s.stream.Push(token.New(s.line, s.col, s.pos, s.pos-s.start, token.Error, fmt.Sprintf("recovered: %v", err)))
			} else {
				panic(recovered) // re-panic if it's not a syntax error
			}
		}
	}()

	// execute the next state while the scanner is not done
	for !s.done() {
		s.executeCurrentState() // for executing the next state, use the same method that testing uses
		if s.current == nil {
			s.current = initial
		}
	}
}

// Done returns a channel, that will be closed when the scanner is done or has
// been closed, AFTER the EOF token was emitted.
func (s *Scanner) Done() <-chan struct{} {
	return s.doneCh
}

// Close will cause this scanner to not execute any more states. The execution
// of the current state cannot be aborted, but the scanner will stop executing
// states after the current state has finished.
func (s *Scanner) Close() error {
	s.closed = true
	return nil
}

// executeCurrentState executes the current scanner state and re-assigns it to
// the output of the current state. This method was refactored out for more
// testability.
func (s *Scanner) executeCurrentState() {
	s.current = s.current(s)
}

// done determines whether the scanner is done with its work. This is the case,
// if either the scanner was closed, or the scanner has reached the end of its
// input.
func (s *Scanner) done() bool {
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
func (s *Scanner) next() rune {
	if s.done() {
		panic(SyntaxError{
			offset:  s.pos,
			line:    s.line,
			col:     s.col,
			message: fmt.Sprintf("state '%v' tried to read another rune, but scanner already reached EOF at offset %d (%d:%d)", s.current, s.pos, s.line, s.col),
		})
	}

	// get the actual next rune
	next := s.input[s.pos]

	// update line and column information
	if next == '\n' { // TODO if next is line delimiter
		s.line++
		s.lastCol = s.col
		s.col = 1
	} else {
		s.col++
	}

	// update current scanner position
	s.pos++

	return next
}

// peek returns the next rune of the scanners input without consuming it.
func (s *Scanner) peek() rune {
	return s.input[s.pos]
}

// goback decrements the scanner's position by one and updates its line and col
// information.
func (s *Scanner) goback() {
	s.pos--

	// update line and column information
	if s.col == 1 {
		s.line--
		s.col = s.lastCol
	} else {
		s.col--
	}
}

// ignore discards all accepted runes. This is done by simply setting the start
// position of the last read token to the current scanner position.
func (s *Scanner) ignore() {
	s.start = s.pos
	s.startLine = s.line
	s.startCol = s.col
}

// checkpoint returns a checkpoint, which represents the complete scanner state
// that is required to return the scanner to the exact position it is in right
// now.
func (s *Scanner) checkpoint() checkpoint {
	return checkpoint{
		start:     s.start,
		pos:       s.pos,
		startLine: s.startLine,
		startCol:  s.startCol,
		line:      s.line,
		lastCol:   s.lastCol,
		col:       s.col,
	}
}

// restore restores the scanner's state to the exact state it was in, when the
// checkpoint was created.
func (s *Scanner) restore(chck checkpoint) {
	s.start = chck.start
	s.pos = chck.pos
	s.startLine = chck.startLine
	s.startCol = chck.startCol
	s.line = chck.line
	s.lastCol = chck.lastCol
	s.col = chck.col
}

// emit pushes a token with the given token type onto the scanner's token
// stream. Afterwards, it sets the start position of the last token to the
// current scanner position.
//
// The following pseudo code illustrates, how a token is created.
//
//	func emit(t token.Type) {
//		tk := token.Token{
//			typ:           t,
//			startPosition: s.start,
//			length:        s.pos-s.start,
//			value:         s.input[s.start:s.pos],
//		}
//		pushToken(tk)
//	}
func (s *Scanner) emit(t token.Type) {
	tok := token.New(s.line, s.col, s.start, s.pos-s.start, t, string(s.input[s.start:s.pos]))
	s.stream.Push(tok)

	s.start = s.pos
	s.startLine = s.line
	s.startCol = s.col
}

// accept accepts exactly one rune matched by the given matcher. This means,
// that: If the next rune is matched by the scanner, it is consumed and ok=true
// is returned. If the next rune is NOT matched, it is unread and ok=false is
// returned. This implies, that accept(Alphanumeric) will actually do nothing if
// the next rune is not Alphanumeric. However, if the next rune is Alphanumeric,
// it will be accepted.
func (s *Scanner) accept(m matcher.M) (ok bool) {
	if m.Matches(s.next()) {
		return true
	}
	s.goback()
	return false
}

// acceptMultiple accepts multiple runes that are matched by the given matcher.
// See the godoc on (*Scanner).accept for more information. The amount of
// matched runes is returned.
func (s *Scanner) acceptMultiple(m matcher.M) (matched uint) {
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
func (s *Scanner) acceptString(str string) bool {
	chck := s.checkpoint()
	for _, r := range str {
		if r != s.next() {
			s.restore(chck)
			return false
		}
	}
	return true
}

// peekString works like (*Scanner).acceptString, except it doesn't consume any
// runes. It just peeks, if the given string lays ahead.
func (s *Scanner) peekString(str string) bool {
	chck := s.checkpoint()
	defer s.restore(chck)

	for _, r := range str {
		if r != s.next() {
			return false
		}
	}
	return true
}
