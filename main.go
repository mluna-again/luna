package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mluna-again/luna/luna"
)

type model struct {
	luna luna.LunaModel
}

func (m model) Init() tea.Cmd {
	return m.luna.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.luna, cmd = m.luna.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return m.luna.View()
}

var initialAnimation string
var initialPet string

func main() {
	flag.StringVar(&initialAnimation, "animation", "idle", "initial animation, can be: idle, sleeping, attacking. default: idle")
	flag.StringVar(&initialPet, "pet", "cat", "initial pet. can be: cat, turtle, bunny. default: cat")
	flag.Parse()

	params := luna.NewLunaParams{
		Animation: initialAnimation,
		Pet:       initialPet,
	}
	errs := params.Validate()
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		os.Exit(1)
	}

	m := model{
		luna.NewLuna(params),
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
