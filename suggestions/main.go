package suggestions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/pm/config"
)

func IsValid(in string) bool {
	if _, err := os.Stat(in); err == nil {
		return true
	}

	return false
}

func Suggestions(in string) (out []string) {
	// fmt.Println(config.GetPasswordStoreDir())
	// fmt.Printf(".%s.", in)

	// path := filepath.Join(config.GetPasswordStoreDir(), in)
	// fmt.Println(path)
	path := fmt.Sprintf("/Users/nicola/.password-store/%s/", in)
	// path := filepath.Join(config.GetPasswordStoreDir(), "/", in)
	if !IsValid(path) {
		path = config.GetPasswordStoreDir(in)
	}

	items, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, item := range items {
		name := item.Name()
		if len(name) > 0 && name[0] == '.' {
			continue
		}

		if item.IsDir() {
			out = append(out, name+"/")
		} else if filepath.Ext(name) == ".gpg" {
			out = append(out, name)
		}
	}

	return out
}
