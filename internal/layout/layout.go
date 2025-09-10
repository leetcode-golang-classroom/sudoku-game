package layout

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

const (
	PanelHeight     = 100 // 上方面板高度
	BoardWidth      = 450
	InputPanelWidth = 200
	ScreenWidth     = BoardWidth + InputPanelWidth
	ScreenHeight    = PanelHeight + 450
	cellSize        = 50
	thinkLineWidth  = 3
	leanLineWidth   = 1
)

type GameLayout struct {
	gameInstance    *game.Game
	difficultyLevel int
	isPlayerWin     bool
	elapsedSeconds  int
}

func (gameLayout *GameLayout) Update() error {
	gameLayout.handleRestartButton()
	gameLayout.handleToggleLevelDifficultButton()
	if gameLayout.isPlayerWin {
		return nil
	}
	gameLayout.detectClearHandler()
	gameLayout.detectClickCell()
	gameLayout.elapsedSeconds = gameLayout.gameInstance.GetElaspedTime()
	gameLayout.DetectCursor()
	gameLayout.DetectInput()
	// 檢查狀態
	gameLayout.isPlayerWin = gameLayout.checkIfPlayerWin()
	return nil
}

func (gameLayout *GameLayout) Draw(screen *ebiten.Image) {
	// 畫出基本背景
	gameLayout.drawBoardBackground(screen)
	// 畫出目前狀態面板
	gameLayout.drawRemainingUnsolvedCount(screen)
	gameLayout.drawBugCount(screen)
	gameLayout.drawBoardStatus(screen)
	gameLayout.drawRestartButton(screen)
	gameLayout.drawTimeLayout(screen)
	gameLayout.drawLevelButtonWithIcon(screen)
	gameLayout.drawClearButton(screen)
	// 畫出 cursor
	gameLayout.drawCursor(screen)
	// 根據遊戲狀態來畫出盤面
	gameLayout.drawCellValuesOnBoard(screen)
	// 畫出盤面格線
	gameLayout.drawLinesOnBoard(screen)
}

// drawCursor - 繪製游標
func (gameLayout *GameLayout) drawCursor(screen *ebiten.Image) {
	cursorBgColor := color.RGBA{0xff, 0, 0, 128}
	targetRow := gameLayout.gameInstance.Board.CursorRow
	targetCol := gameLayout.gameInstance.Board.CursorCol
	if gameLayout.gameInstance.Board.Cells[targetRow][targetCol].Type == game.Preset {
		cursorBgColor = color.RGBA{0xff, 0xff, 0, 128}
	}
	if gameLayout.gameInstance.Board.Cells[targetRow][targetCol].Type == game.Input {
		cursorBgColor = color.RGBA{0x00, 0xff, 0x00, 128}
	}
	if gameLayout.gameInstance.Board.Cells[targetRow][targetCol].Type == game.InputConflict {
		cursorBgColor = color.RGBA{0x00, 0xff, 0xff, 128}
	}
	gameLayout.drawCellBackground(screen,
		gameLayout.gameInstance.Board.CursorRow,
		gameLayout.gameInstance.Board.CursorCol,
		cursorBgColor,
	)
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
			if board.Cells[row][col].Type == game.Preset ||
				board.Cells[row][col].Type == game.Input ||
				board.Cells[row][col].Type == game.InputConflict {
				if row != gameLayout.gameInstance.Board.CursorRow ||
					col != gameLayout.gameInstance.Board.CursorCol {
					gameLayout.drawCellBackground(screen, row, col, getTileBgColor(board.Cells[row][col].Type))
				}
				gameLayout.drawCellValue(screen, row, col, board.Cells[row][col].Value,
					getTileColor(board.Cells[row][col].Type),
				)
			}
			// highlight background
			gameLayout.drawHighLightCell(screen, row, col)
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
		ebitenUtilDrawLine(screen, x, PanelHeight+0, x, ScreenHeight, lineColor, lineWidth)
		// 畫橫線
		ebitenUtilDrawLine(screen, 0, PanelHeight+y, BoardWidth, PanelHeight+y, lineColor, lineWidth)
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
	defaultDifficulty := difficultyOptions[0]
	gameInstance.Board.MakePuzzleFromSolution(int(defaultDifficulty))
	gameInstance.StartTime = time.Now().UTC()
	return &GameLayout{
		gameInstance:    gameInstance,
		difficultyLevel: 0,
		isPlayerWin:     false,
	}
}

// drawCellValue - 畫出 click 之後顯示出來的值
func (*GameLayout) drawCellValue(screen *ebiten.Image, row, col, value int, numColor color.Color) {
	// 繪製數字 (置中)
	textValue := fmt.Sprintf("%d", value)
	textXPos := col*cellSize + (cellSize)/2
	textYPos := PanelHeight + row*cellSize + (cellSize)/2
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
	yPos := PanelHeight + row*cellSize
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

func (gameLayout *GameLayout) drawRemainingUnsolvedCount(screen *ebiten.Image) {
	board := gameLayout.gameInstance.Board
	emojiValue := "⬜"
	emojiXPos := len(emojiValue)
	emojiYPos := cellSize / 2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getIconColor(RemainingCount))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
	value := board.TargetSolvedCount - board.FilledCount
	textValue := fmt.Sprintf("%03d", value)
	textXPos := cellSize
	textYPos := cellSize / 2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignStart
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

func (gameLayout *GameLayout) drawBugCount(screen *ebiten.Image) {
	board := gameLayout.gameInstance.Board
	emojiValue := "🐛"
	emojiXPos := cellSize*3 + len(emojiValue)
	emojiYPos := cellSize / 2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getIconColor(Bug))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
	value := board.ConflictCount
	textValue := fmt.Sprintf("%03d", value)
	textXPos := 4 * cellSize
	textYPos := cellSize / 2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignStart
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

// checkIfPlayerWin - 檢查勝利狀態
func (gameLayout *GameLayout) checkIfPlayerWin() bool {
	board := gameLayout.gameInstance.Board
	remainingCount := board.TargetSolvedCount - board.FilledCount
	conflictCount := board.ConflictCount
	return remainingCount == 0 && conflictCount == 0
}

// drawBoardStatus - 根據是否勝利來畫出不同的提示詞
func (gameLayout *GameLayout) drawBoardStatus(screen *ebiten.Image) {
	emojiValue := "⏳"
	message := "Keep going"
	iconColor := getIconColor(Playing)
	if gameLayout.isPlayerWin {
		emojiValue = "🏆"
		message = "You Win！"
		iconColor = getIconColor(Win)
	}
	emojiXPos := len(emojiValue)
	emojiYPos := cellSize + cellSize/2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(iconColor)
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
	textValue := message
	textXPos := cellSize
	textYPos := cellSize + cellSize/2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignStart
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   25,
	}, textOpts)
}

// drawRestartButton - 繪製重新開始的 Button
func (gameLayout *GameLayout) drawRestartButton(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(ScreenWidth-cellSize/2), cellSize+cellSize/2, 25,
		getIconColor(Button),
		true,
	)
	emojiValue := "🔃"
	emojiXPos := ScreenWidth - cellSize + len(emojiValue)
	emojiYPos := cellSize + cellSize/2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getIconColor(Restart))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
}

func (gameLayout *GameLayout) ResetGameWithLevel() {
	gameInstance := game.NewGame()
	gameInstance.Board.GenerateSolution()
	defaultDifficulty := difficultyOptions[gameLayout.difficultyLevel]
	gameInstance.Board.MakePuzzleFromSolution(int(defaultDifficulty))
	gameInstance.StartTime = time.Now().UTC()
	gameLayout.gameInstance = gameInstance
}
