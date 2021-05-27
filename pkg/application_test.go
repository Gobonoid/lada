package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplication(t *testing.T) {
	t.Run("test can create application", func(t *testing.T) {
		app, err := NewApplication("app name", "1.0.0")
		assert.Nil(t, err)
		assert.Equal(t, "app name", app.Name)
		assert.Equal(t, "1.0.0", app.Version)
	})
}

func TestApplication_AddCommand(t *testing.T) {
	t.Run("test create command", func(t *testing.T) {
		app, _ := NewApplication("app name", "1.0.0")
		app.AddCommand("test", func(terminal *Terminal, arguments Arguments) error {
			return nil
		})

		assert.Contains(t, app.commands, "test")
	})
}
