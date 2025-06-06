package scaffold

import (
	"embed"
	"fmt"
	"os"

	"github.com/sersi-project/core/pkg"

	gotemplate "text/template"
)

func RenderTemplate(t interface{}, templatesFolder embed.FS, templateFile, outputFile, projectName string) error {
	outputPath, err := pkg.CreateOutputFilePath(projectName, outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file path: %w", err)
	}

	tmpl, err := gotemplate.ParseFS(templatesFolder, templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %w", err)
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() {
		if cerr := output.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tmpl.Execute(output, t)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
