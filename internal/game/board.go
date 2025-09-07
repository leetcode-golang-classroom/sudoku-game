package game

import "fmt"

// 檢查 row, col 位置是否可以放置 num
func (board *Board) IsSafe(row, col, num int) bool {
	// 檢查行與列是否有放相同的值
	for i := 0; i < BoardSize; i++ {
		if board.Cells[row][i].Value == num ||
			board.Cells[i][col].Value == num {
			return false
		}
	}

	// 檢查 Box 內是否有相同的值
	boxRow := (row / BoxSize) * BoxSize
	boxCol := (col / BoxSize) * BoxSize
	for br := 0; br < BoxSize; br++ {
		for bc := 0; bc < BoxSize; bc++ {
			if board.Cells[boxRow+br][boxCol+bc].Value == num {
				return false
			}
		}
	}

	return true
}

func (board *Board) Clone() Board {
	copyBoard := *board
	return copyBoard
}

func (board *Board) String() string {
	out := ""
	for r := 0; r < BoardSize; r++ {
		if r%3 == 0 {
			out += "+-------+-------+-------+\n"
		}
		for c := 0; c < BoardSize; c++ {
			if c%3 == 0 {
				out += "| "
			}
			if board.Cells[r][c].Value == 0 {
				out += ". "
			}
			if board.Cells[r][c].Type == Preset {
				out += fmt.Sprintf("%d ", board.Cells[r][c].Value)
			}
		}
		out += "|\n"
	}
	out += "+-------+-------+-------+\n"
	return out
}
