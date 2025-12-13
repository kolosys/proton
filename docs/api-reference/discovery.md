# discovery API

Complete API documentation for the discovery package.

**Import Path:** `github.com/kolosys/proton/internal/discovery`

## Package Documentation



## Types

### Discoverer
Discoverer handles package discovery and parsing

#### Example Usage

```go
// Create a new Discoverer
discoverer := Discoverer{

}
```

#### Type Definition

```go
type Discoverer struct {
}
```

### Constructor Functions

### New

New creates a new package discoverer

```go
func New(cfg *config.Config, projectPath string) *Discoverer
```

**Parameters:**
- `cfg` (*config.Config)
- `projectPath` (string)

**Returns:**
- *Discoverer

## Methods

### DiscoverPackages

DiscoverPackages discovers all packages in the project according to configuration

```go
func (*Discoverer) DiscoverPackages() ([]*PackageInfo, error)
```

**Parameters:**
  None

**Returns:**
- []*PackageInfo
- error

### GetPackagesByCategory

GetPackagesByCategory categorizes packages for easier documentation generation

```go
func (*Discoverer) GetPackagesByCategory(packages []*PackageInfo) map[string][]*PackageInfo
```

**Parameters:**
- `packages` ([]*PackageInfo)

**Returns:**
- map[string][]*PackageInfo

### EnhancedFunc
EnhancedFunc extends doc.Func with additional parameter and return information

#### Example Usage

```go
// Create a new EnhancedFunc
enhancedfunc := EnhancedFunc{
    Params: [],
    Results: [],
    ExampleCode: "example",
    Declaration: "example",
    Doc: "example",
}
```

#### Type Definition

```go
type EnhancedFunc struct {
    *doc.Func
    Params []*Parameter
    Results []*Result
    ExampleCode string
    Declaration string
    Doc string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| **doc.Func | `*doc.Func` |  |
| Params | `[]*Parameter` |  |
| Results | `[]*Result` |  |
| ExampleCode | `string` |  |
| Declaration | `string` | Clean formatted function declaration |
| Doc | `string` | Enhanced documentation (may override doc.Func.Doc) |

### EnhancedType
EnhancedType extends doc.Type with enhanced field information

#### Example Usage

```go
// Create a new EnhancedType
enhancedtype := EnhancedType{
    Fields: [],
    Methods: [],
    Funcs: [],
    TypeKind: "example",
    Declaration: "example",
    Doc: "example",
    ExampleCode: "example",
}
```

#### Type Definition

```go
type EnhancedType struct {
    *doc.Type
    Fields []*Field
    Methods []*EnhancedFunc
    Funcs []*EnhancedFunc
    TypeKind string
    Declaration string
    Doc string
    ExampleCode string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| **doc.Type | `*doc.Type` |  |
| Fields | `[]*Field` |  |
| Methods | `[]*EnhancedFunc` |  |
| Funcs | `[]*EnhancedFunc` |  |
| TypeKind | `string` | struct, interface, type alias, etc. |
| Declaration | `string` | Clean formatted declaration |
| Doc | `string` | Enhanced documentation (may override doc.Type.Doc) |
| ExampleCode | `string` | Usage example code |

### Field
Field represents a struct field

#### Example Usage

```go
// Create a new Field
field := Field{
    Name: "example",
    Type: "example",
    Tag: "example",
    Doc: "example",
}
```

#### Type Definition

```go
type Field struct {
    Name string
    Type string
    Tag string
    Doc string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Type | `string` |  |
| Tag | `string` |  |
| Doc | `string` |  |

### PackageInfo
PackageInfo contains information about a discovered Go package

#### Example Usage

```go
// Create a new PackageInfo
packageinfo := PackageInfo{
    Name: "example",
    Path: "example",
    ImportPath: "example",
    Description: "example",
    Doc: &/* value */{},
    Functions: [],
    Types: [],
    Variables: [],
    Constants: [],
    Examples: [],
    Files: [],
}
```

#### Type Definition

```go
type PackageInfo struct {
    Name string
    Path string
    ImportPath string
    Description string
    Doc *doc.Package
    Functions []*EnhancedFunc
    Types []*EnhancedType
    Variables []*doc.Value
    Constants []*doc.Value
    Examples []*doc.Example
    Files []string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Path | `string` |  |
| ImportPath | `string` |  |
| Description | `string` |  |
| Doc | `*doc.Package` |  |
| Functions | `[]*EnhancedFunc` |  |
| Types | `[]*EnhancedType` |  |
| Variables | `[]*doc.Value` |  |
| Constants | `[]*doc.Value` |  |
| Examples | `[]*doc.Example` |  |
| Files | `[]string` |  |

### Parameter
Parameter represents a function parameter

#### Example Usage

```go
// Create a new Parameter
parameter := Parameter{
    Name: "example",
    Type: "example",
    Doc: "example",
}
```

#### Type Definition

```go
type Parameter struct {
    Name string
    Type string
    Doc string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Type | `string` |  |
| Doc | `string` |  |

### Result
Result represents a function return value

#### Example Usage

```go
// Create a new Result
result := Result{
    Name: "example",
    Type: "example",
    Doc: "example",
}
```

#### Type Definition

```go
type Result struct {
    Name string
    Type string
    Doc string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Type | `string` |  |
| Doc | `string` |  |

## External Links

- [Package Overview](../packages/discovery.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/discovery)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/discovery)
