package tictactoe

type Board struct {
	board [9]uint8
}

func NewBoard() *Board {
	return &Board{board: [9]uint8{}}
}
