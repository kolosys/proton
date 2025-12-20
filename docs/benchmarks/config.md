# config Benchmarks

Performance benchmarks for the config package.

**Import Path:** `github.com/kolosys/proton/internal/config`

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |
| `ConfigLoad-4` | 255290 | 58167 | 884 |
| `ConfigSave-4` | 269458 | 83169 | 218 |

### ConfigLoad-4

- **Nanoseconds per operation:** 255290 ns/op
- **Bytes allocated per operation:** 58167 B/op
- **Allocations per operation:** 884 allocs/op
- **Number of runs:** 5073

### ConfigSave-4

- **Nanoseconds per operation:** 269458 ns/op
- **Bytes allocated per operation:** 83169 B/op
- **Allocations per operation:** 218 allocs/op
- **Number of runs:** 3850

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
