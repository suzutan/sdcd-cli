package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var pipelineGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		pl, err := client.GetPipeline(id)
		if err != nil {
			return err
		}
		return printer().PrintPipeline(*pl)
	},
}
