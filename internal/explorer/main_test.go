package explorer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"

	"github.com/stretchr/testify/assert"
)

func TestDirectoriesAndFiles(t *testing.T) {
	setupStore(t)
	items, err := explorer.DirectoriesAndFiles()

	assert.NoError(t, err)
	assert.Equal(t, []string{"personal/", "personal/site", "root"}, items)
}

func TestDirectories(t *testing.T) {
	setupStore(t)
	items, err := explorer.Directories()

	assert.NoError(t, err)
	assert.Equal(t, []string{"personal/"}, items)
}

func setupStore(t *testing.T) {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
	if err := os.MkdirAll(config.GetPath("personal"), 0700); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"personal/site.gpg", "root.gpg", ".hidden.gpg", "notes.txt"} {
		if err := os.WriteFile(filepath.Join(config.GetPath(""), name), []byte("opaque ciphertext"), 0600); err != nil {
			t.Fatal(err)
		}
	}
}
