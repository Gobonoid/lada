package main

import (
	"fmt"
	lada "github.com/kodemore/lada/pkg"
)

func main() {
	app, _ := lada.NewApplication("input test", "1.0.0")
	app.AddCommand("hello --name[n]=", func(t *lada.Terminal, a lada.Arguments, o lada.Options) error {
		name, _ := o["name"].AsString()
		t.PrettyPrint(fmt.Sprintf("Hello %s", name), lada.Background.MAGENTA)
		message, _ := t.Prompt("What's your name:")
		secret, _ := t.Secret("Tell me your secret")
		t.Printf("Your name is: %s", message)
		t.Printf("Your secret is: %s", secret)

		return nil
	})
	app.Run()
}