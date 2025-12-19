package file

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nicola-strappazzon/pm/check"
	"github.com/nicola-strappazzon/pm/config"
)

var validName = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

func Valid(in string) bool {
	if in == "." || in == "/" {
		return false
	}

	parts := strings.Split(in, string(filepath.Separator))
	for _, part := range parts {
		if part == "" {
			return false
		}
		if !validName.MatchString(part) {
			return false
		}
	}

	return true
}

func AbsolutePath(in string) string {
	return filepath.Clean(path.Join(config.GetDataDirectory(), in))
}

func Exist(in string) bool {
	if _, err := os.Stat(in); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func IsDir(in string) bool {
	info, err := os.Stat(in)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func AddExt(in string) string {
	ext := filepath.Ext(in)

	if ext == ".gpg" {
		return in
	}

	return fmt.Sprintf("%s.gpg", in)
}

func Name(in string) string {
	return strings.TrimSuffix(filepath.Base(in), filepath.Ext(in))
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
