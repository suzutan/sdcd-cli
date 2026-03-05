package cmd

import "github.com/spf13/cobra"

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Manage pipelines",
}

func init() {
	pipelineCmd.AddCommand(pipelineListCmd)
	pipelineCmd.AddCommand(pipelineGetCmd)
	pipelineCmd.AddCommand(pipelineCreateCmd)
	pipelineCmd.AddCommand(pipelineDeleteCmd)
	pipelineCmd.AddCommand(pipelineSyncCmd)
	pipelineCmd.AddCommand(pipelineJobsCmd)
	pipelineCmd.AddCommand(pipelineEventsCmd)
	pipelineCmd.AddCommand(pipelineBuildsCmd)
	pipelineCmd.AddCommand(pipelineStartCmd)
	rootCmd.AddCommand(pipelineCmd)
}
