# {{.Package.Name}} Examples

Examples and code samples for the {{.Package.Name}} package.

## Overview

{{.Package.Description}}

**Import Path:** `{{.Package.ImportPath}}`

{{- if .Package.Examples}}

## Examples

{{- range .Package.Examples}}

### {{.Name}}

{{- if .Doc}}
{{.Doc}}
{{- end}}

```go
{{.Code}}
```

{{- if .Output}}
{{formatExampleOutput .Output}}
{{- end}}

{{- end}}
{{- end}}

## Getting Started

```bash
# Install the package
go get {{.Package.ImportPath}}
```

```go
// Basic usage
package main

import (
    "{{.Package.ImportPath}}"
)

func main() {
    // Your code here
}
```

## More Examples

For more examples and usage patterns:

- [API Reference](../../api-reference/{{.Package.Name}}.md)
- [Package Documentation](../../packages/{{.Package.Name}}.md)
- [pkg.go.dev Examples](https://pkg.go.dev/{{.Package.ImportPath}}#pkg-examples)

## Source Code

- [View Source]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
- [Browse Examples]({{.Repository.URL}}/tree/{{.Repository.Branch}}/examples)
