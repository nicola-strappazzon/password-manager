package edit

import (
	"bytes"
	"errors"
	"os"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/confluentinc/go-editor"
	"github.com/spf13/cobra"
)

var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "edit",
		Short:             "Edit an encrypted item.",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectories,
	}

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

	tmpCard = card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		p.Full(),
	))

	original := bytes.NewBufferString(tmpCard.ToString())
	edit := editor.NewEditor()
	edited, path, err := edit.LaunchTempFile("example", original)
	defer os.Remove(path)

	newCard := card.New(string(edited))
	newCard.Save()

	return err
}
