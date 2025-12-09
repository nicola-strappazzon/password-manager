package config

import (
	"log"
	"os"
	"path"

	"github.com/nicola-strappazzon/pm/env"
)

const PASSWORD_STORE_DIR = ".password-store"

var IGNORE_DIRS = []string{".git", ".public-keys"}

func init() {

}

func GetHomeDir() string {
	base, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return base
}

func GetWorkDirectoryFrom(in string) string {
	return path.Join(GetHomeDir(), env.Get("PM_PATH", PASSWORD_STORE_DIR), in)
}

func GetWorkDirectory() string {
	return GetWorkDirectoryFrom("")
}

func GetPrivateKeyPath() string {
	return "/Users/nicola/Documents/documents/documents/Keys/gpg/private.pgp"
}
