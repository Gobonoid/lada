package lada

type Arguments map[string]Argument
type Parameters map[string]Parameter

func (p Parameters) GetValue(name string) (interface{}, error) {
	return p[name].Value, nil
}

func (p Parameters) GetAsInt(name string) (int, error) {
	return 0, nil
}

func (p Parameters) GetAsString(name string) (string, error) {
	return "", nil
}

func (p Parameters) GetAsFloat(name string) (float32, error) {
	return 0.0, nil
}

func (p Parameters) GetAsBool(name string) (bool, error) {
	return false, nil
}

func (p Parameters) GetAsIntEnum(name string, enum map[string]int) (int, error) {
	return 0, nil
}

func (p Parameters) GetAsStringEnum(name string, enum map[string]string) (string, error) {
	return "", nil
}

func (p Parameters) GetAsStringArray(name string) ([]string, error) {
	return make([]string, 0), nil
}

func (p Parameters) GetAsIntArray(name string) ([]int, error) {
	return make([]int, 0), nil
}

func (p Parameters) GetAsFloatArray(name string) ([]float32, error) {
	return make([]float32, 0), nil
}

func (p Parameters) IsEnabled(name string) (bool, error) {
	return true, nil
}

type Handler func(args Arguments, params Parameters) error

type Command struct {
	Name       string
	Handler    Handler
	Definition *CommandDefinition
}

func NewCommand(definition string, handler Handler) (*Command, error) {
	cmdDefinition, err := NewCommandDefinition(definition)
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
	arguments Arguments
	parameters Parameters
}

func NewCommandInput(cmd string, definition *CommandDefinition) (*commandInput, error) {
	parts := splitCommandStringToParts(cmd)
	if parts[0] != definition.Command() {
		return &commandInput{}, CommandNameDoesNotMatchError.New(parts[0], definition.Command())
	}
	cmdInput := &commandInput{
		commandName: parts[0],
		arguments: make(Arguments),
	}
	argIndex := 0
	i := 1
	iMax := len(parts)
	for i < iMax {
		part := parts[i]

		// long form for parameter or flag
		if part[0:2] == "--" {
			
			continue
		}
		// short form
		if part[0] == '-' {

			// check for grouping
			continue
		}

		if arg, ok := definition.GetArgument(argIndex); ok {
			arg.Value = part
			if cmdArg, ok := cmdInput.arguments[arg.Name]; ok && arg.Wildcard {
				arg.Value = cmdArg.Value + " " + part
			}

			cmdInput.arguments[arg.Name] = arg
		} else {
			return &commandInput{}, UnexpectedArgument.New(part, cmd)
		}

		argIndex++
		i++
	}

	return cmdInput, nil
}

func (c *Command) Execute(cmd string) error {

	return nil
}
