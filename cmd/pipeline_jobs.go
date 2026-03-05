package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	pjPage  int
	pjCount int
)

var pipelineJobsCmd = &cobra.Command{
	Use:   "jobs <id>",
	Short: "List jobs for a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		jobs, err := client.GetPipelineJobs(id, pjPage, pjCount)
		if err != nil {
			return err
		}
		return printer().PrintJobs(jobs)
	},
}

func init() {
	pipelineJobsCmd.Flags().IntVar(&pjPage, "page", 0, "page number")
	pipelineJobsCmd.Flags().IntVar(&pjCount, "count", 0, "items per page")
}
