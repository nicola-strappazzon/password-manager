package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/update"
	"github.com/stretchr/testify/assert"
)

func testUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stdout, stderr, err := run(update.NewCommand(), []string{"work/github", "-f", "username", "-v", "new-user", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("field-without-value", func(t *testing.T) {
		stdout, _, err := run(update.NewCommand(), []string{"work/github", "-f", "host", "-v", "github.com", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "Field 'host' does not have a value. Use 'pm add' to set it.")
	})

	t.Run("source-not-found", func(t *testing.T) {
		stdout, _, err := run(update.NewCommand(), []string{"nonexistent", "-f", "username", "-v", "user"})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "No such file or directory.")
	})
}
