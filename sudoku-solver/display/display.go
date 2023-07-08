package display

import "fmt"

// display the grid
func Display(grid [][]int) {
	// nested loop to print each element of the grid
	for _, arr := range grid {
		for i, el := range arr {
			fmt.Print(el)
			if i < len(grid)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
