# {{.Package.Name}} API

Complete API documentation for the {{.Package.Name}} package.

**Import Path:** `{{.Package.ImportPath}}`

## Package Documentation

{{.Package.Doc.Doc}}

{{- if .Package.Constants}}

## Constants

{{- range .Package.Constants}}

### {{join .Names ", "}}

{{.Doc}}

```go
{{range .Decl.Specs}}{{.}}{{end}}
```

{{- end}}
{{- end}}

{{- if .Package.Variables}}

## Variables

{{- range .Package.Variables}}

### {{join .Names ", "}}

{{.Doc}}

```go
{{range .Decl.Specs}}{{.}}{{end}}
```

{{- end}}
{{- end}}

{{- if .Package.Types}}

## Types

{{- range .Package.Types}}

### {{.Name}}

{{- if .Doc}}
{{.Doc}}
{{- else}}
_No documentation available_
{{- end}}

{{- if .ExampleCode}}

#### Example Usage

```go
{{.ExampleCode}}
```

{{- end}}

#### Type Definition

```go
{{.Declaration}}
```

{{- if eq .TypeKind "interface"}}

## Methods

| Method | Description |
| ------ | ----------- |

{{- range .Methods}}
| `{{.Name}}` | {{.Doc}} |
{{- end}}

{{- end}}

{{- if hasFields .}}

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |

{{- range .Fields}}
| {{formatFieldName .}} | `{{.Type}}` | {{.Doc}} |
{{- end}}

{{- end}}

{{- if .Funcs}}

### Constructor Functions

{{- range .Funcs}}

### {{.Name}}

{{.Doc}}

```go
{{.Declaration}}
```

**Parameters:**

{{- if hasParams .}}
{{- range .Params}}
{{- if .Doc}}

- `{{.Name}}` ({{.Type}}) - {{.Doc}}
  {{- else}}
- `{{.Name}}` ({{.Type}})
  {{- end}}
  {{- end}}
  {{- else}}
  None
  {{- end}}

**Returns:**

{{- if hasResults .}}
{{- range .Results}}
{{- if .Doc}}

- {{.Type}} - {{.Doc}}
  {{- else}}
- {{.Type}}
  {{- end}}
  {{- end}}
  {{- else}}
  None
  {{- end}}

{{- end}}
{{- end}}

{{- if .Methods}}

## Methods

{{- range .Methods}}

### {{.Name}}

{{.Doc}}

```go
{{.Declaration}}
```

**Parameters:**

{{- if hasParams .}}
{{- range .Params}}
{{- if .Doc}}

- `{{.Name}}` ({{.Type}}) - {{.Doc}}
  {{- else}}
- `{{.Name}}` ({{.Type}})
  {{- end}}
  {{- end}}
  {{- else}}
  None
  {{- end}}

**Returns:**

{{- if hasResults .}}
{{- range .Results}}
{{- if .Doc}}

- {{.Type}} - {{.Doc}}
  {{- else}}
- {{.Type}}
  {{- end}}
  {{- end}}
  {{- else}}
  None
  {{- end}}

{{- end}}
{{- end}}

{{- end}}
{{- end}}

{{- if .Package.Functions}}

## Functions

{{- range .Package.Functions}}

### {{.Name}}

{{- if .Doc}}
{{.Doc}}
{{- else}}
_No documentation available_
{{- end}}

```go
{{.Declaration}}
```

**Parameters:**

{{- if hasParams .}}
| Parameter | Type | Description |
|-----------|------|-------------|
{{- range .Params}}
{{- if .Doc}}
| `{{.Name}}` | `{{.Type}}` | {{.Doc}} |
{{- else}}
| `{{.Name}}` | `{{.Type}}` | |
{{- end}}
{{- end}}
{{- else}}
None
{{- end}}

**Returns:**

{{- if hasResults .}}
| Type | Description |
|------|-------------|
{{- range .Results}}
{{- if .Doc}}
| `{{.Type}}` | {{.Doc}} |
{{- else}}
| `{{.Type}}` | |
{{- end}}
{{- end}}
{{- else}}
None
{{- end}}

**Example:**

```go
// Example usage of {{.Name}}
{{.ExampleCode}}
```

{{- end}}
{{- end}}

{{- if .Package.Examples}}

## Code Examples

{{- range .Package.Examples}}

### {{.Name}}

{{- if .Doc}}
{{.Doc}}
{{- end}}

```go
{{.Code}}
```

{{- if .Output}}

**Output:**

```
{{.Output}}
```

{{- end}}

{{- end}}
{{- end}}

## External Links

- [Package Overview](../packages/{{.Package.Name}}.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Package.ImportPath}})
- [Source Code]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
