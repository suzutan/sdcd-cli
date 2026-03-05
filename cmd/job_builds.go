package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	jbPage  int
	jbCount int
)

var jobBuildsCmd = &cobra.Command{
	Use:   "builds <id>",
	Short: "List builds for a job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		builds, err := client.GetJobBuilds(id, jbPage, jbCount)
		if err != nil {
			return err
		}
		return printer().PrintBuilds(builds)
	},
}

func init() {
	jobBuildsCmd.Flags().IntVar(&jbPage, "page", 0, "page number")
	jobBuildsCmd.Flags().IntVar(&jbCount, "count", 0, "items per page")
}
