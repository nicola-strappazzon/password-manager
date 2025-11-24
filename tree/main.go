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

func validFile(in fs.DirEntry) bool {
	if in.IsDir() {
		return false
	}

	return !validExtension(in.Name())
}

func validExtension(in string) bool {
	return strings.HasSuffix(in, ".gpg")
}

func removeExtension(in string) string {
	return strings.TrimSuffix(in, ".gpg")
}

func removeWorkDir(path string) (out string) {
	out = strings.TrimPrefix(path, config.GetWorkDirectory(""))
	out = strings.TrimPrefix(out, "/")

	return out
}

func ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func bold(in string) string {
	return fmt.Sprintf("\033[1;37m%s\033[0m", in)
}

func walk(path string) *Node {
	info, err := os.Stat(path)
	if err != nil {
		info, err = os.Stat(path + ".gpg")
		if err != nil {
			return &Node{}
		}
		path = path + ".gpg"
	}

	node := &Node{
		Name:  removeExtension(filepath.Base(path)),
		Path:  path,
		IsDir: info.IsDir(),
	}

	if !info.IsDir() {
		return node
	}

	files, _ := os.ReadDir(path)
	if err != nil {
		return &Node{}
	}

	for _, file := range files {
		if ignoreDir(file.Name()) {
			continue
		}

		if validFile(file) {
			continue
		}

		childPath := filepath.Join(path, file.Name())
		childNode := walk(childPath)
		node.Children = append(node.Children, childNode)
	}

	return node
}

func WalkFrom(in string) (out *Node) {
	return walk(config.GetWorkDirectory(in))
}

func Walk() (out *Node) {
	return WalkFrom("")
}
