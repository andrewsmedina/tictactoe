package tictactoe

type Board struct {
	board [9]uint8
}

func NewBoard() *Board {
	return &Board{board: [9]uint8{}}
}

func (b *Board) Move(piece, position uint8) {
	b.board[position] = piece
}

type Player struct {
	t uint8
}

var X, O = uint8(1), uint8(2)
