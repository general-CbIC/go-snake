package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
)

// Direction is an ENUM for show snake direction
type Direction int

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

// Field represents width, height and blocks of game field
type Field struct {
	width  int
	height int
	blocks [][]int8
}

// Coordinate is {x, y} point struct in game Field
type Coordinate struct {
	x int
	y int
}

// Snake storages current direction and coordinates of snake.
type Snake struct {
	direction Direction
	head      Coordinate
	body      []Coordinate
}

// Globals
var writer = uilive.New()
var field = new(Field)
var snake = new(Snake)

func (f *Field) fillBlocks() {
	/*
		-1	=> Wall
		0		=> Empty
		1		=> Snake
		2		=> Food
	*/
	f.blocks = [][]int8{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 1, 1, 0, 0, 0, 2, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
}

func (f *Field) render() {
	output := ""

	for _, row := range f.blocks {
		for _, cell := range row {
			switch cell {
			case -1:
				output += "❎"
			case 0:
				output += "🔳"
			case 1:
				output += "🔲"
			case 2:
				output += "🥚"
			default:
				output += "🔳"
			}
		}
		output += "\n"
	}

	fmt.Fprintln(writer, output)
}

func makeStep() {
	moveTail(0, snake.head)

	switch snake.direction {
	case LEFT:
		snake.head.x--
	case UP:
		snake.head.y--
	case RIGHT:
		snake.head.x++
	case DOWN:
		snake.head.y++
	}

	field.blocks[snake.head.y][snake.head.x] = 1
}

func moveTail(i int, coordinate Coordinate) {
	if i == len(snake.body)-1 {
		field.blocks[snake.body[i].y][snake.body[i].x] = 0
	} else {
		moveTail(i+1, snake.body[i])
	}
	snake.body[i] = coordinate
}

func makeTurn() {
	makeStep()
	field.render()
	time.Sleep(time.Second)
	makeTurn()
}

func main() {
	writer.Start()
	defer writer.Stop()

	field.width = 10
	field.height = 10
	field.fillBlocks()

	snake.direction = RIGHT
	snake.head.x = 3
	snake.head.y = 3
	snake.body = make([]Coordinate, 1, 100)
	snake.body[0] = Coordinate{2, 3}

	makeTurn()
}
