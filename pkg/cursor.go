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
const CursorHide = escape + "[?25l"
const CursorShow = escape + "[?25h"
const CursorSgrReset = escape + "[0m"
const CursorResetColor = escape + "[32m"

type Cursor struct {
	writer io.Writer
	hidden bool
	style  []Style
}

func NewCursor(writer io.Writer) (*Cursor, error) {
	return &Cursor{
		writer: writer,
		hidden: false,
	}, nil
}

func (c *Cursor) MoveToNextLine() error {
	return c.execute(fmt.Sprintf(CursorNextLine, 1))
}

func (c *Cursor) MoveToPreviousLine() error {
	return c.execute(fmt.Sprintf(CursorPreviousLine, 1))
}

func (c *Cursor) MoveForward(n int) error {
	return c.execute(fmt.Sprintf(CursorForward, n))
}

func (c *Cursor) MoveBackward(n int) error {
	return c.execute(fmt.Sprintf(CursorBackward, n))
}

func (c *Cursor) MoveUp(n int) error {
	return c.execute(fmt.Sprintf(CursorUp, n))
}

func (c *Cursor) MoveDown(n int) error {
	return c.execute(fmt.Sprintf(CursorDown, n))
}

func (c *Cursor) EraseDisplay() error {
	return c.execute(fmt.Sprintf(CursorEraseInDisplay, 2))
}

func (c *Cursor) EraseLine() error {
	return c.execute(fmt.Sprintf(CursorEraseInLine, 2))
}

func (c *Cursor) SetStyle(style ...Style) error {
	c.style = style
	return c.execute(escape + newSgr(c.style...).Value())
}

func (c *Cursor) ResetStyle() error {
	return c.execute(CursorSgrReset)
}

func (c *Cursor) IsHidden() bool {
	return c.hidden
}

func (c *Cursor) Hide() error {
	if err := c.execute(CursorHide); err != nil {
		return err
	}
	c.hidden = false
	return nil
}

func (c *Cursor) Show() error {
	if err := c.execute(CursorShow); err != nil {
		return err
	}
	c.hidden = false
	return nil
}

func (c *Cursor) execute(s string) error {
	if _, err := fmt.Fprint(c.writer, s); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}

func (c *Cursor) close() error {
	c.Show()
	if _, err := fmt.Fprint(c.writer, "\n"); err != nil {
		return CursorOperationError.CausedBy(err)
	}
	return nil
}