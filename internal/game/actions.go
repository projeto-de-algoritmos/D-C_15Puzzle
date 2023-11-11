package game


type point struct {
	x, y int
}

func diff(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}

func (g Game) isPossibleToSwap() bool {
	ng := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, p := range ng {
		if (p == diff(g.selected, g.empty)) {
			return true
		}
	}
	return false
}

func (g *Game) Swap () {
	if !g.isPossibleToSwap() {
		return
	}
	aux := g.Puzzle[g.selected.x][g.selected.y]
	g.Puzzle[g.selected.x][g.selected.y] = g.Puzzle[g.empty.x][g.empty.y]
	g.Puzzle[g.empty.x][g.empty.y] = aux
	g.empty = g.selected
}

func (g Game) GetSelectedPointer() (row, col int) {
	return g.selected.x, g.selected.y
}

func (g *Game) IncreaseX() {
	if g.selected.x < g.size-1 {
		g.selected.x++
	}
}

func (g *Game) IncreaseY() {
	if g.selected.y < g.size-1 {
		g.selected.y++
	}
}

func (g *Game) DecreaseX() {
	if g.selected.x > 0 {
		g.selected.x--
	}
}

func (g *Game) DecreaseY() {
	if g.selected.y > 0 {
		g.selected.y--
	}
}
