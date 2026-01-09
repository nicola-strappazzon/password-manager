package version_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/version"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand_Metadata(t *testing.T) {
	cmd := version.NewCommand()

	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "Print version number", cmd.Short)
}

func TestNewCommand_PrintsVersion(t *testing.T) {
	cmd := version.NewCommand()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	err := cmd.Execute()
	assert.NoError(t, err)

	out, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, version.VERSION, strings.TrimSpace(string(out)))
}
