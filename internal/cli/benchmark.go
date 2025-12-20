package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	benchmarkOutput string
	benchmarkPath   string
)

// benchmarkCmd represents the benchmark command
var benchmarkCmd = &cobra.Command{
	Use:   "benchmark [package-path]",
	Short: "Run benchmarks and generate benchmark documentation",
	Long: `Run Go benchmarks for the project and generate benchmark documentation.

This command will:
- Run 'go test -bench=. -benchmem ./...' for the project
- Parse benchmark results
- Generate benchmark documentation files

Examples:
  proton benchmark                    # Run benchmarks for current directory
  proton benchmark ./my-package      # Run benchmarks for specific package
  proton benchmark --output results.txt  # Save raw output to file`,
	Args: cobra.MaximumNArgs(1),
	RunE: runBenchmark,
}

func runBenchmark(cmd *cobra.Command, args []string) error {
	// Determine package path
	if len(args) > 0 {
		benchmarkPath = args[0]
	} else {
		benchmarkPath = "./..."
	}

	// Build command
	benchCmd := exec.Command("go", "test", "-bench=.", "-benchmem", benchmarkPath)

	// Run benchmark
	output, err := benchCmd.CombinedOutput()
	if err != nil {
		// Check if it's just because there are no benchmarks
		if strings.Contains(string(output), "no test files") || strings.Contains(string(output), "no packages") {
			fmt.Println("No benchmarks found in the specified path")
			return nil
		}
		return fmt.Errorf("failed to run benchmarks: %w\nOutput: %s", err, string(output))
	}

	// Write output if specified
	if benchmarkOutput != "" {
		absPath, err := filepath.Abs(benchmarkOutput)
		if err != nil {
			return fmt.Errorf("invalid output path: %w", err)
		}
		if err := os.WriteFile(absPath, output, 0644); err != nil {
			return fmt.Errorf("failed to write benchmark output: %w", err)
		}
		fmt.Printf("Benchmark results written to %s\n", absPath)
	} else {
		// Print output to stdout
		fmt.Print(string(output))
	}

	return nil
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)

	// Local flags
	benchmarkCmd.Flags().StringVarP(&benchmarkOutput, "output", "o", "", "output file for benchmark results (default: stdout)")
}
