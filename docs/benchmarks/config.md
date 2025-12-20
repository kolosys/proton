# config Benchmarks

Performance benchmarks for the config package.

**Import Path:** `github.com/kolosys/proton/internal/config`

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |
| `ConfigLoad-4` | 222938 | 57608 | 870 |
| `ConfigSave-4` | 220066 | 81641 | 212 |

### ConfigLoad-4

- **Nanoseconds per operation:** 222938 ns/op
- **Bytes allocated per operation:** 57608 B/op
- **Allocations per operation:** 870 allocs/op
- **Number of runs:** 5067

### ConfigSave-4

- **Nanoseconds per operation:** 220066 ns/op
- **Bytes allocated per operation:** 81641 B/op
- **Allocations per operation:** 212 allocs/op
- **Number of runs:** 6368

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
