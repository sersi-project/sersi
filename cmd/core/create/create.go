package create

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/sersi-project/sersi/internal/tui/menuinput"
	"github.com/sersi-project/sersi/internal/tui/textinput"
	"github.com/sersi-project/sersi/pkg"

	"github.com/spf13/cobra"
)

var customSetup bool

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Genrate Scaffold for Fullstack Application",
	Long:  `Genrate Scaffold for Fullstack Application with customizable options- `,
	Run:   RunCreate,
}

func init() {
	CreateCmd.AddCommand(BackendCmd)
	CreateCmd.AddCommand(FrontendCmd)

	CreateCmd.Flags().StringP("name", "n", "", "Name of project")
	CreateCmd.Flags().Bool("custom", false, "Custom setup")
	CreateCmd.Flags().Bool("dry-run", false, "Dry run for testing")
}

func RunCreate(cmd *cobra.Command, args []string) {
	var tprogram *tea.Program
	totalSteps := 3
	currentStep := 1
	common.PrintLogo()
	fmt.Printf("%s Creating a new full stack project...\n", common.OperationLabel)
	preset := pkg.Preset{}

	projectName, _ := cmd.Flags().GetString("name")
	if projectName == "" {
		tprogram = tea.NewProgram(textinput.InitialModel(totalSteps, currentStep, "Project Name", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}

		if *pn.(textinput.Model).Quitting {
			os.Exit(0)
		}

		projectName = pn.(textinput.Model).Value
		if err := pkg.ValidateName(projectName); err != nil {
			fmt.Println("Invalid project name")
			os.Exit(1)
		}

		if *pn.(textinput.Model).Quitting {
			os.Exit(0)
		}
	}
	currentStep++
	fmt.Printf("\n ├── %s Project name successfully set to: %s\n", common.SuccessLabel, projectName)

	customSetup, _ = cmd.Flags().GetBool("custom")
	if customSetup {
		fmt.Println(common.SuccessLabel + " Custom setup enabled")
	} else {
		stack, _ := cmd.Flags().GetString("stack")
		stackOpts := []string{
			"MongoDB + Express + React + Tailwind (Recommended)",
			"PostgreSQL + FastAPI + React + Tailwind",
			"PostgreSQL + Gin + Svelte + CSS",
			"Custom (Advanced)",
		}

		if stack == "" {
			tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Choose Stack", stackOpts, "Stack"))
			pn, err := tprogram.Run()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error running program:", err)
				os.Exit(1)
			}
			if *pn.(*menuinput.ListModel).Quitting {
				os.Exit(0)
			}

			stack = pn.(*menuinput.ListModel).Choice

			indexOfStack := getIndexOfStack(stack, stackOpts)
			switch indexOfStack {
			case -1:
				fmt.Println("Invalid stack")
				os.Exit(1)
			case 3:
				fmt.Printf(" │\n └── %s Custom setup enabled\n", common.SuccessLabel)
				customSetup = true
			default:
				preset = pkg.StackPresets[indexOfStack]
				fmt.Printf(" │\n └── %s Stack: Selected\n", common.SuccessLabel)
			}
		}
	}

	if customSetup {
		currentStep++
		totalSteps += 2
		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Frontend Framework", []string{"React", "Vue", "Svelte", "Vanilla"}, "Frontend Framework"))
		fm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *fm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendFramework := fm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Frontend CSS", []string{"CSS", "Tailwind", "Bootstrap"}, "Frontend CSS"))
		cssm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *cssm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendCSS := cssm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Frontend Language", []string{"Typescript", "Javascript"}, "Frontend Language"))
		lm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *lm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendLanguage := lm.(*menuinput.ListModel).Choice

		currentStep++
		fmt.Printf("\n%s Configuring frontend...\n", common.OperationLabel)
		fmt.Printf(" \n ├── %s Frontend Language: %s\n", common.SuccessLabel, frontendLanguage)
		fmt.Printf(" │\n ├── %s Frontend Framework: %s\n", common.SuccessLabel, frontendFramework)
		fmt.Printf(" │\n └── %s Frontend CSS: %s\n", common.SuccessLabel, frontendCSS)

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Language", []string{"Go", "Python", "Node", "TypeScript(Node)"}, "Backend Language"))
		blm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *blm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendLanguage := blm.(*menuinput.ListModel).Choice
		backendLanguage = strings.ToLower(backendLanguage)

		var frameworkOpts []string
		switch backendLanguage {
		case "go":
			frameworkOpts = pkg.BackendGoFrameworks
		case "python", "py":
			frameworkOpts = pkg.BackendPythonFrameworks
		case "node", "typescript(node)", "js", "ts", "typescript":
			frameworkOpts = pkg.BackendNodeFrameworks
		default:
			fmt.Println(common.ErrorLabel + " Error validating language: Invalid language")
			os.Exit(1)
		}

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Framework", frameworkOpts, "Backend Framework"))
		bfm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *bfm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendFramework := bfm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Backend Database", []string{"PostgreSQL", "MongoDB", "None"}, "Backend Database"))
		bdm, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *bdm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendDatabase := bdm.(*menuinput.ListModel).Choice
		currentStep++
		fmt.Printf("\n%s Configuring backend...\n", common.OperationLabel)
		fmt.Printf(" │\n ├── %s Backend Language: %s\n", common.SuccessLabel, backendLanguage)
		fmt.Printf(" │\n ├── %s Backend Framework: %s\n", common.SuccessLabel, backendFramework)
		fmt.Printf(" │\n └── %s Backend Database: %s\n", common.SuccessLabel, backendDatabase)

		preset = pkg.Preset{
			Frontend: pkg.FrontendConfig{
				Framework: frontendFramework,
				CSS:       frontendCSS,
				Language:  frontendLanguage,
			},
			Backend: pkg.BackendConfig{
				Language:  backendLanguage,
				Framework: backendFramework,
				Database:  backendDatabase,
			},
			Devops: pkg.DevopsConfig{
				CI:     "github",
				Docker: false,
			},
		}
	}

	tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Project Structure", []string{"Monorepo", "Polyrepos"}, "Project Structure"))
	projectStructure, err := tprogram.Run()
	if err != nil {
		fmt.Println(common.ErrorLabel+" Error running program:", err)
		os.Exit(1)
	}
	if *projectStructure.(*menuinput.ListModel).Quitting {
		os.Exit(0)
	}

	projectStructureChoice := projectStructure.(*menuinput.ListModel).Choice
	monorepo := false
	polyrepos := false

	wd, _ := os.Getwd()
	var frontendOutputPath, backendOutputPath string
	if projectStructureChoice == "Monorepo" {
		monorepo = true
		frontendOutputPath = wd + "/" + projectName + "/frontend"
		backendOutputPath = wd + "/" + projectName + "/backend"
	}

	if projectStructureChoice == "Polyrepos" {
		polyrepos = true
		frontendOutputPath = wd + "/" + projectName + "-frontend"
		backendOutputPath = wd + "/" + projectName + "-backend"
	}

	dryRun, _ := cmd.Flags().GetBool("dry-run")
	if dryRun {
		fmt.Println(common.OperationLabel + " Dry run enabled")
		os.Exit(0)
	}

	fmt.Printf("\n%s Creating frontend project...\n", common.OperationLabel)

	frontendScaffold := frontend.NewFrontendBuilder().
		ProjectName(projectName).
		Language(preset.Frontend.Language).
		Framework(preset.Frontend.Framework).
		CSS(preset.Frontend.CSS).
		Monorepo(monorepo).
		Polyrepos(polyrepos).
		Build()

	if err := frontendScaffold.Generate(); err != nil {
		fmt.Println("Error creating frontend project:", err)
		os.Exit(1)
	}

	fmt.Printf(" \n └── %s Frontend\n", common.SuccessLabel)
	fmt.Printf("     File created: %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(frontendOutputPath))

	fmt.Printf("\n%s Creating backend project...\n", common.OperationLabel)

	backend := backend.NewBackendBuilder().
		ProjectName(projectName).
		Language(preset.Backend.Language).
		Framework(preset.Backend.Framework).
		Database(preset.Backend.Database).
		Monorepo(monorepo).
		Polyrepos(polyrepos).
		Build()

	if err := backend.Generate(); err != nil {
		fmt.Println(common.ErrorLabel+" Error creating backend project:", err)
		os.Exit(1)
	}

	fmt.Printf("\n └── %s Backend\n", common.SuccessLabel)
	fmt.Printf("     File created: %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(backendOutputPath))

	if monorepo {
		config := pkg.NewConfig(projectName, preset.Frontend, preset.Backend, pkg.DevopsConfig{})
		if err := config.GenerateSersiYaml(projectName); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}
	}

	if polyrepos {
		frontendConfig := pkg.NewConfig(projectName, preset.Frontend, pkg.BackendConfig{}, pkg.DevopsConfig{})
		if err := frontendConfig.GenerateSersiYaml(projectName + "-frontend"); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}

		backendConfig := pkg.NewConfig(projectName, pkg.FrontendConfig{}, preset.Backend, pkg.DevopsConfig{})
		if err := backendConfig.GenerateSersiYaml(projectName + "-backend"); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n%s Created project %s\n", common.SuccessLabel, projectName)
}

func getIndexOfStack(stack string, stackOpts []string) int {
	for i, s := range stackOpts {
		if s == stack {
			return i
		}
	}
	return -1
}
