package render

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Avinashreddy47/go-tictactoe/pkg/game"
)

const (
	// ANSI color codes
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
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
	speed := float64(game.InitialSpeed-r.game.Speed) / float64(game.InitialSpeed-game.MaxSpeed)
	level := int(speed * 5) // 5 levels of speed
	if level > 5 {
		level = 5
	}
	return Yellow + "Speed: " + strings.Repeat("★", level) + strings.Repeat("☆", 5-level) + Reset
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
	for i, p := range r.game.Snake {
		if i == 0 {
			board[p.Y][p.X] = Green + "█" + Reset // Head in green
		} else {
			board[p.Y][p.X] = Cyan + "█" + Reset // Body in cyan
		}
	}

	// Place food
	board[r.game.Food.Y][r.game.Food.X] = Red + "●" + Reset

	// Draw header
	fmt.Println(Cyan + "╔════════════════════════════════════════════════════════════╗" + Reset)
	fmt.Println(Cyan + "║" + Reset + "                     SNAKE GAME                     " + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╠════════════════════════════════════════════════════════════╣" + Reset)
	fmt.Println(Cyan + "║" + Reset + " Score: " + Yellow + fmt.Sprintf("%d", r.game.Score) + Reset +
		"    High Score: " + Yellow + fmt.Sprintf("%d", r.game.HighScore) + Reset +
		strings.Repeat(" ", 20) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "║" + Reset + " " + r.getSpeedIndicator() +
		strings.Repeat(" ", 35) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "║" + Reset + " Use W/A/S/D to move, Q to quit" +
		strings.Repeat(" ", 25) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╠════════════════════════════════════════════════════════════╣" + Reset)

	// Draw board
	fmt.Println(Cyan + "║" + Reset + "┌" + strings.Repeat("─", game.Width) + "┐" + Cyan + "║" + Reset)
	for _, row := range board {
		fmt.Print(Cyan + "║" + Reset + "│")
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println("│" + Cyan + "║" + Reset)
	}
	fmt.Println(Cyan + "║" + Reset + "└" + strings.Repeat("─", game.Width) + "┘" + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╚════════════════════════════════════════════════════════════╝" + Reset)
}
