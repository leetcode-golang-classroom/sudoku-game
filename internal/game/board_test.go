package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	board := NewBoard()
	board.Cells[0][0].Value = 5
	board.Cells[0][0].Type = Preset
	board.Cells[1][1].Value = 3
	board.Cells[1][1].Type = Input

	cloned := board.Clone()

	assert.Equal(t, board.Cells[0][0].Value, cloned.Cells[0][0].Value)
	assert.Equal(t, board.Cells[0][0].Type, cloned.Cells[0][0].Type)
	assert.Equal(t, board.Cells[1][1].Value, cloned.Cells[1][1].Value)
	assert.Equal(t, board.Cells[1][1].Type, cloned.Cells[1][1].Type)

	assert.NotSame(t, board.Cells[0][0], cloned.Cells[0][0])
	assert.NotSame(t, board.Cells[1][1], cloned.Cells[1][1])

	cloned.Cells[0][0].Value = 9
	cloned.Cells[0][0].Type = Input

	assert.NotEqual(t, board.Cells[0][0].Value, cloned.Cells[0][0].Value)
	assert.NotEqual(t, board.Cells[0][0].Type, cloned.Cells[0][0].Type)
}

func TestHasUniqueSolution(t *testing.T) {
	board := NewBoard()
	board.GenerateSolution()
	got := board.hasUniqueSolution(digitsShuffled())
	assert.True(t, got, "Generated solution should have unique solution")
}

func TestMakePuzzlePreservesCorrectClueCount(t *testing.T) {
	tests := []struct {
		name        string
		targetClues int
	}{
		{name: "Easy", targetClues: int(Easy)},
		{name: "Medium", targetClues: int(Medium)},
		{name: "Hard", targetClues: int(Hard)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()
			board.GenerateSolution()
			board.MakePuzzleFromSolution(tt.targetClues)

			count := board.presetedCount()
			assert.GreaterOrEqual(t, count, tt.targetClues,
				"Puzzle should have at least target clues")
			assert.LessOrEqual(t, count, tt.targetClues+5,
				"Puzzle clue count should be close to target")
		})
	}
}

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
