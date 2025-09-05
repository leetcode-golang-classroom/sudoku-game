package game

// CellType 代表數獨格子的狀態
type CellType int

const (
	Empty  CellType = iota // 空格
	Preset                 // 題目預設的數字
	Input                  // 玩家輸入的數字
)

// Cell 代表數獨的一個格子
type Cell struct {
	Value int      // 數字，0 表示空格
	Type  CellType // 狀態：Empty、Preset、Input
}
