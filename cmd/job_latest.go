package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var jobLatestCmd = &cobra.Command{
	Use:   "latest-build <id>",
	Short: "View the latest build for a job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		b, err := client.GetLatestBuild(id)
		if err != nil {
			return err
		}
		return printer().PrintBuild(*b)
	},
}
