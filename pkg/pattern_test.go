package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommandPattern(t *testing.T) {
	t.Run("can parse: cmd $arg1 $arg2 --parameter=default\\ value ", func(t *testing.T) {
		cmd := NewCommandPattern("cmd $arg1 $arg2 --parameter=default\\ value")
		err := cmd.Parse()
		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.commandName)
		assert.Contains(t, cmd.arguments, &Argument{"arg1", false})
		assert.Contains(t, cmd.arguments, &Argument{"arg2", false})
		assert.Contains(t, cmd.parameters, &Parameter{
			LongForm: "parameter",
			DefaultValue: "default value",
		})
	})

	t.Run("can parse: cmd ...$wildcard-arg --parameter[P]=default\\ value ", func(t *testing.T) {
		cmd := NewCommandPattern("cmd ...$wildcard-arg --parameter[P]=default\\ value")
		err := cmd.Parse()
		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.commandName)
		assert.Contains(t, cmd.arguments, &Argument{"wildcard-arg", true})
		assert.Contains(t, cmd.parameters, &Parameter{
			LongForm: "parameter",
			ShortForm: "P",
			DefaultValue: "default value",
		})
	})
}
