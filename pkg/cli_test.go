package lada

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCli(t *testing.T) {
	t.Run("Test can create new cli", func(t *testing.T) {
		var cli = NewCli("cli name", "1.0.0")
		assert.Equal(t, "cli name", cli.Name)
	})
}

func TestAddCommand(t *testing.T) {
	t.Run("Test can add a CLI command", func(t *testing.T) {
		var cli = NewCli("cli name", "1.0.0")

		commandString := "subcommand $input $output --num=1 --verbose? --flavour= --word=default --flag?"
		command := cli.AddCommand(commandString, func(args Arguments, params Parameters, flags Flags) {
			//assert.Len(t, args, 2)
			//assert.Len(t, params, 3)
			//assert.Len(t, flags, 2)
			cli.Output().Write([]byte("Test command"))
		})
		err := command.Run()
		assert.Equal(t, nil, err)
	})
}

func TestParse(t *testing.T) {
	commandString := "subcommand adsasd $input $output --num=1 --flavour= --word=default --verbose? --flag?"
	tokens := Parse(commandString)
	assert.Len(t, tokens, 9)
}
