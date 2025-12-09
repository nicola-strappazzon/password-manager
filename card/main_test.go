package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/card"

	"github.com/stretchr/testify/assert"
)

func TestLockup(t *testing.T) {
	yml := `---
password: "#O123456@bc"
`

	assert.Equal(t, card.New(yml).Password, "#O123456@bc")
}
