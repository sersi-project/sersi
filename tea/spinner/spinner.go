package spinner

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"sersi/core"
	"sersi/tea/styles"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type CompletedMsg struct {
	Err error
}

type SpinnerModel struct {
	spinner     spinner.Model
	projectPath string
	quitting    bool
	err         error
	scaffold    *core.Scaffold
	styles      *styles.Styles
}

func InitialSpinnerModel(projectPath string, scaffold *core.Scaffold) SpinnerModel {
	styles := styles.DefaultStyles()
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return SpinnerModel{spinner: s, projectPath: projectPath, scaffold: scaffold, styles: styles}
}

func (m SpinnerModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, runScaffold(m.scaffold))
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case CompletedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}
		m.quitting = true
		return m, tea.Quit
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m SpinnerModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Generating project at %s...\n\n", m.spinner.View(), m.projectPath)
	if m.quitting {
		return "\n       Project created at " + m.projectPath + "\n\n"
	}
	return str
}

func runScaffold(scaffold *core.Scaffold) tea.Cmd {
	return func() tea.Msg {
		return CompletedMsg{Err: scaffold.Execute()}
	}
}
