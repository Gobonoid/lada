package main

import (
	"github.com/kodemore/lada/pkg"
	"github.com/kodemore/lada/pkg/style"
)

func main() {
	app, _ := lada.NewApplication("Print your messages with style", "1.0.0")
	app.AddCommand("styled", func(t *lada.Terminal, _ lada.Arguments) error {
		t.Print("normal text")
		t.PrettyPrint("gray on red", style.Background.Red, style.Foreground.LightGray)
		t.PrettyPrint("bold magenta", style.Foreground.LightMagenta, style.Format.Bold)
		t.PrettyPrint("a blinking text", style.Format.Blink)
		return nil
	})

	app.Run()
}
