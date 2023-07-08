package checker

// call all check func tell double presence
func Checker(x, y, value int, grid [][]int) bool {
	return !CheckRow(y, value, grid) && !CheckColumn(x, value, grid) && !CheckSquare(x, y, value, grid)
}
