package frontend

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sersi-project/core/pkg"
)

type Frontend struct {
	ProjectName string
	Framework   string
	CSS         string
	Language    string
	Monorepo    bool
}

type FrontendBuilder struct {
	config *Frontend
}

func NewFrontendBuilder() *FrontendBuilder {
	return &FrontendBuilder{
		config: &Frontend{},
	}
}

func (fb *FrontendBuilder) ProjectName(name string) *FrontendBuilder {
	fb.config.ProjectName = name
	return fb
}

func (fb *FrontendBuilder) Framework(framework string) *FrontendBuilder {
	fb.config.Framework = strings.ToLower(framework)
	return fb
}

func (fb *FrontendBuilder) CSS(css string) *FrontendBuilder {
	fb.config.CSS = strings.ToLower(css)
	return fb
}

func (fb *FrontendBuilder) Language(lang string) *FrontendBuilder {
	fb.config.Language = strings.ToLower(lang)
	return fb
}

func (fb *FrontendBuilder) Monorepo(monorepo bool) *FrontendBuilder {
	fb.config.Monorepo = monorepo
	return fb
}

func (fb *FrontendBuilder) Build() *Frontend {
	return fb.config
}

func (f *Frontend) Generate() error {
	err := pkg.CreateDirectory(f.ProjectName)
	if err != nil {
		return f.ProcessError(err)
	}

	if f.Monorepo {
		err := pkg.CreateDirectory(filepath.Join(f.ProjectName, "frontend"))
		if err != nil {
			return f.ProcessError(err)
		}
		err = AddPublicFolder(filepath.Join(f.ProjectName, "frontend"))
		if err != nil {
			return f.ProcessError(err)
		}
	} else {
		err = AddPublicFolder(f.ProjectName)
		if err != nil {
			return f.ProcessError(err)
		}
	}

	gtFramework := f.Framework

	if f.Framework == "react" && f.Language == "typescript" {
		gtFramework = "react-ts"
	}

	if f.Framework == "vanilla" && f.Language == "typescript" {
		gtFramework = "vanilla-ts"
	}

	if f.Monorepo {
		goldenTemplate := NewGoldenArchitecture(filepath.Join(f.ProjectName, "frontend"), gtFramework)
		err = goldenTemplate.Generate()
		if err != nil {
			return f.ProcessError(err)
		}
	} else {
		goldenTemplate := NewGoldenArchitecture(f.ProjectName, gtFramework)
		err = goldenTemplate.Generate()
		if err != nil {
			return f.ProcessError(err)
		}
	}

	templateBuilder := NewFTemplateBuilder().
		ProjectName(f.ProjectName).
		Framework(f.Framework).
		CSS(f.CSS).
		Monorepo(f.Monorepo)
	if f.Language == "typescript" {
		templateBuilder.Language("ts")
	} else {
		templateBuilder.Language("js")
	}

	template := templateBuilder.Build()

	err = template.Execute()
	if err != nil {
		return f.ProcessError(err)
	}

	return nil
}

func (f *Frontend) ProcessError(err error) error {
	cleanupErr := pkg.CleanupDirs(f.ProjectName)
	if cleanupErr != nil {
		return fmt.Errorf("failed to cleanup project: %s", cleanupErr.Error())
	}
	return fmt.Errorf("failed to generate project: %s", err.Error())
}
