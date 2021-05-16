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

	t.Run("can parse catch-all and flags: hello $names* --flag-a[f] --flag-b[F]", func(t *testing.T) {
		definition, _ := NewCommandDefinition("hello $names* --flag-a[f] --flag-b[F]")
		input, err := NewCommandInput("hello Bob Boo John --flag-a -F", definition)
		assert.Nil(t, err)

		assert.Equal(t, "hello", input.commandName)
		assert.Contains(t, input.arguments, "names")
		assert.Equal(t, "Bob Boo John", input.arguments["names"].Value)
		assert.Contains(t, input.parameters, "flag-a")
		assert.Contains(t, input.parameters, "flag-b")
	})
}

func TestNewCommand(t *testing.T) {
	t.Run("can create: hello $name", func(t *testing.T) {
		var result string
		cmd, err := NewCommand("hello $name", func(args Arguments, params Parameters) error {
			if name, ok := params["name"]; ok {
				result = fmt.Sprintf("Hello %s", name.Value)
				return nil
			}

			return errors.New("error")
		})

		assert.Nil(t, err)
		assert.Equal(t, 0, cmd.Execute("hello Tom"))
		assert.Equal(t, "Hello Tom", result)
	})
}