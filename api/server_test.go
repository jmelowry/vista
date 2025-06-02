package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamie/vista/internal/repo"
	"github.com/jamie/vista/internal/resource"
)

func TestHandleRepos(t *testing.T) {
	server := NewServer(8080)

	req, err := http.NewRequest("GET", "/repos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.handleRepos)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var repos []repo.Repository
	err = json.Unmarshal(rr.Body.Bytes(), &repos)
	if err != nil {
		t.Errorf("failed to parse response body: %v", err)
	}

	if len(repos) == 0 {
		t.Error("expected non-empty repositories list")
	}
}

func TestHandleRepo(t *testing.T) {
	server := NewServer(8080)

	tests := []struct {
		name     string
		repoID   string
		wantCode int
	}{
		{
			name:     "valid repo",
			repoID:   "ecr-main",
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid repo",
			repoID:   "nonexistent",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/repo/"+tt.repoID, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				server.handleRepo(w, r, tt.repoID)
			})
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if tt.wantCode == http.StatusOK {
				var repository repo.Repository
				err = json.Unmarshal(rr.Body.Bytes(), &repository)
				if err != nil {
					t.Errorf("failed to parse response body: %v", err)
				}

				if repository.ID != tt.repoID {
					t.Errorf("handler returned wrong repository ID: got %v want %v", repository.ID, tt.repoID)
				}
			}
		})
	}
}

func TestHandleRepoResources(t *testing.T) {
	server := NewServer(8080)

	tests := []struct {
		name     string
		repoID   string
		wantCode int
	}{
		{
			name:     "valid repo",
			repoID:   "ecr-main",
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid repo",
			repoID:   "nonexistent",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/repo/"+tt.repoID+"/resources", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				server.handleRepoResources(w, r, tt.repoID)
			})
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if tt.wantCode == http.StatusOK {
				var resources []resource.Resource
				err = json.Unmarshal(rr.Body.Bytes(), &resources)
				if err != nil {
					t.Errorf("failed to parse response body: %v", err)
				}
			}
		})
	}
}

func TestHandleRepoResource(t *testing.T) {
	server := NewServer(8080)

	tests := []struct {
		name       string
		repoID     string
		resourceID string
		wantCode   int
	}{
		{
			name:       "valid resource",
			repoID:     "ecr-main",
			resourceID: "my-app",
			wantCode:   http.StatusOK,
		},
		{
			name:       "invalid repo",
			repoID:     "nonexistent",
			resourceID: "my-app",
			wantCode:   http.StatusNotFound,
		},
		{
			name:       "invalid resource",
			repoID:     "ecr-main",
			resourceID: "nonexistent",
			wantCode:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/repo/"+tt.repoID+"/resource/"+tt.resourceID, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				server.handleRepoResource(w, r, tt.repoID, tt.resourceID)
			})
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if tt.wantCode == http.StatusOK {
				var resource resource.Resource
				err = json.Unmarshal(rr.Body.Bytes(), &resource)
				if err != nil {
					t.Errorf("failed to parse response body: %v", err)
				}

				if resource.ID != tt.resourceID {
					t.Errorf("handler returned wrong resource ID: got %v want %v", resource.ID, tt.resourceID)
				}
			}
		})
	}
}
