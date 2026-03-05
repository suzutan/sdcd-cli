package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suzutan/sdcd-cli/internal/api"
	"github.com/suzutan/sdcd-cli/internal/config"
	"github.com/suzutan/sdcd-cli/internal/output"
)

var (
	cfgFile    string
	ctxFlag    string
	outputFlag string
	noColor    bool

	cfg    *config.Config
	client *api.Client
)

var rootCmd = &cobra.Command{
	Use:   "sdcd",
	Short: "Screwdriver.cd CLI",
	Long:  "sdcd is a CLI tool to interact with Screwdriver.cd instances.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if noClientNeeded(cmd) {
			return nil
		}
		return initClient()
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $XDG_CONFIG_HOME/sdcd-cli/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&ctxFlag, "context", "", "context override")
	rootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "", "output format: table|json|yaml")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable ANSI colors")

	// Shell completion is provided by cobra's built-in completion command.
	rootCmd.CompletionOptions.DisableDefaultCmd = false
}

func initConfig() {
	path := cfgFile
	if path == "" {
		path = config.DefaultPath()
	}
	var err error
	cfg, err = config.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func initClient() error {
	ctx, err := activeContext()
	if err != nil {
		return err
	}
	client = api.NewClient(ctx)
	return nil
}

func activeContext() (*config.Context, error) {
	if ctxFlag != "" {
		for i := range cfg.Contexts {
			if cfg.Contexts[i].Name == ctxFlag {
				return &cfg.Contexts[i], nil
			}
		}
		return nil, fmt.Errorf("context %q not found", ctxFlag)
	}
	return config.ActiveContext(cfg)
}

func printer() output.Printer {
	format := output.Format(outputFlag)
	if format == "" {
		format = output.Format(cfg.Preferences.Output)
	}
	if format == "" {
		format = output.FormatTable
	}
	nc := noColor || cfg.Preferences.NoColor
	return output.NewPrinter(format, nc, os.Stdout)
}

func configPath() string {
	if cfgFile != "" {
		return cfgFile
	}
	return config.DefaultPath()
}

// noClientNeeded returns true for commands that don't need an API client.
func noClientNeeded(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "version", "completion", "help":
		return true
	}
	c := cmd
	for c != nil {
		if c.Name() == "context" {
			return true
		}
		c = c.Parent()
	}
	return false
}
