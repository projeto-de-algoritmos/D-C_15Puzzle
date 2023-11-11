package view

import (
	"15puzzle/internal/game"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type point struct {
	x, y int
}

type board struct {
	t *table.Table
	selected point
	empty point
}



func (t *board) updateBoard() {
	t.t.ClearRows()
}

func newBoard (game *game.Game) *board{

	b := new(board)
	b.selected = point{0, 0}
	b.empty = point{3, 3}

	b.t = table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle().
				Height(3).Width(5).
				Bold(true).
				Align(lipgloss.Center, lipgloss.Center)
				if r, c := game.GetSelectedPointer(); r+1 == row && c == col {
				style = style.Foreground(lipgloss.Color("99"))
			}
			return style
		}).
		Rows(game.Puzzle...)	

	return b
}
