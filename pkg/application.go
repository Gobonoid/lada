package lada

import (
	"os"
	"strings"
)

type ApplicationCommands map[string]*Command

type Application struct {
	Name     string
	Version  string
	terminal *Terminal
	commands ApplicationCommands
}

func NewApplication(name, version string) (*Application, error) {
	terminal, err := NewTerminal()
	if err != nil {
		return &Application{}, err
	}

	return &Application{
		Name:     name,
		Version:  version,
		terminal: terminal,
		commands: make(ApplicationCommands),
	}, nil
}


func (a *Application) AddCommand(format string, handler Handler) error {
	cmd, err := NewCommand(format, handler)
	if err != nil {
		return err
	}
	a.commands[cmd.Name] = cmd
	return nil
}

func (a *Application) Run() int {
	args := os.Args
	if len(args) < 2 {
		// Run help command here
		return 0
	}

	cmdName := args[1]
	cliArgs := strings.Join(args[1:], " ")

	if command, ok := a.commands[cmdName]; ok {
		err := command.Execute(cliArgs, a.terminal)
		a.terminal.close()
		if err != nil {
			return 1
		}

		return 0
	}

	return 1
}

