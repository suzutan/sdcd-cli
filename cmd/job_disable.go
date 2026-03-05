package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var jobDisableCmd = &cobra.Command{
	Use:   "disable <id>",
	Short: "Disable a job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		j, err := client.DisableJob(id)
		if err != nil {
			return err
		}
		fmt.Printf("Job %d disabled (state: %s).\n", j.ID, j.State)
		return nil
	},
}
