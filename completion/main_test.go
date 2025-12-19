package completion_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/completion"

	"github.com/stretchr/testify/assert"
)

func TestDirectoriesAndFiles(t *testing.T) {

	items, err := completion.DirectoriesAndFiles()

	t.Log(items)

	assert.NoError(t, err)

	assert.True(t, true)
}

func TestDirectories(t *testing.T) {
	items, err := completion.Directories()

	t.Log(items)

	assert.NoError(t, err)

	assert.True(t, true)
}
