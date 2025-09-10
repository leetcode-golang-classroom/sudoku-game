package layout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

func (gameLayout *GameLayout) drawHighLightCell(screen *ebiten.Image, row, col int) {
	board := gameLayout.gameInstance.Board
	cursorRow := board.CursorRow
	cursorCol := board.CursorCol

	// 跳過自己
	if row == cursorRow && col == cursorCol {
		return
	}
	// highLight relative row, col
	if row == cursorRow || col == cursorCol {
		gameLayout.drawHighLightCover(screen, row, col)
		return
	}
	boxSize := game.BoxSize
	boxRow := (cursorRow / boxSize) * boxSize
	boxCol := (cursorCol / boxSize) * boxSize
	// highLight Box value check
	for r := 0; r < boxSize; r++ {
		for c := 0; c < boxSize; c++ {
			br := boxRow + r
			bc := boxCol + c
			if br == row && bc == col {
				gameLayout.drawHighLightCover(screen, row, col)
				break
			}
		}
	}
}

func (gameLayout *GameLayout) drawHighLightCover(screen *ebiten.Image, row, col int) {
	xPos := col * cellSize
	yPos := PanelHeight + row*cellSize
	// stable blue
	bgColor := color.RGBA{0x6a, 0x5a, 0xcd, 128}
	if col%3 == 0 {
		xPos += thinkLineWidth
	} else {
		xPos += leanLineWidth
	}
	if row%3 == 0 {
		yPos += thinkLineWidth
	} else {
		yPos += leanLineWidth
	}
	vector.DrawFilledRect(
		screen,
		float32(xPos),
		float32(yPos),
		cellSize-1,
		cellSize-1,
		bgColor,
		false,
	)
}
