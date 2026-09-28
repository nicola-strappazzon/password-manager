package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var DataDir = ".password-manager"
var UserHomeDir = os.UserHomeDir
var GPGIDFile = ".gpg-id"

func GetPath(in string) string {
	home, err := UserHomeDir()
	if err != nil {
		panic(err)
	}

	base := filepath.Join(home, DataDir)
	return filepath.Join(base, in)
}

func GetRecipient() string {
	data, err := os.ReadFile(GetPath(GPGIDFile))
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func SaveRecipient(in string) error {
	path := GetPath(GPGIDFile)

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	content := strings.TrimSpace(in)
	if content == "" {
		return fmt.Errorf("recipient cannot be empty")
	}

	return os.WriteFile(path, []byte(content+"\n"), 0600)
}

func HasNotRecipient() bool {
	return GetRecipient() == ""
}
