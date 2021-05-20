package lada

type Style interface {
	Type() FormatType
	Value() uint8
}

type FormatType uint8

const (
	FormatTypeFormat FormatType = iota
	FormatTypeBackground
	FormatTypeForeground
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
	BOLD      format
	DIM       format
	ITALIC    format
	UNDERLINE format
	BLINK     format
	INVERT    format
	HIDDEN    format
	STRIKE    format
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

func (f format) Name() string {
	return formatters[f]
}

func (f format) Type() FormatType {
	return FormatTypeFormat
}

func (f format) Value() uint8 {
	return uint8(f)
}

type background uint8

var Background = struct {
	DEFAULT      background
	BLACK        background
	RED          background
	GREEN        background
	YELLOW       background
	BLUE         background
	MAGENTA      background
	CYAN         background
	LIGHTGRAY    background
	DARKGRAY     background
	LIGHTRED     background
	LIGHTGREEN   background
	LIGHTYELLOW  background
	LIGHTBLUE    background
	LIGHTMAGENTA background
	LIGHTCYAN    background
	WHITE        background
}{
	DEFAULT:      background(Color.DEFAULT),
	BLACK:        background(Color.BLACK),
	RED:          background(Color.RED),
	GREEN:        background(Color.GREEN),
	YELLOW:       background(Color.YELLOW),
	BLUE:         background(Color.BLUE),
	MAGENTA:      background(Color.MAGENTA),
	CYAN:         background(Color.CYAN),
	LIGHTGRAY:    background(Color.LIGHTGRAY),
	DARKGRAY:     background(Color.DARKGRAY),
	LIGHTRED:     background(Color.LIGHTRED),
	LIGHTGREEN:   background(Color.LIGHTGREEN),
	LIGHTYELLOW:  background(Color.LIGHTYELLOW),
	LIGHTBLUE:    background(Color.LIGHTBLUE),
	LIGHTMAGENTA: background(Color.LIGHTMAGENTA),
	LIGHTCYAN:    background(Color.LIGHTCYAN),
	WHITE:        background(Color.WHITE),
}

func (b background) Type() FormatType {
	return FormatTypeBackground
}

func (b background) Value() uint8 {
	return uint8(b) + 10
}

type foreground uint8

var Foreground = struct {
	DEFAULT      foreground
	BLACK        foreground
	RED          foreground
	GREEN        foreground
	YELLOW       foreground
	BLUE         foreground
	MAGENTA      foreground
	CYAN         foreground
	LIGHTGRAY    foreground
	DARKGRAY     foreground
	LIGHTRED     foreground
	LIGHTGREEN   foreground
	LIGHTYELLOW  foreground
	LIGHTBLUE    foreground
	LIGHTMAGENTA foreground
	LIGHTCYAN    foreground
	WHITE        foreground
}{
	DEFAULT:      foreground(Color.DEFAULT),
	BLACK:        foreground(Color.BLACK),
	RED:          foreground(Color.RED),
	GREEN:        foreground(Color.GREEN),
	YELLOW:       foreground(Color.YELLOW),
	BLUE:         foreground(Color.BLUE),
	MAGENTA:      foreground(Color.MAGENTA),
	CYAN:         foreground(Color.CYAN),
	LIGHTGRAY:    foreground(Color.LIGHTGRAY),
	DARKGRAY:     foreground(Color.DARKGRAY),
	LIGHTRED:     foreground(Color.LIGHTRED),
	LIGHTGREEN:   foreground(Color.LIGHTGREEN),
	LIGHTYELLOW:  foreground(Color.LIGHTYELLOW),
	LIGHTBLUE:    foreground(Color.LIGHTBLUE),
	LIGHTMAGENTA: foreground(Color.LIGHTMAGENTA),
	LIGHTCYAN:    foreground(Color.LIGHTCYAN),
	WHITE:        foreground(Color.WHITE),
}

func (f foreground) Type() FormatType {
	return FormatTypeForeground
}

func (f foreground) Value() uint8 {
	return uint8(f)
}