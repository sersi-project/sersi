package create

import (
	"fmt"
	"os"
	"sersi/common"
	"sersi/core"
	"sersi/tea/menuinput"
	"sersi/tea/spinner"
	"sersi/tea/textinput"
	"sersi/utils"
	"sersi/validator"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)


var cancelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EE4B2B")).Bold(true)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Genrate Scaffold Application",
	Long:  `Genrate Scaffold Application with customizable options`,
	Run:   Run,
}

func init() {
	Cmd.Flags().StringP("name", "n", "", "Name of project")
	Cmd.Flags().StringP("framework", "t", "", "Name of framework for template")
	Cmd.Flags().StringP("css", "s", "", "Styling for template")
	Cmd.Flags().StringP("lang", "l", "", "Javascript or Typescript")
}

func Run(cmd *cobra.Command, args []string) {
	var err error
	var tprogram *tea.Program

	common.PrintLogo()

	scaffoldBuilder := core.NewScaffoldBuilder()

	projectName, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Printf("Error getting project name: %s", err)
		os.Exit(1)
	}

	if projectName != "" {
		if utils.FileExists(projectName) {
			fmt.Printf("Error: %s", "Project already exists")
			os.Exit(1)
		}

		if err = validator.ValidateString(projectName); err != nil {
			fmt.Printf("Error: %s", "Invalid project name")
			os.Exit(1)
		}
		scaffoldBuilder.ProjectName(projectName)
	} else {
		tprogram = tea.NewProgram(textinput.InitialModel("What is this project called?", "Enter project name"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Printf("Error running program: %s", err)
			os.Exit(1)
		}

		projectName = pn.(textinput.Model).Value
		if projectName == "" {
			fmt.Printf("\n%s\n", cancelStyle.Render("Operation cancelled!"))
			os.Exit(0)
		}
		scaffoldBuilder.ProjectName(projectName)
	}

	framework, err := cmd.Flags().GetString("framework")
	if err != nil {
		fmt.Printf("Error getting framework: %s", err)
		os.Exit(1)
	}

	if framework != "" {
		scaffoldBuilder.Framework(framework)
	} else {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput("Framework", []string{"React", "Svelte", "Vanilla", "Vue"}, "Framework"))
		fm, err := tprogram.Run()
		if err != nil {
			fmt.Printf("Error running program: %s", err)
			os.Exit(1)
		}

		if fm.(*menuinput.ListModel).Choice == "" {
			os.Exit(0)
		}
		scaffoldBuilder.Framework(fm.(*menuinput.ListModel).Choice)
	}

	css, err := cmd.Flags().GetString("css")
	if err != nil {
		fmt.Printf("Error getting css: %s", err)
		os.Exit(1)
	}

	if css != "" {
		scaffoldBuilder.CSS(css)
	} else {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput("CSS library", []string{"Tailwind", "Bootstrap", "Traditional"}, "CSS"))
		cm, err := tprogram.Run()
		if err != nil {
			fmt.Printf("Error running program: %s", err)
			os.Exit(1)
		}

		if cm.(*menuinput.ListModel).Choice == "" {
			os.Exit(0)
		}
		scaffoldBuilder.CSS(cm.(*menuinput.ListModel).Choice)
	}

	lang, err := cmd.Flags().GetString("lang")
	if err != nil {
		fmt.Printf("Error getting lang: %s", err)
		os.Exit(1)
	}

	if lang != "" {
		scaffoldBuilder.Language(lang)
	} else {
		tprogram = tea.NewProgram(menuinput.InitialMenuInput("Language", []string{"Javascript", "Typescript"}, "Language"))
		ln, err := tprogram.Run()
		if err != nil {
			fmt.Printf("Error running program: %s", err)
			os.Exit(1)
		}

		if ln.(*menuinput.ListModel).Choice == "" {
			os.Exit(0)
		}
		scaffoldBuilder.Language(ln.(*menuinput.ListModel).Choice)
	}

	projectPath := utils.GetProjectPath(projectName)
	loading := tea.NewProgram(spinner.InitialSpinnerModel(projectPath, scaffoldBuilder.Build()))
	_, err = loading.Run()
	if err != nil {
		fmt.Printf("Error running program: %s", err)
		os.Exit(1)
	}
}
