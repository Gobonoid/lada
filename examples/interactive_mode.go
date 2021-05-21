package main

import (
	//"fmt"
	lada "github.com/kodemore/lada/pkg"
)

func main() {
	t, _ := lada.NewTerminal()

	t.SetInteractiveMode()


	t.Println("Kuniec")
	t.StopInteractiveMode()
	a, _ := t.Prompt("?")
	t.Print(a)
}