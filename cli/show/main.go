package show

import (
	"fmt"
	"slices"

	"github.com/nicola-strappazzon/password-manager/arguments"
	"github.com/nicola-strappazzon/password-manager/card"
	"github.com/nicola-strappazzon/password-manager/clipboard"
	"github.com/nicola-strappazzon/password-manager/completion"
	"github.com/nicola-strappazzon/password-manager/explorer"
	"github.com/nicola-strappazzon/password-manager/openpgp"
	"github.com/nicola-strappazzon/password-manager/otp"
	"github.com/nicola-strappazzon/password-manager/path"
	"github.com/nicola-strappazzon/password-manager/qr"
	"github.com/nicola-strappazzon/password-manager/term"

	"github.com/spf13/cobra"
)

var flagAll bool
var flagCopy bool
var flagQR bool
var flagField string
var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "show path/to/encrypted [flags]",
		Short: "Decrypt and show selected data. By default, it shows the password.",
		Example: "  pm show <TAB>\n" +
			"  pm show aws\n" +
			"  pm show aws -p <passphrase>\n" +
			"  pm show aws -p <passphrase> -a\n" +
			"  pm show aws -p <passphrase> -c\n" +
			"  pm show aws -p <passphrase> -f otp -c\n" +
			"  pm show aws -p <passphrase> -f aws.access_key -c",
		PreRunE:           PreRun,
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all decrypted file")
	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().BoolVarP(&flagQR, "qr", "q", false, "Generate a QR code for the decrypted password")
	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Filter by field name...")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	cmd.MarkFlagsMutuallyExclusive("all", "field", "qr")
	cmd.MarkFlagsMutuallyExclusive("qr", "field")

	cmd.RegisterFlagCompletionFunc("field", completion.SuggestFields)

	return
}

func PreRun(cmd *cobra.Command, args []string) error {
	fields := (&card.Card{}).Fields()
	field, _ := cmd.Flags().GetString("field")

	if field == "" {
		return nil
	}

	if NotInSlice(field, fields) {
		return fmt.Errorf("Invalid field: %s", field)
	}

	return nil
}

func RunCommand(cmd *cobra.Command, args []string) {
	var value string
	var tmpCard card.Card
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.IsNotFile() {
		explorer.PrintTree(p.Absolute())
		return
	}

	tmpCard = card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		p.Full(),
	))

	if flagAll {
		value = tmpCard.ToString()
	}

	if flagField != "" {
		value = tmpCard.GetValue(flagField)
	}

	if flagField == "" && !flagAll {
		value = tmpCard.Password
	}

	if flagField == "password" {
		value = tmpCard.Password
	}

	if flagField == "otp" {
		value = otp.Get(tmpCard.OTP)
	}

	if flagCopy {
		clipboard.Write(value)
		fmt.Fprintf(
			cmd.OutOrStdout(),
			"Copied %s for %s to clipboard.\n",
			flagField,
			p.Path(),
		)
		return
	}

	if flagQR {
		qr.Generate(value)
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), value)
}

func NotInSlice(s string, list []string) bool {
	return !slices.Contains(list, s)
}
