package config

import (
	"os"
	"testing"
)

func TestLoadDefault_FileNotExist(t *testing.T) {
	cfg, err := Load("/nonexistent/path/config.yaml")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if cfg.Preferences.Output != "table" {
		t.Errorf("expected output=table, got %q", cfg.Preferences.Output)
	}
	if cfg.Preferences.PageSize != 50 {
		t.Errorf("expected page-size=50, got %d", cfg.Preferences.PageSize)
	}
}

func TestSaveAndLoad(t *testing.T) {
	cfg := &Config{
		CurrentContext: "prod",
		Contexts: []Context{
			{Name: "prod", APIURL: "https://api.example.com", Token: "tok"},
		},
		Preferences: Preferences{Output: "json", PageSize: 20},
	}
	path := TempConfig(t, cfg)

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.CurrentContext != "prod" {
		t.Errorf("CurrentContext: got %q", loaded.CurrentContext)
	}
	if len(loaded.Contexts) != 1 {
		t.Fatalf("Contexts len: got %d", len(loaded.Contexts))
	}
	if loaded.Contexts[0].APIURL != "https://api.example.com" {
		t.Errorf("APIURL: got %q", loaded.Contexts[0].APIURL)
	}
}

func TestSave_Permission(t *testing.T) {
	cfg := &Config{}
	path := TempConfig(t, cfg)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("expected 0600, got %o", info.Mode().Perm())
	}
}

func TestActiveContext(t *testing.T) {
	cfg := &Config{
		CurrentContext: "staging",
		Contexts: []Context{
			{Name: "prod", APIURL: "https://prod.example.com", Token: "a"},
			{Name: "staging", APIURL: "https://staging.example.com", Token: "b"},
		},
	}
	ctx, err := ActiveContext(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if ctx.Name != "staging" {
		t.Errorf("expected staging, got %q", ctx.Name)
	}
}

func TestActiveContext_NotFound(t *testing.T) {
	cfg := &Config{CurrentContext: "missing"}
	_, err := ActiveContext(cfg)
	if err == nil {
		t.Error("expected error")
	}
}

func TestAddContext(t *testing.T) {
	cfg := &Config{}
	err := AddContext(cfg, Context{Name: "new", APIURL: "https://new.example.com", Token: "t"})
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Contexts) != 1 {
		t.Errorf("expected 1 context, got %d", len(cfg.Contexts))
	}
}

func TestAddContext_Duplicate(t *testing.T) {
	cfg := &Config{
		Contexts: []Context{{Name: "dup", APIURL: "u", Token: "t"}},
	}
	err := AddContext(cfg, Context{Name: "dup"})
	if err == nil {
		t.Error("expected error for duplicate")
	}
}

func TestRemoveContext(t *testing.T) {
	cfg := &Config{
		CurrentContext: "a",
		Contexts: []Context{
			{Name: "a"}, {Name: "b"},
		},
	}
	if err := RemoveContext(cfg, "a"); err != nil {
		t.Fatal(err)
	}
	if len(cfg.Contexts) != 1 {
		t.Errorf("expected 1 context, got %d", len(cfg.Contexts))
	}
	if cfg.CurrentContext != "" {
		t.Errorf("expected empty current-context after removing active context")
	}
}

func TestRemoveContext_NotFound(t *testing.T) {
	cfg := &Config{}
	if err := RemoveContext(cfg, "nonexistent"); err == nil {
		t.Error("expected error")
	}
}

func TestUseContext(t *testing.T) {
	cfg := &Config{
		Contexts: []Context{{Name: "prod"}, {Name: "staging"}},
	}
	if err := UseContext(cfg, "staging"); err != nil {
		t.Fatal(err)
	}
	if cfg.CurrentContext != "staging" {
		t.Errorf("expected staging, got %q", cfg.CurrentContext)
	}
}

func TestUseContext_NotFound(t *testing.T) {
	cfg := &Config{}
	if err := UseContext(cfg, "nonexistent"); err == nil {
		t.Error("expected error")
	}
}
