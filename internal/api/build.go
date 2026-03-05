package api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func (c *Client) GetBuild(id int) (*model.Build, error) {
	var result model.Build
	return &result, c.get(fmt.Sprintf("/v4/builds/%d", id), &result)
}

func (c *Client) StopBuild(id int) (*model.Build, error) {
	var result model.Build
	return &result, c.put(fmt.Sprintf("/v4/builds/%d", id), map[string]string{"status": "ABORTED"}, &result)
}

func (c *Client) GetBuildSteps(id int) ([]model.Step, error) {
	var result []model.Step
	return result, c.get(fmt.Sprintf("/v4/builds/%d/steps", id), &result)
}

func (c *Client) GetBuildLogs(buildID int, stepName string, from int) (*model.LogPage, error) {
	q := url.Values{}
	q.Set("from", strconv.Itoa(from))
	path := fmt.Sprintf("/v4/builds/%d/steps/%s/logs?%s", buildID, url.PathEscape(stepName), q.Encode())

	var lines []model.LogLine
	headers, err := c.doWithHeaders("GET", path, nil, &lines)
	if err != nil {
		return nil, err
	}

	lp := &model.LogPage{Lines: lines}
	if headers.Get("X-More-Data") == "true" {
		if pageStr := headers.Get("X-Next-Page"); pageStr != "" {
			lp.NextPage, _ = strconv.Atoi(pageStr)
		} else if len(lines) > 0 {
			lp.NextPage = lines[len(lines)-1].N + 1
		}
	}
	return lp, nil
}

// GetAllBuildLogs fetches all log lines for a build step across pages.
func (c *Client) GetAllBuildLogs(buildID int, stepName string) ([]model.LogLine, error) {
	var all []model.LogLine
	from := 0
	for {
		lp, err := c.GetBuildLogs(buildID, stepName, from)
		if err != nil {
			return nil, err
		}
		all = append(all, lp.Lines...)
		if lp.NextPage == 0 {
			break
		}
		from = lp.NextPage
	}
	return all, nil
}

func (c *Client) GetBuildArtifacts(id int) ([]string, error) {
	var result []string
	return result, c.get(fmt.Sprintf("/v4/builds/%d/artifacts", id), &result)
}
