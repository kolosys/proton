# {{.Package.Name}} Benchmarks

Performance benchmarks for the {{.Package.Name}} package.

**Import Path:** `{{.Package.ImportPath}}`

{{- if .Benchmarks}}

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |

{{- range .Benchmarks}}
| `{{.Name}}` | {{.NsPerOp}} | {{.BytesPerOp}} | {{.AllocsPerOp}} |
{{- end}}

{{- range .Benchmarks}}

### {{.Name}}

- **Nanoseconds per operation:** {{.NsPerOp}} ns/op
- **Bytes allocated per operation:** {{.BytesPerOp}} B/op
- **Allocations per operation:** {{.AllocsPerOp}} allocs/op
- **Number of runs:** {{.Runs}}

{{- end}}

{{- else}}

## No Benchmarks Available

No benchmark results are available for this package. To add benchmarks:

1. Create a `*_test.go` file in the package directory
2. Add benchmark functions following the pattern:
   ```go
   func BenchmarkFunctionName(b *testing.B) {
       for i := 0; i < b.N; i++ {
           // Your code here
       }
   }
   ```
3. Run `proton benchmark` to generate benchmark results

{{- end}}

## Running Benchmarks

To run benchmarks for this package:

```bash
go test -bench=. -benchmem ./{{packagePath .Package}}
```

To run benchmarks for all packages:

```bash
go test -bench=. -benchmem ./...
```

## External Links

- [Package Overview](../core-concepts/{{.Package.Name}}.md)
- [API Reference](../api-reference/{{.Package.Name}}.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/{{.Package.ImportPath}})
- [Source Code]({{.Repository.URL}}/tree/{{.Repository.Branch}}/{{packagePath .Package}})
