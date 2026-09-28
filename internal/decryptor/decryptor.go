package decryptor

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/term"
)

func Decrypt(passphrase, path string) (card.Card, error) {
	cardReady := openpgp.CardIsReady() == nil

	cardID, err := openpgp.GetRecipientKeyID(path)
	if err != nil {
		return card.Card{}, err
	}

	belongs, err := openpgp.KeyBelongsToRecipient(cardID, config.GetRecipient())
	if err != nil {
		return card.Card{}, err
	}
	if !belongs {
		return card.Card{}, fmt.Errorf("This file is not encrypted for the configured recipient (%s).", config.GetRecipient())
	}

	useCard, err := openpgp.RequiresSmartCard(cardID)
	if err != nil {
		return card.Card{}, err
	}

	if useCard && !cardReady {
		return card.Card{}, fmt.Errorf("This file requires a smartcard to decrypt.")
	}

	if useCard {
		passphrase = term.ReadPassword("Card PIN: ", passphrase)
	} else {
		passphrase = term.ReadPassword("Passphrase: ", passphrase)
	}

	if passphrase == "" {
		if useCard {
			return card.Card{}, fmt.Errorf("Card PIN cannot be empty.")
		}
		return card.Card{}, fmt.Errorf("Passphrase cannot be empty.")
	}

	fileContent, err := openpgp.Decrypt(
		passphrase,
		path,
	)

	return card.New(fileContent), err
}
