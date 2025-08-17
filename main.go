package main

// TODO remove all hardcoded and magical values

import (
	"errors"
	"fmt"
	"os"
	"time"

	"flag"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type animationTickMsg struct{}

func (l lunaModel) animationTick() tea.Cmd {
	shortAnimation := l.activeAnimationCount < 5
	delay := (time.Second * 2) / time.Duration(l.activeAnimationCount)
	if shortAnimation {
		delay = (time.Second * 3) / time.Duration(l.activeAnimationCount)
	}

	if l.activePet == "amogus" {
		delay = time.Millisecond * 40
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
	showHelp             bool
	err                  error
	termH                int
	termW                int
}

type newLunaParams struct {
	animation string
	pet       string
}

func (p newLunaParams) validate() []error {
	if p.animation == "" {
		p.animation = "idle"
	}

	if p.pet == "" {
		p.animation = "cat"
	}

	errs := []error{}
	if !memberOf([]string{"cat", "turtle", "bunny"}, p.pet) {
		errs = append(errs, errors.New("invalid pet"))
	}

	if !memberOf([]string{"idle", "sleeping", "attacking"}, p.animation) {
		errs = append(errs, errors.New("invalid animation"))
	}

	return errs
}

func newLuna(p newLunaParams) lunaModel {
	pets := getPets()

	return lunaModel{
		greetings:            "",
		pets:                 pets,
		activePet:            p.pet,
		activeAnimation:      p.animation,
		activeFrame:          0,
		activeAnimationCount: len(pets["cat"]["idle"]),
		showHelp:             false,
	}
}

func (l lunaModel) Init() tea.Cmd {
	return tea.Batch(tea.HideCursor, l.animationTick())
}

func (l lunaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		l.termH = msg.Height
		l.termW = msg.Width
		return l, nil

	case animationTickMsg:
		animation := l.pets[l.activePet][l.activeAnimation]
		if len(animation) == l.activeFrame+1 {
			l.activeFrame = 0
		} else {
			l.activeFrame++
		}

		return l, l.animationTick()

	case tea.KeyMsg:
		switch msg.String() {
		case "0":
			l.activeAnimation = "idle"
			l.activeFrame = 0
			l.activePet = "amogus"
			l.activeAnimationCount = len(l.pets[l.activePet][l.activeAnimation])
			return l, nil

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

		case "tab":
			l.showHelp = !l.showHelp
			return l, nil

		case "ctrl+c", "q":
			return l, tea.Batch(tea.ShowCursor, tea.Quit)
		}
	}

	return l, nil
}

func (l lunaModel) View() string {
	if l.err != nil {
		return fmt.Sprintf("Oh no: %s", l.err.Error())
	}

	ascii := l.pets[l.activePet][l.activeAnimation][l.activeFrame]
	ascii = lipgloss.Place(l.termW, l.termH, lipgloss.Center, lipgloss.Center, ascii)

	footer := lipgloss.JoinHorizontal(lipgloss.Center, "i - idle ", "s - sleeping ", "a - attacking")
	if !l.showHelp {
		return ascii
	}

	return ascii + footer
}

var initialAnimation string
var initialPet string

func main() {
	flag.StringVar(&initialAnimation, "animation", "idle", "initial animation, can be: idle, sleeping, attacking. default: idle")
	flag.StringVar(&initialPet, "pet", "cat", "initial pet. can be: cat, turtle, bunny. default: cat")
	flag.Parse()

	params := newLunaParams{
		animation: initialAnimation,
		pet:       initialPet,
	}
	errs := params.validate()
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		os.Exit(1)
	}

	p := tea.NewProgram(newLuna(params), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
