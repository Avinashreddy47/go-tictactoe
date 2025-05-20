package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [3][3]string

func NewBoard() Board {
	return Board{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func (b Board) Print() {
	fmt.Println("\n  0 1 2")
	for i, row := range b {
		fmt.Printf("%d ", i)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b Board) IsValidMove(row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	return b[row][col] == " "
}

func (b Board) MakeMove(row, col int, player string) Board {
	newBoard := b
	newBoard[row][col] = player
	return newBoard
}

func (b Board) CheckWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b[i][0] != " " && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if b[0][i] != " " && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return true
		}
	}

	// Check diagonals
	if b[0][0] != " " && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return true
	}
	if b[0][2] != " " && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return true
	}

	return false
}

func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func main() {
	board := NewBoard()
	currentPlayer := "X"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("Enter moves as 'row column' (0-2)")

	for {
		board.Print()
		fmt.Printf("Player %s's turn. Enter row and column (0-2): ", currentPlayer)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coords := strings.Split(input, " ")

		if len(coords) != 2 {
			fmt.Println("Invalid input! Please enter row and column as two numbers.")
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input! Please enter numbers only.")
			continue
		}

		if !board.IsValidMove(row, col) {
			fmt.Println("Invalid move! Try again.")
			continue
		}

		board = board.MakeMove(row, col, currentPlayer)

		if board.CheckWin() {
			board.Print()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		if board.IsFull() {
			board.Print()
			fmt.Println("It's a tie!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
