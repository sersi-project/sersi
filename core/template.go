package core

import (
	"fmt"
	"os"
	"path/filepath"
	gotemplate "text/template"
)

type TemplateMap struct {
	TemplatePath string
	outputFile   string
	isTs         bool
}

var filesToBeCreated = []TemplateMap{
	{"templates/base/index.html.tmpl", "index.html", false},
	{"templates/base/package.json.tmpl", "package.json", false},
	{"templates/base/README.md.tmpl", "README.md", false},
	{"templates/base/styles.css.tmpl", "src/styles.css", false},
	{"templates/base/main.js.tmpl", "src/main.js", false},
	{"templates/base/.gitignore.tmpl", ".gitignore", false},
	{"templates/base/vite.config.js.tmpl", "vite.config.js", false},
	{"templates/base/ts/main.ts.tmpl", "src/main.ts", true},
	{"templates/base/ts/vite.config.ts.tmpl", "vite.config.ts", true},
	// {"templates/base/ts/tsconfig.json.tmpl", "tsconfig.json", true},
	{"templates/base/ts/tsconfig.node.tmpl", "tsconfig.node.json", true},
	{"templates/base/ts/tsconfig.app.json.tmpl", "tsconfig.app.json", true},
}

type TemplateBuilder struct {
	template *Template
}

func NewTemplateBuilder() *TemplateBuilder {
	return &TemplateBuilder{
		template: &Template{},
	}
}

func (b *TemplateBuilder) ProjectName(name string) *TemplateBuilder {
	b.template.ProjectName = name
	return b
}

func (b *TemplateBuilder) Framework(framework string) *TemplateBuilder {
	b.template.Framework = framework
	return b
}

func (b *TemplateBuilder) CSS(css string) *TemplateBuilder {
	b.template.CSS = css
	return b
}

func (b *TemplateBuilder) Language(language string) *TemplateBuilder {
	b.template.Language = language
	return b
}

func (b *TemplateBuilder) Build() *Template {
	return b.template
}

type Template struct {
	ProjectName string
	Framework   string
	CSS         string
	Language    string
}

func (t *Template) Generate() error {
	for _, file := range filesToBeCreated {
		if t.Language == "ts" && file.outputFile == "src/main.js" {
			continue
		}

		if t.Framework == "react" && (file.outputFile == "src/main.js" || file.outputFile == "src/main.ts") {
			continue
		}

		if t.Language == "ts" && file.outputFile == "vite.config.js" {
			continue
		}

		if file.isTs && t.Language != "ts" {
			continue
		}

		err := t.renderTemplate(file.TemplatePath, file.outputFile)
		if err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}
	}
	return nil
}

func (t *Template) renderTemplate(templateFile, outputFile string) error {
	outputPath, err := t.createOutputFilePath(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file path: %w", err)
	}

	tmpl, err := gotemplate.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %w", err)
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() {
		if cerr := output.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tmpl.Execute(output, t)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func (t *Template) createOutputFilePath(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dst := filepath.Join(cwd, t.ProjectName, fileName)
	return dst, nil
}
