package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var (
	pePage  int
	peCount int
)

var pipelineEventsCmd = &cobra.Command{
	Use:   "events <id>",
	Short: "List events for a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		events, err := client.GetPipelineEvents(id, pePage, peCount)
		if err != nil {
			return err
		}
		return printer().PrintEvents(events)
	},
}

func init() {
	pipelineEventsCmd.Flags().IntVar(&pePage, "page", 0, "page number")
	pipelineEventsCmd.Flags().IntVar(&peCount, "count", 0, "items per page")
}
