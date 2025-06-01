package common

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const Logo = `
 ____  _____ ____  ____ ___ 
/ ___|| ____|  _ \/ ___|_ _|
\___ \|  _| | |_) \___ \| | 
 ___) | |___|  _ < ___) | | 
|____/|_____|_| \_|____|___|

SERSI - Skip the boilerplate.`

var(
	link = "https://sersi.dev"
	version = "1.0.0"
	logoStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#36E6E6"))
	infoStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#808080"))
	ErrorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
)

func PrintLogo() {
  fmt.Printf("%s\n", logoStyle.Render(Logo))
	fmt.Printf("%s\n", fmt.Sprintf("v%s    %s\n\n", version, infoStyle.Render(link)))
}