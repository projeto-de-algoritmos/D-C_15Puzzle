package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Menu struct {
	choices []string
	cursor int
	selected map[int]int
	board board
	width, height int
}

func (m Menu) Init() tea.Cmd{
	return nil
}

func (m Menu) View() string{
	var s string
	if m.board.isRender {
		s = m.board.View()
	} else {
		s = "15 PUZZLE\n\n"

		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		s += "\n"
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.board.isRender {
		e, cmd := m.board.Update(msg)
		m.board = e.(board)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "j", tea.KeyDown.String():
			if m.cursor < len(m.selected)-1 {
				m.cursor++
			}
		case "k", tea.KeyUp.String():
			if m.cursor > 0 {
				m.cursor--
			}
		case " ":
			m.board = newBoard(m.selected[m.cursor])
			m.board.isRender = true
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, cmd
}
	

func NewMenu() Menu{
	selected := make(map[int]int)
	selected[0] = 3
	selected[1] = 4
	selected[2] = 5
	selected[3] = 8
	return Menu {
		choices: []string{"Fácil", "Médio", "Difícil", "Desafio"},
		selected: selected,
	}
}
