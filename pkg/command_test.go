package lada

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommand(t *testing.T) {
	t.Run("can create: hello $name", func(t *testing.T) {
		var result string
		cmd, err := NewCommand("hello $name", func(terminal *Terminal, args Arguments) error {
			result = fmt.Sprintf("Hello %s", args.Get("name").Value())
			return nil
		})

		assert.Nil(t, err)
		assert.Equal(t, nil, cmd.Execute("Tom", &Terminal{}))
		assert.Equal(t, "Hello Tom", result)
	})

	t.Run("can retrieve params as specific type", func(t *testing.T) {

		var result int
		cmd, _ := NewCommand("test --parameter[P]= --flag[F]", func(terminal *Terminal, args Arguments) error {
			result, _ = args.Get("parameter").AsInt()

			return nil
		})

		assert.NotEmpty(t, cmd)
		err := cmd.Execute("-P 10", &Terminal{})
		assert.Nil(t, err)

		assert.Equal(t, 10, result)
	})
}