package view

import (
	"log"
	"math/rand"
	"strconv"

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

	newGame := rand.Perm(15)

	b.matrix = make([][]string, 4)
	for i := range b.matrix {
		b.matrix[i] = make([]string, 4)
	}

	for i := 0; i<15; i++ {
		b.matrix[i/4][i%4] =  strconv.Itoa(newGame[i]+1)
	}

	b.matrix[3][3] = "X"

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
