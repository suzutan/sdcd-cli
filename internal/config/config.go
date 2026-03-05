package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DefaultPath returns the default config file path.
func DefaultPath() string {
	base := os.Getenv("XDG_CONFIG_HOME")
	if base == "" {
		home, _ := os.UserHomeDir()
		base = filepath.Join(home, ".config")
	}
	return filepath.Join(base, "sdcd-cli", "config.yaml")
}

// Load reads and parses the config file at path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultConfig(), nil
		}
		return nil, fmt.Errorf("config read error: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("config parse error: %w", err)
	}
	if cfg.Preferences.Output == "" {
		cfg.Preferences.Output = "table"
	}
	if cfg.Preferences.PageSize == 0 {
		cfg.Preferences.PageSize = 50
	}
	return &cfg, nil
}

// Save writes the config to path with permission 0600.
func Save(path string, cfg *Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("config dir create error: %w", err)
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("config marshal error: %w", err)
	}
	return os.WriteFile(path, data, 0600)
}

// ActiveContext returns the Context that matches cfg.CurrentContext.
func ActiveContext(cfg *Config) (*Context, error) {
	for i := range cfg.Contexts {
		if cfg.Contexts[i].Name == cfg.CurrentContext {
			return &cfg.Contexts[i], nil
		}
	}
	return nil, fmt.Errorf("context %q not found", cfg.CurrentContext)
}

// AddContext adds a new context; returns error if name already exists.
func AddContext(cfg *Config, ctx Context) error {
	for _, c := range cfg.Contexts {
		if c.Name == ctx.Name {
			return fmt.Errorf("context %q already exists", ctx.Name)
		}
	}
	cfg.Contexts = append(cfg.Contexts, ctx)
	return nil
}

// RemoveContext removes the named context; returns error if not found.
func RemoveContext(cfg *Config, name string) error {
	for i, c := range cfg.Contexts {
		if c.Name == name {
			cfg.Contexts = append(cfg.Contexts[:i], cfg.Contexts[i+1:]...)
			if cfg.CurrentContext == name {
				cfg.CurrentContext = ""
			}
			return nil
		}
	}
	return fmt.Errorf("context %q not found", name)
}

// UseContext sets cfg.CurrentContext to name; returns error if not found.
func UseContext(cfg *Config, name string) error {
	for _, c := range cfg.Contexts {
		if c.Name == name {
			cfg.CurrentContext = name
			return nil
		}
	}
	return fmt.Errorf("context %q not found", name)
}

func defaultConfig() *Config {
	return &Config{
		Preferences: Preferences{
			Output:   "table",
			NoColor:  false,
			PageSize: 50,
		},
	}
}
