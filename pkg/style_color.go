package lada

type color uint8

var colors = map[color]string{
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

var Color = struct {
	Default      color
	Black        color
	Red          color
	Green        color
	Yellow       color
	Blue         color
	Magenta      color
	Cyan         color
	LightGray    color
	DarkGray     color
	LightRed     color
	LigthGreen   color
	LightYellow  color
	LightBlue    color
	LightMagenta color
	LightCyan    color
	White        color
}{
	Default:      39,
	Black:        30,
	Red:          31,
	Green:        32,
	Yellow:       33,
	Blue:         34,
	Magenta:      35,
	Cyan:         36,
	LightGray:    37,
	DarkGray:     90,
	LightRed:     91,
	LigthGreen:   92,
	LightYellow:  93,
	LightBlue:    94,
	LightMagenta: 95,
	LightCyan:    96,
	White:        97,
}

func (c color) Name() string {
	return colors[c]
}