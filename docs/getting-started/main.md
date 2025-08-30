# Getting Started with main



## Overview

**Import Path:** `github.com/kolosys/proton/examples/basic`



## Installation

### Install the package

```bash
go get github.com/kolosys/proton/examples/basic
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/examples/basic"
)

func main() {
    fmt.Println("main package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with main:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kolosys/proton/examples/basic"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from main!")
}
```

## Key Features

### Functions

- **ExampleHello** - ExampleHello demonstrates basic usage of the library

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/main.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/main/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/examples/basic)
- [Source Code](https://github.com/kolosys/proton/tree/main/examples/basic)
- [GitHub Issues](https://github.com/kolosys/proton/issues)
