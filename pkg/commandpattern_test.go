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
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.arguments, &Argument{Name: "arg1", Wildcard: false})
		assert.Contains(t, cmd.arguments, &Argument{Name: "arg2", Wildcard: false})
		assert.Contains(t, cmd.parameters, &Parameter{
			LongForm: "parameter",
			DefaultValue: "default value",
		})
	})

	t.Run("can parse: cmd ...$wildcard-arg --parameter[P]=default\\ value ", func(t *testing.T) {
		cmd := NewCommandPattern("cmd ...$wildcard-arg --parameter[P]=default\\ value")
		err := cmd.Parse()
		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.arguments, &Argument{Name: "wildcard-arg", Wildcard: true})
		assert.Contains(t, cmd.parameters, &Parameter{
			LongForm: "parameter",
			ShortForm: "P",
			DefaultValue: "default value",
		})
	})

	t.Run("can parse: cmd $arg --flag1[f] --flag2[F] --parameter=", func(t *testing.T) {
		cmd := NewCommandPattern("cmd $arg --flag1[f] --flag2[F] --parameter=")
		err := cmd.Parse()

		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.parameters, &Parameter{LongForm: "parameter"})
		assert.Contains(t, cmd.flags, &Flag{LongForm: "flag1", ShortForm: "f"})
		assert.Contains(t, cmd.flags, &Flag{LongForm: "flag2", ShortForm: "F"})
	})
}
