package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sersi/utils"
)

const baseGoldenTemplatePath = "templates/golden"

type GoldenTemplate struct {
	ProjectName string
	Framework   string
}

func NewGoldenArchitecture(projectName, framework string) *GoldenTemplate {
	return &GoldenTemplate{
		ProjectName: projectName,
		Framework:   framework,
	}
}

func (g *GoldenTemplate) Generate() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectPath := filepath.Join(baseGoldenTemplatePath, g.Framework)
	dst := filepath.Join(cwd, g.ProjectName)

	err = utils.CopyDirectory(projectPath, dst)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Base Framework Injected at %s\n", dst)
	return nil
}
