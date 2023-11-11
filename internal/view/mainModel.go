package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

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
		case "q":
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
