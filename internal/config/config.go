package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Config represents the complete configuration for Proton
type Config struct {
	// Repository information
	Repository Repository `yaml:"repository" mapstructure:"repository"`
	Output     Output     `yaml:"output" mapstructure:"output"`
	Discovery  Discovery  `yaml:"discovery" mapstructure:"discovery"`
	Templates  Templates  `yaml:"templates" mapstructure:"templates"`
	GitBook    GitBook    `yaml:"gitbook" mapstructure:"gitbook"`
	Metadata   Metadata   `yaml:"metadata" mapstructure:"metadata"`
	Generation Generation `yaml:"generation" mapstructure:"generation"`
}

type Repository struct {
	Name        string `yaml:"name" mapstructure:"name"`
	Owner       string `yaml:"owner" mapstructure:"owner"`
	Description string `yaml:"description" mapstructure:"description"`
	ImportPath  string `yaml:"import_path" mapstructure:"import_path"`
	Branch      string `yaml:"branch" mapstructure:"branch"`
	URL         string `yaml:"url" mapstructure:"url"`
}

type Output struct {
	Directory     string `yaml:"directory" mapstructure:"directory"`
	Clean         bool   `yaml:"clean" mapstructure:"clean"`
	GitBookConfig bool   `yaml:"gitbook_config" mapstructure:"gitbook_config"`
}

type Discovery struct {
	Packages      Packages      `yaml:"packages" mapstructure:"packages"`
	APIGeneration APIGeneration `yaml:"api_generation" mapstructure:"api_generation"`
	Examples      Examples      `yaml:"examples" mapstructure:"examples"`
	Guides        Guides        `yaml:"guides" mapstructure:"guides"`
}

type Packages struct {
	AutoDiscover    bool      `yaml:"auto_discover" mapstructure:"auto_discover"`
	IncludePatterns []string  `yaml:"include_patterns" mapstructure:"include_patterns"`
	ExcludePatterns []string  `yaml:"exclude_patterns" mapstructure:"exclude_patterns"`
	ManualPackages  []Package `yaml:"manual_packages" mapstructure:"manual_packages"`
}

type Package struct {
	Name        string `yaml:"name" mapstructure:"name"`
	Path        string `yaml:"path" mapstructure:"path"`
	Description string `yaml:"description" mapstructure:"description"`
}

type APIGeneration struct {
	Enabled           bool `yaml:"enabled" mapstructure:"enabled"`
	IncludeUnexported bool `yaml:"include_unexported" mapstructure:"include_unexported"`
	IncludeTests      bool `yaml:"include_tests" mapstructure:"include_tests"`
	IncludeExamples   bool `yaml:"include_examples" mapstructure:"include_examples"`
}

type Examples struct {
	Enabled      bool     `yaml:"enabled" mapstructure:"enabled"`
	AutoDiscover bool     `yaml:"auto_discover" mapstructure:"auto_discover"`
	Directories  []string `yaml:"directories" mapstructure:"directories"`
}

type Guides struct {
	Enabled             bool          `yaml:"enabled" mapstructure:"enabled"`
	IncludeContributing bool          `yaml:"include_contributing" mapstructure:"include_contributing"`
	IncludeFAQ          bool          `yaml:"include_faq" mapstructure:"include_faq"`
	CustomGuides        []CustomGuide `yaml:"custom_guides" mapstructure:"custom_guides"`
}

type CustomGuide struct {
	Name  string `yaml:"name" mapstructure:"name"`
	File  string `yaml:"file" mapstructure:"file"`
	Title string `yaml:"title" mapstructure:"title"`
}

type Templates struct {
	Directory       string           `yaml:"directory" mapstructure:"directory"`
	CustomTemplates []CustomTemplate `yaml:"custom_templates" mapstructure:"custom_templates"`
}

type CustomTemplate struct {
	Name string `yaml:"name" mapstructure:"name"`
	File string `yaml:"file" mapstructure:"file"`
}

type GitBook struct {
	Title       string           `yaml:"title" mapstructure:"title"`
	Description string           `yaml:"description" mapstructure:"description"`
	Theme       string           `yaml:"theme" mapstructure:"theme"`
	Plugins     []string         `yaml:"plugins" mapstructure:"plugins"`
	Structure   GitBookStructure `yaml:"structure" mapstructure:"structure"`
}

type GitBookStructure struct {
	Readme  string `yaml:"readme" mapstructure:"readme"`
	Summary string `yaml:"summary" mapstructure:"summary"`
}

type Metadata struct {
	Version   string `yaml:"version" mapstructure:"version"`
	GoVersion string `yaml:"go_version" mapstructure:"go_version"`
	Author    string `yaml:"author" mapstructure:"author"`
	License   string `yaml:"license" mapstructure:"license"`
}

type Generation struct {
	DateFormat             string `yaml:"date_format" mapstructure:"date_format"`
	IncludeGeneratedNotice bool   `yaml:"include_generated_notice" mapstructure:"include_generated_notice"`
	IncludeTOC             bool   `yaml:"include_toc" mapstructure:"include_toc"`
	MaxDepth               int    `yaml:"max_depth" mapstructure:"max_depth"`
}

// Load loads configuration from the specified path or discovers it automatically
// Parameters:
// - configPath: The path to the configuration file. If empty, the function will search for a config file in the following locations:
//   - .proton/config.yml
//   - .config/config.yml
//   - The project root directory
//
// - projectPath: The path to the project root directory.
//
// Returns:
// - *Config: The configuration object.
// - error: An error if the configuration file is not found or the configuration is invalid.
func Load(configPath, projectPath string) (*Config, error) {
	v := viper.New()

	// Set defaults
	setDefaults(v)

	// Configure viper
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// Search for config in common locations relative to project
		v.AddConfigPath(filepath.Join(projectPath, ".proton"))
		v.AddConfigPath(filepath.Join(projectPath, ".config"))
		v.AddConfigPath(projectPath)
		v.SetConfigType("yaml")
		v.SetConfigName("config")
	}

	// Read config file if it exists
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found, use defaults
	}

	// Unmarshal into config struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Auto-detect repository information if not provided
	if err := autoDetectRepo(&cfg, projectPath); err != nil {
		return nil, fmt.Errorf("error auto-detecting repository: %w", err)
	}

	// Validate and set remaining defaults
	if err := validateAndSetDefaults(&cfg); err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return &cfg, nil
}

// setDefaults sets default values for configuration
func setDefaults(v *viper.Viper) {
	// Repository defaults
	v.SetDefault("repository.branch", "main")

	// Output defaults
	v.SetDefault("output.directory", "docs")
	v.SetDefault("output.clean", true)
	v.SetDefault("output.gitbook_config", true)

	// Discovery defaults
	v.SetDefault("discovery.packages.auto_discover", true)
	v.SetDefault("discovery.packages.include_patterns", []string{"./..."})
	v.SetDefault("discovery.packages.exclude_patterns", []string{"./vendor/...", "./test/..."})

	v.SetDefault("discovery.api_generation.enabled", true)
	v.SetDefault("discovery.api_generation.include_unexported", false)
	v.SetDefault("discovery.api_generation.include_tests", false)
	v.SetDefault("discovery.api_generation.include_examples", true)

	v.SetDefault("discovery.examples.enabled", true)
	v.SetDefault("discovery.examples.auto_discover", true)

	v.SetDefault("discovery.guides.enabled", true)
	v.SetDefault("discovery.guides.include_contributing", true)
	v.SetDefault("discovery.guides.include_faq", true)

	// GitBook defaults
	v.SetDefault("gitbook.theme", "default")
	v.SetDefault("gitbook.structure.readme", "README.md")
	v.SetDefault("gitbook.structure.summary", "SUMMARY.md")

	// Metadata defaults
	v.SetDefault("metadata.version", "latest")
	v.SetDefault("metadata.license", "MIT")

	// Generation defaults
	v.SetDefault("generation.date_format", "2006-01-02")
	v.SetDefault("generation.include_generated_notice", true)
	v.SetDefault("generation.include_toc", true)
	v.SetDefault("generation.max_depth", 3)
}

// autoDetectRepo attempts to auto-detect repository information
func autoDetectRepo(cfg *Config, projectPath string) error {
	// Try to read go.mod for module information
	goModPath := filepath.Join(projectPath, "go.mod")
	if data, err := os.ReadFile(goModPath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "module ") {
				modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module"))
				if cfg.Repository.ImportPath == "" {
					cfg.Repository.ImportPath = modulePath
				}

				// Extract repository info from module path
				if strings.Contains(modulePath, "github.com/") {
					parts := strings.Split(modulePath, "/")
					if len(parts) >= 3 {
						if cfg.Repository.Owner == "" {
							cfg.Repository.Owner = parts[1]
						}
						if cfg.Repository.Name == "" {
							cfg.Repository.Name = parts[2]
						}
					}
				}
				break
			}
		}
	}

	// Set repository URL if not provided
	if cfg.Repository.URL == "" && cfg.Repository.Owner != "" && cfg.Repository.Name != "" {
		cfg.Repository.URL = fmt.Sprintf("https://github.com/%s/%s", cfg.Repository.Owner, cfg.Repository.Name)
	}

	return nil
}

// validateAndSetDefaults validates the configuration and sets computed defaults
func validateAndSetDefaults(cfg *Config) error {
	// Set GitBook defaults based on repository info
	if cfg.GitBook.Title == "" {
		cfg.GitBook.Title = cfg.Repository.Name
	}
	if cfg.GitBook.Description == "" {
		cfg.GitBook.Description = cfg.Repository.Description
	}

	// Ensure required fields are set
	if cfg.Repository.Name == "" {
		return fmt.Errorf("repository name is required")
	}

	return nil
}

// Save saves the configuration to a file
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
