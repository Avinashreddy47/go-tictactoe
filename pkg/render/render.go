package render

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Avinashreddy47/go-tictactoe/pkg/game"
)

type Renderer struct {
	game *game.Game
}

func NewRenderer(game *game.Game) *Renderer {
	return &Renderer{game: game}
}

func (r *Renderer) ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (r *Renderer) getSpeedIndicator() string {
	speed := float64(r.game.InitialSpeed-r.game.Speed) / float64(r.game.InitialSpeed-r.game.MaxSpeed)
	level := int(speed * 5) // 5 levels of speed
	if level > 5 {
		level = 5
	}
	return "Speed: " + "★"*level + "☆"*(5-level)
}

func (r *Renderer) Draw() {
	r.ClearScreen()

	// Create empty board
	board := make([][]string, game.Height)
	for i := range board {
		board[i] = make([]string, game.Width)
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	// Place snake
	for _, p := range r.game.Snake {
		board[p.Y][p.X] = "█"
	}

	// Place food
	board[r.game.Food.Y][r.game.Food.X] = "●"

	// Draw board
	fmt.Println("Score:", r.game.Score)
	fmt.Println("High Score:", r.game.HighScore)
	fmt.Println(r.getSpeedIndicator())
	fmt.Println("Use W/A/S/D to move, Q to quit")
	fmt.Println("┌" + "─" + "┐")
	for _, row := range board {
		fmt.Print("│")
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println("│")
	}
	fmt.Println("└" + "─" + "┘")
}
