package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestListPipelines(t *testing.T) {
	pipelines := []model.Pipeline{
		{ID: 1, Name: "my-pipeline", State: "ACTIVE"},
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pipelines) //nolint:errcheck
		}),
	})
	result, err := c.ListPipelines(PipelineListParams{})
	if err != nil {
		t.Fatalf("ListPipelines: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
	if result[0].Name != "my-pipeline" {
		t.Errorf("expected name=my-pipeline, got %q", result[0].Name)
	}
}

func TestGetPipeline(t *testing.T) {
	pl := model.Pipeline{ID: 42, Name: "test-pl"}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/42": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pl) //nolint:errcheck
		}),
	})
	result, err := c.GetPipeline(42)
	if err != nil {
		t.Fatalf("GetPipeline: %v", err)
	}
	if result.ID != 42 {
		t.Errorf("expected ID=42, got %d", result.ID)
	}
}
