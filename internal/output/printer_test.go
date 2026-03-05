package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestTablePrinter_PrintPipelines(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	pipelines := []model.Pipeline{
		{ID: 1, Name: "my-pipeline", ScmURI: "github:foo/bar", State: "ACTIVE"},
	}
	if err := p.PrintPipelines(pipelines); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "my-pipeline") {
		t.Errorf("expected 'my-pipeline' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "ACTIVE") {
		t.Errorf("expected 'ACTIVE' in output, got:\n%s", out)
	}
}

func TestJSONPrinter_PrintPipelines(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatJSON, true, &buf)
	pipelines := []model.Pipeline{
		{ID: 42, Name: "json-test"},
	}
	if err := p.PrintPipelines(pipelines); err != nil {
		t.Fatal(err)
	}
	var result []model.Pipeline
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("invalid JSON: %v\noutput:\n%s", err, buf.String())
	}
	if len(result) != 1 || result[0].ID != 42 {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestYAMLPrinter_PrintPipelines(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatYAML, true, &buf)
	pipelines := []model.Pipeline{
		{ID: 7, Name: "yaml-test"},
	}
	if err := p.PrintPipelines(pipelines); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "yaml-test") {
		t.Errorf("expected 'yaml-test' in YAML output, got:\n%s", out)
	}
}

func TestTablePrinter_PrintBuilds(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	builds := []model.Build{
		{ID: 100, JobID: 5, Status: "SUCCESS", SHA: "abc123def456", Number: 10},
	}
	if err := p.PrintBuilds(builds); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "SUCCESS") {
		t.Errorf("expected 'SUCCESS' in output, got:\n%s", out)
	}
}

func TestTablePrinter_PrintSteps(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	code := 0
	steps := []model.Step{
		{Name: "install", Code: &code},
	}
	if err := p.PrintSteps(steps); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "install") {
		t.Errorf("expected 'install' in output, got:\n%s", out)
	}
}
