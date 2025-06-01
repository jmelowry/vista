package repo

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

func GetRepository(id string) *Repository {
	repos := MockRepositories()
	if repo, exists := repos[id]; exists {
		return &repo
	}
	return nil
}

func GetAllRepositories() []Repository {
	repos := MockRepositories()
	result := make([]Repository, 0, len(repos))
	for _, repo := range repos {
		result = append(result, repo)
	}
	return result
}
