package backend_test

import (
	"testing"

	"github.com/sersi-project/sersi/internal/scaffold/backend"
	"github.com/stretchr/testify/assert"
)

func TestBackedBuilder(t *testing.T) {
	builder := backend.NewBackendBuilder()
	builder.ProjectName("test-project").
		Language("javascript")
	backend := builder.Build()
	assert.Equal(t, "test-project", backend.ProjectName)
	assert.Equal(t, "javascript", backend.Language)
}
