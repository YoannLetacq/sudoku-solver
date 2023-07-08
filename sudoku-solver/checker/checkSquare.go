package checker

// check double in the 3x3 square
func CheckSquare(x, y, value int, grid [][]int) bool {
	// calculate the starting coordinates of the square
	sx := int(x/3) * 3
	sy := int(y/3) * 3

	// nested loop to get the offset of the x,y coordinate
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			// if the value is equal to the grid case value
			if grid[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}
