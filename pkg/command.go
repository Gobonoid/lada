package lada

type Handler func(terminal *Terminal, arguments Arguments) error

type Command struct {
	Pattern     *CommandPattern
	Handler     Handler
	Description string
}

func  (c *Command) IsHelpAvailable() bool {
	return c.Description != ""
}

func (c *Command) Verb() string {
	return c.Pattern.Verb()
}

func NewCommand(pattern string, handler Handler) (*Command, error) {
	cPattern, err := NewCommandPattern(pattern)
	if err != nil {
		return &Command{}, err
	}

	cmd := &Command{
		Handler: handler,
		Pattern: cPattern,
	}

	return cmd, nil
}

func (c *Command) Execute(args string, terminal *Terminal) error {
	input, err := NewArguments(args, c.Pattern.Arguments)
	if err != nil {
		return err
	}
	err = c.Handler(terminal, input)
	if err != nil {
		return CommandError.New(c.Pattern.raw).CausedBy(err)
	}

	return nil
}