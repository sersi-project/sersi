package build

import (
	"fmt"
	"os"

	"github.com/sersi-project/core/common"
	"github.com/sersi-project/core/core"
	"github.com/sersi-project/core/hooks"
	"github.com/sersi-project/core/tea/spinner"
	"github.com/sersi-project/core/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var buildStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#CD24CD")).Italic(true)

var Cmd = &cobra.Command{
	Use:   "build",
	Short: "Genrate Scaffold Application from a YAML file",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   Run,
}

func init() {
	Cmd.Flags().StringP("file", "f", "", "File Path")
}

func Run(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	filePath, _ := cmd.Flags().GetString("file")

	if filePath == "" {
		fmt.Println("Error: File path is required")
		os.Exit(1)
	}

	fileParser := core.NewFileParser(filePath)

	fileParserResult, err := fileParser.ExceuteMapping()
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

    hooks := hooks.InitHooks(fileParserResult.Name, fileParserResult.Hooks.PreHook, fileParserResult.Hooks.PostHook)
    err = hooks.RunPreHook()
    if err != nil {
        fmt.Println("Error running pre-hook:", err)
        os.Exit(1)
    }

    fmt.Printf("> %s Creating a new project using %s:\n", buildStyle.Render("Building..."), filePath)

	scaffold := core.NewScaffoldBuilder().
		ProjectName(fileParserResult.Name).
		Framework(fileParserResult.Scaffold.Frontend.Framework).
		CSS(fileParserResult.Scaffold.Frontend.CSS).
		Language(fileParserResult.Scaffold.Frontend.Language).
		Build()

	loading := tea.NewProgram(spinner.InitialSpinnerModel(utils.GetProjectPath(fileParserResult.Name), scaffold))
	_, err = loading.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

    err = hooks.RunPostHook()
    if err != nil {
        fmt.Println("Error running post-hook:", err)
        os.Exit(1)
    }
}
