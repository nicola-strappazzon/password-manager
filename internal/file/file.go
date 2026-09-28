package file

import (
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/password-manager/internal/check"
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

func Remove(path string) error {
	return os.Remove(path)
}

func RemoveEmptyParents(filePath, stopAt string) {
	dir := filepath.Dir(filePath)
	for dir != stopAt {
		entries, err := os.ReadDir(dir)
		if err != nil || len(entries) > 0 {
			break
		}
		if err := os.Remove(dir); err != nil {
			break
		}
		dir = filepath.Dir(dir)
	}
}
