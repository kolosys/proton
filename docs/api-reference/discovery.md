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
    config: &/* value */{},
    projectPath: "example",
    fileSet: &/* value */{},
}
```

#### Type Definition

```go
type Discoverer struct {
    config *config.Config
    projectPath string
    fileSet *token.FileSet
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| config | `*config.Config` |  |
| projectPath | `string` |  |
| fileSet | `*token.FileSet` |  |

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

### autoDiscoverPackages

autoDiscoverPackages automatically discovers packages using the configured patterns

```go
func (*Discoverer) autoDiscoverPackages() ([]*PackageInfo, error)
```

**Parameters:**
  None

**Returns:**
- []*PackageInfo
- error

### determinePackageCategory

determinePackageCategory determines the category of a package for organization

```go
func (*Discoverer) determinePackageCategory(pkg *PackageInfo) string
```

**Parameters:**
- `pkg` (*PackageInfo)

**Returns:**
- string

### enhanceFunction

enhanceFunction extracts detailed parameter and return information from a function

```go
func (*Discoverer) enhanceFunction(fn *doc.Func, astPkg *ast.Package) *EnhancedFunc
```

**Parameters:**
- `fn` (*doc.Func)
- `astPkg` (*ast.Package)

**Returns:**
- *EnhancedFunc

### enhanceType

enhanceType extracts detailed field and method information from a type

```go
func (*Discoverer) enhanceType(typ *doc.Type, astPkg *ast.Package) *EnhancedType
```

**Parameters:**
- `typ` (*doc.Type)
- `astPkg` (*ast.Package)

**Returns:**
- *EnhancedType

### extractExamples

extractExamples extracts example functions from the package

```go
func (*Discoverer) extractExamples(astPkg *ast.Package) []*doc.Example
```

**Parameters:**
- `astPkg` (*ast.Package)

**Returns:**
- []*doc.Example

### extractFieldDoc

extractFieldDoc extracts field documentation from struct field

```go
func (*Discoverer) extractFieldDoc(field *ast.Field) string
```

**Parameters:**
- `field` (*ast.Field)

**Returns:**
- string

### extractFunctionDocumentation

extractFunctionDocumentation performs custom AST traversal to extract function documentation

```go
func (*Discoverer) extractFunctionDocumentation(funcName string, astPkg *ast.Package) string
```

**Parameters:**
- `funcName` (string)
- `astPkg` (*ast.Package)

**Returns:**
- string

### extractInterfaceMethodDoc

extractInterfaceMethodDoc extracts documentation for interface methods

```go
func (*Discoverer) extractInterfaceMethodDoc(typ *doc.Type, methodName string) string
```

**Parameters:**
- `typ` (*doc.Type)
- `methodName` (string)

**Returns:**
- string

### extractMainDescription

extractMainDescription extracts the main description part before Parameters/Returns sections

```go
func (*Discoverer) extractMainDescription(doc string) string
```

**Parameters:**
- `doc` (string)

**Returns:**
- string

### extractManualComments

extractManualComments tries to find comments manually by looking at the file content

```go
func (*Discoverer) extractManualComments(file *ast.File, pos token.Pos) string
```

**Parameters:**
- `file` (*ast.File)
- `pos` (token.Pos)

**Returns:**
- string

### extractParamDoc

extractParamDoc extracts parameter documentation from function documentation

```go
func (*Discoverer) extractParamDoc(doc, paramName string) string
```

**Parameters:**
- `doc` (string)
- `paramName` (string)

**Returns:**
- string

### extractReturnDoc

extractReturnDoc extracts return value documentation from function documentation

```go
func (*Discoverer) extractReturnDoc(doc string, index int) string
```

**Parameters:**
- `doc` (string)
- `index` (int)

**Returns:**
- string

### extractTypeDoc

extractTypeDoc extracts documentation from AST comments for types

```go
func (*Discoverer) extractTypeDoc(commentGroup *ast.CommentGroup) string
```

**Parameters:**
- `commentGroup` (*ast.CommentGroup)

**Returns:**
- string

### extractTypeDocumentation

extractTypeDocumentation performs custom AST traversal to extract type documentation

```go
func (*Discoverer) extractTypeDocumentation(typeName string, astPkg *ast.Package) string
```

**Parameters:**
- `typeName` (string)
- `astPkg` (*ast.Package)

**Returns:**
- string

### formatFuncSignature

formatFuncSignature formats a function signature

```go
func (*Discoverer) formatFuncSignature(funcType *ast.FuncType) string
```

**Parameters:**
- `funcType` (*ast.FuncType)

**Returns:**
- string

### formatType

formatType converts an AST type expression to a string

```go
func (*Discoverer) formatType(expr ast.Expr) string
```

**Parameters:**
- `expr` (ast.Expr)

**Returns:**
- string

### generateExampleCode

generateExampleCode generates basic example code for a function

```go
func (*Discoverer) generateExampleCode(fn *doc.Func) string
```

**Parameters:**
- `fn` (*doc.Func)

**Returns:**
- string

### generateExampleValue

generateExampleValue generates an example value for a type

```go
func (*Discoverer) generateExampleValue(expr ast.Expr) string
```

**Parameters:**
- `expr` (ast.Expr)

**Returns:**
- string

### generateFunctionDeclaration

generateFunctionDeclaration creates a clean function declaration

```go
func (*Discoverer) generateFunctionDeclaration(fn *doc.Func, funcDecl *ast.FuncDecl) string
```

**Parameters:**
- `fn` (*doc.Func)
- `funcDecl` (*ast.FuncDecl)

**Returns:**
- string

### generateInterfaceDeclaration

generateInterfaceDeclaration creates a clean interface declaration

```go
func (*Discoverer) generateInterfaceDeclaration(name string, interfaceType *ast.InterfaceType) string
```

**Parameters:**
- `name` (string)
- `interfaceType` (*ast.InterfaceType)

**Returns:**
- string

### generateStructDeclaration

generateStructDeclaration creates a clean struct declaration

```go
func (*Discoverer) generateStructDeclaration(name string, structType *ast.StructType) string
```

**Parameters:**
- `name` (string)
- `structType` (*ast.StructType)

**Returns:**
- string

### generateTypeExample

generateTypeExample generates a usage example for a type

```go
func (*Discoverer) generateTypeExample(typ *doc.Type, typeSpec *ast.TypeSpec) string
```

**Parameters:**
- `typ` (*doc.Type)
- `typeSpec` (*ast.TypeSpec)

**Returns:**
- string

### getImportPath

getImportPath determines the import path for a package

```go
func (*Discoverer) getImportPath(pkgPath string) string
```

**Parameters:**
- `pkgPath` (string)

**Returns:**
- string

### parseASTPackage

parseASTPackage creates PackageInfo from an AST package

```go
func (*Discoverer) parseASTPackage(astPkg *ast.Package, pkgPath string) (*PackageInfo, error)
```

**Parameters:**
- `astPkg` (*ast.Package)
- `pkgPath` (string)

**Returns:**
- *PackageInfo
- error

### parsePackage

parsePackage parses a single package from a given path

```go
func (*Discoverer) parsePackage(pkgPath string) (*PackageInfo, error)
```

**Parameters:**
- `pkgPath` (string)

**Returns:**
- *PackageInfo
- error

### shouldExcludePackage

shouldExcludePackage checks if a package should be excluded based on patterns

```go
func (*Discoverer) shouldExcludePackage(pkgPath string) bool
```

**Parameters:**
- `pkgPath` (string)

**Returns:**
- bool

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
