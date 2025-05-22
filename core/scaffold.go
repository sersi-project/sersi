package core

import (
	"sersi/utils"
	"strings"
)

type Scaffold struct {
	ProjectName string
	Framework   string
	CSS         string
	Language    string
}

type ScaffoldBuilder struct {
	config *Scaffold
}

func NewScaffoldBuilder() *ScaffoldBuilder {
	return &ScaffoldBuilder{
		config: &Scaffold{},
	}
}
func (sb *ScaffoldBuilder) ProjectName(name string) *ScaffoldBuilder {
	sb.config.ProjectName = name
	return sb
}

func (sb *ScaffoldBuilder) Framework(framework string) *ScaffoldBuilder {
	sb.config.Framework = strings.ToLower(framework)
	return sb
}

func (sb *ScaffoldBuilder) CSS(css string) *ScaffoldBuilder {
	sb.config.CSS = strings.ToLower(css)
	return sb
}

func (sb *ScaffoldBuilder) Language(lang string) *ScaffoldBuilder {
	sb.config.Language = strings.ToLower(lang)
	return sb
}

func (sb *ScaffoldBuilder) Build() *Scaffold {
	return sb.config
}

func (s *Scaffold) Execute() {
	// Create a directory
	err := utils.CreateDirectory(s.ProjectName)
	if err != nil {
		panic(err)
	}

	err = AddPublicFolder(s.ProjectName)
	if err != nil {
		panic(err)
	}

	gtFramework := s.Framework

	if s.Framework == "react" && s.Language == "typescript" {
		gtFramework = "react-ts"
	}

	if s.Framework == "vanilla" && s.Language == "typescript" {
		gtFramework = "vanilla-ts"
	}

	goldenTemplate := NewGoldenArchitecture(s.ProjectName, gtFramework)
	err = goldenTemplate.Generate()
	if err != nil {
		panic(err)
	}

	templateBuilder := NewTemplateBuilder().
		ProjectName(s.ProjectName).
		Framework(s.Framework).
		CSS(s.CSS)
	if s.Language == "typescript" {
		templateBuilder.Language("ts")
	} else {
		templateBuilder.Language("js")
	}

	template := templateBuilder.Build()

	err = template.Generate()
	if err != nil {
		panic(err)
	}
}
