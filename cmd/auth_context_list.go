package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authContextListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, ctx := range cfg.Contexts {
			marker := "  "
			if ctx.Name == cfg.CurrentContext {
				marker = "* "
			}
			fmt.Printf("%s%s\t%s\n", marker, ctx.Name, ctx.APIURL)
		}
		return nil
	},
}
