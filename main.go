package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/mluna-again/luna/luna"
)

type model struct {
	luna  luna.LunaModel
	termH int
	termW int
	bg    *color.Color
}

func (m model) Init() tea.Cmd {
	return m.luna.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termH = msg.Height
		m.termW = msg.Width
	}

	var cmd tea.Cmd
	m.luna, cmd = m.luna.Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	ascii := lipgloss.Place(m.termW, m.termH, lipgloss.Center, lipgloss.Center, m.luna.View().Content)
	var bg color.Color = color.Transparent
	if m.bg != nil {
		bg = *m.bg
	}
	return tea.View{
		Content:         ascii,
		AltScreen:       true,
		BackgroundColor: bg,
	}
}

var initialAnimation string
var initialPet string
var initialVariant string
var name string
var fill string

func main() {
	flag.StringVar(&initialAnimation, "animation", "idle", "initial animation, can be: idle, sleeping, attacking. default: idle")
	flag.StringVar(&initialPet, "pet", "cat", "initial pet. can be: cat, turtle, bunny. default: cat")
	flag.StringVar(&initialVariant, "variant", "default", "initial variant (available for: cat). can be: ragdoll, black. default: black.")
	flag.StringVar(&name, "name", "Luna", "pet's name")
	flag.StringVar(&fill, "fill", "", "Fill background")
	flag.Parse()

	params := luna.NewLunaParams{
		Animation: luna.LunaAnimation(initialAnimation),
		Pet:       luna.LunaPet(initialPet),
		Size:      luna.LARGE,
		Variant:   luna.LunaVariant(initialVariant),
		Name:      name,
	}

	l, errs := luna.NewLuna(params)
	l.SetAutoresize(true)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		os.Exit(1)
	}
	var c *color.Color
	if fill != "" {
		col := lipgloss.Color(fill)
		c = &col
	}
	m := model{
		luna: l,
		bg:   c,
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
