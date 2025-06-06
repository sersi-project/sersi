package spinner

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/sersi-project/sersi/internal/scaffold"
	"github.com/sersi-project/sersi/internal/tui/styles"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var buildStyle = lipgloss.NewStyle().Italic(true)

type errMsg error

type CompletedMsg struct {
	Err error
}

type SpinnerModel struct {
	outputType  string
	spinner     spinner.Model
	projectPath string
	quitting    bool
	err         error
	scaffold    scaffold.Scaffold
	styles      *styles.Styles
}

func InitialSpinnerModel(projectPath, outputType string, scaffold scaffold.Scaffold) SpinnerModel {
	styles := styles.DefaultStyles()
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return SpinnerModel{outputType: outputType, spinner: s, projectPath: projectPath, scaffold: scaffold, styles: styles}
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
		return fmt.Sprintf("Error: %s\n\n(press ctrl+c to exit)\n\nreport this bug at sersi.dev\n", m.err.Error())
	}
	str := fmt.Sprintf("%s Generating project...\n", m.spinner.View())
	if m.quitting {
		return fmt.Sprintf("◉ %s\n", buildStyle.Render("Created "+m.outputType+"..."))
	}
	return str
}

func runScaffold(scaffold scaffold.Scaffold) tea.Cmd {
	if err := scaffold.Generate(); err != nil {
		return func() tea.Msg {
			return CompletedMsg{Err: err}
		}
	}
	return func() tea.Msg {
		return CompletedMsg{Err: scaffold.Generate()}
	}
}
