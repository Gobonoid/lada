package lada

import (
	"fmt"
	"io"
)

const escape = "\033"
const CursorUp = escape + "[%dA"
const CursorDown = escape + "[%dB"
const CursorForward = escape + "[%dC"
const CursorBackward = escape + "[%dD"
const CursorNextLine = escape + "[%dE"
const CursorPreviousLine = escape + "[%dF"
const CursorEraseInDisplay = escape + "[%dJ" // 0 to end of screen, 1 to beginning of screen, 2 entire screen
const CursorEraseInLine = escape + "[%dK"    // 0 to end of line, 1 to beginning of line, 2 entire line
const CursorHide = escape + "[25h"
const CursorShow = escape + "[25l"
const CursorSgrReset = escape + "[0m"

type Cursor struct {
	writer io.Writer
	style  *cursorStyle
}

func NewCursor(writer io.Writer) (*Cursor, error) {
	return &Cursor{
		writer: writer,
	}, nil
}

func (c *Cursor) MoveToNextLine() error {
	if _, err := fmt.Fprintf(c.writer, CursorNextLine, 1); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) MoveToPreviousLine() error {
	if _, err := fmt.Fprintf(c.writer, CursorPreviousLine, 1); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) MoveForward(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorForward, n); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) MoveBackward(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorBackward, n); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) MoveUp(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorUp, n); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) MoveDown(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorDown, n); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) Hide() error {
	if _, err := fmt.Fprint(c.writer, CursorHide); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) Show() error {
	if _, err := fmt.Fprint(c.writer, CursorShow); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) SetStyle(style *cursorStyle) {
	c.style = style
	c.applyStyle()
}

func (c *Cursor) applyStyle() error {
	sgr := c.style.Sgr()
	if _, err := fmt.Fprint(c.writer, sgr); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) Print(text string) error {
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) EraseDisplay() error {
	if _, err := fmt.Fprintf(c.writer, CursorEraseInDisplay, 2); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) EraseLine() error {
	if _, err := fmt.Fprintf(c.writer, CursorEraseInLine, 2); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) Close() error {
	if _, err := fmt.Fprint(c.writer, "\n"); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) ResetStyle() error {
	if _, err := fmt.Fprint(c.writer, CursorSgrReset); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}
