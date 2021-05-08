package lada

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// command $argument ...$argument-list --parameter-without-default[P]= --parameter-with-default=some\ value --flag[F]

type Argument struct {
	Name        string
	Wildcard    bool
	Description string
}

type Parameter struct {
	LongForm     string
	ShortForm    string
	DefaultValue string
	Description  string
}

type Flag struct {
	LongForm    string
	ShortForm   string
	Description string
}

type CommandTemplate struct {
	raw         string
	parts       []string
	commandName string
	parameters  []*Parameter
	flags       []*Flag
	arguments   []*Argument
	Description string
}

func NewCommandTemplate(template string) *CommandTemplate {
	command := &CommandTemplate{
		raw: template,
	}
	command.parts = splitCommandTemplateStringIntoParts(template)
	return command
}

func (c *CommandTemplate) Command() string {
	return c.commandName
}

func splitCommandTemplateStringIntoParts(template string) []string {
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

func makeFlag(value string) (*Flag, error) {
	parts, err := matchParameterName(value)
	if err != nil {
		return &Flag{}, InvalidCommandTemplateIdentifierError.Sprintf(value).CausedBy(err)
	}

	flag := &Flag{LongForm: parts["long"]}
	if short, ok := parts["short"]; ok && len(parts["short"]) > 1 {
		flag.ShortForm = string(short[1])
	}

	return flag, nil
}

func makeParameter(value string) (*Parameter, error) {
	p := strings.Split(value, "=")
	parts, err := matchParameterName(p[0])
	if err != nil {
		return &Parameter{}, InvalidCommandTemplateIdentifierError.Sprintf(value).CausedBy(err)
	}

	parameter := &Parameter{LongForm: parts["long"]}
	if short, ok := parts["short"]; ok && len(parts["short"]) > 1 {
		parameter.ShortForm = string(short[1])
	}

	if len(p) > 1 {
		parameter.DefaultValue = p[1]
	}

	return parameter, nil
}

func (c *CommandTemplate) Parse() error {
	c.flags = make([]*Flag, 0)
	c.arguments = make([]*Argument, 0)
	c.parameters = make([]*Parameter, 0)

	for _, item := range c.parts {
		if item[0:2] == "--" {
			// Flag
			if !strings.ContainsRune(item, '=') {
				flag, err := makeFlag(item[2:])
				if err != nil {
					return CommandTemplateParseError.CausedBy(err)
				}
				c.flags = append(c.flags, flag)
				continue
			}

			// Parameter
			parameter, err := makeParameter(item[2:])
			if err != nil {
				return CommandTemplateParseError.CausedBy(err)
			}
			c.parameters = append(c.parameters, parameter)
			continue
		}

		// arguments
		if item[0] == '$' || (len(item) > 4 && item[0:4] == "...$") {
			argument, err := makeArgument(item)
			if err != nil {
				return CommandTemplateParseError.CausedBy(err)
			}
			c.arguments = append(c.arguments, argument)
			continue
		}

		// command name
		if c.commandName != "" {
			return CommandTemplateParseError.CausedBy(UnexpectedCommandTemplateParameterError.Sprintf(item, c.raw))

		}
		if !parameterNameRegex.MatchString(item) {
			return CommandTemplateParseError.CausedBy(InvalidCommandTemplateIdentifierError.Sprintf(item, c.raw))
		}
		c.commandName = item
	}

	return nil
}

func makeArgument(value string) (*Argument, error) {
	argument := &Argument{}
	if value[0:4] == "...$" {
		if !parameterNameRegex.MatchString(value[4:]) {
			return argument, CommandTemplateParseError.CausedBy(InvalidCommandTemplateIdentifierError.Sprintf(value))
		}
		argument.Wildcard = true
		argument.Name = value[4:]
		return argument, nil
	}
	if !parameterNameRegex.MatchString(value[1:]) {
		return argument, CommandTemplateParseError.CausedBy(InvalidCommandTemplateIdentifierError.Sprintf(value))
	}
	argument.Wildcard = false
	argument.Name = value[1:]
	return argument, nil
}
