package lada

import (
	"regexp"
	"strings"
)

type ArgumentsHelp map[string]string

type ArgumentKind int

const (
	PositionalArgument ArgumentKind = iota
	OptionalArgument
	FlagArgument
)

var argumentNamePattern = regexp.MustCompile(`^(?P<long>[a-zA-Z][a-zA-Z0-9-]+)(?P<short>\[([a-zA-Z])\])?$`)

func parseArgumentName(str string) (map[string]string, error) {
	results := map[string]string{}
	match := argumentNamePattern.FindStringSubmatch(str)
	if match == nil {
		return results, InvalidArgumentNameError.New(str)
	}

	for i, name := range match {
		results[argumentNamePattern.SubexpNames()[i]] = name
	}
	return results, nil
}

type Argument struct {
	Name         string
	Description  string
	ShortName    string
	wildcard     bool
	defaultValue string
	kind         ArgumentKind
}

func NewArgumentFromCommandPattern(p string) (*Argument, error) {
	// --option[O]=default value
	// argument...
	// --option[O]
	arg := &Argument{}

	if len(p) > 2 && p[0:2] == "--" {
		p = p[2:]
		kv := strings.Split(p, "=")
		argName, err := parseArgumentName(kv[0])
		if err != nil {
			return &Argument{}, err
		}
		arg.Name = argName["long"]
		if argName["short"] != "" {
			shortName := argName["short"][1:len(argName["short"])-1]
			arg.ShortName = shortName
		}
		if len(kv) > 1 {
			arg.kind = OptionalArgument
			arg.defaultValue = kv[1]
		} else {
			arg.kind = FlagArgument
		}
		return arg, nil
	}
	if len(p) > 3 && p[len(p) - 3:] == "..." {
		arg.wildcard = true
		p = p[0:len(p) - 3]
	}

	// positional arguments must start with $ sign
	if p[0] != '$' {
		return &Argument{}, InvalidArgumentNameError.New(p)
	}
	argName, err := parseArgumentName(p[1:])
	if err != nil {
		return &Argument{}, err
	}
	arg.Name = argName["long"]
	arg.kind = PositionalArgument

	return arg, nil
}

func (a *Argument) Kind() ArgumentKind {
	return a.kind
}

func (a *Argument) DefaultValue() string {
	if a.kind == FlagArgument {
		return "0"
	}

	return a.defaultValue
}

func (a *Argument) IsWildcard() bool {
	if a.kind == PositionalArgument {
		return a.wildcard
	}

	return false
}

func (a *Argument) IsOptional() bool {
	return !a.IsPositional()
}

func (a *Argument) IsPositional() bool {
	return a.kind == PositionalArgument
}