package openpgp

import (
	"bytes"
	"fmt"
	"os/exec"
)

func IsGPGInstalled() bool {
	cmd := exec.Command("gpg", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func Decrypt(path string) (string, error) {
	var outBuf bytes.Buffer

	cmd := exec.Command("gpg", "--decrypt", path)
	cmd.Stdout = &outBuf

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("Detalles:", outBuf.String())

    return "", nil
}
