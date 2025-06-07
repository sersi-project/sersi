package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type FileParser struct {
	filePath string
}

type Config struct {
	Name      string         `yaml:"name"`
	Structure string         `yaml:"structure"`
	Scaffold  ScaffoldConfig `yaml:"scaffold"`
}

type ScaffoldConfig struct {
	Frontend FrontendConfig `yaml:"frontend"`
	Backend  BackendConfig  `yaml:"backend"`
	Devops   DevopsConfig   `yaml:"devops"`
}

type FrontendConfig struct {
	Framework string `yaml:"framework"`
	CSS       string `yaml:"css"`
	Language  string `yaml:"lang"`
}

type BackendConfig struct {
	Language  string `yaml:"lang"`
	Framework string `yaml:"framework"`
	Database  string `yaml:"database"`
}

type DevopsConfig struct {
	CI     string `yaml:"ci"`
	Docker bool   `yaml:"docker"`
}

func NewFileParser(filePath string) *FileParser {
	return &FileParser{
		filePath: filePath,
	}
}

func NewConfig(projectName string, frontend FrontendConfig, backend BackendConfig, devops DevopsConfig) *Config {
	return &Config{
		Name:      projectName,
		Structure: "monorepo",
		Scaffold: ScaffoldConfig{
			Frontend: frontend,
			Backend:  backend,
			Devops:   devops,
		},
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

func (c *Config) GenerateSersiYaml(filePath string) error {
	yamlData, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %v", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	path := filepath.Join(cwd, filePath, "sersi.yaml")

	err = os.WriteFile(path, yamlData, 0o644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func validateConfig(config *Config) error {
	if FileExists(config.Name) {
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
