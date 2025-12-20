# config Benchmarks

Performance benchmarks for the config package.

**Import Path:** `github.com/kolosys/proton/internal/config`

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |
| `ConfigLoad-8` | 175089 | 57840 | 870 |
| `ConfigSave-8` | 385361 | 81640 | 212 |

### ConfigLoad-8

- **Nanoseconds per operation:** 175089 ns/op
- **Bytes allocated per operation:** 57840 B/op
- **Allocations per operation:** 870 allocs/op
- **Number of runs:** 6436

### ConfigSave-8

- **Nanoseconds per operation:** 385361 ns/op
- **Bytes allocated per operation:** 81640 B/op
- **Allocations per operation:** 212 allocs/op
- **Number of runs:** 4047

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
