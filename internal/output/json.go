package output

import (
	"encoding/json"
	"io"

	"github.com/suzutan/sdcd-cli/internal/model"
)

type jsonPrinter struct {
	w io.Writer
}

func (p *jsonPrinter) encode(v interface{}) error {
	enc := json.NewEncoder(p.w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func (p *jsonPrinter) PrintPipelines(v []model.Pipeline) error  { return p.encode(v) }
func (p *jsonPrinter) PrintPipeline(v model.Pipeline) error     { return p.encode(v) }
func (p *jsonPrinter) PrintJobs(v []model.Job) error             { return p.encode(v) }
func (p *jsonPrinter) PrintJob(v model.Job) error                { return p.encode(v) }
func (p *jsonPrinter) PrintBuilds(v []model.Build) error         { return p.encode(v) }
func (p *jsonPrinter) PrintBuild(v model.Build) error            { return p.encode(v) }
func (p *jsonPrinter) PrintSteps(v []model.Step) error           { return p.encode(v) }
func (p *jsonPrinter) PrintEvents(v []model.Event) error         { return p.encode(v) }
func (p *jsonPrinter) PrintEvent(v model.Event) error            { return p.encode(v) }
func (p *jsonPrinter) PrintSecrets(v []model.Secret) error       { return p.encode(v) }
