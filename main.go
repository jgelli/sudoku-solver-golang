package main

import (
	"errors"
	"fmt"
)

var seduleLength int = 3

var easy = [9][9]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},

	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},

	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

var medium = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 3, 0, 8, 5},
	{0, 0, 1, 0, 2, 0, 0, 0, 0},

	{0, 0, 0, 5, 0, 7, 0, 0, 0},
	{0, 0, 4, 0, 0, 0, 1, 0, 0},
	{0, 9, 0, 0, 0, 0, 0, 0, 0},

	{5, 0, 0, 0, 0, 0, 0, 7, 3},
	{0, 0, 2, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 4, 0, 0, 0, 9},
}

func main() {

	// ok, err := validatePosition(9, 5, 1)
	// position := [2]int{0, 0}
	sudokuSolved, err := solveSudoku(easy)

	if err != nil {
		fmt.Println(err)
	}

	printBoard(sudokuSolved)

	// printBoard(sudoku)

	// fmt.Println(getNextEmptyPosition(sudoku))

}

// func solveSudoku(board [9][9]int, position [2]int) [9][9]int {
// 	number := 1
// 	column := position[0]
// 	line := position[1]

// 	if column == 9 {
// 		line++
// 		column = 0
// 	}

// 	if board[line][column] != 0 {
// 		position[0] = column
// 		position[1] = line
// 		return solveSudoku(board, position)
// 	}

// 	for number <= 9 {
// 		isValid := validatePosition(board, number, column, line)

// 		if !isValid {
// 			number++
// 			continue
// 		}

// 		board[line][column] = number

// 	}

// 	if board[line][column] == 0 {
// 		return board
// 	}

// 	// if column

// 	return board
// }

func solveSudoku(board [9][9]int) ([9][9]int, error) {

	line, column, err := getNextEmptyPosition(board)

	if err != nil {
		return board, nil
	}

	number := 1
	for number <= 9 {
		isValid := validatePosition(board, number, line, column)
		if line == 0 && column == 0 {
			fmt.Println(line, column, number)
		}
		if !isValid {
			number++
			continue
		}

		board[line][column] = number
		printBoard(board)
		board, err = solveSudoku(board)

		if err != nil {
			board[line][column] = 0
			number++
		}
	}

	if board[line][column] == 0 {
		return board, errors.New("Return")
	}

	return solveSudoku(board)
}

func getNextEmptyPosition(board [9][9]int) (int, int, error) {
	for line, columns := range board {
		for column := range columns {
			if board[line][column] == 0 {
				return line, column, nil
			}
		}
	}

	return 0, 0, errors.New("empty position not found")
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	countLine := 0
	for line, columns := range board {
		fmt.Print("| ")
		countColumn := 0
		for column := range columns {
			fmt.Printf("%d ", board[line][column])
			countColumn++
			if countColumn%3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println("")
		countLine++
		if countLine%3 == 0 {
			fmt.Println("+-------+-------+-------+")

		}
	}
}

func validatePosition(board [9][9]int, number, line, column int) bool {
	errors := [3]map[string]int{}

	//Validate sedule
	seduleColumn := column - (column % seduleLength)
	seduleLine := line - (line % seduleLength)

	for y := seduleLine; y < seduleLine+seduleLength; y++ {
		for x := seduleColumn; x < seduleColumn+seduleLength; x++ {
			if number == board[y][x] {
				errors[0] = map[string]int{"column": x, "line": y}
				return false
			}
		}
	}

	//Validate line
	for x := 0; x < len(board[0]); x++ {
		if number == board[line][x] {
			errors[1] = map[string]int{"column": x, "line": line}
			return false
		}
	}

	//Validate column
	for y := 0; y < len(board); y++ {
		if number == board[y][column] {
			errors[2] = map[string]int{"column": y, "line": line}
			return false
		}
	}

	return true
}
