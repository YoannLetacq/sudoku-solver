package solver

import "sudoku/sudoku-solver/checker"

// solve the given grid
func Solver(x, y int, grid [][]int) bool {
	// if we are at the end of the grid
	if y == 9 {
		return true
	}

	if grid[y][x] != 0 {
		// skip the fill case
		return Solver(Next(x, y, grid))
	} else {
		// solve the current case
		for i := range [9]int{} {
			v := i + 1
			// look for doubles
			if checker.Checker(x, y, v, grid) {
				// no double place value
				grid[y][x] = v

				// solve the next case while it's true
				if Solver(Next(x, y, grid)) {
					return true
				}
				// if the value doesn't match set back 0
				grid[y][x] = 0

			}
		}
	}
	return false
}
