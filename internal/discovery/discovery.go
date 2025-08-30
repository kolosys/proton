package discovery

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolosys/proton/internal/config"
)

// PackageInfo contains information about a discovered Go package
type PackageInfo struct {
	Name        string
	Path        string
	ImportPath  string
	Description string
	Doc         *doc.Package
	Functions   []*EnhancedFunc
	Types       []*EnhancedType
	Variables   []*doc.Value
	Constants   []*doc.Value
	Examples    []*doc.Example
	Files       []string
}

// EnhancedFunc extends doc.Func with additional parameter and return information
type EnhancedFunc struct {
	*doc.Func
	Params      []*Parameter
	Results     []*Result
	ExampleCode string
	Declaration string // Clean formatted function declaration
	Doc         string // Enhanced documentation (may override doc.Func.Doc)
}

// EnhancedType extends doc.Type with enhanced field information
type EnhancedType struct {
	*doc.Type
	Fields      []*Field
	Methods     []*EnhancedFunc
	Funcs       []*EnhancedFunc
	TypeKind    string // struct, interface, type alias, etc.
	Declaration string // Clean formatted declaration
	Doc         string // Enhanced documentation (may override doc.Type.Doc)
	ExampleCode string // Usage example code
}

// Parameter represents a function parameter
type Parameter struct {
	Name string
	Type string
	Doc  string
}

// Result represents a function return value
type Result struct {
	Name string
	Type string
	Doc  string
}

// Field represents a struct field
type Field struct {
	Name string
	Type string
	Tag  string
	Doc  string
}

// Discoverer handles package discovery and parsing
type Discoverer struct {
	config      *config.Config
	projectPath string
	fileSet     *token.FileSet
}

// New creates a new package discoverer
func New(cfg *config.Config, projectPath string) *Discoverer {
	return &Discoverer{
		config:      cfg,
		projectPath: projectPath,
		fileSet:     token.NewFileSet(),
	}
}

// DiscoverPackages discovers all packages in the project according to configuration
func (d *Discoverer) DiscoverPackages() ([]*PackageInfo, error) {
	var allPackages []*PackageInfo

	// Auto-discover packages if enabled
	if d.config.Discovery.Packages.AutoDiscover {
		discovered, err := d.autoDiscoverPackages()
		if err != nil {
			return nil, fmt.Errorf("auto-discovery failed: %w", err)
		}
		allPackages = append(allPackages, discovered...)
	}

	// Add manually specified packages
	for _, manualPkg := range d.config.Discovery.Packages.ManualPackages {
		pkgInfo, err := d.parsePackage(manualPkg.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to parse manual package %s: %w", manualPkg.Path, err)
		}

		// Override with manual configuration
		if manualPkg.Name != "" {
			pkgInfo.Name = manualPkg.Name
		}
		if manualPkg.Description != "" {
			pkgInfo.Description = manualPkg.Description
		}

		allPackages = append(allPackages, pkgInfo)
	}

	return allPackages, nil
}

// autoDiscoverPackages automatically discovers packages using the configured patterns
func (d *Discoverer) autoDiscoverPackages() ([]*PackageInfo, error) {
	// Use parser.ParseDir for better AST comment preservation
	// This approach preserves all AST comments needed for documentation extraction

	// Get all Go files in the project
	var allPackages []*PackageInfo

	// Walk through the project directory to find packages
	err := filepath.Walk(d.projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip vendor and test directories unless explicitly included
		if info.IsDir() {
			if strings.Contains(path, "/vendor/") ||
				(strings.Contains(path, "/test/") && !d.config.Discovery.APIGeneration.IncludeTests) {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process .go files
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Skip test files unless explicitly included
		if strings.HasSuffix(info.Name(), "_test.go") && !d.config.Discovery.APIGeneration.IncludeTests {
			return nil
		}

		// Get the package directory
		pkgDir := filepath.Dir(path)

		// Check if we've already processed this package
		for _, pkg := range allPackages {
			if pkg.Path == pkgDir {
				return nil
			}
		}

		// Check if package should be excluded
		relPath, err := filepath.Rel(d.projectPath, pkgDir)
		if err != nil {
			return nil
		}

		if d.shouldExcludePackage(relPath) {
			return nil
		}

		// Parse the package
		pkgInfo, err := d.parsePackage(relPath)
		if err != nil {
			// Log error but continue with other packages
			fmt.Printf("Warning: failed to parse package %s: %v\n", relPath, err)
			return nil
		}

		if pkgInfo != nil {
			allPackages = append(allPackages, pkgInfo)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk project directory: %w", err)
	}

	return allPackages, nil
}

// shouldExcludePackage checks if a package should be excluded based on patterns
func (d *Discoverer) shouldExcludePackage(pkgPath string) bool {
	for _, pattern := range d.config.Discovery.Packages.ExcludePatterns {
		if matched, _ := filepath.Match(pattern, pkgPath); matched {
			return true
		}
		// Also check if the pattern matches any part of the path
		if strings.Contains(pkgPath, strings.Trim(pattern, "*")) {
			return true
		}
	}
	return false
}

// parsePackage parses a single package from a given path
func (d *Discoverer) parsePackage(pkgPath string) (*PackageInfo, error) {
	// Resolve relative path
	fullPath := filepath.Join(d.projectPath, pkgPath)
	if !filepath.IsAbs(pkgPath) {
		fullPath = filepath.Join(d.projectPath, pkgPath)
	} else {
		fullPath = pkgPath
	}

	// Parse package directory with all comment modes
	pkgs, err := parser.ParseDir(d.fileSet, fullPath, func(info os.FileInfo) bool {
		name := info.Name()
		// Skip test files unless explicitly included
		if strings.HasSuffix(name, "_test.go") && !d.config.Discovery.APIGeneration.IncludeTests {
			return false
		}
		return strings.HasSuffix(name, ".go")
	}, parser.ParseComments|parser.AllErrors)

	if err != nil {
		return nil, fmt.Errorf("failed to parse directory %s: %w", fullPath, err)
	}

	// Find the main package (non-test package)
	var mainPkg *ast.Package
	for name, pkg := range pkgs {
		if !strings.HasSuffix(name, "_test") {
			mainPkg = pkg
			break
		}
	}

	if mainPkg == nil {
		return nil, fmt.Errorf("no non-test package found in %s", fullPath)
	}

	return d.parseASTPackage(mainPkg, fullPath)
}

// parseASTPackage creates PackageInfo from an AST package
func (d *Discoverer) parseASTPackage(astPkg *ast.Package, pkgPath string) (*PackageInfo, error) {
	// Create doc package - always use AllDecls for better documentation extraction
	docPkg := doc.New(astPkg, "./", doc.AllDecls)

	// Extract package description from doc
	description := docPkg.Doc
	if description == "" && len(docPkg.Doc) > 0 {
		// Take first sentence as description
		parts := strings.Split(docPkg.Doc, ". ")
		if len(parts) > 0 {
			description = parts[0]
		}
	}

	// Get file list
	var files []string
	for filename := range astPkg.Files {
		files = append(files, filename)
	}

	// Determine import path
	importPath := d.getImportPath(pkgPath)

	// Filter and enhance functions to only include public ones (starting with uppercase)
	var enhancedFuncs []*EnhancedFunc
	for _, fn := range docPkg.Funcs {
		if len(fn.Name) > 0 && strings.ToUpper(fn.Name[:1]) == fn.Name[:1] {
			enhancedFuncs = append(enhancedFuncs, d.enhanceFunction(fn, astPkg))
		}
	}

	// Filter and enhance types to only include public ones (starting with uppercase)
	var enhancedTypes []*EnhancedType
	for _, typ := range docPkg.Types {
		if len(typ.Name) > 0 && strings.ToUpper(typ.Name[:1]) == typ.Name[:1] {
			enhancedTypes = append(enhancedTypes, d.enhanceType(typ, astPkg))
		}
	}

	// Filter variables to only include public ones (starting with uppercase)
	var publicVars []*doc.Value
	for _, v := range docPkg.Vars {
		if len(v.Names) > 0 && len(v.Names[0]) > 0 && strings.ToUpper(v.Names[0][:1]) == v.Names[0][:1] {
			publicVars = append(publicVars, v)
		}
	}

	// Filter constants to only include public ones (starting with uppercase)
	var publicConsts []*doc.Value
	for _, c := range docPkg.Consts {
		if len(c.Names) > 0 && len(c.Names[0]) > 0 && strings.ToUpper(c.Names[0][:1]) == c.Names[0][:1] {
			publicConsts = append(publicConsts, c)
		}
	}

	pkgInfo := &PackageInfo{
		Name:        astPkg.Name,
		Path:        pkgPath,
		ImportPath:  importPath,
		Description: description,
		Doc:         docPkg,
		Functions:   enhancedFuncs,
		Types:       enhancedTypes,
		Variables:   publicVars,
		Constants:   publicConsts,
		Files:       files,
	}

	// Extract examples if enabled
	if d.config.Discovery.APIGeneration.IncludeExamples {
		examples := d.extractExamples(astPkg)
		pkgInfo.Examples = examples
	}

	return pkgInfo, nil
}

// getImportPath determines the import path for a package
func (d *Discoverer) getImportPath(pkgPath string) string {
	// Try to determine import path relative to the module root
	relPath, err := filepath.Rel(d.projectPath, pkgPath)
	if err != nil {
		return ""
	}

	// Clean up the path
	relPath = filepath.ToSlash(relPath)
	if relPath == "." {
		return d.config.Repository.ImportPath
	}

	return d.config.Repository.ImportPath + "/" + relPath
}

// extractExamples extracts example functions from the package
func (d *Discoverer) extractExamples(astPkg *ast.Package) []*doc.Example {
	var examples []*doc.Example

	for _, file := range astPkg.Files {
		fileExamples := doc.Examples(file)
		examples = append(examples, fileExamples...)
	}

	return examples
}

// GetPackagesByCategory categorizes packages for easier documentation generation
func (d *Discoverer) GetPackagesByCategory(packages []*PackageInfo) map[string][]*PackageInfo {
	categories := make(map[string][]*PackageInfo)

	for _, pkg := range packages {
		category := d.determinePackageCategory(pkg)
		categories[category] = append(categories[category], pkg)
	}

	return categories
}

// determinePackageCategory determines the category of a package for organization
func (d *Discoverer) determinePackageCategory(pkg *PackageInfo) string {
	// Determine category based on package path and name
	if strings.Contains(pkg.Path, "/internal/") {
		return "internal"
	}
	if strings.Contains(pkg.Path, "/cmd/") {
		return "commands"
	}
	if strings.HasSuffix(pkg.Name, "_test") {
		return "test"
	}
	if pkg.Name == "main" {
		return "main"
	}

	// Default to public API
	return "api"
}

// enhanceFunction extracts detailed parameter and return information from a function
func (d *Discoverer) enhanceFunction(fn *doc.Func, astPkg *ast.Package) *EnhancedFunc {
	enhanced := &EnhancedFunc{
		Func:        fn,
		Params:      []*Parameter{},
		Results:     []*Result{},
		ExampleCode: d.generateExampleCode(fn),
		Declaration: "",
		Doc:         "",
	}

	// Find the function declaration in the AST
	var funcDecl *ast.FuncDecl
	for _, file := range astPkg.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if f, ok := n.(*ast.FuncDecl); ok && f.Name.Name == fn.Name {
				funcDecl = f
				return false
			}
			return true
		})
		if funcDecl != nil {
			break
		}
	}

	// Use the original documentation from doc.Func for parameter/return parsing
	fullDoc := fn.Doc

	// For the main function description, show only the part before Parameters/Returns sections
	enhanced.Doc = d.extractMainDescription(fullDoc)

	if funcDecl == nil {
		return enhanced
	}

	// Generate clean function declaration
	enhanced.Declaration = d.generateFunctionDeclaration(fn, funcDecl)

	// Extract parameters
	if funcDecl.Type.Params != nil {
		for _, param := range funcDecl.Type.Params.List {
			paramType := d.formatType(param.Type)
			if len(param.Names) > 0 {
				for _, name := range param.Names {
					enhanced.Params = append(enhanced.Params, &Parameter{
						Name: name.Name,
						Type: paramType,
						Doc:  d.extractParamDoc(fullDoc, name.Name),
					})
				}
			} else {
				enhanced.Params = append(enhanced.Params, &Parameter{
					Name: "",
					Type: paramType,
					Doc:  "",
				})
			}
		}
	}

	// Extract return values
	if funcDecl.Type.Results != nil {
		for i, result := range funcDecl.Type.Results.List {
			resultType := d.formatType(result.Type)
			resultName := ""
			if len(result.Names) > 0 {
				resultName = result.Names[0].Name
			}
			enhanced.Results = append(enhanced.Results, &Result{
				Name: resultName,
				Type: resultType,
				Doc:  d.extractReturnDoc(fullDoc, i),
			})
		}
	}

	return enhanced
}

// enhanceType extracts detailed field and method information from a type
func (d *Discoverer) enhanceType(typ *doc.Type, astPkg *ast.Package) *EnhancedType {
	enhanced := &EnhancedType{
		Type:        typ,
		Fields:      []*Field{},
		Methods:     []*EnhancedFunc{},
		Funcs:       []*EnhancedFunc{},
		TypeKind:    "unknown",
		Declaration: "",
		Doc:         typ.Doc, // Start with doc.Type.Doc
		ExampleCode: "",
	}

	// Use custom AST traversal to extract type documentation
	enhanced.Doc = d.extractTypeDocumentation(typ.Name, astPkg)

	// Find the type declaration in the AST
	var typeSpec *ast.TypeSpec
	for _, file := range astPkg.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if ts, ok := n.(*ast.TypeSpec); ok && ts.Name.Name == typ.Name {
				typeSpec = ts
				return false
			}
			return true
		})
		if typeSpec != nil {
			break
		}
	}

	if typeSpec == nil {
		return enhanced
	}

	// Determine type kind and generate clean declaration
	switch t := typeSpec.Type.(type) {
	case *ast.StructType:
		enhanced.TypeKind = "struct"
		enhanced.Declaration = d.generateStructDeclaration(typ.Name, t)
		enhanced.ExampleCode = d.generateTypeExample(typ, typeSpec)
		// Extract struct documentation from AST if doc.Type.Doc is empty
		if enhanced.Doc == "" && typeSpec.Doc != nil {
			enhanced.Doc = d.extractTypeDoc(typeSpec.Doc)
		}
	case *ast.InterfaceType:
		enhanced.TypeKind = "interface"
		enhanced.Declaration = d.generateInterfaceDeclaration(typ.Name, t)
		enhanced.ExampleCode = d.generateTypeExample(typ, typeSpec)
		// Extract interface documentation from AST if doc.Type.Doc is empty
		if enhanced.Doc == "" && typeSpec.Doc != nil {
			enhanced.Doc = d.extractTypeDoc(typeSpec.Doc)
		}
	default:
		enhanced.TypeKind = "type"
		enhanced.Declaration = fmt.Sprintf("type %s %s", typ.Name, d.formatType(typeSpec.Type))
		enhanced.ExampleCode = d.generateTypeExample(typ, typeSpec)
		// Extract type documentation from AST if doc.Type.Doc is empty
		if enhanced.Doc == "" && typeSpec.Doc != nil {
			enhanced.Doc = d.extractTypeDoc(typeSpec.Doc)
		}
	}

	// Extract fields for struct types
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		for _, field := range structType.Fields.List {
			fieldType := d.formatType(field.Type)
			fieldTag := ""
			if field.Tag != nil {
				fieldTag = field.Tag.Value
			}

			if len(field.Names) > 0 {
				for _, name := range field.Names {
					enhanced.Fields = append(enhanced.Fields, &Field{
						Name: name.Name,
						Type: fieldType,
						Tag:  fieldTag,
						Doc:  d.extractFieldDoc(field),
					})
				}
			} else {
				// Embedded field
				enhanced.Fields = append(enhanced.Fields, &Field{
					Name: "",
					Type: fieldType,
					Tag:  fieldTag,
					Doc:  d.extractFieldDoc(field),
				})
			}
		}
	}

	// Enhance methods
	for _, method := range typ.Methods {
		enhancedMethod := d.enhanceFunction(method, astPkg)
		// For interface methods, try to extract documentation from the interface doc
		if enhanced.TypeKind == "interface" && enhancedMethod.Doc == "" {
			enhancedMethod.Doc = d.extractInterfaceMethodDoc(typ, method.Name)
		}
		enhanced.Methods = append(enhanced.Methods, enhancedMethod)
	}

	// Enhance constructor functions
	for _, fn := range typ.Funcs {
		enhanced.Funcs = append(enhanced.Funcs, d.enhanceFunction(fn, astPkg))
	}

	return enhanced
}

// formatType converts an AST type expression to a string
func (d *Discoverer) formatType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + d.formatType(t.X)
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + d.formatType(t.Elt)
		}
		return fmt.Sprintf("[%s]%s", d.formatType(t.Len), d.formatType(t.Elt))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", d.formatType(t.Key), d.formatType(t.Value))
	case *ast.ChanType:
		prefix := "chan"
		if t.Dir == ast.SEND {
			prefix = "chan<-"
		} else if t.Dir == ast.RECV {
			prefix = "<-chan"
		}
		return prefix + " " + d.formatType(t.Value)
	case *ast.FuncType:
		return "func" + d.formatFuncSignature(t)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.SelectorExpr:
		return d.formatType(t.X) + "." + t.Sel.Name
	case *ast.Ellipsis:
		return "..." + d.formatType(t.Elt)
	default:
		return fmt.Sprintf("%T", expr)
	}
}

// formatFuncSignature formats a function signature
func (d *Discoverer) formatFuncSignature(funcType *ast.FuncType) string {
	var result strings.Builder
	result.WriteString("(")

	// Parameters
	if funcType.Params != nil {
		for i, param := range funcType.Params.List {
			if i > 0 {
				result.WriteString(", ")
			}
			if len(param.Names) > 0 {
				for j, name := range param.Names {
					if j > 0 {
						result.WriteString(", ")
					}
					result.WriteString(name.Name)
				}
				result.WriteString(" ")
			}
			result.WriteString(d.formatType(param.Type))
		}
	}

	result.WriteString(")")

	// Results
	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		result.WriteString(" ")
		if len(funcType.Results.List) > 1 {
			result.WriteString("(")
		}
		for i, resultParam := range funcType.Results.List {
			if i > 0 {
				result.WriteString(", ")
			}
			if len(resultParam.Names) > 0 {
				for j, name := range resultParam.Names {
					if j > 0 {
						result.WriteString(", ")
					}
					result.WriteString(name.Name)
					result.WriteString(" ")
				}
			}
			result.WriteString(d.formatType(resultParam.Type))
		}
		if len(funcType.Results.List) > 1 {
			result.WriteString(")")
		}
	}

	return result.String()
}

// extractParamDoc extracts parameter documentation from function documentation
func (d *Discoverer) extractParamDoc(doc, paramName string) string {
	if doc == "" {
		return ""
	}

	// Look for parameter documentation in various formats
	lines := strings.Split(doc, "\n")
	paramSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lower := strings.ToLower(line)

		// Check for parameter section markers
		if strings.HasPrefix(lower, "parameters:") {
			paramSection = true
			continue
		}

		// If we're in parameter section, look for - paramName: format
		if paramSection && strings.HasPrefix(line, "- ") && strings.Contains(line, paramName+":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}

		// Also check for direct paramName: format (without section)
		if strings.HasPrefix(line, paramName+":") {
			return strings.TrimSpace(strings.TrimPrefix(line, paramName+":"))
		}
		// Check for - paramName: format (Go documentation style)
		if strings.HasPrefix(line, "- "+paramName+":") {
			return strings.TrimSpace(strings.TrimPrefix(line, "- "+paramName+":"))
		}
	}
	return ""
}

// extractReturnDoc extracts return value documentation from function documentation
func (d *Discoverer) extractReturnDoc(doc string, index int) string {
	if doc == "" {
		return ""
	}

	// Look for return documentation in various formats
	lines := strings.Split(doc, "\n")
	returnSection := false
	returnCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lower := strings.ToLower(line)

		// Check for return section markers
		if strings.HasPrefix(lower, "returns:") {
			returnSection = true
			continue
		}

		// If we're in return section, look for - Type: description format
		if returnSection && strings.HasPrefix(line, "- ") && strings.Contains(line, ":") {
			if returnCount == index {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1])
				}
			}
			returnCount++
		}
	}
	return ""
}

// extractFieldDoc extracts field documentation from struct field
func (d *Discoverer) extractFieldDoc(field *ast.Field) string {
	if field.Doc != nil {
		var docs []string
		for _, comment := range field.Doc.List {
			// Remove // prefix and clean up
			text := strings.TrimPrefix(comment.Text, "//")
			text = strings.TrimSpace(text)
			if text != "" {
				docs = append(docs, text)
			}
		}
		return strings.TrimSpace(strings.Join(docs, " "))
	}
	if field.Comment != nil {
		var docs []string
		for _, comment := range field.Comment.List {
			// Remove // prefix and clean up
			text := strings.TrimPrefix(comment.Text, "//")
			text = strings.TrimSpace(text)
			if text != "" {
				docs = append(docs, text)
			}
		}
		return strings.TrimSpace(strings.Join(docs, " "))
	}
	return ""
}

// generateExampleCode generates basic example code for a function
func (d *Discoverer) generateExampleCode(fn *doc.Func) string {
	var example strings.Builder

	// Generate a simple function call example
	example.WriteString(fmt.Sprintf("result := %s(", fn.Name))

	// Add placeholder parameters
	example.WriteString("/* parameters */")

	example.WriteString(")")

	return example.String()
}

// generateTypeExample generates a usage example for a type
func (d *Discoverer) generateTypeExample(typ *doc.Type, typeSpec *ast.TypeSpec) string {
	var example strings.Builder

	switch t := typeSpec.Type.(type) {
	case *ast.StructType:
		// Generate a struct instantiation example
		example.WriteString(fmt.Sprintf("// Create a new %s\n", typ.Name))
		example.WriteString(fmt.Sprintf("%s := %s{\n", strings.ToLower(typ.Name), typ.Name))

		// Add example field values
		for i, field := range t.Fields.List {
			if len(field.Names) > 0 {
				fieldName := field.Names[0].Name
				exampleValue := d.generateExampleValue(field.Type)
				example.WriteString(fmt.Sprintf("    %s: %s,", fieldName, exampleValue))
				if i < len(t.Fields.List)-1 {
					example.WriteString("\n")
				}
			}
		}
		example.WriteString("\n}")

	case *ast.InterfaceType:
		// For interfaces, show how to implement them
		example.WriteString(fmt.Sprintf("// Example implementation of %s\n", typ.Name))
		example.WriteString(fmt.Sprintf("type My%s struct {\n", typ.Name))
		example.WriteString("    // Add your fields here\n")
		example.WriteString("}\n\n")

		// Add method stubs
		for _, method := range t.Methods.List {
			if len(method.Names) > 0 {
				methodName := method.Names[0].Name
				if funcType, ok := method.Type.(*ast.FuncType); ok {
					example.WriteString(fmt.Sprintf("func (m My%s) %s(", typ.Name, methodName))
					// Add parameters
					if funcType.Params != nil {
						for i, param := range funcType.Params.List {
							if i > 0 {
								example.WriteString(", ")
							}
							paramType := d.formatType(param.Type)
							example.WriteString(fmt.Sprintf("param%d %s", i+1, paramType))
						}
					}
					example.WriteString(") ")
					// Add return type
					if funcType.Results != nil && len(funcType.Results.List) > 0 {
						example.WriteString(d.formatType(funcType.Results.List[0].Type))
					}
					example.WriteString(" {\n")
					example.WriteString("    // Implement your logic here\n")
					example.WriteString("    return\n")
					example.WriteString("}\n\n")
				}
			}
		}

	default:
		// For other types, show basic usage
		example.WriteString(fmt.Sprintf("// Example usage of %s\n", typ.Name))
		example.WriteString(fmt.Sprintf("var value %s\n", typ.Name))
		example.WriteString("// Initialize with appropriate value")
	}

	return example.String()
}

// generateExampleValue generates an example value for a type
func (d *Discoverer) generateExampleValue(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "string":
			return `"example"`
		case "int", "int8", "int16", "int32", "int64":
			return "42"
		case "uint", "uint8", "uint16", "uint32", "uint64":
			return "42"
		case "float32", "float64":
			return "3.14"
		case "bool":
			return "true"
		default:
			return fmt.Sprintf("%s{}", t.Name)
		}
	case *ast.StarExpr:
		return fmt.Sprintf("&%s{}", d.generateExampleValue(t.X))
	case *ast.ArrayType:
		return "[]"
	case *ast.MapType:
		return "map[]"
	default:
		return "/* value */"
	}
}

// generateFunctionDeclaration creates a clean function declaration
func (d *Discoverer) generateFunctionDeclaration(fn *doc.Func, funcDecl *ast.FuncDecl) string {
	var result strings.Builder

	// Function name
	if funcDecl.Recv != nil {
		// Method
		receiver := d.formatType(funcDecl.Recv.List[0].Type)
		result.WriteString(fmt.Sprintf("func (%s) %s", receiver, fn.Name))
	} else {
		// Function
		result.WriteString(fmt.Sprintf("func %s", fn.Name))
	}

	// Parameters
	result.WriteString("(")
	if funcDecl.Type.Params != nil {
		for i, param := range funcDecl.Type.Params.List {
			if i > 0 {
				result.WriteString(", ")
			}
			if len(param.Names) > 0 {
				for j, name := range param.Names {
					if j > 0 {
						result.WriteString(", ")
					}
					result.WriteString(name.Name)
				}
				result.WriteString(" ")
			}
			result.WriteString(d.formatType(param.Type))
		}
	}
	result.WriteString(")")

	// Return values
	if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) > 0 {
		result.WriteString(" ")
		if len(funcDecl.Type.Results.List) > 1 {
			result.WriteString("(")
		}
		for i, resultParam := range funcDecl.Type.Results.List {
			if i > 0 {
				result.WriteString(", ")
			}
			if len(resultParam.Names) > 0 {
				for j, name := range resultParam.Names {
					if j > 0 {
						result.WriteString(", ")
					}
					result.WriteString(name.Name)
					result.WriteString(" ")
				}
			}
			result.WriteString(d.formatType(resultParam.Type))
		}
		if len(funcDecl.Type.Results.List) > 1 {
			result.WriteString(")")
		}
	}

	return result.String()
}

// generateStructDeclaration creates a clean struct declaration
func (d *Discoverer) generateStructDeclaration(name string, structType *ast.StructType) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("type %s struct {\n", name))

	for _, field := range structType.Fields.List {
		fieldType := d.formatType(field.Type)
		fieldTag := ""
		if field.Tag != nil {
			fieldTag = " " + field.Tag.Value
		}

		if len(field.Names) > 0 {
			for _, fieldName := range field.Names {
				result.WriteString(fmt.Sprintf("    %s %s%s\n", fieldName.Name, fieldType, fieldTag))
			}
		} else {
			// Embedded field
			result.WriteString(fmt.Sprintf("    %s%s\n", fieldType, fieldTag))
		}
	}

	result.WriteString("}")
	return result.String()
}

// generateInterfaceDeclaration creates a clean interface declaration
func (d *Discoverer) generateInterfaceDeclaration(name string, interfaceType *ast.InterfaceType) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("type %s interface {\n", name))

	for _, method := range interfaceType.Methods.List {
		if len(method.Names) > 0 {
			methodName := method.Names[0].Name
			if funcType, ok := method.Type.(*ast.FuncType); ok {
				signature := d.formatFuncSignature(funcType)
				result.WriteString(fmt.Sprintf("    %s%s\n", methodName, signature))
			}
		} else {
			// Embedded interface
			embeddedType := d.formatType(method.Type)
			result.WriteString(fmt.Sprintf("    %s\n", embeddedType))
		}
	}

	result.WriteString("}")
	return result.String()
}

// extractInterfaceMethodDoc extracts documentation for interface methods
func (d *Discoverer) extractInterfaceMethodDoc(typ *doc.Type, methodName string) string {
	if typ.Doc == "" {
		return ""
	}

	// Look for method documentation in the interface doc
	lines := strings.Split(typ.Doc, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		// Look for lines that mention the method name
		if strings.Contains(line, methodName) {
			// Extract the description after the method name
			if strings.Contains(line, " ") {
				parts := strings.SplitN(line, " ", 2)
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1])
				}
			}
			// If no description on same line, look at next line
			if i+1 < len(lines) {
				nextLine := strings.TrimSpace(lines[i+1])
				if nextLine != "" && !strings.Contains(nextLine, " ") {
					return nextLine
				}
			}
		}
	}
	return ""
}

// extractTypeDocumentation performs custom AST traversal to extract type documentation
func (d *Discoverer) extractTypeDocumentation(typeName string, astPkg *ast.Package) string {
	var documentation []string
	found := false

	// Traverse all files in the package
	for _, file := range astPkg.Files {
		// Try to find the type and its documentation
		ast.Inspect(file, func(n ast.Node) bool {
			// Look for type declarations
			if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				// Check if this declaration contains our target type
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.Name == typeName {
						found = true

						// Extract documentation from the declaration
						if genDecl.Doc != nil {
							for _, comment := range genDecl.Doc.List {
								text := strings.TrimPrefix(comment.Text, "//")
								text = strings.TrimSpace(text)
								if text != "" {
									documentation = append(documentation, text)
								}
							}
						}

						// Also check for comments directly on the type spec
						if typeSpec.Doc != nil {
							for _, comment := range typeSpec.Doc.List {
								text := strings.TrimPrefix(comment.Text, "//")
								text = strings.TrimSpace(text)
								if text != "" {
									documentation = append(documentation, text)
								}
							}
						}

						// Also check for comments after the type spec
						if typeSpec.Comment != nil {
							for _, comment := range typeSpec.Comment.List {
								text := strings.TrimPrefix(comment.Text, "//")
								text = strings.TrimSpace(text)
								if text != "" {
									documentation = append(documentation, text)
								}
							}
						}

						// If we still don't have documentation, try to find comments manually
						if len(documentation) == 0 {
							manualDoc := d.extractManualComments(file, typeSpec.Pos())
							if manualDoc != "" {
								documentation = append(documentation, manualDoc)
							}
						}

						return false // Found our type, stop traversal
					}
				}
			}
			return true
		})

		// If we found documentation, no need to check other files
		if len(documentation) > 0 {
			break
		}
	}

	if !found {
		fmt.Printf("WARNING: Type %s not found in AST\n", typeName)
	}

	return strings.TrimSpace(strings.Join(documentation, " "))
}

// extractManualComments tries to find comments manually by looking at the file content
func (d *Discoverer) extractManualComments(file *ast.File, pos token.Pos) string {
	// Get the file content
	fileSet := d.fileSet
	filePos := fileSet.Position(pos)

	// Read the file content
	content, err := os.ReadFile(filePos.Filename)
	if err != nil {
		return ""
	}

	// Convert content to string
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	// Find the line number for the position
	lineNum := filePos.Line - 1 // Convert to 0-based index

	// Look for comments in the lines before the type declaration
	var comments []string
	for i := lineNum - 1; i >= 0 && i >= lineNum-10; i-- { // Look back up to 10 lines
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimSpace(strings.TrimPrefix(line, "//"))
			if comment != "" {
				comments = append([]string{comment}, comments...) // Prepend to maintain order
			}
		} else if strings.HasPrefix(line, "/*") {
			// Handle block comments
			comment := strings.TrimSpace(strings.TrimPrefix(line, "/*"))
			comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
			if comment != "" {
				comments = append([]string{comment}, comments...)
			}
		} else {
			// Stop when we hit non-comment, non-empty line
			break
		}
	}

	return strings.Join(comments, " ")
}

// extractFunctionDocumentation performs custom AST traversal to extract function documentation
func (d *Discoverer) extractFunctionDocumentation(funcName string, astPkg *ast.Package) string {
	var documentation []string
	found := false

	// Traverse all files in the package
	for _, file := range astPkg.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Look for function declarations
			if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Name.Name == funcName {
				found = true
				// Extract documentation from the function declaration
				if funcDecl.Doc != nil {
					for _, comment := range funcDecl.Doc.List {
						text := strings.TrimPrefix(comment.Text, "//")
						text = strings.TrimSpace(text)
						if text != "" {
							documentation = append(documentation, text)
						}
					}
				}
				return false // Found our function, stop traversal
			}
			return true
		})

		// If we found documentation, no need to check other files
		if len(documentation) > 0 {
			break
		}
	}

	if !found {
		fmt.Printf("WARNING: Function %s not found in AST\n", funcName)
	}

	return strings.TrimSpace(strings.Join(documentation, " "))
}

// extractMainDescription extracts the main description part before Parameters/Returns sections
func (d *Discoverer) extractMainDescription(doc string) string {
	if doc == "" {
		return ""
	}

	lines := strings.Split(doc, "\n")
	var mainLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lower := strings.ToLower(line)

		// Stop when we hit Parameters: or Returns: sections
		if strings.HasPrefix(lower, "parameters:") || strings.HasPrefix(lower, "returns:") {
			break
		}

		// Include this line in the main description
		if line != "" {
			mainLines = append(mainLines, line)
		}
	}

	return strings.TrimSpace(strings.Join(mainLines, " "))
}

// extractTypeDoc extracts documentation from AST comments for types
func (d *Discoverer) extractTypeDoc(commentGroup *ast.CommentGroup) string {
	if commentGroup == nil {
		return ""
	}

	var docs []string
	for _, comment := range commentGroup.List {
		// Remove // prefix and clean up
		text := strings.TrimPrefix(comment.Text, "//")
		text = strings.TrimSpace(text)
		if text != "" {
			docs = append(docs, text)
		}
	}
	return strings.TrimSpace(strings.Join(docs, " "))
}
