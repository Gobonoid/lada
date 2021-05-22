package main

import (
	lada "github.com/kodemore/lada/pkg"
)

type C struct {}
func (c C) OnKey(t *lada.Terminal, key lada.Key) bool {

	switch key.Type() {
	case lada.KeyArrowDown:
		t.Println("ArrowDown")
	case lada.KeyArrowUp:
		t.Println("ArrowUp")
	case lada.KeyArrowLeft:
		t.Println("ArrowLeft")
	case lada.KeyArrowRight:
		t.Println("ArrowRight")
	case lada.KeySpace:
		t.Println("Space")
	case lada.KeyBackspace:
		t.Println("Backspace")
	case lada.KeyEnter:
		t.Println("Enter")
	case lada.KeyEscape:
		t.Println("Escape")
		t.Printf("Rune:%d\n", key)
	case lada.KeyTab:
		t.Println("Tab")
	default:
		t.Printf("Rune:%s\n", key.AsString())

		// stop capturing keys
		if key.Equals('q') {
			return false
		}
	}

	// continue capturing keys
	return true
}

func main() {
	t, _ := lada.NewTerminal()
	t.Println("This demo showcases capturing of terminal's keyboard")
	t.CaptureKeys(C{})
}