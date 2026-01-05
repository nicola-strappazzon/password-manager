package ls

import (
	"fmt"

	"github.com/nicola-strappazzon/password-manager/arguments"
	"github.com/nicola-strappazzon/password-manager/completion"
	"github.com/nicola-strappazzon/password-manager/config"
	"github.com/nicola-strappazzon/password-manager/explorer"

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
	out, err := explorer.PrintTree(config.GetPath(arguments.First(args)))
	if err != nil {
		return err
	}

	fmt.Fprint(cmd.OutOrStdout(), out)

	return nil
}
