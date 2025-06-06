package frontend

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/sersi-project/sersi/internal/scaffold"
)

type FTemplateMap struct {
	TemplatePath string
	outputFile   string
}

//go:embed templates/**/*
var templatesFolder embed.FS

var jsFiles = []FTemplateMap{
	{"templates/base/vite.config.tmpl", "vite.config.js"},
}

var tsFiles = []FTemplateMap{
	{"templates/base/vite.config.tmpl", "vite.config.ts"},
	{"templates/base/ts/tsconfig.json.tmpl", "tsconfig.json"},
	{"templates/base/ts/tsconfig.node.tmpl", "tsconfig.node.json"},
	{"templates/base/ts/tsconfig.app.json.tmpl", "tsconfig.app.json"},
}

var baseFiles = []FTemplateMap{
	{"templates/base/index.html.tmpl", "index.html"},
	{"templates/base/package.json.tmpl", "package.json"},
	{"templates/base/README.md.tmpl", "README.md"},
	{"templates/base/styles.css.tmpl", "src/styles.css"},
	{"templates/base/.gitignore.tmpl", ".gitignore"},
	{"templates/base/App.css.tmpl", "src/App.css"},
}

type FTemplateBuilder struct {
	template *FTemplate
}

func NewFTemplateBuilder() *FTemplateBuilder {
	return &FTemplateBuilder{
		template: &FTemplate{},
	}
}

func (b *FTemplateBuilder) ProjectName(name string) *FTemplateBuilder {
	b.template.ProjectName = name
	return b
}

func (b *FTemplateBuilder) Framework(framework string) *FTemplateBuilder {
	b.template.Framework = framework
	return b
}

func (b *FTemplateBuilder) CSS(css string) *FTemplateBuilder {
	b.template.CSS = css
	return b
}

func (b *FTemplateBuilder) Language(language string) *FTemplateBuilder {
	b.template.Language = language
	return b
}

func (b *FTemplateBuilder) Monorepo(monorepo bool) *FTemplateBuilder {
	b.template.Monorepo = monorepo
	return b
}

func (b *FTemplateBuilder) Polyrepos(polyrepos bool) *FTemplateBuilder {
	b.template.Polyrepos = polyrepos
	return b
}

func (b *FTemplateBuilder) Build() *FTemplate {
	return b.template
}

type FTemplate struct {
	ProjectName string
	Framework   string
	CSS         string
	Language    string
	Monorepo    bool
	Polyrepos   bool
}

func (t *FTemplate) Execute() error {
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

func (t *FTemplate) renderAppFile() error {
	var err error
	switch t.Framework {
	case "react":
		if t.Language == "ts" {
			if t.Monorepo {
				err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "frontend/src/App.tsx", t.ProjectName)
				if err != nil {
					return err
				}
			} else {
				err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "src/App.tsx", t.ProjectName)
				if err != nil {
					return err
				}
			}
		} else {
			if t.Monorepo {
				err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "frontend/src/App.jsx", t.ProjectName)
				if err != nil {
					return err
				}

			} else {
				err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "src/App.jsx", t.ProjectName)
				if err != nil {
					return err
				}
			}
		}
	case "vue":
		if t.Monorepo {
			err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "frontend/src/App.vue", t.ProjectName)
			if err != nil {
				return err
			}
		} else {
			err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "src/App.vue", t.ProjectName)
			if err != nil {
				return err
			}
		}
	case "svelte":
		if t.Monorepo {
			err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "frontend/src/App.svelte", t.ProjectName)
			if err != nil {
				return err
			}
		} else {
			err = scaffold.RenderTemplate(t, templatesFolder, "templates/base/App.tmpl", "src/App.svelte", t.ProjectName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *FTemplate) renderFiles(files []FTemplateMap) error {
	for _, file := range files {
		if t.Monorepo {
			file.outputFile = filepath.Join("frontend", file.outputFile)
		}
		err := scaffold.RenderTemplate(t, templatesFolder, file.TemplatePath, file.outputFile, t.ProjectName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *FTemplate) renderMainFile() error {
	if t.Framework == "react" {
		if t.Language == "ts" {
			if t.Monorepo {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "frontend/src/main.tsx", t.ProjectName)
				if err != nil {
					return err
				}
			} else {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "src/main.tsx", t.ProjectName)
				if err != nil {
					return err
				}
			}
		} else {
			if t.Monorepo {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "frontend/src/main.jsx", t.ProjectName)
				if err != nil {
					return err
				}
			} else {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "src/main.jsx", t.ProjectName)
				if err != nil {
					return err
				}
			}
		}
	} else {
		if t.Language == "ts" {
			if t.Monorepo {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "frontend/src/main.ts", t.ProjectName)
				if err != nil {
					return err
				}
			} else {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "src/main.ts", t.ProjectName)
				if err != nil {
					return err
				}
			}
		} else {
			if t.Monorepo {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "frontend/src/main.js", t.ProjectName)
				if err != nil {
					return err
				}
			} else {
				err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/main.tmpl", "src/main.js", t.ProjectName)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
