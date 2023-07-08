package checker

// check double in all griw column
func CheckColumn(x, value int, grid [][]int) bool {
	// iterate over all numbers of a single column
	for i := range [9]int{} {
		// if the value is equal to the grid case
		if grid[i][x] == value {
			return true
		}
	}
	return false
}
