package grid

// create the grid following the args
func New(args []string) [][]int {
	grid := [][]int{}

	for _, arg := range args {
		// skip if the argument is empty
		if arg == "" {
			continue
		}

		arr := make([]int, 9) // Make an array of size 9 filled with zeros

		for i, el := range arg {
			number := int(el - '0')
			// fill the grid with 0 when get a '.'
			if el == '.' {
				number = 0
			}
			arr[i] = number
		}

		grid = append(grid, arr)
	}

	return grid
}
