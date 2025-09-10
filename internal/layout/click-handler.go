package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

func (gameLayout *GameLayout) detectClickCell() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		yPos -= PanelHeight
		// range check
		if (xPos >= 0 && xPos <= BoardWidth) &&
			(yPos >= 0 && yPos <= ScreenHeight) {
			row := yPos / cellSize
			col := xPos / cellSize
			gameLayout.updateCursor(row, col)
		}
	}
}

func (gameLayout *GameLayout) updateCursor(row, col int) {
	if row >= 0 && row < game.BoardSize {
		gameLayout.gameInstance.Board.CursorRow = row
	}
	if col >= 0 && col < game.BoardSize {
		gameLayout.gameInstance.Board.CursorCol = col
	}
}
