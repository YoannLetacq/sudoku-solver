package main

import (
	"fmt"
	"os"

	"sudoku/sudoku-solver/display"
	"sudoku/sudoku-solver/grid"
	"sudoku/sudoku-solver/solver"
)

func main() {
	// get args
	args := os.Args[1:]

	// create a new grid if args are valid
	if !grid.IsValid(args) {
		return
	}
	sudoku := grid.New(args)

	// display the non-solve grid
	fmt.Println("Your grid is:")
	display.Display(sudoku)

	fmt.Println()
	// display the solved grid
	fmt.Println("The solved grid is:")
	// start at 0,0 top left corner of the grid
	solved := solver.Solver(0, 0, sudoku)
	if solved {
		display.Display(sudoku)
	} else {
		fmt.Println("[ERROR] Failed to solve the grid...")
	}
}
