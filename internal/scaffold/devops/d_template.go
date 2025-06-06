package devops

import (
	"embed"

	"github.com/sersi-project/sersi/internal/scaffold"
)

//go:embed templates/base/*
var templatesFolder embed.FS

type DTemplateBuilder struct {
	dTemplate *DTemplate
}

func NewDTemplateBuilder() *DTemplateBuilder {
	return &DTemplateBuilder{
		dTemplate: &DTemplate{},
	}
}

func (b *DTemplateBuilder) ProjectName(projectName string) *DTemplateBuilder {
	b.dTemplate.ProjectName = projectName
	return b
}

func (b *DTemplateBuilder) CI(ci string) *DTemplateBuilder {
	b.dTemplate.CI = ci
	return b
}

func (b *DTemplateBuilder) Docker(docker bool) *DTemplateBuilder {
	b.dTemplate.Docker = docker
	return b
}

func (b *DTemplateBuilder) Monitoring(monitoring string) *DTemplateBuilder {
	b.dTemplate.Monitoring = monitoring
	return b
}

func (b *DTemplateBuilder) Build() *DTemplate {
	return b.dTemplate
}

type DTemplate struct {
	ProjectName string
	CI          string
	Docker      bool
	Monitoring  string
}

func (t *DTemplate) Execute() error {
	switch t.CI {
	case "github":
		return t.ExecuteGithub()
	case "gitlab":
		return t.ExecuteGitlab()
	case "circleci":
		return t.ExecuteCircleci()
	}
	return nil
}

func (t *DTemplate) ProcessError(err error) error {
	return err
}

func (t *DTemplate) ExecuteGithub() error {
	err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/ci.tmpl", "/.github/workflows/ci.yml", t.ProjectName)
	if err != nil {
		return err
	}

	return nil
}

func (t *DTemplate) ExecuteGitlab() error {
	err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/ci.tmpl", ".gitlab-ci.yml", t.ProjectName)
	if err != nil {
		return err
	}

	return nil
}

func (t *DTemplate) ExecuteCircleci() error {
	err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/ci.tmpl", ".circleci/config.yml", t.ProjectName)
	if err != nil {
		return err
	}

	return nil
}
