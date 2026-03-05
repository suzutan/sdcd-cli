package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var pipelineDeleteYes bool

var pipelineDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if !pipelineDeleteYes {
			fmt.Printf("Delete pipeline %d? [y/N] ", id)
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(line)) != "y" {
				fmt.Println("Aborted.")
				return nil
			}
		}
		if err := client.DeletePipeline(id); err != nil {
			return err
		}
		fmt.Printf("Pipeline %d deleted.\n", id)
		return nil
	},
}

func init() {
	pipelineDeleteCmd.Flags().BoolVarP(&pipelineDeleteYes, "yes", "y", false, "skip confirmation")
}
