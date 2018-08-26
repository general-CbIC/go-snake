package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
)

// Field represents game field
type Field struct {
	width  int
	height int
	blocks [][]int8
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
	output := ""
	for _, row := range f.blocks {
		for _, cell := range row {
			switch cell {
			case -1:
				output += "X"
			case 0:
				output += " "
			case 1:
				output += "▆"
			case 2:
				output += "Ѽ"
			default:
				output += " "
			}
			output += "\t"
		}
		output += "\n"
	}

	fmt.Fprintf(writer, output+"\n")
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
