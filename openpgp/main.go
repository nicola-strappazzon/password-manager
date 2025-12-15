package openpgp

import (
	"github.com/nicola-strappazzon/pm/check"
	"github.com/nicola-strappazzon/pm/config"
	"github.com/nicola-strappazzon/pm/file"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
)

func Decrypt(passphrase, path string) string {
	var pgp = crypto.PGP()

	privateKey, err := crypto.NewPrivateKeyFromArmored(
		file.ReadInString(
			config.GetPrivateKeyPath(),
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
			config.GetPublicKeyPath(),
		),
	)
	check.Check(err)

	encHandle, err := pgp.Encryption().Recipient(publicKey).New()
	check.Check(err)

	pgpMessage, err := encHandle.Encrypt([]byte(in))
	check.Check(err)

	return pgpMessage.Bytes()
}
