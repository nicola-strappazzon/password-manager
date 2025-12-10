package openpgp

import (
	"log"

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
	check(err)

	decHandle, err := pgp.Decryption().DecryptionKey(privateKey).New()
	defer decHandle.ClearPrivateParams()
	check(err)

	decrypted, err := decHandle.Decrypt(
		file.ReadInBytes(path),
		crypto.Bytes,
	)
	check(err)

	return decrypted.String()
}

func check(in error) {
	if in != nil {
		log.Fatal(in.Error())
	}
}
