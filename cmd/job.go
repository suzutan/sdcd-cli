package cmd

import "github.com/spf13/cobra"

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Manage jobs",
}

func init() {
	jobCmd.AddCommand(jobViewCmd)
	jobCmd.AddCommand(jobEnableCmd)
	jobCmd.AddCommand(jobDisableCmd)
	jobCmd.AddCommand(jobBuildsCmd)
	jobCmd.AddCommand(jobLatestCmd)
	rootCmd.AddCommand(jobCmd)
}
