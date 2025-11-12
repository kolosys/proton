# Quick Start

This guide will help you get started with proton quickly with a basic example.

## Basic Usage

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "github.com/kolosys/proton/internal/cli"
    "github.com/kolosys/proton/internal/config"
    "github.com/kolosys/proton/internal/discovery"
    "github.com/kolosys/proton/internal/generator"
    "github.com/kolosys/proton/internal/interfaces"
    "github.com/kolosys/proton/internal/templates"
)

func main() {
    // Basic usage example
    fmt.Println("Welcome to proton!")
    
    // TODO: Add your code here
}
```

## Common Use Cases

### Using cli

**Import Path:** `github.com/kolosys/proton/internal/cli`



```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/cli"
)

func main() {
    // Example usage of cli
    fmt.Println("Using cli package")
}
```

#### Available Functions
- **Execute** - Execute adds all child commands to the root command and sets flags appropriately. It initializes the CLI application and runs the root command. Returns error if command execution fails.

For detailed API documentation, see the [cli API Reference](../reference/api-reference/cli.md).

### Using config

**Import Path:** `github.com/kolosys/proton/internal/config`



```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/config"
)

func main() {
    // Example usage of config
    fmt.Println("Using config package")
}
```

#### Available Types
- **APIGeneration** - 
- **Config** - Config represents the complete configuration for Proton
- **CustomGuide** - 
- **CustomTemplate** - 
- **Discovery** - 
- **Examples** - 
- **Generation** - 
- **Guides** - 
- **Metadata** - 
- **Output** - 
- **Package** - 
- **Packages** - 
- **Repository** - 
- **Templates** - 

For detailed API documentation, see the [config API Reference](../reference/api-reference/config.md).

### Using discovery

**Import Path:** `github.com/kolosys/proton/internal/discovery`



```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/discovery"
)

func main() {
    // Example usage of discovery
    fmt.Println("Using discovery package")
}
```

#### Available Types
- **Discoverer** - Discoverer handles package discovery and parsing
- **EnhancedFunc** - EnhancedFunc extends doc.Func with additional parameter and return information
- **EnhancedType** - EnhancedType extends doc.Type with enhanced field information
- **Field** - Field represents a struct field
- **PackageInfo** - PackageInfo contains information about a discovered Go package
- **Parameter** - Parameter represents a function parameter
- **Result** - Result represents a function return value

For detailed API documentation, see the [discovery API Reference](../reference/api-reference/discovery.md).

### Using generator

**Import Path:** `github.com/kolosys/proton/internal/generator`



```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/generator"
)

func main() {
    // Example usage of generator
    fmt.Println("Using generator package")
}
```

#### Available Types
- **Generator** - Generator handles the complete documentation generation process

For detailed API documentation, see the [generator API Reference](../reference/api-reference/generator.md).

### Using interfaces

**Import Path:** `github.com/kolosys/proton/internal/interfaces`

Package interfaces provides core interfaces for Proton documentation generation.


```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/interfaces"
)

func main() {
    // Example usage of interfaces
    fmt.Println("Using interfaces package")
}
```

#### Available Types
- **Capsule** - Capsule represents a time-locked value
- **Codec** - Codec defines how to serialize/deserialize values
- **Configuration** - Configuration holds system configuration
- **Generator** - Generator defines the interface for documentation generators
- **Parser** - Parser defines the interface for parsing Go source code
- **Template** - Template represents a documentation template
- **Writer** - Writer provides file writing capabilities

For detailed API documentation, see the [interfaces API Reference](../reference/api-reference/interfaces.md).

### Using templates

**Import Path:** `github.com/kolosys/proton/internal/templates`



```go
package main

import (
    "fmt"
    "github.com/kolosys/proton/internal/templates"
)

func main() {
    // Example usage of templates
    fmt.Println("Using templates package")
}
```

#### Available Types
- **Context** - Context provides data for template rendering
- **Engine** - Engine handles template rendering for documentation generation
- **PackageContext** - PackageContext provides package-specific data for template rendering

For detailed API documentation, see the [templates API Reference](../reference/api-reference/templates.md).

## Step-by-Step Tutorial

### Step 1: Import the Package

First, import the necessary packages in your Go file:

```go
import (
    "fmt"
    "github.com/kolosys/proton/internal/cli"
    "github.com/kolosys/proton/internal/config"
    "github.com/kolosys/proton/internal/discovery"
    "github.com/kolosys/proton/internal/generator"
    "github.com/kolosys/proton/internal/interfaces"
    "github.com/kolosys/proton/internal/templates"
)
```

### Step 2: Initialize

Set up the basic configuration:

```go
func main() {
    // Initialize your application
    fmt.Println("Initializing proton...")
}
```

### Step 3: Use the Library

Implement your specific use case:

```go
func main() {
    // Your implementation here
}
```

## Running Your Code

To run your Go program:

```bash
go run main.go
```

To build an executable:

```bash
go build -o myapp
./myapp
```

## Configuration Options

proton can be configured to suit your needs. Check the [Core Concepts](../core-concepts/) section for detailed information about configuration options.

## Error Handling

Always handle errors appropriately:

```go
result, err := someFunction()
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Best Practices

- Always handle errors returned by library functions
- Check the API documentation for detailed parameter information
- Use meaningful variable and function names
- Add comments to document your code

## Complete Example

Here's a complete working example:

```go
package main

import (
    "fmt"
    "log"
    "github.com/kolosys/proton/internal/cli"
    "github.com/kolosys/proton/internal/config"
    "github.com/kolosys/proton/internal/discovery"
    "github.com/kolosys/proton/internal/generator"
    "github.com/kolosys/proton/internal/interfaces"
    "github.com/kolosys/proton/internal/templates"
)

func main() {
    fmt.Println("Starting proton application...")
    
    // Add your implementation here
    
    fmt.Println("Application completed successfully!")
}
```

## Next Steps

Now that you've seen the basics, explore:

- **[Core Concepts](../core-concepts/)** - Understanding the library architecture
- **[API Reference](../reference/api-reference/README.md)** - Complete API documentation
- **[Examples](../reference/examples/README.md)** - More detailed examples
- **[Advanced Topics](../advanced/)** - Performance tuning and advanced patterns

## Getting Help

If you run into issues:

1. Check the [API Reference](../reference/api-reference/README.md)
2. Browse the [Examples](../reference/examples/README.md)
3. Visit the [GitHub Issues](https://github.com/kolosys/proton/issues) page

