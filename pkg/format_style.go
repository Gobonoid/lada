package lada

type FormatStyle uint8

var formatStyleNames = map[FormatStyle]string {
	1: "bold",
	2: "dim",
	3: "italic",
	4: "underline",
	5: "blink",
	7: "invert",
	8: "hidden",
	9: "strike",
}

var FormatStyles = struct {
	BOLD      FormatStyle
	DIM       FormatStyle
	ITALIC    FormatStyle
	UNDERLINE FormatStyle
	BLINK     FormatStyle
	INVERT    FormatStyle
	HIDDEN    FormatStyle
	STRIKE    FormatStyle
}{
	BOLD: 1,
	DIM: 2,
	ITALIC: 3,
	UNDERLINE: 4,
	BLINK: 5,
	INVERT: 7,
	HIDDEN: 8,
	STRIKE: 9,
}

func (f FormatStyle) Name() string {
	return formatStyleNames[f]
}


func PrintText(text string, style FormatStyle, color FormatColor) {

}