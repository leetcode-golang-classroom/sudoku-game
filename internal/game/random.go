package game

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// digitsShuffled - 取出亂序數字
func digitsShuffled() []int {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(digits), func(i, j int) { digits[i], digits[j] = digits[j], digits[i] })
	return digits
}

// coords - 取出座標資料
func coords() [][2]int {
	var coordPairs [][2]int
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			coordPairs = append(coordPairs, [2]int{row, col})
		}
	}
	return coordPairs
}

// coordsShuffled - 取出亂數座標
func coordsShuffled() [][2]int {
	coordPairs := coords()
	rand.Shuffle(len(coordPairs), func(i, j int) { coordPairs[i], coordPairs[j] = coordPairs[j], coordPairs[i] })
	return coordPairs
}
