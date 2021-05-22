package lada

const slDisabledItem = "◌"
const slActiveItem = "○"
const slSelectedItem = "⦿"

type SelectUI struct {
	Label      string
	Items      []string
	activeItem int
	height     int
	drawn      bool
}

func NewSelectUI(label string, items []string) *SelectUI {
	l := &SelectUI{
		Label:      label,
		Items:      items,
		activeItem: 0,
		height:     len(items) + 1,
		drawn:      false,
	}

	return l
}

func (s *SelectUI) Display(t *Terminal) error {
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

	return nil
}

func (s *SelectUI) Value() int {
	return s.activeItem
}

func (s *SelectUI) Remove(t *Terminal) error {
	return nil
}

func (s *SelectUI) refresh(t *Terminal) {
	t.Cursor.MoveUp(s.height)
	s.Display(t)
}

func (s *SelectUI) OnKey(t *Terminal, k Key) bool {
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
	s.refresh(t)
	return true
}