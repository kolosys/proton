package benchmarks

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// BenchmarkResult represents a single benchmark result
type BenchmarkResult struct {
	Name        string // Benchmark name (e.g., BenchmarkFunction-8)
	NsPerOp     int64  // Nanoseconds per operation
	AllocsPerOp int64  // Allocations per operation
	BytesPerOp  int64  // Bytes allocated per operation
	Runs        int    // Number of runs
}

// PackageBenchmarks contains benchmark results for a package
type PackageBenchmarks struct {
	PackagePath string
	Results     []*BenchmarkResult
}

// Runner handles benchmark execution and parsing
type Runner struct {
	projectPath string
}

// New creates a new benchmark runner
func New(projectPath string) *Runner {
	return &Runner{
		projectPath: projectPath,
	}
}

// Run executes benchmarks for the specified path pattern
func (r *Runner) Run(pattern string) ([]*PackageBenchmarks, error) {
	// Build command
	benchCmd := exec.Command("go", "test", "-bench=.", "-benchmem", pattern)
	benchCmd.Dir = r.projectPath

	// Run benchmark
	output, err := benchCmd.CombinedOutput()
	if err != nil {
		// Check if it's just because there are no benchmarks
		if strings.Contains(string(output), "no test files") || strings.Contains(string(output), "no packages") {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to run benchmarks: %w\nOutput: %s", err, string(output))
	}

	// Parse output
	return r.parseOutput(string(output))
}

// RunFromFile parses benchmark results from a file
func (r *Runner) RunFromFile(filePath string) ([]*PackageBenchmarks, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read benchmark file: %w", err)
	}

	return r.parseOutput(string(content))
}

// parseOutput parses the output from 'go test -bench=. -benchmem'
func (r *Runner) parseOutput(output string) ([]*PackageBenchmarks, error) {
	var results []*PackageBenchmarks
	var currentPackage *PackageBenchmarks

	scanner := bufio.NewScanner(strings.NewReader(output))

	// Regex patterns for parsing benchmark output
	// Format: BenchmarkName-8         1000000    1234 ns/op    567 B/op    8 allocs/op
	benchmarkPattern := regexp.MustCompile(`^Benchmark(\S+)\s+(\d+)\s+(\d+\.?\d*)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)
	// Format: pkg: github.com/user/repo/pkg
	pkgPattern := regexp.MustCompile(`^pkg:\s+(\S+)`)
	// Format: ok      github.com/user/repo/pkg    0.123s
	packagePattern := regexp.MustCompile(`^ok\s+(\S+)\s+`)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Check if this is a pkg: line (appears before benchmarks)
		if matches := pkgPattern.FindStringSubmatch(line); matches != nil {
			// Save previous package if it exists
			if currentPackage != nil && len(currentPackage.Results) > 0 {
				results = append(results, currentPackage)
			}
			// Start new package
			currentPackage = &PackageBenchmarks{
				PackagePath: matches[1],
				Results:     []*BenchmarkResult{},
			}
			continue
		}

		// Check if this is a package line (appears after benchmarks)
		if matches := packagePattern.FindStringSubmatch(line); matches != nil {
			// Update package path if we don't have one yet
			if currentPackage == nil {
				currentPackage = &PackageBenchmarks{
					PackagePath: matches[1],
					Results:     []*BenchmarkResult{},
				}
			} else if currentPackage.PackagePath == "unknown" {
				currentPackage.PackagePath = matches[1]
			}
			// Save package if it has results
			if len(currentPackage.Results) > 0 {
				results = append(results, currentPackage)
				currentPackage = nil
			}
			continue
		}

		// Check if this is a benchmark result line
		if matches := benchmarkPattern.FindStringSubmatch(line); matches != nil {
			if currentPackage == nil {
				// If no package context, create one
				currentPackage = &PackageBenchmarks{
					PackagePath: "unknown",
					Results:     []*BenchmarkResult{},
				}
			}

			runs, _ := strconv.Atoi(matches[2])
			nsPerOp, _ := strconv.ParseFloat(matches[3], 64)
			bytesPerOp, _ := strconv.Atoi(matches[4])
			allocsPerOp, _ := strconv.Atoi(matches[5])

			result := &BenchmarkResult{
				Name:        matches[1],
				NsPerOp:     int64(nsPerOp),
				BytesPerOp:  int64(bytesPerOp),
				AllocsPerOp: int64(allocsPerOp),
				Runs:        runs,
			}

			currentPackage.Results = append(currentPackage.Results, result)
		}
	}

	// Save last package if it exists
	if currentPackage != nil && len(currentPackage.Results) > 0 {
		results = append(results, currentPackage)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse benchmark output: %w", err)
	}

	return results, nil
}

// GetBenchmarksForPackage returns benchmark results for a specific package path
func (r *Runner) GetBenchmarksForPackage(pkgPath string, allBenchmarks []*PackageBenchmarks) []*BenchmarkResult {
	for _, pkgBench := range allBenchmarks {
		// Try to match package path
		if pkgBench.PackagePath == pkgPath {
			return pkgBench.Results
		}
		// Also try matching by base name
		if filepath.Base(pkgBench.PackagePath) == filepath.Base(pkgPath) {
			return pkgBench.Results
		}
	}
	return nil
}
