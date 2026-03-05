package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TempConfig writes cfg to a temporary file and returns the path.
// The file is cleaned up automatically via t.Cleanup.
func TempConfig(t *testing.T, cfg *Config) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := Save(path, cfg); err != nil {
		t.Fatalf("TempConfig: %v", err)
	}
	return path
}

// TempConfigFile writes raw YAML bytes to a temp file and returns path.
func TempConfigFile(t *testing.T, content []byte) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatalf("TempConfigFile: %v", err)
	}
	return path
}
