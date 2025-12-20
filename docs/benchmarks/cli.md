# cli Benchmarks

Performance benchmarks for the cli package.

**Import Path:** `github.com/kolosys/proton/internal/cli`

## Benchmark Results

The following benchmarks were run using `go test -bench=. -benchmem`:

| Benchmark | ns/op | B/op | allocs/op |
| --------- | ----- | ---- | --------- |
| `CLIExecute-8` | 0 | 0 | 0 |
| `ConfigLoad-8` | 0 | 0 | 0 |

### CLIExecute-8

- **Nanoseconds per operation:** 0 ns/op
- **Bytes allocated per operation:** 0 B/op
- **Allocations per operation:** 0 allocs/op
- **Number of runs:** 1000000000

### ConfigLoad-8

- **Nanoseconds per operation:** 0 ns/op
- **Bytes allocated per operation:** 0 B/op
- **Allocations per operation:** 0 allocs/op
- **Number of runs:** 1000000000

## Running Benchmarks

To run benchmarks for this package:

```bash
go test -bench=. -benchmem ./internal/cli
```

To run benchmarks for all packages:

```bash
go test -bench=. -benchmem ./...
```

## External Links

- [Package Overview](../core-concepts/cli.md)
- [API Reference](../api-reference/cli.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/cli)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/cli)
