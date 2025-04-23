package config

import (
	"log"
	"os"
	"path"
)

var PASSWORD_STORE_DIR = ".password-store"
var IGNORE_DIRS = []string{".git", ".public-keys"}

func GetHomeDir() string {
	base, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return base
}

func GetWorkDirectory(in string) string {
	return path.Join(GetHomeDir(), PASSWORD_STORE_DIR, in)
}
