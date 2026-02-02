package devops

import (
	"path/filepath"

	"github.com/sersi-project/sersi/pkg"
)

type Devops struct {
	ProjectName  string
	Container    string
	CI           string
	outputFolder string
}

type DevopsBuilder struct {
	devops *Devops
}

func NewDevopsBuilder() *DevopsBuilder {
	return &DevopsBuilder{
		devops: &Devops{},
	}
}

func (b *DevopsBuilder) ProjectName(projectName string) *DevopsBuilder {
	b.devops.ProjectName = projectName
	return b
}

func (b *DevopsBuilder) Container(container string) *DevopsBuilder {
	b.devops.Container = container
	return b
}

func (b *DevopsBuilder) CI(ci string) *DevopsBuilder {
	b.devops.CI = ci
	return b
}

func (b *DevopsBuilder) Build() *Devops {
	return b.devops
}

func (b *Devops) Generate(language string) error {
	b.SetOutputFolder()
	if b.outputFolder != "" {
		projectPath := filepath.Join(b.ProjectName, b.outputFolder)
		err := pkg.CreateDirectory(projectPath)
		if err != nil {
			return b.ProcessError(err)
		}
	}

	template := NewDTemplateBuilder().ProjectName(b.ProjectName).CI(b.CI).Container(b.Container).Language(language).Build()
	err := template.Execute()
	if err != nil {
		return b.ProcessError(err)
	}
	return nil
}

func (b *Devops) SetOutputFolder() {
	switch b.CI {
	case "github-actions":
		b.outputFolder = ".github/workflows"
	case "circleci":
		b.outputFolder = ".circleci"
	default:
		b.outputFolder = ""
	}
}

func (b *Devops) ProcessError(err error) error {
	return err
}
