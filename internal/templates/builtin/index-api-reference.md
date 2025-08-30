# API Reference

Complete API documentation for {{.Repository.Name}}.

## Overview

This section contains detailed API documentation for all packages. For package overviews and getting started guides, see the [Packages](../packages/README.md) section.

## Package APIs

{{- range .Packages}}

### [{{.Name}}]({{.Name}}.md)

{{.Description}}

**[→ Full API Documentation]({{.Name}}.md)**

Key APIs:

- Types and interfaces
- Functions and methods
- Constants and variables
- Detailed usage examples

{{- end}}

## Navigation

- **[Packages](../packages/README.md)** - Package overviews and installation
- **[Examples](../examples/README.md)** - Working code examples
- **[Guides](../guides/README.md)** - Best practices and patterns

## External References

- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Repository.ImportPath}}) - Go module documentation
- [GitHub Repository]({{.Repository.URL}}) - Source code and issues
