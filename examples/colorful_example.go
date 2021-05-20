package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("simple_example", "1.0.0")
	app.AddCommand("hello", func(t *l.Terminal, args l.Arguments, params l.Options) error {
		t.Print("normal text")
		t.PrettyPrint("gray on red", l.Background.RED, l.Foreground.LIGHTGRAY)
		t.PrettyPrint("bold magenta", l.Foreground.LIGHTMAGENTA, l.Format.BOLD)

		return nil
	})

	app.Run()
}
