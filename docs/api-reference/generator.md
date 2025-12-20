# generator API

Complete API documentation for the generator package.

**Import Path:** `github.com/kolosys/proton/internal/generator`

## Package Documentation



## Types

### Generator
Generator handles the complete documentation generation process

#### Example Usage

```go
// Create a new Generator
generator := Generator{

}
```

#### Type Definition

```go
type Generator struct {
}
```

### Constructor Functions

### New

New creates a new documentation generator

```go
func New(cfg *config.Config, projectPath string) (*Generator, error)
```

**Parameters:**
- `cfg` (*config.Config)
- `projectPath` (string)

**Returns:**
- *Generator
- error

## Methods

### Generate

Generate performs the complete documentation generation

```go
func (*Generator) Generate() error
```

**Parameters:**
  None

**Returns:**
- error

## External Links

- [Package Overview](../packages/generator.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/generator)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/generator)
