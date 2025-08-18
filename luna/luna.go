package luna

import (
	"errors"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LunaSize string
type LunaPet string
type LunaAnimation string

const (
	SMALL  LunaSize = "sm"
	MEDIUM          = "md"
	LARGE           = "xl"
)

const (
	CAT    LunaPet = "cat"
	BUNNY          = "bunny"
	TURTLE         = "turtle"
	AMOGUS         = "amogus"
)

const (
	IDLE      LunaAnimation = "idle"
	SLEEPING                = "sleeping"
	ATTACKING               = "attacking"
)

type animationTickMsg struct{}

func (l LunaModel) animationTick() tea.Cmd {
	shortAnimation := l.activeAnimationCount < 5
	delay := (time.Second * 2) / time.Duration(l.activeAnimationCount)
	if shortAnimation {
		delay = (time.Second * 3) / time.Duration(l.activeAnimationCount)
	}

	if l.activePet == AMOGUS {
		delay = time.Millisecond * 40
	}

	return tea.Tick(delay, func(t time.Time) tea.Msg {
		return animationTickMsg{}
	})
}

type LunaModel struct {
	greetings            string
	pets                 asciiPets
	activePet            LunaPet
	activeAnimation      LunaAnimation
	activeAnimationCount int
	activeFrame          int
	showHelp             bool
	err                  error
	termH                int
	termW                int
	keysDisabled         bool
	size                 LunaSize
}

type NewLunaParams struct {
	Animation LunaAnimation
	Pet       LunaPet
	Size      LunaSize
}

func (p NewLunaParams) Validate() []error {
	if p.Animation == "" {
		p.Animation = IDLE
	}

	if p.Pet == "" {
		p.Pet = CAT
	}

	if p.Size == "" {
		p.Size = LARGE
	}

	errs := []error{}

	if !memberOf[string]([]string{string(IDLE), string(SLEEPING), string(ATTACKING)}, string(p.Animation)) {
		errs = append(errs, errors.New("invalid animation"))
	}

	if !memberOf[string]([]string{string(CAT), string(BUNNY), string(TURTLE)}, string(p.Pet)) {
		errs = append(errs, errors.New("invalid pet"))
	}

	return errs
}

func NewLuna(p NewLunaParams) (LunaModel, []error) {
	pets := getPets()

	errs := p.Validate()
	if len(errs) > 0 {
		return LunaModel{}, errs
	}

	return LunaModel{
		greetings:            "",
		pets:                 pets,
		activePet:            p.Pet,
		activeAnimation:      p.Animation,
		activeFrame:          0,
		activeAnimationCount: len(pets[CAT][p.Size][IDLE]),
		showHelp:             false,
		size:                 p.Size,
	}, []error{}
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
		frameCount := l.getActiveFrameCount()
		if frameCount == l.activeFrame+1 {
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
			l.SetPet(AMOGUS)
			return l, nil

		case "1":
			l.SetPet(CAT)
			return l, nil

		case "2":
			l.SetPet(BUNNY)
			return l, nil

		case "3":
			l.SetPet(TURTLE)
			return l, nil

		case "s":
			l.SetAnimation(SLEEPING)
			return l, nil

		case "i":
			l.SetAnimation(IDLE)
			return l, nil

		case "a":
			l.SetAnimation(ATTACKING)
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

	ascii := l.getActivePet()

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

func (l *LunaModel) SetPet(name LunaPet) {
	l.activeFrame = 0
	l.activePet = name
	l.activeAnimationCount = len(l.pets[l.activePet][l.size][l.activeAnimation])
}

func (l *LunaModel) SetAnimation(animation LunaAnimation) {
	l.activeAnimation = animation
	l.activeFrame = 0
	l.activeAnimationCount = len(l.pets[l.activePet][l.size][l.activeAnimation])
}

func (l LunaModel) getActivePet() string {
	// yeah this is definetely not gonna crash
	return l.pets[l.activePet][l.size][l.activeAnimation][l.activeFrame]
}

func (l LunaModel) getActiveFrameCount() int {
	return len(l.pets[l.activePet][l.size][l.activeAnimation])
}

func (l *LunaModel) SetSize(size LunaSize) {
	l.size = size
}
