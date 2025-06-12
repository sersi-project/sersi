package scaffolds

import (
	"fmt"
	"os"

	"github.com/sersi-project/sersi/common"
	authorization "github.com/sersi-project/sersi/internal/authorization"
	"github.com/sersi-project/sersi/internal/scaffold"
	"github.com/sersi-project/sersi/pkg"
	"github.com/spf13/cobra"
)

var ScaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Show scaffold of Sersi CLI",
	Long:  `Show scaffold of Sersi CLI`,
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
		fileParser := pkg.NewFileParser(filePath)
		fileParserResult, err := fileParser.ExceuteMapping()
		if err != nil {
			fmt.Println("Error parsing file:", err)
			os.Exit(0)
		}
		err = svc.SaveScaffold(fileParserResult)
		if err != nil {
			fmt.Println("Error saving scaffold:", err)
			os.Exit(1)
		}

		fmt.Println("Scaffold saved successfully")
		os.Exit(0)
	case "list":
		err := svc.GetAllScaffolds()
		if err != nil {
			fmt.Println("Error getting all scaffolds:", err)
			os.Exit(1)
		}
		fmt.Println("All scaffolds:")
		os.Exit(0)
	case "update":
		fmt.Printf("Update scaffold %s", name)
		os.Exit(0)
	case "delete":
		fmt.Printf("Delete scaffold %s", name)
		os.Exit(0)
	default:
		fmt.Println("Invalid action")
		os.Exit(0)
	}
}
