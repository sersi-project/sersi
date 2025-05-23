package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
	Header      lipgloss.Style
	Error       lipgloss.Style
}

func DefaultStyles() *Styles {
	return &Styles{
		BorderColor: lipgloss.Color("36"),
		InputField:  lipgloss.NewStyle().BorderForeground(lipgloss.Color("36")).BorderStyle(lipgloss.NormalBorder()).Padding(1),
		Header:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("36")),
		Error:       lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	}
}
