package explorer

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/store"
)

func Directories() (dirs []string, err error) {
	entries, err := store.List(config.GetPath(""))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.Directory {
			dirs = append(dirs, entry.Path+"/")
		}
	}
	return dirs, nil
}

func DirectoriesAndFiles() (list []string, err error) {
	entries, err := store.List(config.GetPath(""))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		path := entry.Path
		if entry.Directory {
			path += "/"
		}
		list = append(list, path)
	}
	return list, nil
}

func PrintTree(root string) (out string, err error) {
	var lastAtDepth = map[int]bool{}

	return out, filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path == root {
			out += fmt.Sprintf("\033[1;37m%s\033[0m\n", "Password Store")
			return nil
		}

		if d.IsDir() && strings.HasPrefix(d.Name(), ".") && path != root {
			return filepath.SkipDir
		}

		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		if d.Name() == "Makefile" {
			return nil
		}

		rel, _ := filepath.Rel(root, path)
		depth := strings.Count(rel, string(filepath.Separator))

		prefix := ""
		for i := range depth {
			if lastAtDepth[i] {
				prefix += "    "
			} else {
				prefix += "│   "
			}
		}

		isLast := true
		parent := filepath.Dir(path)
		filepath.WalkDir(parent, func(p string, de fs.DirEntry, _ error) error {
			if p != parent && filepath.Dir(p) == parent {
				if p > path {
					isLast = false
				}
			}
			return nil
		})

		if isLast {
			fileName := strings.TrimSuffix(d.Name(), ".gpg")
			lastAtDepth[depth] = true

			if d.IsDir() {
				out += fmt.Sprintf("%s%s \033[1;37m%s\033[0m\n", prefix, "└──", d.Name())
			} else {
				out += fmt.Sprintf("%s%s %s\n", prefix, "└──", fileName)
			}
		} else if d.IsDir() {
			out += fmt.Sprintf("%s%s \033[1;37m%s\033[0m\n", prefix, "├──", d.Name())
			lastAtDepth[depth] = false
		} else {
			fileName := strings.TrimSuffix(d.Name(), ".gpg")
			out += fmt.Sprintf("%s%s %s\n", prefix, "├──", fileName)
			lastAtDepth[depth] = false
		}

		return nil
	})
}
