package lada

import "unicode/utf8"

type Key []byte

type KeyType int

const (
	KeyArrowLeft KeyType = iota
	KeyTab
	KeySpace
	KeyEnter
	KeyEscape
	KeyBackspace
	KeyArrowRight
	KeyArrowUp
	KeyHome
	KeyEnd
	KeyArrowDown
	KeyRune
)

func (k Key) Type() KeyType {
	switch k[0] {
	case 9:
		return KeyTab
	case 32:
		return KeySpace
	case 127:
		return KeyBackspace
	case 10:
		return KeyEnter
	case 27:
		if k[1] == '[' {
			switch k[2] {
			case 'A':
				return KeyArrowUp
			case 'B':
				return KeyArrowDown
			case 'C':
				return KeyArrowRight
			case 'D':
				return KeyArrowLeft
			case 'H':
				return KeyHome
			case 'F':
				return KeyEnd
			}
		}
		// alt + arrows <>, lets ignore alt and return arrows
		if k[1] == 'b' && k[2] == 0 && k[3] == 0 {
			return KeyArrowLeft
		}
		if k[1] == 'f' && k[2] == 0 && k[3] == 0  {
			return KeyArrowRight
		}

		// either escape or some escaped sequence which we dont care about
		return KeyEscape
	}

	return KeyRune
}

func (k Key) AsRune() rune {
	r, _ := utf8.DecodeRune(k)
	return r
}

func (k Key) AsString() string {
	return string(k.AsRune())
}

func (k Key) Equals(r rune) bool {
	if k.Type() == KeyRune && k.AsRune() == r {
		return true
	}

	return false
}