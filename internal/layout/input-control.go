package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

// DetectInput - 處理鍵盤輸入
func (gameLayout *GameLayout) DetectInput() {
	board := gameLayout.gameInstance.Board
	targetRow := board.CursorRow
	targetCol := board.CursorCol
	targetCell := board.Cells[targetRow][targetCol]
	// 偵測合法數字輸入
	// 數字輸入
	for key := ebiten.KeyDigit1; key <= ebiten.KeyDigit9; key++ {
		if inpututil.IsKeyJustPressed(key) {
			handleKeyInput(board, targetCell, key, targetRow, targetCol)
			return
		}
	}

	// 清除輸入
	if inpututil.IsKeyJustPressed(ebiten.Key0) || inpututil.IsKeyJustPressed(ebiten.KeyDelete) {
		handleClearInput(board, targetCell, targetRow, targetCol)
		return
	}
}

// handleClearInput - 處理清除
func handleClearInput(board *game.Board, targetCell *game.Cell,
	targetRow, targetCol int) {
	cellType := targetCell.Type
	// 當遇到題目時
	if cellType == game.Preset {
		return
	}
	// 原本輸入是 conflict
	if cellType == game.InputConflict {
		board.DescreaseConflictCount()
	}
	// 原本輸入非空
	if cellType != game.Empty {
		board.DecreaseFilledCount()
	}
	// 清空目前 Cell 的值
	board.Cells[targetRow][targetCol].Value = 0
	board.Cells[targetRow][targetCol].Type = game.Empty
}

// handleKeyInput - 處理輸入時
func handleKeyInput(board *game.Board, targetCell *game.Cell, key ebiten.Key,
	targetRow, targetCol int) {
	cellType := targetCell.Type
	// 當格子為題目時
	if cellType == game.Preset {
		return
	}
	value := int(key - ebiten.KeyDigit0)
	// 當輸入格為空格時
	if cellType == game.Empty {
		board.IncreaseFilledCount()
	}
	safed := board.IsSafe(targetRow, targetCol, value)
	if !safed {
		handleConflict(board, cellType, targetRow, targetCol)
	} else {
		handleNonConflict(board, cellType, targetRow, targetCol)
	}
	// 更新輸入
	board.Cells[targetRow][targetCol].Value = value
}

// handleConflict - 處理 Conflict Cell
func handleConflict(board *game.Board, cellType game.CellType,
	targetRow, targetCol int) {
	if cellType != game.InputConflict {
		board.IncreaseConflictCount()
	}
	// 標示為 Conflict Input
	board.Cells[targetRow][targetCol].Type = game.InputConflict
}

// handleNonConflict - 處理 Non-Conflict Cell
func handleNonConflict(board *game.Board, cellType game.CellType,
	targetRow, targetCol int) {
	// 當輸入為 Conflict 時
	if cellType == game.InputConflict {
		board.DescreaseConflictCount()
	}
	// 標示為 Input
	board.Cells[targetRow][targetCol].Type = game.Input
}
