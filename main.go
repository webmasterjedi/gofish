package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

func main() {
	m := model{}
	m.spring = harmonica.NewSpring(harmonica.FPS(60), 4.0, 1.4)
	m.bobberTarget = 1
	g := tea.NewProgram(m)
	_, err := g.Run()
	if err != nil {
		fmt.Println("There was a error.")
	}
}
