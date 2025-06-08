package create

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
	Short: "Genrate Scaffold Application",
	Long:  `Genrate Scaffold Application with customizable options`,
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
	fmt.Println("◉ Creating a new full stack project...")
	preset := pkg.Preset{}

	projectName, _ := cmd.Flags().GetString("name")
	if projectName == "" {
		tprogram = tea.NewProgram(textinput.InitialModel(totalSteps, currentStep, "Project Name", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
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
	fmt.Println("◉ Project name:", projectName)

	customSetup, _ = cmd.Flags().GetBool("custom")
	if customSetup {
		fmt.Println("◉ Custom setup enabled")
	} else {
		stack, _ := cmd.Flags().GetString("stack")
		stackOpts := []string{
			"MongoDB + Express + React + Tailwind (Recommended)",
			"PostgreSQL + FastAPI + React + Tailwind",
			"PostgreSQL + Gin + Svelte + CSS",
			"Custom (Advanced)",
		}

		if stack == "" {
			tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Stack", stackOpts, "Stack"))
			pn, err := tprogram.Run()
			if err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
			if *pn.(*menuinput.ListModel).Quitting {
				os.Exit(0)
			}

			stack = pn.(*menuinput.ListModel).Choice

			indexOfStack := getIndexOfStack(stack, stackOpts)
			if indexOfStack == -1 {
				fmt.Println("Invalid stack")
				os.Exit(1)
			} else if indexOfStack == 3 {
				fmt.Println("◉ Custom setup enabled")
				customSetup = true
			} else {
				preset = pkg.StackPresets[indexOfStack]
				fmt.Println("◉ Stack: Selected")
			}
		}
	}

	if customSetup {
		currentStep++
		totalSteps += 2
		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Frontend Framework", []string{"React", "Vue", "Svelte", "Vanilla"}, "Frontend Framework"))
		fm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *fm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendFramework := fm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Frontend CSS", []string{"CSS", "Tailwind", "Bootstrap"}, "Frontend CSS"))
		cssm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *cssm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendCSS := cssm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Frontend Language", []string{"Typescript", "Javascript"}, "Frontend Language"))
		lm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *lm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		frontendLanguage := lm.(*menuinput.ListModel).Choice

		currentStep++
		fmt.Println("◉ Frontend Language:", frontendLanguage)
		fmt.Println("◉ Frontend Framework:", frontendFramework)
		fmt.Println("◉ Frontend CSS:", frontendCSS)

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Language", []string{"Go", "Python", "Node", "TypeScript(Node)"}, "Backend Language"))
		blm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *blm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendLanguage := blm.(*menuinput.ListModel).Choice

		frameworkOpts := []string{}
		if backendLanguage == "Go" {
			frameworkOpts = pkg.BackendGoFrameworks
		} else if backendLanguage == "Python" {
			frameworkOpts = pkg.BackendPythonFrameworks
		} else if backendLanguage == "Node" || backendLanguage == "TypeScript(Node)" {
			frameworkOpts = pkg.BackendNodeFrameworks
		}

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Backend Framework", frameworkOpts, "Backend Framework"))
		bfm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *bfm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendFramework := bfm.(*menuinput.ListModel).Choice

		tprogram = tea.NewProgram(menuinput.InitialMenuInput(3, 2, "Backend Database", []string{"PostgreSQL", "MongoDB", "None"}, "Backend Database"))
		bdm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		if *bdm.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		backendDatabase := bdm.(*menuinput.ListModel).Choice
		currentStep++
		fmt.Println("◉ Backend Language:", backendLanguage)
		fmt.Println("◉ Backend Framework:", backendFramework)
		fmt.Println("◉ Backend Database:", backendDatabase)

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
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if *projectStructure.(*menuinput.ListModel).Quitting {
		os.Exit(0)
	}

	projectStructureChoice := projectStructure.(*menuinput.ListModel).Choice
	monorepo := false
	polyrepos := false
	if projectStructureChoice == "Monorepo" {
		monorepo = true
	}

	if projectStructureChoice == "Polyrepos" {
		polyrepos = true
	}

	dryRun, _ := cmd.Flags().GetBool("dry-run")
	if dryRun {
		fmt.Println("◉ Dry run enabled")
		os.Exit(0)
	}

	fmt.Println("◉ Creating frontend project...")

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

	fmt.Println("◉ Creating backend project...")

	backend := backend.NewBackendBuilder().
		ProjectName(projectName).
		Language(preset.Backend.Language).
		Framework(preset.Backend.Framework).
		Database(preset.Backend.Database).
		Monorepo(monorepo).
		Polyrepos(polyrepos).
		Build()

	if err := backend.Generate(); err != nil {
		fmt.Println("Error creating backend project:", err)
		os.Exit(1)
	}

	if monorepo {
		config := pkg.NewConfig(projectName, preset.Frontend, preset.Backend, pkg.DevopsConfig{})
		if err := config.GenerateSersiYaml(projectName); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}
	}

	if polyrepos {
		frontendConfig := pkg.NewConfig(projectName, preset.Frontend, pkg.BackendConfig{}, pkg.DevopsConfig{})
		if err := frontendConfig.GenerateSersiYaml(projectName); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}

		backendConfig := pkg.NewConfig(projectName, pkg.FrontendConfig{}, preset.Backend, pkg.DevopsConfig{})
		if err := backendConfig.GenerateSersiYaml(projectName); err != nil {
			fmt.Println("Error creating sersi.yaml:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("◉ %s created successfully\n", projectName)
}

func getIndexOfStack(stack string, stackOpts []string) int {
	for i, s := range stackOpts {
		if s == stack {
			return i
		}
	}
	return -1
}
