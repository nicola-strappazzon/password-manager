package decryptor

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/term"
)

func Decrypt(passphrase, path string) (card.Card, error) {
	cardReady := openpgp.CardIsReady() == nil

	cardID, err := openpgp.GetRecipientKeyID(path)
	if err != nil {
		return card.Card{}, err
	}

	useCard, err := openpgp.RequiresSmartCard(cardID)
	if err != nil {
		return card.Card{}, err
	}

	if useCard && !cardReady {
		return card.Card{}, fmt.Errorf("This file requires a smartcard to decrypt.")
	}

	if cardReady && useCard {
		passphrase = term.ReadPassword("Card PIN: ", passphrase)
	} else if !cardReady {
		passphrase = term.ReadPassword("Passphrase: ", passphrase)
	}

	fileContent, err := openpgp.Decrypt(
		passphrase,
		path,
	)

	return card.New(fileContent), err
}
