package layout

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

const (
	ScreenWidth    = 450
	ScreenHeight   = 450
	cellSize       = 50
	thinkLineWidth = 3
	leanLineWidth  = 1
)

type GameLayout struct {
	gameInstance *game.Game
	difficulty   game.Difficulty
}

func (gameLayout *GameLayout) Update() error {
	return nil
}

func (gameLayout *GameLayout) Draw(screen *ebiten.Image) {
	// 畫出基本背景
	gameLayout.drawBoardBackground(screen)
	// 根據遊戲狀態來畫出盤面
	gameLayout.drawCellValuesOnBoard(screen)
	// 畫出盤面格線
	gameLayout.drawLinesOnBoard(screen)
}

// drawBoardBackground - 畫出盤面背景顏色
func (gameLayout *GameLayout) drawBoardBackground(screen *ebiten.Image) {
	boardBgColor := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	boardBackground := ebiten.NewImage(ScreenWidth, ScreenHeight)
	boardBackground.Fill(boardBgColor)
	screen.DrawImage(boardBackground, nil)
}

// drawCellValuesOnBoard - 根據遊戲狀態來畫出盤面
func (gameLayout *GameLayout) drawCellValuesOnBoard(screen *ebiten.Image) {
	board := gameLayout.gameInstance.Board
	for row := 0; row < game.BoardSize; row++ {
		for col := 0; col < game.BoardSize; col++ {
			// draw preset value
			if board.Cells[row][col].Type == game.Preset {
				gameLayout.drawCellBackground(screen, row, col, getTileBgColor(board.Cells[row][col].Type))
				gameLayout.drawCellValue(screen, row, col, board.Cells[row][col].Value,
					getTileColor(board.Cells[row][col].Type),
				)
			}
			// TODO: draw input
		}
	}
}

// drawLinesOnBoard - 畫出目前盤面所需要的格線
func (gameLayout *GameLayout) drawLinesOnBoard(screen *ebiten.Image) {
	// 畫出盤面格線
	for i := 0; i <= game.BoardSize; i++ {
		x := i * cellSize
		y := i * cellSize

		// 預設是細線
		var lineColor color.Color = color.RGBA{0x77, 0x77, 0x77, 0xFF}
		lineWidth := leanLineWidth
		if i%3 == 0 {
			lineColor = color.Black
			lineWidth = thinkLineWidth
		}
		// 畫直線
		ebitenUtilDrawLine(screen, x, 0, x, ScreenHeight, lineColor, lineWidth)
		// 畫橫線
		ebitenUtilDrawLine(screen, 0, y, ScreenWidth, y, lineColor, lineWidth)
	}
}

// ebitenUtilDrawLine - 劃線，從 x1,y1 到 x2, y2 並且寬度為 width
func ebitenUtilDrawLine(screen *ebiten.Image, x1, y1, x2, y2 int,
	lineColor color.Color, width int) {
	img := ebiten.NewImage((x2-x1)+width, (y2-y1)+width)
	img.Fill(lineColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x1), float64(y1))
	screen.DrawImage(img, op)
}

func (gameLayout *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGameLayout() *GameLayout {
	gameInstance := game.NewGame()
	gameInstance.Board.GenerateSolution()
	defaultDifficulty := game.Easy
	gameInstance.Board.MakePuzzleFromSolution(int(defaultDifficulty))
	return &GameLayout{
		gameInstance: gameInstance,
		difficulty:   defaultDifficulty,
	}
}

// drawCellValue - 畫出 click 之後顯示出來的值
func (*GameLayout) drawCellValue(screen *ebiten.Image, row, col, value int, numColor color.Color) {
	// 繪製數字 (置中)
	textValue := fmt.Sprintf("%d", value)
	textXPos := col*cellSize + (cellSize)/2
	textYPos := row*cellSize + (cellSize)/2
	if col%3 == 0 {
		textXPos += thinkLineWidth
	} else {
		textXPos += leanLineWidth
	}
	if row%3 == 0 {
		textYPos += thinkLineWidth
	} else {
		textYPos += leanLineWidth
	}
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(numColor)
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

// drawCellBackground - 繪圖出目前盤面的情況
func (*GameLayout) drawCellBackground(screen *ebiten.Image, row, col int, bgColor color.Color) {
	xPos := col * cellSize
	yPos := row * cellSize
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
