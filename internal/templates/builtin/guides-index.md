# Guides

In-depth guides and best practices for {{.Repository.Name}}.

## Getting Started

- [Installation & Setup](../getting-started.md)
- [Quick Start Guide](quick-start.md)
- [Basic Concepts](concepts.md)

## Best Practices

- [Performance Optimization](performance.md)
- [Error Handling](error-handling.md)
- [Testing Strategies](testing.md)
- [Production Deployment](deployment.md)

## Advanced Topics

- [Architecture Overview](architecture.md)
- [Extending {{.Repository.Name}}](extending.md)
- [Integration Patterns](integration.md)
- [Troubleshooting](troubleshooting.md)

## Package-Specific Guides

{{- range .Packages}}

### {{.Name}}

{{.Description}}

- [{{.Name}} Best Practices]({{.Name}}/best-practices.md) - Recommended patterns and usage

{{- end}}

## Community Resources

- [Contributing Guide](contributing.md)
- [Code of Conduct](code-of-conduct.md)
- [Security Policy](security.md)
- [FAQ](faq.md)

## External Resources

- [GitHub Repository]({{.Repository.URL}})
- [Discussions]({{.Repository.URL}}/discussions)
- [Issues]({{.Repository.URL}}/issues)
