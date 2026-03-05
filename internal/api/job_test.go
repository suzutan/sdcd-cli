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

func TestDisableJob(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/jobs/10": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			state := body["state"]
			resp := model.Job{ID: 10, Name: "main", State: state}
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
