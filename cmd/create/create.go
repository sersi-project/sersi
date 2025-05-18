package create

import (
	"fmt"
	"sersi/core/scaffold"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Genrate Scaffold Application",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   createTemplate,
}

func init() {
	Cmd.Flags().String("name", "my-project", "name of project")
	Cmd.Flags().String("framework", "react", "name of framework to use")
	Cmd.Flags().String("css", "css", "styling for template")
	Cmd.Flags().String("lang", "js", "javascript or Typescript")
}

func createTemplate(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	framework, _ := cmd.Flags().GetString("framework")
	css, _ := cmd.Flags().GetString("css")
	lang, _ := cmd.Flags().GetString("lang")

	fmt.Printf("created template:%10s\n", framework)
	fmt.Printf("using css:%10s\n", css)
	fmt.Printf("using tsconfig.json:%10t\n", lang == "ts")
	scaffold.ScaffoldProject(name, framework, css, lang)
}
