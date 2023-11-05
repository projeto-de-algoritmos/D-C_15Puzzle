package main

import (
	"15puzzle/internal/view"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mainModel := view.InitMainModel()

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		os.Exit(1)
	}

	p := tea.NewProgram(mainModel)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
