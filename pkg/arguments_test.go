package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInputArguments(t *testing.T) {
	t.Run("test positional arguments", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 argN")
		inputArgs, err := NewInputArguments("value1 value2", allArgs)

		assert.Nil(t, err)
		if arg, ok := inputArgs.Get("arg1"); ok {
			assert.Equal(t, "value1", arg.Value())
		} else {
			t.Fail()
		}

		if arg, ok := inputArgs.Get("argN"); ok {
			assert.Equal(t, "value2", arg.Value())
		} else {
			t.Fail()
		}
	})

	t.Run("test missing positional arguments", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 argN")
		args, err := NewInputArguments("value1", allArgs)

		assert.NotNil(t, err)
		assert.Empty(t, args)
	})

	t.Run("test wildcard argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg...")
		args, err := NewInputArguments("value1 value2 value3 value4 value5", allArgs)
		assert.Nil(t, err)

		if arg, ok := args.Get("arg"); ok {
			assert.Equal(t, "value1 value2 value3 value4 value5", arg.Value())
		} else {
			t.Fail()
		}
	})

	t.Run("test arguments with wildcard argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg1 arg2 argN...")
		args, err := NewInputArguments("value1 value2 value3 value4 value5", allArgs)
		assert.Nil(t, err)

		if arg, ok := args.Get("argN"); ok {
			assert.Equal(t, "value3 value4 value5", arg.Value())
		} else {
			t.Fail()
		}
	})

	t.Run("test optional argument", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("--optional= --optional2=default\\ value")
		args, err := NewInputArguments("--optional=123", allArgs)
		assert.Nil(t, err)

		if arg, ok := args.Get("optional"); ok {
			assert.Equal(t, "123", arg.Value())
		} else {
			t.Fail()
		}

		if arg, ok := args.Get("optional2"); ok {
			assert.Equal(t, "default value", arg.Value())
		} else {
			t.Fail()
		}
	})

	t.Run("test optional argument + wildcard", func(t *testing.T) {
		allArgs, _ := NewCommandPatternArguments("arg... --optional=")
		args, err := NewInputArguments("--optional=123 1 2 3 4", allArgs)
		assert.Nil(t, err)

		if arg, ok := args.Get("optional"); ok {
			assert.Equal(t, "123", arg.Value())
		} else {
			t.Fail()
		}

		if arg, ok := args.Get("arg"); ok {
			assert.Equal(t, "1 2 3 4", arg.Value())
		} else {
			t.Fail()
		}
	})
}
