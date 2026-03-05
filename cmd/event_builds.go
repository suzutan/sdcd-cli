package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var eventBuildsCmd = &cobra.Command{
	Use:   "builds <id>",
	Short: "List builds for an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		builds, err := client.GetEventBuilds(id)
		if err != nil {
			return err
		}
		return printer().PrintBuilds(builds)
	},
}
