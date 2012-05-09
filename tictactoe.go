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
	//top
	if b.board[0] != uint8(0) && b.board[0] == b.board[1] && b.board[1] == b.board[2] {
		return true
	}
	//bottom
	if b.board[6] != uint8(0) && b.board[6] == b.board[7] && b.board[7] == b.board[8] {
		return true
	}
	//center
	if b.board[3] != uint8(0) && b.board[3] == b.board[4] && b.board[4] == b.board[5] {
		return true
	}
	//left
	if b.board[0] != uint8(0) && b.board[0] == b.board[3] && b.board[3] == b.board[6] {
		return true
	}
	//middle
	if b.board[1] != uint8(0) && b.board[1] == b.board[4] && b.board[4] == b.board[7] {
		return true
	}
	//right
	if b.board[2] != uint8(0) && b.board[2] == b.board[5] && b.board[5] == b.board[8] {
		return true
	}
	//diagonals
	if b.board[4] != uint8(0) && ((b.board[0] == b.board[4] && b.board[4] == b.board[8]) || (b.board[2] == b.board[4] && b.board[4] == b.board[6])) {
		return true
	}
	return false
}

var X, O = uint8(1), uint8(2)
