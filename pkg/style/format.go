package style

type StyleType uint8

const (
	FormatStyle StyleType = iota
	BackgroundStyle
	ForegroundStyle
)

type format uint8

var formatters = map[format]string{
	1: "bold",
	2: "dim",
	3: "italic",
	4: "underline",
	5: "blink",
	7: "invert",
	8: "hidden",
	9: "strike",
}

var Format = struct {
	Bold      format
	Dim       format
	Italic    format
	Underline format
	Blink     format
	Invert    format
	Hidden    format
	Strike    format
}{
	Bold:      1,
	Dim:       2,
	Italic:    3,
	Underline: 4,
	Blink:     5,
	Invert:    7,
	Hidden:    8,
	Strike:    9,
}

func (f format) Name() string {
	return formatters[f]
}

func (f format) Type() StyleType {
	return FormatStyle
}

func (f format) Value() int {
	return int(f)
}
