package game

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	Width                  = 20
	Height                 = 10
	SpeedIncreaseThreshold = 5   // Increase speed every 5 points
	MaxSpeed               = 50  // Minimum delay between moves in milliseconds
	InitialSpeed           = 200 // Initial delay between moves in milliseconds
	HighScoreFile          = ".snake_highscore.json"

	// Difficulty levels
	Easy   = "easy"
	Medium = "medium"
	Hard   = "hard"

	// Food types
	NormalFood = iota
	SpeedFood
	SlowFood
	DoublePointsFood
	ShrinkFood
)

type Point struct {
	X, Y int
}

type Food struct {
	Point
	Type     int
	Duration time.Duration
}

type Game struct {
	Snake      []Point
	Food       Food
	Direction  string
	Score      int
	HighScore  int
	GameOver   bool
	Speed      int // Current speed (delay in milliseconds)
	Difficulty string
	Effects    map[string]time.Time // Active effects and their end times
}

func NewGame(difficulty string) *Game {
	rand.Seed(time.Now().UnixNano())

	// Set initial speed based on difficulty
	initialSpeed := InitialSpeed
	switch difficulty {
	case Easy:
		initialSpeed = 250
	case Medium:
		initialSpeed = 200
	case Hard:
		initialSpeed = 150
	}

	game := &Game{
		Snake:      []Point{{X: Width / 2, Y: Height / 2}},
		Direction:  "RIGHT",
		Score:      0,
		GameOver:   false,
		Speed:      initialSpeed,
		Difficulty: difficulty,
		Effects:    make(map[string]time.Time),
	}

	game.loadHighScore()
	game.generateFood()
	return game
}

func (g *Game) playSound(sound string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-c", fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlaySync()", sound))
	} else {
		cmd = exec.Command("afplay", sound)
	}
	cmd.Run()
}

func (g *Game) loadHighScore() {
	data, err := os.ReadFile(HighScoreFile)
	if err != nil {
		g.HighScore = 0
		return
	}
	json.Unmarshal(data, &g.HighScore)
}

func (g *Game) saveHighScore() {
	if g.Score > g.HighScore {
		g.HighScore = g.Score
		data, _ := json.Marshal(g.HighScore)
		os.WriteFile(HighScoreFile, data, 0644)
	}
}

func (g *Game) generateFood() {
	for {
		point := Point{
			X: rand.Intn(Width),
			Y: rand.Intn(Height),
		}

		// Make sure food doesn't spawn on snake
		valid := true
		for _, p := range g.Snake {
			if p.X == point.X && p.Y == point.Y {
				valid = false
				break
			}
		}
		if valid {
			// 70% chance for normal food, 30% for special food
			foodType := NormalFood
			if rand.Float32() < 0.3 {
				foodType = rand.Intn(4) + 1 // Random special food type
			}
			g.Food = Food{
				Point:    point,
				Type:     foodType,
				Duration: 10 * time.Second, // Default duration for effects
			}
			break
		}
	}
}

func (g *Game) applyEffect(effect string, duration time.Duration) {
	g.Effects[effect] = time.Now().Add(duration)
}

func (g *Game) hasEffect(effect string) bool {
	if endTime, exists := g.Effects[effect]; exists {
		if time.Now().Before(endTime) {
			return true
		}
		delete(g.Effects, effect)
	}
	return false
}

func (g *Game) Move() {
	head := g.Snake[0]
	newHead := Point{X: head.X, Y: head.Y}

	switch g.Direction {
	case "UP":
		newHead.Y--
	case "DOWN":
		newHead.Y++
	case "LEFT":
		newHead.X--
	case "RIGHT":
		newHead.X++
	}

	// Check for collisions
	if newHead.X < 0 || newHead.X >= Width || newHead.Y < 0 || newHead.Y >= Height {
		g.GameOver = true
		g.playSound("sounds/gameover.wav")
		g.saveHighScore()
		return
	}

	// Check for self collision
	for _, p := range g.Snake {
		if p.X == newHead.X && p.Y == newHead.Y {
			g.GameOver = true
			g.playSound("sounds/gameover.wav")
			g.saveHighScore()
			return
		}
	}

	// Move snake
	g.Snake = append([]Point{newHead}, g.Snake...)

	// Check if food is eaten
	if newHead.X == g.Food.X && newHead.Y == g.Food.Y {
		g.playSound("sounds/eat.wav")

		// Apply food effects
		switch g.Food.Type {
		case SpeedFood:
			g.applyEffect("speed", g.Food.Duration)
			g.Speed = MaxSpeed
		case SlowFood:
			g.applyEffect("slow", g.Food.Duration)
			g.Speed = InitialSpeed
		case DoublePointsFood:
			g.applyEffect("doublePoints", g.Food.Duration)
		case ShrinkFood:
			if len(g.Snake) > 1 {
				g.Snake = g.Snake[:len(g.Snake)-1]
			}
		}

		// Calculate score
		scoreIncrease := 1
		if g.hasEffect("doublePoints") {
			scoreIncrease = 2
		}
		g.Score += scoreIncrease

		g.generateFood()

		// Increase speed based on difficulty (if no speed effect is active)
		if !g.hasEffect("speed") && !g.hasEffect("slow") {
			if g.Score%SpeedIncreaseThreshold == 0 && g.Speed > MaxSpeed {
				switch g.Difficulty {
				case Easy:
					g.Speed -= 20
				case Medium:
					g.Speed -= 30
				case Hard:
					g.Speed -= 40
				}
			}
		}
	} else {
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
}
