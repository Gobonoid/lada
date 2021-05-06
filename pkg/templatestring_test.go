package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateString_Substitute(t *testing.T) {
	t.Run("can substitute: test { value }", func(t *testing.T) {
		var template TemplateString = "test { value }"

		substituted := template.Substitute(map[string]string{"value": "value"})

		assert.Equal(t, "test value", substituted)
	})
}
