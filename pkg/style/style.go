package lada

import (
	"strconv"
	"strings"
)

type Style interface {
	Type() FormatType
	Value() int
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

func (f format) Type() FormatType {
	return FormatTypeFormat
}

func (f format) Value() int {
	return int(f)
}

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

func (b background) Type() FormatType {
	return FormatTypeBackground
}

func (b background) Value() int {
	return int(b) + 10
}

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

func (f foreground) Type() FormatType {
	return FormatTypeForeground
}

func (f foreground) Value() int {
	return int(f)
}

type sgr struct {
	foreground int
	background int
	formats []int
}

func newSgr(style ...Style) sgr {
	result := sgr{formats: make([]int, 0)}
	for _, item := range style {
		switch item.Type() {
		case FormatTypeForeground:
			result.foreground = item.Value()
		case FormatTypeBackground:
			result.background = item.Value()
		case FormatTypeFormat:
			result.formats = append(result.formats, item.Value())
		}
	}
	return result
}

func (s sgr) Value() string {
	result := "["
	var formats []string
	if s.background != 0 {
		formats = append(formats, strconv.Itoa(s.background))
	}

	if s.foreground != 0 {
		formats = append(formats, strconv.Itoa(s.foreground))
	}

	for _, f := range s.formats {
		formats = append(formats, strconv.Itoa(f))
	}

	return result + strings.Join(formats, ";") + "m"
}

type StyleSheet map[string][]Style