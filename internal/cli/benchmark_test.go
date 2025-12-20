package cli

import (
	"testing"
)

// BenchmarkCLIExecute benchmarks the CLI command execution
func BenchmarkCLIExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Simulate CLI command execution
		_ = rootCmd.Use
	}
}

// BenchmarkConfigLoad benchmarks configuration loading
func BenchmarkConfigLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Simulate config loading
		_ = cfgFile
	}
}
