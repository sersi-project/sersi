package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

func ScaffoldProject(name, framework, css, lang string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectPath := filepath.Join(cwd, name)

	err = os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	data := fmt.Sprintf("# Project Options: %s - %s - %s", framework, lang, css)

	err = os.WriteFile(filepath.Join(name, "README.md"), []byte(data), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Project Created at %s\n", projectPath)
	return nil
}
