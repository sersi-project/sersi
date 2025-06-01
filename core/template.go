package core

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	gotemplate "text/template"
)

type TemplateMap struct {
	TemplatePath string
	outputFile   string
}

//go:embed templates/**/*
var templatesFolder embed.FS

var jsFiles = []TemplateMap{
    {"templates/base/vite.config.tmpl", "vite.config.js"},
}

var tsFiles = []TemplateMap{
	{"templates/base/vite.config.tmpl", "vite.config.ts"},
	{"templates/base/ts/tsconfig.json.tmpl", "tsconfig.json"},
	{"templates/base/ts/tsconfig.node.tmpl", "tsconfig.node.json"},
	{"templates/base/ts/tsconfig.app.json.tmpl", "tsconfig.app.json"},
}

var baseFiles = []TemplateMap{
	{"templates/base/index.html.tmpl", "index.html"},
	{"templates/base/package.json.tmpl", "package.json"},
	{"templates/base/README.md.tmpl", "README.md"},
	{"templates/base/styles.css.tmpl", "src/styles.css"},
	{"templates/base/.gitignore.tmpl", ".gitignore"},	
	{"templates/base/App.css.tmpl", "src/App.css"},
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
	err := t.renderFiles(baseFiles)
	if err != nil {
		return fmt.Errorf("failed to render base files: %w", err)
	}

    if t.Language == "ts" {
        err := t.renderFiles(tsFiles)
        if err != nil {
            return fmt.Errorf("failed to render ts files: %w", err)
        }
    } else {
        err := t.renderFiles(jsFiles)
        if err != nil {
            return fmt.Errorf("failed to render js files: %w", err)
        }
    }

    err = t.renderMainFile()
    if err != nil {
        return fmt.Errorf("failed to render main file: %w", err)
    }

    err = t.renderAppFile()
    if err != nil {
        return fmt.Errorf("failed to render app file: %w", err)
    }

	return nil
}

func (t *Template) renderAppFile() error {
    var err error
    switch t.Framework {
    case "react":
        if t.Language == "ts" { 
            err = t.renderTemplate("templates/base/App.tmpl", "src/App.tsx")
            if err != nil {
                return err
            }
        } else {
            err := t.renderTemplate("templates/base/App.tmpl", "src/App.jsx")
            if err != nil {
                return err
            }
        }
    case "vue":
        err = t.renderTemplate("templates/base/App.tmpl", "src/App.vue")
        if err != nil {
            return err
        }
    case "svelte":
        err := t.renderTemplate("templates/base/App.tmpl", "src/App.svelte")
        if err != nil {
            return err
        }
    }

    return nil
}

func (t *Template) renderFiles(files []TemplateMap) error {
	for _, file := range files {
		err := t.renderTemplate(file.TemplatePath, file.outputFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Template) renderMainFile() error {
    if t.Framework == "react" {
        if t.Language == "ts" {
            err := t.renderTemplate("templates/base/main.tmpl", "src/main.tsx")
            if err != nil {
                return err
            }
        } else {
            err := t.renderTemplate("templates/base/main.tmpl", "src/main.jsx")
            if err != nil {
                return err
            }
        }
    } else {
        if t.Language == "ts" {
            err := t.renderTemplate("templates/base/main.tmpl", "src/main.ts")
            if err != nil {
                return err
            }
        } else {
            err := t.renderTemplate("templates/base/main.tmpl", "src/main.js")
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func (t *Template) renderTemplate(templateFile, outputFile string) error {
	outputPath, err := t.createOutputFilePath(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file path: %w", err)
	}

	tmpl, err := gotemplate.ParseFS(templatesFolder, templateFile)
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
