package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInputArguments(t *testing.T) {
	t.Run("test positional arguments", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 argN")
		args, err := NewArguments("value1 value2", allArgs)

		assert.Nil(t, err)

		assert.Equal(t, "value1", args.Get("arg1").Value())
		assert.Equal(t, "value2", args.Get("argN").Value())

	})

	t.Run("test missing positional arguments", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 argN")
		args, err := NewArguments("value1", allArgs)

		assert.NotNil(t, err)
		assert.Empty(t, args)
	})

	t.Run("test wildcard argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg...")
		args, err := NewArguments("value1 value2 value3 value4 value5", allArgs)

		assert.Nil(t, err)
		assert.Equal(t, "value1 value2 value3 value4 value5", args.Get("arg").Value())
	})

	t.Run("test arguments with wildcard argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 arg2 argN...")
		args, err := NewArguments("value1 value2 value3 value4 value5", allArgs)
		assert.Nil(t, err)

		assert.Equal(t, "value3 value4 value5", args.Get("argN").Value())
	})

	t.Run("test optional argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("--optional= --optional2=default\\ value")
		args, err := NewArguments("--optional=123", allArgs)
		assert.Nil(t, err)
		assert.Equal(t, "123", args.Get("optional").Value())
		assert.Equal(t, "default value", args.Get("optional2").Value())
	})

	t.Run("test optional argument + wildcard", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg... --optional=")
		args, err := NewArguments("--optional=123 1 2 3 4", allArgs)
		assert.Nil(t, err)

		assert.Equal(t, "123", args.Get("optional").Value())
		assert.Equal(t, "1 2 3 4", args.Get("arg").Value())
	})
}
