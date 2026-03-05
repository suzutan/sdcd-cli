package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var secretDeleteYes bool

var secretDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if !secretDeleteYes {
			fmt.Printf("Delete secret %d? [y/N] ", id)
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(line)) != "y" {
				fmt.Println("Aborted.")
				return nil
			}
		}
		if err := client.DeleteSecret(id); err != nil {
			return err
		}
		fmt.Printf("Secret %d deleted.\n", id)
		return nil
	},
}

func init() {
	secretDeleteCmd.Flags().BoolVarP(&secretDeleteYes, "yes", "y", false, "skip confirmation")
}
