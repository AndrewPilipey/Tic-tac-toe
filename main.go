package main

import "fmt"

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
)

func main() {
	for {
		fmt.Println("Main menu")
		fmt.Println("1. Start a new game")
		fmt.Println("2. Exit")

		var choise int
		_, _ = fmt.Scan(&choise)

		switch choise {
		case 1:
			playGame()
		case 2:
			fmt.Println("Exiting the programm")
			return
		default:
			fmt.Println("Wrong choice. Try again")
		}
	}
}

func playGame() {
	mapField := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	for move := 1; move <= 9; move++ {
		isZeroNow := move%2 == 0

		if isZeroNow {
			fmt.Println("Turn of 'O'")
		} else {
			fmt.Println("Turn of 'X'")
		}

		printMap(mapField)
		fmt.Println("Enter a number of your turn:")

		cellNumber := getPlayerCellNumber(mapField)
		mapField = makeMove(mapField, cellNumber, isZeroNow)

		if hasWinner(mapField) {
			printMap(mapField)
			if isZeroNow {
				fmt.Println("'O' is winner")
			} else {
				fmt.Println("'X' is winner")
			}
			return
		}
		if move == 9 {
			printMap(mapField)
			fmt.Println("Draw") // ничья
			return
		}
	}
}

func getPlayerCellNumber(mapField [3][3]string) string {
	for {
		var input string
		_, _ = fmt.Scan(&input)

		if !(input == "1" || input == "2" || input == "3" ||
			input == "4" || input == "5" || input == "6" ||
			input == "7" || input == "8" || input == "9") {
			fmt.Println("Wrong input. Please, enter a digit from 1 to 9")
			continue
		}

		if !isMoveCorrect(mapField, input) {
			fmt.Println("Wrong input. Please, enter a number of empty cell")
			continue
		}
		return input
	}
}

func makeMove(mapField [3][3]string, cellNumber string, isZeroNow bool) [3][3]string {
	symbol := "X"
	if isZeroNow {
		symbol = "O"
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if mapField[i][j] == cellNumber {
				mapField[i][j] = symbol
				return mapField
			}
		}
	}
	return mapField
}

func isMoveCorrect(mapField [3][3]string, cellNumber string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if mapField[i][j] == cellNumber {
				return true
			}
		}
	}
	return false
}

func printMap(mapField [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := mapField[i][j]
			switch cell {
			case "X":
				fmt.Print(Red, cell, Reset, " ")
			case "O":
				fmt.Print(Blue, cell, Reset, " ")
			default:
				fmt.Print(cell, " ")
			}
		}
		fmt.Println()
	}
}

func hasWinner(mapField [3][3]string) bool {
	if mapField[0][0] == mapField[1][1] && mapField[0][0] == mapField[2][2] {
		return true
	}
	if mapField[0][2] == mapField[1][1] && mapField[0][2] == mapField[2][0] {
		return true
	}
	if mapField[0][0] == mapField[0][1] && mapField[0][0] == mapField[0][2] {
		return true
	}
	if mapField[1][0] == mapField[1][1] && mapField[1][0] == mapField[1][2] {
		return true
	}
	if mapField[2][0] == mapField[2][1] && mapField[2][0] == mapField[2][2] {
		return true
	}
	if mapField[0][0] == mapField[1][0] && mapField[0][0] == mapField[2][0] {
		return true
	}
	if mapField[0][1] == mapField[1][1] && mapField[0][1] == mapField[2][1] {
		return true
	}
	if mapField[0][2] == mapField[1][2] && mapField[0][2] == mapField[2][2] {
		return true
	}
	return false
}
