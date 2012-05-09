package tictactoe

import "errors"

type Board struct {
	board [9]uint8
}

func NewBoard() *Board {
	return &Board{board: [9]uint8{}}
}

func (b *Board) Move(piece, position uint8) error {
	if b.board[position] != uint8(0) {
		return errors.New("This position is already occupied.")
	}
	b.board[position] = piece
	return nil
}

type Player struct {
	t uint8
}

var X, O = uint8(1), uint8(2)
