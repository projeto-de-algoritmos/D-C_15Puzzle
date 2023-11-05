package view

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type board struct {
	t *table.Table
	x, y int
}

func newBoard () *board{

	b := new(board)
	b.x = 0
	b.y = 1

	row := [][]string {
		{"01", "02", "03", "04"},
		{"05", "06", "07", "08"},
		{"09", "10", "11", "12"},
		{"13", "14", "15", "X"},
	}

	b.t = table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle().
				Height(2).Width(4)

			if col == b.x && row == b.y {
				style = style.Foreground(lipgloss.Color("99"))
			}
			return style
		}).
		Rows(row...)

	return b
}
