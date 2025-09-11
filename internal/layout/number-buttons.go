package layout

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		7, 8, 9,
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
		xPos -= BoardWidth + buttonRectRelativePos.Min.X
		yPos -= (ScreenHeight - 5*cellSize)
		// detect range
		if xPos >= 0 && xPos <= buttonRectRelativePos.Dx() &&
			yPos >= 0 && yPos <= buttonRectRelativePos.Dy() {
			gameLayout.clearCursorValue()
		}
	}
}

func (gameLayout *GameLayout) drawNumberButtons(screen *ebiten.Image) {
	// draw each button
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			gameLayout.drawNumberButton(screen, row, col)
		}
	}
}

func (gameLayout *GameLayout) drawNumberButton(screen *ebiten.Image, row, col int) {
	buttonXPos := BoardWidth + col*cellSize + cellSize
	buttonYPos := (ScreenHeight - 3*cellSize) + row*cellSize
	vector.DrawFilledCircle(screen,
		float32(buttonXPos),
		float32(buttonYPos+cellSize/2),
		25,
		getIconColor(Button),
		true,
	)
	textValue := fmt.Sprintf("%d", numberButtonValues[row][col])
	textXPos := buttonXPos
	textYPos := buttonYPos + cellSize/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(getTileColor(game.Input))
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

func (gameLayout *GameLayout) detectNumberButtonHandler() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		xPos, yPos := ebiten.CursorPosition()
		xPos -= (BoardWidth + cellSize/2)
		yPos -= (ScreenHeight - 3*cellSize)
		// detect range
		if xPos >= 0 && xPos <= 3*cellSize &&
			yPos >= 0 && yPos <= 3*cellSize {
			row := yPos / cellSize
			col := xPos / cellSize
			gameLayout.handleNumberButtonClick(row, col)
		}
		return
	}
}

func (gameLayout *GameLayout) handleNumberButtonClick(row, col int) {
	board := gameLayout.gameInstance.Board
	cursorRow := board.CursorRow
	cursorCol := board.CursorCol
	targetCell := board.Cells[cursorRow][cursorCol]
	cellType := targetCell.Type
	// 當為無法清除的值時
	if cellType == game.Preset {
		return
	}
	// 當輸入格為空格時
	if cellType == game.Empty {
		board.IncreaseFilledCount()
	}
	value := numberButtonValues[row][col]
	safed := board.IsSafe(cursorRow, cursorCol, value)
	if !safed {
		handleConflict(board, cellType, cursorRow, cursorCol)
	} else {
		handleNonConflict(board, cellType, cursorRow, cursorCol)
	}
	// 更新輸入
	board.Cells[cursorRow][cursorCol].Value = value
}
