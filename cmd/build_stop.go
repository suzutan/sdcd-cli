package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var buildStopCmd = &cobra.Command{
	Use:   "stop <id>",
	Short: "Stop (abort) a build",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		b, err := client.StopBuild(id)
		if err != nil {
			return err
		}
		fmt.Printf("Build %d stopped (status: %s).\n", b.ID, b.Status)
		return nil
	},
}
