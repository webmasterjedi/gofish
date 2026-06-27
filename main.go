package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	g := tea.NewProgram(model{})
	_, err := g.Run()
	if err != nil {
		fmt.Println("There was a error.")
	}
}
