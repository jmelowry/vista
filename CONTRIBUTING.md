# Contributing to Vista

Thank you for your interest in contributing to Vista! This document provides guidelines and instructions for contributing to the project.

## Getting Started

1. **Fork the repository**: Create your own fork of Vista to work on.
2. **Clone your fork**: `git clone https://github.com/your-username/vista.git`
3. **Set up upstream**: `git remote add upstream https://github.com/jmelowry/vista.git`
4. **Create a branch**: `git checkout -b feature/your-feature-name`

## Development Workflow

### Setting Up the Development Environment

1. Make sure you have Go 1.24+ installed
2. Run `go mod tidy` to install dependencies
3. Use the provided Makefile commands for common tasks

### Making Changes

1. Make your changes in your feature branch
2. Write or update tests as necessary
3. Ensure all tests pass with `make test`
4. Format your code with `make fmt`
5. Lint your code with `make lint`

### Commit Guidelines

- Use clear, concise commit messages
- Reference issue numbers in commit messages when applicable
- Keep commits focused on a single logical change
- Rebase your branch before submitting a PR

Example commit message:
```
Add ECR repository connector

- Implement basic ECR connector interface
- Add authentication for AWS credentials
- Add unit tests for connector methods

Fixes #42
```

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update the IMPLEMENTATION.md if you're adding new features or changing architecture
3. Ensure your PR passes all CI checks
4. Request a review from a maintainer
5. Address any feedback from reviewers

## Code Style

Vista follows standard Go coding conventions:
- Use `gofmt` to format your code (or `make fmt`)
- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Write clear, concise comments and documentation
- Keep functions small and focused on a single responsibility

## Testing

- Write unit tests for all new functions and methods
- Ensure existing tests pass with your changes
- Use table-driven tests where appropriate
- Mock external dependencies for unit tests

## License

By contributing to Vista, you agree that your contributions will be licensed under the project's MIT license.
