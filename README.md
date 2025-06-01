# Vista

[![Go Version](https://img.shields.io/github/go-mod/go-version/jamie/vista)](https://github.com/jamie/vista)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![CI Status](https://img.shields.io/github/actions/workflow/status/jamie/vista/ci.yml?branch=main&label=CI)](https:/## Contributing

Vista is in early development. If you're interested in using or contributing, please open an issue or start a discussion.

Please see the following documents for more information:
- [CONTRIBUTING.md](CONTRIBUTING.md): Guidelines for contributing to Vista
- [IMPLEMENTATION.md](IMPLEMENTATION.md): Technical implementation details and future plans

---

## Development Approach

Vista is designed so that the API and CLI can be developed in tandem. This allows for rapid iteration, easier testing, and ensures both interfaces remain consistent as features are added.om/jamie/vista/actions)

**Vista** is a clean, unified UI layer for browsing and managing artifact repositories—wherever they live. Vista is not a storage layer; it's a frontend designed to help developers quickly find, inspect, and interact with internal artifacts like container images and build outputs.ista

**Vista** is a clean, unified UI layer for browsing and managing artifact repositories—wherever they live. Vista is not a storage layer; it’s a frontend designed to help developers quickly find, inspect, and interact with internal artifacts like container images and build outputs.

---

## What it does

Vista provides a simple way to:
- Browse and search artifacts across multiple repositories and backends.
- View metadata, tags, and history for each artifact.
- Integrate with popular container registries (ECR, Docker Hub, etc.).
- Offer a consistent, user-friendly interface for artifact discovery.

---

## Artifact Metadata

Artifacts (sometimes called "resources"—see note below) are displayed with rich metadata, such as:

```json
{
  "name": "my-app",
  "type": "container-image",
  "repository": "123456789012.dkr.ecr.us-west-2.amazonaws.com/my-app",
  "tags": ["latest", "v1.2.3"],
  "created": "2025-05-30T12:34:56Z",
  "size": "128MB",
  "digest": "sha256:abc123...",
  "owner": "my-team@example.com"
}
```

> **Note:** The term "artifact" is used here as a generic label for items managed by Vista (e.g., containers, packages, Helm charts). You may also see the term "resource" used to refer to these items in the API and UI.

---

## API (MVP)

Vista API currently implements these endpoints:

- `GET /repos` – List all connected repositories
- `GET /repo/<repoid>` – Get details for a specific repository
- `GET /repo/<repoid>/resources` - List all resources in a repository
- `GET /repo/<repoid>/resource/<resourceid>` - Get details for a specific resource

These endpoints serve as the foundation for both the CLI and UI interfaces.

---

## Current Design

### Architecture

Vista follows a clean, layered architecture designed for modularity and extensibility:

1. **API Layer** (`api/server.go`): HTTP handlers that process requests and responses
2. **Domain Models** (`internal/*/models.go`): Core data structures representing the business domain
3. **Mock Data Layer** (`internal/*/mock.go`): Temporary mock implementations for development

The architecture is deliberately structured to allow for easy replacement of the mock layer with real repository connectors in the future:

```
vista/
├── cmd/                # CLI entrypoints
│   └── vista/
│       └── main.go     # Main application entry point
├── api/                # API server implementation 
│   └── server.go       # HTTP handlers and routing logic
├── internal/           # Internal packages
│   ├── repo/           # Repository-related code
│   │   ├── models.go   # Repository data models
│   │   └── mock.go     # Mock repository data
│   └── resource/       # Resource-related code
│       ├── models.go   # Resource data models
│       └── mock.go     # Mock resource data
├── scripts/            # Helper scripts
│   └── test-api.sh     # API testing script
```

### RESTful API Design

The API follows RESTful design principles with a hierarchical resource structure:

- Resources are scoped to repositories (e.g., `/repo/<repoid>/resource/<resourceid>`)
- Standard HTTP methods are used (currently GET only, with POST/PUT/DELETE planned)
- JSON is used for all request and response bodies
- Proper error handling with appropriate HTTP status codes
- Comprehensive logging for debugging and monitoring

### Data Model

Vista's data model consists of two primary entities:

- **Repository**: A source of artifacts (e.g., ECR, Docker Hub)
  ```go
  type Repository struct {
      ID          string `json:"id"`
      Name        string `json:"name"`
      Type        string `json:"type"`
      URL         string `json:"url"`
      Description string `json:"description,omitempty"`
  }
  ```

- **Resource**: An artifact stored in a repository (e.g., container image)
  ```go
  type Resource struct {
      ID         string   `json:"id"`
      Name       string   `json:"name"`
      Type       string   `json:"type"`
      Repository string   `json:"repository"`
      Tags       []string `json:"tags,omitempty"`
      Created    string   `json:"created,omitempty"`
      Size       string   `json:"size,omitempty"`
      Digest     string   `json:"digest,omitempty"`
      Owner      string   `json:"owner,omitempty"`
  }
  ```

Resources are always scoped to a repository, represented in the API path structure.

---

## Development and Testing

### Using the Makefile

Vista includes a Makefile to simplify common development tasks:

```bash
# Build the binary
make build

# Run the server (default port: 8080)
make run

# Run on a custom port
make run PORT=9000

# Run all tests
make test

# Format code
make fmt

# Show all available commands
make help
```

### Running the Server Manually

To start the API server without using the Makefile:

```bash
go run ./cmd/vista/main.go
```

This will start the server on port 8080 by default. You can specify a different port with the `-port` flag:

```bash
go run ./cmd/vista/main.go -port 9000
```

### Building the Project

```bash
# Using the Makefile (recommended)
make build

# Or manually with go build
go build -o vista ./cmd/vista
```

The compiled binary can then be run directly:

```bash
./vista
```

### Testing the API

The enhanced test script provides comprehensive testing of all API endpoints with proper error checking:

```bash
# Using the Makefile
make test-api

# Or run the script directly
./scripts/test-api.sh
```

The test script checks all endpoints and provides a clear summary of results:

```
===== Vista API Test Suite =====
Testing against: http://localhost:8080

Testing: List all repositories
GET /repos
✓ Status: 200 (expected: 200)
Response:
[
  {
    "id": "ecr-main",
    "name": "ECR Main",
    "type": "ecr",
    "url": "123456789012.dkr.ecr.us-west-2.amazonaws.com",
    "description": "Main ECR repository"
  },
  ...
]

...additional test results...

===== Test Summary =====
Tests passed: 8
Tests failed: 0

All tests passed!
```

### Logging and Debugging

The API server includes comprehensive logging to help with debugging:

```
VISTA API: 2023/10/15 12:34:56 server.go:28: Starting Vista API server on :8080
VISTA API: 2023/10/15 12:34:59 server.go:23: Received request for /repos
VISTA API: 2023/10/15 12:34:59 server.go:44: Handling request for all repositories
VISTA API: 2023/10/15 12:34:59 server.go:48: Returning 2 repositories
```

Logs include:
- Request path information
- Routing decisions
- Repository and resource IDs being accessed
- Error conditions and not found scenarios
- Response details

### Mock Data System

The current implementation uses an in-memory mock data system for development purposes. This approach allows rapid development of the API layer without requiring real repository connections:

#### Mock Data Architecture

- **Repositories**: Defined in `internal/repo/mock.go` with functions to access them
- **Resources**: Defined in `internal/resource/mock.go` with repository-scoped organization

The mock system is organized hierarchically:
1. Repositories are stored in a map by ID
2. Resources are stored in a nested map (repository ID → resource ID → resource)

```go
// Repository mock data structure
func MockRepositories() map[string]Repository {
    return map[string]Repository{
        "ecr-main": {
            ID:          "ecr-main",
            Name:        "ECR Main",
            Type:        "ecr",
            URL:         "123456789012.dkr.ecr.us-west-2.amazonaws.com",
            Description: "Main ECR repository",
        },
        // More repositories...
    }
}

// Resource mock data structure
func MockResources() map[string]map[string]Resource {
    return map[string]map[string]Resource{
        "ecr-main": {
            "my-app": {
                ID:         "my-app",
                Name:       "my-app",
                Type:       "container-image",
                Repository: "123456789012.dkr.ecr.us-west-2.amazonaws.com/my-app",
                Tags:       []string{"latest", "v1.2.3"},
                // More fields...
            },
            // More resources...
        },
        // More repositories...
    }
}
```

#### Accessing Mock Data

The packages provide simple access functions that abstract away the implementation details:

```go
// Repository functions
repo.GetAllRepositories()        // Get all repositories
repo.GetRepository(id)           // Get a specific repository

// Resource functions
resource.GetResourcesForRepo(id) // Get all resources for a repository
resource.GetResource(repoID, id) // Get a specific resource in a repository
```

#### Extending the Mock Data

To add or modify mock data:

1. Edit the appropriate mock.go file
2. Add new entries to the return value of `MockRepositories()` or `MockResources()`
3. Ensure resources are assigned to the correct repository ID

This design makes it easy to test different scenarios and edge cases during development.

---

## Roadmap

- [ ] Basic UI for browsing artifacts
- [ ] Flexible support for loading json metadata from various sources
- [ ] A sync/ cache mechanism for metadata
- [ ] ECR integration
- [ ] Support for other container registries
- [ ] Search and filtering UI
- [ ] Artifact metadata caching
- [ ] CLI for artifact queries
- [ ] S3/local metadata source support

See [IMPLEMENTATION.md](IMPLEMENTATION.md) for detailed technical considerations regarding the implementation of these features.

---

## License

MIT

---

## Contributing

Vista is in early development. If you’re interested in using or contributing, please open an issue or start a discussion.

---

## Development Approach

Vista is designed so that the API and CLI can be developed in tandem. This allows for rapid iteration, easier testing, and ensures both interfaces remain consistent as features are added.

---

## Repository Structure

```
vista/
├── cmd/                # CLI entrypoints (e.g., main.go for CLI)
│   └── vista/
│       └── main.go
├── api/                # API server implementation (handlers, routes)
│   └── server.go
├── internal/           # Internal packages (business logic, models, etc.)
│   ├── repo/
│   ├── resource/
│   └── ... 
├── pkg/                # Exported Go packages (if any, for reuse)
├── web/                # (Optional) UI frontend code
├── scripts/            # Helper scripts (build, dev, etc.)
├── go.mod
├── go.sum
└── README.md
```

- Place CLI code in `cmd/vista/`.
- Place API server code in `api/`.
- Place core logic and types in `internal/`.