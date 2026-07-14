package show

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/clipboard"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/decryptor"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"
	"github.com/nicola-strappazzon/password-manager/internal/otp"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/qr"

	"github.com/spf13/cobra"
)

var flagAll bool
var flagCopy bool
var flagQR bool
var flagField string
var flagPassphrase string
var flagTime int

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "show path/to/encrypted [flags]",
		Short: "Decrypt and show stored data. By default, it shows the password",
		Example: "  pm show <TAB>\n" +
			"  pm show aws\n" +
			"  pm show aws -p <passphrase>\n" +
			"  pm show aws -p <passphrase> -a\n" +
			"  pm show aws -p <passphrase> -c\n" +
			"  pm show aws -p <passphrase> -t 5\n" +
			"  pm show aws -p <passphrase> -f otp -c\n" +
			"  pm show aws -p <passphrase> -f aws.access_key -c",
		PreRunE:           PreRun,
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all fields of the decrypted item")
	cmd.Flags().BoolVarP(&flagCopy, "copy", "c", false, "Copy decrypted password to clipboard")
	cmd.Flags().BoolVarP(&flagQR, "qr", "q", false, "Generate a QR code for the decrypted password")
	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Filter by field name...")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")
	cmd.Flags().IntVarP(&flagTime, "time", "t", 3, "Seconds to wait before showing or copying the OTP code")

	cmd.MarkFlagsMutuallyExclusive("all", "field", "qr")
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

func RunCommand(cmd *cobra.Command, args []string) error {
	var value string
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.ExistDirectory() {
		out, err := explorer.PrintTree(p.Absolute())
		if err != nil {
			return err
		}

		cmd.Print(out)

		return nil
	}

	if !p.Exists() {
		return errors.New("No such file or directory.")
	}

	if p.Size() == 0 {
		return errors.New("File is empty.")
	}

	tmpCard, err := decryptor.Decrypt(flagPassphrase, p.Full())
	if err != nil {
		return err
	}

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

	// When showing the password by default, also surface the OTP code (if the
	// card has one) after the configured delay.
	showsOTP := ShowsOTP(tmpCard)
	delay := time.Duration(flagTime) * time.Second

	if flagCopy {
		if err := clipboard.Write(value); err != nil {
			return err
		}

		field := flagField
		if field == "" {
			field = "password"
		}

		cmd.Printf("Copied %s for %s to clipboard.\n", field, p.Path())

		if showsOTP {
			cmd.PrintErrf("Waiting %s to copy the OTP code...\n", delay)
			time.Sleep(delay)

			if err := clipboard.Write(otp.Get(tmpCard.OTP)); err != nil {
				return err
			}

			cmd.Printf("Copied OTP code for %s to clipboard.\n", p.Path())
		}

		return nil
	}

	if flagQR {
		qr.Generate(value)
		return nil
	}

	cmd.Println(value)

	if showsOTP {
		cmd.PrintErrf("Waiting %s to show the OTP code...\n", delay)
		time.Sleep(delay)

		cmd.Println(otp.Get(tmpCard.OTP))
	}

	return nil
}

func NotInSlice(s string, list []string) bool {
	return !slices.Contains(list, s)
}

// ShowsOTP reports whether the OTP code should be shown alongside the password,
// i.e. the default password view is requested and the card has an OTP secret.
func ShowsOTP(c card.Card) bool {
	return flagField == "" && !flagAll && !c.CheckOTP()
}
