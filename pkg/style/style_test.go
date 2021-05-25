package style

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStyle(t *testing.T) {
	t.Run("can set background", func(t *testing.T) {
		style := NewSgr(Background.Red, Foreground.Green)

		assert.Equal(t, "[32;41m", style)

		fmt.Println(NewSgr(Background.Blue, Foreground.White, Format.Bold, Format.Blink))
	})
}
