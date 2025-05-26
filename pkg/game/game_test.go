package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	// Test initial snake position
	if len(game.Snake) != 1 {
		t.Errorf("Expected snake length to be 1, got %d", len(game.Snake))
	}

	// Test initial direction
	if game.Direction != "RIGHT" {
		t.Errorf("Expected initial direction to be RIGHT, got %s", game.Direction)
	}

	// Test initial score
	if game.Score != 0 {
		t.Errorf("Expected initial score to be 0, got %d", game.Score)
	}

	// Test game over state
	if game.GameOver {
		t.Error("Expected game to not be over initially")
	}
}

func TestMove(t *testing.T) {
	game := NewGame()
	initialLength := len(game.Snake)

	// Test moving right
	game.Move()
	if len(game.Snake) != initialLength {
		t.Errorf("Expected snake length to remain %d, got %d", initialLength, len(game.Snake))
	}

	// Test wall collision
	game.Snake[0] = Point{X: Width, Y: Height / 2}
	game.Move()
	if !game.GameOver {
		t.Error("Expected game to be over after wall collision")
	}
}

func TestFoodGeneration(t *testing.T) {
	game := NewGame()
	initialFood := game.Food

	// Test food generation
	game.generateFood()
	if game.Food == initialFood {
		t.Error("Expected food position to change")
	}

	// Test food not spawning on snake
	for _, p := range game.Snake {
		if p.X == game.Food.X && p.Y == game.Food.Y {
			t.Error("Food spawned on snake")
		}
	}
}
