/*
Copyright © 2025 SERSI
*/
package cmd

import (
	"os"

	"github.com/sersi-project/sersi/cmd/core/build"
	"github.com/sersi-project/sersi/cmd/core/create"
	"github.com/sersi-project/sersi/cmd/core/version"
	"github.com/sersi-project/sersi/cmd/pro/auth"
	hookscmd "github.com/sersi-project/sersi/cmd/pro/hooks"
	scaffoldcmd "github.com/sersi-project/sersi/cmd/pro/scaffolds"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sersi",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiplr line`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommand() {
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(build.BuildCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(hookscmd.HooksCmd)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(scaffoldcmd.ScaffoldCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommand()
}
