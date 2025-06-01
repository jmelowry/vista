package repo

// MockRepositories returns a map of mock repositories for testing and development
func MockRepositories() map[string]Repository {
	return map[string]Repository{
		"ecr-main": {
			ID:          "ecr-main",
			Name:        "ECR Main",
			Type:        "ecr",
			URL:         "123456789012.dkr.ecr.us-west-2.amazonaws.com",
			Description: "Main ECR repository",
		},
		"dockerhub": {
			ID:          "dockerhub",
			Name:        "Docker Hub",
			Type:        "dockerhub",
			URL:         "https://hub.docker.com",
			Description: "Docker Hub registry",
		},
	}
}

// GetRepository returns a repository by ID or nil if not found
func GetRepository(id string) *Repository {
	repos := MockRepositories()
	if repo, exists := repos[id]; exists {
		return &repo
	}
	return nil
}

// GetAllRepositories returns a slice of all repositories
func GetAllRepositories() []Repository {
	repos := MockRepositories()
	result := make([]Repository, 0, len(repos))
	for _, repo := range repos {
		result = append(result, repo)
	}
	return result
}
