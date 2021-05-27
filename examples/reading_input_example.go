package main

import (
	"github.com/kodemore/lada/pkg"
)

func main() {
	app, _ := lada.NewApplication("Input Test", "1.0.0")
	app.AddCommand("input", func(t *lada.Terminal, _ lada.Arguments) error {
		message, _ := t.Prompt("What's your name:")
		secret, _ := t.Secret("Tell me your secret")
		t.Printf("Your name is: %s", message)
		t.Printf("Your secret is: %s", secret)
		return nil
	})
	app.Run()
}