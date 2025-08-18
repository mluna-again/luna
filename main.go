package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mluna-again/luna/luna"
)

type model struct {
	luna  luna.LunaModel
	termH int
	termW int
}

func (m model) Init() tea.Cmd {
	return m.luna.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termH = msg.Height
		m.termW = msg.Width

		if m.termW < 30 {
			m.luna.SetSize(luna.SMALL)
		} else if m.termW < 60 {
			m.luna.SetSize(luna.MEDIUM)
		} else {
			m.luna.SetSize(luna.LARGE)
		}
	}

	var cmd tea.Cmd
	m.luna, cmd = m.luna.Update(msg)

	return m, cmd
}

func (m model) View() string {
	ascii := lipgloss.Place(m.termW, m.termH, lipgloss.Center, lipgloss.Center, m.luna.View())
	return ascii
}

var initialAnimation string
var initialPet string
var initialVariant string

func main() {
	flag.StringVar(&initialAnimation, "animation", "idle", "initial animation, can be: idle, sleeping, attacking. default: idle")
	flag.StringVar(&initialPet, "pet", "cat", "initial pet. can be: cat, turtle, bunny. default: cat")
	flag.StringVar(&initialVariant, "variant", "black", "initial variant (available for: cat). can be: ragdoll, black. default: black.")
	flag.Parse()

	params := luna.NewLunaParams{
		Animation: luna.LunaAnimation(initialAnimation),
		Pet:       luna.LunaPet(initialPet),
		Size:      luna.LARGE,
		Variant:   luna.LunaVariant(initialVariant),
	}

	l, errs := luna.NewLuna(params)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		os.Exit(1)
	}
	m := model{
		luna: l,
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
