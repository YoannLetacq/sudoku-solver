# Sudoku Solver

This project presents a Sudoku solver. This is a Go application that accepts a Sudoku puzzle as input and outputs the solution, assuming the puzzle has a unique solution.

## Installation

Before you can run this program, you need to ensure that you have Go installed on your machine. If you do not, you can download and install it from the official Go website. (https://go.dev/dl/)

## Features

Our Sudoku solver application offers a range of features designed to provide an efficient and user-friendly experience for solving Sudoku puzzles:

1. **Sudoku Solver:** The core feature of this application is to solve a given valid Sudoku puzzle. Using a depth-first search strategy and backtracking, the solver guarantees a solution for all valid puzzles.

2. **Input Validation:** This application checks the input Sudoku puzzle for validity. It ensures that the grid is a 9x9 matrix and the given clues are within a range of 1-9. If the input does not meet these criteria, it returns an error message.

3. **User-Friendly Interface:** The application features a clean, minimalist command-line interface, providing clear prompts and error messages to guide the user through the process of inputting a Sudoku puzzle.

4. **Error Handling:** Our application is designed to handle errors gracefully. If the provided Sudoku puzzle is invalid, the application will notify the user and provide a clear and useful error message.

5. **Performance:** Thanks to our efficient algorithm, the Sudoku solver handles even the toughest puzzles quickly and efficiently.

6. **Open Source:** Our Sudoku solver is open-source, inviting users to contribute, learn from the code, or even modify it for their own purposes.

Please feel free to explore these features and more. If you have any suggestions or feature requests, don't hesitate to open an issue in the GitHub repository!

## Usage

The program accepts Sudoku puzzle strings from the command line. A Sudoku puzzle string consists of 9 lines, each containing 9 characters. Digits represent given values, and '.' characters represent empty cells. Here is an example:

```bash 
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

The output is the solved Sudoku puzzle:

```bash 
Your grid is:
. 9 6 . 4 . . . 1
1 . . . 6 . . 4 .
5 . 4 8 1 . 3 9 .
. . 7 9 5 . . 4 3
. 3 . . 8 . . . .
4 . 5 . 2 3 . 1 8
. 1 . 6 3 . . 5 9
. 5 9 . 7 . 8 3 .
. . 3 5 9 . . . 7

The solved grid is:
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

In case of invalid inputs, the program will return errors:

```bash
 go run . "9.46.3..1" "37.1..2.6" "..6..93.4" "..13..9.5" "56..91..." "82...461." "..79...4." "425.167.." "1.2..75.8"
The Sudoku grid is invalid. Duplicate found at position (row: 7, col: 1).
                                                             
```

```bash 
go run . "9.46.3..1" "37.1..2.6" "..6..93.4" "..13..9.5" "56..91..." "82...461." "..79...4." "425.167.." "1.2."     
[ERROR] Missing number in call
have (4)
want (9)%                                                                               
```

```bash
 go run . "not" "a" "sudoku"                                      
[ERROR] Invalid number format, didn't found alphanumeric.
```

## Package

This project is composed of multiple Go packages:

main: This is the entry point of the application, where we parse the input arguments, initialize the Sudoku grid, display the input and output, and call the Sudoku solver.

* solver: This package provides the Solver function, which uses a backtracking algorithm to solve the Sudoku puzzle.

* grid: This package provides functions for creating and validating a Sudoku grid.

* display: This package provides a function for displaying a Sudoku grid.

* checker: This package provides functions for checking whether a value can be placed in a specific position on the grid.

Please note that the code has been written with good practices in mind: packages are decoupled, functions are small and focused, and variable names are descriptive.


## Contributing

We welcome contributions to this Sudoku solver! If you're interested in improving the project, here are some ways you can contribute:

1. **Report bugs or suggest features:** If you find a bug or have an idea for a feature that could improve the Sudoku solver, please open an issue in the GitHub project.

2. **Improve the code:** If you're able to improve the code, whether it's optimizing the algorithm, refactoring for readability, or fixing a bug, we'd love to see your pull requests.

3. **Improve documentation:** Clear and helpful documentation can make the difference between a good project and a great one. If you see any areas of the README or code comments that could be clearer or more comprehensive, please open a pull request with your improvements.

Here are the steps to contribute:

1. Fork the project on GitHub.
2. Clone your fork: **git clone https://github.com/YoannLetacq/sudoku-solver.git. **
3. Create a new branch: **git checkout -b my-awesome-feature.**
4. Make your changes and commit them: **git commit -m "Add awesome feature".**
5. Push your branch: git push -u origin my-awesome-feature.
6. Open a pull request from your branch on your fork to the **main** branch on the original repository.

Please follow the **Go Code Review Comments** for style guidelines in your contributions.

Thank you for your interest in improving the Sudoku solver!

## License 

This project is licensed under the terms of the MIT license.

## Support

If you encounter any problems or have suggestions for this project, feel free to open an issue. You can do this by clicking on the "Issues" tab at the top of this project's GitHub page, then clicking on the "New Issue" button.

If you have specific questions about the project, you can reach out directly. Contact details are provided below:

* Email: yoannletacq@gmail.com

Please remember to describe your problem/suggestion/question in as much detail as possible. The more information you provide, the easier it will be to address your issue or suggestion.

We appreciate any feedback and contributions. Thank you for your support!

## Author 

Yoann "YoannLetacq" Letacq
