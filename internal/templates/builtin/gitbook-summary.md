# Summary

- [Introduction](README.md)

## Getting Started

- [Packages](packages/README.md)
  {{- range .Packages}}
  {{- if not (isMainPackage .)}}
  - [{{.Name}}](packages/{{.Name}}.md)
    {{- end}}
    {{- end}}

## API Reference

- [API Overview](api-reference/README.md)
  {{- range .Packages}}
  {{- if not (isMainPackage .)}}
  - [{{.Name}} API](api-reference/{{.Name}}.md)
    {{- end}}
    {{- end}}

{{- if .Config.Discovery.Examples.Enabled}}

## Examples

- [Examples Overview](examples/README.md)
  {{- range .Packages}}
  {{- if hasExamples .}}
  - [{{.Name}} Examples](examples/{{.Name}}/README.md)
    {{- end}}
    {{- end}}
    {{- end}}

{{- if .Config.Discovery.Guides.Enabled}}

## Guides

- [Guides Overview](guides/README.md)
  {{- if .Config.Discovery.Guides.IncludeContributing}}
- [Contributing](guides/contributing.md)
  {{- end}}
  {{- if .Config.Discovery.Guides.IncludeFAQ}}
- [FAQ](guides/faq.md)
  {{- end}}
  {{- range .Config.Discovery.Guides.CustomGuides}}
- [{{.Title}}](guides/{{.Name}}.md)
  {{- end}}
  {{- end}}
