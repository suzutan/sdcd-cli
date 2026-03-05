package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var buildStepsCmd = &cobra.Command{
	Use:   "steps <id>",
	Short: "List steps for a build",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		steps, err := client.GetBuildSteps(id)
		if err != nil {
			return err
		}
		return printer().PrintSteps(steps)
	},
}
