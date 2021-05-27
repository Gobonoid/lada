package lada

import (
	"github.com/kodemore/lada/pkg/style"
	"os"
	"strings"
)

func showApplicationHelp(a *Application) {
	t := a.terminal
	showApplicationVersion(a)
	if a.Description != "" {
		a.terminal.Print("\n")
		a.terminal.PrettyPrint("    " + a.Description, style.Format.Dim)
	}
	t.Print("\n\n")
	showApplicationUsage(a)
	t.Print("\n\n")
	showApplicationCommands(a)
	t.Print("\n")
}

func showCommandHelp(a *Application, c *Command) {
	t := a.terminal
	showApplicationVersion(a)
	if c.Description != "" {
		a.terminal.Print("\n")
		a.terminal.PrettyPrint("    " + c.Description, style.Format.Dim)
	}
	t.Print("\n\n")
	t.PrettyPrint("Usage:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")
	t.Print("    ")
	t.Print(os.Args[0] + " ")
	t.PrettyPrint(c.Verb(), style.Format.Underline)
	t.Print(" [--help | -h]")

	columnSize := 4
	for _, arg := range c.Pattern.Arguments {
		if len(arg.Name) > columnSize {
			columnSize = len(arg.Name)
		}
		if arg.IsPositional() {
			continue
		}
		t.Print(" [")
		t.Print("--" + arg.Name)
		if arg.ShortName != "" {
			t.Print(" | -" + arg.ShortName)
		}
		t.Print("]")
	}

	for _, arg := range c.Pattern.Arguments.GetPositionalArguments() {
		t.Print(" <" + arg.Name + ">")
	}

	if !c.IsHelpAvailable() {
		return
	}

	t.Print("\n\n")
	t.PrettyPrint("Arguments:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")

	for _, arg := range c.Pattern.Arguments {
		if arg.IsPositional() {
			continue
		}
		showCommandArgument(t, arg, columnSize)
		t.Print("\n")
	}


	for _, arg := range c.Pattern.Arguments.GetPositionalArguments() {
		showCommandArgument(t, arg, columnSize)
		t.Print("\n")
	}
}

func showCommandArgument(t *Terminal, argument *Argument, columnSize int) {

	t.Print("    ")
	if argument.IsOptional() {
		t.PrettyPrint("--" + argument.Name, style.Foreground.LightCyan)
		if argument.ShortName != "" {
			t.PrettyPrint(", -" + argument.ShortName, style.Foreground.LightCyan)
		}
	} else {
		t.PrettyPrint("<" + argument.Name + ">", style.Foreground.LightCyan)
	}

	spaces := columnSize + 10 - getArgLength(argument)
	t.Print(strings.Repeat(" ", spaces))

	if argument.defaultValue != "" {
		t.PrettyPrint("default value: ", style.Format.Dim)
		t.PrettyPrint(argument.defaultValue, style.Format.Dim, style.Format.Underline)
		t.PrettyPrint(", ", style.Format.Dim)
	}

	if argument.Description != "" {
		t.PrettyPrint(argument.Description, style.Format.Dim)
	}
}

func getArgLength(a *Argument) int {
	columns := len(a.Name) + 2
	if a.IsPositional() {
		return columns
	}
	if a.ShortName != "" {
		return columns + 4
	}
	return columns
}

func showApplicationUsage(a *Application)  {
	t := a.terminal
	args := os.Args
	t.PrettyPrint("Usage:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")
	t.Print("    ")
	t.PrettyPrint(args[0], style.Format.Underline)
	t.Print(" [--version | -v] [--help | -h] ")
	t.PrettyPrint("<command>", style.Foreground.LightCyan)
	t.Print(" [<args>]")
}

func showApplicationVersion(a *Application)  {
	a.terminal.Print(a.Name + " ")
	a.terminal.PrettyPrint(a.Version, style.Foreground.LightGreen, style.Format.Bold)
}

func showApplicationCommands(a *Application)  {
	t := a.terminal
	t.PrettyPrint("Commands:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")
	columnSize := 4
	for key, _ := range a.commands {
		if len(key) > columnSize {
			columnSize = len(key)
		}
	}
	for key, command := range a.commands {
		showApplicationCommand(t, key, command.Description, columnSize)
		a.terminal.Print("\n")
	}
	showApplicationCommand(t,"help", "Displays help dialog", columnSize)
}

func showApplicationCommand(t *Terminal, cmdName string, description string, columnSize int) {
	t.Print("    ")
	t.PrettyPrint(cmdName, style.Foreground.LightCyan)
	if description != "" {
		spaces := columnSize + 4 - len(cmdName)
		t.Print(strings.Repeat(" ", spaces))
		t.PrettyPrint(description, style.Format.Dim)
	}
}
