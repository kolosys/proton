# Proton 📚

[![Go Reference](https://pkg.go.dev/badge/github.com/kolosys/proton.svg)](https://pkg.go.dev/github.com/kolosys/proton)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolosys/proton)](https://goreportcard.com/report/github.com/kolosys/proton)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Proton** is an opinionated, open-source documentation generator specifically designed for Go libraries. It automatically creates comprehensive, GitBook-compatible documentation from your Go source code, comments, and configurable templates.

## 🚀 Features

- **📦 Multi-Package Support** - Handle single and multi-package Go libraries effortlessly
- **🤖 Auto-Discovery** - Automatically discover and document all packages in your project
- **📖 GitBook Integration** - Generate GitBook-compatible documentation with `.gitbook.yml`
- **🎨 Customizable Templates** - Use built-in templates or create your own
- **⚡ GitHub Actions** - Automated documentation deployment with GitHub Actions
- **🔧 Configurable Output** - Flexible output directory and structure configuration
- **📝 API Documentation** - Generate detailed API reference from Go comments
- **💡 Examples & Guides** - Auto-extract examples and generate comprehensive guides
- **🔍 Smart Parsing** - Parse Go AST to extract documentation, types, and examples
- **✅ Configuration Validation** - Validate your configuration before generation
- **🧹 Selective Cleaning** - Preserve custom content while regenerating auto-generated docs

## 📥 Installation

### Install as CLI Tool

```bash
go install github.com/kolosys/proton/cmd/proton@latest
```

### Use as GitHub Action

Add to your `.github/workflows/docs.yml`:

```yaml
name: Generate Documentation

on:
  push:
    branches: [main]

jobs:
  docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: kolosys/proton@v1
        with:
          deploy-to-pages: true
```

## 🏃‍♂️ Quick Start

### 1. Initialize Configuration

```bash
# Initialize in current directory
proton init

# Initialize in specific project
proton init ./my-go-project
```

This creates a `.proton/config.yml` file with sensible defaults.

### 2. Validate Configuration

```bash
# Validate current directory
proton validate

# Validate specific project
proton validate ./my-go-project

# Validate with custom config
proton validate --config custom-config.yml
```

### 3. Generate Documentation

```bash
# Generate with default settings
proton generate

# Generate with custom output directory
proton generate --output my-docs

# Generate with custom configuration
proton generate --config custom-config.yml
```

## ⚙️ Configuration

Proton uses a YAML configuration file (`.proton/config.yml`) to customize documentation generation:

```yaml
repository:
  name: my-library
  owner: myuser
  description: "A fantastic Go library"
  import_path: github.com/myuser/my-library

output:
  directory: docs
  clean: true
  gitbook_config: true

discovery:
  packages:
    auto_discover: true
    exclude_patterns:
      - "./vendor/..."
      - "./test/..."
      - "./.git/..."
      - "**/*_test.go"

  api_generation:
    enabled: true
    include_unexported: false
    include_examples: true

  examples:
    enabled: true
    auto_discover: true

  guides:
    enabled: true
    include_contributing: true
    include_faq: true

gitbook:
  title: My Library Documentation
  description: Complete documentation for My Library
```

### Configuration Schema

See [Configuration Schema](schema/config.yml) for complete documentation of all available options.

## 📚 Generated Documentation Structure

Proton generates a well-organized documentation structure:

```
docs/
├── README.md                    # Main documentation index
├── .gitbook.yml                 # GitBook configuration
├── SUMMARY.md                   # GitBook navigation
├── getting-started/
│   ├── README.md                # Getting started overview
│   └── [package-name].md        # Package-specific getting started guides
├── api-reference/
│   ├── README.md                # API reference index
│   └── [package-name].md        # Package-specific API documentation
├── examples/
│   ├── README.md                # Examples overview
│   └── [example-category]/      # Example categories
│       ├── README.md            # Category overview
│       └── [example-name].md    # Individual examples
└── guides/
    ├── README.md                # Guides overview
    ├── contributing.md          # Contributing guidelines
    ├── faq.md                   # Frequently asked questions
    └── [package-name]/          # Package-specific guides
        └── best-practices.md    # Package best practices
```

## 🎨 Templates

Proton comes with built-in templates that work great out of the box, but you can customize them:

### Built-in Templates

- `index.md` - Main documentation page
- `packages-index.md` - Package overview
- `package.md` - Individual package documentation
- `api-reference.md` - API reference documentation
- `examples-index.md` - Examples overview
- `guides-index.md` - Guides overview
- `contributing.md` - Contributing guidelines
- `faq.md` - FAQ page
- `gitbook-config.yml` - GitBook configuration

### Custom Templates

1. Create a custom templates directory inside `.proton`:

   ```
   .proton/
   ├── config.yml
   └── templates/
       ├── custom-index.md
       └── custom-api.md
   ```

2. Configure in your `.proton/config.yml`:
   ```yaml
   templates:
     directory: ./.proton/templates
     custom_templates:
       - name: index
         file: ./.proton/templates/custom-index.md
   ```

## 🤖 GitHub Action Usage

### Basic Usage

```yaml
- name: Generate Documentation
  uses: kolosys/proton@v1
```

### Advanced Usage

```yaml
- name: Generate Documentation
  uses: kolosys/proton@v1
  with:
    config-file: ".proton/config.yml"
    output-directory: "documentation"
    clean-output: "false"
    deploy-to-pages: true
    deploy-to-branch: "gh-pages"
    auto-commit: true
    commit-message: "📚 Update documentation"
    go-version: "1.24"
    proton-version: "latest"
```

### Action Inputs

| Input              | Description                | Default                                    |
| ------------------ | -------------------------- | ------------------------------------------ |
| `config-file`      | Path to configuration file | `.proton/config.yml`                       |
| `output-directory` | Output directory           | `docs`                                     |
| `clean-output`     | Clean output directory     | `true`                                     |
| `deploy-to-pages`  | Deploy to GitHub Pages     | `false`                                    |
| `deploy-to-branch` | Deploy to specific branch  | ``                                         |
| `auto-commit`      | Auto-commit changes        | `false`                                    |
| `commit-message`   | Commit message             | `📚 Update documentation (auto-generated)` |
| `go-version`       | Go version to use          | `1.24`                                     |
| `proton-version`   | Proton version to install  | `latest`                                   |
| `token`            | GitHub token for auth      | `${{ github.token }}`                      |

### Selective Cleaning

When `clean-output` is set to `false`, Proton preserves custom content in guides and other directories while only regenerating auto-generated content (api-reference, examples, packages, README.md, SUMMARY.md).

## 📖 Examples

### Document a Single Package

```yaml
# .proton/config.yml
repository:
  name: my-package
  import_path: github.com/user/my-package

discovery:
  packages:
    auto_discover: true
    include_patterns: ["."]
```

### Document Multiple Packages

```yaml
# .proton/config.yml
discovery:
  packages:
    auto_discover: true
    include_patterns: ["./..."]
    exclude_patterns:
      - "./internal/..."
      - "./test/..."
      - "./vendor/..."
```

### Custom Package Documentation

```yaml
# .proton/config.yml
discovery:
  packages:
    manual_packages:
      - name: core
        path: ./pkg/core
        description: Core functionality
      - name: utils
        path: ./pkg/utils
        description: Utility functions
```

## 🛠️ Development

### Prerequisites

- Go 1.21 or later
- Git

### Building from Source

```bash
git clone https://github.com/kolosys/proton.git
cd proton
go build -o proton ./cmd/proton
```

### Running Tests

```bash
go test ./...
```

### Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by the need for better Go library documentation
- Built with love for the Go community
- Thanks to all contributors and users

## 📞 Support

- 📄 [Documentation](docs/)
- 🐛 [Issue Tracker](https://github.com/kolosys/proton/issues)
- 💬 [Discussions](https://github.com/kolosys/proton/discussions)
- 📧 [Contact](mailto:support@kolosys.com)

---

**Made with ❤️ by [Kolosys](https://github.com/kolosys)**
