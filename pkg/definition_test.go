package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommandDefinition(t *testing.T) {
	t.Run("can parse: cmd $arg1 $arg2 --parameter=default\\ value ", func(t *testing.T) {
		cmd, err := NewCommandDefinition("cmd $arg1 $arg2 --parameter=default\\ value")
		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.arguments, Argument{Name: "arg1", Wildcard: false})
		assert.Contains(t, cmd.arguments, Argument{Name: "arg2", Wildcard: false})
		assert.Contains(t, cmd.parameters, Parameter{
			Name:         "parameter",
			DefaultValue: "default value",
		})
	})

	t.Run("can parse: cmd $wildcard-arg* --parameter[P]=default\\ value ", func(t *testing.T) {
		cmd, err := NewCommandDefinition("cmd $wildcard-arg* --parameter[P]=default\\ value")
		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.arguments, Argument{Name: "wildcard-arg", Wildcard: true})
		assert.Contains(t, cmd.parameters, Parameter{
			Name:         "parameter",
			ShortForm:    "P",
			DefaultValue: "default value",
		})
	})

	t.Run("can parse: cmd $arg --flag1[f] --flag2[F] --parameter=", func(t *testing.T) {
		cmd, err := NewCommandDefinition("cmd $arg --flag1[f] --flag2[F] --parameter=")

		assert.Nil(t, err)
		assert.Equal(t, "cmd", cmd.Command())
		assert.Contains(t, cmd.parameters, Parameter{Name: "parameter", IsFlag: false})
		assert.Contains(t, cmd.parameters, Parameter{Name: "flag1", ShortForm: "f", IsFlag: true})
		assert.Contains(t, cmd.parameters, Parameter{Name: "flag2", ShortForm: "F", IsFlag: true})
	})

	t.Run("fail to parse: cmd $arg* $arg2", func(t *testing.T) {
		cmd, err := NewCommandDefinition("cmd $arg* $arg2")
		assert.NotNil(t, err)
		assert.Empty(t, cmd)
	})
}
