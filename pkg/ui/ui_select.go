package ui

import (
	"github.com/kodemore/lada/pkg"
	"github.com/kodemore/lada/pkg/style"
)

const slDisabledItem = "◌"
const slActiveItem = "○"
const slSelectedItem = "⦿"

type Select struct {
	Label      string
	Items      []string
	activeItem int
	height     int
	drawn      bool
}

func NewSelect(label string, items []string) *Select {
	l := &Select{
		Label:      label,
		Items:      items,
		activeItem: 0,
		height:     len(items) + 1,
		drawn:      false,
	}

	return l
}

func (s *Select) Display(t *lada.Terminal) error {
	t.Println(s.Label)

	for index, item := range s.Items {
		if index == s.activeItem {
			t.PrettyPrint("  " +slSelectedItem+ " ", style.Foreground.LightBlue)
			t.PrettyPrint(item, style.Foreground.LightBlue, style.Format.Underline)
			t.Print("\n")
			continue
		}
		t.Print("  " + slActiveItem + " ")
		t.Print(item)
		t.Print("\n")
	}

	return nil
}

func (s *Select) Value() int {
	return s.activeItem
}

func (s *Select) Remove(t *lada.Terminal) error {
	t.Cursor.MoveUp(s.height - 1)
	t.Cursor.EraseLine()
	return nil
}

func (s *Select) refresh(t *lada.Terminal) {
	t.Cursor.MoveUp(s.height)
	s.Display(t)
}

func (s *Select) OnKey(t *lada.Terminal, k lada.Key) bool {
	switch k.Type() {
	case lada.KeyArrowUp:
		s.activeItem--
	case lada.KeyArrowDown:
		s.activeItem++
	case lada.KeyEnter:
		return false
	}
	if s.activeItem < 0 {
		s.activeItem = len(s.Items) - 1
	} else if s.activeItem >= len(s.Items) {
		s.activeItem = 0
	}
	s.refresh(t)
	return true
}