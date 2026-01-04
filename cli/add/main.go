package add

import (
	"fmt"
	"slices"

	"github.com/nicola-strappazzon/password-manager/arguments"
	"github.com/nicola-strappazzon/password-manager/card"
	"github.com/nicola-strappazzon/password-manager/completion"
	"github.com/nicola-strappazzon/password-manager/explorer"
	"github.com/nicola-strappazzon/password-manager/openpgp"
	"github.com/nicola-strappazzon/password-manager/path"
	"github.com/nicola-strappazzon/password-manager/term"

	"github.com/spf13/cobra"
)

var flagField string
var flagPassphrase string
var flagValue string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "add",
		Short: "Add or update an encrypted item.",
		Example: "  pm add <TAB>\n" +
			"  pm add aws -p <passphrase> -f password -v 12345\n",
		PreRunE:           PreRun,
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Field name to set or update...")
	cmd.Flags().StringVarP(&flagValue, "value", "v", "", "Value to assign to the field")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	cmd.RegisterFlagCompletionFunc("field", completion.SuggestFields)

	return
}

func PreRun(cmd *cobra.Command, args []string) error {
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	field, _ := cmd.Flags().GetString("field")
	value, _ := cmd.Flags().GetString("value")

	if pathCard == "" {
		return fmt.Errorf("Require to specify path.")
	}

	if p.IsInvalid() {
		return fmt.Errorf("Invalid path. Allowed characters: letters and numbers, not '.', '-' and '_'. The path must not end with '/'.")
	}

	if field == "" {
		return fmt.Errorf("Require to specify field name with --field <name>")
	}

	if value == "" {
		return fmt.Errorf("Require to specify value for field %s with --value <value>", field)
	}

	if NotInSlice(field) {
		return fmt.Errorf("Invalid field: %s", field)
	}

	if field != "" && value == "" {
		return fmt.Errorf("Require to specify value for field %s with --value <value>", field)
	}

	if field == "password" && value != "" {
		return fmt.Errorf("Invalid value: do not provide the password directly. Leave it empty and the tool will prompt for it securely.")
	}

	return nil
}

func RunCommand(cmd *cobra.Command, args []string) {
	var tmpCard = card.Card{}
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.IsNotFile() {
		explorer.PrintTree(p.Absolute())
		return
	}

	// fmt.Println(p.IsFile())

	if p.IsFile() {
		tmpCard = card.New(openpgp.Decrypt(
			term.ReadPassword("Passphrase: ", flagPassphrase),
			p.Full(),
		))
	}

	if flagField == "password" {
		tmpCard.Password = term.ReadPassword(fmt.Sprintf("Enter password for %s: ", p.Path()), "")
	} else {
		tmpCard.SetValue(flagField, flagValue)
	}

	tmpCard.Save()
}

func NotInSlice(s string) bool {
	return !slices.Contains((&card.Card{}).Fields(), s)
}
