package common

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const Logo = `SERSI - Skip the Setup.`

var (
	link       = "https://sersi.dev"
	version    = "1.0.0"
	logoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#36E6E6"))
	infoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#808080"))
	ErrorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
)

func PrintLogo() {
	fmt.Printf("%s\n", logoStyle.Render(Logo))
	fmt.Printf("%s\n", fmt.Sprintf("v%s    %s\n", version, infoStyle.Render(link)))
}
