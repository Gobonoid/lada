package lada

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommandInput(t *testing.T) {
	t.Run("can parse argument: hello $name", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello $name")
		input, err := NewCommandInput("hello Bob", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.arguments, "name")
		assert.Equal(t, "Bob", input.arguments["name"].Value)
	})

	t.Run("can parse catch-all argument: hello $name $names*", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello $name $names*")
		input, err := NewCommandInput("hello Bob Rob John Josh  Smith", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.arguments, "name")
		assert.Equal(t, "Bob", input.arguments["name"].Value)
		assert.Equal(t, "Rob John Josh Smith", input.arguments["names"].Value)
	})

	t.Run("can parse command with flags by full name: hello --flag-a[f] --flag-b[F]", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --flag-a[f] --flag-b[F]")
		input, err := NewCommandInput("hello --flag-a --flag-b", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.parameters, "flag-a")
		assert.Contains(t, input.parameters, "flag-b")
	})

	t.Run("can parse command with grouped flags: hello --flag-a[f] --flag-b[F]", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --flag-a[f] --flag-b[F]")
		input, err := NewCommandInput("hello -fF", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.parameters, "flag-a")
		assert.Contains(t, input.parameters, "flag-b")
	})

	t.Run("can parse parameter with values: hello --parameter= --parameter-2=", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --parameter= --parameter-2=")
		input, err := NewCommandInput("hello --parameter=test --parameter-2=test\\ 2", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.parameters, "parameter")
		assert.Contains(t, input.parameters, "parameter-2")

		assert.Equal(t, "test", input.parameters["parameter"].Value)
		assert.Equal(t, "test 2", input.parameters["parameter-2"].Value)
	})

	t.Run("can parse short name parameter with values: hello --parameter[P]= --parameter-2[p]=", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --parameter[P]= --parameter-2[p]=")
		input, err := NewCommandInput("hello -P test -p test\\ 2", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.parameters, "parameter")
		assert.Contains(t, input.parameters, "parameter-2")

		assert.Equal(t, "test", input.parameters["parameter"].Value)
		assert.Equal(t, "test 2", input.parameters["parameter-2"].Value)
	})

	t.Run("fails when P's value is missing: hello --parameter[P]=", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --parameter[P]=")
		input, err := NewCommandInput("hello -P", definition)
		assert.Empty(t, input)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, MissingParameterValueError))
	})

	t.Run("can retrieve default value: hello --parameter[P]=default", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello --parameter[P]=default")
		input, err := NewCommandInput("hello", definition)
		assert.Nil(t, err)
		assert.Equal(t, "default", input.parameters["parameter"].Value)
	})
}

func TestNewCommand(t *testing.T) {
	t.Run("can create: hello $name", func(t *testing.T) {
		var result string
		cmd, err := NewCommand("hello $name", func(args Arguments, params Parameters) error {
			if name, ok := args["name"]; ok {
				result = fmt.Sprintf("Hello %s", name.Value)
				return nil
			}

			return errors.New("error")
		})

		assert.Nil(t, err)
		assert.Equal(t, nil, cmd.Execute("hello Tom"))
		assert.Equal(t, "Hello Tom", result)
	})

	t.Run("can retrieve params as specific type", func(t *testing.T) {

		var result int
		cmd, _ := NewCommand("test --parameter[P]= --flag[F]", func(args Arguments, params Parameters) error {
			result, _ = params["parameter"].AsInt()
			return nil
		})

		assert.NotEmpty(t, cmd)
		err := cmd.Execute("test -P 10")
		assert.Nil(t, err)

		assert.Equal(t, 10, result)
	})
}