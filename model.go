package main

import (
	"math"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
)

const Title = "Crowley's Ridge Fishing Simulator"

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
	state        GameState
	waitPhrase   string
	caughtFish   Fish
	inventory    []Fish
	totalWeight  float64
	frame        int
	spring       harmonica.Spring
	bobberPos    float64
	bobberVel    float64
	bobberTarget float64
	catchPos     float64
	catchVel     float64
	catchTarget  float64
}
type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd { return tick() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// quit game
		if msg.String() == "q" {
			return m, tea.Quit
		}

		// show/hide log
		if msg.String() == "l" && (m.state == StateIdle || m.state == StateCaught) {
			m.state = StateLog
			return m, nil
		}

		if msg.String() == " " {
			switch m.state {
			case StateIdle: // idle, lets cast a line out
				m.state = StateCasting
				return m, castingLine()

			case StateReeling: // reeling in, get the fish
				m.caughtFish = randomFish()
				m.inventory = append(m.inventory, m.caughtFish)
				m.totalWeight += m.caughtFish.Weight
				m.state = StateCaught
				return m, nil

			case StateCaught: // fish caught
				m.state = StateCasting
				return m, castingLine()

			case StateLog: // log open
				m.state = StateIdle
				return m, nil
			}
		}

	case castMsg:
		m.state = StateWaiting
		m.waitPhrase = randomWaitPhrase()
		return m, waitForBite()

	case fishBiteMsg:
		m.state = StateReeling
		return m, nil

	case tickMsg:
		m.frame++
		switch m.state {
		case StateWaiting:
			// advanced bobberPos here
			m.bobberPos, m.bobberVel = m.spring.Update(m.bobberPos, m.bobberVel, m.bobberTarget)
			// flip direction
			if math.Abs(m.bobberPos-m.bobberTarget) < 0.05 {
				if m.bobberTarget == 1.0 {
					m.bobberTarget = 0.0
				} else {
					m.bobberTarget = 1.0
				}
			}
			return m, tick()
		case StateReeling:
			// advance catch spring
			return m, tick()
		}
		return m, tick()
	}
	return m, tick()
}
