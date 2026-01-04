package path

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/nicola-strappazzon/password-manager/config"
)

// personal/                                             <-- Directory, cli
// personal/passport                                     <-- Path,      cli
// /Users/fulano/.password-manager/personal/             <-- Absolute,  tree
// /Users/fulano/.password-manager/personal/passport.gpg <-- Full,      decript/encript
// passport.gpg                                          <-- Name
// passport                                              <-- Base Name
// gpg                                                   <-- Extencion

type Path string

func (p Path) Path() string {
	return string(p)
}

func (p Path) Directory() string {
	return filepath.Dir(p.Path())
}

func (p Path) Name() string {
	return fmt.Sprintf("%s.gpg", filepath.Base(p.Path()))
}

func (p Path) BaseName() string {
	base := filepath.Base(p.Path())
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func (p Path) Full() string {
	return filepath.Clean(path.Join(config.GetPath(""), p.Directory(), p.Name()))
}

func (p Path) Absolute() string {
	return filepath.Clean(path.Join(config.GetPath(""), p.Directory()))
}

func (p Path) IsDirectory() bool {
	info, err := os.Stat(p.Absolute())
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (p Path) IsNotFile() bool {
	return !p.IsFile()
}

func (p Path) IsFile() bool {
	info, err := os.Stat(p.Full())
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

func (p Path) IsInvalid() bool {
	for _, r := range p.Path() {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '/' || r == '-' || r == '_' {
			continue
		}
		return true
	}

	return false
}
