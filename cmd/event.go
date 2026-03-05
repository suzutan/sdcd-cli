package cmd

import "github.com/spf13/cobra"

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Manage events",
	Long: `Manage Screwdriver.cd events.

An Event is created each time a pipeline starts (e.g. on git push or
via 'sdcd pipeline start'). It groups all Builds triggered in that run.

Get an event ID:
  sdcd pipeline events <pipeline-id>`,
}

func init() {
	eventCmd.AddCommand(eventViewCmd)
	eventCmd.AddCommand(eventBuildsCmd)
	eventCmd.AddCommand(eventStopCmd)
	eventCmd.AddCommand(eventRerunCmd)
	rootCmd.AddCommand(eventCmd)
}
