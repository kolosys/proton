# Getting Started with {{.Package.Name}}

{{.Package.Description}}

## Overview

**Import Path:** `{{.Package.ImportPath}}`

{{.Package.Doc.Doc}}

## Installation

### Install the package

```bash
go get {{.Package.ImportPath}}
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "{{.Package.ImportPath}}"
)

func main() {
    fmt.Println("{{.Package.Name}} package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with {{.Package.Name}}:

```go
package main

import (
    "fmt"
    "log"

    "{{.Package.ImportPath}}"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from {{.Package.Name}}!")
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

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/{{.Package.Name}}.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/{{.Package.Name}}/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Package.ImportPath}})
- [Source Code]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
- [GitHub Issues]({{.Repository.URL}}/issues)
