package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/config"
)

var (
	addAPIURL string
	addToken  string
)

var authContextAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a new context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ctx := config.Context{
			Name:   name,
			APIURL: addAPIURL,
			Token:  addToken,
		}
		if err := config.AddContext(cfg, ctx); err != nil {
			return err
		}
		path := configPath()
		if err := config.Save(path, cfg); err != nil {
			return err
		}
		fmt.Printf("Context %q added.\n", name)
		return nil
	},
}

func init() {
	authContextAddCmd.Flags().StringVar(&addAPIURL, "api-url", "", "API URL (required)")
	authContextAddCmd.Flags().StringVar(&addToken, "token", "", "API token (required)")
	_ = authContextAddCmd.MarkFlagRequired("api-url")
	_ = authContextAddCmd.MarkFlagRequired("token")
}
