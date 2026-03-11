package ls

import (
	"errors"

	"github.com/nicola-strappazzon/password-manager/internal/arguments"
	"github.com/nicola-strappazzon/password-manager/internal/completion"
	"github.com/nicola-strappazzon/password-manager/internal/explorer"
	"github.com/nicola-strappazzon/password-manager/internal/path"

	"github.com/spf13/cobra"
)

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "ls",
		Short:             "List encrypted items in tree format",
		RunE:              RunCommand,
		ValidArgsFunction: completion.SuggestDirectories,
	}

	return
}

func RunCommand(cmd *cobra.Command, args []string) error {
	var pathCard = arguments.First(args)
	var p path.Path = path.Path(pathCard)

	if p.Exists() {
		out, err := explorer.PrintTree(p.Absolute())
		if err != nil {
			return err
		}

		cmd.Print(out)
	} else {
		return errors.New("No such file or directory.")
	}

	return nil
}
