package backend

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/sersi-project/core/internal/scaffold"
)

type BTemplateMap struct {
	TemplatePath string
	outputFile   string
}

//go:embed templates/base/*
var templatesFolder embed.FS

var baseFiles = []BTemplateMap{
	{"templates/base/env.tmpl", ".env"},
	{"templates/base/gitignore.tmpl", ".gitignore"},
}

var jsFiles = []BTemplateMap{
	{"templates/base/main.tmpl", "index.js"},
	{"templates/base/routes.tmpl", "routes/routes.js"},
	{"templates/base/service.tmpl", "services/service.js"},
	{"templates/base/model.tmpl", "models/model.js"},
	{"templates/base/controller.tmpl", "controllers/controller.js"},
	{"templates/base/db.tmpl", "db/dummy_db.js"},
	{"templates/base/package.tmpl", "package.json"},
}

var tsFiles = []BTemplateMap{
	{"templates/base/main.tmpl", "index.ts"},
	{"templates/base/routes.tmpl", "routes/routes.ts"},
	{"templates/base/service.tmpl", "services/service.ts"},
	{"templates/base/model.tmpl", "models/model.ts"},
	{"templates/base/controller.tmpl", "controllers/controller.ts"},
	{"templates/base/db.tmpl", "db/dummy_db.ts"},
}

var goFiles = []BTemplateMap{
	{"templates/base/main.tmpl", "main.go"},
	{"templates/base/routes.tmpl", "routes/routes.go"},
	{"templates/base/service.tmpl", "services/service.go"},
	{"templates/base/model.tmpl", "models/model.go"},
	{"templates/base/controller.tmpl", "handlers/handler.go"},
	{"templates/base/db.tmpl", "db/dummy_db.go"},
}

var pyFiles = []BTemplateMap{
	{"templates/base/main.tmpl", "main.py"},
	{"templates/base/routes.tmpl", "routes/routes.py"},
	{"templates/base/service.tmpl", "services/service.py"},
	{"templates/base/model.tmpl", "models/model.py"},
	{"templates/base/controller.tmpl", "handlers/handler.py"},
	{"templates/base/db.tmpl", "db/dummy_db.py"},
}

type BTemplateBuilder struct {
	template *BTemplate
}

func NewBTemplateBuilder() *BTemplateBuilder {
	return &BTemplateBuilder{
		template: &BTemplate{},
	}
}

func (b *BTemplateBuilder) ProjectName(name string) *BTemplateBuilder {
	b.template.ProjectName = name
	return b
}

func (b *BTemplateBuilder) Framework(framework string) *BTemplateBuilder {
	b.template.Framework = framework
	return b
}

func (b *BTemplateBuilder) Language(language string) *BTemplateBuilder {
	b.template.Language = language
	return b
}

func (b *BTemplateBuilder) Database(database string) *BTemplateBuilder {
	b.template.Database = database
	return b
}

func (b *BTemplateBuilder) Monorepo(monorepo bool) *BTemplateBuilder {
	b.template.Monorepo = monorepo
	return b
}

func (b *BTemplateBuilder) Build() *BTemplate {
	return b.template
}

type BTemplate struct {
	ProjectName string
	Language    string
	Database    string
	Framework   string
	Monorepo    bool
}

func (t *BTemplate) Execute() error {
	switch t.Language {
	case "js":
		err := t.renderFiles(jsFiles)
		if err != nil {
			return fmt.Errorf("failed to render javascript files: %w", err)
		}
	case "ts":
		err := t.renderFiles(tsFiles)
		if err != nil {
			return fmt.Errorf("failed to render typescript files: %w", err)
		}
	case "go":
		err := t.renderFiles(goFiles)
		if err != nil {
			return fmt.Errorf("failed to render go files: %w", err)
		}
	case "py":
		err := t.renderFiles(pyFiles)
		if err != nil {
			return fmt.Errorf("failed to render python files: %w", err)
		}
	}

	err := t.renderFiles(baseFiles)
	if err != nil {
		return fmt.Errorf("failed to render base files: %w", err)
	}

	return nil
}

func (t *BTemplate) renderFiles(files []BTemplateMap) error {
	for _, file := range files {
		if t.Monorepo {
			file.outputFile = filepath.Join("backend", file.outputFile)
		}
		err := scaffold.RenderTemplate(t, templatesFolder, file.TemplatePath, file.outputFile, t.ProjectName)
		if err != nil {
			return fmt.Errorf("failed to render file: %w", err)
		}
	}
	return nil
}
