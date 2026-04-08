package add

import (
	"fmt"
	"slices"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/decryptor"
	"github.com/nicola-strappazzon/password-manager/internal/git"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/spf13/cobra"
)

var flagField string
var flagPassphrase string
var flagValue string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "add",
		Short: "Add or update an encrypted item",
		Example: "  pm add <TAB>\n" +
			"  pm add aws -p <passphrase> -f password -v 12345\n",
		SilenceUsage:      true,
		PreRunE:           PreRun,
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Field name to set or update")
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
		return fmt.Errorf("A path is required.")
	}

	if p.IsInvalid() {
		return fmt.Errorf("Invalid path. Use only letters, numbers, '-', '_', and '/'. The path must not end with '/'.")
	}

	if field == "" {
		return fmt.Errorf("A field name is required.")
	}

	if NotInSlice(field) {
		return fmt.Errorf("Invalid field: %s", field)
	}

	if field == "password" && value != "" {
		cmd.PrintErrln("Warning: Using a password on the command line interface can be insecure.")
	}

	return nil
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var tmpCard = card.Card{}
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.IsFile() {
		var err error
		tmpCard, err = decryptor.Decrypt(flagPassphrase, p.Full())
		if err != nil {
			return err
		}

		if tmpCard.GetValue(flagField) != "" {
			return fmt.Errorf("Field '%s' already exists. Use 'pm update' to modify it.", flagField)
		}
	}

	if flagField == "password" {
		tmpCard.Password = term.ReadPassword("Enter new password: ", flagValue)
	} else {
		tmpCard.SetValue(flagField, flagValue)
	}

	tmpCard.Path = p.Full()

	if err := tmpCard.Save(); err != nil {
		return err
	}

	return git.Commit(fmt.Sprintf("Add %s for %s", flagField, pathCard))
}

func NotInSlice(s string) bool {
	return !slices.Contains((&card.Card{}).Fields(), s)
}
