package cmd

import (
	"github.com/spf13/cobra"
)

var secretListPipelineID int

var secretListCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets for a pipeline",
	RunE: func(cmd *cobra.Command, args []string) error {
		secrets, err := client.ListSecrets(secretListPipelineID)
		if err != nil {
			return err
		}
		return printer().PrintSecrets(secrets)
	},
}

func init() {
	secretListCmd.Flags().IntVar(&secretListPipelineID, "pipeline-id", 0, "pipeline ID (required)")
	_ = secretListCmd.MarkFlagRequired("pipeline-id")
}
