package backend

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sersi-project/sersi/pkg"
)

var archJS = []string{
	"db",
	"controllers",
	"models",
	"routes",
	"services",
}

var archGo = []string{
	"db",
	"handlers",
	"models",
	"routes",
	"services",
}

var archPython = []string{
	"db",
	"models",
	"routes",
	"services",
}

type Backend struct {
	ProjectName string
	Language    string
	Framework   string
	Database    string
	Monorepo    bool
	Polyrepos   bool
}

type BackendBuilder struct {
	backend *Backend
}

func NewBackendBuilder() *BackendBuilder {
	return &BackendBuilder{
		backend: &Backend{},
	}
}

func (b *BackendBuilder) ProjectName(projectName string) *BackendBuilder {
	b.backend.ProjectName = projectName
	return b
}

func (b *BackendBuilder) Language(language string) *BackendBuilder {
	b.backend.Language = b.formatLanguage(language)
	return b
}

func (b *BackendBuilder) Framework(framework string) *BackendBuilder {
	b.backend.Framework = framework
	return b
}

func (b *BackendBuilder) Database(database string) *BackendBuilder {
	b.backend.Database = b.formatDatabase(database)
	return b
}

func (b *BackendBuilder) Monorepo(monorepo bool) *BackendBuilder {
	b.backend.Monorepo = monorepo
	return b
}

func (b *BackendBuilder) Polyrepos(polyrepos bool) *BackendBuilder {
	b.backend.Polyrepos = polyrepos
	return b
}

func (b *BackendBuilder) Build() *Backend {
	return b.backend
}

func (b *Backend) Generate() error {
	err := b.createFolders()
	if err != nil {
		return b.ProcessError(err)
	}
	return nil
}

func (b *Backend) ProcessError(err error) error {
	return err
}

func (b *Backend) createFolders() error {
	var projectPath string
	if b.Monorepo && !b.Polyrepos {
		projectPath = filepath.Join(b.ProjectName, "backend")
	} else if b.Polyrepos && !b.Monorepo {
		b.ProjectName = b.ProjectName + "-backend"
		projectPath = b.ProjectName
	} else if b.Monorepo && b.Polyrepos {
		return fmt.Errorf("invalid project structure")
	} else {
		projectPath = b.ProjectName
	}

	err := pkg.CreateDirectory(projectPath)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	if b.Language == "js" || b.Language == "ts" {
		for _, folder := range archJS {
			err := pkg.CreateDirectory(filepath.Join(projectPath, folder))
			if err != nil {
				return fmt.Errorf("failed to create app directory: %w", err)
			}
		}
	} else if b.Language == "go" {
		for _, folder := range archGo {
			err := pkg.CreateDirectory(filepath.Join(projectPath, folder))
			if err != nil {
				return fmt.Errorf("failed to create app directory: %w", err)
			}
		}
	} else if b.Language == "py" {
		for _, folder := range archPython {
			err := pkg.CreateDirectory(filepath.Join(projectPath, folder))
			if err != nil {
				return fmt.Errorf("failed to create app directory: %w", err)
			}
		}
	} else {
		return fmt.Errorf("unsupported language: %s", b.Language)
	}

	template := NewBTemplateBuilder().
		ProjectName(b.ProjectName).
		Language(b.Language).
		Framework(b.Framework).
		Database(b.Database).
		Monorepo(b.Monorepo).
		Build()

	err = template.Execute()
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func (b *BackendBuilder) formatLanguage(language string) string {
	if i := strings.Index(language, "("); i != -1 {
		language = language[:i]
	}
	language = strings.ToLower(language)

	if language == "javascript" || language == "node" || language == "js" {
		return "js"
	} else if language == "typescript" || language == "ts" {
		return "ts"
	} else if language == "python" || language == "py" {
		return "py"
	} else if language == "go" {
		return "go"
	} else {
		return language
	}
}

func (b *BackendBuilder) formatDatabase(database string) string {
	if i := strings.Index(database, "("); i != -1 {
		database = database[:i]
	}
	database = strings.ToLower(database)
	return database
}
