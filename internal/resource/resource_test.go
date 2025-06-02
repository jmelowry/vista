package resource

import (
	"testing"
)

func TestGetResourcesForRepo(t *testing.T) {
	tests := []struct {
		name       string
		repoID     string
		wantEmpty  bool
		wantCounts map[string]int
	}{
		{
			name:      "ecr-main resources",
			repoID:    "ecr-main",
			wantEmpty: false,
			wantCounts: map[string]int{
				"my-app":      1,
				"api-service": 1,
			},
		},
		{
			name:      "dockerhub resources",
			repoID:    "dockerhub",
			wantEmpty: false,
			wantCounts: map[string]int{
				"nginx": 1,
			},
		},
		{
			name:       "non-existent repo",
			repoID:     "nonexistent",
			wantEmpty:  true,
			wantCounts: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resources := GetResourcesForRepo(tt.repoID)

			if tt.wantEmpty && len(resources) > 0 {
				t.Errorf("Expected empty resources list for repo '%s', got %d resources", tt.repoID, len(resources))
			} else if !tt.wantEmpty && len(resources) == 0 {
				t.Errorf("Expected non-empty resources list for repo '%s', got empty list", tt.repoID)
			}

			resourceCounts := make(map[string]int)
			for _, res := range resources {
				resourceCounts[res.ID]++
			}

			for resID, expectedCount := range tt.wantCounts {
				if count := resourceCounts[resID]; count != expectedCount {
					t.Errorf("Expected %d resources with ID '%s', got %d", expectedCount, resID, count)
				}
			}
		})
	}
}

func TestGetResource(t *testing.T) {
	tests := []struct {
		name       string
		repoID     string
		resourceID string
		wantNil    bool
	}{
		{
			name:       "existing resource",
			repoID:     "ecr-main",
			resourceID: "my-app",
			wantNil:    false,
		},
		{
			name:       "non-existent resource",
			repoID:     "ecr-main",
			resourceID: "nonexistent",
			wantNil:    true,
		},
		{
			name:       "non-existent repo",
			repoID:     "nonexistent",
			resourceID: "my-app",
			wantNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := GetResource(tt.repoID, tt.resourceID)

			if tt.wantNil && resource != nil {
				t.Errorf("Expected nil for resource '%s' in repo '%s', got %+v", tt.resourceID, tt.repoID, resource)
			} else if !tt.wantNil && resource == nil {
				t.Errorf("Expected non-nil for resource '%s' in repo '%s', got nil", tt.resourceID, tt.repoID)
			}

			if resource != nil {
				if resource.ID != tt.resourceID {
					t.Errorf("Expected resource ID to be '%s', got '%s'", tt.resourceID, resource.ID)
				}
			}
		})
	}
}
