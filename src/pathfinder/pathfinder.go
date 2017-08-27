package pathfinder

import (
	"board"
	"math"
)

type Pathfinder struct {
	board  *board.Board
	startX int
	startY int
}

func NewPathfinder(board *board.Board, startX int, startY int) *Pathfinder {
	board.Tiles[startY][startX].Traversed = true

	return &Pathfinder{
		board:  board,
		startX: startX,
		startY: startY,
	}
}

func (p *Pathfinder) Traverse() board.Board {
	return traverse(*p.board, p.startX, p.startY)
}

func traverse(b board.Board, currentX int, currentY int) board.Board {
	if currentX == b.GoalX && currentY == b.GoalY {
		return b
	}

	validSurroundings := getValidSurrounds()
	

	return b
}

func getValidSurrounds() []board.Tile {

}

func distanceAsBirdFlies(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1+x2), 2) + math.Pow(float64(y1+y2), 2))
}
