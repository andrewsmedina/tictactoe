package tictactoe

type Board struct {
	board [9]uint8
}

func NewBoard() *Board {
	return &Board{board: [9]uint8{}}
}

type Player struct {
	t uint8
}

var X, O = uint8(1), uint8(2)
