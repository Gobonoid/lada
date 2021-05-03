package lada

import (
	"errors"
	"fmt"
	"github.com/kodemore/lada/pkg/mock"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestNewCursor(t *testing.T) {
	t.Run("creates new cursor with valid writer", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, err := NewCursor(writer)
			assert.Nil(t, err)
			assert.NotEmpty(t, cursor)

			// writer should contain an escape character, when cursor is created
			writer.Seek(0, 0)
			contents, _ := io.ReadAll(writer)

			assert.Equal(t, fmt.Sprintf(CursorNextLine, 1), string(contents))
		})
	})

	t.Run("fails to create cursor with invalid writer", func(t *testing.T) {
		mock.UseInvalidIoWriterMock(t, func(writer *os.File) {
			cursor, err := NewCursor(writer)
			assert.Empty(t, cursor)
			assert.True(t, errors.Is(err, CursorOperationError))
		})
	})
}


func TestCursor_MoveForward(t *testing.T) {
	t.Run("can move cursor forward", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveForward(1)

			writer.Seek(0, 0)
			contents, _ := io.ReadAll(writer)
			expected := fmt.Sprintf(CursorNextLine, 1) + fmt.Sprintf(CursorForward, 1)
			assert.Equal(t, expected, string(contents))
			assert.Equal(t, 2, cursor.Column())
		})
	})
}

func TestCursor_MoveBackward(t *testing.T) {
	t.Run("can move cursor forward", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveBackward(1)

			writer.Seek(0, 0)
			contents, _ := io.ReadAll(writer)
			expected := fmt.Sprintf(CursorNextLine, 1) + fmt.Sprintf(CursorBackward, 1)
			assert.Equal(t, expected, string(contents))
		})
	})
}