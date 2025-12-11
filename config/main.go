package config

import (
	"log"
	"os"
	"path"

	"github.com/nicola-strappazzon/pm/env"
)

const WORKDIR = ".password-manager"

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
	return path.Join(GetHomeDir(), env.Get("PM_PATH", WORKDIR), in)
}

func GetWorkDirectory() string {
	return GetWorkDirectoryFrom("")
}

func GetPrivateKeyPath() string {
	return env.Get("PM_PRIVATEKEY", "")
}
