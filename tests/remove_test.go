package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/cli/remove"
	"github.com/stretchr/testify/assert"
)

func testRemove(t *testing.T) {
	t.Run("file-and-directory-with-same-name", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"remove-collision", "-f", "password", "-v", "remove-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(add.NewCommand(), []string{"remove-collision/prd", "-f", "password", "-v", "prd-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(remove.NewCommand(), []string{"remove-collision"})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})
}
