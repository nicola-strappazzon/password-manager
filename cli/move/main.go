package move

import (
	"errors"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/file"
	"github.com/nicola-strappazzon/password-manager/internal/openpgp"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/spf13/cobra"
)

var flagPassphrase string

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "move [source] [destination]",
		Short:             "Move an encrypted item",
		SilenceUsage:      true,
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}

	cmd.Flags().StringVarP(&flagPassphrase, "passphrase", "p", "", "Passphrase used to decrypt the GPG-encrypted file")

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("Source and destination paths are required.")
	}

	var src path.Path = path.Path(arguments.First(args))
	var dst path.Path = path.Path(args[1])

	if src.IsInvalid() {
		return errors.New("Source path is invalid.")
	}

	if dst.IsInvalid() {
		return errors.New("Destination path is invalid.")
	}

	if !src.IsFile() {
		return errors.New("No such file or directory.")
	}

	if dst.IsFile() {
		return errors.New("Destination already exists.")
	}

	tmpCard := card.New(openpgp.Decrypt(
		term.ReadPassword("Passphrase: ", flagPassphrase),
		src.Full(),
	))

	tmpCard.Path = dst.Full()
	tmpCard.Save()

	if err := file.Remove(src.Full()); err != nil {
		return err
	}

	file.RemoveEmptyParents(src.Full(), config.GetPath(""))

	return nil
}
