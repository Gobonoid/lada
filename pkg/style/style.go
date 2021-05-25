package style

import (
	"strconv"
	"strings"
)

type Style interface {
	Type() StyleType
	Value() int
}

func NewSgr(style ...Style) string {
	foreground := ""
	background := ""
	formats := make([]string, 0)

	for _, item := range style {
		switch item.Type() {
		case ForegroundStyle:
			foreground = strconv.Itoa(item.Value())
		case BackgroundStyle:
			background = strconv.Itoa(item.Value())
		case FormatStyle:
			formats = append(formats, strconv.Itoa(item.Value()))
		}
	}

	if background != "" {
		formats = append([]string{background}, formats...)
	}

	if foreground != "" {
		formats = append([]string{foreground}, formats...)
	}

	return "[" + strings.Join(formats, ";") + "m"
}