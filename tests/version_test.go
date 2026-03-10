package integration_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/version"
	"github.com/stretchr/testify/assert"
)

func testVersion(t *testing.T) {
	version.VERSION = "v0.0.0-test"
	stdout, _, err := run(version.NewCommand(), []string{})
	assert.NoError(t, err)
	assert.Equal(t, stdout, "v0.0.0-test\n")
}
