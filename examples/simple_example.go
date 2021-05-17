package main

import lada "github.com/kodemore/lada/pkg"

func main() {
	cli, _ := lada.NewCli("simple_example", "1.0.0")
	cli.AddCommand("hello", func(args lada.Arguments, params lada.Parameters) error {
		cli.Write("hello world!")
		return nil
	})

	cli.AddCommand("goodbye", func(args lada.Arguments, params lada.Parameters) error {
		cli.Write("goodbye world!")
		return nil
	})

	cli.Run()
}