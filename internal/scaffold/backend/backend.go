package backend

import (
	"fmt"
	"path/filepath"

	"github.com/sersi-project/core/pkg"
)

var archJS = []string{
	"db",
	"controllers",
	"models",
	"routes",
	"services",
}

var archOther = []string{
	"db",
	"handlers",
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
	b.backend.Language = language
	return b
}

func (b *BackendBuilder) Framework(framework string) *BackendBuilder {
	b.backend.Framework = framework
	return b
}

func (b *BackendBuilder) Database(database string) *BackendBuilder {
	b.backend.Database = database
	return b
}

func (b *BackendBuilder) Monorepo(monorepo bool) *BackendBuilder {
	b.backend.Monorepo = monorepo
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
	if b.Monorepo {
		projectPath = filepath.Join(b.ProjectName, "backend")
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
	} else {
		for _, folder := range archOther {
			err := pkg.CreateDirectory(filepath.Join(projectPath, folder))
			if err != nil {
				return fmt.Errorf("failed to create app directory: %w", err)
			}
		}
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
