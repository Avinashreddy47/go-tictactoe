package game

import (
	"math/rand"
	"time"
)

const (
	Width                  = 20
	Height                 = 10
	SpeedIncreaseThreshold = 5   // Increase speed every 5 points
	MaxSpeed               = 50  // Minimum delay between moves in milliseconds
	InitialSpeed           = 200 // Initial delay between moves in milliseconds
)

type Point struct {
	X, Y int
}

type Game struct {
	Snake     []Point
	Food      Point
	Direction string
	Score     int
	GameOver  bool
	Speed     int // Current speed (delay in milliseconds)
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	game := &Game{
		Snake:     []Point{{X: Width / 2, Y: Height / 2}},
		Direction: "RIGHT",
		Score:     0,
		GameOver:  false,
		Speed:     InitialSpeed,
	}

	game.generateFood()
	return game
}

func (g *Game) generateFood() {
	for {
		g.Food = Point{
			X: rand.Intn(Width),
			Y: rand.Intn(Height),
		}

		// Make sure food doesn't spawn on snake
		valid := true
		for _, p := range g.Snake {
			if p.X == g.Food.X && p.Y == g.Food.Y {
				valid = false
				break
			}
		}
		if valid {
			break
		}
	}
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
		return
	}

	// Check for self collision
	for _, p := range g.Snake {
		if p.X == newHead.X && p.Y == newHead.Y {
			g.GameOver = true
			return
		}
	}

	// Move snake
	g.Snake = append([]Point{newHead}, g.Snake...)

	// Check if food is eaten
	if newHead.X == g.Food.X && newHead.Y == g.Food.Y {
		g.Score++
		g.generateFood()
		// Increase speed every SpeedIncreaseThreshold points
		if g.Score%SpeedIncreaseThreshold == 0 && g.Speed > MaxSpeed {
			g.Speed -= 30 // Decrease delay by 30ms
		}
	} else {
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
}
