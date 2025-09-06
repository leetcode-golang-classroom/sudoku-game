package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePuzzleFromSolution(t *testing.T) {
	tests := []struct {
		name           string
		targetClues    int
		wantCluesCount int
		setup          func() *Board
	}{
		{
			name:           "generate puzzle with Easy Level",
			targetClues:    int(Easy),
			wantCluesCount: int(Easy),
			setup: func() *Board {
				board := NewBoard()
				board.GenerateSolution()
				return board
			},
		},
		{
			name:           "generate puzzle with Medium Level",
			targetClues:    int(Medium),
			wantCluesCount: int(Medium),
			setup: func() *Board {
				board := NewBoard()
				board.GenerateSolution()
				return board
			},
		},
		{
			name:           "generate puzzle with Hard Level",
			targetClues:    int(Hard),
			wantCluesCount: int(Hard),
			setup: func() *Board {
				board := NewBoard()
				board.GenerateSolution()
				return board
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := tt.setup()
			board.MakePuzzleFromSolution(tt.targetClues)
			got := board.presetedCount()
			assert.Equal(t, tt.wantCluesCount, got)
		})
	}
}
