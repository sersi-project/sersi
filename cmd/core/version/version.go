package version     

import (
	"github.com/sersi-project/core/common"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of Sersi CLI",
	Long:  `Show version of Sersi CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}
