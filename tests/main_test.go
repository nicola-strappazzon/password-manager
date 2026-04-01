package integration_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/spf13/cobra"
)

const testPassphrase = "test-passphrase"
const testRecipient = "test@example.com"

func TestMain(m *testing.M) {
	if _, err := exec.LookPath("gpg"); err != nil {
		os.Setenv("PATH", "/opt/homebrew/bin:/usr/local/bin:"+os.Getenv("PATH"))
	}

	gnupghomeDir, err := os.MkdirTemp("", "pm-gnupghome-*")

	if err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.Exit(1)
	}

	vaultDir, err := os.MkdirTemp("", "pm-vault-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.RemoveAll(gnupghomeDir)
		os.Exit(1)
	}

	os.Setenv("GNUPGHOME", gnupghomeDir)

	if err := setupGPG(gnupghomeDir); err != nil {
		fmt.Fprintln(os.Stderr, "key generation error:", err)
		os.RemoveAll(gnupghomeDir)
		os.RemoveAll(vaultDir)
		os.Exit(1)
	}

	config.DataDir = ""
	config.UserHomeDir = func() (string, error) { return vaultDir, nil }
	if err := os.WriteFile(filepath.Join(vaultDir, config.GPGIDFile), []byte(testRecipient+"\n"), 0600); err != nil {
		fmt.Fprintln(os.Stderr, "setup error:", err)
		os.RemoveAll(gnupghomeDir)
		os.RemoveAll(vaultDir)
		os.Exit(1)
	}

	code := m.Run()

	os.Unsetenv("GNUPGHOME")
	os.RemoveAll(gnupghomeDir)
	os.RemoveAll(vaultDir)

	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	t.Run("version", testVersion)
	t.Run("add", testAdd)
	t.Run("ls", testLs)
	t.Run("move", testMove)
	t.Run("update", testUpdate)
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

func setupGPG(dir string) error {
	params := "Key-Type: RSA\n" +
		"Key-Length: 2048\n" +
		"Subkey-Type: RSA\n" +
		"Subkey-Length: 2048\n" +
		"Name-Real: Test\n" +
		"Name-Email: " + testRecipient + "\n" +
		"Expire-Date: 0\n" +
		"Passphrase: " + testPassphrase + "\n" +
		"%commit\n"

	paramsFile := filepath.Join(dir, "keygen.params")
	if err := os.WriteFile(paramsFile, []byte(params), 0600); err != nil {
		return err
	}

	cmd := exec.Command("gpg", "--batch", "--gen-key", paramsFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("gpg keygen failed: %s: %w", out, err)
	}

	return nil
}
