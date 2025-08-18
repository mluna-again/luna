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
type LunaVariant string

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

const (
	CAT_BLACK       LunaVariant = "default"
	CAT_RAGDOLL                 = "ragdoll"
	DEFAULT_VARIANT             = "default"
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
	name                 string
	pets                 asciiPets
	activePet            LunaPet
	activeAnimation      LunaAnimation
	activeVariant        LunaVariant
	activeAnimationCount int
	activeFrame          int
	showHelp             bool
	err                  error
	termH                int
	termW                int
	keysDisabled         bool
	size                 LunaSize
	displayName          bool
}

type NewLunaParams struct {
	Animation LunaAnimation
	Pet       LunaPet
	Size      LunaSize
	Variant   LunaVariant
	Name      string
}

func (p NewLunaParams) Validate() (NewLunaParams, []error) {
	if p.Animation == "" {
		p.Animation = IDLE
	}

	if p.Pet == "" {
		p.Pet = CAT
	}

	if p.Size == "" {
		p.Size = LARGE
	}

	if p.Name == "" {
		p.Name = "Luna"
	}

	if p.Variant == "" {
		p.Variant = getDefaultVariant(p.Pet)
	}
	p.Variant = translateVariant(p.Variant)

	errs := []error{}

	if !memberOf[string]([]string{string(IDLE), string(SLEEPING), string(ATTACKING)}, string(p.Animation)) {
		errs = append(errs, errors.New("invalid animation"))
	}

	if !memberOf[string]([]string{string(CAT), string(BUNNY), string(TURTLE)}, string(p.Pet)) {
		errs = append(errs, errors.New("invalid pet"))
	}

	availableVars := availableVariants(p.Pet)
	if !memberOf[LunaVariant](availableVars, p.Variant) {
		errs = append(errs, errors.New("invalid variant for pet"))
	}

	return p, errs
}

func NewLuna(p NewLunaParams) (LunaModel, []error) {
	pets := getPets()

	params, errs := p.Validate()
	if len(errs) > 0 {
		return LunaModel{}, errs
	}

	return LunaModel{
		greetings:            "",
		pets:                 pets,
		name:                 params.Name,
		activePet:            params.Pet,
		activeAnimation:      params.Animation,
		activeVariant:        params.Variant,
		size:                 params.Size,
		activeFrame:          0,
		activeAnimationCount: len(pets[params.Pet][params.Variant][params.Size][params.Animation]),
		showHelp:             false,
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

		case ">":
			l.SetVariant(nextVariant(l.activeVariant, l.activePet))
			return l, nil

		case " ":
			l.displayName = !l.displayName

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
	if l.displayName {
		name := nameStyle.Render(l.name)
		ascii = lipgloss.JoinVertical(lipgloss.Center, ascii, name)
	}

	footer := lipgloss.JoinHorizontal(lipgloss.Center, "i - idle ", "s - sleeping ", "a - attacking")
	footer = lipgloss.JoinVertical(lipgloss.Center, footer, lipgloss.JoinHorizontal(lipgloss.Center, "> - switch variant"))
	footer = lipgloss.JoinVertical(lipgloss.Center, footer, lipgloss.JoinHorizontal(lipgloss.Center, "1 - cat  2 - bunny  3 - turtle"))
	if !l.showHelp {
		return ascii
	}

	return ascii + "\n" + footer
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
	l.activeAnimationCount = l.getActiveFrameCount()
}

func (l *LunaModel) SetAnimation(animation LunaAnimation) {
	l.activeAnimation = animation
	l.activeFrame = 0
	l.activeAnimationCount = l.getActiveFrameCount()
}

func (l *LunaModel) SetVariant(variant LunaVariant) {
	l.activeVariant = getSelectedVariant(l.activePet, variant)
	l.activeFrame = 0
	l.activeAnimationCount = l.getActiveFrameCount()
}

func (l LunaModel) GetAnimation() LunaAnimation {
	return l.activeAnimation
}

func (l LunaModel) getActivePet() string {
	// yeah this is definetely not gonna crash
	return l.pets[l.activePet][getSelectedVariant(l.activePet, l.activeVariant)][l.size][l.activeAnimation][l.activeFrame]
}

func (l LunaModel) getActiveFrameCount() int {
	return len(l.pets[l.activePet][getSelectedVariant(l.activePet, l.activeVariant)][l.size][l.activeAnimation])
}

func (l *LunaModel) SetSize(size LunaSize) {
	l.size = size
}

func (l *LunaModel) ShowName() {
	l.displayName = true
}

func (l *LunaModel) HideName() {
	l.displayName = false
}

func (l *LunaModel) SetName(name string) {
	l.name = name
}
