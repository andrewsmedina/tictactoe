package tictactoe

import "errors"

type Board struct {
	board  [9]uint8
	player uint8
}

func NewBoard() *Board {
	return &Board{board: [9]uint8{}, player: X}
}

func (b *Board) Move(position uint8) error {
	if b.board[position] != uint8(0) {
		return errors.New("This position is already occupied.")
	}
	b.board[position] = b.player
	b.changePlayer()
	return nil
}

func (b *Board) changePlayer() {
	if b.player == X {
		b.player = O
	} else {
		b.player = X
	}
}

func (b *Board) winner() bool {
	if b.board[0] == X && b.board[1] == X && b.board[2] == X {
		return true
	}
	return false
}

var X, O = uint8(1), uint8(2)
