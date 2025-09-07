package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	tests := []struct {
		name      string
		wantBoard *Board
	}{
		{
			name: "Empty Board",
			wantBoard: &Board{
				Cells: [9][9]*Cell{
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
					{
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
						{0, Empty}, {0, Empty}, {0, Empty},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			assert.Equal(t, tt.wantBoard.Cells, game.Board.Cells)
		})
	}
}
