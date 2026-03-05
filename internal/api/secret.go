package api

import (
	"fmt"
	"net/url"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func (c *Client) ListSecrets(pipelineID int) ([]model.Secret, error) {
	q := url.Values{}
	q.Set("pipelineId", fmt.Sprintf("%d", pipelineID))
	var result []model.Secret
	return result, c.get("/v4/secrets?"+q.Encode(), &result)
}

type CreateSecretParams struct {
	PipelineID int    `json:"pipelineId"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	AllowInPR  bool   `json:"allowInPR"`
}

func (c *Client) CreateSecret(p CreateSecretParams) (*model.Secret, error) {
	var result model.Secret
	return &result, c.post("/v4/secrets", p, &result)
}

type UpdateSecretParams struct {
	Value     *string `json:"value,omitempty"`
	AllowInPR *bool   `json:"allowInPR,omitempty"`
}

func (c *Client) UpdateSecret(id int, p UpdateSecretParams) (*model.Secret, error) {
	var result model.Secret
	return &result, c.put(fmt.Sprintf("/v4/secrets/%d", id), p, &result)
}

func (c *Client) DeleteSecret(id int) error {
	return c.delete(fmt.Sprintf("/v4/secrets/%d", id))
}
