package create

import (
	"fmt"

	"github.com/sersi-project/core/common"

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
}

func RunCreate(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	fmt.Println("◉ Creating a new full stack project...")

	customSetup, _ = cmd.Flags().GetBool("custom")
	if customSetup {
		fmt.Println("◉ Custom setup enabled")
	} else {
		fmt.Println("◉ Custom setup disabled")
	}
}
