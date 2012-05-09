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

func (s *S) TestBoardMove(c *C) {
	board := NewBoard()
	position := uint8(0)
	err := board.Move(position)
	c.Assert(err, IsNil)

	expected := [9]uint8{}
	expected[uint8(0)] = X
	err = board.Move(position)
	c.Assert(err, ErrorMatches, "This position is already occupied.")
}

func (s *S) TestBoardPlayer(c *C) {
	board := NewBoard()
	c.Assert(board.player, Equals, X)

	position := uint8(0)
	board.Move(position)
	c.Assert(board.player, Equals, O)
}

func (s *S) TestChangePlayer(c *C) {
	board := NewBoard()
	c.Assert(board.player, Equals, X)

	board.changePlayer()
	c.Assert(board.player, Equals, O)
}

func (s *S) TestWinner(c *C) {
	board := NewBoard()
	c.Assert(board.winner(), Equals, false)

	board.board[0] = X
	board.board[1] = X
	board.board[2] = X
	c.Assert(board.winner(), Equals, true)
}
