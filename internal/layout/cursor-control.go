package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// DetectCursor - 游標移動
func (gameLayout *GameLayout) DetectCursor() {
	// 游標移動偵測
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		gameLayout.gameInstance.Board.DecreaseCursorRow()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		gameLayout.gameInstance.Board.IncreaseCursorRow()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		gameLayout.gameInstance.Board.DecreaseCursorCol()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		gameLayout.gameInstance.Board.IncreaseCursorCol()
		return
	}
}
