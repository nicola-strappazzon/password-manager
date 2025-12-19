package otp

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/otp"
	"github.com/nicola-strappazzon/pm/term"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

var flagCopy bool
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "otp path/to/encrypted [flags]",
		Short: "Generate an OTP code and optionally put it on the clipboard.",
		Example: "  pm otp <TAB>\n" +
			"  pm otp aws\n" +
			"  pm otp aws -p <passphrase>\n" +
			"  pm otp aws -p <passphrase> -c\n",
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) {
	var v string

	if len(args) == 0 {
		cmd.Help()
		return
	}

	if tree.WalkFrom(arguments.First(args)).IsDir {
		tree.WalkFrom(arguments.First(args)).Print()
		return
	}

	var b = openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		tree.WalkFrom(arguments.First(args)).Path,
	)

	var c = card.New(b)

	v = otp.Get(c.OTP)

	if flagCopy {
		clipboard.Write(v)
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), v)
}
