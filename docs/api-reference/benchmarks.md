# benchmarks API

Complete API documentation for the benchmarks package.

**Import Path:** `github.com/kolosys/proton/internal/benchmarks`

## Package Documentation



## Types

### BenchmarkResult
BenchmarkResult represents a single benchmark result

#### Example Usage

```go
// Create a new BenchmarkResult
benchmarkresult := BenchmarkResult{
    Name: "example",
    NsPerOp: 42,
    AllocsPerOp: 42,
    BytesPerOp: 42,
    Runs: 42,
}
```

#### Type Definition

```go
type BenchmarkResult struct {
    Name string
    NsPerOp int64
    AllocsPerOp int64
    BytesPerOp int64
    Runs int
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` | Benchmark name (e.g., BenchmarkFunction-8) |
| NsPerOp | `int64` | Nanoseconds per operation |
| AllocsPerOp | `int64` | Allocations per operation |
| BytesPerOp | `int64` | Bytes allocated per operation |
| Runs | `int` | Number of runs |

### PackageBenchmarks
PackageBenchmarks contains benchmark results for a package

#### Example Usage

```go
// Create a new PackageBenchmarks
packagebenchmarks := PackageBenchmarks{
    PackagePath: "example",
    Results: [],
}
```

#### Type Definition

```go
type PackageBenchmarks struct {
    PackagePath string
    Results []*BenchmarkResult
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| PackagePath | `string` |  |
| Results | `[]*BenchmarkResult` |  |

### Runner
Runner handles benchmark execution and parsing

#### Example Usage

```go
// Create a new Runner
runner := Runner{

}
```

#### Type Definition

```go
type Runner struct {
}
```

### Constructor Functions

### New

New creates a new benchmark runner

```go
func New(projectPath string) *Runner
```

**Parameters:**
- `projectPath` (string)

**Returns:**
- *Runner

## Methods

### GetBenchmarksForPackage

GetBenchmarksForPackage returns benchmark results for a specific package path

```go
func (*Runner) GetBenchmarksForPackage(pkgPath string, allBenchmarks []*PackageBenchmarks) []*BenchmarkResult
```

**Parameters:**
- `pkgPath` (string)
- `allBenchmarks` ([]*PackageBenchmarks)

**Returns:**
- []*BenchmarkResult

### Run

Run executes benchmarks for the specified path pattern

```go
func (*Runner) Run(pattern string) ([]*PackageBenchmarks, error)
```

**Parameters:**
- `pattern` (string)

**Returns:**
- []*PackageBenchmarks
- error

### RunFromFile

RunFromFile parses benchmark results from a file

```go
func (*Runner) RunFromFile(filePath string) ([]*PackageBenchmarks, error)
```

**Parameters:**
- `filePath` (string)

**Returns:**
- []*PackageBenchmarks
- error

## External Links

- [Package Overview](../packages/benchmarks.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/proton/internal/benchmarks)
- [Source Code](https://github.com/kolosys/proton/tree/main/internal/benchmarks)
