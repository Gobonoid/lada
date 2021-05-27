package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommandFormat(t *testing.T) {
	t.Run("can parse command with single argument", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
	})

	t.Run("can parse command with single wildcard argument", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1...")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: true})
	})

	t.Run("can parse command with multiple arguments", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1 $arg2 $arg3")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg2", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg3", wildcard: false})
	})

	t.Run("can parse command with multiple arguments and wildcard", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1 $argN...")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "argN", wildcard: true})
	})

	t.Run("can parse command with multiple arguments and options", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1 $argN... --option[o]=")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "argN", wildcard: true})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "option", ShortName: "o", wildcard: false, kind: OptionalArgument})
	})

	t.Run("can parse command with multiple arguments and multiple options", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb $arg1 $argN... --option[o]= --option2[O]=")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "argN", wildcard: true})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "option", ShortName: "o", wildcard: false, kind: OptionalArgument})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "option2", ShortName: "O", wildcard: false, kind: OptionalArgument})
	})

	t.Run("can parse command with multiple spaces", func(t *testing.T) {
		cmd, err := NewCommandPattern("verb  $arg1 \n  $argN... --option[o]=   --option2[O]=   ")
		assert.Nil(t, err)
		assert.Equal(t, "verb", cmd.Verb())
		assert.Contains(t, cmd.Arguments, &Argument{Name: "arg1", wildcard: false})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "argN", wildcard: true})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "option", ShortName: "o", wildcard: false, kind: OptionalArgument})
		assert.Contains(t, cmd.Arguments, &Argument{Name: "option2", ShortName: "O", wildcard: false, kind: OptionalArgument})
	})

	t.Run("can parse catch all command", func(t *testing.T) {
		cmd, err := NewCommandPattern("*")
		assert.Nil(t, err)
		assert.Equal(t, "*", cmd.Verb())
		assert.True(t, cmd.IsCatchAll())
	})
}