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
    config: &/* value */{},
    projectPath: "example",
    outputPath: "example",
    discoverer: &/* value */{},
    templates: &/* value */{},
}
```

#### Type Definition

```go
type Generator struct {
    config *config.Config
    projectPath string
    outputPath string
    discoverer *discovery.Discoverer
    templates *templates.Engine
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| config | `*config.Config` |  |
| projectPath | `string` |  |
| outputPath | `string` |  |
| discoverer | `*discovery.Discoverer` |  |
| templates | `*templates.Engine` |  |

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

### cleanOutputDirectory

cleanOutputDirectory removes all files from the output directory Only removes auto-generated directories (api-reference and examples)

```go
func (*Generator) cleanOutputDirectory() error
```

**Parameters:**
  None

**Returns:**
- error

### createTemplateContext

createTemplateContext creates the context object for template rendering

```go
func (*Generator) createTemplateContext(packages []*discovery.PackageInfo) *templates.Context
```

**Parameters:**
- `packages` ([]*discovery.PackageInfo)

**Returns:**
- *templates.Context

### discoverExampleDirectories

discoverExampleDirectories discovers example directories based on configuration

```go
func (*Generator) discoverExampleDirectories() ([]string, error)
```

**Parameters:**
  None

**Returns:**
- []string
- error

### generateAPIDocumentation

generateAPIDocumentation generates API reference documentation

```go
func (*Generator) generateAPIDocumentation(packages []*discovery.PackageInfo, context *templates.Context) error
```

**Parameters:**
- `packages` ([]*discovery.PackageInfo)
- `context` (*templates.Context)

**Returns:**
- error

### generateAdvancedSection

generateAdvancedSection generates advanced section with placeholder files

```go
func (*Generator) generateAdvancedSection(context *templates.Context) error
```

**Parameters:**
- `context` (*templates.Context)

**Returns:**
- error

### generateCoreConcepts

generateCoreConcepts generates core-concepts placeholders for each package

```go
func (*Generator) generateCoreConcepts(packages []*discovery.PackageInfo, context *templates.Context) error
```

**Parameters:**
- `packages` ([]*discovery.PackageInfo)
- `context` (*templates.Context)

**Returns:**
- error

### generateExampleDirectoryDocumentation

generateExampleDirectoryDocumentation generates documentation for a single example directory

```go
func (*Generator) generateExampleDirectoryDocumentation(sourceDir, outputBaseDir string, context *templates.Context) error
```

**Parameters:**
- `sourceDir` (string)
- `outputBaseDir` (string)
- `context` (*templates.Context)

**Returns:**
- error

### generateExampleDirectoryREADME

generateExampleDirectoryREADME generates a README for an example directory

```go
func (*Generator) generateExampleDirectoryREADME(sourceDir, outputDir, relPath string, context *templates.Context) error
```

**Parameters:**
- `sourceDir` (string)
- `outputDir` (string)
- `relPath` (string)
- `context` (*templates.Context)

**Returns:**
- error

### generateExampleFileMarkdown

generateExampleFileMarkdown generates markdown documentation for a single example file

```go
func (*Generator) generateExampleFileMarkdown(sourcePath, outputDir, fileName string) error
```

**Parameters:**
- `sourcePath` (string)
- `outputDir` (string)
- `fileName` (string)

**Returns:**
- error

### generateExampleSubdirectoryDocumentation

generateExampleSubdirectoryDocumentation generates markdown documentation for example subdirectories

```go
func (*Generator) generateExampleSubdirectoryDocumentation(sourceDir, outputDir string, context *templates.Context) error
```

**Parameters:**
- `sourceDir` (string)
- `outputDir` (string)
- `context` (*templates.Context)

**Returns:**
- error

### generateExamplesDocumentation

generateExamplesDocumentation generates examples documentation

```go
func (*Generator) generateExamplesDocumentation(packages []*discovery.PackageInfo, context *templates.Context) error
```

**Parameters:**
- `packages` ([]*discovery.PackageInfo)
- `context` (*templates.Context)

**Returns:**
- error

### generateGettingStartedDocumentation

generateGettingStartedDocumentation generates the getting-started documentation

```go
func (*Generator) generateGettingStartedDocumentation(context *templates.Context) error
```

**Parameters:**
- `context` (*templates.Context)

**Returns:**
- error

### generateMainFiles

generateMainFiles generates the main documentation files

```go
func (*Generator) generateMainFiles(context *templates.Context) error
```

**Parameters:**
- `context` (*templates.Context)

**Returns:**
- error

### isInitialization

isInitialization checks if this is the first time documentation is being generated by checking if getting-started, core-concepts, or advanced directories exist

```go
func (*Generator) isInitialization() bool
```

**Parameters:**
  None

**Returns:**
- bool

## External Links

- [Package Overview](../packages/generator.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/generator)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/generator)
