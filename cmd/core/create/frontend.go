package create

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sersi-project/core/common"
	"github.com/sersi-project/core/internal/scaffold/frontend"
	"github.com/sersi-project/core/internal/tui/menuinput"
	"github.com/sersi-project/core/internal/tui/textinput"
	"github.com/sersi-project/core/pkg"
	"github.com/spf13/cobra"
)

const totalSteps = 5

var FrontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Create a new frontend project",
	Long:  `Create a new frontend project with the given name`,
	Run:   RunFrontend,
}

func init() {
	FrontendCmd.Flags().StringP("name", "n", "", "Name of the project")
	FrontendCmd.Flags().StringP("framework", "t", "", "Name of framework for template")
	FrontendCmd.Flags().StringP("css", "s", "", "Styling for template")
	FrontendCmd.Flags().StringP("lang", "l", "", "Javascript or Typescript")
}

func RunFrontend(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	fmt.Println("◉ Creating a new frontend project...")
	projectName, _ := cmd.Flags().GetString("name")
	framework, _ := cmd.Flags().GetString("framework")
	css, _ := cmd.Flags().GetString("css")
	lang, _ := cmd.Flags().GetString("lang")
	var tprogram *tea.Program

	currentStep := 1

	if projectName == "" {
		tprogram = tea.NewProgram(textinput.InitialModel(totalSteps, currentStep, "Project Name", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		projectName = pn.(textinput.Model).Value
		if err := pkg.ValidateName(projectName); err != nil {
			os.Exit(1)
		}
	} else {
		if err := pkg.ValidateName(projectName); err != nil {
			fmt.Println("Error validating project name:", err)
			os.Exit(1)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", projectName)

	if framework == "" {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Frontend Framework", []string{"React", "Svelte", "Vanilla", "Vue"}, "Framework"))
		fm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		framework = fm.(*menuinput.ListModel).Choice
		if err := pkg.ValidateOptions(strings.ToLower(framework), pkg.FrontendFrameworks); err != nil {
			fmt.Println("Error validating framework:", err)
			os.Exit(1)
		}
		if framework == "" {
			os.Exit(0)
		}
	}
	currentStep++

	fmt.Printf("◉ %s\n", framework)

	if css == "" {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Frontend CSS", []string{"CSS", "Tailwind", "Bootstrap"}, "CSS"))
		cssm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		css = cssm.(*menuinput.ListModel).Choice
		if err := pkg.ValidateOptions(strings.ToLower(css), pkg.FrontendCSS); err != nil {
			fmt.Println("Error validating CSS:", err)
			os.Exit(1)
		}
		if css == "" {
			os.Exit(0)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", css)

	if lang == "" {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Frontend Language", []string{"JavaScript", "TypeScript"}, "Language"))
		langm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		lang = langm.(*menuinput.ListModel).Choice
		if err := pkg.ValidateOptions(strings.ToLower(lang), pkg.FrontendLanguages); err != nil {
			fmt.Println("Error validating language:", err)
			os.Exit(1)
		}
		if lang == "" {
			os.Exit(0)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", lang)

	frontend := frontend.NewFrontendBuilder().
		ProjectName(projectName).
		Language(lang).
		Framework(framework).
		CSS(css).
		Monorepo(false).
		Build()

	fmt.Printf("◉ %s\n", "Building...")
	err := frontend.Generate()
	if err != nil {
		fmt.Println("Error creating frontend project:", err)
		return
	}
	fmt.Printf("◉ %s\n", "Frontend project created successfully")
}
