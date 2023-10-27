package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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

var m = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 0, 0, 0, 0, 8, 0},
	{0, 0, 0, 1, 0, 2, 0, 0, 0},
	{0, 0, 0, 0, 5, 0, 7, 0, 0},
	{0, 0, 4, 0, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 0, 0, 0},
	{5, 0, 0, 0, 0, 0, 0, 7, 3},
	{0, 0, 2, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 4, 0, 0, 0, 9},
}

func main() {
	_, err := solveSudoku(m)

	if err != nil {
		fmt.Println(err)
	}
}

func solveSudoku(board [9][9]int) ([9][9]int, error) {

	line, column, err := getNextEmptyPosition(board)

	if err != nil {
		return board, nil
	}

	validNumbers := getValidNumberByPosition(board, line, column)
	//TODO: Randomize validNumbers
	for _, number := range validNumbers {

		board[line][column] = number
		CallClear()
		printBoard(board)
		board, err = solveSudoku(board)

		if err != nil {
			board[line][column] = 0
			continue
		}
		break
	}

	if board[line][column] == 0 {
		return board, errors.New("Return")
	}

	return solveSudoku(board)
}

func getValidNumberByPosition(board [9][9]int, line int, column int) []int {
	validNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Validate sedule
	seduleColumn := column - (column % seduleLength)
	seduleLine := line - (line % seduleLength)

	for y := seduleLine; y < seduleLine+seduleLength; y++ {
		for x := seduleColumn; x < seduleColumn+seduleLength; x++ {
			for index, number := range validNumbers {
				if number == board[y][x] {
					validNumbers = deleteSliceElement(validNumbers, index)
				}
			}
		}
	}

	//Validate line
	for x := 0; x < len(board[0]); x++ {
		for index, number := range validNumbers {
			if number == board[line][x] {
				validNumbers = deleteSliceElement(validNumbers, index)

			}
		}
	}

	//Validate column
	for y := 0; y < len(board); y++ {
		for index, number := range validNumbers {
			if number == board[y][column] {
				validNumbers = deleteSliceElement(validNumbers, index)
			}
		}
	}

	return validNumbers
}

func deleteSliceElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
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
	var boardStr string
	boardStr = "+-------+-------+-------+\n"
	countLine := 0
	for line, columns := range board {
		boardStr += "| "
		countColumn := 0
		for column := range columns {
			strValue := strconv.Itoa(board[line][column])
			boardStr += strValue + " "
			countColumn++
			if countColumn%3 == 0 {
				boardStr += "| "
			}
		}
		boardStr += "\n"
		countLine++
		if countLine%3 == 0 {
			boardStr += "+-------+-------+-------+\n"
		}
	}
	fmt.Print(boardStr)
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
