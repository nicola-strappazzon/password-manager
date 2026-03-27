package integration_test

import (
	"testing"
)

func testAdd(t *testing.T) {
	t.Run("username", func(t *testing.T) {
		// stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "username", "-v", "user-demo"})
		// assert.Empty(t, stdout)
		// t.Log(stderr)
		// t.Log(err)
		// assert.Empty(t, stderr)
		// assert.NoError(t, err)
	})

	t.Run("password", func(t *testing.T) {
		// stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "password", "-v", "MyStr0ngP@s3w0rd", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.NoError(t, err)
		// assert.Contains(t, stderr, "Warning: Using a password on the command line interface can be insecure.")
	})

	t.Run("url", func(t *testing.T) {
		// stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "url", "-v", "https://github.com", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.Empty(t, stderr)
		// assert.NoError(t, err)
	})

	t.Run("otp", func(t *testing.T) {
		// stdout, stderr, err := run(add.NewCommand(), []string{"github", "-f", "otp", "-v", "246EOSQ2ORPTQRWS", "-p", testPassphrase})
		// assert.Empty(t, stdout)
		// assert.Empty(t, stderr)
		// assert.NoError(t, err)
	})

	// t.Run("field-already-exists", func(t *testing.T) {
	// 	stdout, _, err := run(add.NewCommand(), []string{"github", "-f", "username", "-v", "other-user", "-p", testPassphrase})
	// 	assert.Empty(t, stdout)
	// 	assert.EqualError(t, err, "Field 'username' already exists. Use 'pm update' to modify it.")
	// })
}
