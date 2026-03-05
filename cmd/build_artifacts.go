package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var buildArtifactsCmd = &cobra.Command{
	Use:   "artifacts <id>",
	Short: "List artifacts for a build",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		artifacts, err := client.GetBuildArtifacts(id)
		if err != nil {
			return err
		}
		for _, a := range artifacts {
			fmt.Println(a)
		}
		return nil
	},
}
