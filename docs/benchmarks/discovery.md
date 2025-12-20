# discovery Benchmarks

Performance benchmarks for the discovery package.

**Import Path:** `github.com/kolosys/proton/internal/discovery`

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

## Running Benchmarks

To run benchmarks for this package:

```bash
go test -bench=. -benchmem ./internal/discovery
```

To run benchmarks for all packages:

```bash
go test -bench=. -benchmem ./...
```

## External Links

- [Package Overview](../core-concepts/discovery.md)
- [API Reference](../api-reference/discovery.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/discovery)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/discovery)
