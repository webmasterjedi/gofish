package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

func main() {
	m := model{}
	m.spring = harmonica.NewSpring(harmonica.FPS(60), 5.0, 5.0)
	g := tea.NewProgram(m)
	_, err := g.Run()
	if err != nil {
		fmt.Println("There was a error.")
	}
}
