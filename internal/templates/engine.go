package templates

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kolosys/proton/internal/config"
	"github.com/kolosys/proton/internal/discovery"
)

//go:embed builtin/*.md builtin/*.yml
var builtinTemplates embed.FS

// Engine handles template rendering for documentation generation
type Engine struct {
	config      *config.Config
	projectPath string
	templates   map[string]*template.Template
}

// Context provides data for template rendering
type Context struct {
	Repository config.Repository        `json:"repository"`
	Packages   []*discovery.PackageInfo `json:"packages"`
	Config     *config.Config           `json:"config"`
	Metadata   config.Metadata          `json:"metadata"`
}

// PackageContext provides package-specific data for template rendering
type PackageContext struct {
	*Context
	Package *discovery.PackageInfo `json:"package"`
}

// New creates a new template engine
func New(cfg *config.Config, projectPath string) (*Engine, error) {
	engine := &Engine{
		config:      cfg,
		projectPath: projectPath,
		templates:   make(map[string]*template.Template),
	}

	// Load built-in templates
	if err := engine.loadBuiltinTemplates(); err != nil {
		return nil, fmt.Errorf("failed to load builtin templates: %w", err)
	}

	// Load custom templates if specified
	if cfg.Templates.Directory != "" {
		if err := engine.loadCustomTemplates(); err != nil {
			return nil, fmt.Errorf("failed to load custom templates: %w", err)
		}
	}

	// Load individual custom template overrides
	for _, customTemplate := range cfg.Templates.CustomTemplates {
		if err := engine.loadCustomTemplate(customTemplate.Name, customTemplate.File); err != nil {
			return nil, fmt.Errorf("failed to load custom template %s: %w", customTemplate.Name, err)
		}
	}

	return engine, nil
}

// loadBuiltinTemplates loads the built-in templates from the embedded filesystem
func (e *Engine) loadBuiltinTemplates() error {
	templateNames := []string{
		"index",
		"getting-started-index",
		"getting-started",
		"index-api-reference",
		"api-reference",
		"examples-index",
		"package-examples",
		"guides-index",
		"contributing",
		"faq",
		"package-best-practices",
		"gitbook-config",
		"gitbook-summary",
	}

	for _, name := range templateNames {
		templatePath := fmt.Sprintf("builtin/%s.md", name)
		if name == "gitbook-config" {
			templatePath = "builtin/gitbook-config.yml"
		}

		content, err := builtinTemplates.ReadFile(templatePath)
		if err != nil {
			// Skip missing templates - they're optional
			continue
		}

		tmpl, err := template.New(name).Funcs(e.templateFuncs()).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse builtin template %s: %w", name, err)
		}

		e.templates[name] = tmpl
	}

	return nil
}

// loadCustomTemplates loads templates from the custom templates directory
func (e *Engine) loadCustomTemplates() error {
	templateDir := e.config.Templates.Directory
	if !filepath.IsAbs(templateDir) {
		templateDir = filepath.Join(e.projectPath, templateDir)
	}

	// Check if template directory exists
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return nil // No custom templates directory
	}

	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Only process .md and .yml files
		ext := filepath.Ext(path)
		if ext != ".md" && ext != ".yml" && ext != ".yaml" {
			return nil
		}

		// Determine template name from file path
		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		templateName := strings.TrimSuffix(relPath, ext)
		templateName = strings.ReplaceAll(templateName, string(filepath.Separator), "-")

		return e.loadCustomTemplate(templateName, path)
	})
}

// loadCustomTemplate loads a single custom template
func (e *Engine) loadCustomTemplate(name, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", filePath, err)
	}

	tmpl, err := template.New(name).Funcs(e.templateFuncs()).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", name, err)
	}

	e.templates[name] = tmpl
	return nil
}

// templateFuncs returns the functions available in templates
func (e *Engine) templateFuncs() template.FuncMap {
	return template.FuncMap{
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"title":     strings.Title,
		"join":      strings.Join,
		"replace":   strings.ReplaceAll,
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		"trim":      strings.TrimSpace,
		"indent": func(spaces int, text string) string {
			prefix := strings.Repeat(" ", spaces)
			lines := strings.Split(text, "\n")
			for i, line := range lines {
				if line != "" {
					lines[i] = prefix + line
				}
			}
			return strings.Join(lines, "\n")
		},
		"markdown": func(text string) string {
			// Simple markdown-like formatting for documentation
			return text
		},
		"codeBlock": func(lang, code string) string {
			return fmt.Sprintf("```%s\n%s\n```", lang, code)
		},
		"linkTo": func(path, text string) string {
			return fmt.Sprintf("[%s](%s)", text, path)
		},
		"packagePath": func(pkg *discovery.PackageInfo) string {
			return strings.TrimPrefix(pkg.ImportPath, e.config.Repository.ImportPath+"/")
		},
		"isMainPackage": func(pkg *discovery.PackageInfo) bool {
			return pkg.Name == "main" || strings.Contains(pkg.Path, "/cmd/")
		},
		"hasExamples": func(pkg *discovery.PackageInfo) bool {
			return len(pkg.Examples) > 0
		},
		"formatExampleOutput": func(output string) string {
			if output == "" {
				return ""
			}
			return fmt.Sprintf("// Output:\n// %s", strings.ReplaceAll(output, "\n", "\n// "))
		},
		"hasFields": func(typ *discovery.EnhancedType) bool {
			return len(typ.Fields) > 0
		},
		"hasParams": func(fn *discovery.EnhancedFunc) bool {
			return len(fn.Params) > 0
		},
		"hasResults": func(fn *discovery.EnhancedFunc) bool {
			return len(fn.Results) > 0
		},
		"formatFieldName": func(field *discovery.Field) string {
			if field.Name == "" {
				return "*" + field.Type // Embedded field
			}
			return field.Name
		},
		"formatTag": func(tag string) string {
			if tag == "" {
				return ""
			}
			return "`" + strings.Trim(tag, "`") + "`"
		},
		"typeLink": func(typeName string) string {
			// Simple type linking - could be enhanced for cross-references
			return typeName
		},
	}
}

// RenderToFile renders a template to a file
func (e *Engine) RenderToFile(templateName string, data interface{}, outputPath string) error {
	tmpl, exists := e.templates[templateName]
	if !exists {
		return fmt.Errorf("template %s not found", templateName)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer file.Close()

	// Render template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to render template %s: %w", templateName, err)
	}

	return nil
}

// RenderToString renders a template to a string
func (e *Engine) RenderToString(templateName string, data interface{}) (string, error) {
	tmpl, exists := e.templates[templateName]
	if !exists {
		return "", fmt.Errorf("template %s not found", templateName)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to render template %s: %w", templateName, err)
	}

	return buf.String(), nil
}

// ListTemplates returns a list of available template names
func (e *Engine) ListTemplates() []string {
	var names []string
	for name := range e.templates {
		names = append(names, name)
	}
	return names
}

// HasTemplate checks if a template exists
func (e *Engine) HasTemplate(name string) bool {
	_, exists := e.templates[name]
	return exists
}
