package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GameState int

const (
	StateIdle GameState = iota
	StateCasting
	StateWaiting
	StateReeling
	StateCaught
	StateLog
)

type model struct {
	state       GameState
	castLine    int
	caughtFish  Fish
	inventory   []Fish
	totalWeight float64
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// quit game
		if msg.String() == "q" {
			return m, tea.Quit
		}

		// show/hide log
		if msg.String() == "l" {
			m.state = StateLog
			return m, nil
		}

		if msg.String() == " " {
			switch m.state {
			case StateIdle: // idle, lets cast a line out
				m.state = StateWaiting
				return m, waitForBite()

			case StateReeling: // reeling in, get the fish
				m.caughtFish = randomFish()
				m.inventory = append(m.inventory, m.caughtFish)
				m.totalWeight += m.caughtFish.Weight
				m.state = StateCaught
				return m, nil

			case StateCaught, StateLog:
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
