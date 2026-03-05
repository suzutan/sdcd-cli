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

	// second call must be no-op (cached)
	if err := c.authenticate(); err != nil {
		t.Fatalf("second authenticate: %v", err)
	}
}

func TestAuthenticate_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer srv.Close()

	ctx := &config.Context{APIURL: srv.URL, Token: "bad-token"}
	c := NewClient(ctx)

	if err := c.authenticate(); err == nil {
		t.Error("expected error for 401")
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

func TestDo_Post(t *testing.T) {
	type body struct{ Name string }
	type response struct{ ID int }
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/things": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			if ct := r.Header.Get("Content-Type"); ct != "application/json" {
				http.Error(w, "wrong content-type: "+ct, http.StatusBadRequest)
				return
			}
			var b body
			json.NewDecoder(r.Body).Decode(&b) //nolint:errcheck
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{ID: 42}) //nolint:errcheck
		}),
	})
	var result response
	if err := c.post("/v4/things", body{Name: "test"}, &result); err != nil {
		t.Fatalf("post: %v", err)
	}
	if result.ID != 42 {
		t.Errorf("expected ID=42, got %d", result.ID)
	}
}

func TestDo_Delete(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/things/1": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}),
	})
	if err := c.delete("/v4/things/1"); err != nil {
		t.Fatalf("delete: %v", err)
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

func TestDoWithHeaders_ReturnsHeaders(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/logs": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-More-Data", "true")
			w.Header().Set("X-Next-Page", "3")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[]`)) //nolint:errcheck
		}),
	})
	var result []struct{}
	headers, err := c.doWithHeaders(http.MethodGet, "/v4/logs", nil, &result)
	if err != nil {
		t.Fatalf("doWithHeaders: %v", err)
	}
	if headers.Get("X-More-Data") != "true" {
		t.Errorf("expected X-More-Data=true, got %q", headers.Get("X-More-Data"))
	}
	if headers.Get("X-Next-Page") != "3" {
		t.Errorf("expected X-Next-Page=3, got %q", headers.Get("X-Next-Page"))
	}
}
