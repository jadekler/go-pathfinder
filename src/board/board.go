package board

import "fmt"

type Tile struct {
	X         int
	Y         int
	Blocked   bool
	Traversed bool
}

type Board struct {
	Tiles [][]Tile
	GoalX int
	GoalY int
	height int
	width int
}

/*
 [0][0] [0][1] [0][2]
 [1][0] [1][1] [1][2]
 [2][0] [2][1] [2][2]
*/
func NewBoard(height, width, goalX, goalY int) *Board {
	board := Board{
		Tiles:  make([][]Tile, height),
		GoalX:  goalX,
		GoalY:  goalY,
		height: height,
		width:  width,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			board.Tiles[y] = append(board.Tiles[y], Tile{x, y, false, false})
		}
	}

	return &board
}

func (b *Board) SetBlocked(x int, y int) {
	b.Tiles[y][x].Blocked = true
}

func (b *Board) Draw() {
	fmt.Println()
	for y := b.height - 1; y > 0; y-- {
		for x := 0; x < b.width; x++ {
			t := b.Tiles[y][x]
			if t.Blocked {
				fmt.Print("x")
			} else if t.Traversed {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}