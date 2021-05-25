package style

type foreground uint8

var Foreground = struct {
	Default      foreground
	Black        foreground
	Red          foreground
	Green        foreground
	Yellow       foreground
	Blue         foreground
	Magenta      foreground
	Cyan         foreground
	LightGray    foreground
	DarkGray     foreground
	LightRed     foreground
	LightGreen   foreground
	LightYellow  foreground
	LightBlue    foreground
	LightMagenta foreground
	LightCyan    foreground
	White        foreground
}{
	Default:      foreground(Color.Default),
	Black:        foreground(Color.Black),
	Red:          foreground(Color.Red),
	Green:        foreground(Color.Green),
	Yellow:       foreground(Color.Yellow),
	Blue:         foreground(Color.Blue),
	Magenta:      foreground(Color.Magenta),
	Cyan:         foreground(Color.Cyan),
	LightGray:    foreground(Color.LightGray),
	DarkGray:     foreground(Color.DarkGray),
	LightRed:     foreground(Color.LightRed),
	LightGreen:   foreground(Color.LigthGreen),
	LightYellow:  foreground(Color.LightYellow),
	LightBlue:    foreground(Color.LightBlue),
	LightMagenta: foreground(Color.LightMagenta),
	LightCyan:    foreground(Color.LightCyan),
	White:        foreground(Color.White),
}

func (f foreground) Type() StyleType {
	return ForegroundStyle
}

func (f foreground) Value() int {
	return int(f)
}
