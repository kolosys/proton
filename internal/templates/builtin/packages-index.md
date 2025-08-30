# Packages

Package overviews and getting started guides for {{.Repository.Name}}.

## Overview

{{.Repository.Description}}

## Available Packages

{{- range .Packages}}
{{- if not (isMainPackage .)}}

### [{{.Name}}]({{.Name}}.md)

{{.Description}}

**Import Path:** `{{.ImportPath}}`

**Quick Links:**

- [Package Overview]({{.Name}}.md)
- [API Reference](../api-reference/{{.Name}}.md)
  {{- if hasExamples .}}
- [Examples](../examples/{{.Name}}/README.md)
  {{- end}}

{{- end}}
{{- end}}

## Getting Started

1. **Install the package:**

   ```bash
   go get {{.Repository.ImportPath}}@latest
   ```

2. **Import in your Go code:**

   ```go
   import "{{.Repository.ImportPath}}"
   ```

3. **See package-specific documentation above for usage examples.**

## External Resources

- [GitHub Repository]({{.Repository.URL}})
- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Repository.ImportPath}})
- [Issues & Support]({{.Repository.URL}}/issues)
