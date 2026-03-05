package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	pbPage  int
	pbCount int
)

var pipelineBuildsCmd = &cobra.Command{
	Use:   "builds <id>",
	Short: "List builds for a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		builds, err := client.GetPipelineBuilds(id, pbPage, pbCount)
		if err != nil {
			return err
		}
		return printer().PrintBuilds(builds)
	},
}

func init() {
	pipelineBuildsCmd.Flags().IntVar(&pbPage, "page", 0, "page number")
	pipelineBuildsCmd.Flags().IntVar(&pbCount, "count", 0, "items per page")
}
