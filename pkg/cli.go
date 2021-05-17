package lada

import (
	"fmt"
	"io"
	"strings"
)

import os "os"

type Cli struct {
	Version     string
	Name        string
	cursor		*Cursor
	output      io.Writer
	errorOutput io.Writer
	input       io.Reader
	commands map[string]*Command
}

func NewCli(name string, version string) (*Cli, error) {
	cursor, err := NewCursor(os.Stdin)
	if err != nil {
		return &Cli{}, err
	}

	return &Cli{
		Name:        name,
		Version:     version,
		input:       os.Stdin,
		output:      os.Stdout,
		errorOutput: os.Stderr,
		cursor: cursor,
		commands: map[string]*Command{},
	}, nil
}

func (c *Cli) Output() io.Writer {
	return c.output
}

func (c *Cli) ErrorOutput() io.Writer {
	return c.errorOutput
}

func (c *Cli) Cursor() *Cursor {
	return c.cursor
}

func (c *Cli) Input() io.Reader {
	return c.input
}

func (c *Cli) WriteError(text string) error {
	_, err := fmt.Fprint(c.errorOutput, text)

	return err
}

func (c *Cli) Write(text string) error {
	return c.cursor.Print(text)
}

func (c *Cli) WriteLine(text string) error {
	err := c.cursor.Print(text)
	if err != nil {
		return err
	}
	return c.cursor.MoveToNextLine()
}

func (c *Cli) AddCommand(definition string, handler Handler) error {
	cmd, err := NewCommand(definition, handler)
	if err != nil {
		return err
	}
	c.commands[cmd.Name] = cmd

	return nil
}

func (c *Cli) Run() int {
	args := os.Args
	cmdName := args[1]
	cliArgs := strings.Join(args[1:], " ")

	if command, ok := c.commands[cmdName]; ok {
		err := command.Execute(cliArgs)
		if err != nil {
			return 1
		}

		return 0
	}

	return 1
}