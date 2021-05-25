package lada

import (
	"github.com/kodemore/lada/pkg/style"
	"os"
	"strings"
)

type ApplicationCommands map[string]*Command

type Application struct {
	Name        string
	Version     string
	Description string
	terminal    *Terminal
	commands    ApplicationCommands
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


func (a *Application) AddCommand(format string, description string, handler Handler) error {
	cmd, err := NewCommand(format, handler)
	if err != nil {
		return err
	}
	cmd.Description = description
	a.commands[cmd.Verb()] = cmd
	return nil
}

func (a *Application) Run() int {
	args := os.Args
	if len(args) < 2 {
		showApplicationHelp(a)
		a.terminal.close()
		return 0
	}

	cmdName := args[1]
	cliArgs := strings.Join(args[1:], " ")

	if command, ok := a.commands[cmdName]; ok {
		if strings.Contains(cliArgs, "--help") {
			a.terminal.Print("Show help")

			a.terminal.close()
			return 0
		}

		err := command.Execute(cliArgs, a.terminal)
		a.terminal.close()
		if err != nil {
			return 1
		}

		return 0
	}

	if cmdName == "help" {
		showApplicationHelp(a)
		a.terminal.close()
		return 0
	}

	a.terminal.PrettyPrint("Unknown command ", style.Foreground.Red)
	a.terminal.PrettyPrint(cmdName, style.Foreground.LightRed, style.Format.Underline)
	a.terminal.PrettyPrint(", run ", style.Foreground.Red)

	a.terminal.PrettyPrint(args[0], style.Foreground.LightRed, style.Format.Underline, style.Format.Bold)
	a.terminal.PrettyPrint(" ", style.Foreground.Red)
	a.terminal.PrettyPrint("help", style.Foreground.LightRed, style.Format.Underline, style.Format.Bold)
	a.terminal.PrettyPrint(" to list available commands.", style.Foreground.Red)
	a.terminal.Print("\n")
	a.terminal.close()
	return 1
}