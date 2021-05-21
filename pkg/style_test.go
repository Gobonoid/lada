package lada

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStyle(t *testing.T) {
	t.Run("can set background", func(t *testing.T) {
		style := newSgr(Background.RED, Foreground.GREEN)

		assert.Equal(t, "[41;32m", style.Value())

		fmt.Println(newSgr(Background.BLUE, Foreground.WHITE, Format.BOLD, Format.BLINK).Value())
	})
}
