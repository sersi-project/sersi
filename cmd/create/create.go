package create

import (
	"fmt"
	"os"
	"sersi/core"
	"sersi/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	namePrompt      = "What would you like to name this project"
	frameworkPrompt = "What framework will you be using"
	cssPrompt       = "How would you like to style this project"
	langPrompt      = "What language will you be using"
	defaultName     = "my-project"
	frameworkItems  = []string{"React", "Vue", "Vanilla", "Svelte"}
	cssItems        = []string{"Tailwind", "Bootstrap", "Traditional"}
	langItems       = []string{"Typescript", "Javascript"}
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Genrate Scaffold Application",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   Run,
}

var selectTemplate = promptui.SelectTemplates{
	Active:   "> {{ . | cyan }}",
	Inactive: "  {{ . | faint }}",
	Selected: "✔ {{ . | green }}",
}

func init() {
	Cmd.Flags().StringP("name", "n", defaultName, "name of project")
	Cmd.Flags().StringP("framework", "f", "react", "name of framework to use")
	Cmd.Flags().StringP("css", "s", "css", "styling for template")
	Cmd.Flags().StringP("lang", "l", "js", "javascript or Typescript")
}

func Run(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(model.InitialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %s", err)
		os.Exit(1)
	}

	m := finalModel.(model.Model)
	if m.Name == "" {
		os.Exit(1)
	}
	if m.Framework == "" {
		os.Exit(1)
	}
	if m.Css == "" {
		os.Exit(1)
	}
	if m.Lang == "" {
		os.Exit(1)
	}

	scaffold := core.NewScaffoldBuilder().
		ProjectName(m.Name).
		Framework(m.Framework).
		CSS(m.Css).
		Language(m.Lang).Build()

	scaffold.Execute()
}

// func helperPrompt(cmd *cobra.Command, flagName string, prompt promptui.Prompt) (string, error) {
// 	var value string
// 	if cmd.Flags().Changed(flagName) {
// 		value, _ = cmd.Flags().GetString(flagName)
// 		fmt.Fprintf(cmd.OutOrStderr(), "✔ %s: %s", flagName, value)

// 	} else {
// 		var err error
// 		value, err = prompt.Run()

// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	return value, nil
// }

// func helperSelect(cmd *cobra.Command, flagName string, prompt promptui.Select) (string, error) {
// 	var value string
// 	if cmd.Flags().Changed(flagName) {
// 		value, _ = cmd.Flags().GetString(flagName)
// 		fmt.Fprintf(cmd.OutOrStderr(), "✔ %s: %s", flagName, value)
// 	} else {
// 		var err error
// 		prompt.Templates = &selectTemplate
// 		_, value, err = prompt.Run()

// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	return value, nil
// }
