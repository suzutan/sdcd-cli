package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestGetEvent(t *testing.T) {
	e := model.Event{ID: 200, PipelineID: 1, Status: "SUCCESS"}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events/200": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(e) //nolint:errcheck
		}),
	})
	result, err := c.GetEvent(200)
	if err != nil {
		t.Fatalf("GetEvent: %v", err)
	}
	if result.PipelineID != 1 {
		t.Errorf("expected PipelineID=1, got %d", result.PipelineID)
	}
}

func TestGetEventBuilds(t *testing.T) {
	builds := []model.Build{
		{ID: 10, Status: "SUCCESS"},
		{ID: 11, Status: "FAILURE"},
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events/5/builds": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(builds) //nolint:errcheck
		}),
	})
	result, err := c.GetEventBuilds(5)
	if err != nil {
		t.Fatalf("GetEventBuilds: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 builds, got %d", len(result))
	}
}

func TestStopEvent(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events/300/stop": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
		}),
	})
	if err := c.StopEvent(300); err != nil {
		t.Fatalf("StopEvent: %v", err)
	}
}

func TestRerunEvent(t *testing.T) {
	orig := model.Event{ID: 50, PipelineID: 7, Status: "SUCCESS"}
	newEvent := model.Event{ID: 51, PipelineID: 7, Status: "RUNNING"}

	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events/50": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(orig) //nolint:errcheck
		}),
		"/v4/events": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body RerunEventParams
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			if body.ParentEventID != 50 {
				http.Error(w, "missing parentEventId", http.StatusBadRequest)
				return
			}
			if body.PipelineID != 7 {
				http.Error(w, "wrong pipelineId", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newEvent) //nolint:errcheck
		}),
	})

	result, err := c.RerunEvent(50, "")
	if err != nil {
		t.Fatalf("RerunEvent: %v", err)
	}
	if result.ID != 51 {
		t.Errorf("expected new event ID=51, got %d", result.ID)
	}
}

func TestRerunEvent_WithJob(t *testing.T) {
	orig := model.Event{ID: 60, PipelineID: 3}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/events/60": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(orig) //nolint:errcheck
		}),
		"/v4/events": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body RerunEventParams
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			if body.StartFrom != "deploy" {
				http.Error(w, "wrong startFrom", http.StatusBadRequest)
				return
			}
			resp := model.Event{ID: 61, PipelineID: 3}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})

	result, err := c.RerunEvent(60, "deploy")
	if err != nil {
		t.Fatalf("RerunEvent with job: %v", err)
	}
	if result.ID != 61 {
		t.Errorf("expected ID=61, got %d", result.ID)
	}
}
