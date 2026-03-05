package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

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
		t.Errorf("expected 'my-pipeline' in output:\n%s", out)
	}
	if !strings.Contains(out, "ACTIVE") {
		t.Errorf("expected 'ACTIVE' in output:\n%s", out)
	}
}

func TestTablePrinter_PrintPipeline(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	now := time.Now()
	pl := model.Pipeline{ID: 5, Name: "single", ScmContext: "github", State: "ACTIVE", LastEventID: 10, CreateTime: &now}
	if err := p.PrintPipeline(pl); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "single") {
		t.Errorf("expected 'single' in output:\n%s", out)
	}
}

func TestTablePrinter_PrintJobs(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	jobs := []model.Job{
		{ID: 1, Name: "main", State: "ENABLED", Archived: false},
		{ID: 2, Name: "deploy", State: "DISABLED", Archived: true},
	}
	if err := p.PrintJobs(jobs); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "main") || !strings.Contains(out, "deploy") {
		t.Errorf("expected job names in output:\n%s", out)
	}
}

func TestTablePrinter_PrintJob(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	now := time.Now()
	j := model.Job{ID: 3, PipelineID: 1, Name: "test", State: "ENABLED", CreateTime: &now}
	if err := p.PrintJob(j); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "test") {
		t.Errorf("expected 'test' in output:\n%s", out)
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
		t.Errorf("expected 'SUCCESS' in output:\n%s", out)
	}
	// SHA is truncated to 8 chars
	if !strings.Contains(out, "abc123de") {
		t.Errorf("expected truncated SHA in output:\n%s", out)
	}
}

func TestTablePrinter_PrintBuild(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	now := time.Now()
	b := model.Build{ID: 1, JobID: 2, EventID: 3, Status: "RUNNING", SHA: "deadbeef", Number: 5, CreateTime: &now, StartTime: &now}
	if err := p.PrintBuild(b); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "RUNNING") {
		t.Errorf("expected 'RUNNING' in output:\n%s", out)
	}
}

func TestTablePrinter_PrintSteps(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	code := 0
	now := time.Now()
	steps := []model.Step{
		{Name: "install", Code: &code, StartTime: &now, EndTime: &now},
		{Name: "test"},
	}
	if err := p.PrintSteps(steps); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "install") {
		t.Errorf("expected 'install' in output:\n%s", out)
	}
}

func TestTablePrinter_PrintEvents(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	events := []model.Event{
		{ID: 1, PipelineID: 10, Status: "SUCCESS", SHA: "abcdef123456", Type: "pipeline", Creator: model.EventCreator{Username: "alice"}},
	}
	if err := p.PrintEvents(events); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "alice") {
		t.Errorf("expected 'alice' in output:\n%s", out)
	}
}

func TestTablePrinter_PrintEvent(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	now := time.Now()
	parentID := 9
	e := model.Event{ID: 10, PipelineID: 1, Status: "RUNNING", SHA: "abc", Type: "pipeline",
		Creator: model.EventCreator{Username: "bob"}, CreateTime: &now, ParentEventID: &parentID}
	if err := p.PrintEvent(e); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "bob") {
		t.Errorf("expected 'bob' in output:\n%s", out)
	}
	if !strings.Contains(out, "9") {
		t.Errorf("expected parent event ID in output:\n%s", out)
	}
}

func TestTablePrinter_PrintSecrets(t *testing.T) {
	var buf bytes.Buffer
	p := NewPrinter(FormatTable, true, &buf)
	secrets := []model.Secret{
		{ID: 1, Name: "DB_PASSWORD", PipelineID: 5, AllowInPR: false},
	}
	if err := p.PrintSecrets(secrets); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "DB_PASSWORD") {
		t.Errorf("expected 'DB_PASSWORD' in output:\n%s", out)
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

func TestJSONPrinter_AllMethods(t *testing.T) {
	now := time.Now()
	code := 0

	tests := []struct {
		name string
		fn   func(p Printer) error
	}{
		{"PrintPipeline", func(p Printer) error { return p.PrintPipeline(model.Pipeline{ID: 1}) }},
		{"PrintJobs", func(p Printer) error { return p.PrintJobs([]model.Job{{ID: 1}}) }},
		{"PrintJob", func(p Printer) error { return p.PrintJob(model.Job{ID: 1}) }},
		{"PrintBuilds", func(p Printer) error { return p.PrintBuilds([]model.Build{{ID: 1}}) }},
		{"PrintBuild", func(p Printer) error { return p.PrintBuild(model.Build{ID: 1, CreateTime: &now}) }},
		{"PrintSteps", func(p Printer) error { return p.PrintSteps([]model.Step{{Name: "s", Code: &code}}) }},
		{"PrintEvents", func(p Printer) error { return p.PrintEvents([]model.Event{{ID: 1}}) }},
		{"PrintEvent", func(p Printer) error { return p.PrintEvent(model.Event{ID: 1}) }},
		{"PrintSecrets", func(p Printer) error { return p.PrintSecrets([]model.Secret{{ID: 1}}) }},
	}

	for _, format := range []Format{FormatJSON, FormatYAML} {
		for _, tt := range tests {
			t.Run(string(format)+"/"+tt.name, func(t *testing.T) {
				var buf bytes.Buffer
				pr := NewPrinter(format, true, &buf)
				if err := tt.fn(pr); err != nil {
					t.Errorf("%s/%s: %v", format, tt.name, err)
				}
				if buf.Len() == 0 {
					t.Errorf("%s/%s: empty output", format, tt.name)
				}
			})
		}
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
		t.Errorf("expected 'yaml-test' in YAML output:\n%s", out)
	}
}

func TestColorizeStatus(t *testing.T) {
	tests := []struct {
		status  string
		noColor bool
	}{
		{"RUNNING", false},
		{"SUCCESS", false},
		{"FAILURE", false},
		{"ABORTED", false},
		{"QUEUED", false},
		{"BLOCKED", false},
		{"UNKNOWN", false},
		{"SUCCESS", true}, // noColor=true: output must equal input
	}
	for _, tt := range tests {
		out := ColorizeStatus(tt.status, tt.noColor)
		if tt.noColor && out != tt.status {
			t.Errorf("noColor=true: expected %q, got %q", tt.status, out)
		}
		if !tt.noColor && !strings.Contains(out, tt.status) {
			t.Errorf("expected status %q embedded in colored output %q", tt.status, out)
		}
	}
}
