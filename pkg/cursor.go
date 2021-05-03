package lada

import (
	"fmt"
	"io"
	"strings"
)

const CursorUp = "\033[%dA"
const CursorDown = "\033[%dB"
const CursorForward = "\033[%dC"
const CursorBackward = "\033[%dD"
const CursorNextLine = "\033[%dE"
const CursorPreviousLine = "\033[%dF"
const CursorHorizontalAbsolute = "\033[%dG"
const CursorPosition = "\033[%d;%dH"
const CursorEraseInDisplay = "\033[%dJ"
const CursorEraseInLine = "\033[%dK" // 0 to end of line, 1 to beginning of line, 2 entire line
const CursorHide = "\033[25h"
const CursorShow = "\033[25l"

const CursorSgrReset = "\033[0m"

type Position struct {
	Line int
	Column int
}

type Cursor struct {
	 writer io.Writer
	 line int
	 column int
}

func NewCursor(writer io.Writer) (*Cursor, error) {
	// Reset cursor position to new line and set coordinates to 0,0
	if _, err := fmt.Fprintf(writer, CursorNextLine, 1); err != nil {
		return nil, CursorOperationError.causedBy(err)
	}
	return &Cursor{
		writer: writer,
		line: 1,
		column: 1,
	}, nil
}

func (c *Cursor) updatePosition(text string) {
	split := strings.Split(text, "\n")

	c.line += len(split)
	c.column += len(split[len(split) - 1])
}

func (c *Cursor) MoveToNextLine() error {
	if _, err := fmt.Fprintf(c.writer, CursorNextLine, 1); err != nil {
		return CursorOperationError.causedBy(err)
	}
	c.line += 1
	return nil
}

func (c *Cursor) MoveToPreviousLine() error {
	if c.line <= 1 {
		return CursorOperationError.causedBy(CursorOutOfReachError)
	}

	if _, err := fmt.Fprintf(c.writer, CursorPreviousLine, 1); err != nil {
		return CursorOperationError.causedBy(err)
	}
	c.line -= 1
	return nil
}

func (c *Cursor) MoveForward(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorForward, n); err != nil {
		return CursorOperationError.causedBy(err)
	}
	c.column += n
	return nil
}

func (c *Cursor) MoveBackward(n int) error {
	if _, err := fmt.Fprintf(c.writer, CursorBackward, n); err != nil {
		return CursorOperationError.causedBy(err)
	}
	c.column -= n
	if c.column < 1 {
		c.column = 1
	}

	return nil
}

func (c *Cursor) MoveUp(n int) error {
	if c.line - n < 1 {
		return CursorOperationError.causedBy(CursorOutOfReachError)
	}

	return nil
}