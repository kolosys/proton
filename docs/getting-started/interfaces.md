# Getting Started with interfaces

Package interfaces provides core interfaces for Proton documentation generation.


## Overview

**Import Path:** `github.com/kolosys/proton/internal/interfaces`

Package interfaces provides core interfaces for Proton documentation generation.


## Installation

### Install the package

```bash
go get github.com/kolosys/proton/internal/interfaces
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/interfaces"
)

func main() {
    fmt.Println("interfaces package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with interfaces:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kolosys/proton/internal/interfaces"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from interfaces!")
}
```

## Key Features

### Types

- **Capsule** - Capsule represents a time-locked value

- **Codec** - Codec defines how to serialize/deserialize values

- **Configuration** - Configuration holds system configuration

- **Generator** - Generator defines the interface for documentation generators

- **Parser** - Parser defines the interface for parsing Go source code

- **Template** - Template represents a documentation template

- **Writer** - Writer provides file writing capabilities

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/interfaces.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/interfaces/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/interfaces)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/interfaces)
- [GitHub Issues](https://github.com/kolosys/proton/issues)
