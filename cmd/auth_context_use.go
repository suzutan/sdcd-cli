package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/config"
)

var authContextUseCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Switch to a context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := config.UseContext(cfg, name); err != nil {
			return err
		}
		path := configPath()
		if err := config.Save(path, cfg); err != nil {
			return err
		}
		fmt.Printf("Switched to context %q.\n", name)
		return nil
	},
}
