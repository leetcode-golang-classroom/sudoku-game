package layout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/game"
)

var difficultyOptions = [3]game.Difficulty{
	game.Easy,
	game.Medium,
	game.Hard,
}
var difficultyIcons = map[game.Difficulty]string{
	game.Easy:   "🐣", // 小雞
	game.Medium: "🦊", // 狐狸
	game.Hard:   "🦁", // 獅子
}

func getLevelIconColor(difficulty game.Difficulty) color.Color {
	switch difficulty {
	case game.Medium: // dark Red
		return color.RGBA{0xdc, 0x14, 0x3c, 0xff}
	case game.Hard, game.Easy: // Gold
		return color.RGBA{0xFF, 0xD7, 0x00, 0xFF}
	default: // 其他都是預設 白色
		return color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	}
}

func (gameLayout *GameLayout) drawLevelButtonWithIcon(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(ScreenWidth/2), cellSize+cellSize/2, 25,
		getIconColor(DarkButton),
		true,
	)
	gameDifficulty := difficultyOptions[gameLayout.difficultyLevel]
	levelIcon := difficultyIcons[gameDifficulty]
	emojiValue := levelIcon
	emojiXPos := ScreenWidth / 2
	emojiYPos := cellSize + cellSize/2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getLevelIconColor(gameDifficulty))
	emojiOpts.PrimaryAlign = text.AlignCenter
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
}
