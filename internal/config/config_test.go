package config

import (
	"testing"
)

// BenchmarkConfigLoad benchmarks configuration loading
func BenchmarkConfigLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Load("", ".")
	}
}

// BenchmarkConfigSave benchmarks configuration saving
func BenchmarkConfigSave(b *testing.B) {
	cfg := &Config{
		Repository: Repository{
			Name: "test",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.Save("/tmp/test-config.yml")
	}
}
