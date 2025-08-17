package luna

// TODO remove all hardcoded and magical values

import (
	"errors"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type animationTickMsg struct{}

func (l LunaModel) animationTick() tea.Cmd {
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

type LunaModel struct {
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
	keysDisabled         bool
}

type NewLunaParams struct {
	Animation string
	Pet       string
}

func (p NewLunaParams) Validate() []error {
	if p.Animation == "" {
		p.Animation = "idle"
	}

	if p.Pet == "" {
		p.Animation = "cat"
	}

	errs := []error{}
	if !memberOf([]string{"cat", "turtle", "bunny"}, p.Pet) {
		errs = append(errs, errors.New("invalid pet"))
	}

	if !memberOf([]string{"idle", "sleeping", "attacking"}, p.Animation) {
		errs = append(errs, errors.New("invalid animation"))
	}

	return errs
}

func NewLuna(p NewLunaParams) LunaModel {
	pets := getPets()

	return LunaModel{
		greetings:            "",
		pets:                 pets,
		activePet:            p.Pet,
		activeAnimation:      p.Animation,
		activeFrame:          0,
		activeAnimationCount: len(pets["cat"]["idle"]),
		showHelp:             false,
	}
}

func (l LunaModel) Init() tea.Cmd {
	return tea.Batch(tea.HideCursor, l.animationTick())
}

func (l LunaModel) Update(msg tea.Msg) (LunaModel, tea.Cmd) {
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
		if l.keysDisabled {
			break
		}
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

func (l LunaModel) View() string {
	if l.err != nil {
		return fmt.Sprintf("Oh no: %s", l.err.Error())
	}

	ascii := l.pets[l.activePet][l.activeAnimation][l.activeFrame]

	footer := lipgloss.JoinHorizontal(lipgloss.Center, "i - idle ", "s - sleeping ", "a - attacking")
	if !l.showHelp {
		return ascii
	}

	return ascii + footer
}

func (l *LunaModel) DisableKeys() {
	l.keysDisabled = true
}

func (l *LunaModel) EnableKeys() {
	l.keysDisabled = false
}
