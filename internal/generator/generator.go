package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolosys/proton/internal/config"
	"github.com/kolosys/proton/internal/discovery"
	"github.com/kolosys/proton/internal/templates"
)

// Generator handles the complete documentation generation process
type Generator struct {
	config      *config.Config
	projectPath string
	outputPath  string
	discoverer  *discovery.Discoverer
	templates   *templates.Engine
}

// New creates a new documentation generator
func New(cfg *config.Config, projectPath string) (*Generator, error) {
	// Resolve output path
	outputPath := cfg.Output.Directory
	if !filepath.IsAbs(outputPath) {
		outputPath = filepath.Join(projectPath, outputPath)
	}

	// Create discoverer
	discoverer := discovery.New(cfg, projectPath)

	// Create template engine
	templateEngine, err := templates.New(cfg, projectPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create template engine: %w", err)
	}

	return &Generator{
		config:      cfg,
		projectPath: projectPath,
		outputPath:  outputPath,
		discoverer:  discoverer,
		templates:   templateEngine,
	}, nil
}

// Generate performs the complete documentation generation
func (g *Generator) Generate() error {
	// Clean output directory if requested
	if g.config.Output.Clean {
		if err := g.cleanOutputDirectory(); err != nil {
			return fmt.Errorf("failed to clean output directory: %w", err)
		}
	}

	// Ensure output directory exists
	if err := os.MkdirAll(g.outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Discover packages
	packages, err := g.discoverer.DiscoverPackages()
	if err != nil {
		return fmt.Errorf("package discovery failed: %w", err)
	}

	// Create template context
	context := g.createTemplateContext(packages)

	// Generate main documentation files
	if err := g.generateMainFiles(context); err != nil {
		return fmt.Errorf("failed to generate main files: %w", err)
	}

	// Generate package documentation
	if err := g.generatePackageDocumentation(packages, context); err != nil {
		return fmt.Errorf("failed to generate package documentation: %w", err)
	}

	// Generate API reference documentation
	if g.config.Discovery.APIGeneration.Enabled {
		if err := g.generateAPIDocumentation(packages, context); err != nil {
			return fmt.Errorf("failed to generate API documentation: %w", err)
		}
	}

	// Generate examples documentation
	if g.config.Discovery.Examples.Enabled {
		if err := g.generateExamplesDocumentation(packages, context); err != nil {
			return fmt.Errorf("failed to generate examples documentation: %w", err)
		}
	}

	// Generate guides documentation
	if g.config.Discovery.Guides.Enabled {
		if err := g.generateGuidesDocumentation(context); err != nil {
			return fmt.Errorf("failed to generate guides documentation: %w", err)
		}
	}

	// Generate GitBook configuration
	if g.config.Output.GitBookConfig {
		if err := g.generateGitBookConfig(context); err != nil {
			return fmt.Errorf("failed to generate GitBook configuration: %w", err)
		}
	}

	return nil
}

// cleanOutputDirectory removes all files from the output directory
func (g *Generator) cleanOutputDirectory() error {
	if _, err := os.Stat(g.outputPath); os.IsNotExist(err) {
		return nil // Directory doesn't exist, nothing to clean
	}

	entries, err := os.ReadDir(g.outputPath)
	if err != nil {
		return fmt.Errorf("failed to read output directory: %w", err)
	}

	for _, entry := range entries {
		path := filepath.Join(g.outputPath, entry.Name())
		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to remove %s: %w", path, err)
		}
	}

	return nil
}

// createTemplateContext creates the context object for template rendering
func (g *Generator) createTemplateContext(packages []*discovery.PackageInfo) *templates.Context {
	return &templates.Context{
		Repository: g.config.Repository,
		Packages:   packages,
		Config:     g.config,
		Metadata:   g.config.Metadata,
	}
}

// generateMainFiles generates the main documentation files
func (g *Generator) generateMainFiles(context *templates.Context) error {
	// Generate main README/index
	indexPath := filepath.Join(g.outputPath, "README.md")
	if err := g.templates.RenderToFile("index", context, indexPath); err != nil {
		return fmt.Errorf("failed to generate main index: %w", err)
	}

	return nil
}

// generatePackageDocumentation generates getting-started documentation
func (g *Generator) generatePackageDocumentation(packages []*discovery.PackageInfo, context *templates.Context) error {
	// Create getting-started directory
	gettingStartedDir := filepath.Join(g.outputPath, "getting-started")
	if err := os.MkdirAll(gettingStartedDir, 0755); err != nil {
		return fmt.Errorf("failed to create getting-started directory: %w", err)
	}

	// Generate getting-started index
	gettingStartedIndexPath := filepath.Join(gettingStartedDir, "README.md")
	if err := g.templates.RenderToFile("getting-started-index", context, gettingStartedIndexPath); err != nil {
		return fmt.Errorf("failed to generate getting-started index: %w", err)
	}

	// Generate individual getting-started documentation for each package
	for _, pkg := range packages {
		pkgContext := &templates.PackageContext{
			Context: context,
			Package: pkg,
		}

		pkgPath := filepath.Join(gettingStartedDir, fmt.Sprintf("%s.md", pkg.Name))
		if err := g.templates.RenderToFile("getting-started", pkgContext, pkgPath); err != nil {
			return fmt.Errorf("failed to generate getting-started documentation for package %s: %w", pkg.Name, err)
		}
	}

	return nil
}

// generateAPIDocumentation generates API reference documentation
func (g *Generator) generateAPIDocumentation(packages []*discovery.PackageInfo, context *templates.Context) error {
	// Create API reference directory
	apiDir := filepath.Join(g.outputPath, "api-reference")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		return fmt.Errorf("failed to create API reference directory: %w", err)
	}

	// Generate API reference index
	apiIndexPath := filepath.Join(apiDir, "README.md")
	if err := g.templates.RenderToFile("index-api-reference", context, apiIndexPath); err != nil {
		return fmt.Errorf("failed to generate API reference index: %w", err)
	}

	// Generate individual API documentation for each package
	for _, pkg := range packages {
		pkgContext := &templates.PackageContext{
			Context: context,
			Package: pkg,
		}

		apiPath := filepath.Join(apiDir, fmt.Sprintf("%s.md", pkg.Name))
		if err := g.templates.RenderToFile("api-reference", pkgContext, apiPath); err != nil {
			return fmt.Errorf("failed to generate API reference for package %s: %w", pkg.Name, err)
		}
	}

	return nil
}

// generateExamplesDocumentation generates examples documentation
func (g *Generator) generateExamplesDocumentation(packages []*discovery.PackageInfo, context *templates.Context) error {
	// Create examples directory
	examplesDir := filepath.Join(g.outputPath, "examples")
	if err := os.MkdirAll(examplesDir, 0755); err != nil {
		return fmt.Errorf("failed to create examples directory: %w", err)
	}

	// Generate examples index
	examplesIndexPath := filepath.Join(examplesDir, "README.md")
	if err := g.templates.RenderToFile("examples-index", context, examplesIndexPath); err != nil {
		return fmt.Errorf("failed to generate examples index: %w", err)
	}

	// Discover and generate examples from configured directories
	exampleDirectories, err := g.discoverExampleDirectories()
	if err != nil {
		return fmt.Errorf("failed to discover example directories: %w", err)
	}

	// Generate documentation for each example directory
	for _, exampleDir := range exampleDirectories {
		// Get relative path from project root
		relPath, err := filepath.Rel(g.projectPath, exampleDir)
		if err != nil {
			relPath = filepath.Base(exampleDir)
		}

		// If this is the main examples directory, process its contents directly
		if relPath == "examples" {
			// Process contents directly without creating a nested structure
			entries, err := os.ReadDir(exampleDir)
			if err != nil {
				return fmt.Errorf("failed to read examples directory: %w", err)
			}

			for _, entry := range entries {
				name := entry.Name()

				// Skip config files
				if strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml") {
					continue
				}

				sourcePath := filepath.Join(exampleDir, name)
				destPath := filepath.Join(examplesDir, name)

				if entry.IsDir() {
					// Create subdirectory and generate documentation
					if err := os.MkdirAll(destPath, 0755); err != nil {
						return fmt.Errorf("failed to create subdirectory %s: %w", destPath, err)
					}

					// Generate README for subdirectory
					if err := g.generateExampleDirectoryREADME(sourcePath, destPath, name, context); err != nil {
						return fmt.Errorf("failed to generate README for subdirectory %s: %w", name, err)
					}

					// Recursively generate documentation for subdirectory
					if err := g.generateExampleSubdirectoryDocumentation(sourcePath, destPath, context); err != nil {
						return fmt.Errorf("failed to generate documentation for subdirectory %s: %w", name, err)
					}
				} else if strings.HasSuffix(name, ".go") {
					// Generate markdown for Go files
					if err := g.generateExampleFileMarkdown(sourcePath, examplesDir, name); err != nil {
						return fmt.Errorf("failed to generate markdown for %s: %w", name, err)
					}
				}
			}
		} else {
			// For other directories, create a subdirectory
			if err := g.generateExampleDirectoryDocumentation(exampleDir, examplesDir, context); err != nil {
				return fmt.Errorf("failed to generate examples for directory %s: %w", exampleDir, err)
			}
		}
	}

	return nil
}

// discoverExampleDirectories discovers example directories based on configuration
func (g *Generator) discoverExampleDirectories() ([]string, error) {
	var directories []string

	// Add configured directories
	for _, dir := range g.config.Discovery.Examples.Directories {
		// Resolve relative path
		examplePath := dir
		if !filepath.IsAbs(examplePath) {
			examplePath = filepath.Join(g.projectPath, examplePath)
		}

		// Check if directory exists
		if _, err := os.Stat(examplePath); os.IsNotExist(err) {
			continue // Skip non-existent directories
		}

		directories = append(directories, examplePath)
	}

	// Auto-discover examples if enabled
	if g.config.Discovery.Examples.AutoDiscover {
		// Look for common example directory patterns
		commonPatterns := []string{
			filepath.Join(g.projectPath, "examples"),
			filepath.Join(g.projectPath, "example"),
			filepath.Join(g.projectPath, "samples"),
			filepath.Join(g.projectPath, "sample"),
		}

		for _, pattern := range commonPatterns {
			if _, err := os.Stat(pattern); err == nil {
				directories = append(directories, pattern)
			}
		}

		// Look for examples in package directories
		for _, pkg := range g.config.Discovery.Packages.ManualPackages {
			pkgExamplesDir := filepath.Join(g.projectPath, pkg.Path, "examples")
			if _, err := os.Stat(pkgExamplesDir); err == nil {
				directories = append(directories, pkgExamplesDir)
			}
		}
	}

	return directories, nil
}

// generateExampleDirectoryDocumentation generates documentation for a single example directory
func (g *Generator) generateExampleDirectoryDocumentation(sourceDir, outputBaseDir string, context *templates.Context) error {
	// Get relative path from project root for naming
	relPath, err := filepath.Rel(g.projectPath, sourceDir)
	if err != nil {
		relPath = filepath.Base(sourceDir)
	}

	// Create output directory
	outputDir := filepath.Join(outputBaseDir, relPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	// Generate README for this example directory
	if err := g.generateExampleDirectoryREADME(sourceDir, outputDir, relPath, context); err != nil {
		return fmt.Errorf("failed to generate README for %s: %w", sourceDir, err)
	}

	// Generate markdown documentation for subdirectories and files
	if err := g.generateExampleSubdirectoryDocumentation(sourceDir, outputDir, context); err != nil {
		return fmt.Errorf("failed to generate example documentation from %s: %w", sourceDir, err)
	}

	return nil
}

// generateExampleDirectoryREADME generates a README for an example directory
func (g *Generator) generateExampleDirectoryREADME(sourceDir, outputDir, relPath string, context *templates.Context) error {
	// For now, generate a simple README
	// TODO: Create a proper template for example directories
	readmeContent := fmt.Sprintf(`# Examples: %s

This directory contains examples from: `+sourceDir+`

## Files

`, relPath)

	// List files in the example directory
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read example directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()

		// Skip config files and only list Go files and directories
		if !entry.IsDir() && !strings.HasSuffix(name, ".go") {
			continue
		}

		// For Go files, link to the markdown version
		if strings.HasSuffix(name, ".go") {
			markdownName := strings.TrimSuffix(name, ".go") + ".md"
			readmeContent += fmt.Sprintf("- [%s](%s)\n", name, markdownName)
		} else {
			readmeContent += fmt.Sprintf("- [%s](%s)\n", name, name)
		}
	}

	readmePath := filepath.Join(outputDir, "README.md")
	return os.WriteFile(readmePath, []byte(readmeContent), 0644)
}

// generateExampleSubdirectoryDocumentation generates markdown documentation for example subdirectories
func (g *Generator) generateExampleSubdirectoryDocumentation(sourceDir, outputDir string, context *templates.Context) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()

		// Skip config files
		if strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml") {
			continue
		}

		sourcePath := filepath.Join(sourceDir, name)
		destPath := filepath.Join(outputDir, name)

		if entry.IsDir() {
			// Create subdirectory and generate documentation
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return fmt.Errorf("failed to create subdirectory %s: %w", destPath, err)
			}

			// Generate README for subdirectory
			if err := g.generateExampleDirectoryREADME(sourcePath, destPath, name, context); err != nil {
				return fmt.Errorf("failed to generate README for subdirectory %s: %w", name, err)
			}

			// Recursively generate documentation for subdirectory
			if err := g.generateExampleSubdirectoryDocumentation(sourcePath, destPath, context); err != nil {
				return fmt.Errorf("failed to generate documentation for subdirectory %s: %w", name, err)
			}
		} else if strings.HasSuffix(name, ".go") {
			// Generate markdown for Go files
			if err := g.generateExampleFileMarkdown(sourcePath, outputDir, name); err != nil {
				return fmt.Errorf("failed to generate markdown for %s: %w", name, err)
			}
		}
	}

	return nil
}

// generateExampleFileMarkdown generates markdown documentation for a single example file
func (g *Generator) generateExampleFileMarkdown(sourcePath, outputDir, fileName string) error {
	// Read the source file
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source file %s: %w", sourcePath, err)
	}

	// Create markdown filename (replace .go with .md)
	markdownName := strings.TrimSuffix(fileName, ".go") + ".md"
	markdownPath := filepath.Join(outputDir, markdownName)

	// Generate markdown content
	markdownContent := fmt.Sprintf(`# %s

This example demonstrates basic usage of the library.

## Source Code

`+"```go"+`
%s
`+"```"+`

## Running the Example

To run this example:

`+"```bash"+`
cd %s
go run %s
`+"```"+`

## Expected Output

`+"```"+`
Hello from Proton examples!
`+"```"+`
`,
		strings.TrimSuffix(fileName, ".go"),
		string(content),
		filepath.Base(outputDir),
		fileName,
	)

	// Write markdown file
	return os.WriteFile(markdownPath, []byte(markdownContent), 0644)
}

// generateGuidesDocumentation generates guides documentation
func (g *Generator) generateGuidesDocumentation(context *templates.Context) error {
	// Create guides directory
	guidesDir := filepath.Join(g.outputPath, "guides")
	if err := os.MkdirAll(guidesDir, 0755); err != nil {
		return fmt.Errorf("failed to create guides directory: %w", err)
	}

	// Generate guides index
	guidesIndexPath := filepath.Join(guidesDir, "README.md")
	if err := g.templates.RenderToFile("guides-index", context, guidesIndexPath); err != nil {
		return fmt.Errorf("failed to generate guides index: %w", err)
	}

	// Generate per-package guides
	for _, pkg := range context.Packages {
		pkgContext := &templates.PackageContext{
			Context: context,
			Package: pkg,
		}

		// Create package-specific guides directory
		pkgGuidesDir := filepath.Join(guidesDir, pkg.Name)
		if err := os.MkdirAll(pkgGuidesDir, 0755); err != nil {
			return fmt.Errorf("failed to create guides directory for package %s: %w", pkg.Name, err)
		}

		// Generate package-specific best practices
		bestPracticesPath := filepath.Join(pkgGuidesDir, "best-practices.md")
		if err := g.templates.RenderToFile("package-best-practices", pkgContext, bestPracticesPath); err != nil {
			return fmt.Errorf("failed to generate best practices for package %s: %w", pkg.Name, err)
		}
	}

	// Generate global guides if enabled
	if g.config.Discovery.Guides.IncludeContributing {
		contributingPath := filepath.Join(guidesDir, "contributing.md")
		if err := g.templates.RenderToFile("contributing", context, contributingPath); err != nil {
			return fmt.Errorf("failed to generate contributing guide: %w", err)
		}
	}

	if g.config.Discovery.Guides.IncludeFAQ {
		faqPath := filepath.Join(guidesDir, "faq.md")
		if err := g.templates.RenderToFile("faq", context, faqPath); err != nil {
			return fmt.Errorf("failed to generate FAQ: %w", err)
		}
	}

	// Generate custom guides
	for _, guide := range g.config.Discovery.Guides.CustomGuides {
		guidePath := filepath.Join(guidesDir, fmt.Sprintf("%s.md", guide.Name))
		if err := g.templates.RenderToFile(guide.Name, context, guidePath); err != nil {
			return fmt.Errorf("failed to generate custom guide %s: %w", guide.Name, err)
		}
	}

	return nil
}

// generateGitBookConfig generates the .gitbook.yml configuration file
func (g *Generator) generateGitBookConfig(context *templates.Context) error {
	configPath := filepath.Join(g.outputPath, ".gitbook.yml")
	if err := g.templates.RenderToFile("gitbook-config", context, configPath); err != nil {
		return fmt.Errorf("failed to generate GitBook configuration: %w", err)
	}

	// Also generate SUMMARY.md for GitBook navigation
	summaryPath := filepath.Join(g.outputPath, "SUMMARY.md")
	if err := g.templates.RenderToFile("gitbook-summary", context, summaryPath); err != nil {
		return fmt.Errorf("failed to generate GitBook summary: %w", err)
	}

	return nil
}
