package lada

import "strings"

type CommandPatternArguments []*Argument

func (a CommandPatternArguments) GetArgumentByName(name string) (*Argument, bool) {
	for _, arg := range a {
		if arg.Name == name {
			return arg, true
		}
	}

	return &Argument{}, false
}

func (a CommandPatternArguments) GetArgumentByShortName(name string) (*Argument, bool) {
	for _, arg := range a {
		if arg.ShortName == name {
			return arg, true
		}
	}

	return &Argument{}, false
}

func (a CommandPatternArguments) GetPositionalArguments() []*Argument {
	args := make([]*Argument, 0)
	for _, arg := range a {
		if arg.Kind() != PositionalArgument {
			continue
		}
		args = append(args, arg)
	}

	return args
}

func NewCommandPatternArguments(s string) (CommandPatternArguments, error) {
	items := splitArgumentsString(s)
	arguments := make(CommandPatternArguments, 0)
	argNames := make(map[string]struct{}, 0)

	hasWildcardArgument := false
	for _, item := range items {
		arg, err := NewArgumentFromCommandPattern(item)
		if err != nil {
			return make([]*Argument, 0), err
		}
		if _, exists := argNames[arg.Name]; exists {
			return make([]*Argument, 0), ArgumentAlreadyDefinedError.New(arg.Name, strings.Join(items, " "))
		}
		arguments = append(arguments, arg)
		if hasWildcardArgument && arg.IsWildcard() {
			return make([]*Argument, 0), UnexpectedWildcardArgumentError.New(arg.Name, strings.Join(items, " "))
		}
		if arg.IsWildcard() {
			hasWildcardArgument = true
		}
	}

	return arguments, nil
}
