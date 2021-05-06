package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCli(t *testing.T) {
	t.Run("Test can create new cli", func(t *testing.T) {
		var cli = NewCli("cli name", "1.0.0")
		assert.Equal(t, "cli name", cli.Name)
	})
}
