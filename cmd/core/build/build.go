package build

import (
	"fmt"
	"os"

	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/scaffold"
	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/sersi-project/sersi/internal/tui/spinner"
	"github.com/sersi-project/sersi/pkg"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	buildStyle = lipgloss.NewStyle().Italic(true)
	greenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#22CD24")).Italic(true)
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Genrate Scaffold Application from a YAML file",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   RunBuild,
}

func init() {
	BuildCmd.Flags().StringP("file", "f", "", "File Path")
}

func RunBuild(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	filePath, _ := cmd.Flags().GetString("file")
	fmt.Printf("◉ %s Creating a new project using %s:\n", buildStyle.Render("Building..."), filePath)

	fileParser := pkg.NewFileParser(filePath)

	fileParserResult, err := fileParser.ExceuteMapping()
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	frontend, err := buildFrontend(fileParserResult)
	if err != nil {
		fmt.Println("Error building frontend:", err)
		os.Exit(1)
	}

	backend, err := buildBackend(fileParserResult)
	if err != nil {
		fmt.Println("Error building backend:", err)
		os.Exit(1)
	}

	// devops, err := buildDevops(fileParserResult)
	// if err != nil {
	// 	fmt.Println("Error building devops:", err)
	// 	os.Exit(1)
	// }

	loading := tea.NewProgram(spinner.InitialSpinnerModel(pkg.GetProjectPath(fileParserResult.Name), "frontend", frontend))
	_, err = loading.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	loading = tea.NewProgram(spinner.InitialSpinnerModel(pkg.GetProjectPath(fileParserResult.Name), "backend", backend))
	_, err = loading.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Printf("◉ %s scaffolded successfully!\n\nHappy Coding :)\n", greenStyle.Render(fileParserResult.Name))
}

func buildFrontend(fileParserResult *pkg.Config) (scaffold.Scaffold, error) {
	f := frontend.NewFrontendBuilder().
		ProjectName(fileParserResult.Name).
		Framework(fileParserResult.Scaffold.Frontend.Framework).
		CSS(fileParserResult.Scaffold.Frontend.CSS).
		Language(fileParserResult.Scaffold.Frontend.Language).
		Monorepo(fileParserResult.Structure == "monorepo" || fileParserResult.Structure == "mono").
		Polyrepos(fileParserResult.Structure == "polyrepos" || fileParserResult.Structure == "poly").
		Build()
	return f, nil
}

func buildBackend(fileParserResult *pkg.Config) (scaffold.Scaffold, error) {
	b := backend.NewBackendBuilder().
		ProjectName(fileParserResult.Name).
		Language(fileParserResult.Scaffold.Backend.Language).
		Framework(fileParserResult.Scaffold.Backend.Framework).
		Database(fileParserResult.Scaffold.Backend.Database).
		Monorepo(fileParserResult.Structure == "monorepo" || fileParserResult.Structure == "mono").
		Polyrepos(fileParserResult.Structure == "polyrepos" || fileParserResult.Structure == "poly").
		Build()
	return b, nil
}
