package openpgp

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func runCommand(input []byte, name string, args ...string) ([]byte, error) {
	var stderr bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdin = bytes.NewReader(input)
	cmd.Stderr = &stderr

	return cmd.Output()
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
