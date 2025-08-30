# interfaces API

Complete API documentation for the interfaces package.

**Import Path:** `github.com/kolosys/proton/internal/interfaces`

## Package Documentation

Package interfaces provides core interfaces for Proton documentation generation.


## Types

### Capsule
Capsule represents a time-locked value

#### Example Usage

```go
// Create a new Capsule
capsule := Capsule{
    Value: T{},
    UnlockTime: /* value */,
    CreatedAt: /* value */,
}
```

#### Type Definition

```go
type Capsule struct {
    Value T `json:"value"`
    UnlockTime time.Time `json:"unlock_time"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Value | `T` |  |
| UnlockTime | `time.Time` |  |
| CreatedAt | `time.Time` |  |

### Codec
Codec defines how to serialize/deserialize values

#### Example Usage

```go
// Example implementation of Codec
type MyCodec struct {
    // Add your fields here
}

func (m MyCodec) Encode(param1 T) []byte {
    // Implement your logic here
    return
}

func (m MyCodec) Decode(param1 []byte) T {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Codec interface {
    Encode(value T) ([]byte, error)
    Decode(data []byte) (T, error)
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Configuration
Configuration holds system configuration

#### Example Usage

```go
// Create a new Configuration
configuration := Configuration{
    Debug: true,
    OutputFormat: "example",
    MaxConcurrency: 42,
}
```

#### Type Definition

```go
type Configuration struct {
    Debug bool `yaml:"debug" json:"debug"`
    OutputFormat string `yaml:"output_format" json:"output_format"`
    MaxConcurrency int `yaml:"max_concurrency" json:"max_concurrency"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Debug | `bool` | Debug enables debug mode for verbose logging |
| OutputFormat | `string` | OutputFormat specifies the output format (markdown, html, etc.) |
| MaxConcurrency | `int` | MaxConcurrency limits concurrent operations for performance tuning |

### Generator
Generator defines the interface for documentation generators

#### Example Usage

```go
// Example implementation of Generator
type MyGenerator struct {
    // Add your fields here
}

func (m MyGenerator) Generate(param1 interface{}) error {
    // Implement your logic here
    return
}

func (m MyGenerator) Validate() error {
    // Implement your logic here
    return
}

func (m MyGenerator) Clean() error {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Generator interface {
    Generate(config interface{}) error
    Validate() error
    Clean() error
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Parser
Parser defines the interface for parsing Go source code

#### Example Usage

```go
// Example implementation of Parser
type MyParser struct {
    // Add your fields here
}

func (m MyParser) Parse(param1 []string) interface{} {
    // Implement your logic here
    return
}

func (m MyParser) ExtractTypes() []interface{} {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Parser interface {
    Parse(files []string) (interface{}, error)
    ExtractTypes() []interface{}
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Template
Template represents a documentation template

#### Example Usage

```go
// Example implementation of Template
type MyTemplate struct {
    // Add your fields here
}

func (m MyTemplate) Render(param1 interface{}) string {
    // Implement your logic here
    return
}

func (m MyTemplate) Name() string {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Template interface {
    Render(data interface{}) (string, error)
    Name() string
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Writer
Writer provides file writing capabilities

#### Example Usage

```go
// Example implementation of Writer
type MyWriter struct {
    // Add your fields here
}

func (m MyWriter) WriteFile(param1 string, param2 []byte) error {
    // Implement your logic here
    return
}

func (m MyWriter) Sync() error {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Writer interface {
    io.Writer
    WriteFile(filename string, content []byte) error
    Sync() error
}
```

## Methods

| Method | Description |
| ------ | ----------- |

## External Links

- [Package Overview](../packages/interfaces.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/interfaces)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/interfaces)
