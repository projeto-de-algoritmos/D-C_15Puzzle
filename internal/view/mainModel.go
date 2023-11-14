package view

import (
	"15puzzle/internal/game"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type board struct {
	isRender bool
	board *board
	game *game.Game
	t *table.Table
}

func (m board) Init() tea.Cmd {
	return nil
}

func (b board) View() string {
	return b.t.Render()
}

func (b board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			b.isRender = false
		case "ctrl+c":
			return b, tea.Quit
		case "h", tea.KeyLeft.String():
			b.game.DecreaseY()
		case "l", tea.KeyRight.String():
			b.game.IncreaseY()
		case "j", tea.KeyDown.String():
			b.game.IncreaseX()
		case "k", tea.KeyUp.String():
			b.game.DecreaseX()
		case tea.KeySpace.String():
			b.game.Swap() 
		}	
	}
	return b, cmd
}

func newBoard (size int) board{

	game := game.NewGame(size)

	var b board
	b.game = game
	b.isRender = true

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
