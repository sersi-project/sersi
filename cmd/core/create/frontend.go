package create

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/sersi-project/sersi/internal/tui/menuinput"
	"github.com/sersi-project/sersi/internal/tui/textinput"
	"github.com/sersi-project/sersi/pkg"
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
	FrontendCmd.Flags().Bool("dry-run", false, "Dry run for testing")
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
		if *pn.(*textinput.Model).Quitting {
			os.Exit(0)
		}
		projectName = pn.(textinput.Model).Value
	}
	if err := pkg.ValidateName(projectName); err != nil {
		fmt.Println("Error validating project name:", err)
		os.Exit(1)
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
		if *fm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}
		framework = fm.(*menuinput.ListModel).Choice
		if framework == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(framework), pkg.FrontendFrameworks); err != nil {
		fmt.Println("Error validating framework:", err)
		os.Exit(1)
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
		if *cssm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}
		css = cssm.(*menuinput.ListModel).Choice
		if css == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(css), pkg.FrontendCSS); err != nil {
		fmt.Println("Error validating CSS:", err)
		os.Exit(1)
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
		if *langm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}
		lang = langm.(*menuinput.ListModel).Choice
		if lang == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(lang), pkg.FrontendLanguages); err != nil {
		fmt.Println("Error validating language:", err)
		os.Exit(1)
	}
	fmt.Printf("◉ %s\n", lang)

	dryRun, _ := cmd.Flags().GetBool("dry-run")
	if dryRun {
		fmt.Println("◉ Dry run enabled")
		os.Exit(0)
	}

	frontend := frontend.NewFrontendBuilder().
		ProjectName(projectName).
		Language(lang).
		Framework(framework).
		CSS(css).
		Monorepo(false).
		Polyrepos(false).
		Build()

	fmt.Printf("◉ %s\n", "Building...")
	err := frontend.Generate()
	if err != nil {
		fmt.Println("Error creating frontend project:", err)
		os.Exit(1)
	}

	frontendConfig := pkg.NewConfig(projectName, pkg.FrontendConfig{
		Framework: framework,
		CSS:       css,
		Language:  lang,
	}, pkg.BackendConfig{}, pkg.DevopsConfig{})
	if err := frontendConfig.GenerateSersiYaml(projectName); err != nil {
		fmt.Println("Error creating sersi.yaml:", err)
		os.Exit(1)
	}
	fmt.Printf("◉ %s\n", "Frontend project created successfully")
}
