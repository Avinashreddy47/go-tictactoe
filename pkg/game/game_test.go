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

	// Test initial snake position is in the middle
	head := game.Snake[0]
	if head.X != Width/2 || head.Y != Height/2 {
		t.Errorf("Expected snake to start at (%d,%d), got (%d,%d)", Width/2, Height/2, head.X, head.Y)
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

	// Test self collision
	game = NewGame()
	game.Snake = []Point{
		{X: 5, Y: 5},
		{X: 4, Y: 5},
		{X: 3, Y: 5},
	}
	game.Direction = "LEFT"
	game.Move()
	if !game.GameOver {
		t.Error("Expected game to be over after self collision")
	}

	// Test food eating
	game = NewGame()
	game.Snake[0] = Point{X: game.Food.X - 1, Y: game.Food.Y}
	game.Direction = "RIGHT"
	initialScore := game.Score
	game.Move()
	if game.Score != initialScore+1 {
		t.Errorf("Expected score to increase by 1, got %d", game.Score-initialScore)
	}
	if len(game.Snake) != 2 {
		t.Errorf("Expected snake length to increase by 1, got %d", len(game.Snake))
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

	// Test food within bounds
	if game.Food.X < 0 || game.Food.X >= Width || game.Food.Y < 0 || game.Food.Y >= Height {
		t.Errorf("Food spawned outside bounds at (%d,%d)", game.Food.X, game.Food.Y)
	}
}

func TestDirectionChange(t *testing.T) {
	game := NewGame()

	// Test valid direction changes
	testCases := []struct {
		current  string
		new      string
		expected bool
	}{
		{"RIGHT", "UP", true},
		{"RIGHT", "DOWN", true},
		{"RIGHT", "LEFT", false},
		{"UP", "LEFT", true},
		{"UP", "RIGHT", true},
		{"UP", "DOWN", false},
		{"LEFT", "UP", true},
		{"LEFT", "DOWN", true},
		{"LEFT", "RIGHT", false},
		{"DOWN", "LEFT", true},
		{"DOWN", "RIGHT", true},
		{"DOWN", "UP", false},
	}

	for _, tc := range testCases {
		game.Direction = tc.current
		game.Move()
		game.Direction = tc.new
		game.Move()
		if game.GameOver != tc.expected {
			t.Errorf("Direction change from %s to %s: expected game over %v, got %v",
				tc.current, tc.new, tc.expected, game.GameOver)
		}
	}
}
