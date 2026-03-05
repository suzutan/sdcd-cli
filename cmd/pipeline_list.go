package cmd

import (
	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/api"
)

var (
	plSearch string
	plPage   int
	plCount  int
)

var pipelineListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pipelines",
	RunE: func(cmd *cobra.Command, args []string) error {
		pipelines, err := client.ListPipelines(api.PipelineListParams{
			Search: plSearch,
			Page:   plPage,
			Count:  plCount,
		})
		if err != nil {
			return err
		}
		return printer().PrintPipelines(pipelines)
	},
}

func init() {
	pipelineListCmd.Flags().StringVar(&plSearch, "search", "", "search string")
	pipelineListCmd.Flags().IntVar(&plPage, "page", 0, "page number")
	pipelineListCmd.Flags().IntVar(&plCount, "count", 0, "number of items per page")
}
