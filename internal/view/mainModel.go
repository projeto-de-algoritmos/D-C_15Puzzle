package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	board board
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) View() string {
	return m.board.t.Render()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}	
	}
	return m, nil
}

func InitMainModel() MainModel {
	var m MainModel
	m.board = newBoard();
	return m
}
