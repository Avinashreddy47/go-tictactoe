# Go Snake Game

A simple command-line Snake game implemented in Go.

## Description

This is a classic Snake game that runs in the terminal. Control the snake using W/A/S/D keys to eat food and grow longer. The game ends if the snake hits the wall or itself.

## Prerequisites

* Go 1.21 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Avinashreddy47/go-tictactoe.git
cd go-tictactoe
```

2. Build the game:
```bash
go build ./cmd/snake-game
```

3. Run the game:
```bash
./snake-game
```

## How to Play

1. Use the following keys to control the snake:
   - W: Move Up
   - A: Move Left
   - S: Move Down
   - D: Move Right
   - Q: Quit Game

2. Game Rules:
   - Eat the food (●) to grow longer
   - Avoid hitting the walls
   - Avoid hitting yourself
   - Try to get the highest score possible

## Features

* Simple terminal-based interface
* Score tracking
* Collision detection
* Smooth controls
* Clear visual feedback

## Project Structure

```
.
├── cmd/
│   └── snake-game/
│       └── main.go
├── pkg/
│   ├── game/
│   │   └── game.go
│   ├── input/
│   │   └── input.go
│   └── render/
│       └── render.go
├── go.mod
└── README.md
```

## License

MIT License