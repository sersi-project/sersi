/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/sersi-project/core/cmd/core/build"
	"github.com/sersi-project/core/cmd/core/create"
	"github.com/sersi-project/core/cmd/core/version"
	hookscmd "github.com/sersi-project/core/cmd/pro/hooks"
	logincmd "github.com/sersi-project/core/cmd/pro/login"
	templatescmd "github.com/sersi-project/core/cmd/pro/templates"
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
	rootCmd.AddCommand(logincmd.LoginCmd)
	rootCmd.AddCommand(templatescmd.TemplatesCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommand()
}
