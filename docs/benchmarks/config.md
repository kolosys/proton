# config Benchmarks

Performance benchmarks for the config package.

**Import Path:** `github.com/kolosys/proton/internal/config`

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |
| `ConfigLoad-4` | 228871 | 59769 | 899 |
| `ConfigSave-4` | 178427 | 83249 | 223 |

### ConfigLoad-4

- **Nanoseconds per operation:** 228871 ns/op
- **Bytes allocated per operation:** 59769 B/op
- **Allocations per operation:** 899 allocs/op
- **Number of runs:** 5107

### ConfigSave-4

- **Nanoseconds per operation:** 178427 ns/op
- **Bytes allocated per operation:** 83249 B/op
- **Allocations per operation:** 223 allocs/op
- **Number of runs:** 6836

## Running Benchmarks

To run benchmarks for this package:

```bash
go test -bench=. -benchmem ./internal/config
```

To run benchmarks for all packages:

```bash
go test -bench=. -benchmem ./...
```

## External Links

- [Package Overview](../core-concepts/config.md)
- [API Reference](../api-reference/config.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/config)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/config)
