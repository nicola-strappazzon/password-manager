package add

import (
	"fmt"

	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/card"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/file"
	"github.com/nicola-strappazzon/pm/openpgp"
	"github.com/nicola-strappazzon/pm/term"

	"github.com/spf13/cobra"
)

var flagField string
var flagPassphrase string
var flagValue string

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "add",
		Short: "Add or update an encrypted item.",
		Example: "  pm add <TAB>\n" +
			"  pm add aws -p <passphrase> -f password -v 12345\n",
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().StringVarP(&flagField, "field", "f", "", "Field name to set or update...")
	cmd.Flags().StringVarP(&flagValue, "value", "v", "", "Value to assign to the field")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG file")

	cmd.RegisterFlagCompletionFunc("field", completion.SuggestFields)

	return cmd
}

func RunCommand(cmd *cobra.Command, args []string) {
	var c = card.Card{}
	var p = arguments.First(args)

	if !file.Valid(p) {
		fmt.Println("Invalid path. Allowed characters: letters, numbers, '-' and '_'. The path must not end with '/'.")
		return
	}

	p = file.AbsolutePath(p)
	p = file.AddExt(p)

	if !file.IsDir(p) && file.Exist(p) {
		c = card.New(openpgp.Decrypt(
			term.ReadPassword("Passphrase: ", flagPassphrase),
			p,
		))
	}

	if flagField == "password" {
		c.Password = term.ReadPassword(fmt.Sprintf("Enter password for %s: ", file.Name(p)), "")
	} else {
		c.SetValue(flagField, flagValue)
	}

	file.Save(p, openpgp.Encrypt(c.ToString()))
}
