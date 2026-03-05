package output

import (
	"io"

	"gopkg.in/yaml.v3"

	"github.com/suzutan/sdcd-cli/internal/model"
)

type yamlPrinter struct {
	w io.Writer
}

func (p *yamlPrinter) encode(v interface{}) error {
	return yaml.NewEncoder(p.w).Encode(v)
}

func (p *yamlPrinter) PrintPipelines(v []model.Pipeline) error  { return p.encode(v) }
func (p *yamlPrinter) PrintPipeline(v model.Pipeline) error     { return p.encode(v) }
func (p *yamlPrinter) PrintJobs(v []model.Job) error             { return p.encode(v) }
func (p *yamlPrinter) PrintJob(v model.Job) error                { return p.encode(v) }
func (p *yamlPrinter) PrintBuilds(v []model.Build) error         { return p.encode(v) }
func (p *yamlPrinter) PrintBuild(v model.Build) error            { return p.encode(v) }
func (p *yamlPrinter) PrintSteps(v []model.Step) error           { return p.encode(v) }
func (p *yamlPrinter) PrintEvents(v []model.Event) error         { return p.encode(v) }
func (p *yamlPrinter) PrintEvent(v model.Event) error            { return p.encode(v) }
func (p *yamlPrinter) PrintSecrets(v []model.Secret) error       { return p.encode(v) }
