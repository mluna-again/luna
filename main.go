package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type lunaModel struct {
	greetings string
}

func newLuna() lunaModel {
	return lunaModel{
		greetings: "hello brotha",
	}
}

func (l lunaModel) Init() tea.Cmd {
	return nil
}

func (l lunaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		}
	}
	return l, nil
}

func (l lunaModel) View() string {
	return l.greetings
}

func main() {
	p := tea.NewProgram(newLuna())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
