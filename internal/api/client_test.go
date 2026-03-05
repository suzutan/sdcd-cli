package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/config"
)

func TestAuthenticate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v4/auth/token" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"token":"my-jwt"}`)) //nolint:errcheck
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	ctx := &config.Context{APIURL: srv.URL, Token: "raw-token"}
	c := NewClient(ctx)

	if err := c.authenticate(); err != nil {
		t.Fatalf("authenticate: %v", err)
	}
	if c.jwt != "my-jwt" {
		t.Errorf("expected jwt=my-jwt, got %q", c.jwt)
	}

	// second call should be no-op (cached)
	if err := c.authenticate(); err != nil {
		t.Fatalf("second authenticate: %v", err)
	}
}

func TestDo_Get(t *testing.T) {
	type response struct {
		ID int `json:"id"`
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/pipelines/1": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			if r.Header.Get("Authorization") == "" {
				http.Error(w, "missing auth", http.StatusUnauthorized)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{ID: 1}) //nolint:errcheck
		}),
	})
	var result response
	if err := c.do(http.MethodGet, "/v4/pipelines/1", nil, &result); err != nil {
		t.Fatalf("do: %v", err)
	}
	if result.ID != 1 {
		t.Errorf("expected ID=1, got %d", result.ID)
	}
}

func TestDo_ErrorStatus(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/notfound": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "not found", http.StatusNotFound)
		}),
	})
	if err := c.do(http.MethodGet, "/v4/notfound", nil, nil); err == nil {
		t.Error("expected error for 404")
	}
}
