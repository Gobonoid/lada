package main

import lada "github.com/kodemore/lada/pkg"

func main() {
	app, _ := lada.NewApplication("simple_example", "1.0.0")
	app.AddCommand("hello $name", func(t *lada.Terminal, args lada.Arguments, params lada.Options) error {
		t.Printf("hello, %s!", args["name"].AsString())
		return nil
	})
	app.AddCommand("goodbye", func(t *lada.Terminal, args lada.Arguments, params lada.Options) error {
		t.Print("goodbye world!")
		return nil
	})

	app.Run()
}