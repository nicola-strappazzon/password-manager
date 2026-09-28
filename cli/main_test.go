package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompletionWithoutSetup(t *testing.T) {
	homeDir := t.TempDir()
	oldUserHomeDir := config.UserHomeDir
	config.UserHomeDir = func() (string, error) { return homeDir, nil }
	t.Cleanup(func() { config.UserHomeDir = oldUserHomeDir })

	for _, shell := range []string{"bash", "zsh"} {
		t.Run(shell, func(t *testing.T) {
			output, err := os.CreateTemp(t.TempDir(), "completion")
			require.NoError(t, err)
			oldStdout := os.Stdout
			os.Stdout = output
			t.Cleanup(func() {
				os.Stdout = oldStdout
				output.Close()
			})

			cmd := Load()
			cmd.SetArgs([]string{"completion", shell})
			require.NoError(t, cmd.Execute())

			script, err := os.ReadFile(output.Name())
			require.NoError(t, err)
			assert.Contains(t, string(script), "completion for pm")
			_, err = os.Stat(filepath.Join(homeDir, config.DataDir))
			assert.True(t, os.IsNotExist(err), "completion must not create a password store")
		})
	}
}

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
