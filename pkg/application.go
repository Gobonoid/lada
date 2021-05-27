package lada

import (
	"errors"
	"fmt"
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


func (a *Application) AddCommand(pattern string, handler Handler) {
	cmd, err := NewCommand(pattern, handler)
	if err != nil {
		panic(err)
	}
	a.commands[cmd.Verb()] = cmd
}

func (a *Application) AddHelp(verb string, description string, args ArgumentsHelp) {
	if command, ok := a.commands[verb]; ok {
		command.Description = description
		for argName, argDesc := range args {
			if arg, ok := command.Pattern.Arguments.GetArgumentByName(argName); ok {
				arg.Description = argDesc
			}
		}
		return
	}

	panic("cannot add help to non existing command " + verb)
}

func (a *Application) Run() int {
	return a.RunWithArgs(os.Args...)
}

func (a *Application) RunWithArgs(args ...string) int {
	// no verb provided so we should display help dialog
	if len(args) < 2 {
		showApplicationHelp(a)
		a.terminal.close()
		return 0
	}

	cmdName := args[1]
	if command, ok := a.commands[cmdName]; ok {
		cliArgs := strings.Join(args[2:], " ")
		if strings.Contains(cliArgs, "--help") {
			showCommandHelp(a, command)
			a.terminal.close()
			return 0
		}
		return a.executeCommand(command, cliArgs)
	}

	if cmdName == "help" || cmdName == "--help" || cmdName == "-h" {
		showApplicationHelp(a)
		a.terminal.close()
		return 0
	}

	if cmdName == "--version" || cmdName == "-v" {
		showApplicationVersion(a)
		a.terminal.close()
		return 0
	}

	if cmdName[0] == '-' {
		a.terminal.PrintError("Unknown argument ")
	} else {
		a.terminal.PrintError("Unknown command ")
	}
	a.terminal.PrettyPrint(cmdName, style.Foreground.LightRed, style.Format.Underline)
	a.terminal.PrettyPrint(", run ", style.Foreground.Red)

	a.terminal.PrettyPrint(args[0], style.Foreground.LightRed, style.Format.Underline, style.Format.Bold)
	a.terminal.PrettyPrint(" ", style.Foreground.Red)
	a.terminal.PrettyPrint("help", style.Foreground.LightRed, style.Format.Underline, style.Format.Bold)
	a.terminal.PrettyPrint(" for usage", style.Foreground.Red)
	a.terminal.Print("\n")
	a.terminal.close()
	return 1
}

func (a *Application) executeCommand(command *Command, args string) int {
	err := command.Execute(args, a.terminal)

	if err != nil {
		if errors.Is(err, MissingArgumentValueError) {
			verb := command.Verb()
			if verb == "*" {
				verb = ""
			}
			a.terminal.PrettyPrint(fmt.Sprintf("Missing argument: %s, please run %s --help to learn more", err.Error(), verb), style.Foreground.Red)
			a.terminal.close()
			return 1
		}
		a.terminal.PrettyPrint(err.Error(), style.Foreground.Red)
		a.terminal.close()
		return 1
	}

	a.terminal.close()
	return 0
}