package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	castControl    = "Press SPACE to cast"
	logControl     = "Press l for log"
	quitControl    = "Press q to quit"
	waitingPhrases = []string{
		"Hit this fish whistle dawg!",
		"Patience is key!",
		"Wait for it!",
	}
)

var titleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	Bold(true).
	Foreground(lipgloss.Color("#829B5E")).
	BorderBottom(true).
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("#141b29"))

var subStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	Foreground(lipgloss.Color("#7d96ff"))

var controlsTitleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	Bold(true).
	Foreground(lipgloss.Color("#546faa")).
	BorderBottom(true).
	BorderStyle(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("#141b29"))

var controlStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	Foreground(lipgloss.Color("#7aa2f7"))

var boxStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	Align(lipgloss.Center).
	Width(60).
	Height(10).
	Padding(2).
	Border(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("240"))

var totalRowStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#2e3d5d")).
	BorderTop(true).
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("#141b29")).
	Foreground(lipgloss.Color("#394d75"))

func rarityStyle(rarity int) lipgloss.Style {
	switch rarity {
	case 2:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Background(lipgloss.Color("#2e3d5d"))
	case 3:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Background(lipgloss.Color("#2e3d5d"))
	default:
		return lipgloss.NewStyle().Background(lipgloss.Color("#2e3d5d"))
	}
}

func rarityLabel(rarity int) string {
	baseStyle := rarityStyle(rarity)
	switch rarity {
	case 1:
		return baseStyle.Render("Common")
	case 2:
		return baseStyle.Render("Uncommon")
	case 3:
		return baseStyle.Render("Rare")
	default:
		return baseStyle.Render("Common")
	}
}

func getControls(m model) string {
	controls := controlsTitleStyle.Render("Controls:") + "\n" +
		controlStyle.Render(castControl) + "\n"

	if len(m.inventory) > 0 && m.state != StateLog {
		controls += controlStyle.Render(logControl) + "\n"
	}

	controls += controlStyle.Render(quitControl)
	return controls
}

func randomWaitPhrase() string {
	return waitingPhrases[rand.Intn(len(waitingPhrases))]
}

func getFishLog(m model) string {
	var sb strings.Builder
	var (
		colName   = lipgloss.NewStyle().Width(16).Bold(true).Background(lipgloss.Color("#2e3d5d"))
		colWeight = lipgloss.NewStyle().Width(10).PaddingRight(1).Align(lipgloss.Right).Background(lipgloss.Color("#2e3d5d"))
		colRarity = lipgloss.NewStyle().Width(10).Background(lipgloss.Color("#2e3d5d"))
	)
	sb.WriteString(titleStyle.Render("📕 Catch log"))
	sb.WriteByte('\n')
	sb.WriteByte('\n')
	for _, fish := range m.inventory {
		fmt.Fprintln(&sb, colName.Render(fish.Name)+
			colWeight.Render(fmt.Sprintf("%.1flb", fish.Weight))+
			colRarity.Render(rarityLabel(fish.Rarity)))
	}
	totalRow := subStyle.Bold(true).
		Render("Total Fish: ") +
		subStyle.Foreground(lipgloss.Color("#c3e88d")).
			Render(fmt.Sprintf("%d | ", len(m.inventory))) +
		subStyle.Bold(true).
			Render("Total Weight: ") +
		subStyle.Foreground(lipgloss.Color("#c3e88d")).
			Render(fmt.Sprintf("%.1flb", m.totalWeight))

	sb.WriteString(totalRowStyle.Render(totalRow))
	sb.WriteByte('\n')
	sb.WriteByte('\n')
	sb.WriteString(subStyle.Render(getControls(m)))
	return boxStyle.Render(sb.String())
}

func getIdle(m model) string {
	content := titleStyle.Render(Title) + "\n\n" +
		getControls(m)
	return boxStyle.Render(content)
}

func getCasting() string {
	content := subStyle.Render("Casting your line out!")
	return boxStyle.Render(content)
}

func getWaiting(m model) string {
	content := titleStyle.Render("Wait for a bite...") + "\n\n" +
		subStyle.Render(m.waitPhrase)
	return boxStyle.Render(content)
}

func getReeling() string {
	content := titleStyle.Render("Bing Bong! Fish On!") + "\n\n" +
		subStyle.Render("Press SPACE to set the hook!")
	return boxStyle.Render(content)
}

func getCaught(m model) string {
	content := titleStyle.Render("Fish caught!") + "\n\n" +
		rarityStyle(m.caughtFish.Rarity).Render(fmt.Sprintf("🐟 %s %.1flb", m.caughtFish.Name, m.caughtFish.Weight)) + "\n\n" +
		subStyle.Render(getControls(m))
	return boxStyle.Render(content)
}

func (m model) View() string {
	switch m.state {

	case StateIdle:
		return getIdle(m)

	case StateCasting:
		return getCasting()

	case StateWaiting:
		return getWaiting(m)

	case StateReeling:
		return getReeling()

	case StateCaught:
		return getCaught(m)

	case StateLog:
		return getFishLog(m)

	default:
		return getControls(m)

	}
}
