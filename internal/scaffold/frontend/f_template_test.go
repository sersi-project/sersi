package frontend_test

import (
	"testing"

	"github.com/sersi-project/sersi/internal/scaffold/frontend"
	"github.com/stretchr/testify/assert"
)

func Test_FTemplateBuilder(t *testing.T) {
	builder := frontend.NewFTemplateBuilder()
	builder.ProjectName("test-project").
		Framework("react").
		Language("ts").
		CSS("tailwind").
		Monorepo(true)
	template := builder.Build()
	assert.Equal(t, "test-project", template.ProjectName)
	assert.Equal(t, "react", template.Framework)
	assert.Equal(t, "ts", template.Language)
	assert.Equal(t, "tailwind", template.CSS)
	assert.True(t, template.Monorepo)
}
