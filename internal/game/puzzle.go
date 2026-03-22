package game

// presetBoard - 填滿格子
func (board *Board) presetBoard() bool {
	row, col, foundEmpty := -1, -1, false
	// 找到第一個非空的格子來填
	for r := 0; r < BoardSize && !foundEmpty; r++ {
		for c := 0; c < BoardSize && !foundEmpty; c++ {
			if board.Cells[r][c].Value == 0 && board.Cells[r][c].Type != Preset {
				row, col, foundEmpty = r, c, true
			}
		}
	}

	// 當所有都填滿了回傳 true
	if !foundEmpty {
		return true
	}

	// 隨機取值出來填寫
	for _, digit := range digitsShuffled() {
		// 確認 digit 是否可以填入 row, col
		if board.IsSafe(row, col, digit) {
			// 先填入 row, col 為 digit
			board.Cells[row][col].Type = Preset
			board.Cells[row][col].Value = digit
			// 如果格子填滿則回傳 true
			if board.presetBoard() {
				return true
			}
			// 否則把 row, col 回朔
			board.Cells[row][col].Type = Empty
			board.Cells[row][col].Value = 0
		}
	}
	return false
}

func solveCount(board *Board, digits []int, limit int) int {
	row, col, found := -1, -1, false
	for r := 0; r < BoardSize && !found; r++ {
		for c := 0; c < BoardSize && !found; c++ {
			if board.Cells[r][c].Type == Empty && board.Cells[r][c].Value == 0 {
				row, col, found = r, c, true
			}
		}
	}
	if !found {
		return 1
	}
	count := 0
	for _, digit := range digits {
		if board.IsSafe(row, col, digit) {
			board.Cells[row][col].Type = Preset
			board.Cells[row][col].Value = digit
			count += solveCount(board, digits, limit-count)
			board.Cells[row][col].Type = Empty
			board.Cells[row][col].Value = 0
		}
		if count >= limit {
			break
		}
	}

	return count
}

// hasUniqueSolution - 是否具有唯一解
func (board *Board) hasUniqueSolution(digits []int) bool {
	copyBoard := board.Clone()
	count := solveCountFull(&copyBoard, digits)
	return count == 1
}

// solveCountFull - 計算解答數量（不使用早期終止）
func solveCountFull(board *Board, digits []int) int {
	row, col, found := -1, -1, false
	for r := 0; r < BoardSize && !found; r++ {
		for c := 0; c < BoardSize && !found; c++ {
			if board.Cells[r][c].Type == Empty && board.Cells[r][c].Value == 0 {
				row, col, found = r, c, true
			}
		}
	}
	if !found {
		return 1
	}
	count := 0
	for _, digit := range digits {
		if board.IsSafe(row, col, digit) {
			board.Cells[row][col].Type = Preset
			board.Cells[row][col].Value = digit
			count += solveCountFull(board, digits)
			board.Cells[row][col].Type = Empty
			board.Cells[row][col].Value = 0
		}
	}
	return count
}

// GenerateSolution - 產生解法
func (board *Board) GenerateSolution() {
	// 填入解法
	board.presetBoard()
}

// presetedCount - 計算被先填入的 count
func (board *Board) presetedCount() int {
	count := 0
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if board.Cells[row][col].Type == Preset && board.Cells[row][col].Value != 0 {
				count++
			}
		}
	}
	return count
}

// MakePuzzleFromSolution - 建立題目
func (board *Board) MakePuzzleFromSolution(targetClues int) {
	solution := board.Clone()
	puzzle := solution.Clone()
	order := coordsShuffled()
	digits := digitsShuffled()
	for _, rc := range order {
		if puzzle.presetedCount() <= targetClues {
			break
		}
		r, c := rc[0], rc[1]
		if puzzle.Cells[r][c].Type == Empty && puzzle.Cells[r][c].Value == 0 {
			continue
		}
		tmpType := puzzle.Cells[r][c].Type
		tmpValue := puzzle.Cells[r][c].Value
		puzzle.Cells[r][c].Type = Empty
		puzzle.Cells[r][c].Value = 0
		if !puzzle.hasUniqueSolution(digits) {
			puzzle.Cells[r][c].Type = tmpType
			puzzle.Cells[r][c].Value = tmpValue
		}
	}
	board.TargetSolvedCount = 81 - targetClues
	board.Digits = digits
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			board.Cells[row][col].Type = puzzle.Cells[row][col].Type
			board.Cells[row][col].Value = puzzle.Cells[row][col].Value
		}
	}
}
