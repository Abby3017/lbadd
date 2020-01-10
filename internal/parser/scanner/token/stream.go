package token

import "io"

type Stream interface {
	Peek() Token
	Take() Token
	Push(Token)

	io.Closer
	Closed() bool
}

type stream struct {
	ch chan Token

	peeked *Token

	closed bool
}

func NewStream() Stream {
	return &stream{
		ch:     make(chan Token, 5),
		peeked: nil,
		closed: false,
	}
}

// Peek returns the first element of the stream, waiting for it to become
// available if necessary. However, it does NOT remove the element from the
// stream.
func (s *stream) Peek() Token {
	/*
		The godoc says that this method does not remove a token from the stream.
		That is correct, considering that this object is the stream. However,
		this method may remove a token from the stream's channel. This is a
		subtle, but major difference.
	*/
	s.ensureNotClosed()
	if s.peeked != nil {
		return *s.peeked
	}
	t := s.Take()
	s.peeked = &t
	return t
}

// Take returns the first element of the stream, waiting for it to become
// available if necessary. Take also removes the element from the stream.
func (s *stream) Take() Token {
	s.ensureNotClosed()
	if s.peeked != nil {
		t := *s.peeked
		s.peeked = nil
		return t
	}
	return <-s.ch
}

// Push pushes a token onto the stream, waiting if the stream is full.
func (s *stream) Push(t Token) {
	s.ensureNotClosed()
	s.ch <- t
}

func (s stream) ensureNotClosed() {
	if s.closed {
		panic("token stream is closed")
	}
}

func (s *stream) Close() error {
	s.closed = true
	close(s.ch)
	return nil
}

func (s stream) Closed() bool {
	return s.closed
}
