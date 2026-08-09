package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/cli/show"
	"github.com/stretchr/testify/assert"
)

func testShow(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"show/github", "-f", "password", "-v", "show-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(show.NewCommand(), []string{"show/github", "-p", testPassphrase})
		assert.Equal(t, "show-secret\n", stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("field", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"show/github", "-f", "username", "-v", "show-user", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)

		stdout, stderr, err = run(show.NewCommand(), []string{"show/github", "-f", "username", "-p", testPassphrase})
		assert.Equal(t, "show-user\n", stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("all", func(t *testing.T) {
		stdout, stderr, err := run(show.NewCommand(), []string{"show/github", "-a", "-p", testPassphrase})
		assert.Contains(t, stdout, "password: show-secret\n")
		assert.Contains(t, stdout, "username: show-user\n")
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("invalid-field", func(t *testing.T) {
		stdout, stderr, err := run(show.NewCommand(), []string{"show/github", "-f", "invalid", "-p", testPassphrase})
		assert.Contains(t, stdout, "Usage:")
		assert.Contains(t, stderr, "Error: Invalid field: invalid")
		assert.EqualError(t, err, "Invalid field: invalid")
	})

	t.Run("file-and-directory-with-same-name", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"mongodb", "-f", "password", "-v", "mongodb-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(add.NewCommand(), []string{"mongodb/prd", "-f", "password", "-v", "prd-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(show.NewCommand(), []string{"mongodb", "-p", testPassphrase})
		assert.Equal(t, "mongodb-secret\n", stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})
}
