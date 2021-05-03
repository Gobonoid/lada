package lada

import (
	"io"
	"regexp"
)

import "os"

type Cli struct {
	Version string
	Name string
	output io.Writer
	errorOutput io.Writer
	input io.Reader
}

type Arguments map[string]interface{}
type Parameters map[string]interface{}
type Flags map[string]bool
type Handler func(args Arguments, params Parameters, flags Flags)

type Command struct {
	Handler Handler
	Name string
	Arguments Arguments
	Parameters Parameters
	Flags Flags
}

func NewCli(name string, version string) *Cli {
	return &Cli{
		Name: name,
		Version: version,
		input: os.Stdin,
		output: os.Stdout,
		errorOutput: os.Stderr,
	}
}

func (c *Cli) AddCommand(command string, handler Handler) *Command {
	// [TODO] parse command string
	return &Command{
		Handler: handler,
	}
}

func (c *Cli) Output() io.Writer {
	return c.output
}

func (c *Cli) ErrorOutput() io.Writer {
	return c.errorOutput
}

func (c *Cli) Input() io.Reader {
	return c.input
}

func (c *Command) Run() error {
	c.Handler(c.Arguments, c.Parameters, c.Flags)

	return nil
}

func Parse(commandString string) (string, []string, []string, []string) {
	r := "^(?P<subcommands>(?P<subcommand>\\s*[a-z-_:]+)+)(?P<arguments>(?P<argument>\\s+\\$[a-z]+)*)(?P<parameters>(?P<parameter>\\s*\\-\\-[a-z]+\\s*\\=\\s*[^-]*)*)(?P<flags>(?P<flag>\\s*\\-\\-[a-z]+\\s*\\?)*)$"
	re := regexp.MustCompile(r)
	matches := re.FindStringSubmatch(commandString)

	var (
		command string
		args []string
		params []string
		flags []string
	)

	for key, typeName := range re.SubexpNames() {
		switch typeName {
			case "subcommand":
				command = matches[key]
			case "argument":
				args = append(args, matches[key])
			case "parameter":
				params = append(params, matches[key])
			case "flag":
				flags = append(flags, matches[key])
		}
	}
	return command, args, params, flags
}
