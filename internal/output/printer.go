package output

import (
	"io"

	"github.com/suzutan/sdcd-cli/internal/model"
)

// Format represents output format.
type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

// Printer renders model objects to an io.Writer.
type Printer interface {
	PrintPipelines([]model.Pipeline) error
	PrintPipeline(model.Pipeline) error
	PrintJobs([]model.Job) error
	PrintJob(model.Job) error
	PrintBuilds([]model.Build) error
	PrintBuild(model.Build) error
	PrintSteps([]model.Step) error
	PrintEvents([]model.Event) error
	PrintEvent(model.Event) error
	PrintSecrets([]model.Secret) error
}

// NewPrinter creates a Printer for the given format.
func NewPrinter(format Format, noColor bool, w io.Writer) Printer {
	switch format {
	case FormatJSON:
		return &jsonPrinter{w: w}
	case FormatYAML:
		return &yamlPrinter{w: w}
	default:
		return &tablePrinter{w: w, noColor: noColor}
	}
}
