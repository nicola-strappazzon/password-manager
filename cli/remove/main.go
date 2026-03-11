package remove

import (
	"errors"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/config"
	"github.com/nicola-strappazzon/password-manager/internal/file"
	"github.com/nicola-strappazzon/password-manager/internal/path"
	"github.com/nicola-strappazzon/password-manager/internal/term"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "remove",
		Short:             "Remove an encrypted item",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectoriesAndFiles,
	}
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if !p.IsFile() {
		return errors.New("No such file or directory.")
	}

	if !term.Confirm("Are you sure you would like to delete " + pathCard + "?") {
		return nil
	}

	if err := file.Remove(p.Full()); err != nil {
		return err
	}

	file.RemoveEmptyParents(p.Full(), config.GetPath(""))

	return nil
}
