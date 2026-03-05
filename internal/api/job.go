package api

import (
	"fmt"
	"net/url"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func (c *Client) GetJob(id int) (*model.Job, error) {
	var result model.Job
	return &result, c.get(fmt.Sprintf("/v4/jobs/%d", id), &result)
}

type UpdateJobParams struct {
	State string `json:"state"`
}

func (c *Client) UpdateJob(id int, p UpdateJobParams) (*model.Job, error) {
	var result model.Job
	return &result, c.put(fmt.Sprintf("/v4/jobs/%d", id), p, &result)
}

func (c *Client) EnableJob(id int) (*model.Job, error) {
	return c.UpdateJob(id, UpdateJobParams{State: "ENABLED"})
}

func (c *Client) DisableJob(id int) (*model.Job, error) {
	return c.UpdateJob(id, UpdateJobParams{State: "DISABLED"})
}

func (c *Client) GetJobBuilds(id int, page, count int) ([]model.Build, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if count > 0 {
		q.Set("count", fmt.Sprintf("%d", count))
	}
	path := fmt.Sprintf("/v4/jobs/%d/builds", id)
	if len(q) > 0 {
		path += "?" + q.Encode()
	}
	var result []model.Build
	return result, c.get(path, &result)
}

func (c *Client) GetLatestBuild(jobID int) (*model.Build, error) {
	builds, err := c.GetJobBuilds(jobID, 1, 1)
	if err != nil {
		return nil, err
	}
	if len(builds) == 0 {
		return nil, fmt.Errorf("no builds found for job %d", jobID)
	}
	return &builds[0], nil
}
