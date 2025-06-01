package resource

// MockResources returns a map of repository IDs to a map of resource IDs to resources
func MockResources() map[string]map[string]Resource {
	return map[string]map[string]Resource{
		"ecr-main": {
			"my-app": {
				ID:         "my-app",
				Name:       "my-app",
				Type:       "container-image",
				Repository: "123456789012.dkr.ecr.us-west-2.amazonaws.com/my-app",
				Tags:       []string{"latest", "v1.2.3"},
				Created:    "2025-05-30T12:34:56Z",
				Size:       "128MB",
				Digest:     "sha256:abc123...",
				Owner:      "my-team@example.com",
			},
			"api-service": {
				ID:         "api-service",
				Name:       "api-service",
				Type:       "container-image",
				Repository: "123456789012.dkr.ecr.us-west-2.amazonaws.com/api-service",
				Tags:       []string{"latest", "v2.0.1"},
				Created:    "2025-05-29T10:12:34Z",
				Size:       "95MB",
				Digest:     "sha256:def456...",
				Owner:      "api-team@example.com",
			},
		},
		"dockerhub": {
			"nginx": {
				ID:         "nginx",
				Name:       "nginx",
				Type:       "container-image",
				Repository: "docker.io/library/nginx",
				Tags:       []string{"latest", "1.21.6"},
				Created:    "2025-04-15T08:30:00Z",
				Size:       "142MB",
				Digest:     "sha256:ghi789...",
				Owner:      "nginx-maintainers",
			},
		},
	}
}

// GetResourcesForRepo returns all resources for a repository
func GetResourcesForRepo(repoID string) []Resource {
	resources := MockResources()
	if repoResources, exists := resources[repoID]; exists {
		result := make([]Resource, 0, len(repoResources))
		for _, res := range repoResources {
			result = append(result, res)
		}
		return result
	}
	return []Resource{}
}

// GetResource returns a specific resource from a repository
func GetResource(repoID, resourceID string) *Resource {
	resources := MockResources()
	if repoResources, exists := resources[repoID]; exists {
		if resource, exists := repoResources[resourceID]; exists {
			return &resource
		}
	}
	return nil
}
