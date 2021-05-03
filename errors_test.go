package lada

import (
	"errors"
	"testing"
)

func TestLadaError(t *testing.T) {
	t.Run("Test can create new CliIoReadError", func(t *testing.T) {
		testCause := errors.New("test")
		err := CliIoReadError.causedBy(testCause)
		if !errors.Is(err, CliIoReadError) {
			t.Error("CliIoReadError.Is(CliIoReadError) should be true")
		}
	})
}
