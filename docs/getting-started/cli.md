# Getting Started with cli



## Overview

**Import Path:** `github.com/kolosys/proton/internal/cli`



## Installation

### Install the package

```bash
go get github.com/kolosys/proton/internal/cli
```

### Verify installation

Create a simple test file to verify the package works:

```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/cli"
)

func main() {
    fmt.Println("cli package imported successfully!")
}
```

Run it:

```bash
go run main.go
```

## Quick Start

Here's a basic example to get you started with cli:

```go
package main

import (
    "fmt"
    "log"

    "github.com/kolosys/proton/internal/cli"
)

func main() {
    // TODO: Add basic usage example
    fmt.Println("Hello from cli!")
}
```

## Key Features

### Functions

- **Execute** - Execute adds all child commands to the root command and sets flags appropriately. It initializes the CLI application and runs the root command. Returns error if command execution fails.

## Usage Examples

For more detailed examples, see the [Examples](../examples/README.md) section.

## Next Steps

- [Full API Reference](../api-reference/cli.md) - Complete API documentation
- [Examples](../examples/README.md) - Working examples and tutorials
- [Best Practices](../guides/cli/best-practices.md) - Recommended patterns and usage

## Documentation Links

- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/cli)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/cli)
- [GitHub Issues](https://github.com/kolosys/proton/issues)
