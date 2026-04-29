package openpgp

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

func runCommand(input []byte, name string, args ...string) ([]byte, error) {
	var stderr bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdin = bytes.NewReader(input)
	cmd.Stderr = &stderr

	return cmd.Output()
}

func GetRecipientKeyID(filePath string) (string, error) {
	var keyIDRegex = regexp.MustCompile(`(?m)^:pubkey enc packet:.*keyid\s+([0-9A-F]+)`)

	out, err := runCommand(nil, "gpg", "--list-only", "--list-packets", "--batch", "--no-tty", filePath)
	if err != nil {
		return "", err
	}

	match := keyIDRegex.FindSubmatch(out)
	if match == nil {
		return "", fmt.Errorf("keyid not found in: %s", filePath)
	}

	return string(match[1]), nil
}

func RequiresSmartCard(keyID string) (bool, error) {
	out, err := runCommand(nil, "gpg", "--with-colons", "--list-secret-keys", keyID)
	if err != nil {
		return false, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		fields := strings.Split(line, ":")
		if len(fields) < 15 || !strings.HasPrefix(fields[0], "ssb") {
			continue
		}
		if fields[11] == "e" && fields[14] != "+" {
			return true, nil
		}
	}

	return false, nil
}

func CardStatus() error {
	_, err := runCommand(nil, "gpg", "--card-status")
	return err
}

func CardIsReady() error {
	if err := CardStatus(); err != nil {
		return fmt.Errorf("no smartcard detected")
	}

	out, err := runCommand(
		nil,
		"gpg-connect-agent",
		"SCD SERIALNO",
		"/bye",
	)

	if err != nil {
		return fmt.Errorf("no smartcard detected")
	}

	if !strings.Contains(string(out), "SERIALNO") {
		return fmt.Errorf("smartcard not responding")
	}

	return nil
}

func Encrypt(in, recipient string) (out []byte, err error) {
	out, err = runCommand(
		[]byte(in),
		"gpg",
		"--batch",
		"--yes",
		"--quiet",
		"--encrypt",
		"--armor",
		"--recipient", recipient,
	)
	return
}

func Decrypt(passphrase, filePath string) (out string, err error) {
	var stdout bytes.Buffer

	cmd := exec.Command(
		"gpg",
		"--batch",
		"--yes",
		"--quiet",
		"--no-tty",
		"--pinentry-mode", "loopback",
		"--passphrase-fd", "0",
		"--decrypt", filePath,
	)

	cmd.Stdin = strings.NewReader(passphrase + "\n")
	cmd.Stdout = &stdout
	cmd.Stderr = io.Discard

	err = cmd.Run()
	out = stdout.String()

	return
}
