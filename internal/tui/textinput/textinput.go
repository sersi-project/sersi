package textinput

import (
	"fmt"
	"regexp"

	"github.com/sersi-project/sersi/internal/tui/styles"
	"github.com/sersi-project/sersi/pkg"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	defaultWidth = 20
	dividerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type (
	errMsg error
)

type Model struct {
	textInput  textinput.Model
	header     string
	Value      string
	err        error
	Quitting   *bool
	styles     *styles.Styles
	totalSteps int
	step       int
}

func InitialModel(totalSteps, step int, header, placeholder string) Model {
	quitting := false
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 30
	ti.Validate = validateString

	return Model{
		textInput:  ti,
		header:     header,
		Value:      "",
		err:        nil,
		Quitting:   &quitting,
		styles:     styles.DefaultStyles(),
		totalSteps: totalSteps,
		step:       step,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if *m.Quitting {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			*m.Quitting = true
			return m, tea.Quit
		case "enter":
			if err := validateString(m.textInput.Value()); err != nil {
				m.err = err
				return m, tea.Quit
			}
			m.Value = m.textInput.Value()
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		*m.Quitting = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.Value != "" {
		return ""
	}
	if m.err != nil {
		return fmt.Sprintf("\n%s\n", m.styles.Error.Render(m.err.Error()))
	}

	stepTitle := fmt.Sprintf("\n%s [%s/%s]", lipgloss.NewStyle().Bold(true).Render(m.header), m.styles.StepNumber.Render(fmt.Sprintf("%d", m.step)), m.styles.StepTotal.Render(fmt.Sprintf("%d", m.totalSteps)))
	return lipgloss.JoinVertical(lipgloss.Left, stepTitle, m.textInput.View(), "\n\n(esc to quit)\n")
}

func validateString(s string) error {
	if len(s) < 3 {
		return fmt.Errorf("name is too short: %s", s)
	}
	matched, err := regexp.MatchString("^[a-zA-Z0-9_-]+$", s)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid project name: %s", s)
	}

	if pkg.FileExists(s) {
		return fmt.Errorf("project already exists: %s", s)
	}
	return nil
}
