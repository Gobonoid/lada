package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArgumentFromFormat(t *testing.T) {
	t.Run("can parse positional argument", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("$test")
		assert.Nil(t, err)
		assert.Equal(t, "test", arg.Name)
		assert.Empty(t, arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, PositionalArgument, arg.kind)
	})
	t.Run("can parse wildcard argument", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("$test...")
		assert.Nil(t, err)
		assert.Equal(t, "test", arg.Name)
		assert.Empty(t, arg.ShortName)
		assert.True(t, arg.IsWildcard())
		assert.Equal(t, PositionalArgument, arg.kind)
	})

	t.Run("can parse optional argument", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("--optional=")
		assert.Nil(t, err)
		assert.Equal(t, "optional", arg.Name)
		assert.Empty(t, arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, OptionalArgument, arg.kind)
		assert.Empty(t, arg.DefaultValue())
	})

	t.Run("can parse optional argument with short name", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("--optional[o]=")
		assert.Nil(t, err)
		assert.Equal(t, "optional", arg.Name)
		assert.Equal(t, "o", arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, OptionalArgument, arg.kind)
		assert.Empty(t, arg.DefaultValue())
	})

	t.Run("can parse optional argument with default value", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("--optional[o]=default value")
		assert.Nil(t, err)
		assert.Equal(t, "optional", arg.Name)
		assert.Equal(t, "o", arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, OptionalArgument, arg.kind)
		assert.Equal(t, "default value", arg.DefaultValue())
	})

	t.Run("can parse flag argument", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("--flag")
		assert.Nil(t, err)
		assert.Equal(t, "flag", arg.Name)
		assert.Empty(t, arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, FlagArgument, arg.kind)
		assert.Equal(t, "0", arg.DefaultValue())
	})

	t.Run("can parse flag argument with short name", func(t *testing.T) {
		arg, err := NewArgumentFromCommandPattern("--flag[f]")
		assert.Nil(t, err)
		assert.Equal(t, "flag", arg.Name)
		assert.Equal(t, "f", arg.ShortName)
		assert.False(t, arg.IsWildcard())
		assert.Equal(t, FlagArgument, arg.kind)
		assert.Equal(t, "0", arg.DefaultValue())
	})
}