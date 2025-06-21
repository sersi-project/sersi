package common

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const Logo = `
  ___  ___  ____   ___   ___  
 (  _() __(/  _ \ (  _( )_ _( 
 _) \ | _) )  ' / _) \  _| |_ 
)____))___(|_()_\)____))_____(
                             
`

var (
	Version    = "development"
	Commit     = "environment"
	link       = "https://sersi.dev"
	logoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#36E6E6"))
	infoStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#808080"))
	ErrorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
)

func PrintLogo() {
	fmt.Printf("%s\n", logoStyle.Render(Logo))
	fmt.Printf("%s %s\n", Version, infoStyle.Render(Commit))
	fmt.Printf("%s\n\n\n", infoStyle.Render(link))
}

var (
	SuccessLabel   = lipgloss.NewStyle().Bold(true).PaddingLeft(1).PaddingRight(1).Background(lipgloss.Color("#22CD24")).Render("✓ SUCCESS")
	ErrorLabel     = lipgloss.NewStyle().Bold(true).PaddingLeft(1).PaddingRight(1).Background(lipgloss.Color("#FF0000")).Render("✘ ERROR")
	OperationLabel = lipgloss.NewStyle().Bold(true).PaddingLeft(1).PaddingRight(1).Background(lipgloss.Color("#FFFF00")).Render("◉ OPERATION")
)
