# Go Tic Tac Toe

A simple command-line Tic Tac Toe game implemented in Go.

## Description

This is a two-player Tic Tac Toe game that runs in the terminal. Players take turns placing their marks (X and O) on a 3x3 grid. The first player to get three of their marks in a row (horizontally, vertically, or diagonally) wins.

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-tictactoe.git
cd go-tictactoe
```

2. Run the game:
```bash
go run tictactoe/main.go
```

## How to Play

1. The game board is a 3x3 grid with coordinates:
```
  0 1 2
0     
1     
2     
```

2. Players take turns entering their moves in the format "row column" (e.g., "1 1" for the center)
3. Valid coordinates are 0-2 for both row and column
4. The game alternates between players X and O
5. The first player to get three marks in a row wins
6. If the board fills up with no winner, the game ends in a tie

## Example Moves

- "0 0" - Top-left corner
- "1 1" - Center
- "2 2" - Bottom-right corner

## Features

- Input validation
- Clear board display
- Win condition checking
- Tie game detection
- Alternating players

## License

MIT License