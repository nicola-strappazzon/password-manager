package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/cli/show"
	"github.com/stretchr/testify/assert"
)

func testShow(t *testing.T) {
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
