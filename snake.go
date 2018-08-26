package main

import (
	"time"

	"github.com/gosuri/uilive"
	"github.com/olekukonko/tablewriter"
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
	table := tablewriter.NewWriter(writer)
	table.SetBorder(false)

	for _, row := range f.blocks {
		strRow := make([]string, 10)
		for index, cell := range row {
			switch cell {
			case -1:
				strRow[index] = "X"
			case 0:
				strRow[index] = " "
			case 1:
				strRow[index] = "▆"
			case 2:
				strRow[index] = "Ѽ"
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
