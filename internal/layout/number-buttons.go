package layout

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

var numberButtonValues = [3][3]int{
	{
		1, 2, 3,
	},
	{
		4, 5, 6,
	},
	{
		5, 7, 8,
	},
}

var buttonRectRelativePos = image.Rect(50, 0, 150, 50)

func (gameLayout *GameLayout) drawClearButton(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(BoardWidth+buttonRectRelativePos.Min.X),
		float32((ScreenHeight-5*cellSize)+buttonRectRelativePos.Min.Y),
		float32(buttonRectRelativePos.Dx()),
		float32(buttonRectRelativePos.Dy()),
		getIconColor(Button),
		true,
	)
	textValue := "Clear"
	textXPos := BoardWidth + 2*cellSize
	textYPos := ScreenHeight - 5*cellSize + cellSize/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

func (gameLayout *GameLayout) clearCursorValue() {
	board := gameLayout.gameInstance.Board
	cursorRow := board.CursorRow
	cursorCol := board.CursorCol
	targetCell := board.Cells[cursorRow][cursorCol]
	// 當為無法清除的值時
	if targetCell.Type == game.Preset {
		return
	}
	handleClearInput(board, targetCell, cursorRow, cursorCol)
}

func (gameLayout *GameLayout) detectClearHandler() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		xPos -= BoardWidth
		yPos -= (ScreenHeight - 5*cellSize)
		// detect range
		if xPos >= buttonRectRelativePos.Min.X && xPos <= buttonRectRelativePos.Dx() &&
			yPos >= 0 && yPos <= buttonRectRelativePos.Dy() {
			gameLayout.clearCursorValue()
		}
	}
}
