# templates API

Complete API documentation for the templates package.

**Import Path:** `github.com/kolosys/proton/internal/templates`

## Package Documentation



## Types

### BenchmarkContext
BenchmarkContext provides benchmark-specific data for template rendering

#### Example Usage

```go
// Create a new BenchmarkContext
benchmarkcontext := BenchmarkContext{
    Benchmarks: [],
}
```

#### Type Definition

```go
type BenchmarkContext struct {
    *PackageContext
    Benchmarks []*BenchmarkResult `json:"benchmarks"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| **PackageContext | `*PackageContext` |  |
| Benchmarks | `[]*BenchmarkResult` |  |

### BenchmarkResult
BenchmarkResult represents benchmark data for template rendering

#### Example Usage

```go
// Create a new BenchmarkResult
benchmarkresult := BenchmarkResult{
    Name: "example",
    NsPerOp: 42,
    BytesPerOp: 42,
    AllocsPerOp: 42,
    Runs: 42,
}
```

#### Type Definition

```go
type BenchmarkResult struct {
    Name string `json:"name"`
    NsPerOp int64 `json:"ns_per_op"`
    BytesPerOp int64 `json:"bytes_per_op"`
    AllocsPerOp int64 `json:"allocs_per_op"`
    Runs int `json:"runs"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| NsPerOp | `int64` |  |
| BytesPerOp | `int64` |  |
| AllocsPerOp | `int64` |  |
| Runs | `int` |  |

### Context
Context provides data for template rendering

#### Example Usage

```go
// Create a new Context
context := Context{
    Repository: /* value */,
    Packages: [],
    Config: &/* value */{},
    Metadata: /* value */,
}
```

#### Type Definition

```go
type Context struct {
    Repository config.Repository `json:"repository"`
    Packages []*discovery.PackageInfo `json:"packages"`
    Config *config.Config `json:"config"`
    Metadata config.Metadata `json:"metadata"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Repository | `config.Repository` |  |
| Packages | `[]*discovery.PackageInfo` |  |
| Config | `*config.Config` |  |
| Metadata | `config.Metadata` |  |

### Engine
Engine handles template rendering for documentation generation

#### Example Usage

```go
// Create a new Engine
engine := Engine{

}
```

#### Type Definition

```go
type Engine struct {
}
```

### Constructor Functions

### New

New creates a new template engine

```go
func New(cfg *config.Config, projectPath string) (*Engine, error)
```

**Parameters:**
- `cfg` (*config.Config)
- `projectPath` (string)

**Returns:**
- *Engine
- error

## Methods

### HasTemplate

HasTemplate checks if a template exists

```go
func (*Engine) HasTemplate(name string) bool
```

**Parameters:**
- `name` (string)

**Returns:**
- bool

### ListTemplates

ListTemplates returns a list of available template names

```go
func (*Engine) ListTemplates() []string
```

**Parameters:**
  None

**Returns:**
- []string

### RenderToFile

RenderToFile renders a template to a file

```go
func (*Engine) RenderToFile(templateName string, data interface{}, outputPath string) error
```

**Parameters:**
- `templateName` (string)
- `data` (interface{})
- `outputPath` (string)

**Returns:**
- error

### RenderToString

RenderToString renders a template to a string

```go
func (*Engine) RenderToString(templateName string, data interface{}) (string, error)
```

**Parameters:**
- `templateName` (string)
- `data` (interface{})

**Returns:**
- string
- error

### PackageContext
PackageContext provides package-specific data for template rendering

#### Example Usage

```go
// Create a new PackageContext
packagecontext := PackageContext{
    Package: &/* value */{},
}
```

#### Type Definition

```go
type PackageContext struct {
    *Context
    Package *discovery.PackageInfo `json:"package"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| **Context | `*Context` |  |
| Package | `*discovery.PackageInfo` |  |

## External Links

- [Package Overview](../packages/templates.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/templates)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/templates)
