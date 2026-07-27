package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/cli/move"
	"github.com/nicola-strappazzon/password-manager/cli/show"
	"github.com/stretchr/testify/assert"
)

func testMove(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stdout, stderr, err := run(move.NewCommand(), []string{"github", "work/github", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("source-not-found", func(t *testing.T) {
		stdout, _, err := run(move.NewCommand(), []string{"nonexistent", "other"})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "No such file or directory.")
	})

	t.Run("destination-already-exists", func(t *testing.T) {
		stdout, _, err := run(move.NewCommand(), []string{"work/github", "work/github"})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "Destination already exists.")
	})

	t.Run("missing-arguments", func(t *testing.T) {
		stdout, _, err := run(move.NewCommand(), []string{"work/github"})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "Source and destination paths are required.")
	})

	t.Run("file-and-directory-with-same-name", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"move-collision", "-f", "password", "-v", "move-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(add.NewCommand(), []string{"move-collision/prd", "-f", "password", "-v", "prd-secret", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
		assert.NoError(t, err)

		stdout, stderr, err = run(move.NewCommand(), []string{"move-collision", "move-collision-target", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)

		stdout, stderr, err = run(show.NewCommand(), []string{"move-collision-target", "-p", testPassphrase})
		assert.Equal(t, "move-secret\n", stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})
}
