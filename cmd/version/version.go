package version

import (
	"sersi/common"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of Sersi CLI",
	Long:  `Show version of Sersi CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}
