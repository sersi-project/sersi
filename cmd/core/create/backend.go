package create

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/sersi-project/sersi/internal/tui/menuinput"
	"github.com/sersi-project/sersi/internal/tui/textinput"
	"github.com/sersi-project/sersi/pkg"
	"github.com/spf13/cobra"
)

var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Create a new backend project",
	Long:  `Create a new backend project with the given name`,
	Run:   RunBackend,
}

func init() {
	BackendCmd.Flags().StringP("name", "n", "", "Name of the project")
	BackendCmd.Flags().StringP("framework", "t", "", "Name of framework for template")
	BackendCmd.Flags().StringP("database", "d", "", "Database for template")
	BackendCmd.Flags().StringP("language", "l", "", "Javascript or Typescript")
}

func RunBackend(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	fmt.Println("◉ Creating a new backend project...")

	projectName, _ := cmd.Flags().GetString("name")
	language, _ := cmd.Flags().GetString("language")
	framework, _ := cmd.Flags().GetString("framework")
	database, _ := cmd.Flags().GetString("database")

	currentStep := 1
	if projectName == "" {
		tprogram := tea.NewProgram(textinput.InitialModel(totalSteps, currentStep, "Project Name", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

		if *pn.(textinput.Model).Quitting {
			os.Exit(0)
		}

		projectName = pn.(*textinput.Model).Value
		if err := pkg.ValidateName(projectName); err != nil {
			os.Exit(1)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", projectName)

	if language == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Language", []string{"Node", "Typescript(node)", "Go", "Python"}, "Language"))
		langm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		language = langm.(*menuinput.ListModel).Choice
		if err := pkg.ValidateOptions(strings.ToLower(language), pkg.BackendLanguages); err != nil {
			os.Exit(1)
		}
		if language == "" {
			os.Exit(0)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", language)

	var opts []string
	if language == "JavaScript" || language == "TypeScript" || language == "js" || language == "ts" {
		opts = pkg.BackendNodeFrameworks
	} else if language == "Go" || language == "go" {
		opts = pkg.BackendGoFrameworks
	} else {
		opts = pkg.BackendPythonFrameworks
	}

	var optsTitle []string

	for _, v := range opts {
		optsTitle = append(optsTitle, strings.Title(strings.ToLower(v)))
	}

	if framework == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Framework", optsTitle, "Framework"))
		framem, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		framework = framem.(*menuinput.ListModel).Choice

		if err := pkg.ValidateOptions(strings.ToLower(framework), opts); err != nil {
			fmt.Println("Error validating framework:", err)
			os.Exit(1)
		}
		if framework == "" {
			os.Exit(0)
		}
	}
	currentStep++
	fmt.Printf("◉ %s\n", framework)

	if database == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Database", []string{"PostgreSQL", "MongoDB", "None"}, "Database"))
		db, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		database = db.(*menuinput.ListModel).Choice
		if err := pkg.ValidateOptions(strings.ToLower(database), pkg.DatabaseFrameworks); err != nil {
			fmt.Println("Error validating database:", err)
			os.Exit(1)
		}
		if database == "" {
			os.Exit(0)
		}
	}

	fmt.Printf("◉ %s\n", database)

	backend := backend.NewBackendBuilder().
		ProjectName(projectName).
		Language(language).
		Framework(framework).
		Database(database).
		Monorepo(false).
		Build()

	fmt.Printf("◉ %s\n", "Building...")
	err := backend.Generate()
	if err != nil {
		fmt.Println("Error creating backend project:", err)
		return
	}
	fmt.Println("◉ Backend project created successfully")
}
