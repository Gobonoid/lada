package main

import l "github.com/kodemore/lada/pkg"

func main() {
	app, _ := l.NewApplication("Print your messages with style", "1.0.0")
	app.AddCommand("style", "demo of styled output", func(t *l.Terminal, args l.Arguments, params l.Options) error {
		t.Print("normal text")
		t.PrettyPrint("gray on red", l.Background.Red, l.Foreground.LightGray)
		t.PrettyPrint("bold magenta", l.Foreground.LightMagenta, l.Format.Bold)
		t.PrettyPrint("a blinking text", l.Format.Blink)
		return nil
	})

	app.Run()
}
