package backend_test

import (
	"testing"

	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/stretchr/testify/assert"
)

func TestBTemplateBuilder(t *testing.T) {
	builder := backend.NewBTemplateBuilder()
	builder.ProjectName("test-project").
		Framework("express").
		Language("javascript").
		Database("mongodb").
		Monorepo(true)

	template := builder.Build()
	assert.Equal(t, "test-project", template.ProjectName)
	assert.Equal(t, "express", template.Framework)
	assert.Equal(t, "javascript", template.Language)
	assert.Equal(t, "mongodb", template.Database)
	assert.True(t, template.Monorepo)
}

func TestBTemplateInvalidLanguage(t *testing.T) {
	builder := backend.NewBTemplateBuilder()
	builder.ProjectName("test-project").
		Language("invalid-lang")
	template := builder.Build()

	// Should fail with error for invalid language
	err := template.Execute()
	assert.Error(t, err)
}

func TestBTemplateEmptyProjectName(t *testing.T) {
	builder := backend.NewBTemplateBuilder()
	template := builder.Build()

	// Should fail with error for empty project name
	err := template.Execute()
	assert.Error(t, err)
}
