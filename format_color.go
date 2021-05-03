package lada

type FormatColor uint8

var formatColorNames = map[FormatColor]string {
	39: "default",
	30: "black",
	31: "red",
	32: "green",
	33: "yellow",
	34: "blue",
	35: "magenta",
	36: "cyan",
	37: "lightgray",
	90: "darkgray",
	91: "lightred",
	92: "lightgreen",
	93: "lightyellow",
	94: "lightblue",
	95: "lightmagenta",
	96: "lightcyan",
	97: "white",
}

var FormatColors = struct {
	DEFAULT FormatColor
	BLACK FormatColor
	RED FormatColor
	GREEN FormatColor
	YELLOW FormatColor
	BLUE FormatColor
	MAGENTA FormatColor
	CYAN FormatColor
	LIGHTGRAY FormatColor
	DARKGRAY FormatColor
	LIGHTRED FormatColor
	LIGHTGREEN FormatColor
	LIGHTYELLOW FormatColor
	LIGHTBLUE FormatColor
	LIGHTMAGENTA FormatColor
	LIGHTCYAN FormatColor
	WHITE FormatColor

}{
	DEFAULT: 39,
	BLACK: 30,
	RED: 31,
	GREEN: 32,
	YELLOW: 33,
	BLUE: 34,
	MAGENTA: 35,
	CYAN: 36,
	LIGHTGRAY: 37,
	DARKGRAY: 90,
	LIGHTRED: 91,
	LIGHTGREEN: 92,
	LIGHTYELLOW: 93,
	LIGHTBLUE: 94,
	LIGHTMAGENTA: 95,
	LIGHTCYAN: 96,
	WHITE: 97,
}

func (f FormatColor) Name() string {
	return formatColorNames[f]
}