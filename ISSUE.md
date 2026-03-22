# Issue Report: Sudoku Puzzle Generation Bugs

## Summary

During code review and testing of the Sudoku puzzle generation logic, several critical bugs were discovered that caused incorrect puzzle generation. All bugs have been fixed.

---

## Bugs Found

### 1. Shallow Copy in `Clone()` (board.go)

**Severity:** Critical  
**File:** `internal/game/board.go`

**Problem:**
```go
func (board *Board) Clone() Board {
    copyBoard := *board
    return copyBoard
}
```

The `Board.Cells` field is of type `[BoardSize][BoardSize]*Cell` (array of pointers). A shallow copy like `*board` copies the array structure but all `*Cell` pointers still point to the original cells. Modifying a cloned board's cells would also modify the original board.

**Fix:** Deep copy each Cell:
```go
func (board *Board) Clone() Board {
    copyBoard := *board
    for row := 0; row < BoardSize; row++ {
        for col := 0; col < BoardSize; col++ {
            cell := *board.Cells[row][col]
            copyBoard.Cells[row][col] = &cell
        }
    }
    return copyBoard
}
```

---

### 2. Invalid Pointer Assignment in `MakePuzzleFromSolution` (puzzle.go)

**Severity:** Critical  
**File:** `internal/game/puzzle.go`

**Problem:**
```go
board = &puzzle
```

This statement only changes the local variable `board` inside the function. It has no effect on the caller. Due to the shallow copy bug (Bug #1), the code accidentally worked, but the logic was fundamentally broken.

**Fix:** Explicitly copy cells from puzzle to board:
```go
for row := 0; row < BoardSize; row++ {
    for col := 0; col < BoardSize; col++ {
        board.Cells[row][col].Type = puzzle.Cells[row][col].Type
        board.Cells[row][col].Value = puzzle.Cells[row][col].Value
    }
}
```

---

### 3. Pointer Aliasing Bug in `MakePuzzleFromSolution` (puzzle.go)

**Severity:** Critical  
**File:** `internal/game/puzzle.go`

**Problem:**
```go
tmp := puzzle.Cells[r][c]  // tmp is a pointer to the Cell
puzzle.Cells[r][c].Type = Empty
puzzle.Cells[r][c].Value = 0
// ... check uniqueness ...
if !puzzle.hasUniqueSolution(digits) {
    puzzle.Cells[r][c].Type = tmp.Type  // BUG: tmp.Type is already Empty!
    puzzle.Cells[r][c].Value = tmp.Value
}
```

Since `tmp` is a pointer to the same Cell, when we modify `puzzle.Cells[r][c].Type = Empty`, the value `tmp.Type` also becomes `Empty`. The restore logic always failed, causing:
- Non-unique puzzles to be generated
- Incorrect cell restoration during puzzle generation

**Fix:** Copy values instead of pointers:
```go
tmpType := puzzle.Cells[r][c].Type
tmpValue := puzzle.Cells[r][c].Value
puzzle.Cells[r][c].Type = Empty
puzzle.Cells[r][c].Value = 0
// ... check uniqueness ...
if !puzzle.hasUniqueSolution(digits) {
    puzzle.Cells[r][c].Type = tmpType
    puzzle.Cells[r][c].Value = tmpValue
}
```

---

### 4. Early Termination in `solveCount` (puzzle.go)

**Severity:** Medium  
**File:** `internal/game/puzzle.go`

**Problem:**
```go
if count >= limit {
    break
}
```

The early termination optimization assumed that stopping when `count >= limit` was safe. However, due to random digit ordering in the backtracking search, this optimization could produce inconsistent results:
- Different digit orders might find solutions in different orders
- Some branches might not be explored completely
- This led to non-deterministic solution counting

**Fix:** Added a separate function `solveCountFull` for complete search when checking uniqueness:
```go
func solveCountFull(board *Board, digits []int) int {
    // ... no early termination ...
}
```

---

### 5. Non-deterministic Solution Counting (puzzle.go)

**Severity:** Medium  
**File:** `internal/game/puzzle.go`

**Problem:**
`digitsShuffled()` was called at every recursion level, producing different search orders each time. This caused:
- `hasUniqueSolution` to return inconsistent results
- Same puzzle to sometimes pass and sometimes fail uniqueness checks

**Fix:**
1. Added a `Digits` field to `Board` to store the search order
2. Passed fixed digit slice through the recursion
3. Stored `Digits` in `MakePuzzleFromSolution` for verification

---

## Changes Made

### Modified Files

| File | Changes |
|------|---------|
| `internal/game/sudoku.go` | Added `Digits []int` field to `Board` struct |
| `internal/game/board.go` | Fixed `Clone()` to use deep copy |
| `internal/game/puzzle.go` | Fixed pointer aliasing, added `solveCountFull()`, added `Digits` field support |
| `internal/game/board_test.go` | Added `TestClone`, `TestHasUniqueSolution`, `TestMakePuzzlePreservesCorrectClueCount` tests |

### New Tests Added

- `TestClone` - Verifies deep copy behavior
- `TestHasUniqueSolution` - Verifies generated solutions have unique solutions
- `TestMakePuzzlePreservesCorrectClueCount` - Verifies clue count matches target

---

## Verification

All puzzles generated now have guaranteed unique solutions. Tested across:
- Easy difficulty (38 clues)
- Medium difficulty (32 clues)  
- Hard difficulty (28 clues)

Run tests with:
```bash
go test -v ./internal/game/...
```
