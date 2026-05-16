package backstage

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gobackstage "github.com/datolabs-io/go-backstage/v3"
)

func TestGetClientRequiresHostAndToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		host  string
		token string
	}{
		{
			name:  "missing host",
			host:  "",
			token: "dummy",
		},
		{
			name:  "missing token",
			host:  "http://localhost:7007",
			token: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client, err := getClient(tt.host, tt.token)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if client != nil {
				t.Fatalf("expected nil client on error")
			}
		})
	}
}

func TestGetClientSendsBearerToken(t *testing.T) {
	t.Parallel()

	const expectedToken = "my-secret-token"
	var receivedAuth string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode([]gobackstage.Entity{}); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client, err := getClient(server.URL, expectedToken)
	if err != nil {
		t.Fatalf("getClient returned error: %v", err)
	}

	_, _, err = client.Catalog.Entities.List(context.Background(), &gobackstage.ListEntityOptions{})
	if err != nil {
		t.Fatalf("list entities returned error: %v", err)
	}

	expected := "Bearer " + expectedToken
	if receivedAuth != expected {
		t.Fatalf("expected Authorization header %q, got %q", expected, receivedAuth)
	}
}

func TestGetClientWithMockedCatalog(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/catalog/entities" {
			http.NotFound(w, r)
			return
		}

		entities := []gobackstage.Entity{
			{
				ApiVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Metadata: gobackstage.EntityMeta{
					Name:      "demo-service",
					Namespace: "default",
				},
				Spec: map[string]interface{}{
					"type":  "service",
					"owner": "team-a",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(entities); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	client, err := getClient(server.URL, "dummy-token")
	if err != nil {
		t.Fatalf("getClient returned error: %v", err)
	}

	entities, _, err := client.Catalog.Entities.List(context.Background(), &gobackstage.ListEntityOptions{})
	if err != nil {
		t.Fatalf("list entities returned error: %v", err)
	}
	if len(entities) != 1 {
		t.Fatalf("expected 1 entity, got %d", len(entities))
	}
	if entities[0].Metadata.Name != "demo-service" {
		t.Fatalf("unexpected entity name: %s", entities[0].Metadata.Name)
	}
}
