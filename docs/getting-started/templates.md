# Getting Started with templates



## Overview

**Import Path:** `github.com/kolosys/proton/internal/templates`



## Installation

### Install the package

```bash
go get github.com/kolosys/proton/internal/templates
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/templates"
)

func main() {
    fmt.Println("templates package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with templates:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kolosys/proton/internal/templates"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from templates!")
}
```

## Key Features

### Types

- **Context** - Context provides data for template rendering

- **Engine** - Engine handles template rendering for documentation generation

- **PackageContext** - PackageContext provides package-specific data for template rendering

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/templates.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/templates/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/templates)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/templates)
- [GitHub Issues](https://github.com/kolosys/proton/issues)
