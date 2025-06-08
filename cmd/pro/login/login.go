package logincmd

import (
	"github.com/sersi-project/sersi/common"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Sersi Pro",
	Long:  `Login to Sersi Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}

func init() {
	LoginCmd.Flags().StringP("email", "e", "", "Email of user")
}
