package api

import (
	"fmt"
	"net/url"

	"github.com/suzutan/sdcd-cli/internal/model"
)

type PipelineListParams struct {
	Search string
	Page   int
	Count  int
}

func (c *Client) ListPipelines(p PipelineListParams) ([]model.Pipeline, error) {
	q := url.Values{}
	if p.Search != "" {
		q.Set("search", p.Search)
	}
	if p.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if p.Count > 0 {
		q.Set("count", fmt.Sprintf("%d", p.Count))
	}
	path := "/v4/pipelines"
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result []model.Pipeline
	return result, c.get(path, &result)
}

func (c *Client) GetPipeline(id int) (*model.Pipeline, error) {
	var result model.Pipeline
	return &result, c.get(fmt.Sprintf("/v4/pipelines/%d", id), &result)
}

type CreatePipelineParams struct {
	CheckoutURL string `json:"checkoutUrl"`
	RootDir     string `json:"rootDir,omitempty"`
}

func (c *Client) CreatePipeline(p CreatePipelineParams) (*model.Pipeline, error) {
	var result model.Pipeline
	return &result, c.post("/v4/pipelines", p, &result)
}

func (c *Client) DeletePipeline(id int) error {
	return c.delete(fmt.Sprintf("/v4/pipelines/%d", id))
}

func (c *Client) SyncPipeline(id int) error {
	return c.post(fmt.Sprintf("/v4/pipelines/%d/sync", id), nil, nil)
}

func (c *Client) GetPipelineJobs(id int, page, count int) ([]model.Job, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if count > 0 {
		q.Set("count", fmt.Sprintf("%d", count))
	}
	path := fmt.Sprintf("/v4/pipelines/%d/jobs", id)
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result []model.Job
	return result, c.get(path, &result)
}

func (c *Client) GetPipelineEvents(id int, page, count int) ([]model.Event, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if count > 0 {
		q.Set("count", fmt.Sprintf("%d", count))
	}
	path := fmt.Sprintf("/v4/pipelines/%d/events", id)
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result []model.Event
	return result, c.get(path, &result)
}

func (c *Client) GetPipelineBuilds(id int, page, count int) ([]model.Build, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if count > 0 {
		q.Set("count", fmt.Sprintf("%d", count))
	}
	path := fmt.Sprintf("/v4/pipelines/%d/builds", id)
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result []model.Build
	return result, c.get(path, &result)
}

type StartPipelineParams struct {
	PipelineID    int    `json:"pipelineId"`
	StartFrom     string `json:"startFrom,omitempty"`
	SHA           string `json:"sha,omitempty"`
}

func (c *Client) StartPipeline(p StartPipelineParams) (*model.Event, error) {
	var result model.Event
	return &result, c.post("/v4/events", p, &result)
}
