# {{.Package.Name}} Package

{{.Package.Description}}

## Overview

**Import Path:** `{{.Package.ImportPath}}`

{{.Package.Doc.Doc}}

## Installation

```bash
go get {{.Package.ImportPath}}
```

## Quick Start

```go
package main

import (
    "{{.Package.ImportPath}}"
)

func main() {
    // TODO: Add basic usage example
}
```

## Key Features

{{- if .Package.Types}}

### Types

{{- range .Package.Types}}

- **{{.Name}}** - {{.Doc}}
  {{- end}}
  {{- end}}

{{- if .Package.Functions}}

### Functions

{{- range .Package.Functions}}

- **{{.Name}}** - {{.Doc}}
  {{- end}}
  {{- end}}

## Documentation Links

- [Full API Reference](../api-reference/{{.Package.Name}}.md)
  {{- if hasExamples .Package}}
- [Examples](../examples/{{.Package.Name}}/README.md)
  {{- end}}
- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Package.ImportPath}})

## Source Code

- [View on GitHub]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
- [Browse Files]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
