package api

import (
	"archive/zip"
	"bytes"
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

func TestStopBuild(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/100": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			if body["status"] != "ABORTED" {
				http.Error(w, "wrong status", http.StatusBadRequest)
				return
			}
			resp := model.Build{ID: 100, Status: "ABORTED"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.StopBuild(100)
	if err != nil {
		t.Fatalf("StopBuild: %v", err)
	}
	if result.Status != "ABORTED" {
		t.Errorf("expected ABORTED, got %q", result.Status)
	}
}

func TestGetBuildSteps(t *testing.T) {
	code := 0
	steps := []model.Step{{Name: "install", Code: &code}, {Name: "test"}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/50/steps": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(steps) //nolint:errcheck
		}),
	})
	result, err := c.GetBuildSteps(50)
	if err != nil {
		t.Fatalf("GetBuildSteps: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 steps, got %d", len(result))
	}
	if result[0].Name != "install" {
		t.Errorf("expected install, got %q", result[0].Name)
	}
}

func TestGetBuildLogs(t *testing.T) {
	logs := []model.LogLine{
		{T: 1000, M: "hello", N: 0},
		{T: 2000, M: "world", N: 1},
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/5/steps/install/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if got := r.URL.Query().Get("from"); got != "0" {
				http.Error(w, "expected from=0, got "+got, http.StatusBadRequest)
				return
			}
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

func TestGetBuildLogs_WithNextPageHeader(t *testing.T) {
	logs := []model.LogLine{{T: 1000, M: "line0", N: 0}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/5/steps/install/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if got := r.URL.Query().Get("from"); got != "0" {
				http.Error(w, "expected from=0, got "+got, http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-More-Data", "true")
			w.Header().Set("X-Next-Page", "2")
			json.NewEncoder(w).Encode(logs) //nolint:errcheck
		}),
	})
	lp, err := c.GetBuildLogs(5, "install", 0)
	if err != nil {
		t.Fatalf("GetBuildLogs: %v", err)
	}
	if lp.NextPage != 2 {
		t.Errorf("expected NextPage=2, got %d", lp.NextPage)
	}
}

func TestGetBuildLogs_DeriveNextFromLastLine(t *testing.T) {
	// X-More-Data: true without X-Next-Page => NextPage derived from last line N+1
	logs := []model.LogLine{{T: 1000, M: "line0", N: 4}}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/5/steps/install/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-More-Data", "true")
			json.NewEncoder(w).Encode(logs) //nolint:errcheck
		}),
	})
	lp, err := c.GetBuildLogs(5, "install", 0)
	if err != nil {
		t.Fatalf("GetBuildLogs: %v", err)
	}
	if lp.NextPage != 5 {
		t.Errorf("expected NextPage=5 (last N=4, +1), got %d", lp.NextPage)
	}
}

func TestGetAllBuildLogs_Pagination(t *testing.T) {
	// first page: lines N=0, second page: lines N=1; no X-Next-Page so from is derived
	page0 := []model.LogLine{{T: 1000, M: "line0", N: 0}}
	page1 := []model.LogLine{{T: 2000, M: "line1", N: 1}}

	callCount := 0
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/7/steps/test/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			from := r.URL.Query().Get("from")
			if callCount == 0 {
				if from != "0" {
					http.Error(w, "first call: expected from=0, got "+from, http.StatusBadRequest)
					return
				}
				callCount++
				w.Header().Set("X-More-Data", "true")
				json.NewEncoder(w).Encode(page0) //nolint:errcheck
				return
			}
			if from != "1" {
				http.Error(w, "second call: expected from=1, got "+from, http.StatusBadRequest)
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

func makeZip(files map[string]string) []byte {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for name, content := range files {
		w, _ := zw.Create(name)
		w.Write([]byte(content)) //nolint:errcheck
	}
	zw.Close()
	return buf.Bytes()
}

func TestGetBuildArtifact(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/20/artifacts/manifest.txt": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("hello artifact")) //nolint:errcheck
		}),
	})
	data, err := c.GetBuildArtifact(20, "./manifest.txt")
	if err != nil {
		t.Fatalf("GetBuildArtifact: %v", err)
	}
	if string(data) != "hello artifact" {
		t.Errorf("unexpected content: %q", string(data))
	}
}

func TestGetBuildArtifacts(t *testing.T) {
	zipData := makeZip(map[string]string{
		"./manifest.txt":    "hello",
		"./environment.json": "{}",
	})
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/builds/20/artifacts": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/zip")
			w.Write(zipData) //nolint:errcheck
		}),
	})
	result, err := c.GetBuildArtifacts(20)
	if err != nil {
		t.Fatalf("GetBuildArtifacts: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 artifacts, got %d", len(result))
	}
}
