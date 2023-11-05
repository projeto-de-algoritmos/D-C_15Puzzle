package view

import (
	"15puzzle/internal/game"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	board *board
	g *game.Game
	width, height int
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		m.board.t.Render(),
	)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "h":
			if m.board.selected.x > 0 {
				m.board.selected.x--
			}
		case "l":
			if m.board.selected.x < 3 {
				m.board.selected.x++
			}
		case "j":
			if m.board.selected.y < 3 {
				m.board.selected.y++
			}
		case "k":
			if m.board.selected.y > 0 {
				m.board.selected.y--
			}
		case tea.KeySpace.String():
			m.board.swap()
		}	
	}
	return m, cmd
}

func InitMainModel() MainModel {
	var m MainModel
	m.g = game.NewGame()
	m.board = newBoard(m.g.Puzzle);
	return m
}
