# Summary

- [Introduction](README.md)

## Getting Started

- [Overview](getting-started/overview.md)
- [Installation](getting-started/installation.md)
- [Quick Start](getting-started/quick-start.md)

## Core Concepts

{{- range .Packages}}
{{- if not (isMainPackage .)}}
- [{{.Name}}](core-concepts/{{.Name}}.md)
{{- end}}
{{- end}}

## Advanced

- [Performance Tuning](advanced/performance-tuning.md)
- [Best Practices](advanced/best-practices.md)

## Reference

### API Reference

- [API Overview](reference/api-reference/README.md)
  {{- range .Packages}}
  {{- if not (isMainPackage .)}}
  - [{{.Name}} API](reference/api-reference/{{.Name}}.md)
    {{- end}}
    {{- end}}

{{- if .Config.Discovery.Examples.Enabled}}

### Examples

- [Examples Overview](reference/examples/README.md)
{{- end}}
