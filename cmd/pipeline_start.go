package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/api"
)

var (
	startJob string
	startSHA string
)

var pipelineStartCmd = &cobra.Command{
	Use:   "start <id>",
	Short: "Start a pipeline (create an event)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		p := api.StartPipelineParams{
			PipelineID: id,
			SHA:        startSHA,
		}
		if startJob != "" {
			p.StartFrom = startJob
		}
		event, err := client.StartPipeline(p)
		if err != nil {
			return err
		}
		return printer().PrintEvent(*event)
	},
}

func init() {
	pipelineStartCmd.Flags().StringVar(&startJob, "job", "", "job name to start from")
	pipelineStartCmd.Flags().StringVar(&startSHA, "sha", "", "commit SHA")
}
