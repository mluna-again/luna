package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type lunaModel struct {
	greetings string
	pets      asciiPets
	activePet []string
	frame     int
	err       error
}

func newLuna() lunaModel {
	pets := getPets()

	return lunaModel{
		greetings: "hello brotha",
		pets:      pets,
		activePet: pets.cat.idle,
		frame:     0,
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
	if l.err != nil {
		return fmt.Sprintf("Oh no: %s", l.err.Error())
	}

	ascii := l.activePet[l.frame]

	return ascii
}

func main() {
	p := tea.NewProgram(newLuna())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
