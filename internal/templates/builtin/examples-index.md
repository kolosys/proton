# Examples

Working examples and code samples for {{.Repository.Name}}.

## Overview

This section contains practical examples demonstrating how to use the various packages and features.

## Package Examples

{{- range .Packages}}
{{- if hasExamples .}}

### [{{.Name}}]({{.Name}}/README.md)

{{.Description}}

{{- if .Examples}}
{{- range .Examples}}

- [{{.Name}}]({{$.Name}}/README.md#{{lower .Name}})
  {{- end}}
  {{- end}}

{{- end}}
{{- end}}

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone {{.Repository.URL}}.git
   cd {{.Repository.Name}}
   ```

2. **Navigate to examples:**

   ```bash
   cd examples
   ```

3. **Run an example:**
   ```bash
   go run example-name/main.go
   ```

## Contributing Examples

We welcome contributions of new examples! See our [Contributing Guide](../guides/contributing.md) for details on how to add examples.

### Example Guidelines

- Include clear documentation and comments
- Demonstrate real-world use cases
- Keep examples focused and concise
- Include expected output when applicable

## External Resources

- [GitHub Repository]({{.Repository.URL}})
- [API Reference](../api-reference/README.md)
- [Package Documentation](../packages/README.md)
