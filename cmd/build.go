package cmd

import "github.com/spf13/cobra"

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Manage builds",
	Long: `Manage Screwdriver.cd builds.

A Build is a single execution of a Job within an Event. It contains
Steps (the individual shell commands) and produces logs and artifacts.

Get a build ID:
  sdcd event builds <event-id>
  sdcd job builds <job-id>`,
}

func init() {
	buildCmd.AddCommand(buildViewCmd)
	buildCmd.AddCommand(buildStopCmd)
	buildCmd.AddCommand(buildLogsCmd)
	buildCmd.AddCommand(buildStepsCmd)
	buildCmd.AddCommand(buildArtifactsCmd)
	buildCmd.AddCommand(buildArtifactCmd)
	rootCmd.AddCommand(buildCmd)
}
