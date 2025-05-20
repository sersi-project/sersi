package build

import (
	"fmt"
	"sersi/core"

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

func Run(cmd *cobra.Command, args []string) { // Execute the command logic here
	fmt.Printf("Creating a new project with the following options:\n")
	filePath, _ := cmd.Flags().GetString("file")

	fileParser := core.NewFileParser(filePath)

	fileParserResult, err := fileParser.ExceuteMapping()
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	scaffold := core.NewScaffoldBuilder().
		ProjectName(fileParserResult.Name).
		Framework(fileParserResult.Scaffold.Frontend.Framework).
		CSS(fileParserResult.Scaffold.Frontend.CSS).
		Language(fileParserResult.Scaffold.Frontend.Language).
		Build()

	scaffold.Execute()
}
