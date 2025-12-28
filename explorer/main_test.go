package explorer_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/explorer"

	"github.com/stretchr/testify/assert"
)

func TestDirectoriesAndFiles(t *testing.T) {

	items, err := explorer.DirectoriesAndFiles()

	t.Log(items)

	assert.NoError(t, err)

	assert.True(t, true)
}

func TestDirectories(t *testing.T) {
	items, err := explorer.Directories()

	t.Log(items)

	assert.NoError(t, err)

	assert.True(t, true)
}
