package config

import (
	"path"

	"github.com/nicola-strappazzon/pm/env"
)

const DATA_DIR = ".password-manager"

func GetPath(in string) string {
	return path.Join(env.Get("PM_PATH", DATA_DIR), in)
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
