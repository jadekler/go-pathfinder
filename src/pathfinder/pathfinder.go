package pathfinder

type Pathfinder struct {
	board *Board
}

func NewPathfinder(board *Board) *Pathfinder {
	return &Pathfinder{
		board: board,
	}
}

func (p *Pathfinder) Traverse() []Tile {
	return nil
}
