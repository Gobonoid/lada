package lada

import (
	"strings"
)

type CommandPattern struct {
	raw         string
	verb        string
	Arguments   CommandPatternArguments
}

func NewCommandPattern(pattern string) (*CommandPattern, error) {
	command := &CommandPattern{
		raw: pattern,
	}
	rawArgs := strings.Split(pattern, " ")
	command.verb = rawArgs[0]
	args, err := NewCommandPatternArguments(strings.Join(rawArgs[1:], " "))
	if err != nil {
		return &CommandPattern{}, err
	}
	command.Arguments = args

	return command, nil
}

func (c *CommandPattern) Verb() string {
	return c.verb
}

func (c *CommandPattern) IsCatchAll() bool {
	return c.verb == "*"
}

func splitArgumentsString(format string) []string {
	result := make([]string, 0)
	parts := strings.Split(format, " ")
	escaped := false
	for _, part := range parts {
		if part == "" || part == "\n" {
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