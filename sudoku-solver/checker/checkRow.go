package checker

// check double in all grid row
func CheckRow(y, value int, grid [][]int) bool {
	// iterate over all numbers of a single row
	for i := range [9]int{} {
		// if the value is equal to the grid case value
		if grid[y][i] == value {
			return true
		}
	}
	return false
}
