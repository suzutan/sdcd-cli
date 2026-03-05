package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestGetJob(t *testing.T) {
	j := model.Job{ID: 10, Name: "main", State: "ENABLED"}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/10": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(j) //nolint:errcheck
		}),
	})
	result, err := c.GetJob(10)
	if err != nil {
		t.Fatalf("GetJob: %v", err)
	}
	if result.State != "ENABLED" {
		t.Errorf("expected ENABLED, got %q", result.State)
	}
}

func TestEnableJob(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/10": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			resp := model.Job{ID: 10, State: body["state"]}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.EnableJob(10)
	if err != nil {
		t.Fatalf("EnableJob: %v", err)
	}
	if result.State != "ENABLED" {
		t.Errorf("expected ENABLED, got %q", result.State)
	}
}

func TestDisableJob(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/10": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			resp := model.Job{ID: 10, Name: "main", State: body["state"]}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.DisableJob(10)
	if err != nil {
		t.Fatalf("DisableJob: %v", err)
	}
	if result.State != "DISABLED" {
		t.Errorf("expected DISABLED, got %q", result.State)
	}
}

func TestGetJobBuilds(t *testing.T) {
	builds := []model.Build{{ID: 1, Status: "SUCCESS"}, {ID: 2, Status: "FAILURE"}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/10/builds": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds) //nolint:errcheck
		}),
	})
	result, err := c.GetJobBuilds(10, 0, 0)
	if err != nil {
		t.Fatalf("GetJobBuilds: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 builds, got %d", len(result))
	}
}

func TestGetLatestBuild(t *testing.T) {
	builds := []model.Build{{ID: 99, Status: "RUNNING"}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/5/builds": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds) //nolint:errcheck
		}),
	})
	result, err := c.GetLatestBuild(5)
	if err != nil {
		t.Fatalf("GetLatestBuild: %v", err)
	}
	if result.ID != 99 {
		t.Errorf("expected ID=99, got %d", result.ID)
	}
}

func TestGetLatestBuild_NoBuilds(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/5/builds": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]model.Build{}) //nolint:errcheck
		}),
	})
	_, err := c.GetLatestBuild(5)
	if err == nil {
		t.Error("expected error for no builds")
	}
}
