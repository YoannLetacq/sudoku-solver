package solver
// return the next case 
func Next(x, y int, grid [][]int) (int, int, [][]int) {
	// set the next coordinate
	nextX := (x + 1) % 9
	nextY := y
	// if we are at the end of the line change thr row
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY, grid
}
