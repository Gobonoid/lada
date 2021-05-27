package main

import lada "github.com/kodemore/lada/pkg"

// PositionalArgumentsExample
// Can be executed by running: go run arguments.go positional 1 1
func PositionalArgumentsExample(t *lada.Terminal, arguments lada.Arguments) error {
	t.Printf("Argument #1: %s \n", arguments.Get("arg1").Value())
	t.Printf("Argument #2: %s \n", arguments.Get("arg2").Value())
	return nil
}

// OptionalArgumentsExample
// can be executed by running:
// - go run arguments.go -o 1 -O 2
// - go run arguments.go --optional1=1 --optional2=2
func OptionalArgumentsExample(t *lada.Terminal, arguments lada.Arguments) error {
	t.Printf("Optional #1: %s \n", arguments.Get("optional1").Value())
	t.Printf("Optional #2: %s \n", arguments.Get("optional2").Value())
	return nil
}

func MixedArgumentsExample(t *lada.Terminal, arguments lada.Arguments) error {
	t.Printf("Optional: %s \n", arguments.Get("optional").Value())
	t.Printf("Flag: %s \n", arguments.Get("flag").Value())
	t.Printf("Arg#1: %s \n", arguments.Get("arg1").Value())
	t.Printf("Arg#N: %s \n", arguments.Get("argN").Value())
	return nil
}

func main() {
	app,_ := lada.NewApplication("Arguments example", "1.0.0")
	app.Description = "Example app to show how arguments are working"

	app.AddCommand("positional $arg1 $arg2", PositionalArgumentsExample)
	app.AddCommand("optional --optional1[o]= --optional2[O]=", OptionalArgumentsExample)

	app.AddHelp(
		"positional",
		"Show how positional arguments work",
		lada.ArgumentsHelp{
			"arg1": "argument 1",
			"arg2": "argument 2",
		},
	)

	app.AddHelp("optional",
		"show how optional arguments work",
		lada.ArgumentsHelp{
		"optional1": "optional 1",
		"optional2": "optional 2",
	})

	app.AddCommand("mixed --optional[o]=default\\ value --flag[f] $arg1 $argN...", MixedArgumentsExample)

	app.AddHelp("mixed",
		"show how mix arguments together",
		lada.ArgumentsHelp{
			"optional": "optional argument with a value",
			"flag": "optional flag argument",
			"arg1": "argument on 1st position",
			"argN": "variadic argument",
		})

	app.Run()
}
