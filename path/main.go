package path

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	// "regexp"

	"github.com/nicola-strappazzon/pm/config"
)

// personal/                                             <-- Directory, cli
// personal/passport                                     <-- Path,      cli
// /Users/fulano/.password-manager/personal/             <-- Absolute,  tree
// /Users/fulano/.password-manager/personal/passport.gpg <-- Full,      decript/encript
// passport.gpg                                          <-- Name
// passport                                              <-- Base Name
// gpg                                                   <-- Extencion

type Path string

// var validName = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

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
	return filepath.Clean(path.Join(config.GetDataDirectory(), p.Directory(), p.Name()))
}

func (p Path) Absolute() string {
	return filepath.Clean(path.Join(config.GetDataDirectory(), p.Directory()))
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

// func (p Path) IsInvalid() bool {
// 	if p.Path() == "." || p.Path() == "/" {
// 		return false
// 	}

// 	parts := strings.Split(p.Path(), string(filepath.Separator))
// 	for _, part := range parts {
// 		if part == "" {
// 			return false
// 		}
// 		if !validName.MatchString(part) {
// 			return false
// 		}
// 	}

// 	return true
// }
