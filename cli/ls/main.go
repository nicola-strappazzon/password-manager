package ls

import (
	"errors"
	"fmt"

	"github.com/nicola-strappazzon/password-manager/arguments"
	"github.com/nicola-strappazzon/password-manager/completion"
	"github.com/nicola-strappazzon/password-manager/explorer"
	"github.com/nicola-strappazzon/password-manager/path"

	"github.com/spf13/cobra"
)

func NewCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "ls",
		Short:             "List all encrypted items in tree format.",
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

		fmt.Fprint(cmd.OutOrStdout(), out)
	} else {
		return errors.New("No such file or directory.")
	}

	return nil
}
