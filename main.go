package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type GameState int

const (
	StateIdle    GameState = iota // 0
	StateCasting                  // 1
	StateWaiting                  // 2
	StateReeling                  // 3
	StateCaught                   // 4
)

type model struct {
	state    GameState
	castLine int // was noted that I will be using this later, most likely as a flag to trigger an animation aka update the view.
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == " " {
			m.state = StateCasting
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case StateIdle:
		return "🎣 Fish Game\n\nPress SPACE to cast\nPress q to quit"
	case StateCasting:
		return "🎣 Casting your line out!"
	case StateWaiting:
		return "🎣 Patience! Wait for it. Light that fish whistle dawg" // maybe array of phrases and pick randomly here
	case StateReeling:
		return "🎣 Fish on! Don't let it get away!" // I assume we need to trigger the mini game event from here. This probably triggers something in the View method?
	case StateCaught:
		return "🎣 Fish caught! LFG!!"
	default:
		return "🎣 Fish Game\n\nPress SPACE to cast\nPress q to quit"

	}
}

func main() {
	g := tea.NewProgram(model{})
	_, err := g.Run()
	if err != nil {
		fmt.Println("There was a error.")
	}
}
