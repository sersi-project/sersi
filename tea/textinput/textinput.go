package textinput

import (
	"fmt"
	"regexp"

	"github.com/sersi-project/core/tea/styles"
	"github.com/sersi-project/core/utils"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type Model struct {
	textInput textinput.Model
	header    string
	Value     string
	err       error
	quitting  *bool
	styles    *styles.Styles
}

func InitialModel(header, placeholder string) Model {
	quitting := false
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 30
	ti.Validate = validateString

	return Model{
		textInput: ti,
		header:    header,
		Value:     "",
		err:       nil,
		quitting:  &quitting,
		styles:    styles.DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if *m.quitting {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			*m.quitting = true
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
		*m.quitting = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.Value != "" {
		return fmt.Sprintf("> %-15s: %s\n", m.styles.SelectionTitle.Render("Project Name"), m.styles.Selection.Render(m.Value))
	}
	if m.err != nil {
		return fmt.Sprintf("\n%s\n", m.styles.Error.Render(m.err.Error()))
	}
	return fmt.Sprintf("%s\n%s\n\n%s", m.styles.Header.Render(m.header), m.textInput.View(), "(esc to quit)" + "\n")
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

	if utils.FileExists(s) {
		return fmt.Errorf("project already exists: %s", s)
	}
	return nil
}