package cmd

import "github.com/spf13/cobra"

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage contexts",
}

func init() {
	contextCmd.AddCommand(contextAddCmd)
	contextCmd.AddCommand(contextRemoveCmd)
	contextCmd.AddCommand(contextListCmd)
	contextCmd.AddCommand(contextUseCmd)
	contextCmd.AddCommand(contextCurrentCmd)
	rootCmd.AddCommand(contextCmd)
}
