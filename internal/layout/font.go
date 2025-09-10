package layout

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/fonts"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	emojiFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	s, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.NotoEmojiRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	emojiFaceSource = s
}

func getTileColor(cellType game.CellType) color.Color {
	switch cellType {
	case game.Empty:
		return color.RGBA{0x77, 0x6e, 0x65, 0xff}
	case game.Preset:
		return color.Black
	case game.InputConflict:
		return color.RGBA{0xff, 0, 0, 0xff}
	case game.Input:
		return color.RGBA{0, 0, 0xff, 0xff}
	default:
		return color.RGBA{0xf9, 0xf6, 0xf2, 0xff}
	}
}

// getTileBgColor - Tile背景顏色
func getTileBgColor(cellType game.CellType) color.Color {
	switch cellType {
	case game.Preset: // 顯示亮灰色
		return color.RGBA{0xCD, 0xC9, 0xC9, 0xFF}
	case game.InputConflict: // 顯示粉紅
		return color.RGBA{0xDC, 0x69, 0xB4, 0xFF}
	default: // 其他都是預設 白色
		return color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	}
}

type IconType int

const (
	RemainingCount IconType = iota
	Playing
	Win
	Bug
	Restart
	Button
	IsClock
	DarkButton
)

func getIconColor(iconType IconType) color.Color {
	switch iconType {
	case DarkButton: // 顯示暗藍色
		return color.RGBA{0x70, 0x80, 0x90, 0xff}
	case RemainingCount, Button: // 顯示亮灰色
		return color.RGBA{0xCD, 0xC9, 0xC9, 0xFF}
	case Playing, Bug, Restart: // Green
		return color.RGBA{0x22, 0x8B, 0x22, 0xFF}
	case Win, IsClock: // Gold
		return color.RGBA{0xFF, 0xD7, 0x00, 0xFF}
	default: // 其他都是預設 白色
		return color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	}
}
