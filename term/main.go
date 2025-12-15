package term

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(title, in string) string {
	if len(in) == 0 {
		fmt.Print(title)

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("\r\033[2K")

		return string(bytePassword)
	}

	return in
}
