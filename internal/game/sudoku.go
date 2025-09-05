package game

// Board - 盤面
type Board struct {
	Cells [9][9]Cell
}

// NewBoard 建立一個空的數獨盤面
func NewBoard() *Board {
	board := &Board{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			board.Cells[row][col] = Cell{Value: 0, Type: Empty}
		}
	}
	return board
}

// Game - 遊戲結構
type Game struct {
	board *Board
}

// NewGame - 建構遊戲結構
func NewGame() *Game {
	return &Game{
		board: NewBoard(),
	}
}
