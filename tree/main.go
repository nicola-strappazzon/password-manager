package tree

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicola-strappazzon/pm/config"
)

func ignoreDir(in string) bool {
	for _, item := range config.IGNORE_DIRS {
		if in == item {
			return true
		}
	}
	return false
}

// is file?
func validFile(in fs.DirEntry) bool {
	if in.IsDir() {
		return false
	}

	return !validExtension(in.Name()) // esto se remueve y se compina con otra funcion...
}

func validExtension(in string) bool {
	return strings.HasSuffix(in, ".gpg")
}

func removeExtension(in string) string {
	return strings.TrimSuffix(in, ".gpg")
}

func build(path string) (*Node, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	nodo := &Node{
		Name:  removeExtension(filepath.Base(path)),
		IsDir: info.IsDir(),
	}

	if !info.IsDir() {
		return nodo, nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if ignoreDir(file.Name()) {
			continue
		}

		if validFile(file) {
			continue
		}

		childPath := filepath.Join(path, file.Name())
		childNode, err := build(childPath)
		if err != nil {
			return nil, err
		}
		nodo.Children = append(nodo.Children, childNode)
	}

	return nodo, nil
}

func Build(in string) *Node {
	arbol, err := build(config.GetPasswordStoreDir(in))
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return arbol
}
