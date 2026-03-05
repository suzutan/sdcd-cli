package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/config"
)

var authContextRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := config.RemoveContext(cfg, name); err != nil {
			return err
		}
		path := configPath()
		if err := config.Save(path, cfg); err != nil {
			return err
		}
		fmt.Printf("Context %q removed.\n", name)
		return nil
	},
}
