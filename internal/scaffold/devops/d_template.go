package devops

import (
	"embed"
	"fmt"

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

func (b *DTemplateBuilder) Container(container string) *DTemplateBuilder {
	b.dTemplate.Container = container
	return b
}

func (b *DTemplateBuilder) Language(language string) *DTemplateBuilder {
	b.dTemplate.Language = language
	return b
}

func (b *DTemplateBuilder) Build() *DTemplate {
	return b.dTemplate
}

type DTemplate struct {
	ProjectName string
	CI          string
	Container   string
	Language    string
}

func (t *DTemplate) Execute() error {
	switch t.CI {
	case "github-actions":
		return t.ExecuteGithub()
	case "gitlab-ci":
		return t.ExecuteGitlab()
	case "circleci":
		return t.ExecuteCircleci()
	case "bitbucket-pipelines":
		return t.ExecuteBitbucket()
	default:
		return t.ProcessError(fmt.Errorf("invalid CI provider"))
	}
}

func (t *DTemplate) ProcessError(err error) error {
	return err
}

func (t *DTemplate) ExecuteGithub() error {
	err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/ci.tmpl", ".github/workflows/ci.yml", t.ProjectName)
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

func (t *DTemplate) ExecuteBitbucket() error {
	err := scaffold.RenderTemplate(t, templatesFolder, "templates/base/ci.tmpl", "bitbucket-pipelines.yml", t.ProjectName)
	if err != nil {
		return err
	}

	return nil
}
