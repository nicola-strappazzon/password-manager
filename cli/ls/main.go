package ls

import (
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
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectories,
	}

	return
}

func RunCommand(cmd *cobra.Command, args []string) {
	explorer.PrintTree(config.GetPath(arguments.First(args)))
}
