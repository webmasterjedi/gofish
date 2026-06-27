package main

import (
	"fmt"
	"strings"
)

func getFishLog(m model) string {
	var sb strings.Builder
	for _, fish := range m.inventory {
		fmt.Fprintf(&sb, "- %s | %.1flb | Rarity: %d \n", fish.Name, fish.Weight, fish.Rarity)
	}
	fmt.Fprintf(&sb, "Total Fish: %d\nTotal Weight: %.1flb", len(m.inventory), m.totalWeight)
	return sb.String()
}

func (m model) View() string {
	switch m.state {
	case StateIdle:
		return "🎣 Fish Game\n\nPress SPACE to cast\nPress l for log\nPress q to quit"
	case StateCasting:
		return "🎣 Casting your line out!"
	case StateWaiting:
		return "🎣 Patience! Wait for it. Light that fish whistle dawg" // maybe array of phrases and pick randomly here
	case StateReeling:
		return "🎣 Fish on! Don't let it get away!" // I assume we need to trigger the mini game event from here. This probably triggers something in the View method?
	case StateCaught:
		return fmt.Sprintf("Caught a %s! %.1flb\n\nPress SPACE to continue\nPress l for log\nPress q to quit", m.caughtFish.Name, m.caughtFish.Weight)
	case StateLog:
		return getFishLog(m)

	default:
		return "🎣 Fish Game\n\nPress SPACE to cast\nPress q to quit"

	}
}
