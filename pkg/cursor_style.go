package lada

import (
	"strconv"
	"strings"
)

type cursorStyle struct {
	foreground Color
	background Color
	styles     []Style
}

func NewCursorStyle(foreground Color, background Color, style ...Style) *cursorStyle {
	return &cursorStyle{
		foreground: foreground,
		background: background + 10,
		styles:     style,
	}
}

func (s *cursorStyle) Sgr() string {
	sgr := escape + "["
	sequences := []string{strconv.Itoa(int(s.foreground)), strconv.Itoa(int(s.background))}
	for _, style := range s.styles {
		sequences = append(sequences, strconv.Itoa(int(style)))
	}

	return sgr + strings.Join(sequences, ";") + "m"
}
