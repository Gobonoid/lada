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
