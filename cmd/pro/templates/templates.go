package templatescmd

import (
	"github.com/sersi-project/sersi/common"
	"github.com/spf13/cobra"
)

var TemplatesCmd = &cobra.Command{
	Use:   "template",
	Short: "Show template of Sersi CLI",
	Long:  `Show template of Sersi CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}

func init() {
	TemplatesCmd.Flags().StringP("name", "n", "", "Name of project")
}
