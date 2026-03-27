package decryptor

import (
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/term"
)

func Decrypt(passphrase, path string) (card.Card, error) {
	var fileContent string

	if openpgp.CardStatus() {
		if err := openpgp.CardIsReady(); err == nil {
			fileContent, err = openpgp.DecryptWithCard(
				term.ReadPassword("Card PIN: ", passphrase),
				path,
			)

			if err != nil {
				return card.Card{}, err
			}
		}
	} else {
		fileContent = openpgp.Decrypt(
			term.ReadPassword("Passphrase: ", passphrase),
			path,
		)
	}

	return card.New(fileContent), nil
}
