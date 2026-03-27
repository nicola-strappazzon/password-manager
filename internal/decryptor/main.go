package decryptor

import (
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/term"
)

func Decrypt(passphrase, path string) (card.Card, error) {
	if err := openpgp.CardIsReady(); err == nil {
		passphrase = term.ReadPassword("Card PIN: ", passphrase)
	} else {
		passphrase = term.ReadPassword("Passphrase: ", passphrase)
	}

	fileContent, err := openpgp.Decrypt(
		passphrase,
		path,
	)

	return card.New(fileContent), err
}
