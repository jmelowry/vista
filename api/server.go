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
		s.handleRepos(w, r)
	})

	http.HandleFunc("/repo/", func(w http.ResponseWriter, r *http.Request) {
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

	repos := repo.GetAllRepositories()

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

	if len(pathParts) == 0 || pathParts[0] == "" {
		http.Error(w, "Repository ID is required", http.StatusBadRequest)
		return
	}

	repoID := pathParts[0]

	switch {
	case len(pathParts) == 1:
		s.handleRepo(w, r, repoID)

	case len(pathParts) == 2 && pathParts[1] == "resources":
		s.handleRepoResources(w, r, repoID)

	case len(pathParts) == 3 && pathParts[1] == "resource":
		s.handleRepoResource(w, r, repoID, pathParts[2])

	default:
		http.Error(w, "Invalid path", http.StatusNotFound)
	}
}

func (s *Server) handleRepo(w http.ResponseWriter, r *http.Request, repoID string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	repository := repo.GetRepository(repoID)
	if repository == nil {
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

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

	repository := repo.GetRepository(repoID)
	if repository == nil {
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

	resources := resource.GetResourcesForRepo(repoID)

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

	repository := repo.GetRepository(repoID)
	if repository == nil {
		http.Error(w, fmt.Sprintf("Repository '%s' not found", repoID), http.StatusNotFound)
		return
	}

	res := resource.GetResource(repoID, resourceID)
	if res == nil {
		http.Error(w, fmt.Sprintf("Resource '%s' not found in repository '%s'", resourceID, repoID), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
