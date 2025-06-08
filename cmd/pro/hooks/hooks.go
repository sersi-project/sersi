package hookscmd

import (
	"github.com/sersi-project/sersi/common"
	"github.com/spf13/cobra"
)

type action string

const (
	actionSave   action = "save"
	actionView   action = "view"
	actionUpdate action = "update"
	actionDelete action = "delete"
)

var HooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Show hooks of Sersi CLI",
	Long:  `Show hooks of Sersi CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}

func init() {
	HooksCmd.Flags().StringP("name", "n", "", "Name of project")
	HooksCmd.Flags().StringP("file-path", "f", "", "File path of project")
	HooksCmd.Flags().StringP("action", "a", "", "Action to perform (save, view, update, delete)")
}
