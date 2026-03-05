package cmd

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"github.com/suzutan/sdcd-cli/internal/api"
)

var (
	secretUpdateValue     string
	secretUpdateAllowInPR bool
)

var secretUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		p := api.UpdateSecretParams{}
		if cmd.Flags().Changed("value") {
			p.Value = &secretUpdateValue
		} else if !cmd.Flags().Changed("allow-in-pr") {
			// no value flag: prompt for it
			fmt.Fprint(os.Stderr, "Secret value: ")
			b, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				return fmt.Errorf("read password: %w", err)
			}
			fmt.Fprintln(os.Stderr)
			v := string(b)
			p.Value = &v
		}
		if cmd.Flags().Changed("allow-in-pr") {
			p.AllowInPR = &secretUpdateAllowInPR
		}
		s, err := client.UpdateSecret(id, p)
		if err != nil {
			return err
		}
		fmt.Printf("Secret %d (%s) updated.\n", s.ID, s.Name)
		return nil
	},
}

func init() {
	secretUpdateCmd.Flags().StringVar(&secretUpdateValue, "value", "", "new secret value (prompted if omitted)")
	secretUpdateCmd.Flags().BoolVar(&secretUpdateAllowInPR, "allow-in-pr", false, "allow secret in PR builds")
}
