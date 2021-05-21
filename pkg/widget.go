package lada

type SelectList struct {
	Label string
	Items []string
	activeItem int
	size int
}

func NewSelectList(label string, items []string) *SelectList {
	l := &SelectList{
		Label: label,
		Items: items,
		activeItem: 0,
	}

	return l
}

/*
○
⦿
◌
*/