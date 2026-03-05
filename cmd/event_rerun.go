package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var rerunJob string

var eventRerunCmd = &cobra.Command{
	Use:   "rerun <id>",
	Short: "Rerun an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		e, err := client.RerunEvent(id, rerunJob)
		if err != nil {
			return err
		}
		return printer().PrintEvent(*e)
	},
}

func init() {
	eventRerunCmd.Flags().StringVar(&rerunJob, "job", "", "job name to start from")
}
