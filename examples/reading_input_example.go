package main

import lada "github.com/kodemore/lada/pkg"

func main() {
	app, _ := lada.NewApplication("input test", "1.0.0")

	app.AddCommand("hello", func(t *lada.Terminal, a lada.Arguments, o lada.Options) error {
		message, _ := t.Prompt("Put your hello message:")
		t.Printf("Your message is: %s", message)
		return nil
	})

	app.Run()
}