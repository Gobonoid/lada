package lada

const slDisabledItem = "◌"
const slActiveItem = "○"
const slSelectedItem = "⦿"

type SelectList struct {
	Label      string
	Items      []string
	activeItem int
	height     int
	drawn      bool
}

func NewSelectList(label string, items []string) *SelectList {
	l := &SelectList{
		Label:      label,
		Items:      items,
		activeItem: 0,
		height:     len(items) + 1,
		drawn:      false,
	}

	return l
}

func (s *SelectList) Display(t *Terminal) {
	t.Println(s.Label)

	for index, item := range s.Items {
		if index == s.activeItem {
			t.PrettyPrint("  " + slSelectedItem + " ", Foreground.LightBlue)
			t.PrettyPrint(item, Foreground.LightBlue, Format.Underline)
			t.Print("\n")
			continue
		}
		t.Print("  " + slActiveItem + " ")
		t.Print(item)
		t.Print("\n")
	}
}

func (s *SelectList) Refresh(t *Terminal) {
	t.Cursor.MoveUp(s.height)
	s.Display(t)
}

func (s *SelectList) OnKey(t *Terminal, k Key) bool {
	switch k.Type() {
	case KeyArrowUp:
		s.activeItem--
	case KeyArrowDown:
		s.activeItem++
	case KeyEnter:
		return false
	}
	if s.activeItem < 0 {
		s.activeItem = len(s.Items) - 1
	} else if s.activeItem >= len(s.Items) {
		s.activeItem = 0
	}
	s.Refresh(t)
	return true
}