package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var eventViewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		e, err := client.GetEvent(id)
		if err != nil {
			return err
		}
		return printer().PrintEvent(*e)
	},
}
