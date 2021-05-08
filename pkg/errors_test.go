package lada

import (
	"errors"
	"testing"
)

func TestLadaError(t *testing.T) {
	t.Run("Test can create new CliIoReadError", func(t *testing.T) {
		testCause := errors.New("test")
		err := IoReaderError.CausedBy(testCause)
		if !errors.Is(err, IoReaderError) {
			t.Error("CliIoReadError.Is(CliIoReadError) should be true")
		}
	})
}
