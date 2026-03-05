package cmd

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var artifactOut string

var buildArtifactCmd = &cobra.Command{
	Use:   "artifact <id> <name>",
	Short: "Download a build artifact file to stdout or a file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		data, err := client.GetBuildArtifact(id, args[1])
		if err != nil {
			return err
		}
		if artifactOut != "" {
			return os.WriteFile(artifactOut, data, 0600)
		}
		_, err = os.Stdout.Write(data)
		return err
	},
}

func init() {
	buildArtifactCmd.Flags().StringVar(&artifactOut, "out", "", "write output to file instead of stdout")
}
