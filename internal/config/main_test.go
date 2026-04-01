package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecipient(t *testing.T) {
	home := t.TempDir()
	oldUserHomeDir := UserHomeDir
	oldDataDir := DataDir
	oldGPGIDFile := GPGIDFile
	UserHomeDir = func() (string, error) { return home, nil }
	DataDir = ".pm-test"
	GPGIDFile = ".gpg-id"
	t.Cleanup(func() {
		UserHomeDir = oldUserHomeDir
		DataDir = oldDataDir
		GPGIDFile = oldGPGIDFile
	})

	t.Run("returns empty when file does not exist", func(t *testing.T) {
		assert.Equal(t, "", GetRecipient())
	})

	t.Run("reads and trims gpg id file", func(t *testing.T) {
		err := os.MkdirAll(filepath.Join(home, DataDir), 0700)
		assert.NoError(t, err)
		err = os.WriteFile(filepath.Join(home, DataDir, GPGIDFile), []byte("test@example.com\n"), 0600)
		assert.NoError(t, err)
		assert.Equal(t, "test@example.com", GetRecipient())
	})
}

func TestSaveRecipient(t *testing.T) {
	home := t.TempDir()
	oldUserHomeDir := UserHomeDir
	oldDataDir := DataDir
	oldGPGIDFile := GPGIDFile
	UserHomeDir = func() (string, error) { return home, nil }
	DataDir = ".pm-test"
	GPGIDFile = ".gpg-id"
	t.Cleanup(func() {
		UserHomeDir = oldUserHomeDir
		DataDir = oldDataDir
		GPGIDFile = oldGPGIDFile
	})

	err := SaveRecipient("demo@example.com")
	assert.NoError(t, err)
	data, err := os.ReadFile(filepath.Join(home, DataDir, GPGIDFile))
	assert.NoError(t, err)
	assert.Equal(t, "demo@example.com\n", string(data))

	err = SaveRecipient("   ")
	assert.Error(t, err)
}
