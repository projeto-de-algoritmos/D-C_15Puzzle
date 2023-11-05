package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	board *board
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) View() string {
	return m.board.t.Render()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "h":
			if m.board.x > 0 {
				m.board.x--
			}
		case "l":
			if m.board.x < 3 {
				m.board.x++
			}
		case "j":
			if m.board.y < 4 {
				m.board.y++
			}
		case "k":
			if m.board.y > 1 {
				m.board.y--
			}
		}	
	}
	return m, cmd
}

func InitMainModel() MainModel {
	var m MainModel
	m.board = newBoard();
	return m
}
