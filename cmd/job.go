package cmd

import "github.com/spf13/cobra"

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Manage jobs",
	Long: `Manage Screwdriver.cd jobs.

A Job is a named stage defined in screwdriver.yaml (e.g. "main", "deploy").
Each time a Job runs it creates a Build. Jobs can be enabled or disabled.

Get a job ID:
  sdcd pipeline jobs <pipeline-id>`,
}

func init() {
	jobCmd.AddCommand(jobViewCmd)
	jobCmd.AddCommand(jobEnableCmd)
	jobCmd.AddCommand(jobDisableCmd)
	jobCmd.AddCommand(jobBuildsCmd)
	jobCmd.AddCommand(jobLatestCmd)
	rootCmd.AddCommand(jobCmd)
}
