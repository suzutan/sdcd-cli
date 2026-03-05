package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"github.com/suzutan/sdcd-cli/internal/api"
)

var (
	secretCreatePipelineID int
	secretCreateName       string
	secretCreateValue      string
)

var secretCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a secret",
	RunE: func(cmd *cobra.Command, args []string) error {
		value := secretCreateValue
		if value == "" {
			fmt.Fprint(os.Stderr, "Secret value: ")
			b, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				return fmt.Errorf("read password: %w", err)
			}
			fmt.Fprintln(os.Stderr)
			value = string(b)
		}
		s, err := client.CreateSecret(api.CreateSecretParams{
			PipelineID: secretCreatePipelineID,
			Name:       secretCreateName,
			Value:      value,
		})
		if err != nil {
			return err
		}
		fmt.Printf("Secret %d (%s) created.\n", s.ID, s.Name)
		return nil
	},
}

func init() {
	secretCreateCmd.Flags().IntVar(&secretCreatePipelineID, "pipeline-id", 0, "pipeline ID (required)")
	secretCreateCmd.Flags().StringVar(&secretCreateName, "name", "", "secret name (required)")
	secretCreateCmd.Flags().StringVar(&secretCreateValue, "value", "", "secret value (prompted if omitted)")
	_ = secretCreateCmd.MarkFlagRequired("pipeline-id")
	_ = secretCreateCmd.MarkFlagRequired("name")
}
