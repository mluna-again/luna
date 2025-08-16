package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type animationTickMsg struct{}

func animationTick(frameCount int) tea.Cmd {
	shortAnimation := frameCount < 5
	delay := (time.Second * 2) / time.Duration(frameCount)
	if shortAnimation {
		delay = (time.Second * 3) / time.Duration(frameCount)
	}

	return tea.Tick(delay, func(t time.Time) tea.Msg {
		return animationTickMsg{}
	})
}

type lunaModel struct {
	greetings            string
	pets                 asciiPets
	activePet            string
	activeAnimation      string
	activeAnimationCount int
	activeFrame          int
	err                  error
}

func newLuna() lunaModel {
	pets := getPets()

	return lunaModel{
		greetings:            "",
		pets:                 pets,
		activePet:            "cat",
		activeAnimation:      "idle",
		activeFrame:          0,
		activeAnimationCount: len(pets["cat"]["idle"]),
	}
}

func (l lunaModel) Init() tea.Cmd {
	return animationTick(l.activeAnimationCount)
}

func (l lunaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case animationTickMsg:
		animation := l.pets[l.activePet][l.activeAnimation]
		if len(animation) == l.activeFrame+1 {
			l.activeFrame = 0
		} else {
			l.activeFrame++
		}

		return l, animationTick(l.activeAnimationCount)

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			l.activeAnimation = "idle"
			l.activeFrame = 0
			l.activePet = "cat"
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

		case "2":
			l.activeAnimation = "idle"
			l.activeFrame = 0
			l.activePet = "turtle"
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

		case "3":
			l.activeAnimation = "idle"
			l.activeFrame = 0
			l.activePet = "bunny"
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

		case "s":
			l.activeAnimation = "sleeping"
			l.activeFrame = 0
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

		case "i":
			l.activeAnimation = "idle"
			l.activeFrame = 0
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

		case "a":
			l.activeAnimation = "attacking"
			l.activeFrame = 0
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

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

	ascii := l.pets[l.activePet][l.activeAnimation][l.activeFrame]

	footer := lipgloss.JoinHorizontal(lipgloss.Center, "i - idle ", "s - sleeping ", "a - attacking")

	return ascii + footer
}

func main() {
	p := tea.NewProgram(newLuna())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
