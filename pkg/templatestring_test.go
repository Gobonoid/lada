package lada

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTemplateString_Substitute(t *testing.T) {
	t.Run("can substitute: test { value }", func(t *testing.T) {
		var template TemplateString = "test { value }"
		substituted, err := template.Substitute(map[string]string{"value": "value"})

		assert.Nil(t, err)
		assert.Equal(t, "test value", substituted)
	})
}

func TestTemplateString_SubstituteWithFilters(t *testing.T) {
	t.Run("can substitute with uppercase: test { value| uppercase }", func(t *testing.T) {
		var template TemplateString = "test {value | uppercase }"
		params := map[string]string{"value": "value"}
		filters := Filters{
			"uppercase": func(s string) string {
				return strings.ToUpper(s)
			},
		}
		parsed, err := template.SubstituteWithFilters(params, filters)

		assert.Nil(t, err)
		assert.Equal(t, "test VALUE", parsed)
	})

	t.Run("can substitute with two filters: test { value | ucfirst | uclast }", func(t *testing.T) {
		var template TemplateString = "test { value | ucfirst | uclast }"
		params := map[string]string{"value": "value"}
		filters := Filters{
			"ucfirst": func(s string) string {
				return strings.ToUpper(string(s[0])) + s[1:]
			},
			"uclast": func(s string) string {
				return s[0:len(s)-1] + strings.ToUpper(string(s[len(s)-1]))
			},
		}
		parsed, err := template.SubstituteWithFilters(params, filters)

		assert.Nil(t, err)
		assert.Equal(t, "test ValuE", parsed)
	})
}
