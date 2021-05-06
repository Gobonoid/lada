package lada

import (
	"fmt"
	"github.com/kodemore/lada/pkg/mock"
	"github.com/stretchr/testify/assert"
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
			contents := mock.ReadFileAsString(writer)
			assert.Empty(t, contents)
		})
	})
}

func TestCursor_MoveForward(t *testing.T) {
	t.Run("can move cursor forward", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveForward(1)

			contents := mock.ReadFileAsString(writer)
			expected := fmt.Sprintf(CursorForward, 1)
			assert.Equal(t, expected, contents)
		})
	})
}

func TestCursor_MoveBackward(t *testing.T) {
	t.Run("can move cursor forward", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveBackward(5)

			contents := mock.ReadFileAsString(writer)
			expected := fmt.Sprintf(CursorBackward, 5)
			assert.Equal(t, expected, contents)
		})
	})

	t.Run("cannot exceed minimum column value of 1", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveBackward(5)

			contents := mock.ReadFileAsString(writer)
			expected := fmt.Sprintf(CursorBackward, 5)
			assert.Equal(t, expected, contents)
		})
	})
}

func TestCursor_MoveUp(t *testing.T) {
	t.Run("can move cursor up", func(t *testing.T) {
		mock.UseIoWriterMock(t, func(writer *os.File) {
			cursor, _ := NewCursor(writer)
			cursor.MoveUp(2)

			contents := mock.ReadFileAsString(writer)
			expected := fmt.Sprintf(CursorUp, 2)
			assert.Equal(t, expected, contents)
		})
	})
}
