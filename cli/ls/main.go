package ls

import (
	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/config"
	"github.com/nicola-strappazzon/pm/explorer"

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
