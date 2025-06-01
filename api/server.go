package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jamie/vista/internal/repo"
	"github.com/jamie/vista/internal/resource"
)

var logger = log.New(os.Stdout, "VISTA API: ", log.LstdFlags|log.Lshortfile)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Received request for /repos")
		s.handleRepos(w, r)
	})

	http.HandleFunc("/repo/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Received request for %s", r.URL.Path)
		s.handleRepoRequests(w, r)
	})

	addr := fmt.Sprintf(":%d", s.port)
	logger.Printf("Starting Vista API server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleRepos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Printf("Handling request for all repositories")

	repos := repo.GetAllRepositories()

	logger.Printf("Returning %d repositories", len(repos))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(repos); err != nil {
		logger.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleRepoRequests(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/repo/")

	pathParts := strings.Split(path, "/")

	logger.Printf("Original URL path: %s", r.URL.Path)
	logger.Printf("After trimming prefix: %s", path)
	logger.Printf("Path parts: %v", pathParts)

	if len(pathParts) == 0 || pathParts[0] == "" {
		http.Error(w, "Repository ID is required", http.StatusBadRequest)
		return
	}

	repoID := pathParts[0]
	logger.Printf("Repository ID: %s", repoID)

	switch {
	case len(pathParts) == 1:
		logger.Printf("Routing to handleRepo for %s", repoID)
		s.handleRepo(w, r, repoID)

	case len(pathParts) == 2 && pathParts[1] == "resources":
		logger.Printf("Routing to handleRepoResources for %s", repoID)
		s.handleRepoResources(w, r, repoID)

	case len(pathParts) == 3 && pathParts[1] == "resource":
		logger.Printf("Routing to handleRepoResource for %s, resource %s", repoID, pathParts[2])
		s.handleRepoResource(w, r, repoID, pathParts[2])

	default:
		logger.Printf("No route match for path: %s with %d parts", path, len(pathParts))
		http.Error(w, "Invalid path", http.StatusNotFound)
	}
}

func (s *Server) handleRepo(w http.ResponseWriter, r *http.Request, repoID string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Printf("Handling repository request for: %s", repoID)

	repository := repo.GetRepository(repoID)
	if repository == nil {
		logger.Printf("Repository '%s' not found", repoID)
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

	logger.Printf("Returning repository data for %s", repoID)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(repository); err != nil {
		logger.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleRepoResources(w http.ResponseWriter, r *http.Request, repoID string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Printf("Handling resources for repo: %s", repoID)

	repository := repo.GetRepository(repoID)
	if repository == nil {
		logger.Printf("Repository '%s' not found", repoID)
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

	resources := resource.GetResourcesForRepo(repoID)

	logger.Printf("Returning %d resources for repo %s", len(resources), repoID)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resources); err != nil {
		logger.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleRepoResource(w http.ResponseWriter, r *http.Request, repoID, resourceID string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Printf("Handling resource request for repo: %s, resource: %s", repoID, resourceID)

	repository := repo.GetRepository(repoID)
	if repository == nil {
		logger.Printf("Repository '%s' not found", repoID)
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

	res := resource.GetResource(repoID, resourceID)
	if res == nil {
		logger.Printf("Resource '%s' not found in repository '%s'", resourceID, repoID)
		http.Error(w, fmt.Sprintf("Resource '%s' not found in repository '%s'", resourceID, repoID), http.StatusNotFound)
		return
	}

	logger.Printf("Returning resource %s for repo %s", resourceID, repoID)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
