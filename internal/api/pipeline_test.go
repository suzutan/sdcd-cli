package api

import (
	"encoding/json"
	"net/http"
	"net/url"
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
	if len(result) != 1 || result[0].Name != "my-pipeline" {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestListPipelines_WithSearch(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("search") != "foo" {
				http.Error(w, "missing search param", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]model.Pipeline{{ID: 2, Name: "foo-pipeline"}}) //nolint:errcheck
		}),
	})
	result, err := c.ListPipelines(PipelineListParams{Search: "foo", Page: 1, Count: 10})
	if err != nil {
		t.Fatalf("ListPipelines with search: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
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

func TestCreatePipeline(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body CreatePipelineParams
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			resp := model.Pipeline{ID: 99, Name: "new-pl"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.CreatePipeline(CreatePipelineParams{CheckoutURL: "https://github.com/foo/bar"})
	if err != nil {
		t.Fatalf("CreatePipeline: %v", err)
	}
	if result.ID != 99 {
		t.Errorf("expected ID=99, got %d", result.ID)
	}
}

func TestDeletePipeline(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/5": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}),
	})
	if err := c.DeletePipeline(5); err != nil {
		t.Fatalf("DeletePipeline: %v", err)
	}
}

func TestSyncPipeline(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/3/sync": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
		}),
	})
	if err := c.SyncPipeline(3); err != nil {
		t.Fatalf("SyncPipeline: %v", err)
	}
}

func TestGetPipelineJobs(t *testing.T) {
	jobs := []model.Job{{ID: 1, Name: "main"}, {ID: 2, Name: "publish"}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/10/jobs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(jobs) //nolint:errcheck
		}),
	})
	result, err := c.GetPipelineJobs(10, 0, 0)
	if err != nil {
		t.Fatalf("GetPipelineJobs: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 jobs, got %d", len(result))
	}
}

func TestGetPipelineEvents(t *testing.T) {
	events := []model.Event{{ID: 100, PipelineID: 10}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/10/events": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(events) //nolint:errcheck
		}),
	})
	result, err := c.GetPipelineEvents(10, 0, 0)
	if err != nil {
		t.Fatalf("GetPipelineEvents: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 event, got %d", len(result))
	}
}

func TestGetPipelineBuilds(t *testing.T) {
	builds := []model.Build{{ID: 200, Status: "SUCCESS"}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/10/builds": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds) //nolint:errcheck
		}),
	})
	result, err := c.GetPipelineBuilds(10, 0, 0)
	if err != nil {
		t.Fatalf("GetPipelineBuilds: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 build, got %d", len(result))
	}
}

func TestStartPipeline(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body StartPipelineParams
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck

			q, _ := url.ParseQuery(r.URL.RawQuery)
			_ = q

			resp := model.Event{ID: 500, PipelineID: body.PipelineID, Status: "RUNNING"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.StartPipeline(StartPipelineParams{PipelineID: 10, StartFrom: "main"})
	if err != nil {
		t.Fatalf("StartPipeline: %v", err)
	}
	if result.Status != "RUNNING" {
		t.Errorf("expected RUNNING, got %q", result.Status)
	}
}
