package ls

import (
	"github.com/nicola-strappazzon/pm/arguments"
	"github.com/nicola-strappazzon/pm/completion"
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "ls",
		Short:             "List all encrypted items in tree format.",
		Run:               RunCommand,
		ValidArgsFunction: completion.SuggestDirectories,
	}

	return cmd
}

func RunCommand(cmd *cobra.Command, args []string) {
	tree.WalkFrom(arguments.First(args)).Print()
}
