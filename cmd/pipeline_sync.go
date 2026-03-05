package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var pipelineSyncCmd = &cobra.Command{
	Use:   "sync <id>",
	Short: "Sync a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := client.SyncPipeline(id); err != nil {
			return err
		}
		fmt.Printf("Pipeline %d synced.\n", id)
		return nil
	},
}
