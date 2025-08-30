# Getting Started

This guide will help you get up and running quickly with {{.Repository.Name}}.

## Installation

### Requirements

- Go 1.21 or later
- No external dependencies required

### Install via go get

```bash
go get {{.Repository.ImportPath}}@latest
```

### Install specific version

```bash
go get {{.Repository.ImportPath}}@v0.1.0
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "{{.Repository.ImportPath}}"
)

func main() {
    // Your code here
    fmt.Println("Hello from {{.Repository.Name}}!")
}
```

## Available Packages

{{.Repository.Name}} provides the following packages:

{{- range .Packages}}

### [{{.Name}}]({{.Name}}.md)

{{.Description}}

**Quick Links:**

- [Getting Started]({{.Name}}.md) - Installation and getting started
- [API Reference](../api-reference/{{.Name}}.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples
- [Best Practices](../guides/{{.Name}}/best-practices.md) - Recommended patterns

{{- end}}

## Next Steps

- [API Reference](../api-reference/README.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Guides](../guides/README.md) - In-depth guides and best practices
- [GitHub Repository]({{.Repository.URL}})

## Need Help?

- [GitHub Issues]({{.Repository.URL}}/issues)
- [GitHub Discussions]({{.Repository.URL}}/discussions)
- [FAQ](../guides/faq.md)
