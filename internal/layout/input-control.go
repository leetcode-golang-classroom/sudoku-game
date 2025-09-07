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
			if targetCell.Type != game.Preset {
				value := int(key - ebiten.KeyDigit0)
				// 檢查輸入的值是否為放入是否能會造成 Conflict
				if !board.IsSafe(targetRow, targetCol, value) {
					// 標示為 Conflict Input
					board.Cells[targetRow][targetCol].Type = game.InputConflict
				} else {
					board.Cells[targetRow][targetCol].Type = game.Input
				}
				// 更新輸入
				board.Cells[targetRow][targetCol].Value = value
			}
			return
		}
	}

	// 清除輸入
	if inpututil.IsKeyJustPressed(ebiten.Key0) || inpututil.IsKeyJustPressed(ebiten.KeyDelete) {
		if targetCell.Type != game.Preset {
			// 清空目前 Cell 的值
			board.Cells[targetRow][targetCol].Value = 0
			board.Cells[targetRow][targetCol].Type = game.Empty
		}
		return
	}
}
