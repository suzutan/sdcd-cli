package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/config"
)

// NewMockServer creates a test HTTP server with the given routes and returns a Client
// configured to use it. The server is closed automatically when the test finishes.
func NewMockServer(t *testing.T, routes map[string]http.Handler) *Client {
	t.Helper()
	mux := http.NewServeMux()

	// Register auth endpoint
	mux.HandleFunc("/v4/auth/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token":"test-jwt-token"}`)) //nolint:errcheck
	})

	for pattern, handler := range routes {
		mux.Handle(pattern, handler)
	}

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	ctx := &config.Context{
		Name:   "test",
		APIURL: srv.URL,
		Token:  "test-token",
	}
	return NewClient(ctx)
}
