package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authContextCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current context",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfg.CurrentContext == "" {
			fmt.Println("(none)")
			return nil
		}
		fmt.Println(cfg.CurrentContext)
		return nil
	},
}
