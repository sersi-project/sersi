package frontend

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sersi-project/core/pkg"
)

//go:embed templates/golden/react
var baseReactPath embed.FS

//go:embed templates/golden/vue
var baseVuePath embed.FS

//go:embed templates/golden/svelte
var baseSveltePath embed.FS

//go:embed templates/golden/react-ts
var baseReactTsPath embed.FS

//go:embed templates/golden/vanilla
var baseVanillaPath embed.FS

//go:embed templates/golden/vanilla-ts
var baseVanillaTsPath embed.FS

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

	var projectPath embed.FS

	switch g.Framework {
	case "react":
		projectPath = baseReactPath
	case "vue":
		projectPath = baseVuePath
	case "svelte":
		projectPath = baseSveltePath
	case "react-ts":
		projectPath = baseReactTsPath
	case "vanilla":
		projectPath = baseVanillaPath
	case "vanilla-ts":
		projectPath = baseVanillaTsPath
	default:
		return fmt.Errorf("invalid framework: %s", g.Framework)
	}

	dst := filepath.Join(cwd, g.ProjectName)

	err = pkg.CopyDirectory(projectPath, "templates/golden/"+g.Framework, dst)
	if err != nil {
		return err
	}

	return nil
}
