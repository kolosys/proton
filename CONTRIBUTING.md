# Contributing to Proton

We love your input! We want to make contributing to Proton as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## Development Process

We use GitHub to sync code to and from our public repository. We'll use GitHub to track issues and feature requests, as well as accept pull requests.

## Pull Requests

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for convenience commands)

### Setting Up Development Environment

1. **Fork and clone the repository:**

   ```bash
   git clone https://github.com/your-username/proton.git
   cd proton
   ```

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Build the project:**

   ```bash
   go build -o proton ./cmd/proton
   ```

4. **Run tests:**

   ```bash
   go test ./...
   ```

5. **Test the CLI:**
   ```bash
   ./proton --help
   ```

### Project Structure

```
proton/
├── cmd/proton/              # CLI entry point
├── internal/
│   ├── cli/                 # CLI commands
│   ├── config/              # Configuration handling
│   ├── discovery/           # Package discovery logic
│   ├── generator/           # Documentation generation
│   └── templates/           # Template engine
│       └── builtin/         # Built-in templates
├── examples/                # Example configurations
├── schema/                  # Configuration schema
├── templates/               # Template files for docs
├── .github/workflows/       # GitHub Actions
└── tests/                   # Integration tests
```

## Types of Contributions

### 🐛 Bug Reports

Great Bug Reports tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

### 💡 Feature Requests

We track feature requests as GitHub issues. When creating a feature request, please include:

- Clear description of the problem the feature would solve
- Detailed description of the desired behavior
- Example use cases
- Any implementation ideas (optional)

### 🔧 Code Contributions

#### Areas Where We'd Love Help

1. **Template Improvements**

   - New built-in templates
   - Template function enhancements
   - Better GitBook integration

2. **Package Discovery**

   - Better Go module parsing
   - Support for more project structures
   - Performance improvements

3. **Documentation Generation**

   - Enhanced API documentation
   - Better example extraction
   - Support for more documentation formats

4. **CLI Experience**

   - New commands and flags
   - Better error messages
   - Interactive configuration

5. **GitHub Action**
   - More deployment options
   - Better caching
   - Additional output formats

#### Development Guidelines

1. **Code Style**

   - Follow Go conventions and style
   - Use `gofmt` to format your code
   - Run `go vet` to check for common errors
   - Write meaningful commit messages

2. **Testing**

   - Write tests for new functionality
   - Ensure existing tests pass
   - Add integration tests for CLI commands
   - Include examples in your tests

3. **Documentation**

   - Update README.md for user-facing changes
   - Add/update function and package documentation
   - Update configuration schema if needed
   - Include examples for new features

4. **Compatibility**
   - Maintain backward compatibility when possible
   - Follow semantic versioning principles
   - Document breaking changes clearly

### 📝 Documentation Contributions

Documentation improvements are always welcome:

- Fix typos or clarify confusing sections
- Add examples and use cases
- Improve API documentation
- Create tutorials and guides
- Translate documentation (future)

## Submitting Changes

### Commit Message Format

We follow conventional commits:

```
type(scope): description

[optional body]

[optional footer]
```

Types:

- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:

```
feat(cli): add validate command for configuration files
fix(discovery): handle edge case in package parsing
docs(readme): update installation instructions
```

### Pull Request Process

1. **Create a branch** from `main` with a descriptive name:

   ```bash
   git checkout -b feat/add-validation-command
   ```

2. **Make your changes** following the development guidelines

3. **Test thoroughly**:

   ```bash
   go test ./...
   go build ./cmd/proton
   ./proton generate --help
   ```

4. **Commit your changes** with conventional commit messages

5. **Push to your fork** and create a pull request

6. **Fill out the PR template** with:
   - Clear description of changes
   - Motivation and context
   - Screenshots (if applicable)
   - Testing performed
   - Checklist completion

### Review Process

1. Maintainers will review your PR
2. Feedback will be provided within a few days
3. Address any requested changes
4. Once approved, maintainers will merge the PR

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/config

# Run integration tests
go test ./tests/...
```

### Writing Tests

- Place unit tests in the same package as the code
- Use descriptive test names: `TestConfigLoad_WithValidFile`
- Include table-driven tests for multiple scenarios
- Test error conditions and edge cases
- Mock external dependencies when needed

## Release Process

Releases are handled by maintainers:

1. Version tags follow semantic versioning (v1.2.3)
2. Release notes are generated from commit messages
3. GitHub Actions automatically builds and publishes releases
4. Docker images and binaries are published automatically

## Community

- GitHub Discussions for questions and ideas
- GitHub Issues for bugs and feature requests
- Follow [@kolosys](https://github.com/kolosys) for updates

## License

By contributing, you agree that your contributions will be licensed under the same MIT License that covers the project. Feel free to contact the maintainers if that's a concern.

## Questions?

Don't hesitate to reach out! You can:

- Open an issue for questions
- Start a discussion for ideas
- Contact maintainers directly

Thank you for contributing to Proton! 🚀
