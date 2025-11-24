package ls

import (
	"github.com/nicola-strappazzon/pm/tree"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List all encrypted items in tree format.",
		Run: func(cmd *cobra.Command, args []string) {
			tree.WalkFrom(GetFirstArg(args)).Print()
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return tree.Walk().List(), cobra.ShellCompDirectiveNoFileComp
		},
	}

	return cmd
}

func GetFirstArg(in []string) string {
	if len(in) == 1 {
		return in[0]
	}

	return ""
}
