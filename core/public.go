package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sersi/utils"
)

const pathToPublic = "templates/base/public"

func AddPublicFolder(projectName string) error {
	dst := filepath.Join(projectName, "public")

	err := utils.CreateDirectory(dst)
	if err != nil {
		return fmt.Errorf("failed to create public folder: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	projectPath := filepath.Join(cwd, projectName, "public")

	err = utils.CopyDirectory(pathToPublic, projectPath)
	if err != nil {
		return fmt.Errorf("failed to copy public folder: %w", err)
	}

	fmt.Printf("✅ Public folder created at %s\n", projectPath)

	return nil
}
