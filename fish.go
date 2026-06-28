package main

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	castMsg     struct{}
	fishBiteMsg struct{}
	Fish        struct {
		Name   string
		Weight float64
		Rarity int
	}
)

var fishCatalog = map[string]Fish{
	"lmb":     {Name: "Largemouth Bass", Weight: 4.02, Rarity: 2},
	"smb":     {Name: "Smallmouth Bass", Weight: 3.02, Rarity: 2},
	"spot":    {Name: "Spotted Bass", Weight: 1.8, Rarity: 1},
	"wcrap":   {Name: "White Crappie", Weight: 1.66, Rarity: 2},
	"bcrap":   {Name: "Black Crappie", Weight: 1.16, Rarity: 1},
	"gill":    {Name: "Bluegill Sunfish", Weight: 0.23, Rarity: 1},
	"flatcat": {Name: "Flathead Catfish", Weight: 66.6, Rarity: 3},
	"carp":    {Name: "Carp", Weight: 33.26, Rarity: 2},
}

func randomFish() Fish {
	keys := make([]string, 0, len(fishCatalog))
	for k := range fishCatalog {
		keys = append(keys, k)
	}
	return fishCatalog[keys[rand.Intn(len(keys))]]
}

func castingLine() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		return castMsg{}
	}
}

func waitForBite() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(2+rand.Intn(5)) * time.Second
		time.Sleep(duration)
		return fishBiteMsg{}
	}
}
