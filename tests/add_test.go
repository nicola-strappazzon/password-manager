package integration_test

import (
	"os"
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/add"
	"github.com/nicola-strappazzon/password-manager/internal/decryptor"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/stretchr/testify/assert"
)

func testAdd(t *testing.T) {
	t.Run("username", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "username", "-v", "user-demo"})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("password", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "password", "-v", "MyStr0ngP@s3w0rd", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.NoError(t, err)
		assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
	})

	t.Run("url", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "url", "-v", "https://github.com", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("phone", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "phone", "-v", "+34123456789", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("puk", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "puk", "-v", "12345678", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)

		card, err := decryptor.Decrypt(testPassphrase, path.Path("github").Full())
		assert.NoError(t, err)
		assert.Equal(t, "+34123456789", card.Phone)
		assert.Equal(t, "12345678", card.PUK)
	})

	t.Run("otp", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "otp", "-v", "246EOSQ2ORPTQRWS", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)
	})

	t.Run("otp-with-whitespace-and-hyphens", func(t *testing.T) {
		stdout, stderr, err := run(add.NewCommand(), []string{"github_whitespace_otp", "-f", "otp", "-v", "246E-OSQ2\tORPT-QRWS", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.Empty(t, stderr)
		assert.NoError(t, err)

		card, err := decryptor.Decrypt(testPassphrase, path.Path("github_whitespace_otp").Full())
		assert.NoError(t, err)
		assert.Equal(t, "246EOSQ2ORPTQRWS", card.OTP)
		assert.NoError(t, os.Remove(path.Path("github_whitespace_otp").Full()))
	})

	t.Run("field-already-exists", func(t *testing.T) {
		stdout, _, err := run(add.NewCommand(), []string{"github", "-f", "username", "-v", "other-user", "-p", testPassphrase})
		assert.Empty(t, stdout)
		assert.EqualError(t, err, "Field 'username' already exists. Use 'pm update' to modify it.")
	})
}
