package otp

import (
	"errors"
	"fmt"

	"github.com/nicola-strappazzon/password-manager/arguments"
	"github.com/nicola-strappazzon/password-manager/card"
	"github.com/nicola-strappazzon/password-manager/clipboard"
	"github.com/nicola-strappazzon/password-manager/completion"
	"github.com/nicola-strappazzon/password-manager/explorer"
	"github.com/nicola-strappazzon/password-manager/openpgp"
	"github.com/nicola-strappazzon/password-manager/otp"
	"github.com/nicola-strappazzon/password-manager/path"
	"github.com/nicola-strappazzon/password-manager/term"

	"github.com/spf13/cobra"
)

var flagCopy bool
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "otp path/to/encrypted [flags]",
		Short: "Generate an OTP code and optionally copy it to the clipboard.",
		Example: "  pm otp <TAB>\n" +
			"  pm otp aws\n" +
			"  pm otp aws -p <passphrase>\n" +
			"  pm otp aws -p <passphrase> -c\n",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy OTP code to clipboard")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var otpValue string
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)
	var tmpCard card.Card

	if p.ExistDirectory() {
		out, err := explorer.PrintTree(p.Absolute())
		if err != nil {
			return err
		}

		fmt.Fprint(cmd.OutOrStdout(), out)

		return nil
	}

	if !p.Exists() {
		return errors.New("No such file or directory.")
	}

	tmpCard = card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		p.Full(),
	))

	if tmpCard.CheckOTP() {
		fmt.Fprintf(cmd.OutOrStdout(), "The %s card does not have an OTP token.\n", p.Path())
		return nil
	}

	otpValue = otp.Get(tmpCard.OTP)

	if flagCopy {
		if err := clipboard.Write(otpValue); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Copied OTP code for %s to clipboard.\n", p.Path())
		return nil
	}

	fmt.Fprintln(cmd.OutOrStdout(), otpValue)

	return nil
}
