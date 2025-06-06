package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor    lipgloss.Color
	InputField     lipgloss.Style
	Header         lipgloss.Style
	Error          lipgloss.Style
	SelectionTitle lipgloss.Style
	Selection      lipgloss.Style
	Result         lipgloss.Style
	Cancel         lipgloss.Style
	StepNumber     lipgloss.Style
	StepTotal      lipgloss.Style
}

func DefaultStyles() *Styles {
	return &Styles{
		BorderColor:    lipgloss.Color("#36E6E6"),
		InputField:     lipgloss.NewStyle().BorderForeground(lipgloss.Color("#ffffff")).BorderStyle(lipgloss.ASCIIBorder()).Padding(1),
		Header:         lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#cd24cd")).Foreground(lipgloss.Color("#ffffff")),
		Error:          lipgloss.NewStyle().Background(lipgloss.Color("#EE4B2B")).Bold(true).Foreground(lipgloss.Color("#ffffff")),
		SelectionTitle: lipgloss.NewStyle().Foreground(lipgloss.Color("#36E6E6")).Italic(true),
		Selection:      lipgloss.NewStyle().Bold(true),
		Result:         lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")).Bold(true),
		Cancel:         lipgloss.NewStyle().Foreground(lipgloss.Color("#EE4B2B")).Bold(true),
		StepNumber:     lipgloss.NewStyle().Foreground(lipgloss.Color("#36E6E6")).Italic(true),
		StepTotal:      lipgloss.NewStyle().Foreground(lipgloss.Color("#36E6E6")).Italic(true),
	}
}
