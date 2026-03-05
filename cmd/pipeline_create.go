package cmd

import (
	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/api"
)

var (
	createCheckoutURL string
	createRootDir     string
)

var pipelineCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a pipeline",
	RunE: func(cmd *cobra.Command, args []string) error {
		pl, err := client.CreatePipeline(api.CreatePipelineParams{
			CheckoutURL: createCheckoutURL,
			RootDir:     createRootDir,
		})
		if err != nil {
			return err
		}
		return printer().PrintPipeline(*pl)
	},
}

func init() {
	pipelineCreateCmd.Flags().StringVar(&createCheckoutURL, "checkout-url", "", "checkout URL (required)")
	pipelineCreateCmd.Flags().StringVar(&createRootDir, "root-dir", "", "root directory")
	_ = pipelineCreateCmd.MarkFlagRequired("checkout-url")
}
