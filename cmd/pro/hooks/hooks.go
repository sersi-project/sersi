package hooks

import (
	"github.com/sersi-project/sersi/common"
	"github.com/spf13/cobra"
)

var HooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Custom hooks store actions for Sersi Pro (create, view, update, delete, use)",
	Long:  `Custom hooks store actions for Sersi Pro (create, view, update, delete, use)`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}

func init() {
	HooksCmd.Flags().StringP("name", "n", "", "Name of project")
	HooksCmd.Flags().StringP("file-path", "f", "", "File path of project")
	HooksCmd.Flags().StringP("action", "a", "", "Action to perform (create, view, update, delete)")
}
