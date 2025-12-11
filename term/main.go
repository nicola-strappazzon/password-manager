package term

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(in string) string {
	if len(in) == 0 {
		fmt.Print("Passphrase: ")

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")

		return string(bytePassword)
	}

	return in
}
