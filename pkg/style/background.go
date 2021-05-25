package style

type background uint8

var Background = struct {
	Default      background
	Black        background
	Red          background
	Green        background
	Yellow       background
	Blue         background
	Magenta      background
	Cyan         background
	LightGray    background
	DarkGray     background
	LightRed     background
	LightGreen   background
	LightYellow  background
	LightBlue    background
	LightMagenta background
	LightCyan    background
	White        background
}{
	Default:      background(Color.Default),
	Black:        background(Color.Black),
	Red:          background(Color.Red),
	Green:        background(Color.Green),
	Yellow:       background(Color.Yellow),
	Blue:         background(Color.Blue),
	Magenta:      background(Color.Magenta),
	Cyan:         background(Color.Cyan),
	LightGray:    background(Color.LightGray),
	DarkGray:     background(Color.DarkGray),
	LightRed:     background(Color.LightRed),
	LightGreen:   background(Color.LigthGreen),
	LightYellow:  background(Color.LightYellow),
	LightBlue:    background(Color.LightBlue),
	LightMagenta: background(Color.LightMagenta),
	LightCyan:    background(Color.LightCyan),
	White:        background(Color.White),
}

func (b background) Type() StyleType {
	return BackgroundStyle
}

func (b background) Value() int {
	return int(b) + 10
}
