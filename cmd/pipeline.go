package cmd

import "github.com/spf13/cobra"

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Manage pipelines",
	Long: `Manage Screwdriver.cd pipelines.

A Pipeline is the top-level resource that defines your CI/CD configuration
(screwdriver.yaml). It contains Jobs and produces Events when started.

Resource hierarchy:
  Pipeline
    ├─ Job   (defined in screwdriver.yaml)
    └─ Event (created when a pipeline starts)
         └─ Build (one per Job in the Event)

Get a pipeline ID:
  sdcd pipeline list`,
}

func init() {
	pipelineCmd.AddCommand(pipelineListCmd)
	pipelineCmd.AddCommand(pipelineViewCmd)
	pipelineCmd.AddCommand(pipelineCreateCmd)
	pipelineCmd.AddCommand(pipelineDeleteCmd)
	pipelineCmd.AddCommand(pipelineSyncCmd)
	pipelineCmd.AddCommand(pipelineJobsCmd)
	pipelineCmd.AddCommand(pipelineEventsCmd)
	pipelineCmd.AddCommand(pipelineBuildsCmd)
	pipelineCmd.AddCommand(pipelineStartCmd)
	rootCmd.AddCommand(pipelineCmd)
}
