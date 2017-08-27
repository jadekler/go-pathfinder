package main

import (
	"pathfinder"
	"board"
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
	b := board.NewBoard(10, 10, 9, 8)
	b.SetBlocked(8, 8)
	b.SetBlocked(8, 7)
	b.SetBlocked(9, 7)
	b.Draw()

	p := pathfinder.NewPathfinder(b, 3, 4)
	path := p.Traverse()
	path.Draw()
}

