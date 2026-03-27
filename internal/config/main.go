package config

import (
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/password-manager/internal/env"
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

func GetRecipient() string {
	return env.Get("PM_RECIPIENT", "")
}

func HasNotRecipient() bool {
	return GetRecipient() == ""
}
