package main

import (
	"fmt"
	"strconv"
	"strings"

	"calc/stack"
	"calc/display"
)

var s stack.Stack
var p []float64

func blankCalc() {
	display.ClearTerm()
	fmt.Println(" D: ")
	fmt.Println(" C: ")
	fmt.Println(" B: ")
	fmt.Println(" A: ")
	fmt.Println("-------------")
	fmt.Print(". ")
	display.GoToend()
}

func updateValues() {

	p = s.Preview(4)

	for i := 0; i < 4; i++ {
		display.ClearLine(uint16(4 - i), 5) // move cursor and clean line

		if i < len(p) {
			fmt.Print(p[i])
		}
	}

	display.GoToend()
}

func updateValue(value float64, row uint16) {
	display.MoveCursorTo(row, 5)
	fmt.Print(value)
	display.GoToend()
}

func parseInput(i string) {

}

func waitForInput() {
	for {
		display.ClearLine(6, 3)
		display.MoveCursorTo(6, 3)
		i := display.GetInput()
		parseInput(i)
	}
}



func main() {
	blankCalc()
	waitForInput()
	updateValue(s.Pop().Value, 4)
}
