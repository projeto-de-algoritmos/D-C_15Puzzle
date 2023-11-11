package view

import (
	"15puzzle/internal/game"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	board *board
	game *game.Game
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

		case "h", tea.KeyLeft.String():
			m.game.DecreaseY()
		case "l", tea.KeyRight.String():
			m.game.IncreaseY()
		case "j", tea.KeyDown.String():
			m.game.IncreaseX()
		case "k", tea.KeyUp.String():
			m.game.DecreaseX()
		case tea.KeySpace.String():
			m.game.Swap() 
		}	
	}
	return m, cmd
}

func InitMainModel() MainModel {
	var m MainModel
	m.game = game.NewGame(4)
	m.board = newBoard(m.game);
	return m
}
