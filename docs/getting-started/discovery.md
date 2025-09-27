# Getting Started with discovery



## Overview

**Import Path:** `github.com/kolosys/proton/internal/discovery`



## Installation

### Install the package

```bash
go get github.com/kolosys/proton/internal/discovery
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/discovery"
)

func main() {
    fmt.Println("discovery package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with discovery:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kolosys/proton/internal/discovery"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from discovery!")
}
```

## Key Features

### Types

- **Discoverer** - Discoverer handles package discovery and parsing

- **EnhancedFunc** - EnhancedFunc extends doc.Func with additional parameter and return information

- **EnhancedType** - EnhancedType extends doc.Type with enhanced field information

- **Field** - Field represents a struct field

- **PackageInfo** - PackageInfo contains information about a discovered Go package

- **Parameter** - Parameter represents a function parameter

- **Result** - Result represents a function return value

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/discovery.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/discovery/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/discovery)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/discovery)
- [GitHub Issues](https://github.com/kolosys/proton/issues)
