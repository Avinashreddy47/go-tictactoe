package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Avinashreddy47/snake-game/pkg/game"
)

type InputHandler struct {
	game *game.Game
}

func NewInputHandler(game *game.Game) *InputHandler {
	return &InputHandler{game: game}
}

func (h *InputHandler) HandleInput() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter direction (W/A/S/D) or Q to quit: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading input: %v\n", err)
		return false
	}
	input = strings.TrimSpace(strings.ToUpper(input))

	log.Printf("Received input: %s\n", input)

	switch input {
	case "W":
		if h.game.Direction != "DOWN" {
			h.game.Direction = "UP"
			log.Println("Direction changed to UP")
		} else {
			log.Println("Cannot change direction to UP when moving DOWN")
		}
	case "S":
		if h.game.Direction != "UP" {
			h.game.Direction = "DOWN"
			log.Println("Direction changed to DOWN")
		} else {
			log.Println("Cannot change direction to DOWN when moving UP")
		}
	case "A":
		if h.game.Direction != "RIGHT" {
			h.game.Direction = "LEFT"
			log.Println("Direction changed to LEFT")
		} else {
			log.Println("Cannot change direction to LEFT when moving RIGHT")
		}
	case "D":
		if h.game.Direction != "LEFT" {
			h.game.Direction = "RIGHT"
			log.Println("Direction changed to RIGHT")
		} else {
			log.Println("Cannot change direction to RIGHT when moving LEFT")
		}
	case "Q":
		log.Println("Game quit")
		return false
	default:
		log.Println("Invalid input")
	}
	return true
}
