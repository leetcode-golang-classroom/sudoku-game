package game

const (
	BoardSize = 9
	BoxSize   = 3
)

// Board - 盤面
type Board struct {
	Cells     [BoardSize][BoardSize]*Cell
	CursorRow int
	CursorCol int
}

// NewBoard 建立一個空的數獨盤面
func NewBoard() *Board {
	board := &Board{}
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			board.Cells[row][col] = &Cell{Value: 0, Type: Empty}
		}
	}
	return board
}

// Game - 遊戲結構
type Game struct {
	Board *Board
}

// NewGame - 建構遊戲結構
func NewGame() *Game {
	return &Game{
		Board: NewBoard(),
	}
}

func (board *Board) IncreaseCursorRow() {
	if board.CursorRow < BoardSize-1 {
		board.CursorRow++
		return
	}
}

func (board *Board) DecreaseCursorRow() {
	if board.CursorRow >= 1 {
		board.CursorRow--
		return
	}
}

func (board *Board) IncreaseCursorCol() {
	if board.CursorCol < BoardSize-1 {
		board.CursorCol++
		return
	}
}

func (board *Board) DecreaseCursorCol() {
	if board.CursorCol >= 1 {
		board.CursorCol--
		return
	}
}
