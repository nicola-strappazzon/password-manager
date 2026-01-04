package arguments_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/arguments"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	assert.Equal(t, "", arguments.First([]string{""}))
	assert.Equal(t, "foo/path", arguments.First([]string{"foo/path"}))
	assert.Equal(t, "foo/path", arguments.First([]string{"foo/path", "foo/", ""}))
}
