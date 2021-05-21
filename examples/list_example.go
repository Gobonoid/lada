package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("simple_example", "1.0.0")
	app.AddCommand("hello", func(t *l.Terminal, args l.Arguments, params l.Options) error {
		//value, err := t.SelectList("Pick an item", "item 1", "item 2", "item 3")

		return nil
	})

	app.Run()
}

