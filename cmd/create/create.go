package create

import (
	"fmt"
	"sersi/core"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	namePrompt      = "What would you like to name this project"
	frameworkPrompt = "What framework will you be using"
	cssPrompt       = "How would you like to style this project"
	langPrompt      = "What language will you be using"
	defaultName     = "my-project"
	frameworkItems  = []string{"React", "Vue", "Vanilla"}
	cssItems        = []string{"Tailwind", "Bootstrap", "Traditional"}
	langItems       = []string{"Typescript", "Javascript"}
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Genrate Scaffold Application",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   runPrompts,
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

func runPrompts(cmd *cobra.Command, args []string) {
	name, err := helperPrompt(cmd, "name", promptui.Prompt{Label: namePrompt, Default: defaultName})
	if err != nil {
		fmt.Sprintf("Error Occured: %s", err.Error())
	}

	framework, err := helperSelect(cmd, "framework", promptui.Select{Label: frameworkPrompt, Items: frameworkItems})
	if err != nil {
		fmt.Sprintf("Error Occured: %s", err.Error())
	}

	css, err := helperSelect(cmd, "css", promptui.Select{Label: cssPrompt, Items: cssItems})
	if err != nil {
		fmt.Sprintf("Error Occured: %s", err.Error())
	}

	lang, err := helperSelect(cmd, "lang", promptui.Select{Label: langPrompt, Items: langItems})
	if err != nil {
		fmt.Sprintf("Error Occured: %s", err.Error())
	}

	scaffold := core.NewScaffoldBuilder().
		ProjectName(name).
		Framework(framework).
		CSS(css).
		Language(lang).
		Build()

	scaffold.Execute()
}

func helperPrompt(cmd *cobra.Command, flagName string, prompt promptui.Prompt) (string, error) {
	var value string
	if cmd.Flags().Changed(flagName) {
		value, _ = cmd.Flags().GetString(flagName)
		fmt.Fprintf(cmd.OutOrStderr(), "✔ %s: %s", flagName, value)

	} else {
		var err error
		value, err = prompt.Run()

		if err != nil {
			return "", err
		}
	}

	return value, nil
}

func helperSelect(cmd *cobra.Command, flagName string, prompt promptui.Select) (string, error) {
	var value string
	if cmd.Flags().Changed(flagName) {
		value, _ = cmd.Flags().GetString(flagName)
		fmt.Fprintf(cmd.OutOrStderr(), "✔ %s: %s", flagName, value)
	} else {
		var err error
		prompt.Templates = &selectTemplate
		_, value, err = prompt.Run()

		if err != nil {
			return "", err
		}
	}

	return value, nil
}
