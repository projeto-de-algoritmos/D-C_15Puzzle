package main

import (
	"15puzzle/internal/view"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mainModel := view.InitMainModel()
	p := tea.NewProgram(mainModel)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
