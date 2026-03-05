package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var eventStopCmd = &cobra.Command{
	Use:   "stop <id>",
	Short: "Stop an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := client.StopEvent(id); err != nil {
			return err
		}
		fmt.Printf("Event %d stopped.\n", id)
		return nil
	},
}
