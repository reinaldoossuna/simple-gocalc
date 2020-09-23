package main

import (
	"fmt"
	"strconv"

	"calc/stack"
	"calc/display"
)

var s stack.Stack

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
	n, _ := strconv.ParseFloat("1.33", 64)
	s.Push(n)
	updateValue(s.Pop().Value, 4)
	waitForInput()
}
