package auth

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sersi-project/sersi/common"
	"github.com/sersi-project/sersi/internal/authorization"
	"github.com/sersi-project/sersi/internal/tui/logininput"
	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login to Sersi Pro",
	Long:  `Login to Sersi Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		common.PrintLogo()
	},
}

func init() {
	addLoginFlags(loginCmd)
	AuthCmd.AddCommand(loginCmd)
	AuthCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of authentication",
	Long:  `Show status of authentication`,
	Run:   runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	common.PrintLogo()
	userID, ok := authorization.CheckStatus()
	if ok {
		fmt.Printf("%s You are logged in as %s\n", common.SuccessLabel, lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(userID))
		os.Exit(0)
	}
	fmt.Println(common.ErrorLabel + " You are not logged in")
	os.Exit(0)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Sersi Pro",
	Long:  `Login to Sersi Pro`,
	Run:   runLogin,
}

func addLoginFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("email", "e", "", "Email")
	cmd.Flags().BoolP("bypass-check", "b", false, "Bypass login check")
}

func runLogin(cmd *cobra.Command, args []string) {
	common.PrintLogo()

	bypassCheck, _ := cmd.Flags().GetBool("bypass-check")
	if !bypassCheck {
		userID, ok := authorization.CheckStatus()
		if ok {
			fmt.Printf("%s You are already logged in as %s\n", common.SuccessLabel, lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(userID))
			os.Exit(0)
		}
	}

	var email, password string

	email, _ = cmd.Flags().GetString("email")
	password, _ = cmd.Flags().GetString("password")

	if email == "" || password == "" {
		tprogram := tea.NewProgram(logininput.InitialModel())
		pn, err := tprogram.Run()
		if err != nil {
			fmt.Println(common.ErrorLabel+" Error running program:", err)
			os.Exit(1)
		}
		email = pn.(logininput.Model).Inputs[0].Value()
		password = pn.(logininput.Model).Inputs[1].Value()
	}

	err := authorization.Login(email, password)
	if err != nil {
		fmt.Println(common.ErrorLabel+" Error logging in:", err)
		os.Exit(1)
	}

	fmt.Println(common.SuccessLabel + " Logged in successfully")
}
