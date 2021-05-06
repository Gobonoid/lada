package lada

type Style uint8

var styleNames = map[Style]string{
	1: "bold",
	2: "dim",
	3: "italic",
	4: "underline",
	5: "blink",
	7: "invert",
	8: "hidden",
	9: "strike",
}

var Styles = struct {
	BOLD      Style
	DIM       Style
	ITALIC    Style
	UNDERLINE Style
	BLINK     Style
	INVERT    Style
	HIDDEN    Style
	STRIKE    Style
}{
	BOLD:      1,
	DIM:       2,
	ITALIC:    3,
	UNDERLINE: 4,
	BLINK:     5,
	INVERT:    7,
	HIDDEN:    8,
	STRIKE:    9,
}

func (f Style) Name() string {
	return styleNames[f]
}
