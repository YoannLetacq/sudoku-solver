package grid

import "fmt"

// validate the grid format
func IsFormatted(grid [][]int) (bool, int, int) {
	validRows := isValidRows(grid)
	if !validRows {
		return false, -1, -1
	}

	row, col := isValidColumns(grid)
	if row != -1 {
		return false, row, col
	}

	row, col = isValidSubgrids(grid)
	if row != -1 {
		return false, row, col
	}

	return true, -1, -1
}

func isValidRows(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			num := grid[i][j]
			if num != 0 {
				if seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}
	return true
}

func isValidColumns(grid [][]int) (int, int) {
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			num := grid[i][j]
			if num != 0 {
				if seen[num] {
					return i, j
				}
				seen[num] = true
			}
		}
	}
	return -1, -1
}

func isValidSubgrids(grid [][]int) (int, int) {
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			seen := make(map[int]bool)
			for i := r; i < r+3; i++ {
				for j := c; j < c+3; j++ {
					num := grid[i][j]
					if num != 0 {
						if seen[num] {
							return i, j
						}
						seen[num] = true
					}
				}
			}
		}
	}
	return -1, -1
}

// validate the argument format
func IsValid(args []string) bool {
	fmt.Println(len(args))
	// verify the number of row
	if len(args) > 10 {
		fmt.Printf("[ERROR] Too much argument in call\nhave (%d)\nwant (9)", len(args))
		return false
	}
	if len(args) > 10 {
		fmt.Printf("[ERROR] Missing argument in call\nhave (%d)\nwant (9)", len(args))
		return false
	}

	// verify the number of element per row
	lines := args[1:]
	for _, el := range lines {
		if len(el) > 10 {
			fmt.Printf("[ERROR] Too much argument in call\nhave (%d)\nwant (9)", len(lines))
			return false
		}
		if len(el) < 10 {
			fmt.Printf("[ERROR] Missing argument in call\nhave (%d)\nwant (9)", len(lines))
			return false
		}
	}

	// define valid grid based on composition
	grid := New(args)
	valid, row, col := IsFormatted(grid)
	if valid {
		fmt.Println("The Sudoku grid is valid.")
		return true
	} else {
		fmt.Printf("The Sudoku grid is invalid. Duplicate found at position (%d, %d).\n", row, col)
		return false
	}
}
