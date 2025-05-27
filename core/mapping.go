package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sersi-project/core/utils"

	yaml "gopkg.in/yaml.v2"
)

type FileParser struct {
	filePath string
}

type Config struct {
	Name     string         `yaml:"name"`
	Scaffold ScaffoldConfig `yaml:"scaffold"`
}

type ScaffoldConfig struct {
	Frontend FrontendConfig `yaml:"frontend"`
}

type FrontendConfig struct {
	Framework string `yaml:"framework"`
	CSS       string `yaml:"css"`
	Language  string `yaml:"lang"`
}

func NewFileParser(filePath string) *FileParser {
	return &FileParser{
		filePath: filePath,
	}
}

func (fp *FileParser) ExceuteMapping() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %v", err)
	}
	path := filepath.Join(cwd, "sersi.yaml")
	if fp.filePath != "" {
		path = filepath.Join(cwd, fp.filePath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	if config.Scaffold.Frontend.Language == "" {
		config.Scaffold.Frontend.Language = "javascript"
	}

	err = validateConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("error in config: %v", err)
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if utils.FileExists(config.Name) {
		return fmt.Errorf("project already exists")
	}

	if config.Name == "" {
		return fmt.Errorf("parameter: name is required")
	}

	if config.Scaffold.Frontend.Framework == "" {
		return fmt.Errorf("parameter: framework is required")
	}

	if config.Scaffold.Frontend.CSS == "" {
		return fmt.Errorf("parameter: css is required")
	}
	return nil
}
