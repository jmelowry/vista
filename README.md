# Vista

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

- `GET /repos` – List all connected repositories
- `GET /repo/<repoid>` – Get details for a specific repository
- `GET /resource/<resourceid>` – Get details for a specific resource (artifact, package, container, etc.)

---

## CLI (Planned)

```bash
vista info my-app                   # Show metadata for an artifact
vista list                         # List all artifacts
vista tags my-app                  # List tags for an artifact
```

---

## Requirements

- Go 1.21+
- Access to supported artifact repositories (ECR, Docker Hub, etc.)
- (Optional) S3 or local filesystem for metadata caching

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