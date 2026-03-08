package file

import (
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/password-manager/check"
)

func Name(in string) string {
	return filepath.Base(in)
}

func ReadInString(in string) string {
	return string(ReadInBytes(in))
}

func ReadInBytes(in string) []byte {
	data, err := os.ReadFile(in)
	check.Check(err)

	return data
}

func Save(path string, content []byte) {
	os.MkdirAll(filepath.Dir(path), 0700)
	check.Check(os.WriteFile(path, content, 0600))
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
