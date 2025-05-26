package input

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/Avinashreddy47/go-tictactoe/pkg/game"
)

type InputHandler struct {
	game *game.Game
}

func NewInputHandler(game *game.Game) *InputHandler {
	return &InputHandler{game: game}
}

func (h *InputHandler) HandleInput() bool {
	// Set up non-blocking input
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan byte, 1)

	go func() {
		char, _ := reader.ReadByte()
		ch <- char
	}()

	// Wait for input with timeout
	select {
	case char := <-ch:
		switch char {
		case 'w', 'W':
			if h.game.Direction != "DOWN" {
				h.game.Direction = "UP"
				log.Println("Direction changed to UP")
			} else {
				log.Println("Cannot change direction to UP when moving DOWN")
			}
		case 's', 'S':
			if h.game.Direction != "UP" {
				h.game.Direction = "DOWN"
				log.Println("Direction changed to DOWN")
			} else {
				log.Println("Cannot change direction to DOWN when moving UP")
			}
		case 'a', 'A':
			if h.game.Direction != "RIGHT" {
				h.game.Direction = "LEFT"
				log.Println("Direction changed to LEFT")
			} else {
				log.Println("Cannot change direction to LEFT when moving RIGHT")
			}
		case 'd', 'D':
			if h.game.Direction != "LEFT" {
				h.game.Direction = "RIGHT"
				log.Println("Direction changed to RIGHT")
			} else {
				log.Println("Cannot change direction to RIGHT when moving LEFT")
			}
		case 'q', 'Q':
			log.Println("Game quit")
			return false
		case 'p', 'P':
			h.game.TogglePause()
		}
	case <-time.After(time.Duration(h.game.Speed) * time.Millisecond):
		// No input received within timeout
	}

	return true
}
