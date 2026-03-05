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
