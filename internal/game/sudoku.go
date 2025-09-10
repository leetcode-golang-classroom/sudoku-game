package game

import "time"

const (
	BoardSize = 9
	BoxSize   = 3
)

// Board - 盤面
type Board struct {
	Cells             [BoardSize][BoardSize]*Cell // 格子內容
	CursorRow         int                         // Cursor row position
	CursorCol         int                         // Cursor col position
	TargetSolvedCount int                         // 需要解決的格子數 = 81 - clues
	FilledCount       int                         // 目前填入格子數
	ConflictCount     int                         // 不符合規則的格子數
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
	Board     *Board
	StartTime time.Time // 遊戲開始時間
}

// NewGame - 建構遊戲結構
func NewGame() *Game {
	return &Game{
		Board:     NewBoard(),
		StartTime: time.Now().UTC(),
	}
}

func (game *Game) GetElaspedTime() int {
	return int(time.Since(game.StartTime).Seconds())
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

func (board *Board) IncreaseConflictCount() {
	board.ConflictCount++
}

func (board *Board) DescreaseConflictCount() {
	board.ConflictCount--
}

func (board *Board) IncreaseFilledCount() {
	board.FilledCount++
}

func (board *Board) DecreaseFilledCount() {
	board.FilledCount--
}

// ResetBoardToDefault - 重開始功能
func (board *Board) ResetBoardToDefault() {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if board.Cells[row][col].Type != Preset {
				board.Cells[row][col].Type = Empty
				board.Cells[row][col].Value = 0
			}
		}
	}
	board.FilledCount = 0
	board.ConflictCount = 0
}
