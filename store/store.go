// Package store lists the directory structure and encrypted item names of a
// password-manager store without reading or decrypting their contents.
package store

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Directory bool   `json:"directory"`
}

var ErrInvalidPath = errors.New("invalid store path")

func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".password-manager"), nil
}

func validPath(path string) bool {
	if path == "." {
		return true
	}
	if !filepath.IsLocal(path) {
		return false
	}
	for _, part := range strings.Split(filepath.ToSlash(path), "/") {
		if part == "" || strings.HasPrefix(part, ".") {
			return false
		}
	}
	return true
}

// List recursively lists visible directories and .gpg items. Paths are relative
// to the store; item paths omit the .gpg suffix, matching the CLI.
func List(rootPath string) ([]Entry, error) {
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return nil, err
	}
	defer root.Close()
	entries := make([]Entry, 0)
	err = fs.WalkDir(root.FS(), ".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == "." {
			return nil
		}
		if strings.HasPrefix(entry.Name(), ".") {
			if entry.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return nil
		}
		if entry.IsDir() {
			entries = append(entries, Entry{Name: entry.Name(), Path: path, Directory: true})
		} else if entry.Type().IsRegular() && strings.HasSuffix(entry.Name(), ".gpg") {
			entries = append(entries, Entry{Name: strings.TrimSuffix(entry.Name(), ".gpg"), Path: strings.TrimSuffix(path, ".gpg")})
		}
		return nil
	})
	return entries, err
}

// Items lists only immediate encrypted items in the selected directory.
func Items(rootPath, directory string) ([]Entry, error) {
	if directory == "" {
		directory = "."
	}
	if !validPath(directory) {
		return nil, ErrInvalidPath
	}
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return nil, err
	}
	defer root.Close()
	current := ""
	for _, part := range strings.Split(filepath.ToSlash(directory), "/") {
		current = filepath.Join(current, part)
		info, err := root.Lstat(current)
		if err != nil {
			return nil, err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil, ErrInvalidPath
		}
	}
	file, err := root.Open(directory)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	children, err := file.ReadDir(-1)
	if err != nil {
		return nil, err
	}
	items := make([]Entry, 0)
	for _, entry := range children {
		if !entry.Type().IsRegular() || strings.HasPrefix(entry.Name(), ".") || !strings.HasSuffix(entry.Name(), ".gpg") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".gpg")
		items = append(items, Entry{Name: name, Path: filepath.ToSlash(filepath.Join(directory, name))})
	}
	return items, nil
}
