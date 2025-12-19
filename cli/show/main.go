package show

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/clipboard"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/config"
	"github.com/nicola-strappazzon/pm/explorer"
	"github.com/nicola-strappazzon/pm/file"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/otp"
	"github.com/nicola-strappazzon/pm/qr"
	"github.com/nicola-strappazzon/pm/term"

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
		Short: "Show and decrypt selected data. By default, it shows the password.",
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
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	cmd.MarkFlagsMutuallyExclusive("all", "field")
	cmd.MarkFlagsMutuallyExclusive("all", "qr")
	cmd.MarkFlagsMutuallyExclusive("qr", "field")
	cmd.MarkFlagsMutuallyExclusive("qr", "copy")

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
	var v string

	if len(args) == 0 {
		cmd.Help()
		return
	}

	if file.IsDir(config.GetDataDirectoryFrom(arguments.First(args))) {
		explorer.PrintTree(config.GetDataDirectoryFrom(arguments.First(args)))
		return
	}

	var b = openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		fmt.Sprintf("%s.gpg", config.GetDataDirectoryFrom(arguments.First(args))),
	)

	var c = card.New(b)

	if flagAll {
		v = b
	}

	if flagField != "" {
		v = c.GetValue(flagField)
	}

	if !flagAll && flagField == "" {
		v = c.Password
	}

	if flagField == "otp" {
		v = otp.Get(c.OTP)
	}

	if flagCopy {
		clipboard.Write(v)
		return
	}

	if flagQR {
		qr.Generate(v)
		return
	}

	fmt.Fprintln(cmd.OutOrStdout(), v)
}

func NotInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return false
		}
	}
	return true
}
