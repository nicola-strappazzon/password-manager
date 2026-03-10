package integration_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
	"github.com/ProtonMail/gopenpgp/v3/profile"
	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/spf13/cobra"
)

const testPassphrase = "test-passphrase"

func TestMain(m *testing.M) {
	keysDir, err := os.MkdirTemp("", "pm-keys-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.Exit(1)
	}

	vaultDir, err := os.MkdirTemp("", "pm-vault-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.Exit(1)
	}

	if err := generateKeys(keysDir); err != nil {
		fmt.Fprintln(os.Stderr, "key generation error:", err)
		os.RemoveAll(keysDir)
		os.RemoveAll(vaultDir)
		os.Exit(1)
	}

	config.DataDir = ""
	config.UserHomeDir = func() (string, error) { return vaultDir, nil }

	os.Setenv("PM_PUBLICKEY", filepath.Join(keysDir, "public.asc"))
	os.Setenv("PM_PRIVATEKEY", filepath.Join(keysDir, "private.asc"))

	code := m.Run()

	os.Unsetenv("PM_PUBLICKEY")
	os.Unsetenv("PM_PRIVATEKEY")
	os.RemoveAll(keysDir)
	os.RemoveAll(vaultDir)

	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	t.Run("version", testVersion)
	t.Run("add", testAdd)
	t.Run("ls", testLs)
	t.Run("move", testMove)
	t.Run("generate", testGenerate)
}

func run(cmd *cobra.Command, args []string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return stdout.String(), stderr.String(), err
}

func generateKeys(dir string) error {
	pgp := crypto.PGPWithProfile(profile.Default())
	key, err := pgp.KeyGeneration().AddUserId("Test", "test@example.com").New().GenerateKey()
	if err != nil {
		return err
	}
	defer key.ClearPrivateParams()

	lockedKey, err := crypto.PGP().LockKey(key, []byte(testPassphrase))
	if err != nil {
		return err
	}

	pubKey, err := key.ToPublic()
	if err != nil {
		return err
	}

	pubArmored, err := pubKey.Armor()
	if err != nil {
		return err
	}

	privArmored, err := lockedKey.Armor()
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dir, "public.asc"), []byte(pubArmored), 0600); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, "private.asc"), []byte(privArmored), 0600)
}
