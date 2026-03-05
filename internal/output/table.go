package output

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/suzutan/sdcd-cli/internal/model"
)

type tablePrinter struct {
	w       io.Writer
	noColor bool
}

func (p *tablePrinter) newTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(p.w)
	t.SetStyle(table.StyleLight)
	return t
}

func fmtTime(tp interface{ String() string }) string {
	if tp == nil {
		return ""
	}
	return tp.String()
}

func (p *tablePrinter) PrintPipelines(pipelines []model.Pipeline) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"ID", "NAME", "SCM URI", "STATE"})
	for _, pl := range pipelines {
		t.AppendRow(table.Row{pl.ID, pl.Name, pl.ScmURI, ColorizeStatus(pl.State, p.noColor)})
	}
	t.Render()
	return nil
}

func (p *tablePrinter) PrintPipeline(pl model.Pipeline) error {
	t := p.newTable()
	t.AppendRow(table.Row{"ID", pl.ID})
	t.AppendRow(table.Row{"Name", pl.Name})
	t.AppendRow(table.Row{"SCM URI", pl.ScmURI})
	t.AppendRow(table.Row{"SCM Context", pl.ScmContext})
	t.AppendRow(table.Row{"State", ColorizeStatus(pl.State, p.noColor)})
	t.AppendRow(table.Row{"Last Event ID", pl.LastEventID})
	if pl.CreateTime != nil {
		t.AppendRow(table.Row{"Created", pl.CreateTime.String()})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
	return nil
}

func (p *tablePrinter) PrintJobs(jobs []model.Job) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"ID", "NAME", "STATE", "ARCHIVED"})
	for _, j := range jobs {
		t.AppendRow(table.Row{j.ID, j.Name, ColorizeStatus(j.State, p.noColor), j.Archived})
	}
	t.Render()
	return nil
}

func (p *tablePrinter) PrintJob(j model.Job) error {
	t := p.newTable()
	t.AppendRow(table.Row{"ID", j.ID})
	t.AppendRow(table.Row{"Pipeline ID", j.PipelineID})
	t.AppendRow(table.Row{"Name", j.Name})
	t.AppendRow(table.Row{"State", ColorizeStatus(j.State, p.noColor)})
	t.AppendRow(table.Row{"Archived", j.Archived})
	if j.CreateTime != nil {
		t.AppendRow(table.Row{"Created", j.CreateTime.String()})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
	return nil
}

func (p *tablePrinter) PrintBuilds(builds []model.Build) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"ID", "JOB ID", "STATUS", "SHA", "NUMBER"})
	for _, b := range builds {
		sha := b.SHA
		if len(sha) > 8 {
			sha = sha[:8]
		}
		t.AppendRow(table.Row{b.ID, b.JobID, ColorizeStatus(b.Status, p.noColor), sha, fmt.Sprintf("%.0f", b.Number)})
	}
	t.Render()
	return nil
}

func (p *tablePrinter) PrintBuild(b model.Build) error {
	t := p.newTable()
	t.AppendRow(table.Row{"ID", b.ID})
	t.AppendRow(table.Row{"Job ID", b.JobID})
	t.AppendRow(table.Row{"Event ID", b.EventID})
	t.AppendRow(table.Row{"Status", ColorizeStatus(b.Status, p.noColor)})
	t.AppendRow(table.Row{"SHA", b.SHA})
	t.AppendRow(table.Row{"Number", fmt.Sprintf("%.0f", b.Number)})
	if b.CreateTime != nil {
		t.AppendRow(table.Row{"Created", b.CreateTime.String()})
	}
	if b.StartTime != nil {
		t.AppendRow(table.Row{"Started", b.StartTime.String()})
	}
	if b.EndTime != nil {
		t.AppendRow(table.Row{"Ended", b.EndTime.String()})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
	return nil
}

func (p *tablePrinter) PrintSteps(steps []model.Step) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"NAME", "CODE", "START", "END"})
	for _, s := range steps {
		code := ""
		if s.Code != nil {
			code = fmt.Sprintf("%d", *s.Code)
		}
		start := ""
		if s.StartTime != nil {
			start = s.StartTime.Format("2006-01-02 15:04:05")
		}
		end := ""
		if s.EndTime != nil {
			end = s.EndTime.Format("2006-01-02 15:04:05")
		}
		t.AppendRow(table.Row{s.Name, code, start, end})
	}
	t.Render()
	return nil
}

func (p *tablePrinter) PrintEvents(events []model.Event) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"ID", "PIPELINE ID", "STATUS", "SHA", "TYPE", "CREATOR"})
	for _, e := range events {
		sha := e.SHA
		if len(sha) > 8 {
			sha = sha[:8]
		}
		t.AppendRow(table.Row{e.ID, e.PipelineID, ColorizeStatus(e.Status, p.noColor), sha, e.Type, e.Creator.Username})
	}
	t.Render()
	return nil
}

func (p *tablePrinter) PrintEvent(e model.Event) error {
	t := p.newTable()
	t.AppendRow(table.Row{"ID", e.ID})
	t.AppendRow(table.Row{"Pipeline ID", e.PipelineID})
	t.AppendRow(table.Row{"Status", ColorizeStatus(e.Status, p.noColor)})
	t.AppendRow(table.Row{"SHA", e.SHA})
	t.AppendRow(table.Row{"Type", e.Type})
	t.AppendRow(table.Row{"Creator", e.Creator.Username})
	if e.CreateTime != nil {
		t.AppendRow(table.Row{"Created", e.CreateTime.String()})
	}
	if e.ParentEventID != nil {
		t.AppendRow(table.Row{"Parent Event ID", *e.ParentEventID})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
	return nil
}

func (p *tablePrinter) PrintSecrets(secrets []model.Secret) error {
	t := p.newTable()
	t.AppendHeader(table.Row{"ID", "NAME", "PIPELINE ID", "ALLOW IN PR"})
	for _, s := range secrets {
		t.AppendRow(table.Row{s.ID, s.Name, s.PipelineID, s.AllowInPR})
	}
	t.Render()
	return nil
}

var _ = fmtTime // suppress unused warning
