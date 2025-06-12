package common

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const Logo = `SERSI - Skip the Setup.`

var (
	link       = "https://sersi.dev"
	version    = "0.0.0-alpha"
	logoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#36E6E6"))
	infoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#808080"))
	ErrorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
)

func PrintLogo() {
	fmt.Printf("%s\n", logoStyle.Render(Logo))
	fmt.Printf("%s\n", version)
	fmt.Printf("%s\n\n\n", infoStyle.Render(link))
}
