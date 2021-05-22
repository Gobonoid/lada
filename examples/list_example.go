package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("Simple Application", "1.0.0")
	app.Description = "This application showcases how select list works"

	app.AddCommand("list-demo", "runs list demo", func(t *l.Terminal, args l.Arguments, params l.Options) error {
		items := []string{"item 1", "item 2", "item 3", "exit"}
		list := l.NewSelectUI("Pick an Item", items)
		for {
			t.Display(list)
			t.Printf("You have selected item: %s \n", items[list.Value()])

			if list.Value() == 3 {
				break
			}
		}
		return nil
	})

	app.Run()
}

