package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var logsStep string

var buildLogsCmd = &cobra.Command{
	Use:   "logs <id>",
	Short: "Show logs for a build step",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if logsStep == "" {
			return fmt.Errorf("--step is required")
		}
		lines, err := client.GetAllBuildLogs(id, logsStep)
		if err != nil {
			return err
		}
		for _, line := range lines {
			ts := time.Unix(0, line.T*int64(time.Millisecond)).UTC().Format("15:04:05")
			fmt.Printf("[%s] %s\n", ts, line.M)
		}
		return nil
	},
}

func init() {
	buildLogsCmd.Flags().StringVar(&logsStep, "step", "", "step name (required)")
}
