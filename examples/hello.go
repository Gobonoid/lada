package main

import (
	"github.com/kodemore/lada/pkg"
)

func HelloCommand(t *lada.Terminal, args lada.Arguments) error {
	t.Printf("hello, %s!", args.Get("name").Value())
	return nil
}

func GoodbyeCommand(t *lada.Terminal, args lada.Arguments) error {
	t.Print("goodbye world!")
	return nil
}

func main() {
	app, _ := lada.NewApplication("Hello Application", "1.0.0")
	app.AddCommand("hello $name", HelloCommand)
	app.AddCommand("goodbye", GoodbyeCommand)
	app.Run()
}