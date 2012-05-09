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
