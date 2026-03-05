package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var jobEnableCmd = &cobra.Command{
	Use:   "enable <id>",
	Short: "Enable a job",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		j, err := client.EnableJob(id)
		if err != nil {
			return err
		}
		fmt.Printf("Job %d enabled (state: %s).\n", j.ID, j.State)
		return nil
	},
}
