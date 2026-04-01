package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestPersistentPreRunE(t *testing.T) {
	homeDir := t.TempDir()
	oldUserHomeDir := config.UserHomeDir
	oldDataDir := config.DataDir
	config.UserHomeDir = func() (string, error) { return homeDir, nil }
	config.DataDir = ""
	t.Cleanup(func() {
		config.UserHomeDir = oldUserHomeDir
		config.DataDir = oldDataDir
	})

	t.Run("allows version without setup", func(t *testing.T) {
		cmd := &cobra.Command{Use: "version"}
		err := PersistentPreRunE(cmd, []string{})
		assert.NoError(t, err)
	})

	t.Run("requires setup for other commands", func(t *testing.T) {
		cmd := &cobra.Command{Use: "ls"}
		err := PersistentPreRunE(cmd, []string{})
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "Run 'pm setup'"))
	})

	t.Run("allows commands once gpg-id exists", func(t *testing.T) {
		err := os.WriteFile(filepath.Join(homeDir, ".gpg-id"), []byte("test@example.com\n"), 0600)
		assert.NoError(t, err)

		cmd := &cobra.Command{Use: "ls"}
		err = PersistentPreRunE(cmd, []string{})
		assert.NoError(t, err)
	})
}
