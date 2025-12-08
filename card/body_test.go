package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/stretchr/testify/assert"
)

func TestBodyReadByLine(t *testing.T) {
	counter := 0
	body := card.Body(`
line1
line2
line3
`)

	t.Run("counter", func(t *testing.T) {
		body.ReadByLine(func(line string) (exit bool) {
			counter++
			return
		})

		assert.Equal(t, 3, counter)
	})

	t.Run("interrupt", func(t *testing.T) {
		counter = 0

		body.ReadByLine(func(line string) (exit bool) {
			counter++
			return true
		})

		assert.Equal(t, 1, counter)
	})
}
