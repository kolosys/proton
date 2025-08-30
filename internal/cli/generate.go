package cli

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kolosys/proton/internal/config"
	"github.com/kolosys/proton/internal/generator"
)

var (
	outputDir   string
	clean       bool
	configPath  string
	projectPath string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [project-path]",
	Short: "Generate documentation for a Go project",
	Long: `Generate comprehensive documentation for a Go project.

This command will:
- Discover Go packages in the project
- Parse Go source code and comments
- Generate GitBook-compatible documentation
- Create .gitbook.yml configuration
- Apply custom templates if configured

Examples:
  proton generate                    # Generate docs for current directory
  proton generate ./my-project      # Generate docs for specific project
  proton generate --output docs     # Generate with custom output directory
  proton generate --clean=false     # Don't clean output directory`,
	Args: cobra.MaximumNArgs(1),
	RunE: runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Determine project path
	if len(args) > 0 {
		projectPath = args[0]
	} else {
		projectPath = "."
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return fmt.Errorf("invalid project path: %w", err)
	}
	projectPath = absPath

	// Load configuration
	cfg, err := config.Load(configPath, projectPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Override with command line flags if provided
	if outputDir != "" {
		cfg.Output.Directory = outputDir
	}
	if cmd.Flags().Changed("clean") {
		cfg.Output.Clean = clean
	}

	// Create generator
	gen, err := generator.New(cfg, projectPath)
	if err != nil {
		return fmt.Errorf("failed to create generator: %w", err)
	}

	// Generate documentation
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	fmt.Printf("Documentation generated successfully in %s\n", cfg.Output.Directory)
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Local flags
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "", "output directory (default: docs)")
	generateCmd.Flags().BoolVar(&clean, "clean", true, "clean output directory before generation")
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "", "path to configuration file")

	// Bind flags to viper
	viper.BindPFlag("output.directory", generateCmd.Flags().Lookup("output"))
	viper.BindPFlag("output.clean", generateCmd.Flags().Lookup("clean"))
}
