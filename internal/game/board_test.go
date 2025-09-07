package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSafe(t *testing.T) {
	type coord struct {
		Row int
		Col int
	}
	tests := []struct {
		name        string
		targetCoord coord
		targetValue int
		setup       func() *Board
		want        bool
	}{
		{
			name: "check if board is safe to put 1, 2 with value 9, should be false, for setup board.Cells[1][1].Value=9",
			targetCoord: coord{
				Row: 1,
				Col: 2,
			},
			targetValue: 9,
			setup: func() *Board {
				board := NewBoard()
				board.Cells[1][1].Type = Preset
				board.Cells[1][1].Value = 9
				return board
			},
			want: false,
		},
		{
			name: "check if board is safe to put 2, 5 with value 9, should be true, for setup board.Cells[1][1].Value=9, boards.Cells[4][3].Value = 9",
			targetCoord: coord{
				Row: 2,
				Col: 5,
			},
			targetValue: 9,
			setup: func() *Board {
				board := NewBoard()
				board.Cells[1][1].Type = Preset
				board.Cells[1][1].Value = 9
				board.Cells[4][3].Type = Preset
				board.Cells[4][3].Value = 9
				return board
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := tt.setup()
			got := board.IsSafe(tt.targetCoord.Row, tt.targetCoord.Col, tt.targetValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
