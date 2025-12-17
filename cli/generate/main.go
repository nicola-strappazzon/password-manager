package generate

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nicola-strappazzon/pm/clipboard"

	"github.com/spf13/cobra"
)

var flagCopy bool
var flagLength int
var flagLowercase bool
var flagNumbers bool
var flagSymbols bool
var flagUppercase bool

const (
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	NUMBERS   = "0123456789"
	SPECIAL   = "!@#$%^&*()-_=+[]{}<>?,."
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a random password with configurable length, symbols, and numbers.",
		Run:   RunCommand,
	}

	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().BoolVarP(&flagLowercase, "no-lowercase", "l", false, "Exclude lowercase")
	cmd.Flags().BoolVarP(&flagNumbers, "no-numbers", "n", false, "Exclude numbers")
	cmd.Flags().BoolVarP(&flagSymbols, "no-symbols", "s", false, "Exclude symbols")
	cmd.Flags().BoolVarP(&flagUppercase, "no-uppercase", "u", false, "Exclude uppercase")
	cmd.Flags().IntVarP(&flagLength, "length", "m", 16, "Password max length, by default is 16 characters")

	return cmd
}

func RunCommand(cmd *cobra.Command, args []string) {
	var v string

	if flagLength < 4 || flagLength > 255 {
		fmt.Println("Invalid length. Must be between 4 and 255 characters.")
		return
	}

	v = GeneratePassword()

	if flagCopy {
		clipboard.Write(v)
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), v)
}

func GeneratePassword() string {
	var charset string

	p := make([]byte, flagLength)

	if !flagNumbers {
		charset += NUMBERS
	}
	if !flagLowercase {
		charset += LOWERCASE
	}
	if !flagUppercase {
		charset += UPPERCASE
	}
	if !flagSymbols {
		charset += SPECIAL
	}

	for i := range p {
		p[i] = charset[rand.Intn(len(charset))]
	}

	return string(p)
}
