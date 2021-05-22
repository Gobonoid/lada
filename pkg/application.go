package lada

import (
	"os"
	"strings"
)

type ApplicationCommands map[string]*Command

type Application struct {
	Name     string
	Version  string
	Description string
	StyleSheet StyleSheet
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


func (a *Application) AddCommand(format string, description string, handler Handler) error {
	cmd, err := NewCommand(format, handler)
	if err != nil {
		return err
	}
	cmd.Description = description
	a.commands[cmd.Name] = cmd
	return nil
}

func (a *Application) Run() int {
	args := os.Args
	if len(args) < 2 {
		a.showHelpDialog()
		a.terminal.close()
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

	if cmdName == "help" {
		a.showHelpDialog()
		a.terminal.close()
		return 0
	}

	a.terminal.PrettyPrint("Unknown command ", Foreground.Red)
	a.terminal.PrettyPrint(cmdName, Foreground.LightRed, Format.Underline)
	a.terminal.PrettyPrint(", run ", Foreground.Red)

	a.terminal.PrettyPrint(args[0], Foreground.LightRed, Format.Underline, Format.Bold)
	a.terminal.PrettyPrint(" ", Foreground.Red)
	a.terminal.PrettyPrint("help", Foreground.LightRed, Format.Underline, Format.Bold)
	a.terminal.PrettyPrint(" to list available commands.", Foreground.Red)
	a.terminal.Print("\n")
	a.terminal.close()
	return 1
}

func (a *Application) showHelpDialog() {
	a.showApplicationVersionDialog()
	a.terminal.Print("\n\n")
	a.showUsageDialog()
	a.terminal.Print("\n\n")
	a.showApplicationCommandsDialog()
	a.terminal.Print("\n")
}

func (a *Application) showUsageDialog() {
	args := os.Args
	a.terminal.PrettyPrint("Usage:", Format.Bold, Foreground.LightGreen)
	a.terminal.Print("\n")
	a.terminal.Print("    ")
	a.terminal.PrettyPrint(args[0], Format.Underline)
	a.terminal.Print(" [<options>] ")
	a.terminal.PrettyPrint("<command>", Foreground.LightCyan)
	a.terminal.Print(" [<arg>] ... [<arg n>] ")
}

func (a *Application) showApplicationVersionDialog() {
	a.terminal.Print(a.Name + " ")
	a.terminal.PrettyPrint(a.Version, Foreground.LightGreen, Format.Bold)
	if a.Description != "" {
		a.terminal.Print("\n")
		a.terminal.PrettyPrint("    " + a.Description, Format.Dim)
	}
}

func (a *Application) showApplicationCommandsDialog() {
	a.terminal.PrettyPrint("Commands:", Format.Bold, Foreground.LightGreen)
	a.terminal.Print("\n")
	longestCommandName := 4
	for key, _ := range a.commands {
		if len(key) > longestCommandName {
			longestCommandName = len(key)
		}
	}
	for key, command := range a.commands {
		a.showCommandDetailsDialog(key, command.Description, longestCommandName)
		a.terminal.Print("\n")
	}
	a.showCommandDetailsDialog("help", "Displays help dialog", longestCommandName)
}

func (a *Application) showCommandDetailsDialog(cmdName string, description string, columnSize int) {
	a.terminal.Print("    ")
	a.terminal.PrettyPrint(cmdName, Foreground.LightCyan)
	if description != "" {
		spaces := columnSize + 4 - len(cmdName)
		a.terminal.Print(strings.Repeat(" ", spaces))
		a.terminal.PrettyPrint(description, Format.Dim)
	}
}