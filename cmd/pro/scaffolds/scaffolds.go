package scaffolds

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sersi-project/sersi/common"
	authorization "github.com/sersi-project/sersi/internal/authorization"
	"github.com/sersi-project/sersi/internal/scaffold"
	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/sersi-project/sersi/internal/tui/menuinput"
	"github.com/sersi-project/sersi/internal/tui/textinput"
	"github.com/sersi-project/sersi/pkg"
	"github.com/spf13/cobra"
)

var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffold store actions for Sersi Pro (save, view, update, delete, use)",
	Long:  `Scaffold store actions for Sersi Pro (save, view, update, delete, use)`,
	Run:   runScaffold,
}

func init() {
	ScaffoldCmd.Flags().StringP("name", "n", "", "Name of project")
	ScaffoldCmd.Flags().StringP("action", "a", "", "Action to perform (save, view, update, delete)")
	ScaffoldCmd.Flags().StringP("file-path", "f", "", "File path of project")
}

func runScaffold(cmd *cobra.Command, args []string) {
	common.PrintLogo()

	_, ok := authorization.CheckStatus()
	if !ok {
		fmt.Println("You are not logged in")
		os.Exit(0)
	}

	svc := scaffold.NewScaffoldService()
	name, _ := cmd.Flags().GetString("name")
	action, _ := cmd.Flags().GetString("action")
	filePath, _ := cmd.Flags().GetString("file-path")

	switch action {
	case "save":
		if filePath == "" {
			fmt.Printf("%s --file-path required to save template", common.ErrorLabel)
			os.Exit(0)
		}
		fileParser := pkg.NewFileParser(filePath)
		fileParserResult, err := fileParser.ExceuteMapping()
		if err != nil {
			fmt.Println("Error parsing file:", err)
			os.Exit(0)
		}

		fmt.Printf("%s Saving scaffold to store \n", common.OperationLabel)
		err = svc.SaveScaffold(fileParserResult)
		if err != nil {
			fmt.Printf("\n ├── %s Error saving scaffold: %s\n", common.ErrorLabel, err.Error())
			os.Exit(0)
		}

		fmt.Printf("\n ├── %s Scaffold saved successfully", common.SuccessLabel)
		os.Exit(0)
	case "list":
		fmt.Printf("\n%s Retreiving stored scaffolds..\n", common.OperationLabel)
		var scaffoldNames []string
		list, err := svc.GetAllScaffolds()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error getting all scaffolds:", err)
			os.Exit(1)
		}
		for _, scaffold := range list {
			scaffoldNames = append(scaffoldNames, scaffold.Name)
		}

		noOfItems := len(list)

		tprogram := tea.NewProgram(menuinput.InitialMenuInput(noOfItems, noOfItems, "Select Scaffold from Store", scaffoldNames, "scaffold"))
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		if *pn.(*menuinput.ListModel).Quitting {
			os.Exit(0)
		}

		scaffoldName := pn.(*menuinput.ListModel).Choice
		fmt.Printf("\n  ├── %s Selected scaffold: %s\n", common.SuccessLabel, scaffoldName)

		var config pkg.Config
		for _, scaffold := range list {
			if scaffold.Name == scaffoldName {
				config = scaffold
				break
			}
		}

		fmt.Printf("\n  ├── %s Scaffold: %s\n", common.SuccessLabel, config.Name)

		if config.Scaffold.Frontend.Framework != "" {
			fmt.Printf("\n%s Building Frontend\n", common.OperationLabel)
			frontend, err := buildFrontend(&config)
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error building frontend:", err)
				os.Exit(1)
			}

			err = frontend.Generate()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error generating frontend:", err)
				os.Exit(1)
			}
			fmt.Printf("\n  ├── %s Frontend generated successfully\n", common.SuccessLabel)
		}

		if config.Scaffold.Backend.Language != "" {
			fmt.Printf("\n%s Building Backend: %s\n", common.OperationLabel, config.Scaffold.Backend.Language)
			backend, err := buildBackend(&config)
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error building backend:", err)
				os.Exit(1)
			}

			err = backend.Generate()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error generating backend:", err)
				os.Exit(1)
			}
			fmt.Printf("\n  ├── %s Backend generated successfully\n", common.SuccessLabel)
			fmt.Printf("\n%s Scaffold generated successfully\n", common.SuccessLabel)
		}
		os.Exit(0)
	case "use":
		if name == "" {
			fmt.Println("Please provide a scaffold name")
			tprogram := tea.NewProgram(textinput.InitialModel(1, 1, "Scaffold Name", "Enter scaffold name"))
			pn, err := tprogram.Run()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error running program:", err)
				os.Exit(1)
			}
			if *pn.(textinput.Model).Quitting {
				os.Exit(0)
			}

			name = pn.(textinput.Model).Value
			if err := pkg.ValidateName(name); err != nil {
				fmt.Println("Invalid project name")
				os.Exit(1)
			}
		}

		scaffold, err := svc.GetScaffold(name)
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error getting scaffold:", err)
			os.Exit(1)
		}

		if scaffold.Scaffold.Frontend.Framework != "" {
			frontend, err := buildFrontend(scaffold)
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error building frontend:", err)
				os.Exit(1)
			}

			err = frontend.Generate()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error generating frontend:", err)
				os.Exit(1)
			}
		}

		if scaffold.Scaffold.Backend.Language != "" {
			backend, err := buildBackend(scaffold)
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error building backend:", err)
				os.Exit(1)
			}

			err = backend.Generate()
			if err != nil {
				fmt.Println(common.ErrorLabel+" Error generating backend:", err)
				os.Exit(1)
			}
		}

		fmt.Printf("\n%s Scaffold %s used successfully\n", common.SuccessLabel, name)
		os.Exit(0)
	case "update":
		if filePath == "" {
			fmt.Printf("%s --file-path required to update template", common.ErrorLabel)
			os.Exit(0)
		}
		fileParser := pkg.NewFileParser(filePath)
		fileParserResult, err := fileParser.ExceuteMapping()
		if err != nil {
			fmt.Printf("\n ├── %s Error parsing file: %s", common.ErrorLabel, err.Error())
			os.Exit(0)
		}

		fmt.Printf("%s Updating scaffold in store \n", common.OperationLabel)
		err = svc.UpdateScaffold(fileParserResult)
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error updating scaffold:", err)
			os.Exit(1)
		}

		fmt.Printf("\n  ├── %s Updated scaffold - %s - successfully\n", common.SuccessLabel, filePath)
		os.Exit(0)
	case "delete":
		if name == "" {
			fmt.Printf("%s Please provide a name using the --name flag", common.ErrorLabel)
			os.Exit(0)
		}

		fmt.Printf("\n%s Delete scaffold %s\n", common.OperationLabel, name)
		err := svc.DeleteScaffold(name)
		if err != nil {
			fmt.Printf("\n ├── %s Error deleting scaffold: %s", common.ErrorLabel, err.Error())
			os.Exit(0)
		}
		fmt.Printf("\n  ├── %s Deleted scaffold - %s - successfully\n", common.SuccessLabel, name)
		os.Exit(0)
	default:
		fmt.Printf("\n%s Invalid action \n", common.ErrorLabel)
		fmt.Printf("Allow actions -> save, update, delete, list, use")
		os.Exit(0)
	}
}

func buildFrontend(fileParserResult *pkg.Config) (scaffold.Scaffold, error) {
	f := frontend.NewFrontendBuilder().
		ProjectName(fileParserResult.Name).
		Framework(fileParserResult.Scaffold.Frontend.Framework).
		CSS(fileParserResult.Scaffold.Frontend.CSS).
		Language(fileParserResult.Scaffold.Frontend.Language).
		Monorepo(fileParserResult.Structure == "monorepo" || fileParserResult.Structure == "mono").
		Polyrepos(fileParserResult.Structure == "polyrepos" || fileParserResult.Structure == "poly").
		Build()
	return f, nil
}

func buildBackend(fileParserResult *pkg.Config) (scaffold.Scaffold, error) {
	b := backend.NewBackendBuilder().
		ProjectName(fileParserResult.Name).
		Language(fileParserResult.Scaffold.Backend.Language).
		Framework(fileParserResult.Scaffold.Backend.Framework).
		Database(fileParserResult.Scaffold.Backend.Database).
		Monorepo(fileParserResult.Structure == "monorepo" || fileParserResult.Structure == "mono").
		Polyrepos(fileParserResult.Structure == "polyrepos" || fileParserResult.Structure == "poly").
		Build()
	return b, nil
}
