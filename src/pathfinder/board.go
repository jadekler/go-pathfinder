package pathfinder

type Tile struct {
	X       int
	Y       int
	Blocked bool
}

type Board struct {
	tiles  [][]Tile
	startX int
	startY int
	goalX  int
	goalY  int
}

/*
 [0][0] [0][1] [0][2]
 [1][0] [1][1] [1][2]
 [2][0] [2][1] [2][2]
*/
func NewBoard(height, width, startX, startY, goalX, goalY int) *Board {
	board := Board{
		tiles:  make([][]Tile, height),
		startX: startX,
		startY: startY,
		goalX:  goalX,
		goalY:  goalY,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			board.tiles[y] = append(board.tiles[y], Tile{x, y, false})
		}
	}

	return &board
}

func (b *Board) SetBlocked(x int, y int) {
	b.tiles[y][x].Blocked = true
}
