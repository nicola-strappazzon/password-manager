package integration_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/file"
	"github.com/stretchr/testify/assert"
)

func testFile(t *testing.T) {
	sourceDir := t.TempDir()
	sourcePath := filepath.Join(sourceDir, "front.png")
	original := []byte("demo-binary-content")
	assert.NoError(t, os.WriteFile(sourcePath, original, 0600))

	t.Run("add", func(t *testing.T) {
		stdout, stderr, err := run(file.NewCommand(), []string{"github", "-i", sourcePath, "-p", testPassphrase})
		assert.NoError(t, err)
		assert.Empty(t, stderr)
		assert.Equal(t, "Added file front.png to the GPG-encrypted container github.\n", stdout)
	})

	t.Run("list", func(t *testing.T) {
		stdout, stderr, err := run(file.NewCommand(), []string{"github", "-p", testPassphrase})
		assert.NoError(t, err)
		assert.Empty(t, stderr)
		assert.Contains(t, stdout, "Files inside github:\n")
		assert.Contains(t, stdout, " - front.png (")
	})

	t.Run("export", func(t *testing.T) {
		outputPath := filepath.Join(t.TempDir(), "restored-front.png")
		stdout, stderr, err := run(file.NewCommand(), []string{"github", "-e", "front.png", "-o", outputPath, "-p", testPassphrase})
		assert.NoError(t, err)
		assert.Empty(t, stderr)
		assert.Equal(t, "Saved file front.png to "+outputPath+".\n", stdout)

		restored, readErr := os.ReadFile(outputPath)
		assert.NoError(t, readErr)
		assert.Equal(t, original, restored)
	})

	t.Run("delete", func(t *testing.T) {
		stdout, stderr, err := run(file.NewCommand(), []string{"github", "-d", "front.png", "-p", testPassphrase})
		assert.NoError(t, err)
		assert.Empty(t, stderr)
		assert.Equal(t, "Deleted file front.png from the GPG-encrypted container github.\n", stdout)
	})

	t.Run("list-after-delete", func(t *testing.T) {
		stdout, stderr, err := run(file.NewCommand(), []string{"github", "-p", testPassphrase})
		assert.NoError(t, err)
		assert.Empty(t, stderr)
		assert.Contains(t, stdout, "Files inside github:\n")
		assert.NotContains(t, stdout, "front.png")
	})
}
