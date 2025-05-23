package build

import (
	"fmt"
	"os"
	"sersi/common"
	"sersi/core"
	"sersi/tea/spinner"
	"sersi/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

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
	fmt.Println(common.Logo)
	filePath, _ := cmd.Flags().GetString("file")
	fmt.Printf("Creating a new project using %s:\n", filePath)

	fileParser := core.NewFileParser(filePath)

	fileParserResult, err := fileParser.ExceuteMapping()
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

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
}
