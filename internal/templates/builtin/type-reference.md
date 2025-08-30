# {{.Type.Name}} Type Reference

{{.Type.Doc}}

## Definition

```go
{{.Type.Decl}}
```

{{- if hasFields .}}

## Fields

{{- range .Fields}}

### {{formatFieldName .}}

{{.Doc}}

| Field                 | Type                 | Description |
| --------------------- | -------------------- | ----------- |
| {{formatFieldName .}} | `{{typeLink .Type}}` | {{.Doc}}    |

{{- if .Tag}}
**Tags:** {{formatTag .Tag}}
{{- end}}

{{- end}}
{{- end}}

{{- if .Type.Interfaces}}

## Implements

{{- range .Type.Interfaces}}

- [`{{.Name}}`]({{.Name}}.md)
  {{- end}}
  {{- end}}

{{- if .Methods}}

## Methods

{{- range .Methods}}

### {{.Name}}

{{.Doc}}

```go
{{.Decl}}
```

**Parameters:**

{{- if hasParams .}}
| Parameter | Type | Description |
|-----------|------|-------------|
{{- range .Params}}
| `{{.Name}}` | `{{typeLink .Type}}` | {{.Doc}} |
{{- end}}
{{- else}}
None
{{- end}}

**Returns:**

{{- if hasResults .}}
| Type | Description |
|------|-------------|
{{- range .Results}}
| `{{typeLink .Type}}` | {{.Doc}} |
{{- end}}
{{- else}}
None
{{- end}}

**Example:**

```go
instance := {{$.Type.Name}}{}
result := instance.{{.Name}}({{range $i, $p := .Params}}{{if $i}}, {{end}}{{$p.Name}}{{end}})
```

{{- end}}
{{- end}}

{{- if .Funcs}}

## Constructor Functions

{{- range .Funcs}}

### {{.Name}}

{{.Doc}}

```go
{{.Decl}}
```

**Parameters:**

{{- if hasParams .}}
| Parameter | Type | Description |
|-----------|------|-------------|
{{- range .Params}}
| `{{.Name}}` | `{{typeLink .Type}}` | {{.Doc}} |
{{- end}}
{{- else}}
None
{{- end}}

**Returns:**

{{- if hasResults .}}
| Type | Description |
|------|-------------|
{{- range .Results}}
| `{{typeLink .Type}}` | {{.Doc}} |
{{- end}}
{{- else}}

- `{{$.Type.Name}}` - New instance of {{$.Type.Name}}
  {{- end}}

**Example:**

```go
instance := {{.Name}}({{range $i, $p := .Params}}{{if $i}}, {{end}}/* {{$p.Name}} */{{end}})
```

{{- end}}
{{- end}}

## Usage Examples

```go
package main

import (
    "{{$.Repository.ImportPath}}"
)

func main() {
    // Create new instance
    {{- if .Funcs}}
    {{- $constructor := index .Funcs 0}}
    instance := {{$constructor.Name}}({{range $i, $p := $constructor.Params}}{{if $i}}, {{end}}/* {{$p.Name}} */{{end}})
    {{- else}}
    instance := {{.Type.Name}}{}
    {{- end}}

    {{- if .Methods}}
    {{- $method := index .Methods 0}}

    // Call method
    result := instance.{{$method.Name}}({{range $i, $p := $method.Params}}{{if $i}}, {{end}}/* {{$p.Name}} */{{end}})
    {{- end}}
}
```

## External Links

- [Package Overview](../packages/{{$.Package.Name}}.md)
- [Full API Reference]({{$.Package.Name}}.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/{{$.Package.ImportPath}}#{{.Type.Name}})
- [Source Code]({{$.Repository.URL}}/tree/{{$.Repository.Branch}}/{{packagePath $.Package}})
