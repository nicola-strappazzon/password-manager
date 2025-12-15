package config

import (
	"os"
	"path"

	"github.com/nicola-strappazzon/pm/check"
	"github.com/nicola-strappazzon/pm/env"
)

const DATA_DIR = ".password-manager"

var IGNORE_DIRS = []string{".git", ".public-keys"}

func init() {

}

func GetHomeDir() string {
	base, err := os.UserHomeDir()
	check.Check(err)

	return base
}

func GetDataDirectoryFrom(in string) string {
	return path.Join(GetHomeDir(), env.Get("PM_PATH", DATA_DIR), in)
}

func GetDataDirectory() string {
	return GetDataDirectoryFrom("")
}

func GetPrivateKeyPath() string {
	return env.Get("PM_PRIVATEKEY", "")
}

func GetPublicKeyPath() string {
	return env.Get("PM_PUBLICKEY", "")
}
