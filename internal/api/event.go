package api

import (
	"fmt"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func (c *Client) GetEvent(id int) (*model.Event, error) {
	var result model.Event
	return &result, c.get(fmt.Sprintf("/v4/events/%d", id), &result)
}

func (c *Client) GetEventBuilds(id int) ([]model.Build, error) {
	var result []model.Build
	return result, c.get(fmt.Sprintf("/v4/events/%d/builds", id), &result)
}

func (c *Client) StopEvent(id int) error {
	return c.put(fmt.Sprintf("/v4/events/%d/stop", id), nil, nil)
}

type RerunEventParams struct {
	PipelineID    int    `json:"pipelineId"`
	StartFrom     string `json:"startFrom,omitempty"`
	ParentEventID int    `json:"parentEventId"`
}

func (c *Client) RerunEvent(id int, jobName string) (*model.Event, error) {
	// Get the original event to find pipelineId
	orig, err := c.GetEvent(id)
	if err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}
	p := RerunEventParams{
		PipelineID:    orig.PipelineID,
		ParentEventID: id,
	}
	if jobName != "" {
		p.StartFrom = jobName
	}
	var result model.Event
	return &result, c.post("/v4/events", p, &result)
}
