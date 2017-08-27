package main

import (
	"pathfinder"
	"fmt"
)

func main() {
	/*					9,9
	o o o o o o o o o o
	o o o o o o o o x .
	o o o o o o o o x x
	o o o o o o o o o o
	o o o o o o o o o o
	o o o * o o o o o o
	o o o o o o o o o o
	o o o o o o o o o o
	o o o o o o o o o o
	o o o o o o o o o o 9,0
	 */
	b := pathfinder.NewBoard(10, 10, 3, 4, 9, 8)

	p := pathfinder.NewPathfinder(b)
	path := p.Traverse()

	fmt.Println(path)
}

