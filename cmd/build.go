package cmd

import "github.com/spf13/cobra"

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Manage builds",
}

func init() {
	buildCmd.AddCommand(buildGetCmd)
	buildCmd.AddCommand(buildStopCmd)
	buildCmd.AddCommand(buildLogsCmd)
	buildCmd.AddCommand(buildStepsCmd)
	buildCmd.AddCommand(buildArtifactsCmd)
	buildCmd.AddCommand(buildArtifactCmd)
	rootCmd.AddCommand(buildCmd)
}
