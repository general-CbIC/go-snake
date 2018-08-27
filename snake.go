package main

import (
	"time"

	"github.com/gosuri/uilive"
	"github.com/olekukonko/tablewriter"
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
		{-1, 0, 0, 1, 0, 0, 0, 2, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
}

func (f *Field) render(writer *uilive.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetBorder(false)

	for _, row := range f.blocks {
		strRow := make([]string, 10)
		for index, cell := range row {
			switch cell {
			case -1:
				strRow[index] = "❎"
			case 0:
				strRow[index] = " "
			case 1:
				strRow[index] = "🔲"
			case 2:
				strRow[index] = "🥚"
			default:
				strRow[index] = " "
			}
		}
		table.Append(strRow)
	}

	table.Render()
}

func makeTurn(field *Field, writer *uilive.Writer) {
	field.render(writer)
	time.Sleep(time.Second)
	makeTurn(field, writer)
}

func main() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	field := new(Field)
	field.width = 10
	field.height = 10
	field.fillBlocks()

	makeTurn(field, writer)
}
