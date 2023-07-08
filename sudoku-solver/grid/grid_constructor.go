package grid

// create the grid following the args
func New(args []string) [][]int {
	arr := []int{}
	grid := [][]int{}
	var number int
	//
	for _, arg := range args {
		for _, el := range arg {
			number = int(el - '0')
			// fill the grid with 0 when get a '.'
			if number == '.'-'0' {
				number = 0
				arr = append(arr, number)
			} else {
				arr = append(arr, number)
			}
		}
		grid = append(grid, arr)
		arr = nil
	}
	return grid
}
