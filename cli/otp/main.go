package otp

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/explorer"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/otp"
	"github.com/nicola-strappazzon/pm/path"
	"github.com/nicola-strappazzon/pm/term"

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

	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy OTP code to clipboard")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) {
	var otpValue string
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)
	var tmpCard card.Card

	if p.IsNotFile() {
		explorer.PrintTree(p.Absolute())
		return
	}

	tmpCard = card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		p.Full(),
	))

	if tmpCard.CheckOTP() {
		fmt.Fprintf(cmd.OutOrStdout(), "This card dont have OTP.\n")
		return
	}

	otpValue = otp.Get(tmpCard.OTP)

	if flagCopy {
		clipboard.Write(otpValue)
		fmt.Fprintf(cmd.OutOrStdout(), "Copied OTP code for %s to clipboard.\n", p.Path())
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), otpValue)
}
