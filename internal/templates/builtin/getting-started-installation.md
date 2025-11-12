# Installation

This guide will help you install and set up {{.Repository.Name}} in your Go project.

## Prerequisites

Before installing {{.Repository.Name}}, ensure you have:

- **Go {{.Metadata.GoVersion}}** or later installed
- A Go module initialized in your project (run `go mod init` if needed)
- Access to the GitHub repository (for private repositories)

## Installation Steps

### Step 1: Install the Package

Use `go get` to install {{.Repository.Name}}:

```bash
go get {{.Repository.ImportPath}}
```

This will download the package and add it to your `go.mod` file.

### Step 2: Import in Your Code

Import the package in your Go source files:

```go
import "{{.Repository.ImportPath}}"
```

{{- if gt (len .Packages) 1}}

### Multiple Packages

{{.Repository.Name}} includes several packages. Import the ones you need:

{{- range .Packages}}
{{- if ne .Name "main"}}

```go
// {{.Description}}
import "{{.ImportPath}}"
```
{{- end}}
{{- end}}
{{- end}}

### Step 3: Verify Installation

Create a simple test file to verify the installation:

```go
package main

import (
    "fmt"
    "{{.Repository.ImportPath}}"
)

func main() {
    fmt.Println("{{.Repository.Name}} installed successfully!")
}
```

Run the test:

```bash
go run main.go
```

## Updating the Package

To update to the latest version:

```bash
go get -u {{.Repository.ImportPath}}
```

To update to a specific version:

```bash
go get {{.Repository.ImportPath}}@v1.2.3
```

## Installing a Specific Version

To install a specific version of the package:

```bash
go get {{.Repository.ImportPath}}@v1.0.0
```

Check available versions on the [GitHub releases page]({{.Repository.URL}}/releases).

## Development Setup

If you want to contribute or modify the library:

### Clone the Repository

```bash
git clone {{.Repository.URL}}.git
cd {{.Repository.Name}}
```

### Install Dependencies

```bash
go mod download
```

### Run Tests

```bash
go test ./...
```

## Troubleshooting

### Module Not Found

If you encounter a "module not found" error:

1. Ensure your `GOPATH` is set correctly
2. Check that you have network access to GitHub
3. Try running `go clean -modcache` and reinstall

### Private Repository Access

For private repositories, configure Git to use SSH or a personal access token:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Or set up GOPRIVATE:

```bash
export GOPRIVATE={{.Repository.ImportPath}}
```

## Next Steps

Now that you have {{.Repository.Name}} installed, check out the [Quick Start Guide](quick-start.md) to learn how to use it.

## Additional Resources

- [Quick Start Guide](quick-start.md)
- [API Reference](../reference/api-reference/README.md)
- [Examples](../reference/examples/README.md)

