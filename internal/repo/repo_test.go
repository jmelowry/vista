package repo

import (
	"testing"
)

func TestGetAllRepositories(t *testing.T) {
	repos := GetAllRepositories()
	if len(repos) == 0 {
		t.Error("Expected non-empty repository list")
	}

	foundECR := false
	foundDockerHub := false

	for _, repo := range repos {
		switch repo.ID {
		case "ecr-main":
			foundECR = true
			if repo.Type != "ecr" {
				t.Errorf("Expected ecr-main repository type to be 'ecr', got '%s'", repo.Type)
			}
		case "dockerhub":
			foundDockerHub = true
			if repo.Type != "dockerhub" {
				t.Errorf("Expected dockerhub repository type to be 'dockerhub', got '%s'", repo.Type)
			}
		}
	}

	if !foundECR {
		t.Error("Expected to find 'ecr-main' repository")
	}
	if !foundDockerHub {
		t.Error("Expected to find 'dockerhub' repository")
	}
}

func TestGetRepository(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		wantRepo bool
	}{
		{
			name:     "get existing repository",
			id:       "ecr-main",
			wantRepo: true,
		},
		{
			name:     "get non-existent repository",
			id:       "nonexistent",
			wantRepo: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := GetRepository(tt.id)

			if tt.wantRepo && repo == nil {
				t.Errorf("Expected to find repository with ID '%s', got nil", tt.id)
			} else if !tt.wantRepo && repo != nil {
				t.Errorf("Expected nil for non-existent repository ID '%s', got %+v", tt.id, repo)
			}

			if repo != nil && repo.ID != tt.id {
				t.Errorf("Expected repository ID to be '%s', got '%s'", tt.id, repo.ID)
			}
		})
	}
}
