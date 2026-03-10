package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/ls"
	"github.com/stretchr/testify/assert"
)

func testLs(t *testing.T) {
	t.Run("ls", func(t *testing.T) {
		stdout, stderr, err := run(ls.NewCommand(), []string{})
		assert.Empty(t, stderr)
		assert.NoError(t, err)
		assert.Equal(t, "\x1b[1;37mPassword Store\x1b[0m\n└── github\n", stdout)
	})
}
