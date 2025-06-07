package backend_test

import (
	"testing"

	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/stretchr/testify/assert"
)

func Test_BackedBuilder(t *testing.T) {
	builder := backend.NewBackendBuilder()
	builder.ProjectName("test-project").
		Language("js").
		Framework("express").
		Database("postgresql")
	backend := builder.Build()
	assert.Equal(t, "test-project", backend.ProjectName)
	assert.Equal(t, "js", backend.Language)
	assert.Equal(t, "express", backend.Framework)
	assert.Equal(t, "postgresql", backend.Database)
}

func Test_FormatLanguage(t *testing.T) {
	t.Run("javascript(node)", func(t *testing.T) {
		builder := backend.NewBackendBuilder()
		builder.ProjectName("test-project").
			Language("Javascript(node)")
		backend := builder.Build()
		assert.Equal(t, "js", backend.Language)
	})

	t.Run("typescript(ts)", func(t *testing.T) {
		builder := backend.NewBackendBuilder()
		builder.ProjectName("test-project").
			Language("Typescript(ts)")
		backend := builder.Build()
		assert.Equal(t, "ts", backend.Language)
	})

	t.Run("python", func(t *testing.T) {
		builder := backend.NewBackendBuilder()
		builder.ProjectName("test-project").
			Language("python")
		backend := builder.Build()
		assert.Equal(t, "py", backend.Language)
	})
}

func Test_FormatDatabase(t *testing.T) {
	t.Run("postgresql", func(t *testing.T) {
		builder := backend.NewBackendBuilder()
		builder.ProjectName("test-project").
			Database("postgresql")
		backend := builder.Build()
		assert.Equal(t, "postgresql", backend.Database)
	})

	t.Run("mongodb", func(t *testing.T) {
		builder := backend.NewBackendBuilder()
		builder.ProjectName("test-project").
			Database("MongoDB")
		backend := builder.Build()
		assert.Equal(t, "mongodb", backend.Database)
	})
}
