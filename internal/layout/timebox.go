package layout

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (gameLayout *GameLayout) drawTimeLayout(screen *ebiten.Image) {
	elapsedSeconds := gameLayout.elapsedSeconds
	secondPart := elapsedSeconds % 60
	minutePart := (elapsedSeconds / 60) % 60
	textValue := fmt.Sprintf("%02d:%02d", minutePart, secondPart)
	textXPos := ScreenWidth - cellSize/4 + len(textValue)
	textYPos := cellSize / 2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignEnd
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, textValue, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
	emojiValue := "⏰"
	emojiXPos := ScreenWidth - 3*cellSize + len(emojiValue)
	emojiYPos := cellSize / 2
	emojiOpts := &text.DrawOptions{}
	emojiOpts.ColorScale.ScaleWithColor(getIconColor(IsClock))
	emojiOpts.PrimaryAlign = text.AlignStart
	emojiOpts.SecondaryAlign = text.AlignCenter
	emojiOpts.GeoM.Translate(float64(emojiXPos), float64(emojiYPos))
	text.Draw(screen, emojiValue, &text.GoTextFace{
		Source: emojiFaceSource,
		Size:   30,
	}, emojiOpts)
}
