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
	Long:  `Create a new backend project using Sersi Scaffold`,
	Run:   RunBackend,
}

func init() {
	BackendCmd.Flags().StringP("name", "n", "", "Name of the project")
	BackendCmd.Flags().StringP("framework", "t", "", "Name of framework for template")
	BackendCmd.Flags().StringP("database", "d", "", "Database for template")
	BackendCmd.Flags().StringP("lang", "l", "", "Javascript or Typescript")
	BackendCmd.Flags().Bool("dry-run", false, "Dry run for testing")
}

func RunBackend(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	fmt.Printf("\n%s Creating a new backend project...\n", common.OperationLabel)

	projectName, _ := cmd.Flags().GetString("name")
	language, _ := cmd.Flags().GetString("lang")
	framework, _ := cmd.Flags().GetString("framework")
	database, _ := cmd.Flags().GetString("database")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	currentStep := 1
	if projectName == "" {
		tprogram := tea.NewProgram(textinput.InitialModel(totalSteps, currentStep, "Project Name", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Printf("\n%s Error running program: %s\n", common.ErrorLabel, err)
			os.Exit(1)
		}

		if *pn.(textinput.Model).Quitting {
			os.Exit(0)
		}

		projectName = pn.(textinput.Model).Value
	}
	if err := pkg.ValidateName(projectName); err != nil {
		fmt.Printf("\n%s Error validating project name: %s\n", common.ErrorLabel, err)
		os.Exit(1)
	}
	currentStep++
	fmt.Printf("\n ├── %s Project name successfully set to: %s\n", common.SuccessLabel, projectName)

	if language == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Language", []string{"Node", "Typescript(node)", "Go", "Python"}, "Language"))
		langm, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		language = langm.(*menuinput.ListModel).Choice
		if language == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(language), pkg.BackendLanguages); err != nil {
		fmt.Printf("\n%s Error validating language: %s\n", common.ErrorLabel, err)
		os.Exit(1)
	}
	currentStep++
	fmt.Printf(" │\n ├── %s Language successfully set to: %s\n", common.SuccessLabel, language)

	var opts []string

	switch language {
	case "Node", "Typescript(node)", "node", "js", "ts", "typescript(node)", "typescript":
		opts = pkg.BackendNodeFrameworks
	case "Go", "go":
		opts = pkg.BackendGoFrameworks
	case "Python", "python", "py":
		opts = pkg.BackendPythonFrameworks
	default:
		fmt.Printf("\n%s Error validating language: Invalid language\n", common.ErrorLabel)
		fmt.Println("Allowed languages:", pkg.BackendLanguages)
		os.Exit(1)
	}
	var optsTitle []string

	for _, v := range opts {
		optsTitle = append(optsTitle, strings.Title(strings.ToLower(v))) //nolint
	}

	if framework == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Backend Framework", optsTitle, "Framework"))
		framem, err := tprogram.Run()
		if err != nil {
			fmt.Printf("\n%s Error running program: %s\n", common.ErrorLabel, err)
			os.Exit(1)
		}
		framework = framem.(*menuinput.ListModel).Choice
		if framework == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(framework), opts); err != nil {
		fmt.Printf("\n%s Error validating framework: %s\n", common.ErrorLabel, err)
		os.Exit(1)
	}
	currentStep++
	fmt.Printf(" │\n ├── %s Framework successfully set to: %s\n", common.SuccessLabel, framework)

	if database == "" {
		tprogram := tea.NewProgram(menuinput.InitialMenuInput(totalSteps, currentStep, "Database", []string{"PostgreSQL", "MongoDB", "None"}, "Database"))
		db, err := tprogram.Run()
		if err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

		database = db.(*menuinput.ListModel).Choice
		if database != "None" {
			fmt.Printf(" │\n ├── %s Please note that this version does not currently support databases: defaulting to none\n", common.OperationLabel)
			database = "none"
		}

		if database == "" {
			os.Exit(0)
		}
	}
	if err := pkg.ValidateOptions(strings.ToLower(database), pkg.DatabaseFrameworks); err != nil {
		fmt.Printf("\n%s Error validating database: %s\n", common.ErrorLabel, err)
		os.Exit(1)
	}

	fmt.Printf(" │\n ├── %s Database successfully set to: %s\n", common.SuccessLabel, database)

	if dryRun {
		fmt.Printf(" │\n ├── %s Dry run enabled\n", common.SuccessLabel)
		os.Exit(0)
	}

	backend := backend.NewBackendBuilder().
		ProjectName(projectName).
		Language(language).
		Framework(framework).
		Database(database).
		Monorepo(false).
		Polyrepos(false).
		Build()

	backendConfig := pkg.NewConfig(projectName, pkg.FrontendConfig{}, pkg.BackendConfig{
		Framework: framework,
		Language:  language,
		Database:  database,
	}, pkg.DevopsConfig{})

	fmt.Printf("\n%s Building...\n", common.OperationLabel)
	err := backend.Generate()
	if err != nil {
		fmt.Printf("\n%s Error creating backend project: %s\n", common.ErrorLabel, err)
		return
	}

	if err := backendConfig.GenerateSersiYaml(projectName); err != nil {
		fmt.Printf("\n%s Error creating sersi.yaml: %s\n", common.ErrorLabel, err)
		os.Exit(1)
	}

	fmt.Printf(" │\n └── %s Backend project created successfully\n", common.SuccessLabel)
}
