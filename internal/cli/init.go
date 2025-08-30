package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/kolosys/proton/internal/config"
)

var (
	force bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-path]",
	Short: "Initialize a new Proton configuration",
	Long: `Initialize a new Proton configuration file in the specified project directory.

This command creates a .proton/config.yml file with sensible defaults for your project.
It will attempt to auto-detect repository information from go.mod and git configuration.

Examples:
  proton init                    # Initialize in current directory
  proton init ./my-project      # Initialize in specific project
  proton init --force           # Overwrite existing configuration`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	// Determine project path
	projectPath := "."
	if len(args) > 0 {
		projectPath = args[0]
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return fmt.Errorf("invalid project path: %w", err)
	}
	projectPath = absPath

	// Check if configuration already exists
	configDir := filepath.Join(projectPath, ".proton")
	configPath := filepath.Join(configDir, "config.yml")

	if _, err := os.Stat(configPath); err == nil && !force {
		return fmt.Errorf("configuration file already exists at %s (use --force to overwrite)", configPath)
	}

	// Create a default configuration
	cfg := createDefaultConfig(projectPath)

	// Auto-detect repository information
	if err := autoDetectProjectInfo(cfg, projectPath); err != nil {
		fmt.Printf("Warning: Could not auto-detect some project information: %v\n", err)
	}

	// Create configuration directory
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create configuration directory: %w", err)
	}

	// Save configuration
	if err := saveConfig(cfg, configPath); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	fmt.Printf("✅ Proton configuration initialized at %s\n", configPath)
	fmt.Printf("📝 Edit the configuration file to customize documentation generation\n")
	fmt.Printf("🚀 Run 'proton generate' to generate documentation\n")

	return nil
}

func createDefaultConfig(projectPath string) *config.Config {
	return &config.Config{
		Repository: config.Repository{
			// Branch will be auto-detected from Git
		},
		Output: config.Output{
			Directory:     filepath.Join(projectPath, "docs"),
			Clean:         true,
			GitBookConfig: true,
		},
		Discovery: config.Discovery{
			Packages: config.Packages{
				AutoDiscover:    true,
				IncludePatterns: []string{"./..."},
				ExcludePatterns: []string{"./vendor/...", "./test/...", "./.git/...", "**/*_test.go"},
			},
			APIGeneration: config.APIGeneration{
				Enabled:           true,
				IncludeUnexported: false,
				IncludeTests:      false,
				IncludeExamples:   true,
			},
			Examples: config.Examples{
				Enabled:      true,
				AutoDiscover: true,
			},
			Guides: config.Guides{
				Enabled:             true,
				IncludeContributing: true,
				IncludeFAQ:          true,
			},
		},
		GitBook: config.GitBook{
			Theme: "default",
			Structure: config.GitBookStructure{
				Readme:  "README.md",
				Summary: "SUMMARY.md",
			},
		},
		Metadata: config.Metadata{
			Version: "latest",
			License: "MIT",
		},
		Generation: config.Generation{
			DateFormat:             "2006-01-02",
			IncludeGeneratedNotice: true,
			IncludeTOC:             true,
			MaxDepth:               3,
		},
	}
}

func autoDetectProjectInfo(cfg *config.Config, projectPath string) error {
	// Try to read go.mod for module information
	goModPath := filepath.Join(projectPath, "go.mod")
	if data, err := os.ReadFile(goModPath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "module ") {
				modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module"))
				cfg.Repository.ImportPath = modulePath

				// Extract repository info from module path
				if strings.Contains(modulePath, "github.com/") {
					parts := strings.Split(modulePath, "/")
					if len(parts) >= 3 {
						cfg.Repository.Owner = parts[1]
						cfg.Repository.Name = parts[2]
					}
				}
				break
			}
		}
	}

	// Try to get Git information
	if err := detectGitInfo(cfg, projectPath); err == nil {
		// Git detection succeeded, use Git info as primary source
		if cfg.Repository.Owner != "" && cfg.Repository.Name != "" {
			cfg.Repository.URL = fmt.Sprintf("https://github.com/%s/%s", cfg.Repository.Owner, cfg.Repository.Name)
			cfg.GitBook.Title = cfg.Repository.Name
		}
	} else {
		// Fallback to go.mod derived info
		if cfg.Repository.Owner != "" && cfg.Repository.Name != "" {
			cfg.Repository.URL = fmt.Sprintf("https://github.com/%s/%s", cfg.Repository.Owner, cfg.Repository.Name)
			cfg.GitBook.Title = cfg.Repository.Name
		}
	}

	// Try to get description from README
	readmePath := filepath.Join(projectPath, "README.md")
	if data, err := os.ReadFile(readmePath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				cfg.Repository.Description = line
				cfg.GitBook.Description = line
				break
			}
		}
	}

	return nil
}

// detectGitInfo extracts repository information from Git configuration
func detectGitInfo(cfg *config.Config, projectPath string) error {
	// Get current branch
	if branch, err := exec.Command("git", "branch", "--show-current").Output(); err == nil {
		cfg.Repository.Branch = strings.TrimSpace(string(branch))
	}

	// Get remote URL
	if remoteURL, err := exec.Command("git", "config", "--get", "remote.origin.url").Output(); err == nil {
		url := strings.TrimSpace(string(remoteURL))

		// Parse SSH or HTTPS URLs
		if strings.HasPrefix(url, "git@github.com:") {
			// SSH format: git@github.com:owner/repo.git
			parts := strings.Split(strings.TrimSuffix(url, ".git"), ":")
			if len(parts) == 2 {
				repoParts := strings.Split(parts[1], "/")
				if len(repoParts) == 2 {
					cfg.Repository.Owner = repoParts[0]
					cfg.Repository.Name = repoParts[1]
				}
			}
		} else if strings.Contains(url, "github.com/") {
			// HTTPS format: https://github.com/owner/repo.git
			parts := strings.Split(strings.TrimSuffix(url, ".git"), "/")
			if len(parts) >= 2 {
				// Find github.com in the parts
				for i, part := range parts {
					if part == "github.com" && i+2 < len(parts) {
						cfg.Repository.Owner = parts[i+1]
						cfg.Repository.Name = parts[i+2]
						break
					}
				}
			}
		}
	}

	return nil
}

func saveConfig(cfg *config.Config, path string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write configuration file: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Local flags
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite existing configuration")
}
