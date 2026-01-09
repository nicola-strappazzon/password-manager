package config

import (
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/password-manager/env"
)

var DataDir = ".password-manager"
var UserHomeDir = os.UserHomeDir

func GetPath(in string) string {
	home, err := UserHomeDir()
	if err != nil {
		panic(err)
	}

	base := filepath.Join(home, DataDir)
	return filepath.Join(base, in)
}

func GetPrivateKey() string {
	return env.Get("PM_PRIVATEKEY", "")
}

func GetPublicKey() string {
	return env.Get("PM_PUBLICKEY", "")
}

func HasNotPrivateKey() bool {
	return GetPrivateKey() == ""
}

func HasNotPublicKey() bool {
	return GetPublicKey() == ""
}
