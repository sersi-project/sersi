package build

import (
	"fmt"
	"os"

	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/scaffold"
	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/sersi-project/sersi/pkg"

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
	Long:  `Genrate Scaffold Application fromn a YAML file`,
	Run:   RunBuild,
}

func init() {
	BuildCmd.Flags().StringP("file", "f", "", "File Path")
}

func RunBuild(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	filePath, _ := cmd.Flags().GetString("file")
	fmt.Printf(" %s Creating a new project using %s:\n", common.OperationLabel, filePath)

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

	fmt.Printf("\n %s Generating frontend...\n", common.OperationLabel)
	err = frontend.Generate()
	if err != nil {
		fmt.Println("Error generating frontend:", err)
		os.Exit(1)
	}
	fmt.Printf("\n ├── %s Frontend generated successfully\n", common.SuccessLabel)

	fmt.Printf("\n %s Generating backend...\n", common.OperationLabel)
	err = backend.Generate()
	if err != nil {
		fmt.Println("Error generating backend:", err)
		os.Exit(1)
	}
	fmt.Printf("\n ├── %s Backend generated successfully\n", common.SuccessLabel)

	fmt.Printf("\n %s scaffolded successfully!\n\nHappy Coding :)\n", common.SuccessLabel)
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
