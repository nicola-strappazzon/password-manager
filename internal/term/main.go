package term

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func ReadLine(question string) string {
	fmt.Print(question + " ")
	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(answer)
}

func Confirm(question string) bool {
	fmt.Print(question + " [y/N]: ")
	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(strings.ToLower(answer)) == "y"
}
