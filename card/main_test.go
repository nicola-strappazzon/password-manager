package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/stretchr/testify/assert"
)

func TestLockup(t *testing.T) {
	c := card.Card{Body: `
testPassword
password: testPasswordTwo
username: aaa
email: user@com.local
host: 127.0.0.1
empty:
`}

	assert.Equal(t, "testPasswordTwo", c.Lockup("password"))
	// assert.Equal(t, c.Lockup("username"), "aaa")
	// assert.Equal(t, c.Lockup("email"), "user@com.local")
	// assert.Equal(t, c.Lockup("host"), "127.0.0.1")
	// assert.Equal(t, c.Lockup("empty"), "")
	// assert.Equal(t, c.Lockup("no-exist"), "")
}
