package lada

import (
	"strings"
)

type Arguments map[string]Argument
type Options map[string]Option

type Handler func(terminal *Terminal, arguments Arguments, options Options) error

type Command struct {
	Name       string
	Handler    Handler
	Definition *CommandFormat
}

func NewCommand(format string, handler Handler) (*Command, error) {
	cmdDefinition, err := NewCommandFormat(format)
	if err != nil {
		return &Command{}, err
	}

	cmd := &Command{
		Name: cmdDefinition.commandName,
		Handler: handler,
		Definition: cmdDefinition,
	}

	return cmd, nil
}

type commandInput struct {
	commandName string
	arguments   Arguments
	parameters  Options
}

func NewCommandInput(cmd string, definition *CommandFormat) (*commandInput, error) {
	parts := splitCommandFormat(cmd)
	if parts[0] != definition.Command() {
		return &commandInput{}, CommandNameDoesNotMatchError.New(parts[0], definition.Command())
	}
	cmdInput := &commandInput{
		commandName: parts[0],
		arguments: make(Arguments),
		parameters: make(Options),
	}
	argIndex := 0
	i := 1
	iMax := len(parts)
	for i < iMax {
		part := parts[i]

		// long form for parameter or flag
		if len(part) > 1 && part[0:2] == "--" {
			p := strings.Split(part[2:], "=")
			param, ok := definition.GetParameter(p[0])
			if !ok {
				return &commandInput{}, UnexpectedParameterError.New(p[0], cmd)
			}
			if param.IsFlag && len(p) > 1 {
				return &commandInput{}, UnexpectedFlagValueError.New(param.Name)
			}

			if param.IsFlag {
				param.Value = "1"
			} else {
				param.Value = p[1]
			}
			cmdInput.parameters[param.Name] = param
			i++
			continue
		}

		// short form of parameter or flag
		if part[0] == '-' {
			f := part[1:]

			// flag group
			if len(f) > 1 {
				for _, c := range f {
					flag, ok := definition.GetParameter(string(c))
					if !ok {
						return &commandInput{}, UnknownParameterError.New(string(c), cmd)
					}

					if !flag.IsFlag {
						return &commandInput{}, UnexpectedParameterError.New(string(c), cmd)
					}

					flag.Value = "1"
					cmdInput.parameters[flag.Name] = flag
				}
				i++
				continue
			}

			parameter, ok := definition.GetParameter(f)

			if !ok {
				return &commandInput{}, UnknownParameterError.New(f, cmd)
			}

			if parameter.IsFlag {
				parameter.Value = "1"
				cmdInput.parameters[parameter.Name] = parameter
				i++
				continue
			}

			if i + 1 >= iMax {
				return &commandInput{}, MissingParameterValueError.New(f)
			}

			parameter.Value = parts[i+1]
			cmdInput.parameters[parameter.Name] = parameter
			i += 2
			continue
		}

		if arg, ok := definition.GetArgument(argIndex); ok {
			arg.Value = part
			if cmdArg, ok := cmdInput.arguments[arg.Name]; ok && arg.Wildcard {
				arg.Value = cmdArg.Value + " " + part
			}

			cmdInput.arguments[arg.Name] = arg
		} else {
			return &commandInput{}, UnexpectedArgumentError.New(part, cmd)
		}

		argIndex++
		i++
	}
	// add default parameters
	for _, parameter := range definition.parameters {
		if _, ok := cmdInput.parameters[parameter.Name]; !ok && len(parameter.DefaultValue) > 0 {
			parameter.Value = parameter.DefaultValue
			cmdInput.parameters[parameter.Name] = parameter
		}
	}

	return cmdInput, nil
}

func (c *Command) Execute(cmd string, terminal *Terminal) error {
	input, err := NewCommandInput(cmd, c.Definition)
	if err != nil {
		return CommandError.New(cmd).CausedBy(err)
	}
	err = c.Handler(terminal, input.arguments, input.parameters)
	if err != nil {
		return CommandError.New(cmd).CausedBy(err)
	}

	return nil
}
