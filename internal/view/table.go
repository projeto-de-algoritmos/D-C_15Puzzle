package view

import (
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type point struct {
	x, y int
}

type board struct {
	t *table.Table
	matrix [][]string
	selected point
	empty point
}

func diff(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}

func isPossibleToSwap(b board) bool {
	ng := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, p := range ng {
		if (p == diff(b.selected, b.empty)) {
			return true
		}
	}
	return false
}

func (b *board) swap() {
	if !isPossibleToSwap(*b) {
		return
	}	
	log.Println("Swap:", b.selected, b.empty)
	aux := b.matrix[b.selected.y][b.selected.x]
	b.matrix[b.selected.y][b.selected.x] = b.matrix[b.empty.y][b.empty.x]	
	b.matrix[b.empty.y][b.empty.x] = aux

	b.empty = b.selected

	b.updateBoard()
}

func (b *board) updateBoard() {
	b.t = table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle().
				Height(0).Width(4)
			if col == b.selected.x && row == b.selected.y {
				style = style.Foreground(lipgloss.Color("99"))
			}
			return style
		}).
		Headers(b.matrix[0]...).
		Rows(b.matrix[1:4]...)
}

func newBoard () *board{

	b := new(board)
	b.selected = point{0, 0}
	b.empty = point{3, 3}

	b.matrix = [][]string {
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
				Height(0).Width(4)
			if col == b.selected.x && row == b.selected.y {
				style = style.Foreground(lipgloss.Color("99"))
			}
			return style
		}).
		Headers(b.matrix[0]...).
		Rows(b.matrix[1:4]...)

	return b
}
