package cmd

import "github.com/spf13/cobra"

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Manage events",
}

func init() {
	eventCmd.AddCommand(eventViewCmd)
	eventCmd.AddCommand(eventBuildsCmd)
	eventCmd.AddCommand(eventStopCmd)
	eventCmd.AddCommand(eventRerunCmd)
	rootCmd.AddCommand(eventCmd)
}
