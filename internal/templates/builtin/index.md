# {{.Repository.Name}} Documentation

{{.Repository.Description}}

## Quick Navigation

### 🚀 [Getting Started](getting-started/README.md)

Everything you need to get up and running with {{.Repository.Name}}.

### 📚 [API Reference](api-reference/README.md)

Complete API documentation for all packages.

### 📖 [Examples](examples/README.md)

Working examples and tutorials.

### 📘 [Guides](guides/README.md)

In-depth guides and best practices.

## Package Overview

{{- range .Packages}}

### {{.Name}}

{{.Description}}

- [Getting Started](getting-started/{{.Name}}.md)
- [API Reference](api-reference/{{.Name}}.md)
- [Examples](examples/README.md)
- [Best Practices](guides/{{.Name}}/best-practices.md)
  {{- end}}

## External Resources

- [GitHub Repository]({{.Repository.URL}})
- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Repository.ImportPath}})
- [Issues & Support]({{.Repository.URL}}/issues)

## Contributing

See our [Contributing Guide](guides/contributing.md) to get started.
