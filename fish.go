package main

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	fishBiteMsg struct{}
	Fish        struct {
		Name   string
		Weight float64
		Rarity int
	}
)

var fishCatalog = map[string]Fish{
	"bass":     {Name: "Bass", Weight: 1.2, Rarity: 1},
	"trout":    {Name: "Trout", Weight: 0.8, Rarity: 1},
	"salmon":   {Name: "Salmon", Weight: 3.5, Rarity: 2},
	"pike":     {Name: "Pike", Weight: 5.0, Rarity: 2},
	"sturgeon": {Name: "Sturgeon", Weight: 12.0, Rarity: 3},
}

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
