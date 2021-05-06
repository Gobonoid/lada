package examples
/*
cli = lada.NewCli(name: "whatever", version: "1.0.0")

// Magic implementation

constraints = NewConstraints()
constraints.SetParameterContrain("flavour", Constrain.StringEnum("cherry", "apple"))
constraints.SetParameterContrain("num", Contrain.IntegerRange(8, 2))

handler = func(args, parameters, flags) {

}
cli.AddCommandWithConstraints("dsl template", handler, contrains)

command := cli.AddCommand("subcommand $input $output --num=1 --verbose? --flavour= --word=default --flag?", func (arguments, flags) {
	cli.WriteLine("Ass hole")
	cli.ShowPrompt("Are you sure", ["yes"])
})
command.SetParameterContraint("flavour", Constraint.StringEnum("cherry", "apple"))
command.SetParameterContraint("num", Contraint.IntegerRange(8, 2))


// Verbose implementation

myCmd := cli.NewCommand("subcommand")
myCmd.AddArgs(["inputfile", "outputfile"])
myCmd.AddParams([
	// name, min, max, default
	cli.ParamInt("num", 1, 8, 2),
	// name, default
	cli.ParamBool("verbose", false),
	// name, array of strings
	cli.ParamEnum("flavour", ["peanut", "cherry", "apple"])
	// name, default
	cli.ParamString("word", "any")
])
// name
myCmd.addFlag("flag")

myCmd.Handler(func(command myCmd) {
	args := myCmd.getArgs()
	params := myCmd.getParams()

	// Command logic here

	myCmd.WriteLine("command complete")
})

// Subcommands
// lada generate

// Arguments
// lada generate ./filename.txt --split=,

// Parameters
// int --num=2
// bool --verbose=true
// enum --flavour=[peanut,cherry,apple]
// string --word=any

// Flags
// --flag

 */