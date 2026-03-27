package integration_test

import (
	"testing"
)

func testUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// stdout, stderr, err := run(update.NewCommand(), []string{"work/github", "-f", "username", "-v", "new-user", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.Empty(t, stderr)
		// assert.NoError(t, err)

		// stdout, stderr, err = run(show.NewCommand(), []string{"work/github", "-f", "username", "-p", testPassphrase})
		// assert.Equal(t, "new-user\n", stdout)
		// assert.Empty(t, stderr)
		// assert.NoError(t, err)
	})

	t.Run("password", func(t *testing.T) {
		// stdout, stderr, err := run(update.NewCommand(), []string{"work/github", "-f", "password", "-v", "NewP@ssw0rd!", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		// assert.NoError(t, err)

		// stdout, _, err = run(show.NewCommand(), []string{"work/github", "-p", testPassphrase})
		// assert.Equal(t, "NewP@ssw0rd!\n", stdout)
		// assert.NoError(t, err)
	})

	t.Run("field-without-value", func(t *testing.T) {
		// stdout, _, err := run(update.NewCommand(), []string{"work/github", "-f", "host", "-v", "github.com", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.EqualError(t, err, "Field 'host' does not have a value. Use 'pm add' to set it.")
	})

	t.Run("source-not-found", func(t *testing.T) {
		// stdout, _, err := run(update.NewCommand(), []string{"nonexistent", "-f", "username", "-v", "user"})
		// assert.Empty(t, stdout)
		// assert.EqualError(t, err, "No such file or directory.")
	})
}
