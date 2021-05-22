package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("simple_example", "1.0.0")
	app.AddCommand("hello", func(t *l.Terminal, args l.Arguments, params l.Options) error {

		items := []string{"item 1", "item 2", "item 3"}
		list := l.NewSelectUI("Pick an Item", items)
		t.Display(list)
		t.Printf("You have selected item: %s", items[list.Value()])

		return nil
	})

	app.Run()
}

