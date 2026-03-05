package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var buildGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a build",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		b, err := client.GetBuild(id)
		if err != nil {
			return err
		}
		return printer().PrintBuild(*b)
	},
}
