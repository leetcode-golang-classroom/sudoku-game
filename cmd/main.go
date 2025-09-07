package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/sudoku-game/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.ScreenWidth, layout.ScreenHeight)
	ebiten.SetWindowTitle("Sudoku Board")
	gameLayout := layout.NewGameLayout()
	if err := ebiten.RunGame(gameLayout); err != nil {
		log.Fatal(err)
	}
}
