package tictactoe

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestNewBoard(c *C) {
	board := NewBoard()
	expected := [9]uint8{}
	c.Assert(board.board, DeepEquals, expected)
}

func (s *S) TestX(c *C) {
	c.Assert(X, Equals, uint8(1))
}

func (s *S) TestO(c *C) {
	c.Assert(O, Equals, uint8(2))
}
