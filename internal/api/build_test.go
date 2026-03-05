package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestGetBuild(t *testing.T) {
	b := model.Build{ID: 100, Status: "SUCCESS"}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/100": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(b) //nolint:errcheck
		}),
	})
	result, err := c.GetBuild(100)
	if err != nil {
		t.Fatalf("GetBuild: %v", err)
	}
	if result.Status != "SUCCESS" {
		t.Errorf("expected SUCCESS, got %q", result.Status)
	}
}

func TestGetBuildLogs(t *testing.T) {
	logs := []model.LogLine{
		{T: 1000, M: "hello", N: 0},
		{T: 2000, M: "world", N: 1},
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/5/steps/install/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// no X-More-Data header = no more pages
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(logs) //nolint:errcheck
		}),
	})
	lp, err := c.GetBuildLogs(5, "install", 0)
	if err != nil {
		t.Fatalf("GetBuildLogs: %v", err)
	}
	if len(lp.Lines) != 2 {
		t.Errorf("expected 2 lines, got %d", len(lp.Lines))
	}
	if lp.NextPage != 0 {
		t.Errorf("expected no next page, got %d", lp.NextPage)
	}
}

func TestGetAllBuildLogs_Pagination(t *testing.T) {
	page0 := []model.LogLine{{T: 1000, M: "line0", N: 0}}
	page1 := []model.LogLine{{T: 2000, M: "line1", N: 1}}

	callCount := 0
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/7/steps/test/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if callCount == 0 {
				callCount++
				w.Header().Set("X-More-Data", "true")
				json.NewEncoder(w).Encode(page0) //nolint:errcheck
				return
			}
			json.NewEncoder(w).Encode(page1) //nolint:errcheck
		}),
	})
	lines, err := c.GetAllBuildLogs(7, "test")
	if err != nil {
		t.Fatalf("GetAllBuildLogs: %v", err)
	}
	if len(lines) != 2 {
		t.Errorf("expected 2 total lines, got %d", len(lines))
	}
}
