// Package interfaces provides core interfaces for Proton documentation generation.
package interfaces

import (
	"io"
	"time"
)

// Generator defines the interface for documentation generators
type Generator interface {
	// Generate creates documentation from the provided configuration
	// config: The configuration object containing generation settings
	// Returns: error if generation fails
	Generate(config interface{}) error

	// Validate checks if the configuration is valid
	// Returns: error if validation fails
	Validate() error

	// Clean removes generated files
	// Returns: error if cleanup fails
	Clean() error
}

// Parser defines the interface for parsing Go source code
type Parser interface {
	// Parse extracts documentation from Go source files
	Parse(files []string) (interface{}, error)

	// ExtractTypes returns type information from parsed source
	ExtractTypes() []interface{}
}

// Template represents a documentation template
type Template interface {
	// Render generates documentation content
	Render(data interface{}) (string, error)

	// Name returns the template name
	Name() string
}

// Codec defines how to serialize/deserialize values
type Codec[T any] interface {
	// Encode serializes a value to the writer
	Encode(value T) ([]byte, error)

	// Decode deserializes data from the reader
	Decode(data []byte) (T, error)
}

// Capsule represents a time-locked value
type Capsule[T any] struct {
	Value      T         `json:"value"`
	UnlockTime time.Time `json:"unlock_time"`
	CreatedAt  time.Time `json:"created_at"`
}

// Writer provides file writing capabilities
type Writer interface {
	io.Writer

	// WriteFile writes content to a file
	WriteFile(filename string, content []byte) error

	// Sync ensures all data is written
	Sync() error
}

// Configuration holds system configuration
type Configuration struct {
	// Debug enables debug mode for verbose logging
	Debug bool `yaml:"debug" json:"debug"`

	// OutputFormat specifies the output format (markdown, html, etc.)
	OutputFormat string `yaml:"output_format" json:"output_format"`

	// MaxConcurrency limits concurrent operations for performance tuning
	MaxConcurrency int `yaml:"max_concurrency" json:"max_concurrency"`
}
