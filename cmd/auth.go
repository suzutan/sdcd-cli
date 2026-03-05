package cmd

import "github.com/spf13/cobra"

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication and context management",
}

var authContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage contexts",
}

func init() {
	authContextCmd.AddCommand(authContextAddCmd)
	authContextCmd.AddCommand(authContextRemoveCmd)
	authContextCmd.AddCommand(authContextListCmd)
	authContextCmd.AddCommand(authContextUseCmd)
	authContextCmd.AddCommand(authContextCurrentCmd)
	authCmd.AddCommand(authContextCmd)
	rootCmd.AddCommand(authCmd)
}
