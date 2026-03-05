package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/suzutan/sdcd-cli/internal/model"
)

func TestListSecrets(t *testing.T) {
	secrets := []model.Secret{
		{ID: 1, Name: "MY_SECRET", PipelineID: 10},
	}
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/secrets": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(secrets) //nolint:errcheck
		}),
	})
	result, err := c.ListSecrets(10)
	if err != nil {
		t.Fatalf("ListSecrets: %v", err)
	}
	if len(result) != 1 || result[0].Name != "MY_SECRET" {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestCreateSecret(t *testing.T) {
	c := NewMockServer(t, map[string]http.Handler{
		"/v4/secrets": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "wrong method", http.StatusMethodNotAllowed)
				return
			}
			var body CreateSecretParams
			json.NewDecoder(r.Body).Decode(&body) //nolint:errcheck
			resp := model.Secret{ID: 99, Name: body.Name, PipelineID: body.PipelineID}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(resp) //nolint:errcheck
		}),
	})
	result, err := c.CreateSecret(CreateSecretParams{
		PipelineID: 10,
		Name:       "NEW_SECRET",
		Value:      "secret-val",
	})
	if err != nil {
		t.Fatalf("CreateSecret: %v", err)
	}
	if result.Name != "NEW_SECRET" {
		t.Errorf("expected NEW_SECRET, got %q", result.Name)
	}
}
