package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("simple_example", "1.0.0")
	app.AddCommand("hello", func(t *l.Terminal, args l.Arguments, params l.Options) error {

		items := []string{"item 1", "item 2", "item 3"}
		value, _ := t.SelectList("Pick an item", items)
		t.Printf("You have selected item: %s", items[value])

		return nil
	})

	app.Run()
}

