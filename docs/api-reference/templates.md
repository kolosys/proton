# templates API

Complete API documentation for the templates package.

**Import Path:** `github.com/kolosys/proton/internal/templates`

## Package Documentation



## Types

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
