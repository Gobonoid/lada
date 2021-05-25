package lada

import (
	"fmt"
	"strings"
)

type Arguments struct {
	raw       string
	arguments CommandPatternArguments
	values    []ArgumentValue
}

func (a *Arguments) Get(s string) ArgumentValue {
	for _, arg := range a.values {
		if arg.argument.Name == s {
			return arg
		}
	}

	for _, arg := range a.arguments {
		if arg.Name == s {
			return ArgumentValue{arg, ""}
		}
	}

	panic(fmt.Sprintf("Trying to reach non-existent argument %s", s))
}

func NewArguments(s string, arguments CommandPatternArguments) (Arguments, error) {
	args := Arguments{
		raw: s,
		arguments: arguments,
		values: make([]ArgumentValue, 0),
	}
	err := args.parse()
	if err != nil {
		return Arguments{}, err
	}

	return args, nil
}


func (a *Arguments) parse() error {
	var wildcardArg ArgumentValue
	positionalArguments := a.arguments.GetPositionalArguments()
	parts := splitArgumentsString(a.raw)
	positionalArgIndex := 0

	for i := 0; i < len(parts); i++ {
		part := parts[i]
		// long form optional arg
		if len(part) > 1 && part[0:2] == "--" {
			kv := strings.Split(part[2:], "=")
			if arg, ok := a.arguments.GetArgumentByName(kv[0]); ok {
				if len(kv) > 1 && arg.Kind() == FlagArgument {
					return UnexpectedArgumentValue.New(arg.Name)
				}
				if arg.Kind() == FlagArgument {
					a.values = append(a.values, NewArgumentValue(arg, "1"))
					continue
				}
				if len(kv) < 2 {
					return MissingArgumentValueError.New(arg.Name)
				}
				a.values = append(a.values, NewArgumentValue(arg, kv[1]))
			} else {
				return UnknownArgumentError.New(kv[0])
			}
			continue
		}

		// short form optional arg
		if part[0] == '-' {

			// multiple flags at once
			if len(part) > 2 {
				for _, c := range part[1:] {
					if arg, ok := a.arguments.GetArgumentByShortName(string(c)); ok {
						if arg.Kind() == FlagArgument {
							a.values = append(a.values, NewArgumentValue(arg, "1"))
						// we dont expect here non flag arguments
						} else {
							return UnexpectedArgumentError.New(arg.Name, a.raw)
						}
					// argument is not found in the pattern
					} else {
						return UnknownArgumentError.New(string(c))
					}
				}
				continue
			}
			// single flag or optional argument
			if arg, ok := a.arguments.GetArgumentByShortName(string(part[1])); ok {
				switch arg.Kind() {
				case FlagArgument:
					a.values = append(a.values, NewArgumentValue(arg, "1"))
				case OptionalArgument:
					// we have to pick the value from next item
					// if we run out of scope we should return an error
					if i + 1 >= len(parts) {
						return MissingArgumentValueError.New(arg.Name)
					}
					a.values = append(a.values, NewArgumentValue(arg, parts[i+1]))
					// skip next item as it is already appended to args as value
					i++
					continue

				default:
					return UnknownArgumentError.New(string(part[1]))
				}
			} else {
				return UnknownArgumentError.New(string(part[1]))
			}
			continue
		}

		// positional argument
		if positionalArgIndex >= len(positionalArguments) {
			if wildcardArg.argument != nil {
				wildcardArg.value += " " + part
				continue
			} else {
				return UnexpectedArgumentError.New(fmt.Sprintf("arg#%d", positionalArgIndex), a.raw)
			}
		}
		arg := a.arguments[positionalArgIndex]
		if arg.IsWildcard() {
			wildcardArg = NewArgumentValue(arg, part)
			a.values = append(a.values, wildcardArg)
		} else {
			a.values = append(a.values, NewArgumentValue(arg, part))
		}

		positionalArgIndex++
	}

	if len(positionalArguments) > positionalArgIndex {
		return MissingArgumentValueError.New(fmt.Sprintf("arg#%d", positionalArgIndex))
	}

	// copy al wildcard values to correct index
	for index, arg := range a.values {
		if arg.argument.IsWildcard() {
			a.values[index] = wildcardArg
		}
	}

	return nil
}