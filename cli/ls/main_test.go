package ls_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/ls"
	"github.com/nicola-strappazzon/password-manager/config"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.DataDir = ""

	orig := config.UserHomeDir
	defer func() { config.UserHomeDir = orig }()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.Exit(1)
	}

	config.UserHomeDir = func() (string, error) {
		return fmt.Sprintf("%s/../../testdata/vault", wd), nil
	}

	os.Setenv("PM_PUBLICKEY", "testdata/gpg/pubkey.asc")

	code := m.Run()

	os.Unsetenv("PM_PUBLICKEY")
	os.Exit(code)
}

func TestCommand(t *testing.T) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	cmd := ls.NewCommand()
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	err := cmd.Execute()
	assert.NoError(t, err)

	assert.Empty(t, stderr.String())
	assert.Equal(t, "\033[1;37mPassword Store\033[0m\n└── google\n", stdout.String())
}
