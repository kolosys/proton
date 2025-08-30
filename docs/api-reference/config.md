# config API

Complete API documentation for the config package.

**Import Path:** `github.com/kolosys/proton/internal/config`

## Package Documentation



## Types

### APIGeneration
_No documentation available_

#### Example Usage

```go
// Create a new APIGeneration
apigeneration := APIGeneration{
    Enabled: true,
    IncludeUnexported: true,
    IncludeTests: true,
    IncludeExamples: true,
}
```

#### Type Definition

```go
type APIGeneration struct {
    Enabled bool `yaml:"enabled" mapstructure:"enabled"`
    IncludeUnexported bool `yaml:"include_unexported" mapstructure:"include_unexported"`
    IncludeTests bool `yaml:"include_tests" mapstructure:"include_tests"`
    IncludeExamples bool `yaml:"include_examples" mapstructure:"include_examples"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Enabled | `bool` |  |
| IncludeUnexported | `bool` |  |
| IncludeTests | `bool` |  |
| IncludeExamples | `bool` |  |

### Config
Config represents the complete configuration for Proton

#### Example Usage

```go
// Create a new Config
config := Config{
    Repository: Repository{},
    Output: Output{},
    Discovery: Discovery{},
    Templates: Templates{},
    GitBook: GitBook{},
    Metadata: Metadata{},
    Generation: Generation{},
}
```

#### Type Definition

```go
type Config struct {
    Repository Repository `yaml:"repository" mapstructure:"repository"`
    Output Output `yaml:"output" mapstructure:"output"`
    Discovery Discovery `yaml:"discovery" mapstructure:"discovery"`
    Templates Templates `yaml:"templates" mapstructure:"templates"`
    GitBook GitBook `yaml:"gitbook" mapstructure:"gitbook"`
    Metadata Metadata `yaml:"metadata" mapstructure:"metadata"`
    Generation Generation `yaml:"generation" mapstructure:"generation"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Repository | `Repository` | Repository information |
| Output | `Output` |  |
| Discovery | `Discovery` |  |
| Templates | `Templates` |  |
| GitBook | `GitBook` |  |
| Metadata | `Metadata` |  |
| Generation | `Generation` |  |

### Constructor Functions

### Load

Load loads configuration from the specified path or discovers it automatically

```go
func Load(configPath, projectPath string) (*Config, error)
```

**Parameters:**

- `configPath` (string) - The path to the configuration file. If empty, the function will search for a config file in the following locations:

- `projectPath` (string) - The path to the project root directory.

**Returns:**

- *Config - The configuration object.

- error - An error if the configuration file is not found or the configuration is invalid.

## Methods

### Save

Save saves the configuration to a file

```go
func (*Config) Save(path string) error
```

**Parameters:**
- `path` (string)

**Returns:**
- error

### CustomGuide
_No documentation available_

#### Example Usage

```go
// Create a new CustomGuide
customguide := CustomGuide{
    Name: "example",
    File: "example",
    Title: "example",
}
```

#### Type Definition

```go
type CustomGuide struct {
    Name string `yaml:"name" mapstructure:"name"`
    File string `yaml:"file" mapstructure:"file"`
    Title string `yaml:"title" mapstructure:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| File | `string` |  |
| Title | `string` |  |

### CustomTemplate
_No documentation available_

#### Example Usage

```go
// Create a new CustomTemplate
customtemplate := CustomTemplate{
    Name: "example",
    File: "example",
}
```

#### Type Definition

```go
type CustomTemplate struct {
    Name string `yaml:"name" mapstructure:"name"`
    File string `yaml:"file" mapstructure:"file"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| File | `string` |  |

### Discovery
_No documentation available_

#### Example Usage

```go
// Create a new Discovery
discovery := Discovery{
    Packages: Packages{},
    APIGeneration: APIGeneration{},
    Examples: Examples{},
    Guides: Guides{},
}
```

#### Type Definition

```go
type Discovery struct {
    Packages Packages `yaml:"packages" mapstructure:"packages"`
    APIGeneration APIGeneration `yaml:"api_generation" mapstructure:"api_generation"`
    Examples Examples `yaml:"examples" mapstructure:"examples"`
    Guides Guides `yaml:"guides" mapstructure:"guides"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Packages | `Packages` |  |
| APIGeneration | `APIGeneration` |  |
| Examples | `Examples` |  |
| Guides | `Guides` |  |

### Examples
_No documentation available_

#### Example Usage

```go
// Create a new Examples
examples := Examples{
    Enabled: true,
    AutoDiscover: true,
    Directories: [],
}
```

#### Type Definition

```go
type Examples struct {
    Enabled bool `yaml:"enabled" mapstructure:"enabled"`
    AutoDiscover bool `yaml:"auto_discover" mapstructure:"auto_discover"`
    Directories []string `yaml:"directories" mapstructure:"directories"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Enabled | `bool` |  |
| AutoDiscover | `bool` |  |
| Directories | `[]string` |  |

### Generation
_No documentation available_

#### Example Usage

```go
// Create a new Generation
generation := Generation{
    DateFormat: "example",
    IncludeGeneratedNotice: true,
    IncludeTOC: true,
    MaxDepth: 42,
}
```

#### Type Definition

```go
type Generation struct {
    DateFormat string `yaml:"date_format" mapstructure:"date_format"`
    IncludeGeneratedNotice bool `yaml:"include_generated_notice" mapstructure:"include_generated_notice"`
    IncludeTOC bool `yaml:"include_toc" mapstructure:"include_toc"`
    MaxDepth int `yaml:"max_depth" mapstructure:"max_depth"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DateFormat | `string` |  |
| IncludeGeneratedNotice | `bool` |  |
| IncludeTOC | `bool` |  |
| MaxDepth | `int` |  |

### GitBook
_No documentation available_

#### Example Usage

```go
// Create a new GitBook
gitbook := GitBook{
    Title: "example",
    Description: "example",
    Theme: "example",
    Plugins: [],
    Structure: GitBookStructure{},
}
```

#### Type Definition

```go
type GitBook struct {
    Title string `yaml:"title" mapstructure:"title"`
    Description string `yaml:"description" mapstructure:"description"`
    Theme string `yaml:"theme" mapstructure:"theme"`
    Plugins []string `yaml:"plugins" mapstructure:"plugins"`
    Structure GitBookStructure `yaml:"structure" mapstructure:"structure"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Title | `string` |  |
| Description | `string` |  |
| Theme | `string` |  |
| Plugins | `[]string` |  |
| Structure | `GitBookStructure` |  |

### GitBookStructure
_No documentation available_

#### Example Usage

```go
// Create a new GitBookStructure
gitbookstructure := GitBookStructure{
    Readme: "example",
    Summary: "example",
}
```

#### Type Definition

```go
type GitBookStructure struct {
    Readme string `yaml:"readme" mapstructure:"readme"`
    Summary string `yaml:"summary" mapstructure:"summary"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Readme | `string` |  |
| Summary | `string` |  |

### Guides
_No documentation available_

#### Example Usage

```go
// Create a new Guides
guides := Guides{
    Enabled: true,
    IncludeContributing: true,
    IncludeFAQ: true,
    CustomGuides: [],
}
```

#### Type Definition

```go
type Guides struct {
    Enabled bool `yaml:"enabled" mapstructure:"enabled"`
    IncludeContributing bool `yaml:"include_contributing" mapstructure:"include_contributing"`
    IncludeFAQ bool `yaml:"include_faq" mapstructure:"include_faq"`
    CustomGuides []CustomGuide `yaml:"custom_guides" mapstructure:"custom_guides"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Enabled | `bool` |  |
| IncludeContributing | `bool` |  |
| IncludeFAQ | `bool` |  |
| CustomGuides | `[]CustomGuide` |  |

### Metadata
_No documentation available_

#### Example Usage

```go
// Create a new Metadata
metadata := Metadata{
    Version: "example",
    GoVersion: "example",
    Author: "example",
    License: "example",
}
```

#### Type Definition

```go
type Metadata struct {
    Version string `yaml:"version" mapstructure:"version"`
    GoVersion string `yaml:"go_version" mapstructure:"go_version"`
    Author string `yaml:"author" mapstructure:"author"`
    License string `yaml:"license" mapstructure:"license"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Version | `string` |  |
| GoVersion | `string` |  |
| Author | `string` |  |
| License | `string` |  |

### Output
_No documentation available_

#### Example Usage

```go
// Create a new Output
output := Output{
    Directory: "example",
    Clean: true,
    GitBookConfig: true,
}
```

#### Type Definition

```go
type Output struct {
    Directory string `yaml:"directory" mapstructure:"directory"`
    Clean bool `yaml:"clean" mapstructure:"clean"`
    GitBookConfig bool `yaml:"gitbook_config" mapstructure:"gitbook_config"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Directory | `string` |  |
| Clean | `bool` |  |
| GitBookConfig | `bool` |  |

### Package
_No documentation available_

#### Example Usage

```go
// Create a new Package
package := Package{
    Name: "example",
    Path: "example",
    Description: "example",
}
```

#### Type Definition

```go
type Package struct {
    Name string `yaml:"name" mapstructure:"name"`
    Path string `yaml:"path" mapstructure:"path"`
    Description string `yaml:"description" mapstructure:"description"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Path | `string` |  |
| Description | `string` |  |

### Packages
_No documentation available_

#### Example Usage

```go
// Create a new Packages
packages := Packages{
    AutoDiscover: true,
    IncludePatterns: [],
    ExcludePatterns: [],
    ManualPackages: [],
}
```

#### Type Definition

```go
type Packages struct {
    AutoDiscover bool `yaml:"auto_discover" mapstructure:"auto_discover"`
    IncludePatterns []string `yaml:"include_patterns" mapstructure:"include_patterns"`
    ExcludePatterns []string `yaml:"exclude_patterns" mapstructure:"exclude_patterns"`
    ManualPackages []Package `yaml:"manual_packages" mapstructure:"manual_packages"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AutoDiscover | `bool` |  |
| IncludePatterns | `[]string` |  |
| ExcludePatterns | `[]string` |  |
| ManualPackages | `[]Package` |  |

### Repository
_No documentation available_

#### Example Usage

```go
// Create a new Repository
repository := Repository{
    Name: "example",
    Owner: "example",
    Description: "example",
    ImportPath: "example",
    Branch: "example",
    URL: "example",
}
```

#### Type Definition

```go
type Repository struct {
    Name string `yaml:"name" mapstructure:"name"`
    Owner string `yaml:"owner" mapstructure:"owner"`
    Description string `yaml:"description" mapstructure:"description"`
    ImportPath string `yaml:"import_path" mapstructure:"import_path"`
    Branch string `yaml:"branch" mapstructure:"branch"`
    URL string `yaml:"url" mapstructure:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Owner | `string` |  |
| Description | `string` |  |
| ImportPath | `string` |  |
| Branch | `string` |  |
| URL | `string` |  |

### Templates
_No documentation available_

#### Example Usage

```go
// Create a new Templates
templates := Templates{
    Directory: "example",
    CustomTemplates: [],
}
```

#### Type Definition

```go
type Templates struct {
    Directory string `yaml:"directory" mapstructure:"directory"`
    CustomTemplates []CustomTemplate `yaml:"custom_templates" mapstructure:"custom_templates"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Directory | `string` |  |
| CustomTemplates | `[]CustomTemplate` |  |

## External Links

- [Package Overview](../packages/config.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/config)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/config)
