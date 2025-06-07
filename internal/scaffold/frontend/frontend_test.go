package frontend_test

import (
	"testing"

	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/stretchr/testify/assert"
)

func Test_Frontend(t *testing.T) {
	builder := frontend.NewFrontendBuilder()
	builder.ProjectName("test-project").
		Framework("react").
		Language("ts").
		CSS("tailwind").
		Monorepo(true)
	frontend := builder.Build()
	assert.Equal(t, "test-project", frontend.ProjectName)
	assert.Equal(t, "react", frontend.Framework)
	assert.Equal(t, "ts", frontend.Language)
	assert.Equal(t, "tailwind", frontend.CSS)
	assert.True(t, frontend.Monorepo)
}
