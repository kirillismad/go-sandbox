# Go Sandbox

A comprehensive Go development sandbox with examples, algorithms, and various integrations.

## 📋 Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Branch Protection](#branch-protection)
- [Contributing](#contributing)
- [Documentation](#documentation)

## 🚀 Overview

This repository serves as a sandbox for Go development, containing:

- **Algorithms & Data Structures**: Queue, Stack, Sorting, Search implementations
- **HTTP APIs**: RESTful services with Swagger documentation
- **gRPC Services**: Protocol buffer definitions and implementations
- **Database Integration**: SQL operations with various databases
- **Message Queues**: Kafka and NATS integrations
- **Caching**: Redis and in-memory cache implementations
- **Concurrency**: Go concurrency patterns and examples
- **Rate Limiting**: Various rate limiting strategies

## 📁 Project Structure

```
├── algos/           # Algorithms and data structures
├── cache/           # Caching implementations
├── concurrency/     # Concurrency patterns
├── docs/            # Project documentation
├── grpc/            # gRPC services and definitions
├── http/            # HTTP API implementations
├── kafka/           # Kafka integration examples
├── nats/            # NATS integration examples
├── rate_limiters/   # Rate limiting implementations
├── sql/             # Database operations and examples
├── tasks/           # Coding challenges and solutions
├── utils/           # Utility functions
└── .github/         # GitHub configuration and workflows
```

## 🛠 Getting Started

### Prerequisites

- Go 1.21 or later
- Task (optional, for running predefined tasks)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kirillismad/go-sandbox.git
   cd go-sandbox
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests to verify installation:
   ```bash
   go test ./...
   ```

### Using Task Runner

This project includes a `Taskfile.yml` for common operations:

```bash
# Install required tools and dependencies
task init

# Generate Swagger documentation
task swag

# Build the project
task build
```

## 🔄 Development Workflow

### Standard Workflow

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** and ensure tests pass:
   ```bash
   go test ./...
   ```

3. **Commit your changes**:
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

4. **Push to your branch**:
   ```bash
   git push origin feature/your-feature-name
   ```

5. **Create a Pull Request** on GitHub

### Code Quality

- All code must pass existing tests
- New features should include appropriate tests
- Follow Go best practices and conventions
- Use meaningful commit messages

## 🛡️ Branch Protection

This repository implements branch protection rules to ensure code quality and security. **Direct pushes to the master branch are not allowed**.

### Key Protection Rules

- ✅ **Pull requests required** for all changes to master
- ✅ **Code review required** before merging
- ✅ **Status checks must pass** (CI/CD pipeline)
- ✅ **Conversation resolution required**
- ✅ **Branch must be up to date** before merging

### For Detailed Information

See our comprehensive [Branch Protection Guide](docs/BRANCH_PROTECTION.md) for:
- Setting up branch protection rules
- Managing repository access
- Code review workflows
- Troubleshooting common issues
- Best practices for contributors and maintainers

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

1. **Read the [Branch Protection Guide](docs/BRANCH_PROTECTION.md)** to understand our workflow
2. **Check existing issues** or create a new one to discuss your changes
3. **Fork the repository** and create a feature branch
4. **Follow our coding standards** and include tests
5. **Submit a pull request** with a clear description

### Code Owners

This repository uses GitHub's CODEOWNERS feature for automatic review assignment. See [`.github/CODEOWNERS`](.github/CODEOWNERS) for current assignments.

### Issue Templates

When reporting bugs or requesting features, please use our issue templates:
- [Bug Report](.github/ISSUE_TEMPLATE/bug_report.md)
- [Feature Request](.github/ISSUE_TEMPLATE/feature_request.md)

## 📚 Documentation

- [Branch Protection Guide](docs/BRANCH_PROTECTION.md) - Complete guide to repository access management
- [API Documentation](http/echo_example/docs/) - Swagger documentation for HTTP APIs
- [Shortcuts Reference](shortcuts.md) - Development environment shortcuts

## 🔧 Available Commands

### HTTP Server Examples

```bash
# Run Echo HTTP server with Swagger UI
go run main.go petproj

# Access Swagger UI at http://localhost:8080/docs
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test package
go test ./algos/queue
```

### Building

```bash
# Build all packages
go build ./...

# Build specific binary
go build -o bin/server ./main.go
```

## 🚦 CI/CD Pipeline

This project includes a comprehensive CI/CD pipeline that runs on every pull request:

- **Testing**: Automated test execution with coverage reporting
- **Linting**: Code quality checks with golangci-lint
- **Security**: Vulnerability scanning and security analysis
- **Building**: Compilation verification across multiple targets
- **Dependency Validation**: Go module verification and vulnerability checks

The pipeline must pass before any PR can be merged to master.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🔗 Related Projects

- [Go Documentation](https://golang.org/doc/)
- [GitHub Branch Protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/about-protected-branches)

---

For questions or support, please create an issue using our [issue templates](.github/ISSUE_TEMPLATE/).