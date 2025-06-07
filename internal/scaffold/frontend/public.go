package frontend

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sersi-project/sersi/pkg"
)

//go:embed templates/base/public/**
var publicFolder embed.FS

func AddPublicFolder(projectName string) error {
	dst := filepath.Join(projectName, "public")

	err := pkg.CreateDirectory(dst)
	if err != nil {
		return fmt.Errorf("failed to create public folder: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	projectPath := filepath.Join(cwd, projectName, "public")

	err = pkg.CopyDirectory(publicFolder, "templates/base/public", projectPath)
	if err != nil {
		return fmt.Errorf("failed to copy public folder: %w", err)
	}

	return nil
}
