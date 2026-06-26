package main

import (
	"fmt"
	"math/rand"
	"time"

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

type Fish struct {
	Name   string
	Weight float64
	Rarity int
}

var fishCatalog = map[string]Fish{
	"bass":     {Name: "Bass", Weight: 1.2, Rarity: 1},
	"trout":    {Name: "Trout", Weight: 0.8, Rarity: 1},
	"salmon":   {Name: "Salmon", Weight: 3.5, Rarity: 2},
	"pike":     {Name: "Pike", Weight: 5.0, Rarity: 2},
	"sturgeon": {Name: "Sturgeon", Weight: 12.0, Rarity: 3},
}

type model struct {
	state      GameState
	castLine   int // was noted that I will be using this later, most likely as a flag to trigger an animation aka update the view.
	caughtFish Fish
}

type fishBiteMsg struct{}

func randomFish() Fish {
	keys := make([]string, 0, len(fishCatalog))
	for k := range fishCatalog {
		keys = append(keys, k)
	}
	return fishCatalog[keys[rand.Intn(len(keys))]]
}

func waitForBite() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(2+rand.Intn(5)) * time.Second
		time.Sleep(duration)
		return fishBiteMsg{}
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == " " {
			switch m.state {
			case StateIdle: // idle, lets cast a line out
				m.state = StateWaiting
				return m, waitForBite()
			case StateReeling: // reeling in, get the fish
				m.caughtFish = randomFish()
				m.state = StateCaught
				return m, nil
			case StateCaught:
				m.state = StateIdle
				return m, nil
			}
		}
	case fishBiteMsg:
		m.state = StateReeling
		return m, nil
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
		return fmt.Sprintf("Caught a %s! %.1flb", m.caughtFish.Name, m.caughtFish.Weight)
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
