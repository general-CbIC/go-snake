package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Direction is an ENUM for show snake direction
type Direction uint8

const (
	// LEFT is left direction (0)
	LEFT Direction = 0
	// UP is up direction (1)
	UP Direction = 1
	// RIGHT is right direction (2)
	RIGHT Direction = 2
	// DOWN is down direction (3)
	DOWN Direction = 3
)

// Coordinate is {x, y} point struct in game Field
type Coordinate struct {
	x int
	y int
}

// Snake storages current direction and coordinates of snake.
type Snake struct {
	*tl.Entity
	direction Direction
	head      Coordinate
	body      []Coordinate
}

// Tick - method for snake control
func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			snake.direction = RIGHT
		case tl.KeyArrowLeft:
			snake.direction = LEFT
		case tl.KeyArrowUp:
			snake.direction = UP
		case tl.KeyArrowDown:
			snake.direction = DOWN
		case 0:
			switch event.Ch {
			case 100:
				snake.direction = RIGHT
			case 97:
				snake.direction = LEFT
			case 119:
				snake.direction = UP
			case 115:
				snake.direction = DOWN
			}
		}
	}
}

// Draw - updates each frame
func (snake *Snake) Draw(screen *tl.Screen) {
	x, y := snake.Position()
	snake.SetPosition(x+1, y)
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	// Field
	level.AddEntity(tl.NewRectangle(20, 1, 140, 42, tl.ColorWhite))

	// Walls
	level.AddEntity(tl.NewRectangle(20, 1, 140, 1, tl.ColorRed))
	level.AddEntity(tl.NewRectangle(20, 1, 1, 42, tl.ColorRed))
	level.AddEntity(tl.NewRectangle(160, 1, 1, 42, tl.ColorRed))
	level.AddEntity(tl.NewRectangle(21, 42, 140, 1, tl.ColorRed))

	// Snake
	snake := Snake{Entity: tl.NewEntity(40, 20, 1, 1)}
	snake.body = make([]Coordinate, 100)

	snake.SetCell(0, 0, &tl.Cell{Ch: '🔲'})

	level.AddEntity(snake.Entity)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(10.0)
	game.Start()
}
