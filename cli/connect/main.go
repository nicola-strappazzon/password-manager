package connect

import (
	"errors"
	"os"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/connector"
	"github.com/nicola-strappazzon/password-manager/internal/decryptor"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"
	"github.com/nicola-strappazzon/password-manager/internal/path"

	"github.com/spf13/cobra"
)

var flagPassphrase string
var flagCmd bool

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "connect path/to/encrypted [flags]",
		Short: "Connect to the database stored in an item using its credentials",
		Example: "  pm connect <TAB>\n" +
			"  pm connect production/db\n" +
			"  pm connect production/db -p <passphrase>\n",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().BoolVar(&flagCmd, "cmd", false, "Print the database client command without executing it")
	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)
	var tmpCard card.Card

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

	if flagCmd {
		out, err := connector.StringForPrint(tmpCard)
		if err != nil {
			return err
		}

		cmd.Println(out)

		return nil
	}

	sql, err := connector.Command(tmpCard)
	if err != nil {
		return err
	}

	sql.Stdin = os.Stdin
	sql.Stdout = os.Stdout
	sql.Stderr = os.Stderr

	return sql.Run()
}
