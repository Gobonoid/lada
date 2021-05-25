package lada

import (
	"github.com/kodemore/lada/pkg/style"
	"os"
	"strings"
)

func showApplicationHelp(a *Application) {
	t := a.terminal
	showApplicationVersion(a)
	t.Print("\n\n")
	showApplicationUsage(a)
	t.Print("\n\n")
	showApplicationCommands(a)
	t.Print("\n")
}

func showApplicationUsage(a *Application)  {
	t := a.terminal
	args := os.Args
	t.PrettyPrint("Usage:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")
	t.Print("    ")
	t.PrettyPrint(args[0], style.Format.Underline)
	t.Print(" [--version] [--help] [<options>] ")
	t.PrettyPrint("<command>", style.Foreground.LightCyan)
	t.Print(" [<args>]")
}

func showApplicationVersion(a *Application)  {
	a.terminal.Print(a.Name + " ")
	a.terminal.PrettyPrint(a.Version, style.Foreground.LightGreen, style.Format.Bold)
	if a.Description != "" {
		a.terminal.Print("\n")
		a.terminal.PrettyPrint("    " + a.Description, style.Format.Dim)
	}
}

func showApplicationCommands(a *Application)  {
	t := a.terminal
	t.PrettyPrint("Commands:", style.Format.Bold, style.Foreground.LightGreen)
	t.Print("\n")
	longestCommandName := 4
	for key, _ := range a.commands {
		if len(key) > longestCommandName {
			longestCommandName = len(key)
		}
	}
	for key, command := range a.commands {
		showApplicationCommand(t, key, command.Description, longestCommandName)
		a.terminal.Print("\n")
	}
	showApplicationCommand(t,"help", "Displays help dialog", longestCommandName)
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
