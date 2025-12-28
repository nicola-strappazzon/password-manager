package explorer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/password-manager/config"
)

func Directories() (dirs []string, err error) {
	basePath := config.GetPath("")

	err = filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			return nil
		}

		rel, _ := filepath.Rel(basePath, path)
		if rel == "." {
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") {
			return filepath.SkipDir
		}

		dirs = append(dirs, rel+"/")

		return nil
	})

	return
}

func DirectoriesAndFiles() (list []string, err error) {
	basePath := config.GetPath("")

	err = filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(basePath, path)
		if rel == "." {
			return nil
		}

		rel = filepath.ToSlash(rel)

		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}

			if !strings.HasSuffix(rel, "/") {
				rel += "/"
			}

			list = append(list, rel)
			return nil
		}

		if strings.HasSuffix(d.Name(), ".gpg") {
			noExt := strings.TrimSuffix(rel, ".gpg")
			list = append(list, noExt)
		}

		return nil
	})

	return
}

func PrintTree(root string) error {
	var lastAtDepth = map[int]bool{}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
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

		if path == root {
			fmt.Printf("\033[1;37m%s\033[0m\n", "Password Store")
			return nil
		}

		rel, _ := filepath.Rel(root, path)
		depth := strings.Count(rel, string(filepath.Separator))

		prefix := ""
		for i := 0; i < depth; i++ {
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
			fmt.Println(prefix + "└── " + fileName)
			lastAtDepth[depth] = true
		} else if d.IsDir() {
			fmt.Println(prefix + "├── " + fmt.Sprintf("\033[1;37m%s\033[0m", d.Name()))
			lastAtDepth[depth] = false
		} else {
			fileName := strings.TrimSuffix(d.Name(), ".gpg")
			fmt.Println(prefix + "├── " + fileName)
			lastAtDepth[depth] = false
		}

		return nil
	})
}
