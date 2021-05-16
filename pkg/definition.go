package lada

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// command $argument $argument-list* --parameter-without-default[P]= --parameter-with-default=some\ value --flag[F]

type Argument struct {
	Name        string
	Wildcard    bool
	Description string
	Value 		string
}

type Parameter struct {
	Name         string
	ShortForm    string
	DefaultValue string
	Description  string
	Value        string
	IsFlag       bool
}

type CommandDefinition struct {
	raw         string
	parts       []string
	commandName string
	parameters  []Parameter
	arguments   []Argument
	Description string
}

func (c *CommandDefinition) Command() string {
	return c.commandName
}

func (c *CommandDefinition) GetArgument(position int) (Argument, bool) {
	if position >= len(c.arguments) {
		if c.arguments[len(c.arguments) - 1].Wildcard {
			return c.arguments[len(c.arguments) - 1], true
		}
		return Argument{}, false
	}

	return c.arguments[position], true
}

func (c *CommandDefinition) GetParameter(name string) (Parameter, bool) {
	for _, p := range c.parameters {
		if p.ShortForm == name {
			return p, true
		}

		if p.Name == name {
			return p, true
		}
	}
	return Parameter{}, false
}

func NewCommandDefinition(definition string) (*CommandDefinition, error) {
	command := &CommandDefinition{
		raw: definition,
	}
	command.parts = splitCommandStringToParts(definition)
	if err:= command.Parse(); err != nil {
		return &CommandDefinition{}, err
	}

	return command, nil
}

func splitCommandStringToParts(template string) []string {
	result := make([]string, 0)
	parts := strings.Split(template, " ")
	escaped := false
	for _, part := range parts {
		if part == "" {
			continue
		}
		resultLength := len(result)
		if escaped {
			result[resultLength-1] += " " + part
		} else {
			result = append(result, part)
			resultLength += 1
		}
		escaped = false

		if part[len(part)-1] == '\\' {
			escaped = true
			result = result[:resultLength-1]
			result = append(result, part[0:len(part)-1])
		}
	}
	// trim whitespace from each item in result
	for index, item := range result {
		result[index] = strings.TrimSpace(item)
	}
	return result
}

var parameterNameRegex = regexp.MustCompile(`^(?P<long>[a-z][a-z-0-9-]+)(?P<short>\[([a-zA-Z])\])?$`)

func matchParameterName(str string) (map[string]string, error) {
	results := map[string]string{}
	match := parameterNameRegex.FindStringSubmatch(str)
	if match == nil {
		return results, errors.New(
			fmt.Sprintf(
				"`%s` does not conform name pattern `([a-z][a-z-0-9]+)`",
				str,
			),
		)
	}

	for i, name := range match {
		results[parameterNameRegex.SubexpNames()[i]] = name
	}
	return results, nil
}

func makeFlag(value string) (Parameter, error) {
	parts, err := matchParameterName(value)
	if err != nil {
		return Parameter{}, InvalidCommandIdentifierError.New(value).CausedBy(err)
	}

	flag := Parameter{Name: parts["long"], IsFlag: true}
	if short, ok := parts["short"]; ok && len(parts["short"]) > 1 {
		flag.ShortForm = string(short[1])
	}

	return flag, nil
}

func makeParameter(value string) (Parameter, error) {
	p := strings.Split(value, "=")
	parts, err := matchParameterName(p[0])
	if err != nil {
		return Parameter{}, InvalidCommandIdentifierError.New(value).CausedBy(err)
	}

	parameter := Parameter{Name: parts["long"], IsFlag: false}
	if short, ok := parts["short"]; ok && len(parts["short"]) > 1 {
		parameter.ShortForm = string(short[1])
	}

	if len(p) > 1 {
		parameter.DefaultValue = p[1]
	}

	return parameter, nil
}

func (c *CommandDefinition) Parse() error {
	c.arguments = make([]Argument, 0)
	c.parameters = make([]Parameter, 0)
	wildCardArgPresent := false

	for _, item := range c.parts {
		if item[0:2] == "--" {
			// Flag
			if !strings.ContainsRune(item, '=') {
				flag, err := makeFlag(item[2:])
				if err != nil {
					return CommandDefinitionParseError.CausedBy(err)
				}
				c.parameters = append(c.parameters, flag)
				continue
			}

			// Parameter
			parameter, err := makeParameter(item[2:])
			if err != nil {
				return CommandDefinitionParseError.CausedBy(err)
			}
			c.parameters = append(c.parameters, parameter)
			continue
		}

		// arguments
		if item[0] == '$' {
			argument, err := makeArgument(item)
			if err != nil {
				return CommandDefinitionParseError.CausedBy(err)
			}

			if wildCardArgPresent {
				return CommandDefinitionParseError.CausedBy(UnexpectedWildcardArgumentError.New(item, c.raw))
			}

			if argument.Wildcard {
				wildCardArgPresent = true
			}

			c.arguments = append(c.arguments, argument)
			continue
		}

		// command name
		if c.commandName != "" {
			return CommandDefinitionParseError.CausedBy(UnexpectedCommandParameterError.New(item, c.raw))

		}
		if !parameterNameRegex.MatchString(item) {
			return CommandDefinitionParseError.CausedBy(InvalidCommandIdentifierError.New(item, c.raw))
		}
		c.commandName = item
	}

	return nil
}

func makeArgument(value string) (Argument, error) {
	argument := Argument{}
	if value[len(value) - 1] == '*' {
		name := value[1:len(value) - 1]
		if !parameterNameRegex.MatchString(name) {
			return argument, CommandDefinitionParseError.CausedBy(InvalidCommandIdentifierError.New(value))
		}
		argument.Wildcard = true
		argument.Name = name
		return argument, nil
	}
	name := value[1:]
	if !parameterNameRegex.MatchString(name) {
		return argument, CommandDefinitionParseError.CausedBy(InvalidCommandIdentifierError.New(value))
	}
	argument.Wildcard = false
	argument.Name = name
	return argument, nil
}
