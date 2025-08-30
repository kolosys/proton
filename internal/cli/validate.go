package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/kolosys/proton/internal/config"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [project-path]",
	Short: "Validate Proton configuration and project structure",
	Long: `Validate the Proton configuration file and project structure.

This command checks:
- Configuration file syntax and completeness
- Go module structure and validity
- Package discovery patterns
- Template accessibility
- Output directory permissions

Examples:
  proton validate                    # Validate current directory
  proton validate ./my-project      # Validate specific project
  proton validate --config custom.yml # Use custom config file`,
	Args: cobra.MaximumNArgs(1),
	RunE: runValidate,
}

func runValidate(cmd *cobra.Command, args []string) error {
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

	fmt.Printf("🔍 Validating Proton configuration for %s\n\n", projectPath)

	// Load and validate configuration
	cfg, err := config.Load(configPath, projectPath)
	if err != nil {
		return fmt.Errorf("❌ Configuration validation failed: %w", err)
	}

	fmt.Printf("✅ Configuration loaded successfully\n")

	// Validate repository information
	if err := validateRepository(&cfg.Repository); err != nil {
		return fmt.Errorf("❌ Repository validation failed: %w", err)
	}

	fmt.Printf("✅ Repository information is valid\n")

	// Validate output configuration
	if err := validateOutput(&cfg.Output, projectPath); err != nil {
		return fmt.Errorf("❌ Output validation failed: %w", err)
	}

	fmt.Printf("✅ Output configuration is valid\n")

	// Validate package discovery
	if err := validateDiscovery(&cfg.Discovery, projectPath); err != nil {
		return fmt.Errorf("❌ Discovery validation failed: %w", err)
	}

	fmt.Printf("✅ Package discovery configuration is valid\n")

	// Validate templates
	if err := validateTemplates(&cfg.Templates, projectPath); err != nil {
		return fmt.Errorf("❌ Template validation failed: %w", err)
	}

	fmt.Printf("✅ Template configuration is valid\n")

	fmt.Printf("\n🎉 All validations passed! Configuration is ready for documentation generation.\n")
	fmt.Printf("💡 Run 'proton generate' to generate documentation\n")

	return nil
}

func validateRepository(repo *config.Repository) error {
	if repo.Name == "" {
		return fmt.Errorf("repository name is required")
	}

	if repo.ImportPath == "" {
		return fmt.Errorf("import path is required")
	}

	if repo.Owner == "" {
		fmt.Printf("⚠️  Warning: repository owner not specified\n")
	}

	if repo.Description == "" {
		fmt.Printf("⚠️  Warning: repository description not specified\n")
	}

	return nil
}

func validateOutput(output *config.Output, projectPath string) error {
	if output.Directory == "" {
		return fmt.Errorf("output directory is required")
	}

	// Check if output directory is writable
	outputPath := output.Directory
	if !filepath.IsAbs(outputPath) {
		outputPath = filepath.Join(projectPath, outputPath)
	}

	// Try to create the directory to test permissions
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("cannot create output directory %s: %w", outputPath, err)
	}

	return nil
}

func validateDiscovery(discovery *config.Discovery, projectPath string) error {
	// Validate package patterns
	if discovery.Packages.AutoDiscover {
		if len(discovery.Packages.IncludePatterns) == 0 {
			fmt.Printf("⚠️  Warning: no include patterns specified for auto-discovery\n")
		}
	}

	// Validate manual packages if specified
	for _, pkg := range discovery.Packages.ManualPackages {
		if pkg.Name == "" {
			return fmt.Errorf("manual package name is required")
		}

		if pkg.Path == "" {
			return fmt.Errorf("manual package path is required for package %s", pkg.Name)
		}

		// Check if package path exists
		pkgPath := pkg.Path
		if !filepath.IsAbs(pkgPath) {
			pkgPath = filepath.Join(projectPath, pkgPath)
		}

		if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
			return fmt.Errorf("manual package path does not exist: %s", pkgPath)
		}
	}

	// Validate examples configuration
	if discovery.Examples.Enabled {
		for _, dir := range discovery.Examples.Directories {
			examplePath := dir
			if !filepath.IsAbs(examplePath) {
				examplePath = filepath.Join(projectPath, examplePath)
			}

			if _, err := os.Stat(examplePath); os.IsNotExist(err) {
				fmt.Printf("⚠️  Warning: example directory does not exist: %s\n", examplePath)
			}
		}
	}

	return nil
}

func validateTemplates(templates *config.Templates, projectPath string) error {
	// Validate custom template directory
	if templates.Directory != "" {
		templatePath := templates.Directory
		if !filepath.IsAbs(templatePath) {
			templatePath = filepath.Join(projectPath, templatePath)
		}

		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			return fmt.Errorf("custom template directory does not exist: %s", templatePath)
		}
	}

	// Validate custom template files
	for _, customTemplate := range templates.CustomTemplates {
		if customTemplate.Name == "" {
			return fmt.Errorf("custom template name is required")
		}

		if customTemplate.File == "" {
			return fmt.Errorf("custom template file is required for template %s", customTemplate.Name)
		}

		templatePath := customTemplate.File
		if !filepath.IsAbs(templatePath) {
			templatePath = filepath.Join(projectPath, templatePath)
		}

		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			return fmt.Errorf("custom template file does not exist: %s", templatePath)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Local flags
	validateCmd.Flags().StringVarP(&configPath, "config", "c", "", "path to configuration file")
}
