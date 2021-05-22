package lada

type UIKeyboardListener interface {
	OnKey(t *Terminal, key Key) bool
}

type UIElement interface {
	Display(t *Terminal) error
}

type UIRemovable interface {
	Remove(t *Terminal) error
}