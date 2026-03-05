package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var jobViewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View a job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		j, err := client.GetJob(id)
		if err != nil {
			return err
		}
		return printer().PrintJob(*j)
	},
}
