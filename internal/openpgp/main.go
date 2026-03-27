package openpgp

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/nicola-strappazzon/password-manager/internal/check"
	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/file"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
)

func runCommand(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = io.Discard
	return cmd.Output()
}

func Decrypt(passphrase, path string) string {
	var pgp = crypto.PGP()

	privateKey, err := crypto.NewPrivateKeyFromArmored(
		file.ReadInString(
			config.GetPrivateKey(),
		),
		[]byte(passphrase),
	)
	check.Check(err)

	decHandle, err := pgp.Decryption().DecryptionKey(privateKey).New()
	defer decHandle.ClearPrivateParams()

	check.Check(err)

	decrypted, err := decHandle.Decrypt(
		file.ReadInBytes(path),
		crypto.Bytes,
	)
	check.Check(err)

	return decrypted.String()
}

func Encrypt(in string) []byte {
	var pgp = crypto.PGP()

	publicKey, err := crypto.NewKeyFromArmored(
		file.ReadInString(
			config.GetPublicKey(),
		),
	)
	check.Check(err)

	encHandle, err := pgp.Encryption().Recipient(publicKey).New()
	check.Check(err)

	pgpMessage, err := encHandle.Encrypt([]byte(in))
	check.Check(err)

	return pgpMessage.Bytes()
}

func CardStatus() bool {
	_, err := runCommand("gpg", "--card-status")
	return err == nil
}

func CardIsReady() error {
	out, err := runCommand(
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

func DecryptWithCard(passphrase, filePath string) (string, error) {
	var out bytes.Buffer

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
	cmd.Stdout = &out
	cmd.Stderr = io.Discard

	err := cmd.Run()

	return out.String(), err
}
