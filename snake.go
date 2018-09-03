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

// globals
var snakeCell = &tl.Cell{Ch: '🔲'}

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
	snake.moveTail(snake.head)

	switch snake.direction {
	case RIGHT:
		snake.head.x++
	case LEFT:
		snake.head.x--
	case UP:
		snake.head.y--
	case DOWN:
		snake.head.y++
	}

	screen.RenderCell(snake.head.x, snake.head.y, snakeCell)

	for _, bodyCell := range snake.body {
		screen.RenderCell(bodyCell.x, bodyCell.y, snakeCell)
	}
}

func (snake *Snake) moveTail(coordinate Coordinate) {
	snake.body = append(snake.body[1:], coordinate)
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
	snake := new(Snake)
	snake.Entity = tl.NewEntity(0, 0, 1, 1)
	snake.head = Coordinate{40, 20}
	snake.body = make([]Coordinate, 4, 100)
	snake.body[0] = Coordinate{38, 22}
	snake.body[1] = Coordinate{38, 21}
	snake.body[2] = Coordinate{38, 20}
	snake.body[3] = Coordinate{39, 20}

	snake.direction = RIGHT

	snake.SetCell(0, 0, snakeCell)

	level.AddEntity(snake)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(10.0)
	game.Start()
}
